# Go 语言学习路线（零基础 · 内容自包含版）

> 本路线以菜鸟教程的知识体系为底座，并把 Go 官方教程（入门、创建模块、泛型、数据库、
> Gin Web 服务）的实际内容融合进各阶段正文，教程中引用而未展开的部分也已补全——
> 全文读完不需要再去查任何外部资料。
> 预计总时长：6~7 周（每天 1~2 小时）；阶段七至九内容较重，节奏可自行放缓。

---

## 阶段一：跑起来再说 —— 环境与工具链（2~3 天）

**目标**：搭好环境，理解 Go "编译型语言" 的工作方式，并跑通第一个模块化项目。

### 1. 安装与环境变量
- 官网下载安装包（Mac 也可用 `brew install go`），安装后执行 `go version` 验证
- 直接装最新稳定版即可：本路线按 Go 1.22+ 的语言行为撰写（如循环变量按迭代独立），所有示例在新版上可直接运行
- 两个重要环境变量：
  - `GOROOT`：Go 安装目录（编译器、标准库所在）
  - `GOPATH`：工作目录；`go install` 编译出的可执行程序会放在 `$GOPATH/bin`（默认 `~/go/bin`）

### 2. 必会命令

| 命令 | 作用 |
|---|---|
| `go run .` | 编译并运行当前目录的程序（不保留二进制） |
| `go build .` | 编译生成可执行文件 |
| `go install .` | 编译并把可执行文件安装到 `$GOPATH/bin`，之后可在任意目录直接敲命令运行 |
| `go mod init <模块路径>` | 初始化模块，生成 go.mod |
| `go mod tidy` | 自动下载/补全/清理依赖（最常用） |
| `go fmt ./...` | 统一代码格式（Go 强制统一风格，没有格式之争） |

### 3. 第一个程序（单文件版，先热身）
```go
package main    // 每个 Go 文件必须属于一个包；main 包是程序入口包

import "fmt"    // 导入标准库的 fmt 包（格式化输入输出）

func main() {   // main 函数是程序入口，无参数、无返回值
	fmt.Println("Hello, World!")
}
```
要点：`main` 包 + `main` 函数二者缺一不可；未被使用的 import 会**编译报错**（不是警告）。

### 4. 正式的起步方式：模块化 Hello World
官方推荐的姿势不是裸跑单文件，而是先创建模块（模块 = 用一个 go.mod 管理的一组包，是依赖管理的最小单位）：
```bash
$ mkdir hello && cd hello
$ go mod init example/hello
go: creating new go.mod: module example/hello
```
> 模块路径的命名：真实项目中通常是代码仓库地址（如 `github.com/mymodule`），这样别人才能用 Go 工具下载你的模块。练习时随意取（如 `example/hello`）即可。

### 5. 调用外部模块的代码
Go 生态的包都发布在 pkg.go.dev 上（官方的包发现站点），可以直接搜索想用的功能。来试试官方教程的经典例子——引用 `rsc.io/quote` 模块：

```go
package main

import (
	"fmt"
	"rsc.io/quote"   // 外部模块里的 quote 包
)

func main() {
	fmt.Println(quote.Go())   // 调用该包的导出函数 Go()
}
```
执行 `go mod tidy`（它会自动找到并下载 rsc.io/quote，同时生成 go.sum 校验文件），然后运行：
```bash
$ go mod tidy
go: finding module for package rsc.io/quote
go: found rsc.io/quote in rsc.io/quote v1.5.2

$ go run .
Don't communicate by sharing memory, share memory by communicating.
```
输出的这句话正是 Go 的并发哲学——它出自 Go 谚语，在阶段八还会见到它。

### 练习
- [ ] 跑通两种 Hello World（单文件版 + 模块版），并用 `go build` 生成二进制后脱离源码直接运行
- [ ] 引入 `rsc.io/quote`，试试该包的其他函数（如 quote.Hello()）
- [ ] 把缩进改成 tab/空格混用，执行 `go fmt` 观察自动修正

### 过关自检
- `go run`、`go build`、`go install` 三者的区别？
- go.mod 是干什么的？`go mod tidy` 什么时候用？

---

## 阶段二：语言的骨架 —— 基本语法与数据（第 1 周）

**目标**：能读懂并写出顺序执行的简单程序。

### 1. 基础语法规则
- 行分隔符：一行代表一个语句结束，**结尾的分号可省略**；打算将多条语句写在一行才用 `;` 分隔
- 注释：单行 `//`，多行 `/* ... */`（不可嵌套）
- 标识符：字母/数字/下划线组成，不能以数字开头；**区分大小写**
- 空格有意义的场景：`fruit=apples+oranges` 不留空格会被判定为标识符错误

### 2. 数据类型

| 分类 | 类型 | 说明 |
|---|---|---|
| 布尔 | `bool` | 只有 true / false |
| 整型 | `int8/16/32/64`、`uint8/16/32/64`、`int/uint` | int 大小跟平台有关（64 位系统为 64 位） |
| 浮点 | `float32` / `float64` | 默认推荐 float64 |
| 复数 | `complex64` / `complex128` | 科学计算用 |
| 字符 | `byte`（=uint8，ASCII）、`rune`（=int32，UTF-8 字符） | 处理中文必须用 rune |
| 字符串 | `string` | UTF-8 编码、**不可变**；可用 `+` 拼接 |
| 其他 | `uintptr`（存指针值）、派生类型（指针/数组/切片/结构体/函数/接口/通道） | 后续阶段逐一学 |

### 3. 变量
```go
var a int = 10      // 指定类型并赋值
var b = 10          // 类型推断
c := 10             // 短声明：只能在函数内使用
var x, y int = 1, 2 // 多变量声明
var (               // 因式分解式声明
	e int
	f string
)
```
- 零值机制：只声明不赋值时有默认值——数值 `0`、布尔 `false`、字符串 `""`、指针 `nil`
- `:=` 的本质是"声明 + 初始化"二合一，等价于先 `var message string` 再赋值
- 值类型 vs 引用类型：值类型（int/数组/结构体等）赋值即拷贝；引用类型（切片/Map/通道等）赋值共享底层数据

### 4. 常量与 iota
```go
const Pi = 3.14159
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)
const ( // iota：从 0 开始，每新增一行 const 自动 +1
	Mon = iota // 0
	Tue        // 1
	Wed        // 2
)
const ( // iota + 位移的经典用法
	B = 1 << (10 * iota) // 1        (1<<0)
	KB                   // 1024     (1<<10)
	MB                   // 1048576  (1<<20)
)
```

### 5. 运算符
- 算术：`+ - * / %` `++` `--`（注意：`++`/`--` 是**语句不是表达式**，只有后置，不能写 `a = i++`）
- 关系：`== != > < >= <=`
- 逻辑：`&& || !`
- 位运算：`& | ^ << >> &^`（`^` 同时表示异或和按位取反）
- 赋值：`= += -= *= /= %= <<= >>= &= ^= |=`
- 优先级：`*` 系 > `+` 系 > 关系 > `&&` > `||`

### 6. 类型转换——必须显式
```go
var a int = 42
var b float64 = float64(a)  // 不能直接把 a 赋给 float64 变量
var c int = int(b)
```
数值转换可能丢失精度；字符串与数字互转需借助 `strconv` 包（见阶段四）。

### 7. 格式化输出速查（fmt 的常用动词）
```go
fmt.Printf("%v 是最通用的打印\n", 42)     // %v 任意值（默认形态）
fmt.Printf("%d 整数 %s 字符串 %f 浮点\n", 1, "a", 3.14)
fmt.Printf("%q 带引号 %T 类型 %t 布尔\n", "hi", 3.14, true)
msg := fmt.Sprintf("Hi, %v. Welcome!", "Gladys")  // Sprintf：格式化但不打印，返回字符串
```

### 练习
- [ ] 打印各类型变量的零值
- [ ] 用 iota 定义星期常量；用 `1 << (10*iota)` 定义 KB/MB/GB
- [ ] 验证 `17/5` 是整数除法，再练习 int↔float64 显式转换
- [ ] 用 Sprintf 拼一句 "Hi, 你的名字. Welcome!" 并打印

### 过关自检
- `:=` 与 `var` 的使用限制区别？
- rune 和 byte 的区别？为什么遍历中文串要转 `[]rune`？
- %v、%q、%T 分别输出什么？

---

## 阶段三：程序的逻辑 —— 流程控制与函数（第 2 周上半）

**目标**：写出有分支、循环、函数复用的程序，并掌握 defer。

### 1. 条件语句
```go
if a < 20 { ... }                  // 条件不加括号，必须有花括号
if v := 10 * 2; v > 15 { ... }    // if 可带初始化语句，v 仅在 if 块内有效

switch day {                       // 默认每个 case 自动 break，不穿透
case 1: ...
case 6, 7: ...
default: ...
}
switch {                           // 无表达式 switch：替代长串 if-else
case score >= 90: ...
default: ...
}
switch x := 5; x { ... }           // 也可带初始化语句
```

### 2. 循环——只有 for
```go
for i := 0; i < 10; i++ { ... }   // 经典三段式
for i < 10 { ... }                // 形似 while
for { ... }                       // 死循环
for k, v := range collection {}   // 遍历
```
- `break`：跳出循环（配合标签可跳出多重循环 `break LABEL`）
- `continue`：跳过本次循环
- `goto`：无条件跳转到指定标签（不推荐滥用）

### 3. 函数
```go
func max(a, b int) int {          // 参数类型相同时可合并声明
	if a > b { return a }
	return b
}
func swap(x, y string) (string, string) {  // 多返回值
	return y, x
}
func sum(nums ...int) int {       // 可变参数，nums 在函数内是切片
	total := 0
	for _, n := range nums { total += n }
	return total
}
```
- 参数是**值传递**；想修改外部数据需传指针或引用类型
- 函数是一等公民：可赋值给变量、可作为参数和返回值
- 匿名函数与闭包：
```go
add := func(a, b int) int { return a + b }   // 匿名函数赋给变量
func() { fmt.Println("立即执行") }()          // 立即调用
func counter() func() int {                   // 闭包：内层函数引用外层变量
	i := 0
	return func() int { i++; return i }
}
```
- 递归：经典练习是阶乘、斐波那契、求平方根（牛顿迭代）；优点是代码简洁，缺点是深递归消耗栈

### 4. defer——延迟调用（官方教程与标准库中无处不在）
`defer` 语句会把函数调用推迟到**外层函数返回之前**执行，三个规则必须记住：
```go
func main() {
	defer fmt.Println("world")   // 规则1：压栈延迟
	fmt.Println("hello")
	// 输出顺序：hello → world
}

// 规则2：defer 的参数在 defer 语句执行时就求值（不是真正调用时）
func a() {
	i := 0
	defer fmt.Println(i)   // 打印 0，虽然函数返回前 i 已变成 1
	i++
	return
}

// 规则3：多个 defer 后进先出（LIFO），像叠盘子一样，最后 defer 的最先执行
func b() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// 输出顺序：3 → 2 → 1
}
```
经典用途：打开文件后立刻 `defer f.Close()`、加锁后立刻 `defer mu.Unlock()`——"资源在手，归还先行"的编程习惯。阶段八、九会大量用到。

### 5. 作用域
- 局部变量：函数/代码块内声明
- 全局变量：函数外声明，整个包可用；**首字母大写 = 可被其他包访问（导出）**
- 形式参数：函数内当作局部变量

### 练习
- [ ] for-range 打印九九乘法表
- [ ] 写 `divMod(a, b)` 同时返回商和余数
- [ ] 闭包计数器：每次调用返回 1、2、3……
- [ ] 递归实现斐波那契第 n 项
- [ ] 写一个函数用三个 defer 验证 LIFO 顺序

### 过关自检
- Go 有哪三种 for 形态？
- 闭包捕获的变量何时释放？
- defer 的三条规则是什么？为什么打开文件后要立即 defer Close？

---

## 阶段四：数据组织 —— 数组、切片与 Map（第 2 周下半）

**目标**：掌握三大数据容器；切片是重点中的重点。

### 1. 数组——定长、值类型
```go
var arr [5]int                    // 长度是类型的一部分：[5]int ≠ [10]int
arr2 := [5]int{1, 2, 3, 4, 5}
arr3 := [...]int{1, 2, 3}         // 自动推断长度
```
赋值和传参会**整体拷贝**，这是它与切片的本质区别。

### 2. 切片——变长、引用类型（核心）
- 底层结构：指向底层数组的指针 + 长度 len + 容量 cap
```go
s := make([]int, 5)        // len=5, cap=5，元素为零值
s2 := make([]int, 0, 10)   // len=0, cap=10
s3 := []int{1, 2, 3}       // 声明时省略长度即为切片
```
- `len(s)` 长度、`cap(s)` 容量
- 截取：`s[1:3]`（左闭右开）——**与原切片共享底层数组，改 s2 会影响 s1**
- `append(s, elem)`：追加；若超过 cap 会自动扩容（分配更大的底层数组并拷贝）
- `copy(dst, src)`：真正的独立复制
- nil 切片：`var s []int` 长度为 0，引用底层数组为 nil
- 随机取元素（官方教程用法）：`s[rand.Intn(len(s))]`——用 math/rand 生成下标

### 3. Range 遍历与空标识符
```go
for i, v := range slice {}   // 数组/切片：下标、值
for k, v := range m {}       // Map：键、值
for i, c := range "Go语言" {} // 字符串：字节下标、rune
for v := range ch {}         // 通道：不断取值直到通道关闭
for _, v := range slice {}   // _ 空标识符：不需要的返回值用它丢弃（Go 惯例，不可用其他名字代替）
```
空标识符 `_` 是 Go 的惯例符号：range 只要值时丢弃下标、函数只要一个返回值时丢弃另一个（如 `_, ok := m[k]`），后续还会反复出现。

### 4. Map——键值对集合
```go
m := make(map[string]int)       // 初始化（nil map 不能写入，写入会 panic）
m["a"] = 1                      // 增/改
v := m["a"]                     // 查（不存在时返回零值）
v, ok := m["a"]                 // comma-ok：判断 key 是否存在
delete(m, "a")                  // 删
country := map[string]string{   // 声明时初始化
	"CN": "China", "US": "USA",
}
```
- Map 是**无序**的，每次遍历顺序可能不同
- map 的典型用途之一：把一组值映射成另一组值（如"姓名 → 问候语"），阶段七的综合实战会用到

### 5. 字符串与数字互转
- `string ↔ []byte`、`string ↔ []rune` 显式转换
- `strconv.Itoa(42)` 数字转字符串、`strconv.Atoi("42")` 字符串转数字（注意返回两个值，第二个是 error）

### 6. 日常高频工具包：strings / slices / time
写真实程序前先认识三个最常用的标准库，不必背，混个脸熟、用时回来查：

```go
// strings：字符串处理
strings.Contains("hello world", "world")      // 是否包含
strings.HasPrefix(s, "he") / strings.HasSuffix(s, "lo")   // 前缀 / 后缀
strings.ToUpper(s) / strings.ToLower(s)       // 大小写转换
strings.Split("a,b,c", ",")                   // 分割 → []string{"a","b","c"}
strings.Join([]string{"a", "b"}, "-")         // 拼接 → "a-b"
strings.TrimSpace("  hi  ")                   // 去首尾空白
strings.ReplaceAll(s, "old", "new")           // 全量替换
strings.Fields("a  b\tc")                     // 按任意空白切词 → 统计词频的利器

// slices（Go 1.21+）：切片工具（基于泛型实现）
slices.Sort(nums)                 // 升序排序
slices.Contains(nums, 42)         // 是否包含
slices.Index(nums, 42)            // 找下标，没有返回 -1
// 按自定义规则排序用 sort 包：
sort.Slice(users, func(i, j int) bool { return users[i].Age < users[j].Age })

// time：时间处理——Go 的格式化布局是“参考时间”写法，独树一帜
now := time.Now()
now.Format("2006-01-02 15:04:05")  // ⚠ 不是随便的数字！它是 1月2日 15点4分5秒 2006年
// 想输出什么格式，照着参考时间“抄”一份即可；解析也用它：
t, err := time.Parse("2006-01-02", "2026-09-12")   // 字符串 → time.Time
time.Since(start)                  // 耗时统计（返回 Duration）
time.Sleep(2 * time.Second)        // 休眠
```
注意 strings.Fields + map 统计 + sort.Slice 排序，这三板斧组合就是“词频 Top N”的完整解法——毕业项目 2 的核心。

### 练习
- [ ] 统计英文句子中每个单词的出现次数
- [ ] 用 strings.Fields 切词 + map 统计 + sort.Slice 排序，输出词频 Top 3
- [ ] 对切片截取后 append，观察原切片是否被影响并解释
- [ ] 实现"删除切片第 i 个元素"的函数

### 过关自检
- len 与 cap 的区别？append 扩容时发生了什么？
- 为什么两个切片截取自同一底层数组会互相影响？copy 为什么不会？
- strings.Fields 与 strings.Split(" ") 有何区别？如何按结构体字段排序？time 格式化为什么要写 2006-01-02？

---

## 阶段五：类型的深度 —— 指针、结构体与方法（第 3 周上半）

**目标**：能建模现实事物，理解值语义。

### 1. 指针
```go
var a int = 20
var ip *int      // 声明指针变量
ip = &a          // & 取地址
fmt.Println(*ip) // * 解引用取值 → 20
var ptr *int     // 只声明未赋值 → nil 指针
```
- Go **没有指针运算**（不能 `ip++` 移动指针），比 C 安全
- 指针传参的意义：函数内可修改调用方的变量（突破值传递）

### 2. 内存分配：new 与 make 的区别
```go
p := new(int)          // 分配零值内存，返回指针 *int（*p == 0）
q := new(Circle)       // 任何类型都可用 new
s := make([]int, 0, 10) // make 只用于 切片/Map/通道，返回初始化好且可用的 T 本身（不是指针）
m := make(map[string]int)
```
一句话记忆：**new 返回指针，make 返回初始化好的对象**；日常用 make 更多，new 较少直接使用。

### 3. 结构体
```go
type Book struct {
	title  string
	author string
	id     int
}
book := Book{"Go 入门", "菜鸟教程", 1}     // 按顺序初始化
book2 := Book{title: "Go", id: 2}         // 按字段名初始化（官方教程推荐的复合字面量写法）
fmt.Println(book.title)                    // 点号访问成员

var bp *Book = &book
fmt.Println(bp.title)                      // 结构体指针访问成员自动解引用
```
- 结构体是**值类型**：赋值、传参都是整体拷贝
- 结构体指针传参可避免大结构体拷贝、且能在函数内修改原对象

### 4. 方法（带接收者的函数）
```go
type Circle struct { radius float64 }

func (c Circle) Area() float64 {            // 值接收者：操作副本
	return 3.14159 * c.radius * c.radius
}
func (c *Circle) Scale(f float64) {         // 指针接收者：能修改原对象
	c.radius *= f
}
```
选择原则：需要修改自身、或结构体较大时用指针接收者；只需读取用值接收者。

### 5. "继承"——结构体嵌入（组合）
```go
type Person struct { name string; age int }
func (p Person) SayHello() { fmt.Println("你好，我是", p.name) }

type Employee struct {
	Person                       // 匿名嵌入：方法被"提升"
	company string
}
e := Employee{Person{"张三", 30}, "ABC"}
e.SayHello()                    // 直接调用"父类"方法
```
Go 没有类、没有 extends，通过**组合**实现代码复用；嵌入后可覆盖同名方法实现类似"重写"。

### 练习
- [ ] 定义 `Rectangle`，分别用值/指针接收者实现 `Area()` 与 `Scale()`，体会差异
- [ ] 用结构体嵌入组合 Person → Employee，验证方法提升

### 过关自检
- 值接收者与指针接收者的本质区别？
- new 和 make 各返回什么、分别用于什么场景？

---

## 阶段六：抽象的艺术 —— 接口、断言与泛型（第 3 周下半）

**目标**：掌握 Go 的多态与通用编程。

### 1. 接口——隐式实现
```go
type Shape interface {
	Area() float64
	Perimeter() float64
}
// Circle 只要实现了这两个方法，就自动满足 Shape，无需任何声明
func totalArea(shapes []Shape) float64 { ... }  // 多态：统一处理不同类型
```
- 接口特点：方法集隐式实现、值/指针接收者影响实现关系（指针接收者的方法只有指针类型算实现）
- 空接口 `interface{}`（别名 `any`）：可接收任何值
- 接口组合：接口可嵌入其他接口形成更大的接口
- 接口变量底层 = 动态类型 + 动态值；接口零值是 `nil`

### 2. 标准库中最重要的接口：Stringer 与 error
标准库大量用"隐式实现"提供扩展点，最典型的两个：
```go
// fmt.Stringer：实现 String() 方法后，打印时自动调用（类似自定义 toString）
type Person struct { Name string; Age int }
func (p Person) String() string {
	return fmt.Sprintf("%v (%v 岁)", p.Name, p.Age)
}
fmt.Println(Person{"张三", 30})   // 输出：张三 (30 岁)

// error 本身就是接口：type error interface { Error() string }
// 任何实现 Error() string 的类型都可以当 error 返回——阶段七细讲
```

### 3. 类型断言
```go
var i interface{} = "hello"
s := i.(string)        // 断言失败会 panic
s, ok := i.(string)    // comma-ok：失败不 panic，ok 为 false

switch v := i.(type) { // 类型分支
case int:    fmt.Println("整数", v)
case string: fmt.Println("字符串", v)
default:     fmt.Println("未知类型")
}
```

### 4. 泛型——从两个重复函数到一个通用函数
官方泛型教程用"求和函数"演示泛型的价值，完整走一遍这个过程比背语法有效得多。

**第一步：没有泛型时，只能为每种类型写一个函数**
```go
// SumInts 对 int64 类型的 map 求和
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}
// SumFloats 对 float64 类型的 map 求和——逻辑完全重复！
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}
```
**第二步：引入类型参数，合并成一个函数**
```go
// [K comparable, V int64 | float64] 是类型参数列表：
// K 约束为可比较类型（能作 map 的键），V 只允许 int64 或 float64
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
```
**第三步：调用——可显式传类型参数，也可让编译器自动推断**
```go
ints := map[string]int64{"first": 34, "second": 12}
floats := map[string]float64{"first": 35.98, "second": 26.99}

fmt.Printf("%v %v\n",
	SumIntsOrFloats[string, int64](ints),     // 显式指定
	SumIntsOrFloats(floats))                   // 自动推断（推荐）
```
**第四步：用约束接口让声明更整洁**
```go
type Number interface {
	int64 | float64      // 约束接口：把"允许的类型"收拢到一个具名接口里
}
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m { s += v }
	return s
}
```
泛型补充要点：
- 常用约束：`any`（任意类型）、`comparable`（可比较）、自定义约束接口
- 泛型结构体同样支持：
```go
type Stack[T any] struct { items []T }
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }
func (s *Stack[T]) Pop() T {
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v
}
```
- 泛型是编译期实例化，无运行时开销

### 练习
- [ ] 定义 Shape 接口，让圆和矩形隐式实现，统一计算面积总和
- [ ] 给你的结构体实现 String() 方法，观察打印变化
- [ ] 亲手重走"求和函数泛型化"四步，再写一个泛型 Map 函数（对切片每个元素做变换）
- [ ] 用 type-switch 处理 interface{} 中的多种类型

### 过关自检
- 什么叫"隐式实现"？接口变量底层存了什么？
- `v, ok := x.(T)` 失败时行为与 `v := x.(T)` 有何不同？
- 类型参数、类型约束分别解决什么问题？

---

## 阶段七：健壮与规范 —— 错误处理与工程化（第 4 周）

**目标**：写出能正确处理失败、可复用、可测试的工程代码。

### 1. 错误处理——error 是普通值
```go
// error 本质是一个接口：type error interface { Error() string }
err1 := errors.New("文件不存在")                 // errors 包创建错误
err2 := fmt.Errorf("打开 %s 失败: %w", f, err1)  // %w 包装底层错误，保留错误链

val, err := strconv.Atoi("abc")
if err != nil {                                 // 标准范式：立即检查
	fmt.Println("出错了:", err)
	return
}

// 自定义错误：实现 Error() 方法即可
type MyError struct { Code int; Msg string }
func (e *MyError) Error() string { return fmt.Sprintf("[%d] %s", e.Code, e.Msg) }

// 错误链判断
errors.Is(err, targetErr)      // 判断错误链上是否包含目标
var me *MyError
errors.As(err, &me)            // 提取错误链上特定类型
```
**panic 与 recover——只用于不可恢复的场景**
```go
defer func() {
	if r := recover(); r != nil {  // recover 必须放在 defer 中
		fmt.Println("已捕获 panic:", r)
	}
}()
panic("严重错误")   // 程序崩溃前会先执行所有 defer
```
原则：可预期的失败用 error 返回；程序逻辑彻底崩坏（如数组越界、nil 写入）才 panic。

### 2. log 包——错误信息的正式输出
比 fmt.Println 更正式的日志方式（官方教程的标配用法）：
```go
log.SetPrefix("greetings: ")   // 给日志加前缀，方便定位来源
log.SetFlags(0)                // 0 = 不打印时间戳/文件行号（默认会打印）
log.Fatal(err)                 // 打印错误并立即退出程序（exit status 1）
```

### 3. 内置单元测试——go test
Go 自带测试框架，零配置。约定：
- 测试文件必须以 `_test.go` 结尾（如 `greetings_test.go`），和被测代码同包
- 测试函数以 `Test` 开头，参数是 `t *testing.T`，用 `t.Errorf` 报告失败
```go
package greetings

import (
	"regexp"
	"testing"
)

// 测试正常输入：返回值应包含名字且无错误
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Errorf(`Hello("Gladys") = %q, %v, want match for %q, nil`, msg, err, want)
	}
}

// 测试异常输入：空名字应返回错误
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
```
运行：`go test`（加 `-v` 显示每个用例详情）。**边开发边写测试**是官方强调的习惯。

Go 社区最惯用的其实是**表格驱动测试**：把所有用例组织成一个切片，循环执行，以后加用例只需加一行：

```go
func TestHello(t *testing.T) {
	tests := []struct {
		name    string // 用例说明
		in      string // 输入
		wantErr bool   // 期望是否出错
	}{
		{"正常名字", "Gladys", false},
		{"空名字应报错", "", true},
		{"中文名字", "张三", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {    // 子测试：每个用例独立报告 PASS/FAIL
			msg, err := Hello(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello(%q) error = %v, wantErr %v", tt.in, err, tt.wantErr)
				return
			}
			if !tt.wantErr && msg == "" {
				t.Errorf("Hello(%q) 返回了空消息", tt.in)
			}
		})
	}
}
```

### 4. 包与模块
```bash
go mod init example.com/myapp   # 初始化模块，生成 go.mod
go get github.com/some/package  # 添加依赖
go mod tidy                     # 清理/补全依赖
go install .                    # 编译并安装当前程序到 $GOPATH/bin
```
- `go.mod`：记录模块名、Go 版本、依赖列表；`go.sum`：依赖校验和，保证构建可重现
- 包的可见性：**标识符首字母大写 = 导出**，小写 = 包内私有
- 依赖下载慢时可设置国内代理：`go env -w GOPROXY=https://goproxy.cn,direct`

### 5. 综合实战：双模块项目（官方 create-module 教程完整版）
这个项目把本阶段全部知识串起来：**greetings**（被调用的库模块）+ **hello**（调用方程序），是 Go 工程化的标准形态。

**步骤 1：创建 greetings 模块**
```bash
mkdir greetings && cd greetings
go mod init example.com/greetings
```

**步骤 2：编写 greetings/greetings.go（最终完整版）**
```go
package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello 为指定的人返回问候语。
// 函数名的 H 大写 = 导出，供其他包调用。
func Hello(name string) (string, error) {
	if name == "" {   // 参数校验：无法处理的输入返回 error
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil   // 成功时 error 位置返回 nil
}

// Hellos 为一组人返回"姓名 → 问候语"的映射。
// 注意：不修改已发布的 Hello 函数签名，而是新增函数——保持向后兼容的工程思维。
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {   // _ 丢弃不需要的下标
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// randomFormat 随机返回一种问候格式。
// r 小写 = 未导出，仅本包可用（包的封装性）。
func randomFormat() string {
	formats := []string{          // 切片：长度可变的数组
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]   // 随机下标取元素
}
```

**步骤 3：创建 hello 调用方模块并本地关联**
```bash
cd .. && mkdir hello && cd hello
go mod init example.com/hello
```
在 hello/go.mod 中把依赖指到本地目录（发布后改为真实仓库地址）：
```
module example.com/hello

go 1.22

require example.com/greetings v0.0.0

replace example.com/greetings => ../greetings
```

**步骤 4：编写 hello/hello.go**
```go
package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	// 输出示例（每次运行的问候格式随机）：
	// map[Darrin:Hail, Darrin! Well met! Gladys:Hi, Gladys. Welcome! Samantha:Great to see you, Samantha!]
}
```

**步骤 5：补上测试与安装**
- 把第 3 节的测试代码存为 `greetings/greetings_test.go`，在 hello 目录运行 `go test ./...` 全绿
- `go install .` 后，`~/go/bin/hello` 可在任意目录直接执行

### 练习
- [ ] 完整走通上述双模块项目，并故意传空名字观察 log.Fatal 的行为
- [ ] 给 randomFormat 再加两种格式；给 greetings 增加 `Goodbye(name string) (string, error)` 并补测试

### 过关自检
- error 与 panic 的适用边界？`_test.go` 和 `TestXxx` 命名约定是什么？
- 什么情况下"新增函数"比"修改函数签名"更好？
- replace 指令解决什么问题？
- 表格驱动测试相比写多个独立测试函数好在哪？

---

## 阶段八：Go 的灵魂 —— 并发编程（第 5 周）⭐

**目标**：掌握 goroutine 与 channel——选择 Go 的核心理由，并用 context 管理超时与取消。

### 1. Goroutine——轻量级线程
```go
go sayHello("A")     // go 关键字：启动一个 goroutine，函数立即返回
// goroutine 初始栈仅几 KB，单机可轻松运行百万个
```
- 主 goroutine（main 函数）退出时，所有子 goroutine 立即终止——这就是常见的"程序没输出就结束了"的原因

### 2. Channel——goroutine 间通信
```go
ch := make(chan int)        // 无缓冲：发送阻塞直到有人接收（同步）
ch2 := make(chan int, 10)   // 有缓冲：装满前发送不阻塞（异步）
ch <- 10                    // 发送
v := <-ch                   // 接收
close(ch)                   // 关闭（发送方关闭；关闭后再发送会 panic）
v, ok := <-ch               // ok=false 表示通道已关闭且取空
for v := range ch {}        // 持续接收直到通道关闭
```
设计哲学（阶段一 quote.Go() 输出的那句话）：**不要通过共享内存来通信，而要通过通信来共享内存**。

### 3. Select——多路复用
```go
select {
case v := <-ch1:
	fmt.Println("ch1:", v)
case ch2 <- 100:
	fmt.Println("发送到 ch2")
case <-time.After(3 * time.Second):  // 超时控制
	fmt.Println("超时")
default:                             // 所有通道都没准备好时执行（非阻塞）
	fmt.Println("无数据")
}
```

### 4. 同步工具
```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
	wg.Add(1)               // 计数 +1
	go func(id int) {
		defer wg.Done()     // 计数 -1（defer 的标准用法）
		work(id)
	}(i)
}
wg.Wait()                    // 阻塞直到计数归零

var mu sync.Mutex            // 互斥锁：保护共享数据
mu.Lock(); counter++; mu.Unlock()
```
常见坑：① 主 goroutine 提前退出；② 向已关闭通道发送数据 panic；③ 循环启动 goroutine 时捕获循环变量（仅限 Go 1.21 及更早——Go 1.22 起循环变量按迭代独立，此坑已消失，但读老代码时仍会遇到）；④ Map 并发读写会崩溃，必须加锁或用 sync.Map。

### 5. 发起网络请求——net/http 客户端
并发练习（如爬虫）离不开发 HTTP 请求，标准库 net/http 几行搞定：

```go
// GET 最简形态
resp, err := http.Get("https://go.dev")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()          // 响应体是资源，必须关闭（defer 的又一次兑现）
body, err := io.ReadAll(resp.Body)
fmt.Println(resp.StatusCode, len(body))

// POST JSON 的标准三步
req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonData))
if err != nil { return err }
req.Header.Set("Content-Type", "application/json")
resp2, err := http.DefaultClient.Do(req)
defer resp2.Body.Close()

// 带超时的客户端——生产必备！http.DefaultClient 默认无超时，慢接口会拖死程序
client := &http.Client{Timeout: 10 * time.Second}
resp3, err := client.Get(url)
```

> 示例中的 io.ReadAll（把响应流读成 []byte）与 bytes.NewReader（把 []byte 包装成请求体读取器）互为逆操作，是 io / bytes 两个包里最常用的函数，先混个脸熟即可。

### 6. context——超时与取消的标准传递方式
真实服务必须回答两个问题：“这个请求最多允许跑多久？”“客户端断开了，后续工作还要不要做？”Go 的答案是 context：在调用链中显式传递截止时间与取消信号。**约定 ctx 永远是函数的第一个参数**。它的用法完全建立在本阶段刚学的 select 与 channel 之上：

```go
func worker(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):   // 模拟一段耗时工作
		return nil
	case <-ctx.Done():                     // 上游超时/取消时 Done 通道被关闭
		return ctx.Err()                    // DeadlineExceeded 或 Canceled
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()   // 惯例：拿到 cancel 立即 defer，防止 context 泄漏
	if err := worker(ctx); err != nil {
		log.Println("未完成:", err)   // 2 秒即中断，不会等满 5 秒
	}
}
```
三个要点：
- `context.Background()` 是根 context，main 的起点；`WithTimeout / WithCancel / WithValue` 从父 context 派生子 context，随调用链层层传递
- 标准库全面接入：`db.QueryContext(ctx, ...)`、`http.NewRequestWithContext(ctx, ...)`、Gin 里用 `c.Request.Context()` 取当前请求的 ctx（阶段九 Gin 一节会实际接入）
- ctx 只作参数传递，不要存进结构体、不要长期持有

### 练习
- [ ] 生产者-消费者：3 个生产者往 channel 投递、1 个消费者汇总
- [ ] 给生产者-消费者加 3 秒整体超时：用 context.WithTimeout + select 控制未完成时优雅退出
- [ ] 用 WaitGroup 并发统计 1~100 万内的素数个数
- [ ] 并发转账程序：先裸并发观察数据竞争，再用 Mutex 修复
- [ ] 官方 Tour 的经典综合题——并发 Web 爬虫：给定起始 URL，用 goroutine 并发抓取页面（用第 5 节的 http.Get）、提取链接、去重后继续爬取，限制并发数不超过阈值。综合考查 goroutine、channel、互斥锁与程序设计能力，是本阶段最好的检验

### 过关自检
- 无缓冲与有缓冲通道的阻塞行为差异？
- 向已关闭的 channel 发送 / 接收分别发生什么？
- WaitGroup 和 channel 各适合什么等待场景？
- 为什么拿到 cancel 后要立即 `defer cancel()`？ctx 约定放在参数列表什么位置？
- 为什么生产代码要自定义 http.Client 的 Timeout？resp.Body 不 Close 会怎样？

---

## 阶段九：落地能力 —— 文件、数据库、Web 服务与毕业实战（第 6 周）

**目标**：补齐实战三件套（文件 IO、数据库、Web API），完成毕业项目。

### 1. 文件处理速查
```go
// 读
data, _ := os.ReadFile("a.txt")                 // 小文件一次性读
f, _ := os.Open("a.txt")                        // 打开（只读），用完 f.Close()
scanner := bufio.NewScanner(f)                  // 大文件逐行读
for scanner.Scan() { line := scanner.Text() }

// 写
os.WriteFile("b.txt", []byte("内容"), 0644)      // 一次性写
f, _ := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 追加
w := bufio.NewWriter(f)                         // 缓冲写（大量数据推荐）
w.WriteString("hello"); w.Flush()

// 其他高频操作
os.Remove("a.txt")                              // 删除
os.Rename("a.txt", "b.txt")                     // 重命名/移动
os.Stat("a.txt")                                // 文件信息；配合 os.IsNotExist 判断是否存在
os.MkdirAll("a/b/c", 0755)                      // 递归建目录
filepath.Join("a", "b", "c.txt")                // 跨平台路径拼接
io.Copy(dst, src)                               // 文件复制
filepath.WalkDir(root, fn)                      // 递归遍历目录
```
> 更规范的开法是打开后立即 `defer f.Close()`——阶段三的 defer 在这里兑现。

### 2. 正则表达式
```go
re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)  // 编译（MustCompile 失败即 panic）
re.MatchString("日期 2026-09-12")                    // 是否匹配
re.FindAllString(s, -1)                              // 找出全部匹配
sub := re.FindStringSubmatch("2026-09-12")           // 分组捕获: [整串, 2026, 09, 12]
re.ReplaceAllString(s, "$3/$2/$1")                   // 替换（$n 引用分组）
re.Split("a,b,,c", -1)                               // 分割
```

### 3. JSON 序列化
```go
type Person struct {
	Name string `json:"name"`       // tag 指定 JSON 字段名；不写 tag 会直接用大写的字段名
	Age  int    `json:"age"`
}
b, _ := json.Marshal(p)             // 结构体 → JSON
json.Unmarshal(b, &p2)              // JSON → 结构体（必须传指针）
```

### 4. 数据库访问——database/sql 标准库
以 MySQL 为例（官方教程场景：老爵士唱片库），三步走：

**第一步：建库建表**
```sql
create database recordings;
use recordings;
CREATE TABLE album (
	id     INT AUTO_INCREMENT NOT NULL,
	title  VARCHAR(128) NOT NULL,
	artist VARCHAR(255) NOT NULL,
	price  DECIMAL(5,2) NOT NULL,
	PRIMARY KEY (`id`)
);
INSERT INTO album (title, artist, price) VALUES
	('Blue Train', 'John Coltrane', 56.99),
	('Giant Steps', 'John Coltrane', 63.99),
	('Jeru', 'Gerry Mulligan', 17.99),
	('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

**第二步：连接数据库**
```bash
go mod init example/data-access
go get github.com/go-sql-driver/mysql
```
```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"   // 驱动包：只用它的配置功能
	// 若不需要驱动的高级功能，常见写法是空导入：_ "github.com/go-sql-driver/mysql"
	// 作用是执行驱动的 init()，把自己注册进 database/sql
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	cfg := mysql.NewConfig()             // 用配置对象拼连接串，避免手写 DSN 出错
	cfg.User = os.Getenv("DBUSER")       // 从环境变量读账号密码，不硬编码
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "recordings"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())   // sql.Open 不真正连接，只准备句柄
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {   // Ping 才真正建立连接，启动时尽早失败
		log.Fatal(err)
	}
	fmt.Println("Connected!")
}
```

**第三步：CRUD 三板斧（所有错误都按 (值, error) 惯例处理）**
```go
// 查多行：Query + rows.Next 循环 + Scan 装填 + defer rows.Close
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name) // ? 占位符防注入
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()                       // 必须关闭，否则连接泄漏
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {       // 遍历结束后统一检查迭代错误
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// 查单行：QueryRow + Scan；用 sql.ErrNoRows 判断"没查到"
func albumByID(id int64) (Album, error) {
	var alb Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}
	return alb, nil
}

// 增删改：Exec；用 LastInsertId 拿自增主键
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
```
> 记忆口诀：查多条 Query、查一条 QueryRow、写数据 Exec；SQL 一律用 `?` 占位符传参，永不相接字符串。

### 5. Web 服务——Gin 框架写 REST API
目标：一个专辑管理 API，两个资源端点。

**端点设计（先设计 API 再写代码）**

| 方法 | 路径 | 功能 |
|---|---|---|
| GET | /albums | 返回全部专辑 |
| GET | /albums/:id | 按 ID 返回单张专辑 |
| POST | /albums | 新增专辑 |

**完整实现**
```bash
mkdir web-service-gin && cd web-service-gin
go mod init example/web-service-gin
go get github.com/gin-gonic/gin     # 拉取 Gin 依赖
```
```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album 结构体：json tag 决定序列化后的字段名（小写下划线风格是 API 惯例）
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 内存存储种子数据（重启即失；生产中换成上一节的数据库）
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// 处理函数的签名固定：参数是 *gin.Context
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)   // 200 + 缩进格式化的 JSON
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {   // 请求体 JSON → 结构体
		return                                      // BindJSON 失败时已自动回了 400
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)    // 201 + 新资源
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")                      // 取路径参数 /albums/:id 中的 :id
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})  // 404 + JSON 提示
}

func main() {
	router := gin.Default()                  // 带日志和恢复中间件的路由器
	router.GET("/albums", getAlbums)         // 注册路由：方法 + 路径 + 处理函数
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")             // 启动服务（默认 8080）
}
```
**启动并验证**
```bash
$ go run .
# 另开一个终端：
$ curl http://localhost:8080/albums
$ curl http://localhost:8080/albums/2
$ curl http://localhost:8080/albums \
    --include --header "Content-Type: application/json" \
    --request POST \
    --data '{"id":"4","title":"The Modern Sound of Betty Carter","artist":"Betty Carter","price":49.99}'
```
第三条命令会收到 `201 Created`；再次 GET /albums 就能看到新数据。这个 60 行的程序就是 Go Web 后端的最小完整闭环。

**接入 context（兑现阶段八的伏笔）**：每个 Gin 处理函数里都能用 `c.Request.Context()` 拿到“当前请求”的 ctx，传给任何支持它的耗时调用：
```go
func getAlbums(c *gin.Context) {
	ctx := c.Request.Context()   // 请求级 ctx：客户端断开时自动取消
	// 换成真实耗时操作时传入：db.QueryContext(ctx, ...)、http.NewRequestWithContext(ctx, ...)
	c.IndentedJSON(http.StatusOK, albums)
}
```
这就是“请求生命周期管理”的标准方式——客户端中途断开，下游的数据库查询 / HTTP 调用会随之中止，不再浪费资源。

### 6. 命令行参数——os.Args 与 flag
CLI 工具（包括毕业项目 1）需要读命令行参数，两种方式按需选择：

```go
// 方式一：os.Args——最原始的参数切片
// 运行 ./todo add 买牛奶  →  os.Args = ["./todo", "add", "买牛奶"]
fmt.Println(os.Args)      // os.Args[0] 是程序名，参数从 os.Args[1] 开始

// 方式二：flag——带类型解析、默认值和自动 -h 帮助，写工具的标准姿势
func main() {
	n := flag.Int("n", 10, "输出 Top N 条")   // 返回 *int：参数名、默认值、说明
	w := flag.String("w", "", "关键词")        // 返回 *string
	flag.Parse()                             // 定义完全部 flag 后调用一次
	fmt.Println(*n, *w)                      // 运行：./tool -n 20 -w=error
}

// 子命令风格（todo add / todo list）用 os.Args[1] 手动分发：
if len(os.Args) < 2 {
	fmt.Println("用法: todo <add|list|done> [参数]")
	return
}
switch os.Args[1] {
case "add":  // os.Args[2:] 是子命令自己的参数
case "list":
}
```

### 过关自检
- os.Args[0] 存的是什么？flag.Int 返回什么类型？flag.Parse() 应在何时调用？
- 查多行、查单行、增删改分别对应 database/sql 的哪三个方法？

### 毕业项目（三选一，融合全部所学）
1. **命令行待办清单**：结构体 + 切片 + JSON + 文件持久化 + os.Args 命令解析
2. **并发日志分析器**：bufio 逐行读大文件 → 正则提取 IP/URL → goroutine + channel 分片统计 → 输出 Top N
3. **专辑 REST API 完整版**：把第 5 节的内存版改造为"第 4 节数据库存储 + Gin 暴露 API + DELETE/PUT 端点 + _test.go 测试"——做完它你就具备了初级 Go 后端的全部要素

### 最终自检（整条路线）
- [ ] 能从零创建 go mod 项目、拆分多包并写测试
- [ ] 能讲清：切片扩容机制、接口隐式实现、channel 阻塞规则、defer 三规则
- [ ] 能独立写出 200 行以上、正确处理 error 与并发的完整程序
- [ ] 能用 database/sql 完成 CRUD、用 Gin 暴露 REST API
- [ ] 会用 context 控制超时取消、用 http.Client 发请求、用 flag/os.Args 写 CLI 工具

---

## 学习原则

1. **每段代码亲手敲**：看懂 ≠ 会写；改一改示例观察行为变化是最好的理解方式
2. **从第一天就用官方工具**：`go fmt` 格式化、`go vet` 静态检查、`go test` 随手测
3. **并发阶段放慢速度**：阶段八是全路线核心，值得花一周反复练习
4. **两遍学习法**：第一遍照本路线走完并做完毕业项目；第二遍精读官方《Effective Go》（惯用法指南）与 FAQ（回答"为什么 Go 不做 XX"），把代码从"能跑"提升到"地道"
5. **学完之后的方向**：真实项目（Gin + database/sql 已具备雏形）→ Docker 容器化部署 → 进阶并发模式（流水线、worker pool、errgroup）→ 关注官方博客跟进版本新特性

---

## 附录：学完本路线后的延伸资料（可选）

| 资料 | 定位 |
|---|---|
| A Tour of Go | 交互式教程，浏览器直接跑代码，适合做第二遍练习 |
| Effective Go | 官方"惯用写法指南"，复习阶段的精读材料 |
| Go FAQ | 官方问答，解答语言设计取舍 |
| 标准库文档（pkg.go.dev） | API 手册；fmt/os/io/net/http 最常用 |
| 语言规范（Language Specification） | 语法的最终权威，遇到边界争议时查证 |
| 内存模型（Memory Model） | 并发可见性规则，深入并发后阅读 |
| Go 官方博客 | 每个版本新特性的深度文章 |

**后续深入主题（按需选学）**：sync/atomic 原子操作、接口 nil 陷阱（装入了 nil 指针的接口并不等于 nil）、init 函数与包初始化顺序、交叉编译（GOOS/GOARCH）、go work 多模块工作区、pprof 性能分析、worker pool / 流水线并发模式、reflect 反射。
