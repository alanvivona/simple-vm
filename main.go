package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/alanvivona/simple-vm/hardware"

	"github.com/alanvivona/simple-vm/isa"
	"github.com/alanvivona/simple-vm/link/link"

	"github.com/sirupsen/logrus"
)

func main() {
	var filePath = flag.String("i", "", "Input file path")
	var verboseMode = flag.Bool("v", false, "Verbose output")
	var manualMode = flag.Bool("m", false, "Manual execution")
	flag.Parse()

	if verboseMode != nil && *verboseMode {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

	if filePath == nil || len(*filePath) < 1 {
		logrus.Fatalf("Input file path missing")
	}
	inBytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		logrus.Error(err)
		logrus.Fatalf("Can't read file %s", *filePath)
	}

	logrus.Info("Checking file integrity and extracting bytecode")
	bytecode, err := link.ExtractExecutable(inBytes)
	if err != nil {
		logrus.Error(err)
		logrus.Fatalf("Can't extract bytecode from file %s", *filePath)
	}
	logrus.Infof("Loaded %db from code section", len(bytecode))

	vh, err := hardware.Create("fakeone", 48)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Infof("Initialized Hardware")
	logrus.Infof("CPU: %s", vh.CPU.Model)
	logrus.Infof("Memory: %db", len(vh.Mem))

	InstructionSet := isa.Create()

	logrus.Warn("Press a key to start execution")
	// Wait for user input to execute next instruction
	fmt.Scanln()

	for true {
		clearScreen()
		pc := vh.CPU.Registers[vh.CPU.PCIndex]
		if uint(pc) >= uint(len(bytecode)) {
			logrus.Fatalf("No bytecode left. End of execution at pc=0x%02x", pc)
			break
		}

		opID := bytecode[pc]
		microcode, exists := InstructionSet[opID]
		if !exists {
			logrus.Fatalf("Unrecognized op identifier: 0x%x", opID)
			break
		}
		args := bytecode[pc+1 : pc+isa.InstructionSize]
		logrus.Infof("Executing bytecode: 0x%02x 0x%02x 0x%02x", opID, args[0], args[1])
		err := microcode.Exec(&vh, args)

		vh.CPU.Print()
		vh.Mem.Print()

		if err != nil {
			logrus.Error(err)
			break
		}

		if manualMode != nil && (*manualMode) == true {
			logrus.Warn("Press a key to continue execution")
			// Wait for user input to execute next instruction
			fmt.Scanln()
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
	logrus.Infof("Code execution finished")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
