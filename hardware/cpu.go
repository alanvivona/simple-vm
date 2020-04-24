package hardware

import (
	"fmt"
)

type AbstractCPU map[uint]int

type VirtualCPU struct {
	State AbstractCPU
	Arch  map[string]uint
	Model string
}

func (c VirtualCPU) SetByIndex(what int, where uint) error {
	if where > uint(len(c.State)) {
		return fmt.Errorf("Unrecognized register index %d", where)
	}
	c.State[where] = what
	return nil
}

func (c VirtualCPU) SetByName(what int, where string) error {
	index, exists := c.Arch[where]
	if !exists {
		return fmt.Errorf("Unrecognized register name %s", where)
	}
	c.State[index] = what
	return nil
}

func (c VirtualCPU) GtByIndex(index uint) (int, error) {
	if index > uint(len(c.State)) {
		return 0, fmt.Errorf("Unrecognized register index %d", index)
	}
	return c.State[index], nil
}

func (c VirtualCPU) GetByName(name string) (int, error) {
	index, exists := c.Arch[name]
	if !exists {
		return 0, fmt.Errorf("Unrecognized register name %s", name)
	}
	return c.State[index], nil
}

func (c VirtualCPU) Print() {
	fmt.Printf("= CPU state =\n")
	for name, index := range c.Arch {
		fmt.Printf("%s: 0x%08x\n", name, c.State[index])
	}
}
