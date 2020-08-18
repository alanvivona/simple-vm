package isa

import (
	"errors"

	"github.com/alanvivona/simple-vm/hardware"
)

type Instructions map[byte]Executable

type Executable interface {
	Exec(h *hardware.VirtualHardware, args []byte) error
}

type Microcode struct {
	Def        func(h *hardware.VirtualHardware, args []byte) error
	IsMemWrite bool
	IsExecFlow bool
}

func (m Microcode) Exec(h *hardware.VirtualHardware, args []byte) error {
	if m.IsMemWrite {
		sp := h.CPU.Registers[h.CPU.SPIndex]
		if uint(sp) >= uint(len(h.Mem)-8) {
			return errors.New("Out of memory")
		}
	}
	err := m.Def(h, args)
	if !m.IsExecFlow {
		h.CPU.Registers[h.CPU.PCIndex] += InstructionSize
	}
	return err
}

const InstructionSize = 3

func Create() Instructions {
	return Instructions{
		// Setup
		0x00: start,
		0x0f: end,

		// Set registers
		0x10: setr,
		0x11: setv,

		// Stack
		0x20: putr,
		0x21: putv,
		0x22: get,

		// Arithmetic
		0x30: addr,
		0x31: addv,
		0x32: subr,
		0x33: subv,
		0x34: dec,
		0x35: inc,

		// Logic
		0x40: not,
		0x41: neg,
		0x42: andr,
		0x43: andv,
		0x44: orr,
		0x45: orv,
		0x46: xorr,
		0x47: xorv,

		// Program Flow (NOT IMPLEMENTED)
		//0x50: jmp,
		//0x51: jer,
		//0x52: jev,
		//0x53: jner,
		//0x54: jnev,
		//0x55: call,
		//0x56: ret,
	}
}

var jmp = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		return h.CPU.Set(h.CPU.PCIndex, bytecode[1])
	},
	IsMemWrite: false,
	IsExecFlow: true,
}

//var jev = Microcode{
//Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
//v, err := h.CPU.Get(uint(bytecode[0]))
//if err != nil {
//return err
//}
//if v == bytecode[1] {
//return h.CPU.Set(h.CPU.PCIndex, bytecode[1])
//}
//},
//IsMemWrite: false,
//IsExecFlow: true,
//}

var not = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v, err := h.CPU.Get(uint(bytecode[0]))
		if err != nil {
			return err
		}
		err = h.CPU.Set(uint(bytecode[0]), ^v)
		if err != nil {
			return err
		}
		return nil
	},
	IsMemWrite: false,
}

var neg = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v, err := h.CPU.Get(uint(bytecode[0]))
		if err != nil {
			return err
		}
		err = h.CPU.Set(uint(bytecode[0]), (^v)+1)
		if err != nil {
			return err
		}
		return nil
	},
	IsMemWrite: false,
}

var orr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v2, err := h.CPU.Get(uint(bytecode[1]))
		if err != nil {
			return err
		}
		return orv.Def(h, []byte{bytecode[0], v2})
	},
	IsMemWrite: false,
}

var orv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v1, err := h.CPU.Get(uint(bytecode[0]))
		if err != nil {
			return err
		}

		err = h.CPU.Set(uint(bytecode[0]), v1|bytecode[1])
		if err != nil {
			return err
		}
		return nil
	},
	IsMemWrite: false,
}

var andv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v1, err := h.CPU.Get(uint(bytecode[0]))
		if err != nil {
			return err
		}

		err = h.CPU.Set(uint(bytecode[0]), v1&bytecode[1])
		if err != nil {
			return err
		}
		return nil
	},
	IsMemWrite: false,
}

var andr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v2, err := h.CPU.Get(uint(bytecode[1]))
		if err != nil {
			return err
		}
		return andv.Def(h, []byte{bytecode[0], v2})
	},
	IsMemWrite: false,
}

var xorv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v1, err := h.CPU.Get(uint(bytecode[0]))
		if err != nil {
			return err
		}

		err = h.CPU.Set(uint(bytecode[0]), v1^bytecode[1])
		if err != nil {
			return err
		}
		return nil
	},
	IsMemWrite: false,
}

var xorr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v2, err := h.CPU.Get(uint(bytecode[1]))
		if err != nil {
			return err
		}
		return xorv.Def(h, []byte{bytecode[0], v2})
	},
	IsMemWrite: false,
}

var start = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		for k := range h.CPU.Registers {
			h.CPU.Registers[k] = 0
		}
		return nil
	},
	IsMemWrite: false,
}

var end = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		return errors.New("End Execution")
	},
	IsMemWrite: false,
}

var setr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v := h.CPU.Registers[uint(bytecode[1])]
		return h.CPU.Set(uint(bytecode[0]), v)
	},
	IsMemWrite: false,
}

var setv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		return h.CPU.Set(uint(bytecode[0]), bytecode[1])
	},
	IsMemWrite: false,
}

var putr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v := h.CPU.Registers[uint(bytecode[0])]
		return putv.Def(h, []byte{v})
	},
	IsMemWrite: true,
}

var putv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		sp := h.CPU.Registers[h.CPU.SPIndex]
		err := h.Mem.Set(int(sp), bytecode[0])
		if err != nil {
			return err
		}
		err = h.CPU.Set(h.CPU.SPIndex, sp+1)
		if err != nil {
			return err
		}
		return nil
	},
	IsMemWrite: true,
}

var get = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		sp, err := h.CPU.Get(h.CPU.SPIndex)
		if err != nil {
			return err
		}
		v, err := h.Mem.Get(int(sp) - 1)
		if err != nil {
			return err
		}
		err = h.CPU.Set(uint(bytecode[0]), v)
		if err != nil {
			return err
		}
		h.CPU.Set(h.CPU.SPIndex, sp-1)
		return nil
	},
	IsMemWrite: false,
}

var addr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v, err := h.CPU.Get(uint(bytecode[1]))
		if err != nil {
			return err
		}
		addv.Def(h, []byte{bytecode[0], v})
		return nil
	},
	IsMemWrite: false,
}

var addv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v, err := h.CPU.Get(uint(bytecode[0]))
		if err != nil {
			return err
		}
		h.CPU.Set(uint(bytecode[0]), v+bytecode[1])
		return nil
	},
	IsMemWrite: false,
}

var subv = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		comp := (^bytecode[1]) + 0x01
		return addv.Def(h, []byte{bytecode[0], comp})
	},
	IsMemWrite: false,
}

var subr = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		v, err := h.CPU.Get(uint(bytecode[1]))
		if err != nil {
			return err
		}
		return subv.Def(h, []byte{bytecode[0], v})
	},
	IsMemWrite: false,
}

var inc = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		return addv.Def(h, []byte{bytecode[0], 0x01})
	},
	IsMemWrite: false,
}

var dec = Microcode{
	Def: func(h *hardware.VirtualHardware, bytecode []byte) error {
		return subv.Def(h, []byte{bytecode[0], 0x01})
	},
	IsMemWrite: false,
}
