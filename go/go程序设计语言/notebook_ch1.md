## Ch1
### 1.1 Hello, World
- go run xxx.go 直接运行程序
- go build xxx.go 编译连接生成可执行文件
- 通过包来组织，文件开头用package声明属于哪个包。用import来导入包
- 同C语言等一样，从main函数开始执行
- **换行符会影响go代码解析（比如 { 不和func在同一行 会编译错误）**
- go fmt 可以格式化指定的文件中的代码。前提是代码可以编译过
- goimports可以按需管理导入，不是标准发布，用go get golang.org/x/tools/cmd/goimports 可以获得

### 1.2 Command-Line Arguments
- 切片，类似于Python，索引从0开始，s[i]表示第i+1个元素。切片不包含最后一个索引，s[i:j]表示第i+1个元素到第j个元素。s[:]等于s[0:len(s)]
- os.Args os.Args[0]是命令本身的名字，见下面的输出
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println(os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
/*
[root@bigdata go]# go run ch1echo1.go aa bb cc dd
/tmp/go-build908279927/b001/exe/ch1echo1
aa bb cc dd

[root@bigdata go]# go build ch1echo1.go
[root@bigdata go]# ./ch1echo1 aa bb cc dd
./ch1echo1
aa bb cc dd
*/
```

-  var声明变量，可以声明时初始化，**没有明确初始化会隐式初始化为空，变量类型写在变量名后面**
- for是唯一循环语句，**{ 必须和for在同一行，小括号不是必须的**
```go
//无限循环
for {

}
//循环10次
for i := 0; i < 10; i++ {

}
```
- **i++是语句不是表达式，j=i++不合法，且仅支持后缀，--i不合法**
- **:=用于短变量声明，会根据初始值给予合适的类型**
- range 产生一对值，索引和索引对应的值。如果某个值不想处理，可以用空标识符_（下划线）
```go
for _, arg := range os.Args[1:] {
    
}
```
- strings.Join(os.Args[1:], sep)  将一系列字符串连接为一个字符串，之间用sep来分隔
- var stringSlice [] string 新建切片

### 1.3 Finding Duplicate Lines
- map存储键值对，提供常数级别的检索。key值需要是能用==进行比较的类型，value可以是任意类型。make函数用来创建一个空的map
```go
//counts 是key为string类型 value为int类型的map
counts := make(map[string]int)

counts[input.Text()]++ //key不在map中时会自动插入，value为0值
//上面的代码等价于如下两行
line := input.Text()
counts[line] = counts[line] + 1
```
- range map迭代的顺序是随机的，这一点同其他语言差不多
```go
for line, n := range counts {
}
```
- bufio 处理输入输出的包，最常用的特性是Scanner，读取输入，并拆分成行或者单词。是处理行输入的最简单的方式。Scanner从标准输入读取，每次调用input.Scan()读取一行，并把最后的换行符删掉，读取的结果可以通过调用input.Text()。Scan()函数在有输入时返回true，没有输入时返回false。

```go
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
    counts[input.Text()]++
}
```
- fmt.Printf函数和C语言中的printf函数差不多，通过一些表达式来格式化输出。第一个参数是格式化字符串，用来指定后面的参数如何被格式化，每个参数由%+字符来指定，%d输出十进制整数，%s是字符串，这些转换go程序员称之为verbs，下面列出一些常见的转换
```go
%d 十进制整形数
%x, %o, %b 十六进制、八进制、二进制整数
%f, %g, %e 浮点数如: 3.141593 3.141592653589793 3.141593e+00
%t 布尔值: true or false
%c 字符 (Unicodecode point)
%s 字符串
%q 带引号的字符串 "abc" or rune 'c'
%v 内置的任何类型
%T 任何类型
%% 真正的百分号
```
- 字符串字面值可以包含\t（tab）\n（换行）等转义序列（escape sequences）来表示不可见字符。Printf函数默认不会换行。f结尾的函数如log.Printf,fmt.Errorf使用和fmt.Printf一样的格式化策略。ln结尾（Println）以%v处理参数，并且最后会换行。
```go
fmt.Printf("%d\t%s\n", n, line)
```
- os.Open返回两个值，第一个是一个打开的文件（\*os.File）,这个文件后面会被Scanner读取，第二个返回值是内置的错误类型，如果返回值等于内置类型nil，则文件成功打开，当返回值不为nil时，文件打开出错，错误码中描述的是错误内容。文件读取到末尾时，Close函数会关闭文件，并释放资源。一种简单的处理是Fprintf和%v输出一段信息到标准err流，错误处理后，dup会处理下一个文件。
```go
f, err := os.Open(arg)
if err != nil {
fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
continue
}
```

- 函数和其他包级别的实体可以用任意顺序声明。
- map是make创建的数据结构的引用（reference）。当map传递给一个函数时，函数接收到那个引用的一份拷贝，所以被调用的函数对底层数据结构的修改对外层的引用来说也是可见的

- ioutil.ReadFile(在io/ioutil包内)返回能转换成字符串的字节切片，以便于能被strings.Split切割，
```go
data, err := ioutil.ReadFile(filename)
if err != nil {
fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
}
for _, line := range strings.Split(string(data), "\n") {//range切片返回索引和值
counts[line]++
}
```
