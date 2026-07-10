#ifndef WRAPPER_H
#define WRAPPER_H

// 与 wrapper.cpp 中的声明一致（C 接口）
typedef void* CalculatorHandle;

CalculatorHandle Calculator_new();
void Calculator_delete(CalculatorHandle handle);
int Calculator_add(CalculatorHandle handle, int a, int b);
void Calculator_greet(CalculatorHandle handle, const char* name);

#endif