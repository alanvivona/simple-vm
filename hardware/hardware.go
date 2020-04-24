package hardware

import (
	"errors"
	"fmt"
)

type VirtualHardware struct {
	CPU VirtualCPU
	Mem VirtualMemory
}

func Create(cpuModel string, memorySize uint) (VirtualHardware, error) {
	virtHard := VirtualHardware{}
	cpu, err := createCPU(cpuModel)
	if err != nil {
		return virtHard, err
	}
	virtHard.CPU = cpu

	mem, err := createMemory(memorySize)
	if err != nil {
		return virtHard, err
	}
	virtHard.Mem = mem

	return virtHard, nil
}

func createCPU(model string) (VirtualCPU, error) {
	newCPU := VirtualCPU{}
	newCPU.Model = model

	switch model {
	case "fakeone":
		registers := []string{"ra", "rb", "rc", "rd", "sp", "pc"}
		newCPU.State = make(AbstractCPU)
		newCPU.Arch = make(map[string]uint)
		for i, name := range registers {
			newCPU.Arch[name] = uint(i)
			newCPU.State[uint(i)] = 0
		}
	default:
		return newCPU, fmt.Errorf("Model %s not found", model)
	}

	return newCPU, nil
}

func createMemory(size uint) (VirtualMemory, error) {
	if size%8 != 0 {
		return nil, errors.New("Memory size shold be a multiple of 8")
	}
	return make(VirtualMemory, size), nil
}
