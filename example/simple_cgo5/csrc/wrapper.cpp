#include "sdk.hpp"
#include <string>

// 使用 extern "C" 阻止 C++ 名称修饰，使函数按 C 方式导出
extern "C" {

// 定义一个不透明指针（隐藏 C++ 对象细节）
typedef void* CalculatorHandle;

// 创建 Calculator 对象，返回句柄
CalculatorHandle Calculator_new() {
    return new Calculator();
}

// 删除对象
void Calculator_delete(CalculatorHandle handle) {
    delete static_cast<Calculator*>(handle);
}

// 调用 add 方法
int Calculator_add(CalculatorHandle handle, int a, int b) {
    Calculator* calc = static_cast<Calculator*>(handle);
    return calc->add(a, b);
}

// 调用 greet 方法（接收 C 字符串）
void Calculator_greet(CalculatorHandle handle, const char* name) {
    Calculator* calc = static_cast<Calculator*>(handle);
    std::string str(name);
    calc->greet(str);
}

} // extern "C"