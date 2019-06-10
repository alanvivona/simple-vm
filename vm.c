#include <stdio.h>
#include <string.h>

/* 
    Virtual machine description

    Memory:
        0x000   .stack  size: 512b 0x200, perms: rw-
        0x200   .text   size: 512b 0x200, perms: r-x
    
    CPU Registers:
        r0-r9   16-bit general purpouse
        sp      16-bit stack pointer
        pc      16-bit program counter
        lr      16-bit link register

    ISA (ARMv6 inspired instruction set):
        vstr
            this is the only instruction capable of storing into memory
            vstr r0, [r1]       store r0 into *r1

        vldr
            this is the only instruction capable of loading from memory
            vldr r0, [r1]       laod r0 from *r1
        
        vmov
            copy immediate value or register content to another register
            vmov r0, r1         copy values from r1 into r0
            vmov r0, #5         copy immediate value #5 into r0

        vadd
            add immediate value or register content with another register
            stores result into 1st operand
            vadd r0, r0, r1     r0=r0+r1
            vadd r0, r1, #5     r0=r1+5

        veor
            perform bitwise exclusive or (XOR) on an immediate value or register content with another register
            stores result into 1st operand
            veor r0, r0, r1     r0=r0^r1
            veor r0, r1, #5     r0=r1^5

        vand
            perform bitwise AND on an immediate value or register content with another register
            stores result into 1st operand
            vand r0, r0, r1     r0=r0&r1
            vand r0, r1, #5     r0=r1&5

        vor
            perform bitwise OR on an immediate value or register content with another register
            stores result into 1st operand
            vor  r0, r0, r1     r0=r0|r1
            vor  r0, r1, #5     r0=r1|5
        
        vb
            branches to another location (updates pc register)
            value can be an offset or register
            vb   r0             branches to *r0
            vb   #5             branches to offset +5
        
        vbl
            branch and link to another location (updates pc register and link register)
            value can be an offset or register
            vbl   r0            branch and link to *r0
            vbl   #5            branch and link to offset +5

        vret
            branch to link register content (updates pc register and empties the link register value)
            vret

*/

/* Virtual memory map */
typedef struct vmem_section {
    bool perms[3]{true, false, true};
    char space[512];
} vmem_section_text;

typedef struct vmem_section {
    bool perms[3]{true, true, false};
    char space[512];
} vmem_section_stack;

typedef struct vmem { 
    vmem_section_text   text;
    vmem_section_stack  stack;
} vmem;

/* Virtual ISA implementation */
int vstr(val, dst){
    vmem.stack[dst] = val;    
}

int vldr(dst, src){
    dst = vmem.stack[src];    
}

int vmov(dst, src){
    dst = src;
}

int vadd(dst, opa, opb){
    dst=opa+opb;
}

int veor(dst, opa, opb){
    dst=opa^opb;
}

int vand(dst, opa, opb){
    dst=opa&opb;
}

int vor(dst, opa, opb){
    dst=opa|opb;
}

int vb(dst, offset){
    vcpu.pc = dst;
}            
    
int vbl(dst, offset){

}

int vret(){

}

int main(int argc, char const *argv[])
{

    if (argc < 2)
    {
        printf("Only %d arguments provided", argc);
        return 0;
    }
    
    for (size_t i = 0; i < sizeof(text_section); i++)
    {
        
    }
    
    
    return 0;
}
