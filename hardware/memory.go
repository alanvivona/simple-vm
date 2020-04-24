package hardware

import (
	"encoding/hex"
	"fmt"
)

type VirtualMemory []byte

func (m VirtualMemory) check(index int) error {
	if index >= len(m) || index < 0 {
		return fmt.Errorf("Invalid memory addres 0x%x", index)
	}
	return nil
}

func (m VirtualMemory) Get(index int) (byte, error) {
	if err := m.check(index); err != nil {
		return 0x00, err
	}
	return m[index], nil
}

func (m VirtualMemory) Set(index int, val byte) error {
	err := m.check(index)
	if err == nil {
		m[index] = val
	}
	return err
}

func (m VirtualMemory) Print() {
	fmt.Printf("= Memory [%d bytes] [%d words] =\n", len(m), len(m)/8)
	fmt.Printf("%s\n", hex.Dump(m))
}
