// Code generated by command: go run fe_amd64_asm.go -out ../fe_amd64.s -stubs ../fe_amd64.go. DO NOT EDIT.

//go:build amd64 && gc && !purego

#include "textflag.h"

// func feMul(out *Element, a *Element, b *Element)
TEXT ·feMul(SB), NOSPLIT, $0-24
	MOVQ a+8(FP), CX
	MOVQ b+16(FP), BX

	// r0 = a0√ób0
	MOVQ (CX), AX
	MULQ (BX)
	MOVQ AX, DI
	MOVQ DX, SI

	// r0 += 19√óa1√ób4
	MOVQ   8(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   32(BX)
	ADDQ   AX, DI
	ADCQ   DX, SI

	// r0 += 19√óa2√ób3
	MOVQ   16(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   24(BX)
	ADDQ   AX, DI
	ADCQ   DX, SI

	// r0 += 19√óa3√ób2
	MOVQ   24(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   16(BX)
	ADDQ   AX, DI
	ADCQ   DX, SI

	// r0 += 19√óa4√ób1
	MOVQ   32(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   8(BX)
	ADDQ   AX, DI
	ADCQ   DX, SI

	// r1 = a0√ób1
	MOVQ (CX), AX
	MULQ 8(BX)
	MOVQ AX, R9
	MOVQ DX, R8

	// r1 += a1√ób0
	MOVQ 8(CX), AX
	MULQ (BX)
	ADDQ AX, R9
	ADCQ DX, R8

	// r1 += 19√óa2√ób4
	MOVQ   16(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   32(BX)
	ADDQ   AX, R9
	ADCQ   DX, R8

	// r1 += 19√óa3√ób3
	MOVQ   24(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   24(BX)
	ADDQ   AX, R9
	ADCQ   DX, R8

	// r1 += 19√óa4√ób2
	MOVQ   32(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   16(BX)
	ADDQ   AX, R9
	ADCQ   DX, R8

	// r2 = a0√ób2
	MOVQ (CX), AX
	MULQ 16(BX)
	MOVQ AX, R11
	MOVQ DX, R10

	// r2 += a1√ób1
	MOVQ 8(CX), AX
	MULQ 8(BX)
	ADDQ AX, R11
	ADCQ DX, R10

	// r2 += a2√ób0
	MOVQ 16(CX), AX
	MULQ (BX)
	ADDQ AX, R11
	ADCQ DX, R10

	// r2 += 19√óa3√ób4
	MOVQ   24(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   32(BX)
	ADDQ   AX, R11
	ADCQ   DX, R10

	// r2 += 19√óa4√ób3
	MOVQ   32(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   24(BX)
	ADDQ   AX, R11
	ADCQ   DX, R10

	// r3 = a0√ób3
	MOVQ (CX), AX
	MULQ 24(BX)
	MOVQ AX, R13
	MOVQ DX, R12

	// r3 += a1√ób2
	MOVQ 8(CX), AX
	MULQ 16(BX)
	ADDQ AX, R13
	ADCQ DX, R12

	// r3 += a2√ób1
	MOVQ 16(CX), AX
	MULQ 8(BX)
	ADDQ AX, R13
	ADCQ DX, R12

	// r3 += a3√ób0
	MOVQ 24(CX), AX
	MULQ (BX)
	ADDQ AX, R13
	ADCQ DX, R12

	// r3 += 19√óa4√ób4
	MOVQ   32(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   32(BX)
	ADDQ   AX, R13
	ADCQ   DX, R12

	// r4 = a0√ób4
	MOVQ (CX), AX
	MULQ 32(BX)
	MOVQ AX, R15
	MOVQ DX, R14

	// r4 += a1√ób3
	MOVQ 8(CX), AX
	MULQ 24(BX)
	ADDQ AX, R15
	ADCQ DX, R14

	// r4 += a2√ób2
	MOVQ 16(CX), AX
	MULQ 16(BX)
	ADDQ AX, R15
	ADCQ DX, R14

	// r4 += a3√ób1
	MOVQ 24(CX), AX
	MULQ 8(BX)
	ADDQ AX, R15
	ADCQ DX, R14

	// r4 += a4√ób0
	MOVQ 32(CX), AX
	MULQ (BX)
	ADDQ AX, R15
	ADCQ DX, R14

	// First reduction chain
	MOVQ   $0x0007ffffffffffff, AX
	SHLQ   $0x0d, DI, SI
	SHLQ   $0x0d, R9, R8
	SHLQ   $0x0d, R11, R10
	SHLQ   $0x0d, R13, R12
	SHLQ   $0x0d, R15, R14
	ANDQ   AX, DI
	IMUL3Q $0x13, R14, R14
	ADDQ   R14, DI
	ANDQ   AX, R9
	ADDQ   SI, R9
	ANDQ   AX, R11
	ADDQ   R8, R11
	ANDQ   AX, R13
	ADDQ   R10, R13
	ANDQ   AX, R15
	ADDQ   R12, R15

	// Second reduction chain (carryPropagate)
	MOVQ   DI, SI
	SHRQ   $0x33, SI
	MOVQ   R9, R8
	SHRQ   $0x33, R8
	MOVQ   R11, R10
	SHRQ   $0x33, R10
	MOVQ   R13, R12
	SHRQ   $0x33, R12
	MOVQ   R15, R14
	SHRQ   $0x33, R14
	ANDQ   AX, DI
	IMUL3Q $0x13, R14, R14
	ADDQ   R14, DI
	ANDQ   AX, R9
	ADDQ   SI, R9
	ANDQ   AX, R11
	ADDQ   R8, R11
	ANDQ   AX, R13
	ADDQ   R10, R13
	ANDQ   AX, R15
	ADDQ   R12, R15

	// Store output
	MOVQ out+0(FP), AX
	MOVQ DI, (AX)
	MOVQ R9, 8(AX)
	MOVQ R11, 16(AX)
	MOVQ R13, 24(AX)
	MOVQ R15, 32(AX)
	RET

// func feSquare(out *Element, a *Element)
TEXT ·feSquare(SB), NOSPLIT, $0-16
	MOVQ a+8(FP), CX

	// r0 = l0√ól0
	MOVQ (CX), AX
	MULQ (CX)
	MOVQ AX, SI
	MOVQ DX, BX

	// r0 += 38√ól1√ól4
	MOVQ   8(CX), AX
	IMUL3Q $0x26, AX, AX
	MULQ   32(CX)
	ADDQ   AX, SI
	ADCQ   DX, BX

	// r0 += 38√ól2√ól3
	MOVQ   16(CX), AX
	IMUL3Q $0x26, AX, AX
	MULQ   24(CX)
	ADDQ   AX, SI
	ADCQ   DX, BX

	// r1 = 2√ól0√ól1
	MOVQ (CX), AX
	SHLQ $0x01, AX
	MULQ 8(CX)
	MOVQ AX, R8
	MOVQ DX, DI

	// r1 += 38√ól2√ól4
	MOVQ   16(CX), AX
	IMUL3Q $0x26, AX, AX
	MULQ   32(CX)
	ADDQ   AX, R8
	ADCQ   DX, DI

	// r1 += 19√ól3√ól3
	MOVQ   24(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   24(CX)
	ADDQ   AX, R8
	ADCQ   DX, DI

	// r2 = 2√ól0√ól2
	MOVQ (CX), AX
	SHLQ $0x01, AX
	MULQ 16(CX)
	MOVQ AX, R10
	MOVQ DX, R9

	// r2 += l1√ól1
	MOVQ 8(CX), AX
	MULQ 8(CX)
	ADDQ AX, R10
	ADCQ DX, R9

	// r2 += 38√ól3√ól4
	MOVQ   24(CX), AX
	IMUL3Q $0x26, AX, AX
	MULQ   32(CX)
	ADDQ   AX, R10
	ADCQ   DX, R9

	// r3 = 2√ól0√ól3
	MOVQ (CX), AX
	SHLQ $0x01, AX
	MULQ 24(CX)
	MOVQ AX, R12
	MOVQ DX, R11

	// r3 += 2√ól1√ól2
	MOVQ   8(CX), AX
	IMUL3Q $0x02, AX, AX
	MULQ   16(CX)
	ADDQ   AX, R12
	ADCQ   DX, R11

	// r3 += 19√ól4√ól4
	MOVQ   32(CX), AX
	IMUL3Q $0x13, AX, AX
	MULQ   32(CX)
	ADDQ   AX, R12
	ADCQ   DX, R11

	// r4 = 2√ól0√ól4
	MOVQ (CX), AX
	SHLQ $0x01, AX
	MULQ 32(CX)
	MOVQ AX, R14
	MOVQ DX, R13

	// r4 += 2√ól1√ól3
	MOVQ   8(CX), AX
	IMUL3Q $0x02, AX, AX
	MULQ   24(CX)
	ADDQ   AX, R14
	ADCQ   DX, R13

	// r4 += l2√ól2
	MOVQ 16(CX), AX
	MULQ 16(CX)
	ADDQ AX, R14
	ADCQ DX, R13

	// First reduction chain
	MOVQ   $0x0007ffffffffffff, AX
	SHLQ   $0x0d, SI, BX
	SHLQ   $0x0d, R8, DI
	SHLQ   $0x0d, R10, R9
	SHLQ   $0x0d, R12, R11
	SHLQ   $0x0d, R14, R13
	ANDQ   AX, SI
	IMUL3Q $0x13, R13, R13
	ADDQ   R13, SI
	ANDQ   AX, R8
	ADDQ   BX, R8
	ANDQ   AX, R10
	ADDQ   DI, R10
	ANDQ   AX, R12
	ADDQ   R9, R12
	ANDQ   AX, R14
	ADDQ   R11, R14

	// Second reduction chain (carryPropagate)
	MOVQ   SI, BX
	SHRQ   $0x33, BX
	MOVQ   R8, DI
	SHRQ   $0x33, DI
	MOVQ   R10, R9
	SHRQ   $0x33, R9
	MOVQ   R12, R11
	SHRQ   $0x33, R11
	MOVQ   R14, R13
	SHRQ   $0x33, R13
	ANDQ   AX, SI
	IMUL3Q $0x13, R13, R13
	ADDQ   R13, SI
	ANDQ   AX, R8
	ADDQ   BX, R8
	ANDQ   AX, R10
	ADDQ   DI, R10
	ANDQ   AX, R12
	ADDQ   R9, R12
	ANDQ   AX, R14
	ADDQ   R11, R14

	// Store output
	MOVQ out+0(FP), AX
	MOVQ SI, (AX)
	MOVQ R8, 8(AX)
	MOVQ R10, 16(AX)
	MOVQ R12, 24(AX)
	MOVQ R14, 32(AX)
	RET
