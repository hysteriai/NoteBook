在赋值语句中，顶层const属性可以忽略，底层const不能忽略，双方必须有相同的底层const属性，或者底层const属性可以转换（非常量可以转换成常量，反之则不行）。

通俗来说，顶层const表示变量自己是const，不能修改当前这个变量的值。底层const表示当前变量所指向（引用）的值是一个不能修改的const值。

代码示例：
- 顶层const
```c++
const int i = 1;  //顶层const，i的值不能改变
i = 3;     //报错

int a = 1;
int b = 4;
int *const p = &a;  //p是顶层const，即p的值不能更改，但是可以通过p修改a那块地址的值
p = &b;  //错误
*p = b;  //正确，a的值被改为4
```

- 底层const
```c++
int i = 5;
const int *p = &i;  //p是底层const，p的值可以改变，当然i的值也能改变，但是不能通过p改变i的值
*p = 3; //错误，不能通过p改i的值
p = 0; //正确，p不是顶层const
i = 6; //正确

```
- 即是顶层又是底层const
```c++
int i = 10;
const int *const p = &i;  //p的值不能改变，也不能通过p改变所指向地址的值，前面的const是底层const后面的const是顶层const
```

- 赋值操作顶层const不受什么影响
```c++
int i = 10;
const int j = i;//j是顶层const
int k = j;//正确
```
- 底层const不能忽略，双方必须有相同的底层const属性，或者底层const属性可以转换（非常量可以转换成常量，反之则不行）。
声明引用的const都是底层const
```c++
const int i = 10; 
const int *p = &i;//正确，&i具有底层const属性，p也有底层const属性
int *const p2 = &i; //错误，p2有顶层const属性，但是没有底层const属性
int &a = i;  //错误，不能将普通的int绑定到const int 上
const int &b = i;  //正确



int k = 1;
const int &j = k;//正确，底层const可以从无到有, 不能通过j改变k的值，但是k的值可以改变
const int *p3 = &k;//正确，底层const可以从无到有
```
