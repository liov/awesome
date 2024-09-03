msys2下都不行，放弃，有空试试linux
# goat 跑不成功
go install github.com/gorse-io/goat@latest
clang -O3 -S -c xxx.c -o xxx.asm -mno-red-zone, -mstackrealign, -mllvm, -inline-threshold=1000,
-fno-asynchronous-unwind-tables, -fno-exceptions, -fno-rtti

clang -O3 -c xxx.c -o xxx.o -mno-red-zone, -mstackrealign, -mllvm, -inline-threshold=1000,
-fno-asynchronous-unwind-tables, -fno-exceptions, -fno-rtti

objdump -d xxx.o --insn-width 16

goat avx.c -O3 -mavx2


# c2goasm  linux 14.0.0 -O3可行 | msys2 18+ 不行,有未识别符号
go install github.com/minio/c2goasm@latest
go install github.com/minio/asm2plan9s@latest
clang -O3 -S -c avx.c -o asm/avx.asm -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -mavx2
c2goasm -a asm/avx.asm asm/avx.s

msys2 clang 18 和clang version 17.0.3 Target: x86_64-pc-windows-msvc一致
```asm
	.text
	.def	@feat.00;
	.scl	3;
	.type	0;
	.endef
	.globl	@feat.00
.set @feat.00, 0
	.intel_syntax noprefix
	.file	"avx.c"
	.def	Avx2SsdInt16;
	.scl	2;
	.type	32;
	.endef
	.globl	Avx2SsdInt16                    # -- Begin function Avx2SsdInt16
	.p2align	4, 0x90
Avx2SsdInt16:                           # @Avx2SsdInt16
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	vmovdqu	ymm0, ymmword ptr [rcx]
	vpsubw	ymm0, ymm0, ymmword ptr [rdx]
	vpmaddwd	ymm0, ymm0, ymm0
	vmovdqu	ymmword ptr [r8], ymm0
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
                                        # -- End function
	.addrsig

```
linux clang 14
```asm
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

```