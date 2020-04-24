package isa

import (
	"../hardware"
)

type Executable func(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)

type Instructions map[string]Executable{
	"start": start,
	"end": end,
	"set": set,
	"put": put,
	"get": get,
	"cmp": cmp,
	"if": ifExecute,
	"dec": dec,
	"inc": inc,
	"add": add,
	"sub": sub,
	"call": call,
	"ret": ret,
}

func start(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	*cpu = hardware.AbstractCPU{
		RA, RB, RC, RD : 0,
		SP             : 0,
		PC             : 0,
	}
}

func end(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	cpu = nil
	mem = nil
}

func set(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func put(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func get(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func cmp(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func ifExecute(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func add(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func sub(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	
}

func dec(cpu *hardware.AbstractCPU, mem *hardware.MadeUpMemory, args []int64) (int64, error)  {
	sub(cpu, mem, args[0], 1)
}