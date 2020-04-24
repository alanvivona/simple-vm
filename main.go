package main

import (
	"./hardware"
	"github.com/sirupsen/logrus"
)

func main() {

	cpu, err := hardware.CreateCPU("fakeone")
	if err != nil {
		logrus.Fatal(err)
		return
	}

	logrus.Infof("Created CPU: %s", cpu.Model)
	cpu.Print()
}
