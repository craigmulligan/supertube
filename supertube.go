package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/docker/docker/pkg/namesgenerator"
	"io"
	"log"
	"net/http"
	"os"
)

func getUrl(pipeName string) string {
	return fmt.Sprintf("https://patchbay.pub/%s", pipeName)
}

func readData(pipeName string) error {
	url := getUrl(pipeName)
	resp, err := http.Get(url)
	io.Copy(os.Stdout, resp.Body)
	return err
}

func writeData(pipeName string) error {
	url := getUrl(pipeName)
	stdReader := bufio.NewReader(os.Stdin)

	fmt.Fprintf(os.Stderr, "To read from this supertube on another machine use the following:\n\t > supertube -r %s", pipeName)
	_, err := http.Post(url, "plain/text", stdReader)

	return err
}

func usage() {
	fmt.Printf("Usage: %s [FLAGS] [pipeName]\n\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	showHelp := flag.Bool("h", false, "Print Usage.")
	readerMode := flag.Bool("r", false, "Reader mode, the  default is Writer mode.")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}

	var pipeName string
	var err error

	if flag.NArg() < 1 {
		if *readerMode {
			fmt.Println("A pipeName is required in ReaderMode\n")
			flag.Usage()
			os.Exit(1)
		}
		pipeName = namesgenerator.GetRandomName(0)
	} else {
		pipeName = flag.Args()[0]
	}

	if *readerMode {
		err = readData(pipeName)
	} else {
		err = writeData(pipeName)
	}

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
