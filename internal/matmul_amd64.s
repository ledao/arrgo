// +build !noasm !appengine

#define NOSPLIT 7

// func dotProd(a,b []float64) (float64)
TEXT ·DotProd(SB), NOSPLIT, $0
	// a data ptr
	MOVQ 	a_base+0(FP), R8
	MOVQ 	a_len+8(FP), SI
	MOVQ 	b_base+24(FP), R9
	XORQ 	DI, DI
	PXOR 	X0, X0
	
	// zero len return
	CMPQ  	SI, $0
	JE   	dotp_end

	// check tail
	SUBQ 	$2, SI
	JL   	dotp_tail

	CMPB	·FmaSupt(SB), $1
	JE 	dotp_fma_loop

dotp_loop:
	MOVOU	(R8)(DI*8), X1
	MULPD	(R9)(DI*8), X1
	ADDPD	X1, X0
	ADDQ	$2, DI
	CMPQ	DI, SI
	JLE	dotp_loop
dotp_tail:
	ADDQ 	$1, SI
	CMPQ 	DI, SI
	JNE   	dotp_end
	MOVSD 	(R8)(DI*8), X1
	MULSD 	(R9)(DI*8), X1
	ADDSD 	X1, X0
	JMP 	dotp_end

dotp_fma_loop:
	MOVOU 	(R8)(DI*8), X1
	// VMFADD231PD X1, (R9)(DI*8), X0 (x0 += x1*(R9)
	BYTE	$0xC4; BYTE $0xC2; BYTE $0xF1; BYTE $0xB8; BYTE $0x04; BYTE $0xF9
	ADDQ	$2, DI
	CMPQ	DI, SI
	JLE	dotp_fma_loop
dotp_fma_tail:
	ADDQ	$1, SI
	CMPQ	DI, SI
	JNE 	dotp_end
	MOVSD	(R8)(DI*8), X1
	// VMFADD231SD X1, (R9)(DI*8), X0 (x0 += x1*x2)
	BYTE	$0xC4; BYTE $0xC2; BYTE $0xF1; BYTE $0xB9; BYTE $0x04; BYTE $0xF9
dotp_end:
	CMPB 	·Sse3Supt(SB), $1
	JE	dotp_sse3
	MOVAPD	X0, X1
	UNPCKHPD X1, X0
	ADDPD 	X1, X0
	MOVSD	X0, ret+48(FP)
	RET
dotp_sse3:
	BYTE $0x66; BYTE $0x0F; BYTE $0x7C; BYTE $0xC0
	//HADDPD X0, X0  //Added in 1.6
	MOVSD	X0, ret+48(FP)
	RET
