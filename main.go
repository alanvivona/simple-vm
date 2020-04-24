package main

import (
	"strconv"
	"strings"

	"./hardware"
	"./isa"
	"github.com/sirupsen/logrus"
)

func main() {

	vh, err := hardware.Create("fakeone", 48)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Infof("Created Hardware")
	logrus.Infof("CPU: %s", vh.CPU.Model)
	vh.CPU.Print()
	logrus.Infof("Memory: %d", len(vh.Mem))
	vh.Mem.Print()

	instruction := isa.Create()

	code := []string{
		"start",
		"start 1",
		"end",
	}

	for lineNumber, line := range code {
		line := strings.Split(strings.ToLower(line), " ")
		logrus.Infof("Executing line %d: %+v", lineNumber, line)

		instruction, exists := instruction[line[0]]
		if !exists {
			logrus.Fatalf("Mnemonic '%s' not found for line %d", line[0], lineNumber)
		}

		args := []int{}
		for _, lineSlice := range line[1:] {
			ImmediateValue, err := strconv.ParseInt(lineSlice, 16, 64)
			if err == nil {
				args = append(args, int(ImmediateValue))
				continue
			}

			registerIndex, exists := vh.CPU.Arch[lineSlice]
			if !exists {
				logrus.Fatalf("Unrecognized value %s on line %d", lineSlice, lineNumber)
				return
			}
			args = append(args, int(registerIndex))
		}

		keepGoing, err := instruction.Exec(&vh, args)
		if err != nil {
			logrus.Error(err)
		}
		if !keepGoing {
			break
		}
	}

	logrus.Infof("Code execution finished")
}
