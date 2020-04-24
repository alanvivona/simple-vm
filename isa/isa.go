package isa

import (
	"fmt"

	"../hardware"
)

type Instructions map[string]Executable

type Executable interface {
	Exec(h *hardware.VirtualHardware, args []int) (bool, error)
}

type Microcode struct {
	Def   func(h *hardware.VirtualHardware, args []int) (bool, error)
	ArgsQ uint
}

func (m Microcode) Exec(h *hardware.VirtualHardware, args []int) (bool, error) {
	if uint(len(args)) != m.ArgsQ {
		return false, fmt.Errorf("Mismatch args quantity. Have: %d %+v, Want: %d", len(args), args, m.ArgsQ)
	}
	return m.Def(h, args)
}

func Create() Instructions {
	return Instructions{
		"start": start,
		"end":   end,
		//"set":   set,
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
	Def: func(h *hardware.VirtualHardware, args []int) (bool, error) {
		for k := range h.CPU.State {
			h.CPU.State[k] = 0
		}
		return true, nil
	},
}

var end = Microcode{
	Def: func(h *hardware.VirtualHardware, args []int) (bool, error) {
		return false, nil
	},
}

//func set(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func put(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func get(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func cmp(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func ifExecute(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func add(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func sub(h *hardware.VirtualHardware, args []int) (bool, error) {
//
//}
//
//func dec(h *hardware.VirtualHardware, args []int) (bool, error) {
////sub(cpu, mem, args[0], 1)
//}
