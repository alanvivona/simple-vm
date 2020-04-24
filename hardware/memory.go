package hardware

import (
	"encoding/hex"
	"fmt"
)

type VirtualMemory []byte

func (m VirtualMemory) Print() {
	fmt.Printf("= Memory [%d bytes] [%d words] =\n", len(m), len(m)/8)
	fmt.Printf("%s\n", hex.Dump(m))
}
