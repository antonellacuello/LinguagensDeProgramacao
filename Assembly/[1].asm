%include "io.inc"

SECTION .data
fat    dd 5
res    dd 1

SECTION .text
global CMAIN

CMAIN:
    mov eax, 1
    mov ecx, [fat]

loop_fatorial:
    cmp ecx, 1
    jl fim
    imul eax, ecx
    dec ecx
    jmp loop_fatorial

fim:
    mov [res], eax
    PRINT_UDEC 4, eax
    mov eax, 0
    ret
