# simple-vm
I wanted to write a really simple vm just to get an idea of the inner workings of one.  
Ended up with a fake arch, isa, assembler, dissasembler, executable file format and of course, the VM.  

WIP:  
- utils module with some function that are replicated over the different utilities  
- better folder structure and pack it into a proper mo module + bin install  

##  Changelog  
v1 half baked version implemented in C  
v2 reimplemntation in go. WIP  


## Asm cli
WIP:  
- Define a better encoding so we can add ops like jmp, jne, je, call, ret, etc  

Usage:  
> cd simple-vm/asm  
> cat sample.asm | go run main.go -o sample.bin -v  
> cat sample.bin | xxd  

## VM
WIP:  
- Implement cli args for reading a bin file  
- Implement jmp, jne, je, call, ret, etc  


## ISA 
WIP:  
- Write a full desc after finishing with the VM impl  


## File format  
WIP:
- Write the header def after finishing the "linker"  


## Asm + "Link" + Run a yz file inside the VM

> cd out/ && rm *
> cat ../examples/test.asm | go run ../asm/main.go -v -o test.bin  
> go run ../link/main.go -v -i test.bin -o test.yz  
> cat test.bin | xxd  
> cat test.yz | xxd  
> go run ../vm/main.go -v -i test.yz  