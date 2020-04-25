package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"os"

	"./asm"
	"github.com/sirupsen/logrus"
)

func main() {

	var outFilePath = flag.String("o", "out.bin", "Output file path")
	var verboseMode = flag.Bool("v", false, "Verbose output")
	flag.Parse()

	if verboseMode != nil && *verboseMode {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	f, err := os.Create(*outFilePath)
	FAIL(err)
	logrus.Infof("Writting output to %s", *outFilePath)
	defer f.Close()

	stdinscanner := bufio.NewScanner(os.Stdin)

	lineNumber := 0
	for stdinscanner.Scan() {
		text := stdinscanner.Text()
		lineNumber++
		bytecode, errs := asm.AsmLine(text)
		if errs != nil && len(errs) > 0 {
			for _, err := range errs {
				logrus.WithFields(logrus.Fields{"line": lineNumber}).Error(err)
			}
		}
		stringRep := make([]byte, hex.EncodedLen(len(bytecode)))
		hex.Encode(stringRep, bytecode)
		logrus.Infof("%d:\t%s\t# from `%s`\n", lineNumber, stringRep, text)

		_, err := f.Write(bytecode)
		FAIL(err)
	}

	err = stdinscanner.Err()
	FAIL(err)
}

func FAIL(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
