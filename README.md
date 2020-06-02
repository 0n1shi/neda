# NEDA (NEs DisAssembler)

A disassembler for NES ROM which written in Golang (still not perfect...).

## Usage

```bash
$ neda -h
Name: NEDA (NEs rom DisAssembler)
Version: 0.7.0
Usage of neda:
        neda [Options ...]
Options
  -header-only
        display header only
  -rom string
        rom file path
```

## Example

Outout will be like below.

```asm
magic number: NES
program bank count: 2 (32768 bytes)
character bank count: 1 (16384 bytes)
mapper type: 0 (NROM)

0x8000: 78            sei
0x8001: D8            cld
0x8002: A9 10         lda #$10
0x8004: 8D 00 20      sta $2000
0x8007: A2 FF         ldx #$FF
0x8009: 9A            txs
0x800A: AD 02 20      lda $2002
0x800D: 10 FB         bpl $8108
0x800F: AD 02 20      lda $2002
0x8012: 10 FB         bpl $810D
0x8014: A0 FE         ldy #$FE
0x8016: A2 05         ldx #$05
0x8018: BD D7 07      lda $07D7, X
0x801B: C9 0A         cmp #$0A
0x801D: B0 0C         bcs $8029
0x801F: CA            dex
0x8020: 10 F6         bpl $8116
0x8022: AD FF 07      lda $07FF
0x8025: C9 A5         cmp #$A5
0x8027: D0 02         bne $8029
0x8029: A0 D6         ldy #$D6
0x802B: 20 CC 90      jsr $90CC
0x802E: 8D 11 40      sta $4011
0x8031: 8D 70 07      sta $0770
0x8034: A9 A5         lda #$A5
0x8036: 8D FF 07      sta $07FF
0x8039: 8D A7 07      sta $07A7
0x803C: A9 0F         lda #$0F
0x803E: 8D 15 40      sta $4015
0x8041: A9 06         lda #$06
0x8043: 8D 01 20      sta $2001
0x8046: 20 20 82      jsr $8220
0x8049: 20 19 8E      jsr $8E19
0x804C: EE 74 07      inc $0774
0x804F: AD 78 07      lda $0778
0x8052: 09 80         ora #$80
0x8054: 20 ED 8E      jsr $8EED
0x8057: 4C 57 80      jmp $8057

0x805A: 01 A4         ora ($A4, X)
0x805C: C8            iny
0x805D: EC 10 00      cpx $0010
0x8060: 41 41         eor ($41, X)
0x8062: 4C 34 3C      jmp $3C34

0x8065: 44 54 68 7C   .db $44 $54 $68 $7C
0x8069: A8 BF DE EF   .db $A8 $BF $DE $EF
0x806D: 03 8C 8C 8C   .db $03 $8C $8C $8C
0x8071: 8D 03 03 03   .db $8D $03 $03 $03
0x8075: 8D 8D 8D 8D   .db $8D $8D $8D $8D
0x8079: 8D 8D 8D 8D   .db $8D $8D $8D $8D
0x807D: 8D 8D 8D 00   .db $8D $8D $8D $00
0x8081: 40            .db $40

0x8082: AD 78 07      lda $0778
0x8085: 29 7F         and #$7F
0x8087: 8D 78 07      sta $0778
0x808A: 29 7E         and #$7E
0x808C: 8D 00 20      sta $2000
0x808F: AD 79 07      lda $0779

:
```
