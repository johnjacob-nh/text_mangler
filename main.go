package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text_mangler/manglers"
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
			log.Fatal(err)
		}
	}

	fmt.Println("bloop")
	if *outPtr == "" {
		out = os.Stdout
	} else {
		pth := path.Clean(*outPtr)
		out, err = os.Open(pth)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer in.Close()
	defer out.Close()

	mangler, ok := manglers.Registry[*manglerPtr]
	if !ok {
		err = fmt.Errorf("mangler with name %s not found in registry", *manglerPtr)
		log.Fatal(err)
	}

	output := mangler(in)
	_, err = out.Write(output)
	if err != nil {
		log.Fatal(err)
	}
}
