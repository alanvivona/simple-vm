package hardware

import (
	"fmt"
	"sort"
)

type VirtualCPU struct {
	Arch      map[string]uint
	Registers map[uint]byte
	Model     string
	SPIndex   uint
	PCIndex   uint
}

func (c VirtualCPU) Set(registerIndex uint, value byte) error {
	if registerIndex > uint(len(c.Registers)) {
		return fmt.Errorf("Unrecognized register index 0x%x", registerIndex)
	}
	c.Registers[registerIndex] = value
	return nil
}

func (c VirtualCPU) Get(index uint) (byte, error) {
	if index > uint(len(c.Registers)) {
		return 0, fmt.Errorf("Unrecognized register index 0x%x", index)
	}
	return c.Registers[index], nil
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
		fmt.Printf("%12s: 0x%02x", name, c.Registers[index])
		if index%2 == 0 {
			fmt.Println()
		}
	}
}
