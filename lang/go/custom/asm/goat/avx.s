//go:build !noasm && amd64
// Code generated by GoAT. DO NOT EDIT.

TEXT ·avx2_ssd_int16(SB), $0-32
	MOVQ a+0(FP), DI
	MOVQ b+8(FP), SI
	MOVQ c+16(FP), DX
	BYTE $0x55                   // pushq	%rbp
	WORD $0x8948; BYTE $0xe5     // movq	%rsp, %rbp
	LONG $0xf8e48348             // andq	$-8, %rsp
	LONG $0x016ffec5             // vmovdqu	(%rcx), %ymm0
	LONG $0x02f9fdc5             // vpsubw	(%rdx), %ymm0, %ymm0
	LONG $0xc0f5fdc5             // vpmaddwd	%ymm0, %ymm0, %ymm0
	LONG $0x7f7ec1c4; BYTE $0x00 // vmovdqu	%ymm0, (%r8)
	WORD $0x8948; BYTE $0xec     // movq	%rbp, %rsp
	BYTE $0x5d                   // popq	%rbp
	WORD $0xf8c5; BYTE $0x77     // vzeroupper
	BYTE $0xc3                   // retq

