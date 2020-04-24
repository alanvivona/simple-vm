package main

import (
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

	InstructionSet := isa.Create()

	bytecode := []byte{
		0x00, 0x00, 0x00, //"start"
		0xf1, 0x00, 0x01,
		0x01, 0x01, 0x00,
		0xf1, 0x02, 0x02,
		0x01, 0x03, 0x02,
		0xff, 0x00, 0x00, //"end"
	}

	for true {
		pc := vh.CPU.Registers[vh.CPU.PCIndex]
		if uint(pc) >= uint(len(bytecode)-3) {
			logrus.Fatalf("No bytecode left. Early end of execution at pc=0x%02x", pc)
			break
		}

		opID := bytecode[pc]
		microcode, exists := InstructionSet[opID]
		if !exists {
			logrus.Fatalf("Unrecognized op identifier: 0x%x", opID)
			break
		}
		args := bytecode[pc+1 : pc+3]
		logrus.Infof("Executing bytecode: 0x%02x 0x%02x 0x%02x", opID, args[0], args[1])
		keepGoing, err := microcode.Exec(&vh, args)

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
