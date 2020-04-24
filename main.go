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
	logrus.Infof("Memory: %d", len(vh.Mem))

	instruction := isa.Create()

	code := []string{
		"start",
		"set ra 1",
		"set 1 2",
		"set rc 0x3",
		"set 3 0x4",
		"set pc 0xfffffffffffffff",
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
			if strings.HasPrefix(lineSlice, "0x") {
				lineSlice = lineSlice[2:]
			}
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

		vh.CPU.Print()
		vh.Mem.Print()

		if err != nil {
			logrus.Error(err)
			break
		}
		if !keepGoing {
			break
		}
	}

	logrus.Infof("Code execution finished")
}
