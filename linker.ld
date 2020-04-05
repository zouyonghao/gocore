/* 对应boot.asm中的 boot */
ENTRY(boot)
OUTPUT_FORMAT("binary")
SECTIONS {
	. = 0x7c00; /* the BIOS will jump to location 0x7c00 */
	.text :
	{
        *(.boot)
		*(.text)
	}
    
    .rodata :
    {
    	*(.rodata)
    }

    .data :
    {
    	*(.data)
    }

    .bss :
    {
    	*(.bss)
    }
}

