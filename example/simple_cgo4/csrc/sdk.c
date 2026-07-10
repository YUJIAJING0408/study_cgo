#include <stdio.h>
#include "api.h"

int add(int a, int b) {
    return a + b;
}

void greet(const char* name) {
    printf("Hello, %s from C!\n", name);
}