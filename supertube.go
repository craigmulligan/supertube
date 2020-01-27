package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func readData(pipeName string) error {
	resp, err := http.Get(fmt.Sprintf("https://patchbay.pub/%s", pipeName))
	io.Copy(os.Stdout, resp.Body)
	return err
}

func writeData(pipeName string) error {
	stdReader := bufio.NewReader(os.Stdin)
	_, err := http.Post(fmt.Sprintf("https://patchbay.pub/%s", pipeName), "plain/text", stdReader)
	return err
}

func usage() {
	fmt.Printf("Usage: %s [FLAGS] pipeName\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	showHelp := flag.Bool("h", false, "Print Usage.")
	readerMode := flag.Bool("r", false, "reader mode.")
	flag.Parse()

	if flag.NArg() < 1 || *showHelp {
		flag.Usage()
		return
	}

	pipeName := flag.Args()[0]

	if *readerMode {
		readData(pipeName)
	} else {
		writeData(pipeName)
	}
}
