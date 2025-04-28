%include "io.inc"

SECTION .data
n           dd 17              
divisor     dd 2               
primo       dd 1                  

SECTION .text
global CMAIN
CMAIN:

    mov eax, [n]             
    cmp eax, 2                 
    JL nao_primo

verificar_divisores:
    mov ebx, [divisor]       
    cmp ebx, eax               
    jge fim_teste            

    mov edx, 0                 
    mov ecx, ebx              
    div ecx                    

    cmp edx, 0                 
    je nao_primo               

    mov ebx, [divisor]         
    add ebx, 1                 
    mov [divisor], ebx         
    mov eax, [n]               
    jmp verificar_divisores    

nao_primo:
    mov dword [primo], 0       
    jmp fim_teste    

fim_teste:
    PRINT_UDEC 4, [primo]   ; se 1 -> primo, se 0 -> não é primo
    PRINT_CHAR 10            

    mov eax, 0               
    ret