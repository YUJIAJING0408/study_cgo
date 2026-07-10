#include "sdk.hpp"
#include <iostream>

int Calculator::add(int a, int b) {
    return a + b;
}

void Calculator::greet(const std::string& name) {
    std::cout << "Hello, " << name << " from C++!" << std::endl;
}