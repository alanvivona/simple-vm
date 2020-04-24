package isa

import (
	"errors"
	"strconv"

	"../hardware"
)

type Instructions map[byte]Executable

type Executable interface {
	Exec(h *hardware.VirtualHardware, args []byte) (bool, error)
}

type Microcode struct {
	Def        func(h *hardware.VirtualHardware, args []byte) (bool, error)
	ArgsQ      uint
	IsMemWrite bool
}

func (m Microcode) Exec(h *hardware.VirtualHardware, args []byte) (bool, error) {
	if m.IsMemWrite {
		sp := h.CPU.Registers[h.CPU.SPIndex]
		if uint(sp) >= uint(len(h.Mem)-8) {
			return false, errors.New("Out of memory")
		}
	}
	keepGoing, err := m.Def(h, args)
	h.CPU.Registers[h.CPU.PCIndex] += 3
	return keepGoing, err
}

func Create() Instructions {
	return Instructions{
		0x00: start,
		0xff: end,
		//0x01: setr,
		0xf1: setv,
		//"put":   put,
		//"get":   get,
		//"cmp":   cmp,
		//"if":    ifExecute,
		//"dec":   dec,
		//"inc":   inc,
		//"add":   add,
		//"sub":   sub,
		//"call":  call,
		//"ret":   ret,
	}
}

var start = Microcode{
	Def: func(h *hardware.VirtualHardware, args []byte) (bool, error) {
		for k := range h.CPU.Registers {
			h.CPU.Registers[k] = 0
		}
		return true, nil
	},
	ArgsQ:      0,
	IsMemWrite: false,
}

var end = Microcode{
	Def: func(h *hardware.VirtualHardware, args []byte) (bool, error) {
		return false, nil
	},
	ArgsQ:      0,
	IsMemWrite: false,
}

var setv = Microcode{
	Def: func(h *hardware.VirtualHardware, args []byte) (bool, error) {
		return true, h.CPU.Set(uint(args[0]), args[1])
	},
	ArgsQ:      2,
	IsMemWrite: false,
}

var putv = Microcode{
	Def: func(h *hardware.VirtualHardware, args []byte) (bool, error) {
		sp := h.CPU.Registers[h.CPU.SPIndex]
		sp++
		for i, v := range []byte(strconv.FormatInt(int64(args[0]), 16)) {
			h.Mem[uint(sp)+uint(i)] = v
		}
		return true, nil
	},
	ArgsQ:      2,
	IsMemWrite: true,
}

//
//func get(h *hardware.VirtualHardware, args []byte) (bool, error) {
//
//}
//
//func cmp(h *hardware.VirtualHardware, args []byte) (bool, error) {
//
//}
//
//func ifExecute(h *hardware.VirtualHardware, args []byte) (bool, error) {
//
//}
//
//func add(h *hardware.VirtualHardware, args []byte) (bool, error) {
//
//}
//
//func sub(h *hardware.VirtualHardware, args []byte) (bool, error) {
//
//}
//
//func dec(h *hardware.VirtualHardware, args []byte) (bool, error) {
////sub(cpu, mem, args[0], 1)
//}
