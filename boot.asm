section .boot
bits 16
global boot
boot:
	cli

	;enable A20 line to access more than 1MiB memory
	mov ax, 0x2401 ;pass param
	int 0x15       ;call BIOS function through interrupts


	;BIOS itself will print something unkown if commented the next two line code.
	mov ax, 0x3
	int 0x10      ;call BIOS function, set it to a known value


	mov ah, 0x2    ;read sectors
	mov al, 6      ;sectors to read
	mov ch, 0      ;cylinder idx
	mov dh, 0      ;head idx
	mov cl, 2      ;sector idx
	mov dl, [disk] ;disk idx
	mov bx, copy_target;target pointer
	int 0x13


	;enter Protected Mode to access 32bit instructions and registers
	lgdt [gdt_pointer]    ;load gdt table. it will set gdtr, a 48 bits register

	mov eax, cr0
	or eax,0x1
	mov cr0, eax          ;enable Protect Mode by set a special bit on CR0 register


    ;set the remaining segments to point at the data segment.
	mov ax, DATA_SEG
	mov ds, ax
	mov es, ax
	mov fs, ax
	mov gs, ax
	mov ss, ax
	jmp CODE_SEG:boot2
gdt_start:                ;the first entry of gdt table, 64 bits
	dq 0x0                ;0x00000000(4 bytes) as a null segment
gdt_code:                 ;the second entry, set the segment full 4G (flat mode)
	dw 0xFFFF             ;limit_low, set limit 0xFFFFF(20 Bit). multiplied by 4096
	dw 0x0                ;base_low, set base 0x00000000(4 bytes)
	db 0x0                ;base_middle
	db 10011010b          ;present(1B),ring level(2B),type(1B)
	                      ;executable(1B),direction(1B),readable/writable(1B),access(1B)
	db 11001111b          ;granularity(1B),size(1B),00b(2B),limit_high(4B)
	db 0x0                ;base_high
gdt_data:                 ;the same as code except...
	dw 0xFFFF
	dw 0x0
	db 0x0
	db 10010010b          ;except, the executable bit 0 means data selector
	db 11001111b
	db 0x0
gdt_end:
gdt_pointer:              ;set gdtr by lgdt
	dw gdt_end - gdt_start;16 bits, the size of gdt table
	dd gdt_start          ;the address of gdt table
disk:
	db 0x0
CODE_SEG equ gdt_code - gdt_start;code segment offset for later use
DATA_SEG equ gdt_data - gdt_start;data segment offset for later use

times 510 - ($-$$) db 0   ;fill zero
dw 0xaa55                 ;magic word,the 511b,512b must be 0xaa,0x55


copy_target:
bits 32
	hello: db "Hello more than 512 bytes world!",0
boot2:
	;print hello
	mov esi,hello
	mov ebx,0xb8000
.loop:
	lodsb
	or al,al
	jz boot3
	or eax,0x0F00
	mov word [ebx], ax
	add ebx,2
	jmp .loop

boot3:
	;goto the GO world
	mov esp,kernel_stack_top
	extern gocore.kernel.Main
	call gocore.kernel.Main

	cli
halt:
	hlt
	jmp halt


section .bss
align 4
kernel_stack_bottom: equ $
	resb 16384 ; 16 KiB
kernel_stack_top:
