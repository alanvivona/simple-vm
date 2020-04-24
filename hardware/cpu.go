package hardware

import (
	"fmt"
)

type AbstractCPU map[uint]int

type ConcreteCPU struct {
	State AbstractCPU
	Arch  map[string]uint
	Model string
}

func (c ConcreteCPU) SetByIndex(what int, where uint) error {
	if where > uint(len(c.State)) {
		return fmt.Errorf("Unrecognized register index %d", where)
	}
	c.State[where] = what
	return nil
}

func (c ConcreteCPU) SetByName(what int, where string) error {
	index, exists := c.Arch[where]
	if !exists {
		return fmt.Errorf("Unrecognized register name %s", where)
	}
	c.State[index] = what
	return nil
}

func (c ConcreteCPU) GtByIndex(index uint) (int, error) {
	if index > uint(len(c.State)) {
		return 0, fmt.Errorf("Unrecognized register index %d", index)
	}
	return c.State[index], nil
}

func (c ConcreteCPU) GetByName(name string) (int, error) {
	index, exists := c.Arch[name]
	if !exists {
		return 0, fmt.Errorf("Unrecognized register name %s", name)
	}
	return c.State[index], nil
}

func (c ConcreteCPU) Print() {
	fmt.Printf("= CPU state =\n")
	for name, index := range c.Arch {
		fmt.Printf("%s: %d\n", name, c.State[index])
	}
}
