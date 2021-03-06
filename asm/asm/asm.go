package asm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const TYPE_NUM = 0
const TYPE_REG = 1

type instructionDef struct {
	members []int
	opByte  byte
}

func (idef instructionDef) parse(textSlice []string) ([]byte, error) {
	bytecode := []byte{idef.opByte, 0x00, 0x00}

	expectedFieldsQ := len(textSlice) - 1
	givenFieldsQ := len(idef.members)
	if expectedFieldsQ != givenFieldsQ {
		return bytecode, fmt.Errorf("Wrong number of fields on instruction '%+v'. Got &d. Expected %d", textSlice, givenFieldsQ, expectedFieldsQ)
	}

	for i, mtype := range idef.members {
		mIndex := i + 1
		mValue := textSlice[mIndex]

		switch mtype {
		case TYPE_REG:
			v, exists := registers[mValue]
			if !exists {
				return nil, fmt.Errorf("Unrecognized register `%s`", mValue)
			}
			bytecode[mIndex] = v
		case TYPE_NUM:
			base := 10

			if strings.HasPrefix(mValue, "0x") {
				base = 16
				mValue = mValue[2:]
			} else if strings.HasSuffix(mValue, "h") {
				base = 16
				mValue = mValue[0 : len(mValue)-1]
			} else if len(mValue) > 1 && strings.HasPrefix(mValue, "0") {
				base = 8
				mValue = mValue[1:]
			} else if strings.HasSuffix(mValue, "b") {
				base = 2
				mValue = mValue[0 : len(mValue)-1]
			}

			v, err := strconv.ParseUint(mValue, base, 8)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse numeric field '%s' on instruction '%+v'", mValue, textSlice)
			}
			bytecode[mIndex] = byte(v)
		default:
			logrus.Fatalf("Bad member type definition for instruction 0x%x (%s)", idef.opByte, textSlice[0])
			panic("Assembler definition is broken")
		}
	}

	return bytecode, nil
}

var mnemonics = map[string][]instructionDef{

	"start": []instructionDef{
		instructionDef{
			members: []int{},
			opByte:  0x00,
		},
	},

	"end": []instructionDef{
		instructionDef{
			members: []int{},
			opByte:  0x0f,
		},
	},

	"set": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG, TYPE_REG},
			opByte:  0x10,
		},
		instructionDef{
			members: []int{TYPE_REG, TYPE_NUM},
			opByte:  0x11,
		},
	},

	"put": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG},
			opByte:  0x20,
		},
		instructionDef{
			members: []int{TYPE_NUM},
			opByte:  0x21,
		},
	},

	"get": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG},
			opByte:  0x22,
		},
	},

	"add": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG, TYPE_REG},
			opByte:  0x30,
		},
		instructionDef{
			members: []int{TYPE_REG, TYPE_NUM},
			opByte:  0x31,
		},
	},

	"sub": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG, TYPE_REG},
			opByte:  0x32,
		},
		instructionDef{
			members: []int{TYPE_REG, TYPE_NUM},
			opByte:  0x33,
		},
	},

	"dec": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG},
			opByte:  0x34,
		},
	},

	"inc": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG},
			opByte:  0x35,
		},
	},

	"not": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG},
			opByte:  0x40,
		},
	},

	"neg": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG},
			opByte:  0x41,
		},
	},

	"and": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG, TYPE_REG},
			opByte:  0x42,
		},
		instructionDef{
			members: []int{TYPE_REG, TYPE_NUM},
			opByte:  0x43,
		},
	},

	"or": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG, TYPE_REG},
			opByte:  0x44,
		},
		instructionDef{
			members: []int{TYPE_REG, TYPE_NUM},
			opByte:  0x45,
		},
	},

	"xor": []instructionDef{
		instructionDef{
			members: []int{TYPE_REG, TYPE_REG},
			opByte:  0x46,
		},
		instructionDef{
			members: []int{TYPE_REG, TYPE_NUM},
			opByte:  0x47,
		},
	},
}

var registers = map[string]byte{
	// this must match the VM hardware CPU register order
	"ra": 0x00,
	"rb": 0x01,
	"rc": 0x02,
	"rd": 0x03,
	"sp": 0x04,
	"pc": 0x05,
}

func AsmLine(input string) ([]byte, []error) {
	// Get rid of comment section
	processedInput := strings.Split(input, "#")[0]
	// Get rid of surrounding spaces and tabs
	processedInput = strings.Trim(processedInput, " \t\r")
	if len(processedInput) == 0 {
		return []byte{}, nil
	}

	posibleInstruction := normalize(processedInput)
	if l := len(posibleInstruction); l < 1 || l > 3 {
		return []byte{}, []error{fmt.Errorf("Invalid instruction: `%s`", input)}
	}

	parseErrors := []error{}
	for _, p := range mnemonics[posibleInstruction[0]] {
		bytecode, err := p.parse(posibleInstruction)
		if err != nil {
			parseErrors = append(parseErrors, err)
		} else {
			return bytecode, nil
		}
	}

	return nil, parseErrors
}

func normalize(text string) []string {
	processedInput := strings.ToLower(text)
	processedInput = strings.ReplaceAll(processedInput, ",", " ")
	processedInput = strings.ReplaceAll(processedInput, "\t", " ")
	processedInput = strings.ReplaceAll(processedInput, "\r", " ")
	// processedInput is now
	// "[no trail][lowercase op][N spaces][lowercase val1?][N spaces][lowercase val2?]"
	nonEmptySections := []string{}
	for _, section := range strings.Split(processedInput, " ") {
		if len(section) > 0 {
			nonEmptySections = append(nonEmptySections, section)
		}
	}
	return nonEmptySections
}
