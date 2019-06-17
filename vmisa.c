#ifndef VM_ISA_H
#define VM_ISA_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

/* 
    ISA (ARMv6 inspired instruction set):
        vnop
            no operaion
            vnop                does nothing

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

void vnop(){
}

void vadd(uint32_t op1, uint32_t op2, uint32_t* res){
    *res = op1 + op2;
}

void vsub(uint32_t op1, uint32_t op2, uint32_t* res){
    *res = op1 - op2;
}

void vand(uint32_t op1, uint32_t op2, uint32_t* res){
    *res = op1 & op2;
}

void vor(uint32_t op1, uint32_t op2, uint32_t* res){
    *res = op1 | op2;
}

void veor(uint32_t op1, uint32_t op2, uint32_t* res){
    *res = op1 ^ op2;
}

void vmov(uint32_t* dst, uint32_t* src){
    *dst = *src;
}

void vb(uint32_t* src, uint32_t* dst, uint32_t* offset){
    // Relative Branch by offset
    if (*offset)
    {
        *src = *src + *offset;
        return;
    }
    // Absolute Branch
    *src = *dst;
    return;
}

#endif /* VM_ISA_H */
