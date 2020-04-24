package hardware

import (
	"fmt"
	"sort"
)

type AbstractCPU map[uint]int

type VirtualCPU struct {
	State AbstractCPU
	Arch  map[string]uint
	Model string
}

func (c VirtualCPU) Set(where uint, what int) error {
	if where > uint(len(c.State)) {
		return fmt.Errorf("Unrecognized register index %d", where)
	}
	c.State[where] = what
	return nil
}

func (c VirtualCPU) Get(index uint) (int, error) {
	if index > uint(len(c.State)) {
		return 0, fmt.Errorf("Unrecognized register index %d", index)
	}
	return c.State[index], nil
}

func (c VirtualCPU) Print() {
	fmt.Printf("= CPU state =\n")
	registreNames := make([]string, 0)
	for name, _ := range c.Arch {
		registreNames = append(registreNames, name)
	}
	sort.Strings(registreNames)
	for _, name := range registreNames {
		index := c.Arch[name]
		fmt.Printf("%12s: 0x%015x", name, c.State[index])
		if index%2 == 0 {
			fmt.Println()
		}
	}
}
