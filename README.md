# Simple VM in Golang
I wanted to write a really simple vm just to get an idea of the inner workings of one.  
Ended up with a fake arch, isa, assembler, dissasembler, executable file format and of course, the VM.  

##  Changelog  
v1   half baked version implemented in C  
v2   reimplemntation in go.  
v2.1 gomodules and new examples.

# Usage

## Assembler Module  
> go run ./asm/main.go -h  

The module takes input from stdin and outputs the binary to a file.  

> cat examples/sample2.asm | go run ./asm/main.go -o out/sample.bin -v  
> cat out/sample.bin | xxd    

## Linker Module  
> go run ./link/main.go -h  

The module takes a path to a raw binary file and saves an executable.  
Executable files have a header and a checksum.

> go run ./link/main.go -v -i sample.bin -o sample.yz    
> cat sample.bin | xxd    
> cat sample.yz | xxd    

## Virtual Machine Module  

The module checks and runs an executable file.  

> go run ./main.go -v -i out/sample.yz    

## Instrution Set

#### Setup
`start` : Start machine execution  
`end` : End machine execution  

#### Set registers
`set` : Set register to register or immediate value

#### Stack
`put` : Push register or immediate value into the stack    
`get` : Fetch value from the stack into register   

#### Arithmetic
`add` : Add register or immediate value to register    
`sub` : Substract register or immediate value to register   
`dec` : Decrement register value by one   
`inc` : Increment register value by one  

#### Logic
`not` : Negate register value   
`neg` : One's complement of register value  
`and` : Logic AND between regiter or immediate value and target register  
`or`  : Logic OR between regiter or immediate value and target register  
`xor` : Logic XOR between regiter or immediate value and target register  

## The YZ File format  
| Section       | Subsection     | Description     |
| :------------- | :----------: | -----------: |
|  Header |  | Describes file structure |
| | Magic Number / File Signature | Two bytes signature. Value: 0x5959 ("YY") |
| | Checksum | Sha512 of the code section |
| | Size field | 64 integer indicating next's section size |
| | Next section header | Label. Next section's name. Value: "code" |
| | Section delimiter | Indicates section end. Value: 0x6060 ("ZZ") |
| Code | |  Raw binary to be run |

# Demo

## Assemble + Create executable + Run inside the VM  

Reset the output folder:  
> rm -rf ./out; mkdir out  

Create a binary file by running the source code through the assembler:  
> cat ./examples/sample2.asm | go run ./asm/main.go -v -o test.bin    

Create an executable by using the linker:
> go run ./link/main.go -v -i test.bin -o test.yz    

Notice the difference between raw binary and executable:  
> cat test.bin | xxd    
> cat test.yz | xxd    

You should see a checksum, magic number and header end were added to the executable.  

Run the executable in the virtual machine:  
> go run ./vm/main.go -v -i test.yz    