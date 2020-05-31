// a simple implementation just for compilation
int __builtin_memcmp(void *cs, void *ct, unsigned int count)
{
    const unsigned char *su1, *su2;
    int res = 0;

    for (su1 = cs, su2 = ct; 0 < count; ++su1, ++su2, count--)
        if ((res = *su1 - *su2) != 0)
            break;
    return res;
}

// static unsigned mem = (unsigned)0x100000;
// extern unsigned __kernel_end();

// void *sys_malloc(long unsigned int size)
// {
//     static char *p;
//     void *q;

//     if (p == 0)
//     {
//         p = (char *)((unsigned)__kernel_end() & 0xFFFFF000) + 0x1000;
//         //p += 7 & -(uintptr)p;
//     }

//     size += 7 & -size;

//     q = p;
//     p += size;
//     return q;
// }