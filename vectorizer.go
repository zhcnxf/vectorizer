package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
	"os"
)

var help bool

func init() {
	flag.BoolVar(&help, "h", false, "Print this help information")
}

var regions []Region

func main() {
	flag.Parse()
	if help || flag.NArg() > 1 {
		PrintHelp()
	} else {
		var reader io.Reader
		var err error
		if flag.NArg() > 0 {
			reader, err = os.Open(flag.Arg(0))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			reader = os.Stdin
		}
		img, _, err := image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
		Print(img)
	}
}

func PrintHelp() {
	fmt.Fprintf(os.Stderr, "Usage of %[1]s:\n\t%[1]s [option] <file>\n", os.Args[0])
	flag.PrintDefaults()
}
