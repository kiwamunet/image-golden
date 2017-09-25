# image-golden


**image-golden** is Specify with url and download the online converted image.  
Test whether the downloaded image is the same as the image downloaded last time.  
Those used for testing are compared using binary comparison, psnr comparison, ssim comparison.

## Installation:

```git clone https://github.com/kiwamunet/image-golden.git```

## Features:

- Prevent degeneration
- Parameters can be defined in toml.
- You can compare ssim, psnr

## Bugs:

## Use:

```go run main.go -result_flie result/hoge.txt```

## Result:

```
++ golden test starting .....

++ prepare test .....

++ config checking .....
ServerDomain is :http://example.io

++ image Downloading .....
downloading. url is :http://example.io/sample.jpg
downloading. url is :http://example.io/sample.jpg?width=100
downloading. url is :http://example.io/sample.jpg?height=100
downloading. url is :http://example.io/sample.jpg?format=png
downloading. url is :http://example.io/sample.jpg?format=webp
downloading. url is :http://example.io/sample.png?format=jpg
downloading. url is :http://example.io/sample.jpg?grayscale=90

++ Conpare Image Result .....
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.jpg 
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.jpg?width=100 
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.jpg?height=100 
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.jpg?format=png 
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.jpg?format=webp 
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.png?format=jpg 
Compare result. 
	BYTE is : true 
	SSIM is : 1.000000 
	PSNR is : 0.000000 
	FILE is : sample.jpg?grayscale=90 



Comprete Test.

Fin.
```
