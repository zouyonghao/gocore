void __go_runtime_error (int i){}

// a simple implementation just for compilation
int __builtin_memcmp(void *p1, void *p2, int len) {
    int* ip1 = p1;
    int *ip2 = p2;
    for (int i = 0; i < len; i++) {
        if (ip1[i] != ip2[i]) {
            return -1;
        }
    }
    return 0;
}