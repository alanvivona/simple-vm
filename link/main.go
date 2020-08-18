package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"os"

	"github.com/alanvivona/simple-vm/link"
	"github.com/sirupsen/logrus"
)

func main() {
	var outFilePath = flag.String("o", "out.yz", "Output file path")
	var inFilePath = flag.String("i", "", "Input file path")
	var verboseMode = flag.Bool("v", false, "Verbose output")
	flag.Parse()

	if verboseMode != nil && *verboseMode {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	outfile, err := os.Create(*outFilePath)
	FAIL(err)
	logrus.Infof("Output file %s", *outFilePath)
	outfile.Close()

	if inFilePath == nil || len(*inFilePath) < 1 {
		FAIL(errors.New("Input file path missing"))
	}
	inBytes, err := ioutil.ReadFile(*inFilePath)
	FAIL(err)

	outBytes, err := link.Link(inBytes)
	FAIL(err)

	err = ioutil.WriteFile(*outFilePath, outBytes, 0400)
	FAIL(err)

	logrus.Infof("Wrote %d (0x%x) bytes to file %s", len(outBytes), len(outBytes), *outFilePath)
	logrus.Infof("OK")
}

func FAIL(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
