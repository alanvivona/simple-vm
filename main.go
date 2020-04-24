package main

import (
	"./hardware"
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
}
