package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	golden "github.com/kiwamunet/image-golden"
)

func main() {
	res := run(os.Args)
	log.Println(res)
}

func run(args []string) bool {

	flagSet := flag.NewFlagSet("cmd line", flag.ExitOnError)
	var (
		resultFile = flagSet.String("result_flie", "", "list of result_flie")
		isByte     = flagSet.Bool("byte", false, "list of byte")
		ssim       = flagSet.Float64("ssim", 0, "list of ssim")
		psnr       = flagSet.Float64("psnr", 0, "list of psnr")
	)
	if err := flagSet.Parse(args[1:]); err != nil {
		log.Fatalf(fmt.Sprintf("System error: Termination Err:%v", err))
	}

	t := golden.Th{
		IsByteCheck: *isByte,
	}

	if *ssim != 0 {
		t.IsSsimCheck = true
		t.Ssim = *ssim
	}
	if *psnr != 0 {
		t.IsPsnrCheck = true
		t.Psnr = *psnr
	}

	o := &golden.Opt{
		ResultFile: *resultFile,
		Threshold:  t,
	}
	return golden.Golden(o)
}
