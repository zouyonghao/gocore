GCCGO := i686-elf-gccgo
GCC := i686-elf-gcc
NASM := nasm
QEMU := qemu

default: build

build: gocore.bin

gocore.bin: boot.o kernel.o runtime/libgo.so
	$(GCC) -T linker.ld -o gocore.bin -ffreestanding -nostdlib boot.o kernel.o runtime/libgo.so -lgcc

boot.o: boot.asm
	$(NASM) -felf32 boot.asm -o boot.o

kernel.o: kernel.go
	$(GCCGO) -c kernel.go -fgo-prefix=gocore

runtime/libgo.so: runtime/libgo.c
	cd runtime; \
	$(GCC) -shared -c libgo.c -o libgo.so -std=gnu99 -ffreestanding

run-qemu:
	$(QEMU) -kernel gocore.bin

clean:
	rm -f *.o
	rm -f **/*.so
	rm -f *.bin
