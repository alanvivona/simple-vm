# simple-vm
A rather simple vm implementation

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