package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"text_mangler/manglers"
)

func getFlags() (inFilename, outFilename, manglerFlag string) {
	inPtr := flag.String("infile", "", "file to read from")
	outPtr := flag.String("outfile", "", "file to write to")
	manglerPtr := flag.String("mangler", "", "which text transformer to use")

	flag.Parse()
	return *inPtr, *outPtr, *manglerPtr
}

func getFileHandles(inFilename, outFilename string) (*os.File, *os.File, error) {
	var in, out *os.File
	var err error
	if inFilename == "" {
		in = os.Stdin
	} else {
		pth := path.Clean(inFilename)
		in, err = os.Open(pth)
		if err != nil {
			return nil, nil, err
		}
	}

	if outFilename == "" {
		out = os.Stdout
	} else {
		pth := path.Clean(outFilename)
		out, err = os.OpenFile(pth, os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return nil, nil, err
		}
	}
	return in, out, nil
}

func main() {

	inFilename, outFilename, manglerFlag := getFlags()

	in, out, err := getFileHandles(inFilename, outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	defer out.Close()


	mangler, ok := manglers.Registry[manglerFlag]
	if !ok {
		err = fmt.Errorf("mangler with name %s not found in registry", manglerFlag)
		log.Fatal(err)
	}

	output := mangler(in)
	_, err = out.Write(output)
	if err != nil {
		log.Fatal(err)
	}
}
