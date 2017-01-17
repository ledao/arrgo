//+build !noasm,!appengine

#define NOSPLIT 7

TEXT ·findBool(SB), NOSPLIT, $0	
	MOVQ vals_base+0(FP), R8
	MOVQ vals_len+8(FP), SI
	MOVB find+24(FP), R10
	PXOR X0,X0

	CMPQ SI, $0
	JE failed
	
	CMPB R10, $0
	JE loop
	
	MOVQ R10, X0
	PUNPCKLBW X0, X0
	PSHUFLW $0, X0, X0
	PUNPCKLQDQ X0, X0
loop:
	MOVOU X0, X1
	PCMPEQB (R8), X1
	PMOVMSKB X1, R9
	BSFL R9,R9
	JNZ fnd
	ADDQ $16, R8
	SUBQ $16, SI
	JG loop
	JMP failed
fnd:
	CMPQ R9, SI
	JG failed
	MOVB $1, flg+32(FP)
	RET
failed:
	MOVB $0, flg+32(FP)
	RET

	
