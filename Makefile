### Build params

CC_CROSS = i686-elf-gcc
LD_CROSS = i686-elf-ld
GO_CROSS = i686-elf-gccgo
OBJCOPY = i686-elf-objcopy
PREPROC = $(CC_CROSS) -E -x c -P
CC = gcc
LD = ld
ASM = nasm -f elf
CFLAGS_CROSS = -Werror -nostdlib -fno-builtin -nostartfiles -nodefaultlibs
GOFLAGS_CROSS = -static  -Werror -nostdlib -nostartfiles -nodefaultlibs 
INCLUDE_DIRS = -I.
QEMU = qemu-system-i386

### Sources

CORE_SOURCES = loader.o interupts.o asm.o asm.go.o asm.gox regs.go.o regs.gox ptr.go.o ptr.gox color.go.o color.gox video.go.o video.gox gdt.go.o gdt.gox idt.go.o idt.gox pit.go.o pit.gox kbd.go.o kbd.gox runtime.go.o libgo.o kernel.go.o

SOURCE_OBJECTS = $(CORE_SOURCES)
 
### Targets

all: kernel.bin

clean:
	rm -f $(SOURCE_OBJECTS) $(TEST_EXECS) kernel.bin kernel.o

boot: kernel.bin
	$(QEMU) -kernel kernel.bin -m 1024

boot-iso: kernel.iso
	qemu-system-i386 -cdrom kernel.iso

### Rules

%.o: %.s
	$(ASM) $(INCLUDE_DIRS) -o $@ $<

%.gox: %.go.o
	$(OBJCOPY) -j .go_export $< $@

%.go.o: %.go
	$(GO_CROSS) $(GOFLAGS_CROSS) $(INCLUDE_DIRS) -o $@ -c $<

libgo.o: asm.o
	$(CC_CROSS) $(CFLAGS_CROSS) libgo.c -o libgo.o

test.o: asm.o
	$(CC_CROSS) $(CFLAGS_CROSS) test.c -o test.o

kernel.bin: $(SOURCE_OBJECTS)
	$(LD_CROSS) -T link.ld -o kernel.bin $(SOURCE_OBJECTS)

debug: kernel.bin
	gnome-terminal -- /bin/bash -c "$(QEMU) -S -s -d in_asm -D q.log -parallel stdio -kernel kernel.bin -serial null"
	sleep 2
	gnome-terminal -- /bin/bash -c "gdb -q -tui -x tools/gdbinit"

kernel.iso: kernel.bin
	cp kernel.bin isodir/boot/kernel.bin
	grub-mkrescue -o kernel.iso isodir
