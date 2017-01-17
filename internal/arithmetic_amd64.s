// +build !noasm !appengine

#define NOSPLIT 7

// func initasm()(a,a2 bool)
// pulled from runtime/asm_amd64.s
TEXT ·initasm(SB), NOSPLIT, $0
	MOVQ $1, R15

	MOVQ $1, AX
	CPUID

	ANDL $0x1, CX
	CMPL CX, $0x1
	CMOVQEQ R15, R9
	MOVB R9, ·Sse3Supt(SB)
	XORQ R9, R9

	MOVQ $1, AX
	CPUID
	ANDL $0x18001000, CX
	CMPL CX, $0x18001000
	CMOVQEQ R15, R9
	MOVB R9, ·FmaSupt(SB) // set numgo·fmaSupt
	XORQ R9, R9

	ANDL $0x18000000, CX 
	CMPL CX, $0x18000000
	JNE  noavx

	// For XGETBV, OSXSAVE bit is required and sufficient
	MOVQ $0, CX

	// Check for FMA capability
	// XGETBV
	BYTE $0x0F; BYTE $0x01; BYTE $0xD0
	
	ANDL $6, AX
	CMPL AX, $6                        // Check for OS support of YMM registers
	JNE  noavx
	MOVB $1, ·AvxSupt(SB)              // set numgo·avxSupt

	// Check for AVX2 capability
	MOVL $7, AX
	XORQ CX, CX
	CPUID
	ANDL $0x20, BX         // check for AVX2 bit
	CMPL BX, $0x20
	CMOVQEQ R15, R9
	MOVB R9, ·Avx2Supt(SB) // set numgo·avx2Supt
	XORQ R9, R9
	RET

noavx:
	MOVB $0, ·FmaSupt(SB) // set numgo·fmaSupt
	MOVB $0, ·AvxSupt(SB) // set numgo·avxSupt
	MOVB $0, ·Avx2Supt(SB) // set numgo·avx2Supt
	RET

// func AddC(c float64, d []float64)
TEXT ·AddC(SB), NOSPLIT, $0
	// data ptr
	MOVQ d+8(FP), R10

	// n = data len
	MOVQ d_len+16(FP), SI

	// zero len return
	CMPQ SI, $0
	JE   ACEND

	// check tail
	SUBQ $4, SI
	JL   ACTAIL

	// avx support test
	LEAQ c+0(FP), R9
	CMPB ·AvxSupt(SB), $1
	JE   AVX_AC
	CMPB ·Avx2Supt(SB), $1
	JE  AVX2_AC

	// load multiplier
	MOVSD  (R9), X0
	SHUFPD $0, X0, X0

ACLOOP:  // Unrolled x2 d[i]|d[i+1] += c
	MOVUPD 0(R10), X1
	MOVUPD 16(R10), X2
	ADDPD  X0, X1
	ADDPD  X0, X2
	MOVUPD X1, 0(R10)
	MOVUPD X2, 16(R10)
	ADDQ   $32, R10
	SUBQ   $4, SI
	JGE    ACLOOP
	JMP    ACTAIL

// NEED AVX INSTRUCTION CODING FOR THIS TO WORK
AVX2_AC: // Until AVX2 is known
AVX_AC:
	// VBROADCASTD (R9), Y0
	BYTE $0xC4; BYTE $0xC2; BYTE $0x7D; BYTE $0x19; BYTE $0x01

AVX_ACLOOP:
	// VADDPD (R10),Y0,Y1
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7D; BYTE $0x58; BYTE $0x0A

	// VMOVDQA Y1, (R10)
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7E; BYTE $0x7F; BYTE $0x0A
	
	ADDQ $32, R10
	SUBQ $4, SI
	JGE  AVX_ACLOOP
	//VZEROUPPER
	BYTE $0xC5; BYTE $0xF8; BYTE $0x77

ACTAIL:  // Catch len % 4 == 0
	ADDQ $4, SI
	JE   ACEND

ACTL:  // Calc the last values individually d[i] += c
	MOVSD 0(R10), X1
	ADDSD X0, X1
	MOVSD X1, 0(R10)
	ADDQ  $8, R10
	SUBQ  $1, SI
	JG    ACTL

ACEND:
	RET

// func subtrC(c float64, d []float64)
TEXT ·SubtrC(SB), NOSPLIT, $0
	// data ptr
	MOVQ d+8(FP), R10

	// n = data len
	MOVQ d_len+16(FP), SI

	// zero len return
	CMPQ SI, $0
	JE   SCEND

	// check tail
	SUBQ $4, SI
	JL   SCTAIL

	// load multiplier
	MOVSD  c+0(FP), X0
	SHUFPD $0, X0, X0

SCLOOP:  // load d[i] | d[i+1]
	MOVUPD 0(R10), X1
	MOVUPD 16(R10), X2
	SUBPD  X0, X1
	SUBPD  X0, X2
	MOVUPD X1, 0(R10)
	MOVUPD X2, 16(R10)
	ADDQ   $32, R10
	SUBQ   $4, SI
	JGE    SCLOOP

SCTAIL:
	ADDQ $4, SI
	JE   SCEND

SCTL:
	MOVSD 0(R10), X1
	SUBSD X0, X1
	MOVSD X1, 0(R10)
	ADDQ  $8, R10
	SUBQ  $1, SI
	JG    SCTL

SCEND:
	RET

// func multC(c float64, d []float64)
TEXT ·MultC(SB), NOSPLIT, $0
	MOVQ d_base+8(FP), R10
	MOVQ d_len+16(FP), SI

	// zero len return
	CMPQ SI, $0
	JE   MCEND
	SUBQ $4, SI
	JL   MCTAIL

	// load multiplier
	MOVSD  c+0(FP), X0
	SHUFPD $0, X0, X0

MCLOOP:  // load d[i] | d[i+1]
	MOVUPD 0(R10), X1
	MOVUPD 16(R10), X2
	MULPD  X0, X1
	MULPD  X0, X2
	MOVUPD X1, 0(R10)
	MOVUPD X2, 16(R10)
	ADDQ   $32, R10
	SUBQ   $4, SI
	JGE    MCLOOP

MCTAIL:
	ADDQ $4, SI
	JE   MCEND

MCTL:
	MOVSD 0(R10), X1
	MULSD X0, X1
	MOVSD X1, 0(R10)
	ADDQ  $8, R10
	SUBQ  $1, SI
	JG    MCTL

MCEND:
	RET

// func divC(c float64, d []float64)
TEXT ·DivC(SB), NOSPLIT, $0
	// data ptr
	MOVQ d+8(FP), R10

	// n = data len
	MOVQ d_len+16(FP), SI

	// zero len return
	CMPQ SI, $0
	JE   DCEND

	// check tail
	SUBQ $4, SI
	JL   DCTAIL

	// load multiplier
	MOVSD  c+0(FP), X0
	SHUFPD $0, X0, X0

DCLOOP:  // load d[i] | d[i+1]
	MOVUPD 0(R10), X1
	MOVUPD 16(R10), X2
	DIVPD  X0, X1
	DIVPD  X0, X2
	MOVUPD X1, 0(R10)
	MOVUPD X2, 16(R10)
	ADDQ   $32, R10
	SUBQ   $4, SI
	JGE    DCLOOP 

DCTAIL:
	ADDQ $4, SI
	JE   DCEND

DCTL:
	MOVSD 0(R10), X1
	DIVSD X0, X1
	MOVSD X1, 0(R10)
	ADDQ  $8, R10
	SUBQ  $1, SI
	JG    DCTL

DCEND:
	RET

// func add(a,b []float64)
TEXT ·Add(SB), NOSPLIT, $0
	// a data ptr
	MOVQ a_base+0(FP), R8

	// a len
	MOVQ a_len+8(FP), SI

	// b data ptr
	MOVQ b_base+24(FP), R9
	MOVQ R9, R10

	// b len
	MOVQ b_len+32(FP), DI
	MOVQ DI, R11

	// zero len return
	CMPQ SI, $0
	JE   AEND

	// check tail
	SUBQ $2, SI
	JL   ATAIL

ALD:
	CMPQ DI, $1
	JE   ALT
	SUBQ $2, DI
	JGE  ALO
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

ALO:
	MOVUPD (R9), X1
	ADDQ   $16, R9
	JMP    ALOOP

ALT:
	MOVLPD (R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD (R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

ALOOP:
	MOVUPD (R8), X0
	ADDPD  X1, X0
	MOVUPD X0, (R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    ALD

ATAIL:
	ADDQ $2, SI
	JE   AEND

ATL:
	MOVSD (R8), X0
	MOVSD (R9), X1
	ADDSD X1, X0
	MOVSD X0, (R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JG    ATL

AEND:
	RET

// func vadd(a,b[]float64)
// req:  len(a) == len(b)
TEXT ·Vadd(SB), NOSPLIT, $0
	// a data ptr
	MOVQ a_base+0(FP), R8

	// a len
	MOVQ a_len+8(FP), SI

	// b data ptr
	MOVQ b_base+24(FP), R9

	// zero len return
	CMPQ SI, $0
	JE   vadd_exit

	// check tail
	SUBQ $8, SI
	JL   vadd_tail

	// AVX vs SSE
	CMPB ·AvxSupt(SB), $1
	JE   vadd_avx_loop

vadd_loop:
	MOVUPD (R9), X1
	MOVUPD 16(R9), X3
	MOVUPD 32(R9), X5
	MOVUPD 48(R9), X7
	
	MOVUPD (R8), X0
	ADDPD  X1, X0
	MOVUPD 16(R8), X2
	ADDPD  X3, X2
	MOVUPD 32(R8), X4
	ADDPD  X5, X4
	MOVUPD 48(R8), X6
	ADDPD  X7, X6
	
	MOVUPD X0, (R8)
	MOVUPD X2, 16(R8)
	MOVUPD X4, 32(R8)
	MOVUPD X6, 48(R8)
	ADDQ   $64, R8
	ADDQ   $64, R9
	SUBQ   $8, SI
	JGE    vadd_loop

vadd_tail:
	ADDQ $8, SI
	JE   vadd_exit

vadd_tail_loop:
	MOVSD (R8), X15
	MOVSD (R9), X14
	ADDSD X14, X15
	MOVSD X15, (R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JGE   vadd_tail_loop
	JMP   vadd_exit
	
vadd_avx_loop:
	//VMOVDQU (R9), Y0
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7E; BYTE $0x6F; BYTE $0x01
	//VMOVDQU 32(R9), Y1
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7E; BYTE $0x6F; BYTE $0x49; BYTE $0x20

	// VADDPD (R8),Y0,Y0
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7D; BYTE $0x58; BYTE $0x00
	// VADDPD 32(R10),Y1,Y1
	BYTE $0xC4; BYTE $0xC1; BYTE $0x75; BYTE $0x58; BYTE $0x48; BYTE $0x20

	//VMOVDQA Y0, (R8)
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7E; BYTE $0x7F; BYTE $0x00
	//VMOVDQA Y1, 32(R8)
	BYTE $0xC4; BYTE $0xC1; BYTE $0x7E; BYTE $0x7F; BYTE $0x48; BYTE $0x20

	
	ADDQ $64, R8
	ADDQ $64, R9
	SUBQ $8, SI
	JGE  vadd_avx_loop
	//VZEROUPPER
	BYTE $0xC5; BYTE $0xF8; BYTE $0x77
	ADDQ $8, SI
	JE   vadd_exit
	JMP  vadd_tail_loop

vadd_exit:
	RET

// func hadd(st uint64, a []float64)
// req:  len(a) == len(b)
TEXT ·Hadd(SB), NOSPLIT, $0
	// a data ptr
	MOVQ a_base+8(FP), R8
	MOVQ R8, R9

	// a len
	MOVQ a_len+16(FP), SI
	MOVQ st+0(FP), CX
	MOVQ CX,  DI
	ANDQ $1, DI
	

	CMPQ CX, $1
	JE hadd_exit
	CMPQ SI, $0
	JE   hadd_exit
	CMPQ CX, $8
	JG hadd_big_stride
	CMPB ·Sse3Supt(SB), $1
	JE hadd_sse3_head

hadd_big_stride:
	// AVX vs SSE
	CMPB ·AvxSupt(SB), $1
	//JE   hadd_avx_head
	CMPB ·Sse3Supt(SB), $1
	JE hadd_sse3_head
hadd_head:
	PXOR X0, X0
	MOVQ CX, DI
	SUBQ $1, DI
hadd_loop:
	ADDPD (R8), X0
	ADDQ $16, R8
	SUBQ $2, DI
	JG hadd_loop
	JZ hadd_tail
	MOVAPD X0, X1
	UNPCKHPD X1, X0
	ADDPD X1,X0
	MOVQ X0, (R9)
	ADDQ $8, R9
	SUBQ CX, SI
	JG hadd_head
	JMP hadd_exit
hadd_tail:
	ADDSD (R8), X0
	MOVAPD X0, X1
	UNPCKHPD X1, X0
	ADDPD X1,X0
	MOVQ X0, (R9)
	ADDQ $8, R9
	SUBQ CX, SI
	JZ hadd_exit
	MOVQ 8(R8), X0
	MOVQ CX, DI
	SUBQ $2, DI
	ADDQ $16, R8
	JMP hadd_loop
hadd_sse3_head:
	PXOR X0, X0
	MOVQ CX, DI
	SUBQ $1, DI
hadd_sse3_loop:
	ADDPD (R8), X0
	ADDQ $16, R8
	SUBQ $2, DI
	JG hadd_sse3_loop
	JZ hadd_sse3_tail
	BYTE $0x66; BYTE $0x0F; BYTE $0x7C; BYTE $0xC0
	// HADDPD X0, X0  //Added in 1.6
	MOVQ X0, (R9)
	ADDQ $8, R9
	SUBQ CX, SI
	JG hadd_sse3_head
	JMP hadd_exit
hadd_sse3_tail:
	ADDSD (R8), X0
	BYTE $0x66; BYTE $0x0F; BYTE $0x7C; BYTE $0xC0
	// HADDPD X0, X0  //Added in 1.6
	MOVQ X0, (R9)
	ADDQ $8, R9
	SUBQ CX, SI
	JZ hadd_exit
	MOVQ 8(R8), X0
	MOVQ CX, DI
	SUBQ $2, DI
	ADDQ $16, R8
	JMP hadd_sse3_loop
hadd_exit:	
	RET

	
// func subtr(a,b []float64)
TEXT ·Subtr(SB), NOSPLIT, $0
	// a data ptr
	MOVQ a_base+0(FP), R8

	// a len
	MOVQ a_len+8(FP), SI

	// b data ptr
	MOVQ b_base+24(FP), R9
	MOVQ R9, R10

	// b len
	MOVQ b_len+32(FP), DI
	MOVQ DI, R11

	// zero len return
	MOVQ $0, AX
	CMPQ AX, SI
	JE   SEND

	// check tail
	SUBQ $2, SI
	JL   STAIL

SLD:
	SUBQ $1, DI
	JE   SLT
	SUBQ $1, DI
	JGE  SLO
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

SLO:
	MOVUPD 0(R9), X1
	ADDQ   $16, R9
	JMP    SLOOP

SLT:
	MOVLPD 0(R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD 0(R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

SLOOP:
	MOVUPD 0(R8), X0
	SUBPD  X1, X0
	MOVUPD X0, 0(R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    SLD

STAIL:
	ADDQ $2, SI
	JE   SEND

STL:
	MOVSD 0(R8), X0
	MOVSD 0(R9), X1
	SUBSD X1, X0
	MOVSD X0, 0(R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JG    STL

SEND:
	RET

// func mult(a,b []float64)
TEXT ·Mult(SB), NOSPLIT, $0
	// a data ptr
	MOVQ a_base+0(FP), R8

	// a len
	MOVQ a_len+8(FP), SI

	// b data ptr
	MOVQ b_base+24(FP), R9
	MOVQ R9, R10

	// b len
	MOVQ b_len+32(FP), DI
	MOVQ DI, R11

	// zero len return
	MOVQ $0, AX
	CMPQ AX, SI
	JE   MEND

	// check tail
	SUBQ $2, SI
	JL   MTAIL

MLD:
	SUBQ $1, DI
	JE   MLT
	SUBQ $1, DI
	JGE  MLO
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

MLO:
	MOVUPD 0(R9), X1
	ADDQ   $16, R9
	JMP    MLOOP

MLT:
	MOVLPD 0(R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD 0(R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

MLOOP:
	MOVUPD 0(R8), X0
	MULPD  X1, X0
	MOVUPD X0, 0(R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    MLD

MTAIL:
	ADDQ $2, SI
	JE   MEND

MTL:
	MOVSD 0(R8), X0
	MOVSD 0(R9), X1
	MULSD X1, X0
	MOVSD X0, 0(R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JG    MTL

MEND:
	RET

// func div(a,b []float64)
TEXT ·Div(SB), NOSPLIT, $0
	// a data ptr
	MOVQ a_base+0(FP), R8

	// a len
	MOVQ a_len+8(FP), SI

	// b data ptr
	MOVQ b_base+24(FP), R9
	MOVQ R9, R10

	// b len
	MOVQ b_len+32(FP), DI
	MOVQ DI, R11

	// zero len return
	MOVQ $0, AX
	CMPQ AX, SI
	JE   DEND

	// check tail
	SUBQ $2, SI
	JL   DTAIL

DLD:
	SUBQ $1, DI
	JE   DLT
	SUBQ $1, DI
	JGE  DLO
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

DLO:
	MOVUPD 0(R9), X1
	ADDQ   $16, R9
	JMP    DLOOP
DLT:
	MOVLPD 0(R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD 0(R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

DLOOP:
	MOVUPD 0(R8), X0
	DIVPD  X1, X0
	MOVUPD X0, 0(R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    DLD

DTAIL:
	ADDQ $2, SI
	JE   DEND
DTL:
	MOVSD 0(R8), X0
	MOVSD 0(R9), X1
	DIVSD X1, X0
	MOVSD X0, 0(R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JG    DTL

DEND:
	RET

// func fma12(a float64, x,b []float64)
// x[i] = a*x[i]+b[i]
TEXT ·Fma12(SB), NOSPLIT, $0
	// a ptr
	MOVSD  a+0(FP), X2
	SHUFPD $0, X2, X2

	// x data ptr
	MOVQ x_base+8(FP), R8

	// x len
	MOVQ x_len+16(FP), SI

	// b data ptr
	MOVQ b_base+32(FP), R9
	MOVQ R9, R10

	// b len
	MOVQ b_len+40(FP), DI
	MOVQ DI, R11

	// zero len return
	CMPQ SI, $0
	JE   F12END

	// check tail
	SUBQ $2, SI
	JL   F12TAIL

F12LD:
	CMPQ DI, $1
	JE   F12LT
	SUBQ $2, DI
	JGE  F12LO
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

F12LO:
	MOVUPD (R9), X1
	ADDQ   $16, R9
	JMP    F12LOOP

F12LT:
	MOVLPD (R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD (R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

F12LOOP:
	MOVUPD (R8), X0
	MULPD  X2, X0
	ADDPD  X1, X0
	MOVUPD X0, (R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    F12LD
	JMP    F12TAIL

F12LDF:
	CMPQ DI, $1
	JE   F12LTF
	SUBQ $2, DI
	JGE  F12LOF
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

F12LOF:
	MOVUPD (R9), X1
	ADDQ   $16, R9
	JMP    F12LOOPF

F12LTF:
	MOVLPD (R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD (R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

F12LOOPF:
	MOVUPD (R8), X0

	// VMFADD213PD X0, X1, X2
	BYTE   $0xC4; BYTE $0xE2; BYTE $0xF1; BYTE $0x98; BYTE $0xC2
	MOVUPD X0, (R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    F12LDF

F12TAIL:
	ADDQ $2, SI
	JE   F12END

F12TL:
	MOVSD (R8), X0
	MOVSD (R9), X1
	MULPD X2, X0
	ADDPD X1, X0
	MOVSD X0, (R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JG    F12TL

F12END:
	RET

// func fma21(a float64, x,b []float64)
// x[i] = x[i]*b[i]+a
TEXT ·Fma21(SB), NOSPLIT, $0
	// a ptr
	MOVSD  a+0(FP), X2
	SHUFPD $0, X2, X2

	// x data ptr
	MOVQ x_base+8(FP), R8

	// x len
	MOVQ x_len+16(FP), SI

	// b data ptr
	MOVQ b_base+32(FP), R9
	MOVQ R9, R10

	// b len
	MOVQ b_len+40(FP), DI
	MOVQ DI, R11

	// zero len return
	CMPQ SI, $0
	JE   F21END

	// check tail
	SUBQ $2, SI
	JL   F21TAIL

F21LD:
	CMPQ DI, $1
	JE   F21LT
	SUBQ $2, DI
	JGE  F21LO
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

F21LO:
	MOVUPD (R9), X1
	ADDQ   $16, R9
	JMP    F21LOOP

F21LT:
	MOVLPD (R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD (R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

F21LOOP:
	MOVUPD (R8), X0
	MULPD  X1, X0
	ADDPD  X2, X0
	MOVUPD X0, (R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    F21LD
	JMP    F21TAIL

F21LDF:
	CMPQ DI, $1
	JE   F21LTF
	SUBQ $2, DI
	JGE  F21LOF
	MOVQ R10, R9
	MOVQ R11, DI
	SUBQ $2, DI

F21LOF:
	MOVUPD (R9), X1
	ADDQ   $16, R9
	JMP    F21LOOPF

F21LTF:
	MOVLPD (R9), X1
	MOVQ   R10, R9
	MOVQ   R11, DI
	MOVHPD (R9), X1
	SUBQ   $1, DI
	ADDQ   $8, R9

F21LOOPF:
	MOVUPD (R8), X0

	// VMFADD213PD X0, X1, X2
	BYTE   $0xC4; BYTE $0xE2; BYTE $0xF1; BYTE $0xA8; BYTE $0xC2
	MOVUPD X0, (R8)
	ADDQ   $16, R8
	SUBQ   $2, SI
	JGE    F21LDF

F21TAIL:
	ADDQ $2, SI
	JE   F21END

F21TL:
	MOVSD (R8), X0
	MOVSD (R9), X1
	MULPD X1, X0
	ADDPD X2, X0
	MOVSD X0, (R8)
	ADDQ  $8, R8
	ADDQ  $8, R9
	SUBQ  $1, SI
	JG    F21TL

F21END:
	RET
