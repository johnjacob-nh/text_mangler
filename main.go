package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"text_mangler/manglers/parse_location_csv"
)



func main() {

	inPtr := flag.String("infile", "", "file to read from")
	outPtr := flag.String("outfile", "", "file to write to")
	manglerPtr := flag.String("mangler", "", "which text transformer to use")

	flag.Parse()

	var in, out *os.File
	var err error
	var mangler func(*os.File) []byte
	if *inPtr == "" {
		in = os.Stdin
	} else {
		pth := path.Clean(*inPtr)
		in, err = os.Open(pth)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}

	if *outPtr == "" {
		out = os.Stdout
	} else {
		pth := path.Clean(*outPtr)
		out, err = os.Open(pth)
		if err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}
	defer in.Close()
	defer out.Close()

	switch *manglerPtr {
	case "p":
		fallthrough
	case "parse_location":
		mangler = parse_location_csv.Mangle
	}

	output := mangler(in)
	_, err = out.Write(output)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}
}
