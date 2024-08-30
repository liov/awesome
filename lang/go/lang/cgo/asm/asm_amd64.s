// System V AMD64 ABI
// func asmCallCAdd(cfun uintptr, a, b int64) int64
TEXT ·AsmCallCAdd(SB),  $0
    MOVQ cfun+0(FP), AX // cfun
    MOVQ a+8(FP),    DI // a
    MOVQ b+16(FP),   SI // b
    CALL AX
    MOVQ AX, ret+24(FP)
    RET

// func SyscallWrite_Darwin(fd int, msg string) int
TEXT ·SyscallWrite_Darwin(SB),  $0
    MOVQ $(0x2000000+4), AX // #define SYS_write 4
    MOVQ fd+0(FP),       DI
    MOVQ msg_data+8(FP), SI
    MOVQ msg_len+16(FP), DX
    SYSCALL
    MOVQ AX, ret+0(FP)
    RET


TEXT ·_mm512_mul_to(SB), $0-32
 MOVQ a+0(FP), DI
 MOVQ b+8(FP), SI
 MOVQ c+16(FP), DX
 MOVQ n+24(FP), CX
 BYTE $0x55               // pushq %rbp
 WORD $0x8948; BYTE $0xe5 // movq %rsp, %rbp
 LONG $0xf8e48348         // andq $-8, %rsp
 LONG $0x0f498d4c         // leaq 15(%rcx), %r9
 WORD $0x8548; BYTE $0xc9 // testq %rcx, %rcx
 LONG $0xc9490f4c         // cmovnsq %rcx, %r9
 LONG $0x04e9c149         // shrq $4, %r9
 WORD $0x8944; BYTE $0xc8 // movl %r9d, %eax
 WORD $0xe0c1; BYTE $0x04 // shll $4, %eax
 WORD $0xc129             // subl %eax, %ecx
 WORD $0x8545; BYTE $0xc9 // testl %r9d, %r9d
 JLE  LBB3_6
 LONG $0xff418d41         // leal -1(%r9), %eax
 WORD $0x8945; BYTE $0xc8 // movl %r9d, %r8d
 LONG $0x03e08341         // andl $3, %r8d
 WORD $0xf883; BYTE $0x03 // cmpl $3, %eax
 JB   LBB3_4
 LONG $0xfce18341         // andl $-4, %r9d
 WORD $0xf741; BYTE $0xd9 // negl %r9d

LBB3_3:
 LONG $0x487cf162; WORD $0x0710             // vmovups (%rdi), %zmm0
 LONG $0x487cf162; WORD $0x0659             // vmulps (%rsi), %zmm0, %zmm0
 LONG $0x487cf162; WORD $0x0211             // vmovups %zmm0, (%rdx)
 LONG $0x487cf162; WORD $0x4710; BYTE $0x01 // vmovups 64(%rdi), %zmm0
 LONG $0x487cf162; WORD $0x4659; BYTE $0x01 // vmulps 64(%rsi), %zmm0, %zmm0
 LONG $0x487cf162; WORD $0x4211; BYTE $0x01 // vmovups %zmm0, 64(%rdx)
 LONG $0x487cf162; WORD $0x4710; BYTE $0x02 // vmovups 128(%rdi), %zmm0
 LONG $0x487cf162; WORD $0x4659; BYTE $0x02 // vmulps 128(%rsi), %zmm0, %zmm0
 LONG $0x487cf162; WORD $0x4211; BYTE $0x02 // vmovups %zmm0, 128(%rdx)
 LONG $0x487cf162; WORD $0x4710; BYTE $0x03 // vmovups 192(%rdi), %zmm0
 LONG $0x487cf162; WORD $0x4659; BYTE $0x03 // vmulps 192(%rsi), %zmm0, %zmm0
 LONG $0x487cf162; WORD $0x4211; BYTE $0x03 // vmovups %zmm0, 192(%rdx)
 LONG $0x00c78148; WORD $0x0001; BYTE $0x00 // addq $256, %rdi
 LONG $0x00c68148; WORD $0x0001; BYTE $0x00 // addq $256, %rsi
 LONG $0x00c28148; WORD $0x0001; BYTE $0x00 // addq $256, %rdx
 LONG $0x04c18341                           // addl $4, %r9d
 JNE  LBB3_3

LBB3_4:
 WORD $0x8545; BYTE $0xc0 // testl %r8d, %r8d
 JE   LBB3_6

LBB3_5:
 LONG $0x487cf162; WORD $0x0710 // vmovups (%rdi), %zmm0
 LONG $0x487cf162; WORD $0x0659 // vmulps (%rsi), %zmm0, %zmm0
 LONG $0x487cf162; WORD $0x0211 // vmovups %zmm0, (%rdx)
 LONG $0x40c78348               // addq $64, %rdi
 LONG $0x40c68348               // addq $64, %rsi
 LONG $0x40c28348               // addq $64, %rdx
 LONG $0xffc08341               // addl $-1, %r8d
 JNE  LBB3_5

LBB3_6:
 WORD $0xf983; BYTE $0x07 // cmpl $7, %ecx
 JLE  LBB3_8
 LONG $0x0710fcc5         // vmovups (%rdi), %ymm0
 LONG $0x0659fcc5         // vmulps (%rsi), %ymm0, %ymm0
 LONG $0x0211fcc5         // vmovups %ymm0, (%rdx)
 LONG $0x20c78348         // addq $32, %rdi
 LONG $0x20c68348         // addq $32, %rsi
 LONG $0x20c28348         // addq $32, %rdx
 WORD $0xc183; BYTE $0xf8 // addl $-8, %ecx

LBB3_8:
 WORD $0xc985             // testl %ecx, %ecx
 JLE  LBB3_14
 WORD $0xc989             // movl %ecx, %ecx
 LONG $0xff418d48         // leaq -1(%rcx), %rax
 WORD $0x8941; BYTE $0xc8 // movl %ecx, %r8d
 LONG $0x03e08341         // andl $3, %r8d
 LONG $0x03f88348         // cmpq $3, %rax
 JAE  LBB3_15
 WORD $0xc031             // xorl %eax, %eax
 JMP  LBB3_11

LBB3_15:
 WORD $0xe183; BYTE $0xfc // andl $-4, %ecx
 WORD $0xc031             // xorl %eax, %eax

LBB3_16:
 LONG $0x0410fac5; BYTE $0x87   // vmovss (%rdi,%rax,4), %xmm0
 LONG $0x0459fac5; BYTE $0x86   // vmulss (%rsi,%rax,4), %xmm0, %xmm0
 LONG $0x0411fac5; BYTE $0x82   // vmovss %xmm0, (%rdx,%rax,4)
 LONG $0x4410fac5; WORD $0x0487 // vmovss 4(%rdi,%rax,4), %xmm0
 LONG $0x4459fac5; WORD $0x0486 // vmulss 4(%rsi,%rax,4), %xmm0, %xmm0
 LONG $0x4411fac5; WORD $0x0482 // vmovss %xmm0, 4(%rdx,%rax,4)
 LONG $0x4410fac5; WORD $0x0887 // vmovss 8(%rdi,%rax,4), %xmm0
 LONG $0x4459fac5; WORD $0x0886 // vmulss 8(%rsi,%rax,4), %xmm0, %xmm0
 LONG $0x4411fac5; WORD $0x0882 // vmovss %xmm0, 8(%rdx,%rax,4)
 LONG $0x4410fac5; WORD $0x0c87 // vmovss 12(%rdi,%rax,4), %xmm0
 LONG $0x4459fac5; WORD $0x0c86 // vmulss 12(%rsi,%rax,4), %xmm0, %xmm0
 LONG $0x4411fac5; WORD $0x0c82 // vmovss %xmm0, 12(%rdx,%rax,4)
 LONG $0x04c08348               // addq $4, %rax
 WORD $0x3948; BYTE $0xc1       // cmpq %rax, %rcx
 JNE  LBB3_16

LBB3_11:
 WORD $0x854d; BYTE $0xc0 // testq %r8, %r8
 JE   LBB3_14
 LONG $0x820c8d48         // leaq (%rdx,%rax,4), %rcx
 LONG $0x86148d48         // leaq (%rsi,%rax,4), %rdx
 LONG $0x87048d48         // leaq (%rdi,%rax,4), %rax
 WORD $0xf631             // xorl %esi, %esi

LBB3_13:
 LONG $0x0410fac5; BYTE $0xb0 // vmovss (%rax,%rsi,4), %xmm0
 LONG $0x0459fac5; BYTE $0xb2 // vmulss (%rdx,%rsi,4), %xmm0, %xmm0
 LONG $0x0411fac5; BYTE $0xb1 // vmovss %xmm0, (%rcx,%rsi,4)
 LONG $0x01c68348             // addq $1, %rsi
 WORD $0x3949; BYTE $0xf0     // cmpq %rsi, %r8
 JNE  LBB3_13

LBB3_14:
 WORD $0x8948; BYTE $0xec // movq %rbp, %rsp
 BYTE $0x5d               // popq %rbp
 WORD $0xf8c5; BYTE $0x77 // vzeroupper
 BYTE $0xc3               // retq
