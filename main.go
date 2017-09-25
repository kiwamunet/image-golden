package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/kiwamunet/image-golden/imagediff"
)

const (
	newDir     = "output/new"
	oldDir     = "output/old"
	configName = "config.toml"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	ServerDomain string  `toml:"serverDomain"`
	Param        []Param `toml:"param"`
}

type Param struct {
	File  string `toml:"file"`
	Query string `toml:"query"`
}

var (
	resultFile = flag.String("result_flie", "", "list of result_flie")
)

func init() {
	flag.Parse()
}
func main() {
	run(os.Args)
}

func run(args []string) {

	log.Println(*resultFile)

	// result file
	f, err := createFile()
	defer f.Close()
	if err != nil {
		log.Fatalf("System error: Termination Err:%v", err)
	}

	Log(f, "\n++ golden test starting .....")
	Log(f, "\n++ prepare test .....")
	if ok := prepareOutputDir(f); !ok {
		Log(f, fmt.Sprintf("System error: Termination Err:%v", err))
		os.Exit(0)
	}

	// config
	Log(f, "\n++ config checking .....")
	c := getConfig(configName, f)

	Log(f, "\n++ image Downloading .....")

	files := []string{}
	for _, v := range c.Server.Param {
		fileName := v.File + v.Query
		url := c.Server.ServerDomain + "/" + fileName
		Log(f, fmt.Sprintf("downloading. url is :%s", url))
		err := downLoadImages(url, fileName)
		if err != nil {
			Log(f, fmt.Sprintf("System error: Err:%v", err))
		}
		files = append(files, fileName)
	}

	Log(f, "\n++ Conpare Image Result .....")
	for _, file := range files {
		byteComp, ssim, psnr, err := compareImage(file)
		if err != nil {
			Log(f, fmt.Sprintf("System error: Err:%v", err))
		}
		Log(f, fmt.Sprintf("Compare result. \n\tBYTE is : %t \n\tSSIM is : %f \n\tPSNR is : %f \n\tFILE is : %s ", byteComp, ssim, psnr, file))
	}

	Log(f, "\n\n\nComprete Test.")
	Log(f, "\nFin.")
	log.Println("\n\n===== Test is Completed!! Please look at [ result/result-XXXX.txt ] =====\n")
	return
}

func getConfig(configName string, f *os.File) Config {
	var config Config
	_, err := toml.DecodeFile(configName, &config)
	if err != nil {
		Log(f, fmt.Sprintf("System error: Termination Err:%v", err))
		os.Exit(0)
	}

	Log(f, fmt.Sprintf("ServerDomain is :%s", config.Server.ServerDomain))
	return config
}

func createFile() (*os.File, error) {
	var fileName string
	if *resultFile != "" {
		fileName = *resultFile
	} else {
		fileName = "result/result-" + time.Now().Format("2006-01-02-15:04:05") + ".txt"
	}

	return os.Create(fileName)
}

func Log(f *os.File, str string) {
	_, err := f.WriteString(str + "\n")
	if err != nil {
		Log(f, fmt.Sprintf("System error: Termination Err:%v", err))
		os.Exit(0)
	}
}

func prepareOutputDir(f *os.File) bool {
	if err := os.RemoveAll(oldDir); err != nil {
		Log(f, fmt.Sprintf("%v", err))
	}
	if err := os.Rename(newDir, oldDir); err != nil {
		Log(f, fmt.Sprintf("%v", err))
	}
	if err := os.MkdirAll(newDir, 0777); err != nil {
		Log(f, fmt.Sprintf("%v", err))
	}
	if !checkOutputDir() {
		return false
	}
	return true
}

func checkOutputDir() bool {
	_, err := os.Stat(newDir)
	if err != nil {
		return false
	}
	_, err = os.Stat(oldDir)
	if err != nil {
		return false
	}
	return true
}

func compareImage(fileName string) (bool, float64, float64, error) {

	newData, err := ioutil.ReadFile(newDir + "/" + fileName)
	if err != nil {
		return false, 0, 0, err
	}
	oldData, err := ioutil.ReadFile(oldDir + "/" + fileName)
	if err != nil {
		return false, 0, 0, err
	}

	ssim, psnr, err := imagediff.ImageComp(newDir+"/"+fileName, oldDir+"/"+fileName)
	if err != nil {
		return false, 0, 0, err
	}

	return bytes.Equal(newData, oldData), ssim, psnr, nil
}

func downLoadImages(url string, fileName string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	file, err := os.Create(newDir + "/" + fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, response.Body)
	return nil
}
