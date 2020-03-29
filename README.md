# Gocore

## setup
To compile gocore You need a gccgo cross-compiler

1. build a target i386/i686 gcc cross-compiler with go enabled, follow the article http://wiki.osdev.org/GCC_Cross-Compiler

2. install nasm from your repositories

3. install qemu for test

## compiler & run!

1. compile: `make GCC=i686-elf-gcc GCCGO=i686-elf-gccgo`, replace `GCC` and `GCCGO` with your binary name

2. run on qemu: `make run-qemu QEMU=qemu-system-i386`, replace `QEMU` with your target binary name

## 阅读材料

交叉编译

https://wiki.osdev.org/GCC_Cross-Compiler

我用的`binutils-2.34`和`gcc-7.5.0`，按照教程走就可以了，除了`gcc configure`这一步改成

```
../gcc-x.y.z/configure --target=$TARGET --prefix="$PREFIX" --disable-nls --enable-languages=c,c++,go --without-headers
```

可以再参考这里

https://wiki.osdev.org/Go_Bare_Bones

