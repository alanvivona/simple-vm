package main

import (
	"fmt"
	"os"
	"os/exec"

	"./hardware"
	"./isa"
	"github.com/sirupsen/logrus"
)

func main() {
	clearScreen()

	vh, err := hardware.Create("fakeone", 48)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Infof("Created Hardware")
	logrus.Infof("CPU: %s", vh.CPU.Model)
	logrus.Infof("Memory: %db", len(vh.Mem))

	InstructionSet := isa.Create()

	bytecode := []byte{
		0x00, 0x00, 0x00, // start
		0xf1, 0x00, 0x01, // set ra 0x01
		0x01, 0x01, 0x00, // set rb ra
		0xf1, 0x02, 0x02, // set rc 0x02
		0x01, 0x03, 0x02, // set rd rc

		0x02, 0x00, 0x00, // put ra
		0xf2, 0x44, 0x00, // put 0x44
		0x03, 0x00, 0x00, // get ra

		0xf6, 0x03, 0x02, // add rd 0x44
		0x06, 0x03, 0x02, // add rd ra
		0xf7, 0x03, 0x02, // sub rd 0x44
		0x07, 0x03, 0x02, // sub rd ra

		0x08, 0x00, 0x00, // dec ra
		0x09, 0x00, 0x00, // inc ra

		0x0a, 0x00, 0x02, // not ra
		0x0b, 0x00, 0x02, // neg ra
		0x0c, 0x00, 0x01, // and ra rb
		0xfc, 0x00, 0x55, // and ra 0x55
		0x0d, 0x00, 0x01, // or  ra rb
		0xfd, 0x00, 0x55, // or  ra 0x55
		0x0e, 0x00, 0x01, // xor ra rb
		0xfe, 0x00, 0x55, // xor ra 0x55

		0xff, 0x00, 0x00, // end
	}

	for true {
		pc := vh.CPU.Registers[vh.CPU.PCIndex]
		if uint(pc) >= uint(len(bytecode)-isa.InstructionSize) {
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

		// Wait for user input to execute next instruction
		fmt.Scanln()
		clearScreen()
	}
	logrus.Infof("Code execution finished")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
