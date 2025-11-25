/**
 * C++ Basics - Variables và Data Types
 */
#include <iostream>
#include <string>

using namespace std;

int main() {
    // 1. Primitive Data Types
    
    // Integer types
    int integer = 42;
    short shortInt = 32767;
    long longInt = 2147483647L;
    long long longLongInt = 9223372036854775807LL;
    
    // Floating point types
    float floatNum = 3.14f;
    double doubleNum = 3.14159;
    long double longDouble = 3.141592653589793238L;
    
    // Character types
    char character = 'A';
    wchar_t wideChar = L'A';
    
    // Boolean type
    bool isTrue = true;
    bool isFalse = false;
    
    // 2. String
    string str = "Hello, C++!";
    
    // 3. Arrays
    int numbers[5] = {1, 2, 3, 4, 5};
    int numbers2[] = {1, 2, 3};
    
    // 4. Pointers
    int x = 10;
    int* ptr = &x;
    cout << "Value: " << x << endl;
    cout << "Address: " << ptr << endl;
    cout << "Value via pointer: " << *ptr << endl;
    
    // 5. References
    int y = 20;
    int& ref = y;
    ref = 30;
    cout << "y = " << y << endl;  // y is now 30
    
    // 6. Constants
    const int MAX_SIZE = 100;
    const double PI = 3.14159;
    
    // 7. Auto keyword (C++11)
    auto autoInt = 42;
    auto autoDouble = 3.14;
    auto autoString = "Hello";
    
    // 8. Output
    cout << "Integer: " << integer << endl;
    cout << "Double: " << doubleNum << endl;
    cout << "String: " << str << endl;
    cout << "Boolean: " << isTrue << endl;
    
    return 0;
}

