package hardware

import "fmt"

func CreateCPU(model string) (*ConcreteCPU, error) {
	newCPU := ConcreteCPU{}
	newCPU.Model = model

	switch model {
	case "fakeone":
		registers := []string{"ra", "rb", "rc", "rd", "sp", "pc"}
		newCPU.State = make(AbstractCPU)
		newCPU.Arch = make(map[string]uint)
		for i, name := range registers {
			newCPU.Arch[name] = uint(i)
			newCPU.State[uint(i)] = 0
		}
	default:
		return nil, fmt.Errorf("Model %s not found", model)
	}

	return &newCPU, nil
}
