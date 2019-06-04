### 1、类和成员
- 一个类就是一个用户自定义类型，由一些内置类型、其他用户自定义类型和一些函数组成。这些用来定义类的组成部分称为成员
- 在成员函数中，一个成员名字表示调用该成员函数的对象中对应的成员，及下面var.mf()中的m表示var.m
```C++
class X{
public:
    int m;//数据成员
    int mf(int v){//函数成员
        int old = m;
        m = v; 
        return old;
    }
}
    X var;
    var.m = 7;
    int x = var.mf(9);
```

### 2、接口和实现
- 我们通常把类看做是接口+实现。接口是类声明的一部分，用户可以直接访问他，实现是类声明的另一部分，用户只能间接访问他。
- 公有的接口用public:标识，实现用private:来标识。所以，一个类声明可以理解为下面：
```C++
class X{
public:
    //共有成员:用户接口（可被所有人访问）
    //函数
    //类型
    //数据（通常最好为private）
private:
    //私有成员：实现细节（只能被类的成员访问）
    //函数
    //类型
    //数据
}
```
- 类成员默认是私有的
- 如果类只包含数据的话，接口和实现间的区别就没什么意义。**struct用来描述没有私有实现细节的类。一个结构体就是一个成员默认为公有属性的类**

### 4、类的演化
#### 类的构造
- 与类同名的构造函数是特殊的成员函数，称为构造函数。如果类有需要参数的构造函数，忘记用构造函数初始化对象，编译器会报错。
```C++
struct Date{
 int y, m, d;
 Date(int y, int m, int d);
 void add_day(int n);
};

Date my_birthday;//报错
Date last {2000, 2,14};//正确，口语风格
Date next = {2014, 2, 14};//正确
Date cri = Date{1976, 12, 24}；//正确，比较啰嗦
Date last(2000, 2, 14);//正确，旧的口语风格
```
- 因为C++中认为一个结构体就是成员默认为共有属性的类，所以可以在struct中定义成员函数，测试代码如下，clang编译器C++17标准编译通过
```C++
#include <iostream>
using namespace std;

struct Date{
    int y, m, d;
    Date(int a, int b, int c);
    void add_day(int n);
};

Date::Date(int a, int b, int c)
{
    y = a;
    m = b;
    d = c;
    return;
}

void Date::add_day(int n)
{
    d += n;
    return;
}

int main()
{
    Date a  {2000, 1, 1};
    cout << a.y << " " << a.m << " " << a.d << endl;
    a.add_day(1);
    cout << a.y << " " << a.m << " " << a.d << endl;

    return 0;
}
```

#### 保持细节私有

