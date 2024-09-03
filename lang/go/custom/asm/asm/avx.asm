	.text
	.intel_syntax noprefix
	.file	"avx.c"
	.globl	Avx2SsdInt16                    # -- Begin function Avx2SsdInt16
	.p2align	4, 0x90
	.type	Avx2SsdInt16,@function
Avx2SsdInt16:                           # @Avx2SsdInt16
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	vmovdqu	ymm0, ymmword ptr [rdi]
	vpsubw	ymm0, ymm0, ymmword ptr [rsi]
	vpmaddwd	ymm0, ymm0, ymm0
	vmovdqu	ymmword ptr [rdx], ymm0
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.Lfunc_end0:
	.size	Avx2SsdInt16, .Lfunc_end0-Avx2SsdInt16
                                        # -- End function
	.ident	"Ubuntu clang version 14.0.0-1ubuntu1.1"
	.section	".note.GNU-stack","",@progbits
	.addrsig
