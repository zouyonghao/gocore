GCCGO := i686-elf-gccgo
GCC := i686-elf-gcc
NASM := nasm
QEMU := qemu-system-i386

GOFLAGS_CROSS = -static -Werror -nostdlib

default: build

build: gocore.bin

gocore.bin: boot.o kernel.o runtime/libgo.so runtime/runtime.o
	$(GCC) -T linker.ld -o gocore.bin -ffreestanding -nostdlib boot.o runtime/runtime.o kernel.o runtime/libgo.so -lgcc

boot.o: boot.asm
	$(NASM) -felf32 boot.asm -g -o boot.o

kernel.o: kernel.go
	$(GCCGO) -c kernel.go kernel/*.go -fgo-prefix=gocore

runtime/runtime.o: runtime/runtime.go
	cd runtime; \
	$(GCCGO) $(GOFLAGS_CROSS) -c runtime.go

runtime/libgo.so: runtime/libgo.c
	cd runtime; \
	$(GCC) -shared -c libgo.c -o libgo.so -std=gnu99 -ffreestanding

debug: gocore.bin
	gnome-terminal -- /bin/bash -c "$(QEMU) -S -s -d in_asm -D q.log -parallel stdio -hda $< -serial null"
	sleep 2
	gnome-terminal -- /bin/bash -c "gdb -q -tui -x tools/gdbinit"

run-qemu:
	$(QEMU) -fda gocore.bin

clean:
	rm -f *.o
	rm -f **/*.o
	rm -f **/*.s
	rm -f q.log
	rm -f **/*.so
	rm -f *.bin
