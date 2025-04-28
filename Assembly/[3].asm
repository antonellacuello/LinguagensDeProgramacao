%include "io.inc"

SECTION .data
n   dd 5
i   dd 1

SECTION .text
global CMAIN
CMAIN:

	mov edx, [n]
        mov ecx,[i]

loop_impares:
	cmp ecx, edx 		
	jg fim       		
	
	mov eax, ecx 		
	and eax, 1   		
	cmp eax, 0   		
	je pular_print

	PRINT_UDEC 4, ecx
	PRINT_CHAR 10

pular_print:
	inc ecx      		
	jmp loop_impares 	

fim:
	mov eax, 0
	ret          		