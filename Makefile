GCCGO := i686-elf-gccgo
GCC := i686-elf-gcc
NASM := nasm
QEMU := qemu-system-i386

default: build

build: gocore.bin

gocore.bin: boot.o kernel.o runtime/libgo.so runtime/alg.o
	$(GCC) -T linker.ld -o gocore.bin -ffreestanding -nostdlib boot.o kernel.o runtime/libgo.so runtime/alg.o -lgcc

boot.o: boot.asm
	$(NASM) -felf32 boot.asm -o boot.o

kernel.o: kernel.go
	$(GCCGO) -c kernel.go mm/pmm.go -fgo-prefix=gocore

runtime/runtime.o: runtime/alg.go
	cd runtime; \
	$(GCCGO) -c alg.go

runtime/libgo.so: runtime/libgo.c
	cd runtime; \
	$(GCC) -shared -c libgo.c -o libgo.so -std=gnu99 -ffreestanding

run-qemu:
	$(QEMU) -fda gocore.bin

clean:
	rm -f *.o
	rm -f **/*.so
	rm -f *.bin
