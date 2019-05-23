### Ch1
#### 1.1
- go run xxx.go 直接运行程序
- go build xxx.go 编译连接生成可执行文件
- 通过包来组织，文件开头用package声明属于哪个包。用import来导入包
- 同C语言等一样，从main函数开始执行
- **换行符会影响go代码解析（比如 { 不和func在同一行 会编译错误）**
- go fmt 可以格式化指定的文件中的代码。前提是代码可以编译过
- goimports可以按需管理导入，不是标准发布，用go get golang.org/x/tools/cmd/goimports 可以获得

#### 1.2
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
#### 1.3
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
