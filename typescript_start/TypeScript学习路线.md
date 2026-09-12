# TypeScript 学习路线：零基础 → Node.js 后端工程实践

> **总时长**：约 31 天（4~5 周），每天 4~5 小时，合计约 130~150 小时
> **版本基线**：TypeScript 5.x（含 5.0 标准装饰器与旧版实验装饰器的差异说明）、Node.js 18 LTS 及以上
> **内容构成**：菜鸟教程 TypeScript 教程章节（React/Vue 前端实战两章按方向移入“后续深入”）+ 菜鸟 JavaScript 教程核心章节（按 Node.js 后端方向裁剪，浏览器/DOM 章节不纳入）+ TypeScript 官方 Handbook（基础篇、类型变换篇、参考篇全部章节，JSX 章节移入“后续深入”）三源融合，互相补全
> **读者设定**：零基础（JavaScript 也不会）→ 全面深入（含装饰器、类型体操）→ 以 Node.js 后端为目标场景
> **学习方式**：所有代码示例都可直接在本地运行；每阶段末尾有练习与过关自检，自检全部能回答再进入下一阶段

---

## 总览：为什么这条路线长这样

TypeScript 是 JavaScript 的**超集**：所有合法的 JavaScript 代码都是合法的 TypeScript 代码，TS 在 JS 之上叠加了一套**静态类型系统**，最终编译回纯 JavaScript 运行。这决定了两件事：

1. **必须先学 JavaScript**。类型在编译后会被全部抹除，运行时行为完全由 JS 决定。不懂 JS 就学 TS，等于不懂走路学跑步。
2. **TS 的价值在编译期，不在运行期**。下面这个例子是全路线的"第一课"：

```typescript
// JavaScript：运行时才会报错
function greet(name) {
    return "Hello, " + name.toUpperCase();
}
greet(123); // 运行时报错：name.toUpperCase is not a function

// TypeScript：编译时就能发现错误
function greet(name: string): string {
    return "Hello, " + name.toUpperCase();
}
greet(123); // 编译错误：Argument of type 'number' is not assignable to parameter of type 'string'
```

| 问题场景 | JavaScript 的困境 | TypeScript 的解决方式 |
|---|---|---|
| 函数参数传错类型 | 运行时才报错 | 编译阶段即报错，IDE 实时提示 |
| 访问不存在的属性 | 返回 undefined，行为难预测 | 编译器直接报错 |
| 大型项目重构 | 改一处不知道哪里会崩 | 类型系统自动追踪所有引用 |
| 团队协作 | 函数签名全靠注释 | 类型签名即文档，自动补全 |

TypeScript 不是银弹，先看清它的边界：类型信息**编译后全部抹除**（运行时无类型）、相比纯 JS 多一道编译步骤、`any` 会被滥用、类型体操有门槛。了解边界和了解能力同样重要。

**阶段地图**：

| 阶段 | 主题 | 天数 |
|---|---|---|
| 一 | 跑通工具链 + JavaScript 核心（上）：变量、类型、运算符、控制流 | 3 天 |
| 二 | JavaScript 核心（下）：函数、对象、数组、解构、异步 | 4 天 |
| 三 | TypeScript 入门：类型注解、基础类型、联合与交叉 | 3 天 |
| 四 | TypeScript 核心：函数类型、接口、类、类型收窄 | 4 天 |
| 五 | 泛型与类型变换（类型体操 · 上） | 3 天 |
| 六 | 类型系统深水区（类型体操 · 下）：工具类型、类型兼容性 | 3 天 |
| 七 | 工程化与语言机制：模块、声明文件、tsconfig、装饰器、生成器 | 4 天 |
| 八 | Node.js 后端实战：Express API + 类型分层 + 单元测试 | 3 天 |
| 九 | 毕业项目 | 4 天 |

---

## 阶段一：跑通工具链 + JavaScript 核心（上）

> **目标**：搭好 Node.js + TypeScript 开发环境，跑通「写 .ts → 编译 → 运行」的完整闭环；掌握 JavaScript 的变量、数据类型、运算符与控制流。本阶段末尾你将能读懂并编写简单的 JS 脚本。

### 1.1 环境搭建

需要安装三样东西：

1. **Node.js**（含 npm 包管理器）：从官网下载 LTS 版本（18 及以上），安装后在终端验证：

```bash
node -v   # 输出类似 v20.11.0
npm -v    # 输出版本号
```

2. **TypeScript 编译器**：通过 npm 全局安装（国内网络可先切换镜像源）：

```bash
# 可选：切换国内镜像加速
npm config set registry https://registry.npmmirror.com

# 全局安装 TypeScript，得到 tsc 命令
npm install -g typescript

tsc -v     # 输出 Version 5.x.x
```

3. **VS Code**：微软出品的编辑器，对 TypeScript 有一级支持（智能提示、错误标红、跳转定义、悬浮文档、快速修复）。装好后打开任意 .ts 文件即可体验。

### 1.2 第一个程序：看清 TS 的工作方式 ⭐

新建 `hello.ts`：

```typescript
const hello: string = "Hello World!";
console.log(hello);
```

在终端执行：

```bash
tsc hello.ts     # 编译：同目录生成 hello.js
node hello.js    # 运行：输出 Hello World!
```

打开生成的 `hello.js`，你会发现**类型注解消失了**：

```javascript
var hello = "Hello World!";
console.log(hello);
```

这就是 TypeScript 的本质：**类型只活在编译期，编译产物是纯 JavaScript**。整个流程：`.ts 源码 → tsc 编译 → .js → node 执行`。

`main.ts` 里如果写了类型错误，`tsc` 会直接报错并指出行号——但注意：**即使报错，tsc 默认仍会生成 .js**（可用 `noEmitOnError` 阻止）。

**tsc 常用命令速查**：

| 命令 | 作用 |
|---|---|
| `tsc app.ts` | 编译单个文件 |
| `tsc file1.ts file2.ts` | 同时编译多个文件 |
| `tsc --watch` / `tsc -w` | 监视模式，文件变化自动重新编译 |
| `tsc --declaration` | 额外生成 `.d.ts` 类型声明文件（阶段七详解） |
| `tsc --target ES2020` | 指定编译目标的 JS 版本 |
| `tsc --removeComments` | 删除注释 |
| `tsc --sourceMap` | 生成 .map 源码映射文件，用于调试时对应源码 |
| `tsc --noImplicitAny` | 隐式 any 报错（严格模式的一部分，阶段七详解） |

> 💡 伏笔：`--strict` 严格模式家族为什么推荐"始终开启"，会在阶段七讲 tsconfig 时系统展开；现在只需记住命令行用法。

### 1.3 变量声明：var / let / const ⭐

JavaScript 有三种声明变量的关键字。历史上只有 `var`（ES5），ES6（2015）引入了 `let` 和 `const`：

```javascript
var a = 1;    // 函数作用域，可重复声明，可重新赋值
let b = 2;    // 块级作用域，不可重复声明，可重新赋值
const c = 3;  // 块级作用域，不可重复声明，不可重新赋值
```

**var 的三个坑**（理解了它们就理解了为什么现代代码只用 let/const）：

**坑 1：没有块级作用域**——`var` 声明的变量在 `{}` 外依然可见：

```javascript
{
    var x = 2;
}
console.log(x); // 2，变量"泄漏"到块外

{
    let y = 2;
}
// console.log(y); // ReferenceError: y is not defined
```

**坑 2：块内重新声明会覆盖块外变量**：

```javascript
var x = 10;
{
    var x = 2;   // 覆盖了外面的 x
}
console.log(x);  // 2

var i = 5;
for (var i = 0; i < 10; i++) { }  // 循环变量也是同一个 i
console.log(i);  // 10
```

**坑 3：循环捕获问题**——经典的 `var + setTimeout` 面试题：

```javascript
for (var i = 0; i < 3; i++) {
    setTimeout(function () { console.log(i); }, 100);
}
// 输出 3 3 3：三个回调捕获的是同一个函数作用域的 i，循环结束时 i 已是 3

for (let i = 0; i < 3; i++) {
    setTimeout(function () { console.log(i); }, 100);
}
// 输出 0 1 2：let 为每次迭代创建一个全新的变量环境
```

**const 的注意点**：`const` 锁的是"变量名到值的绑定"，不是值本身的不可变：

```javascript
const kitty = { name: "Aurora" };
kitty.name = "Rory";   // 合法：修改对象内部状态
kitty.numLives = 9;    // 合法：新增属性
// kitty = {};         // 错误：不能重新赋值
```

> **实践准则**（最小权限原则）：默认全用 `const`，确实需要重新赋值才用 `let`，永远不用 `var`。本路线后续所有示例遵循此准则。

### 1.4 数据类型：动态类型意味着什么

JavaScript 是**动态类型语言**：同一个变量可以先后赋不同类型的值：

```javascript
let x;            // undefined
x = 5;            // 现在 x 是数字
x = "John";       // 现在 x 是字符串
```

这与 TypeScript 的静态检查形成对照——TS 在编译期就会拒绝上面的第三行。JS 的数据类型分两大类：

**原始类型（值类型）**：`string`、`number`、`boolean`、`null`、`undefined`、`symbol`（ES6 新增，表示独一无二的值）
**引用类型（对象类型）**：`Object`、`Array`、`Function`，以及特殊的 `RegExp`（正则）和 `Date`（日期）

各类型要点：

```javascript
let str = "hello";            // 单引号双引号均可；模板字符串用反引号（1.6 节）
let n1 = 3.14;                // JS 只有一种数字类型（64 位浮点），不分整数浮点
let n2 = 123e5;               // 科学计数法：12300000
let big = 999999999999999;    // 整数最多安全表示 15 位
console.log(0.1 + 0.2);       // 0.30000000000000004 —— 浮点精度问题
let ok = true;                // boolean 只有 true / false
let empty = null;             // "空值"：明确的空对象引用
let notSet;                   // undefined：声明了但没赋值
```

**用 typeof 检测类型**（运行时），以及它的两个著名怪癖：

```javascript
typeof "abc"       // "string"
typeof 3.14        // "number"
typeof true        // "boolean"
typeof undefined   // "undefined"
typeof function(){} // "function"
typeof {}          // "object"
typeof [1,2,3]     // "object"  ← 怪癖 1：数组也是 "object"
typeof null        // "object"  ← 怪癖 2：历史遗留 bug，null 被检测为 "object"
```

正确检测数组用 `Array.isArray([1,2,3])`（返回 true）。检测 null 用严格相等 `x === null`。

**NaN 与 Infinity**：

```javascript
let x = 1000 / "Apple";   // NaN：非数字
isNaN(x);                 // true
1000 / 0;                 // Infinity（除以 0 得无穷大，不报错）
```

### 1.5 运算符

```javascript
// 算术：+ - * / % **（** 是 ES2016 指数运算符，2 ** 3 === 8）
// 赋值：= += -= *= /= %= **=
let a = 5;
a += 2;  // 等价于 a = a + 2

// 比较运算符
5 == "5";   // true  （宽松相等：先做类型转换再比较 —— 危险！）
5 === "5";  // false （严格相等：类型不同直接 false）
5 != "5";   // false
5 !== "5";  // true
// 铁律：永远使用 === 和 !==，禁止 == 和 !=

// 逻辑运算符：&& || !
true && "hello";   // "hello"（逻辑与：第一个为真返回第二个）
false || "hi";     // "hi"  （逻辑或：第一个为假返回第二个）
!0;                // true

// 三元（条件）运算符
let age = 20;
let type = age >= 18 ? "成年" : "未成年";

// 短路求值：&& 左边为假时右边不执行，|| 左边为真时右边不执行
let user = null;
let name = user && user.name;        // null，不会报错
let displayName = user?.name || "游客"; // "游客"（可选链阶段四讲）
```

**位运算符**（了解即可，Node.js 后端偶用于权限位运算）：`& | ^ ~ << >> >>>`。

### 1.6 流程控制

**条件语句**：

```javascript
let time = 14;
if (time < 10) {
    console.log("早上好");
} else if (time < 20) {
    console.log("今天好");
} else {
    console.log("晚上好");
}
```

**switch**（注意 break 与 default，不写 break 会"穿透"执行下一个 case）：

```javascript
const day = new Date().getDay();   // 0=周日, 1=周一 ...
switch (day) {
    case 6:
        console.log("星期六");
        break;
    case 0:
        console.log("星期日");
        break;
    default:
        console.log("期待周末");
}
```

**循环**：

```javascript
// for 经典三段式（初始化; 条件; 步进）——三段均可省略
for (let i = 0; i < 5; i++) {
    console.log(i);
}

// while：条件为真一直执行（谨防死循环：条件里的变量必须变化）
let count = 0;
while (count < 5) {
    count++;
}

// do...while：先执行一次再判断条件（至少执行一次）
let j = 10;
do {
    console.log(j);   // 输出 10，即使条件一开始就是假
} while (j < 5);

// for...in：遍历对象的键（也常用于数组，但不推荐，见下）
const person = { fname: "Bill", lname: "Gates", age: 62 };
for (const key in person) {
    console.log(key, person[key]);   // fname Bill / lname Gates / age 62
}

// for...of：遍历可迭代对象的值（数组首选，阶段七讲迭代器协议时深入）
const cars = ["BMW", "Volvo", "Saab"];
for (const car of cars) {
    console.log(car);
}
```

**break / continue**：break 跳出整个循环，continue 跳过本轮进入下一轮。

```javascript
for (let i = 0; i < 10; i++) {
    if (i === 3) break;      // i 为 3 时整个循环结束
    if (i % 2 === 0) continue; // 跳过偶数
    console.log(i);          // 输出 1
}
```

**模板字符串**（拼字符串的正确姿势，反引号 + `${}`）：

```javascript
const name = "RUNOOB";
const age = 30;
const msg = `My name is ${name} and I'm ${age} years old.`;
const multi = `第一行
第二行`;                       // 反引号内可直接换行
const total = `Total: ${(10 * 1.25).toFixed(2)}`;  // ${} 里可以放任意表达式
```

**注释**：单行 `//`，多行 `/* ... */`。

### 阶段一练习

- [ ] 安装 Node.js、TypeScript、VS Code，`tsc -v` 能输出版本号
- [ ] 编写 `hello.ts`，手动走完「tsc 编译 → 查看 hello.js → node 运行」三步，观察类型注解在产物中消失
- [ ] 分别用 `tsc --watch` 和 `--removeComments` 各编译一次，观察行为差异
- [ ] 在 JS 文件里复现 var 的三个坑（块泄漏、覆盖、setTimeout 打印 3 3 3），再用 let 修复
- [ ] 写一个脚本：给定一个数字数组，用 for 循环求最大值；再用 `for...of` 求所有偶数的和
- [ ] 用 switch 写一个函数：输入 1~7 返回星期几的中文名，超出范围返回"无效"
- [ ] 验证并记录：`typeof null`、`typeof []`、`0.1 + 0.2 === 0.3` 各是什么结果

### 阶段一过关自检

1. TypeScript 代码能直接在 Node.js 上运行吗？为什么？描述 .ts 从编写到运行的完整链路。
2. var、let、const 三者在作用域和重复赋值上的区别是什么？"var + setTimeout 打印 3 3 3"的原因是什么？
3. `const obj = {}` 之后还能给 obj 加属性吗？为什么？
4. `==` 和 `===` 的区别是什么？项目里应该用哪个？
5. `typeof null` 返回什么？正确判断数组和 null 的方法各是什么？
6. break 和 continue 的区别？do...while 和 while 的区别？

---

## 阶段二：JavaScript 核心（下）——函数、对象、数组与异步

> **目标**：掌握 JS 的函数（含箭头函数与 this）、对象与原型链、数组高阶方法、解构与展开、JSON、错误处理，以及贯穿 Node.js 后端核心的异步编程（回调 → Promise → async/await）。这是 JS 底座中最重要的一段。

### 2.1 函数 ⭐

**声明与调用**：

```javascript
// 函数声明
function add(a, b) {
    return a + b;
}
console.log(add(1, 2));   // 3

// 函数表达式（匿名函数赋给变量）
const mul = function (a, b) {
    return a * b;
};

// 有返回值的函数：return 后函数立即停止；无 return 的函数返回 undefined
function greet(name) {
    return "Hello, " + name;
}
```

**参数的特性**：

- JS 传参个数不匹配**不报错**：少传的参数是 undefined，多传的直接被忽略：

```javascript
function intro(name, job) {
    console.log(`${name}, the ${job}`);
}
intro("Harry");              // "Harry, the undefined"
intro("Harry", "Wizard", 1); // 第三个参数被忽略
```

- **默认参数**（ES6）：

```javascript
function calc(price, rate = 0.5) {
    return price * rate;
}
calc(1000);      // 500
calc(1000, 0.3); // 300
```

- **剩余参数**（`...` 收集不定数量的参数为数组）：

```javascript
function sum(...nums) {
    let total = 0;
    for (const n of nums) {   // nums 就是一个普通数组，可遍历
        total += n;
    }
    return total;
}
sum(1, 2, 3);        // 6
sum(10, 10, 10, 10); // 40
```

- **递归**：

```javascript
function factorial(n) {
    if (n <= 0) return 1;
    return n * factorial(n - 1);
}
console.log(factorial(6)); // 720
```

**作用域与闭包（概念先建立）**：

```javascript
// 局部变量：函数内声明，函数外不可见
function f() {
    const local = "函数内";
}
// console.log(local); // ReferenceError

// 闭包：内部函数可以访问外部函数的变量，即使外部函数已执行完毕
function makeCounter() {
    let count = 0;                      // count 被闭包"记住"
    return function () {
        return ++count;
    };
}
const counter = makeCounter();
counter(); // 1
counter(); // 2 —— count 的状态被保留了
```

闭包是 JS 原型链、模块化、回调的底层机制，后续阶段（类、模块）会反复见到它。

### 2.2 箭头函数 ⭐

```javascript
// 完整形式
const add = (a, b) => {
    return a + b;
};
// 单表达式可省略 {} 和 return
const add2 = (a, b) => a + b;
// 单参数可省略括号
const double = n => n * 2;
// 无参数必须写空括号
const hello = () => console.log("hi");
```

箭头函数与普通函数的核心区别：**普通函数的 this 在调用时确定，箭头函数的 this 在定义时捕获外层作用域的 this**。看这个经典对比：

```javascript
function Person1() {
    this.name = "Alice";
    setTimeout(function () {
        // 普通函数：this 指向调用者（这里是定时器环境），拿不到 name
        console.log("普通函数:", this.name);   // undefined
    }, 100);
}

function Person2() {
    this.name = "Bob";
    setTimeout(() => {
        // 箭头函数：捕获定义时外层的 this（即 Person2 的实例）
        console.log("箭头函数:", this.name);   // Bob
    }, 100);
}
new Person1();
new Person2();
```

> 💡 伏笔：类的方法里用箭头函数绑定 this 的实战场景，在阶段四"类"中兑现。

### 2.3 对象、原型与 this ⭐

**对象字面量**（JS 中组织数据的第一方式）：

```javascript
const person = {
    firstName: "John",
    lastName: "Doe",
    age: 50,
    greet: function () {              // 方法：值是函数的属性
        return `Hi, ${this.firstName}`;
    },
};
person.firstName;       // 点访问
person["lastName"];     // 方括号访问（键是变量或含特殊字符时必须用）
person.greet();         // 调用方法
person.email = "a@b.c"; // 动态添加属性
```

**this 的规则**（JS 面试必考，Node 后端回调中也常见）：

| 场景 | this 指向 |
|---|---|
| 对象方法中 | 调用该方法的对象 |
| 单独使用 / 普通函数中 | 全局对象（Node 中是 global）；严格模式下 undefined |
| 箭头函数中 | 定义时外层作用域的 this |
| call/apply 显式绑定 | 被指定的对象 |

```javascript
const p1 = {
    fullName: function () { return this.firstName + " " + this.lastName; }
};
const p2 = { firstName: "John", lastName: "Doe" };
p1.fullName.call(p2);   // "John Doe" —— call 把 this 绑到 p2 上
```

**原型与原型链**（理解到"能解释方法从哪来"的程度即可）：

```javascript
function Person(first, last) {          // 构造函数（类语法出现前的"类"）
    this.firstName = first;
    this.lastName = last;
}
// 所有实例共享的方法挂在原型上
Person.prototype.getName = function () {
    return this.firstName + " " + this.lastName;
};

const bob = new Person("Bob", "Dai");
bob.getName();          // Bob Dai —— bob 自身没有 getName，沿原型链找到

// 访问属性时：先在对象自身找 → 找不到去它的原型找 → 再去原型的原型找 → 直到 Object.prototype（null）
// 这就是"原型链"。数组的方法（push/map/...）都挂在 Array.prototype 上
```

`Object.create(proto)` 可以显式指定原型创建对象。现代开发直接用 `class`（阶段四），但 class 本质仍是这套原型机制。

### 2.4 数组与高阶方法 ⭐

```javascript
const arr = ["Saab", "Volvo", "BMW"];   // 字面量创建（推荐）
arr[0];            // "Saab"（下标从 0 开始）
arr.length;        // 3
arr[arr.length - 1]; // 最后一个元素
arr[9];            // 不存在的下标返回 undefined
```

**常用方法速查**：

| 方法 | 作用 | 返回 |
|---|---|---|
| push / pop | 尾部追加 / 弹出 | 新长度 / 弹出元素 |
| unshift / shift | 头部插入 / 移除 | 新长度 / 移除元素 |
| indexOf / includes | 查找下标 / 是否包含 | 下标或 -1 / 布尔 |
| slice(start, end) | 浅拷贝片段（不改原数组） | 新数组 |
| splice(start, n, ...items) | 删除并插入（**改原数组**） | 被删元素数组 |
| concat | 拼接数组 | 新数组 |
| join(sep) | 拼成字符串 | 字符串 |
| reverse / sort | 反转 / 排序（**改原数组**） | 原数组 |
| find / findIndex | 按条件找第一个元素/下标 | 元素或 undefined |
| forEach | 遍历（无返回） | undefined |
| map | 逐个变换 | 新数组 |
| filter | 过滤 | 新数组 |
| reduce | 归纳累积 | 累积值 |

**高阶三件套**（Node 后端数据处理天天用）：

```javascript
const nums = [1, 2, 3, 4, 5];

nums.map(n => n * 2);              // [2, 4, 6, 8, 10]
nums.filter(n => n % 2 === 0);     // [2, 4]
nums.reduce((sum, n) => sum + n, 0); // 15（0 是初始值）
nums.find(n => n > 3);             // 4
const users = [
    { name: "A", age: 30 },
    { name: "B", age: 20 },
];
users.sort((a, b) => a.age - b.age); // 按 age 升序（比较函数返回正负数）
```

### 2.5 字符串、JSON 与内置对象

**字符串常用操作**：

```javascript
const s = "Hello World";
s.length;              // 11
s[0];                  // "H"（下标访问，字符串不可变，s[0]="X" 无效）
s.indexOf("World");    // 6（找不到返回 -1）
s.includes("lo");      // true
s.slice(0, 5);         // "Hello"（切片）
s.toUpperCase();       // "HELLO WORLD"
s.toLowerCase();
"  hi  ".trim();       // "hi"（去两端空白——处理用户输入必做）
"a,b,c".split(",");    // ["a", "b", "c"]
s.replace("World", "TS"); // "Hello TS"（只替换第一个匹配）
"5".padStart(2, "0");  // "05"（补零，格式化日期常用）
```

**Math 与 Number**：

```javascript
Math.round(4.7);        // 5（四舍五入）
Math.floor(4.7);        // 4（向下取整）
Math.ceil(4.1);         // 5（向上取整）
Math.random();          // [0, 1) 随机数
Math.floor(Math.random() * 11); // 0~10 的随机整数
Math.max(1, 5, 3);      // 5
Math.abs(-7);           // 7
Number.parseInt("42px");  // 42
Number.parseFloat("3.14abc"); // 3.14
(3.14159).toFixed(2);   // "3.14"（注意返回字符串！）
Number.isInteger(5);    // true
Number.isNaN(NaN);      // true（可靠的 NaN 判断，NaN === NaN 是 false）
```

**Date 基础**：

```javascript
const now = new Date();              // 当前时间
now.getFullYear();                   // 年（getMonth() 返回 0~11，要 +1！）
now.getDate();                       // 日
now.toISOString();                   // "2024-01-15T08:00:00.000Z"（API 时间标准格式）
Date.now();                          // 当前毫秒时间戳
```

**JSON**（后端接口的数据交换格式）：

```javascript
// 对象 → JSON 字符串
const obj = { name: "RUNOOB", tags: ["a", "b"], nested: { x: 1 } };
const json = JSON.stringify(obj);   // '{"name":"RUNOOB","tags":["a","b"],"nested":{"x":1}}'

// JSON 字符串 → 对象
const back = JSON.parse('{"sites":[{"name":"Runoob"}]}');
back.sites[0].name;                 // "Runoob"

// JSON 语法规则：键必须双引号；值只能是 string/number/boolean/null/对象/数组
// JSON.stringify 会丢弃 undefined 和函数
```

### 2.6 解构与展开 ⭐（官方 Handbook 内容，菜鸟教程未覆盖）

**数组解构**：

```javascript
const input = [1, 2];
const [first, second] = input;       // first=1, second=2

// 交换变量（不用临时变量）
let a = 1, b = 2;
[a, b] = [b, a];

// 剩余元素
const [head, ...rest] = [1, 2, 3, 4];  // head=1, rest=[2,3,4]

// 跳过元素
const [, second2, , fourth] = [1, 2, 3, 4];  // second2=2, fourth=4
```

**对象解构**：

```javascript
const o = { a: "foo", b: 12, c: "bar" };
const { a, b } = o;                    // a="foo", b=12（c 被忽略）

// 重命名：a: newName 读作 "a as newName"
const { a: newName } = o;              // newName = "foo"

// 默认值（属性为 undefined 或缺失时生效）
const withDefault = { a: "foo" };
const { a: fromObj, b: bVal = 1001 } = withDefault;   // bVal → 1001

// 剩余属性
const { a: omitted, ...passthrough } = o;
// passthrough = { b: 12, c: "bar" } —— 剔除属性的经典手法
```

**函数参数解构**（配置对象的最常见用法）：

```javascript
// 参数直接解构 + 默认值
function draw({ shape, xPos = 0, yPos = 0 }) {
    console.log(shape, xPos, yPos);
}
draw({ shape: "circle" });            // circle 0 0
```

⚠ 易错：对象解构的冒号是"重命名"不是"类型"；整个解构语句作为表达式赋值时要加括号 `({ a, b } = { a: 1, b: 2 })`，否则 `{` 被解析为块语句。

**展开运算符（spread，解构的逆操作）**：

```javascript
// 数组展开
const firstArr = [1, 2], secondArr = [3, 4];
const bothPlus = [0, ...firstArr, ...secondArr, 5];  // [0,1,2,3,4,5]

// 对象展开：浅拷贝与覆盖（后展开的覆盖先展开的同名属性）
const defaults = { host: "localhost", port: 3000, debug: false };
const config = { ...defaults, port: 8080 };
// { host: "localhost", port: 8080, debug: false }

// 展开是浅拷贝：嵌套对象仍是引用
// 展开会丢失对象的方法（只保留自有可枚举属性）
```

### 2.7 错误处理

```javascript
// try / catch / finally
function parse(input) {
    try {
        const n = Number(input.trim());
        if (input === "") throw new Error("输入不能为空");
        if (isNaN(n)) throw new Error("不是有效数字");
        if (n > 100) throw new RangeError(`数字 ${n} 太大了`);
        return n;
    } catch (err) {
        console.error("错误:", err.message);   // Error 对象有 message 属性
        return null;
    } finally {
        console.log("处理完成");  // 无论成功失败都执行（清理资源的去处）
    }
}
parse("abc");   // 错误: 不是有效数字 / 处理完成
```

- `throw` 可以抛出任意值，但**工程实践永远抛 `new Error("描述")` 或其子类**。
- `finally` 里的代码总会执行，常用于关闭文件、释放连接。

### 2.8 异步编程 ⭐⭐（Node.js 后端的生命线）

**为什么需要异步**：JS 是单线程语言，一次只能执行一个任务。读文件、网络请求这类耗时操作如果同步等待，整个程序（服务器）就会卡死。异步的思路是：发起耗时操作后**不等待**，先继续执行后面的代码，操作完成后再回来处理结果。

```javascript
// 执行顺序是理解异步的第一关：
console.log("1");
setTimeout(() => console.log("2"), 0);   // 回调进入任务队列，稍后执行
console.log("3");
// 输出：1 3 2 —— setTimeout 即使延迟 0 也会等主线程代码跑完
```

**第一代：回调函数**——"完成后调用这个函数"：

```javascript
setTimeout(function () {
    console.log("3 秒后执行");
}, 3000);
// Node.js 读文件也是回调风格：fs.readFile(path, (err, data) => { ... })
```

回调的致命问题——**回调地狱**（嵌套依赖导致代码不可读、错误难处理）：

```javascript
getData(function (a) {
    getMoreData(a, function (b) {
        getMoreData(b, function (c) {
            console.log(c);   // 向右无限膨胀……
        });
    });
});
```

**第二代：Promise**——代表"一个异步操作的最终结果"的对象。三种状态：`pending`（进行中）→ `fulfilled`（成功）或 `rejected`（失败），**状态一旦改变不可逆**。

```javascript
// 创建
const p = new Promise((resolve, reject) => {
    const success = true;
    if (success) {
        resolve("成功的结果");      // pending → fulfilled
    } else {
        reject(new Error("失败原因")); // pending → rejected
    }
});

// 消费：then 处理成功，catch 处理失败，finally 总会执行
p.then(result => console.log(result))
 .catch(err => console.error(err.message))
 .finally(() => console.log("结束"));

// 链式调用：每个 then 返回新 Promise，把值传给下一个
Promise.resolve(1)
    .then(n => n * 2)       // 2
    .then(n => n + 10)      // 12
    .then(n => console.log(n));  // 12
```

**Promise 静态方法**：

```javascript
// Promise.all：等全部成功（任一失败则整体失败）
Promise.all([fetchUser(1), fetchUser(2)])
    .then(([u1, u2]) => { /* 结果数组，顺序与传入一致 */ });

// Promise.allSettled：等全部结束（无论成败），返回每个的状态与值
Promise.allSettled([p1, p2, p3]).then(results => {
    results.forEach(r => console.log(r.status)); // "fulfilled" / "rejected"
});

// Promise.race：第一个完成的（无论成败）的结果 —— 实现超时控制的经典手段
// Promise.resolve / Promise.reject：快速创建已定型的 Promise
```

**第三代：async/await（ES2017）**——用同步的写法写异步，**本质仍是 Promise**：

```javascript
// async 函数：总是返回 Promise；返回值自动包装，抛错自动变成 rejected
async function getData() {
    return { name: "Alice" };      // 等价于 Promise.resolve({name: "Alice"})
}

// await：只能在 async 函数内使用（模块顶层也可，ES2022）
// 等待 Promise 完成并取出值；rejected 则抛出异常
function delay(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function main() {
    console.log("开始");
    await delay(1000);             // 暂停 1 秒（不阻塞主线程）
    const data = await getData();
    console.log(data.name);
}
main();
```

**错误处理用 try/catch**：

```javascript
async function fetchUser() {
    try {
        const result = await mayFail();
        return result;
    } catch (error) {
        console.error("捕获:", error.message);
        return null;
    }
}
```

**并行 vs 串行**（性能关键）：

```javascript
// ❌ 串行：两次请求依次等待，总耗时 = 100 + 200
const u1 = await fetchUser(1);
const u2 = await fetchUser(2);

// ✅ 并行：同时发起，总耗时 = max(100, 200)
const [u1, u2] = await Promise.all([fetchUser(1), fetchUser(2)]);
```

**最佳实践**：现代代码一律用 async/await；忘记写 await 会拿到 Promise 对象而不是值（值是 undefined 或 [object Promise]）；能用 Promise.all 并行就不要串行 await。

### 阶段二练习

- [ ] 写一个 `makeCounter()` 闭包函数，连续调用输出 1、2、3
- [ ] 用 map/filter/reduce 一行式实现：给定商品数组 `[{name, price}]`，取出价格 > 100 的商品的名称数组，并计算它们的总价
- [ ] 写函数 `formatDate(date)` 返回 `"2024-01-15"` 格式（padStart 练习）
- [ ] 解构练习：从 `const config = { db: { host, port, user }, redis: {...} }` 中一次取出 host/port 并重命名为 dbHost/dbPort
- [ ] 写一个返回 Promise 的 `delay(ms)` 函数，然后串行执行两次、再用 Promise.all 并行执行两次，用计时对比耗时差异
- [ ] 实现 `safeJsonParse(str)`：解析失败返回 null 而不是抛错
- [ ] 用 async/await + try/catch 重写"依次读取三个模拟接口，任一失败打印错误并返回 null"的流程

### 阶段二过关自检

1. 箭头函数和普通函数在 this 上的区别？什么场景必须用箭头函数？
2. 描述原型链：访问 `obj.foo` 时查找顺序是怎样的？`bob.getName()` 中方法实际存放在哪里？
3. map、filter、reduce 各自的返回值是什么？哪些数组方法会修改原数组？
4. 对象解构中 `{ a: newName }` 的冒号是什么意思？展开运算符拷贝是深还是浅？
5. Promise 的三种状态是什么？状态能从 fulfilled 变回 pending 吗？
6. Promise.all 和 Promise.allSettled 的区别？分别适合什么场景？
7. `async` 函数的返回值是什么类型？`await` 一个 rejected 的 Promise 会发生什么？
8. 串行 await 和 Promise.all 并行的区别是什么？

---
## 阶段三：TypeScript 入门——类型注解与基础类型

> **目标**：正式进入 TypeScript。掌握类型注解语法、类型推断的边界、全部基础类型与特殊类型（any/unknown/void/never）、元组、枚举、联合/交叉/字面量类型。学完本阶段，你能为任意变量和函数写出正确的类型。

### 3.1 类型系统是怎么工作的 ⭐

官方 Handbook 的核心观点：**类型 = 描述一个值能做什么**。静态类型系统在代码运行**之前**就对"哪些操作合法"做出预测：

```typescript
// 访问不存在的属性：JS 返回 undefined（静默失败），TS 直接报错
const user = {
    name: "Daniel",
    age: 26,
};
user.location;  // 错误：Property 'location' does not exist

// 拼写错误：TS 帮你抓 typo
"Hello".toLocaleLowercase();   // 错误：应该是 toLocaleLowerCase

// 未调用的函数
Math.random < 0.5;             // 错误：Math.random 是函数不是值

// 逻辑错误：不可达分支
const value = Math.random() < 0.5 ? "a" : "b";
if (value !== "a") {
} else if (value === "b") {    // 错误：这个分支永远不可能到达
}
```

三个必须建立的认知：

1. **类型是可选的但强烈推荐**：TS 完全兼容 JS，不加类型也能编译，但没有类型的 TS 约等于 JS。
2. **类型会被擦除**：编译产物里没有任何类型信息，`string`、`interface` 等全部消失。运行时无法依赖类型做判断——需要运行时校验要用代码自己写。
3. **TS 是结构化类型系统**：只看"形状"是否匹配，不要求显式声明继承关系（阶段六深入）。这也是它与 Java/C# 名义类型系统的本质区别。

### 3.2 类型注解与类型推断 ⭐

**注解语法**：`变量名: 类型`，函数参数和返回值同理：

```typescript
let age: number = 25;
let username: string = "runoob";
let isActive: boolean = true;

function add(a: number, b: number): number {
    return a + b;
}

// 无返回值的函数用 void
function log(msg: string): void {
    console.log(msg);
}
```

**类型推断**：TS 会根据上下文自动确定类型，不必处处手写：

```typescript
let count = 0;           // 推断为 number
let name = "RUNOOB";     // 推断为 string
const flag = true;       // const + 字面量 → 推断为字面量类型 true

let numbers = [1, 2, 3]; // 推断为 number[]
const config = { host: "localhost", port: 3000 };
// 推断为 { host: string; port: number }

// 函数返回值推断
function double(n: number) {   // 返回类型自动推断为 number
    return n * 2;
}

// 上下文推断：回调参数类型由所在位置决定
const doubled = numbers.map(n => n * 2);  // n 自动推断为 number
```

**推断失败时**：没有初始值、没有上下文的参数会退化为 `any`（在开启 `noImplicitAny` 的严格模式下报错，提示你显式标注）。

> **实践准则**：函数参数必须显式标注（调用方无法推断）；局部变量通常靠推断，只在推断不出来或想让类型更宽时才标注。过度标注（`const name: string = "Alice"`）是反模式。

### 3.3 基础类型全家福 ⭐

| 类型 | 描述 | 示例 |
|---|---|---|
| `string` | 文本 | `let s: string = "hi"` |
| `number` | 数字（整数+浮点，不分家） | `let n: number = 3.14` |
| `boolean` | 布尔 | `let b: boolean = true` |
| `array` | 数组，两种写法 | `number[]` 或 `Array<number>` |
| `tuple` | 定长定类型的数组 | `[string, number]` |
| `enum` | 命名常量集合 | `enum Color { Red }` |
| `object` | 非原始类型 | `{ name: string }` |
| `any` | 任意类型（逃生舱） | 见 3.4 |
| `unknown` | 不确定类型（安全版 any） | 见 3.4 |
| `void` | 无返回值 | `function f(): void` |
| `null` / `undefined` | 空值 / 未定义 | 见下 |
| `never` | 永不出现的值 | 见 3.4 |

**数组**：

```typescript
let scores: number[] = [90, 85, 92];
let tags: Array<string> = ["ts", "js"];   // 泛型写法，与上面等价
let mixed: (string | number)[] = [1, "a"]; // 联合类型数组
```

**null 与 undefined**：默认情况下 null/undefined 是所有类型的子类型（`let x: number` 可以赋 null）。但开启 `strictNullChecks`（strict 模式的一部分，**生产项目必开**）后：

```typescript
let x: number;
x = 1;          // 正确
x = undefined;  // 错误：undefined 不能赋给 number
x = null;       // 错误

// 一个变量确实可能为空时，显式声明联合：
let y: number | null | undefined;
y = 1; y = null; y = undefined;   // 都合法
```

**object 类型的局限**：`object` 只表示"非原始类型"，访问具体属性仍会报错，实际中几乎总是用接口或类型字面量描述对象形状（阶段四）。

### 3.4 特殊类型：any / unknown / void / never ⭐

```typescript
// any：关闭类型检查的逃生舱——变量可以是任何类型、调用任何方法都不报错
let anything: any = 42;
anything = "hello";        // 合法
anything.foo.bar();        // 编译不报错，运行时可能爆炸
// 适用场景仅限：渐进迁移的老代码、确实无法确定类型的第三方数据
// 工程铁律：能不用就不用；开启 noImplicitAny 防止隐式 any

// unknown：安全版 any——可以接收任何值，但使用前必须收窄
let value: unknown = "Hello";
// value.length;          // 错误：unknown 不能直接访问属性
if (typeof value === "string") {
    console.log(value.length);  // 在守卫分支内合法
}
// 处理外部输入（JSON.parse、请求体）首选 unknown，而不是 any

// void：函数没有返回值
function log(msg: string): void { console.log(msg); }
// 声明 void 变量没有意义（只能赋 null/undefined）

// never：永不出现的值。两种典型来源：
function throwError(message: string): never {
    throw new Error(message);   // 抛错——函数永远不会正常返回
}
function infiniteLoop(): never {
    while (true) {}             // 死循环——永远走不到结尾
}
```

**对比记忆**：

| 对比 | 结论 |
|---|---|
| any vs unknown | 两者都能装任意值；any 放弃检查，unknown 强制先收窄再使用。**新代码一律 unknown** |
| never vs void | void = "没有返回值"（函数正常结束但无值）；never = "根本不会有返回"（抛错/死循环） |
| never 的妙用 | 它是所有类型的子类型；联合类型收窄到只剩 never 说明分支穷尽了（阶段四的穷尽检查） |

### 3.5 元组（Tuple）

已知长度、每个位置类型固定的数组：

```typescript
let person: [string, number] = ["Alice", 25];
let point: [number, number] = [10, 20];

person[0];         // string 类型
person[1];         // number 类型
// person[2];      // 错误：长度固定

// 元组可以解构（类型按位置对应）：
const [name2, age2] = person;   // name2: string, age2: number

// 越界解构报错（普通数组不会）：
// const [a, b, c] = person;    // 错误：没有第 3 个元素

// 带标签的可选元组（API 场景）：
type Range = [start: number, end: number, step?: number];
```

适合表达"固定结构的小数据"：坐标、键值对、HTTP 状态行等。**超过 3 个元素或结构会变化时，应该用对象类型。**

### 3.6 枚举（Enum）

枚举定义一组命名常量，消灭魔法数字/字符串：

```typescript
// 数字枚举：默认从 0 自增；也可指定起点
enum Direction {
    Up = 1,     // 显式从 1 开始
    Down,       // 2
    Left,       // 3
    Right,      // 4
}
Direction.Up;        // 1
Direction[1];        // "Up" —— 反向映射（数字枚举特有）

// 字符串枚举：每个成员必须显式初始化，更推荐（调试时值可读、序列化友好）
enum Color {
    Red = "RED",
    Green = "GREEN",
}
function paint(color: Color): void {
    console.log(`Painting in ${color}`);
}
paint(Color.Red);      // Painting in RED
// paint("red");       // 错误：字符串 "red" 不是 Color 类型

// const 枚举：编译时完全内联，不生成运行时对象（体积更小）
const enum HttpStatus {
    OK = 200,
    NotFound = 404,
}
const status = HttpStatus.OK;   // 编译后直接变成 const status = 200

// 异构枚举（字符串+数字混用）：语法上允许，实践中不要用
```

⚠ 枚举是 TS 里少数**编译后产生运行时代码**的类型特性（普通类型都会被擦除）。

> 💡 简单场景的替代品是字面量联合类型（3.8 节）：`type Direction = "up" | "down"`，零运行时成本。枚举与字面量联合的选择在阶段七工程化时再权衡。

### 3.7 联合类型与交叉类型 ⭐

**联合类型 `|`**："或"——值是多种类型之一：

```typescript
type ID = string | number;
let userId: ID = "abc-123";
userId = 456;             // 合法
// userId = true;         // 错误：boolean 不在联合范围内

// 联合类型数组：数组本身是 number[] 或 string[]（不是混合数组）
let arr: number[] | string[];
arr = [1, 2, 4];
arr = ["Runoob", "Google"];

// 使用联合类型的值时，只能访问所有成员共有的属性：
function printId(id: string | number) {
    console.log(`ID: ${id}`);
    // id.toUpperCase();  // 错误：number 没有这个方法
    if (typeof id === "string") {
        console.log(id.toUpperCase());  // 收窄后在分支内可用（阶段四详解）
    }
}
```

**交叉类型 `&`**："与"——同时具备多个类型的全部成员：

```typescript
interface Person {
    name: string;
    age: number;
}
interface Worker {
    company: string;
    salary: number;
}
type Employee = Person & Worker;   // 四个属性全都要有，缺一不可

const emp: Employee = {
    name: "Alice", age: 25, company: "Google", salary: 100000,
};
```

⚠ 交叉不兼容的类型得到 never：`type Bad = string & number;`（不存在一个值既是 string 又是 number）。同名但类型冲突的属性交叉后也会变成 never——这是交叉类型最常见的坑。

### 3.8 字面量类型与 as const ⭐

限制变量只能取特定值，是联合类型最实用的形态：

```typescript
// 字符串字面量联合：状态机的正确表达
type Status = "pending" | "active" | "completed";
let s: Status = "active";
// s = "actived";   // 错误：拼写错误直接被编译器抓住

// 数字字面量
type Weekday = 1 | 2 | 3 | 4 | 5 | 6 | 7;
let today: Weekday = 1;

// 布尔字面量：boolean 本质就是 true | false 的别名
let flag: true | false;

// 函数参数是最常见用法：
function setDirection(dir: "up" | "down" | "left" | "right") { }
setDirection("up");
// setDirection("diagonal");  // 错误
```

**const 推断与 as const**：

```typescript
// let 推断为宽类型，const 推断为最窄字面量
let a = "hello";     // string（还能重新赋值）
const b = "hello";   // "hello"（不可能变了）

// as const：把整个对象/数组冻结为只读字面量
const colors1 = ["red", "green"];            // string[]
const colors2 = ["red", "green"] as const;   // readonly ["red", "green"]
// colors2.push("yellow");   // 错误：只读数组不能修改
// colors2[0] = "blue";      // 错误

const config = { host: "localhost", port: 3000 } as const;
// config 的类型：{ readonly host: "localhost"; readonly port: 3000 }
// config.port = 8080;       // 错误：只读
```

**判别字段的雏形**（阶段四可辨识联合的基础）：

```typescript
interface Circle { kind: "circle"; radius: number; }
interface Rect { kind: "rect"; width: number; height: number; }
type Shape = Circle | Rect;
// kind 是"判别属性"，用它区分联合成员——阶段四会展开完整玩法
```

### 阶段三练习

- [ ] 给阶段二的 `sum`、`formatDate`、`safeJsonParse` 补上完整类型注解（参数+返回值）
- [ ] 声明一个 `User` 形状的对象字面量类型（含 name: string、age: number、email?: string），创建两个合法实例（一个带 email 一个不带）
- [ ] 定义 `httpStatus` 联合类型：200 | 404 | 500，写函数根据状态码返回对应消息（用 switch）
- [ ] 分别用枚举和字符串字面量联合实现"订单状态"，对比编译产物差异（打开 .js 文件看）
- [ ] 写一个函数 `format(value: string | number)`：字符串转大写，数字保留两位小数（先用 typeof 分支——类型收窄的预演）
- [ ] 用 `as const` 定义一个颜色常量表，尝试修改元素验证编译报错
- [ ] 声明 `[string, number, boolean?]` 类型的元组并解构前两个元素

### 阶段三过关自检

1. 类型注解会出现在编译产物中吗？这对"运行时校验外部数据"意味着什么？
2. 类型推断什么时候会退化成 any？怎么避免？
3. any 和 unknown 的区别？处理 `JSON.parse` 的结果应该用哪个，为什么？
4. void 和 never 的区别？各举一个返回它们的函数。
5. strictNullChecks 开启后，`let x: number` 还能赋 null 吗？正确的可空类型怎么写？
6. `string & number` 的结果是什么？为什么？
7. `const x = "hello"` 推断出的类型是什么？`let` 呢？
8. 数字枚举和字符串枚举各自的特点？const 枚举编译后变成什么？

---

## 阶段四：TypeScript 核心——函数类型、接口、类与类型收窄

> **目标**：掌握 TS 四大核心：函数类型的完整表达（含重载）、接口（含索引签名与可辨识联合）、类（修饰符/继承/抽象类/静态成员/#private）、类型收窄（类型守卫全家桶）。这是日常开发 80% 的类型场景。

### 4.1 函数类型的完整表达 ⭐⭐

**函数类型表达式**——描述"函数长什么样"：

```typescript
// (参数: 类型) => 返回类型
type Transformer = (input: string) => number;
const toNumber: Transformer = (s) => parseInt(s, 10);

// 作为参数（回调的典型场景）
function greeter(fn: (a: string) => void) {
    fn("Hello, World");
}
function printToConsole(s: string) { console.log(s); }
greeter(printToConsole);
```

⚠ 参数名**必填**：`(string) => void` 的意思是"一个参数名叫 string、类型 any 的函数"。

**调用签名与构造签名**——函数本身也是对象，可以带属性；也可以被 new：

```typescript
// 调用签名：描述"可调用 + 有属性"（注意用 : 不是 =>）
type DescribableFunction = {
    description: string;
    (someArg: number): boolean;
};

// 构造签名：描述"可用 new 调用"
type SomeConstructor = {
    new (s: string): SomeObject;
};
// 两者可共存（如 Date 既可直接调用又可 new）：
interface CallOrConstruct {
    (n?: number): string;
    new (s: string): Date;
}
```

**参数的四种形态**：

```typescript
// 1. 可选参数 ?（必须在必需参数之后）
function buildName(firstName: string, lastName?: string) {
    return lastName ? `${firstName} ${lastName}` : firstName;
}
buildName("Bob");              // 合法
buildName("Bob", "Adams");     // 合法
// buildName("Bob", "Adams", "Sr.");  // 错误：参数过多

// 2. 默认参数（等价于可选，但函数体内一定有值）
function calc(price: number, rate: number = 0.5) {
    return price * rate;
}

// 3. 剩余参数 ...rest
function buildAll(firstName: string, ...rest: string[]) {
    return firstName + " " + rest.join(" ");
}
buildAll("Joseph", "Samuel", "Lucas");

// 4. 参数解构 + 类型（类型写在整个解构模式之后）
function draw({ shape, xPos = 0, yPos = 0 }: { shape: string; xPos?: number; yPos?: number }) {
    console.log(shape, xPos, yPos);
}
```

⚠ **回调参数不要写可选**：`(callback: (arg: any, index?: number) => void)` 的含义是"实现可能只传一个参数"，这会让使用 `index.toFixed()` 的回调报错。写回调类型时，除非你确实打算不传该参数，否则不要加 `?`（参数少的函数永远可以赋给参数多的函数类型）。

**函数重载**——同名函数的多种调用形态：先写若干**重载签名**（只有声明），最后写一个**实现签名**（参数用宽类型兼容所有形态）：

```typescript
function makeDate(timestamp: number): Date;                    // 重载 1
function makeDate(m: number, d: number, y: number): Date;      // 重载 2
function makeDate(mOrTimestamp: number, d?: number, y?: number): Date {
    if (d !== undefined && y !== undefined) {
        return new Date(y, m - 1, d);
    }
    return new Date(mOrTimestamp);
}
makeDate(1700000000000);     // 匹配重载 1
makeDate(1, 15, 2024);       // 匹配重载 2
// makeDate(1, 15);          // 错误：不匹配任何重载签名
```

⚠ 实现签名对外不可见；调用方只能按重载签名的形态调用。能用联合类型/可选参数解决的，优先不用重载。

**箭头函数与 this（兑现阶段二的伏笔）**：类中把方法写成箭头函数属性，可保证 this 始终指向实例——回调场景不再丢失 this：

```typescript
class Counter {
    count = 0;
    // 箭头函数属性：this 永远绑定为实例
    increment = () => { this.count++; };
}
const c = new Counter();
setTimeout(c.increment, 100);   // 即使被拆出来传，this 也不丢
```

### 4.2 接口（Interface）⭐⭐

接口描述对象的"形状"——有哪些属性、什么类型、必选还是可选：

```typescript
interface User {
    id: number;                 // 必选
    name: string;
    email?: string;             // 可选（? 后缀）
    readonly role: string;      // 只读（首次赋值后不可改）
    sayHi: () => string;        // 方法成员
}

const admin: User = {
    id: 1,
    name: "RUNOOB",
    role: "admin",
    sayHi: () => "Hi",
};
// admin.role = "user";   // 错误：只读属性
```

**可选属性的读取**：可选意味着 `类型 | undefined`，使用前要处理空值（默认值是惯用法）：

```typescript
function greet(user: User) {
    const email = user.email ?? "未填写";   // 空值合并（4.6 节）
}
```

**索引签名**——键名不固定时的"字典"模式：

```typescript
interface StringArray {
    [index: number]: string;      // 数字索引 → string
}
interface Ages {
    [name: string]: number;       // 字符串索引 → number
}
const ages: Ages = {};
ages["runoob"] = 15;

// 字符串索引会约束所有属性的类型：
interface NumberDictionary {
    [key: string]: number;
    length: number;     // 合法：number 匹配
    // name: string;    // 错误：string 不匹配索引签名 number
}
// 解决：把索引签名写成联合 [key: string]: number | string

// 只读索引签名
interface ReadonlyStringArray {
    readonly [index: number]: string;
}
```

**接口继承**—— extends 支持单继承也支持多继承（逗号分隔）：

```typescript
interface Person {
    age: number;
}
interface Musician extends Person {
    instrument: string;
}
// 多继承
interface IParent1 { v1: number; }
interface IParent2 { v2: number; }
interface Child extends IParent1, IParent2 { }
const obj: Child = { v1: 12, v2: 23 };
```

**多余属性检查**：把对象**字面量**直接赋给接口类型时，不允许出现接口里没有的属性（这是字面量的特殊检查，非字面量对象不受限——原因见阶段六结构化类型）：

```typescript
interface SquareConfig { color?: string; width?: number; }
function createSquare(config: SquareConfig) { /* ... */ }

createSquare({ colour: "red", width: 100 });
// 错误：colour 不存在于 SquareConfig（很可能只是拼错 color，TS 在帮你）

const outer = { colour: "red", width: 100 };
createSquare(outer);   // 合法：变量不经过字面量检查（多余属性被忽略）
```

**可辨识联合（Discriminated Unions）⭐**——用判别字段（字面量类型）区分联合成员，是 TS 最推荐的数据建模模式：

```typescript
interface Circle {
    kind: "circle";            // 判别属性：字面量类型
    radius: number;
}
interface Rectangle {
    kind: "rectangle";
    width: number;
    height: number;
}
interface Triangle {
    kind: "triangle";
    base: number;
    height: number;
}
type Shape = Circle | Rectangle | Triangle;

function getArea(shape: Shape): number {
    switch (shape.kind) {              // kind 相当于运行时"标签"
        case "circle":
            return Math.PI * shape.radius ** 2;     // 此分支内 shape 是 Circle
        case "rectangle":
            return shape.width * shape.height;      // 此分支内是 Rectangle
        case "triangle":
            return 0.5 * shape.base * shape.height;
    }
}
```

对比反面教材——一个接口塞所有可选属性：`{ kind: ...; radius?: number; width?: number }`，每个分支都得判空、还可能出现 kind 与属性不匹配的非法状态。可辨识联合从类型层面杜绝了非法状态。

**接口 vs 类型别名（type）**：

| 对比项 | interface | type |
|---|---|---|
| 对象类型 | ✅ 主战场 | ✅ 也可以 |
| 联合/元组/原始类型别名 | ❌ | ✅ 独有能力 |
| 声明合并（同名自动合并） | ✅ 支持 | ❌ 重复定义报错 |
| extends 继承 | ✅ | 用 & 交叉模拟 |
| 编译性能 | 略优（简单对象类型推荐） | 复杂类型变换必须用 |

实践：公共对象形状用 interface；联合、工具类型、函数签名用 type；团队统一即可。

### 4.3 类（Class）⭐⭐

**基本结构**：字段 + 构造函数 + 方法：

```typescript
class Car {
    engine: string;                        // 字段声明
    constructor(engine: string) {          // 构造函数：new 时自动调用
        this.engine = engine;              // this 指向实例
    }
    disp(): void {                         // 方法
        console.log("发动机: " + this.engine);
    }
}
const car = new Car("V8");
car.disp();

// 字段初始化器：声明时直接赋值（可省构造函数）
class Point2D {
    x = 0;         // 推断为 number，实例创建时自动初始化
    y = 0;
}
```

⚠ **字段必须初始化**：strict 模式下，`name: string;` 没有在构造函数中赋值会报错（strictPropertyInitialization）。解决：构造函数里赋值、声明时给初始值、或用确定赋值断言 `name!: string;`（表示"我保证稍后会有值"）。

**访问修饰符**——控制成员从哪里可访问：

```typescript
class Animal {
    readonly name: string;          // 只读：构造后不可改（构造函数内可赋值）
    private age: number;            // 私有：仅类内部
    protected species: string;      // 受保护：类内部 + 子类
    public id: number;              // 公有（默认，可省略 public）

    constructor(name: string, age: number, species: string) {
        this.name = name;
        this.age = age;
        this.species = species;
    }
    public introduce(): string {    // 默认 public
        return `I'm ${this.name}`;
    }
    get info(): string {            // getter：像属性一样访问
        return `${this.name} (${this.age})`;
    }
}

class Dog extends Animal {
    private breed: string;
    constructor(name: string, age: number, breed: string) {
        super(name, age, "犬");     // 必须先调 super() 再用 this
        this.breed = breed;
    }
    describe(): string {
        // 子类可访问 protected，不可访问 private
        return `${this.name} is a ${this.species}, breed: ${this.breed}`;
    }
}
const dog = new Dog("旺财", 3, "金毛");
console.log(dog.info);      // getter 不加括号
// console.log(dog.age);    // 错误：private
```

**参数属性**——构造函数参数加修饰符，自动完成"声明 + 赋值"：

```typescript
class Point {
    constructor(
        public x: number,        // 自动生成 public x 字段并赋值
        public y: number,
        private z: number,       // 自动生成 private z
    ) {}
    getZ() { return this.z; }
}
const p = new Point(1, 2, 3);
console.log(p.x, p.getZ());
```

**private vs #private（软私有与硬私有）**：TS 的 `private` 只在编译期检查，运行时仍可通过 `obj["secret"]` 或 JS 代码访问（软私有）；JS 原生 `#field` 是运行时真私有（硬私有）：

```typescript
class MySafe {
    private secretKey = 12345;    // 编译期检查
    #realSecret = 67890;          // 运行时也拿不到
}
const s = new MySafe();
// s.secretKey;        // 编译错误；但 JS 里 s["secretKey"] 能拿到
// s.#realSecret;      // 语法级禁止，任何方式都拿不到
```

**静态成员**——属于类本身而非实例，通过类名访问：

```typescript
class Numbers {
    static sval = 10;              // 静态字段
    numVal = 13;                   // 实例字段
    static printSval(): void {
        console.log(Numbers.sval);
        // 静态方法内不能访问 this.numVal（实例成员不存在）
    }
}
Numbers.sval;          // 10（不需要 new）
Numbers.printSval();
```

**继承、重写与多态**：

```typescript
class Animal {
    constructor(protected name: string) {}
    speak(): void {
        console.log(`${this.name} 发出声音`);
    }
}
class Cat extends Animal {
    speak(): void {                          // 方法重写
        console.log(`${this.name} 喵喵喵`);
    }
}
class Dog extends Animal {
    speak(): void {
        console.log(`${this.name} 汪汪汪`);
        super.speak();                       // super 调用父类版本
    }
}

// 多态：父类型引用指向子类实例，调同一个方法各有表现
const animals: Animal[] = [new Cat("小白"), new Dog("旺财")];
animals.forEach(a => a.speak());

// instanceof：运行时检查实例类型（也是类型守卫，见 4.4）
const d = new Dog("旺财");
console.log(d instanceof Animal);   // true（原型链上有）
```

⚠ **初始化顺序陷阱**：JS 类的初始化顺序是"父类字段 → 父类构造 → 子类字段 → 子类构造"。父类构造函数里打印 this.name 时，看到的是父类字段的值，子类同名字段尚未初始化：

```typescript
class Base {
    name = "base";
    constructor() { console.log(this.name); }   // 打印 "base"
}
class Derived extends Base {
    name = "derived";   // 在父类构造之后才赋值
}
new Derived();          // 输出 "base" 而不是 "derived"
```

⚠ **子类必须遵循父类的契约**：重写方法不能比父类更严格（父类可选参数，子类改成必选 → 编译错误），否则"用父类型引用调用"时可能崩。

**抽象类**——不能实例化、只能被继承的"半成品基类"：

```typescript
abstract class Shape {
    constructor(protected name: string) {}
    abstract area(): number;      // 抽象方法：只声明，子类必须实现
    describe(): void {            // 具象方法：子类直接复用
        console.log(`${this.name} 的面积是 ${this.area().toFixed(2)}`);
    }
}
// const s = new Shape("x");      // 错误：抽象类不能 new

class Rect extends Shape {
    constructor(private w: number, private h: number) {
        super("矩形");
    }
    area(): number { return this.w * this.h; }   // 必须实现
}
new Rect(4, 5).describe();        // 矩形的面积是 20.00

// 抽象类可以作为类型（接收任何子类实例）：
function printArea(s: Shape) { s.describe(); }
```

抽象类 vs 接口：抽象类**有实现、有构造函数、单继承**；接口**纯形状描述、可多继承、更轻**。描述"是什么形状"用接口；多个子类共享大量代码时用抽象类。

**implements 与 extends 的区别**：

```typescript
interface Pingable {
    ping(): void;
}
class Sonar implements Pingable {    // implements：检查类是否满足接口
    ping() { console.log("ping!"); } // 满足
}
// class Ball implements Pingable {  // 编译错误：没有 ping

// ⚠ implements 只是"检查"，不改变类本身的类型推导：
class NameChecker implements Checkable {
    check(s) { return s.toLowerCase() === "ok"; }   // s 仍是隐式 any！
}
// 接口的参数类型不会反向"注入"实现——这是新手最大的误解

// extends：真正继承实现（一个类只能 extends 一个父类）
class Animal { move() {} }
class Dog extends Animal { woof() {} }
```

### 4.4 类型收窄（Narrowing）⭐⭐

联合类型的值在**使用前**必须先确定具体类型，TS 通过分析代码的执行路径自动完成"收窄"。

**typeof 守卫**（原始类型）：

```typescript
function padLeft(padding: number | string, input: string): string {
    if (typeof padding === "number") {
        return " ".repeat(padding) + input;   // 此分支：number
    }
    return padding + input;                    // 此分支：string
}
// typeof 可识别："string" | "number" | "bigint" | "boolean"
//               | "symbol" | "undefined" | "object" | "function"
// 注意：typeof null === "object"（历史 bug），别用它判 null
```

**真值收窄**：

```typescript
// 这些值在条件中为假：0、NaN、""、0n、null、undefined
// 其余都为真
function printAll(strs: string | string[] | null) {
    if (strs && typeof strs === "object") {   // && 先排除 null/空串
        for (const s of strs) { console.log(s); }
    }
}
// 用 Boolean(x) 或 !!x 显式转布尔
```

⚠ 真值检查对原始类型有隐患（空字符串也会被过滤），需要精确时用显式比较。

**相等性收窄**：

```typescript
function example(x: string | number, y: string | boolean) {
    if (x === y) {
        // 两者相等 ⇒ 类型必相同 ⇒ 只能都是 string
        x.toUpperCase();   // OK：string
    }
}
function printAll2(strs: string | string[] | null) {
    if (strs !== null) {   // 精确排除 null（不会误伤空串）
        // strs: string | string[]
    }
}
// 宽松相等也支持：x == null 同时排除 null 和 undefined
```

**in 守卫**（属性存在性）：

```typescript
type Fish = { swim: () => void };
type Bird = { fly: () => void };
function move(animal: Fish | Bird) {
    if ("swim" in animal) {
        animal.swim();     // true 分支：Fish（可选属性会同时出现在两侧）
    } else {
        animal.fly();
    }
}
```

**instanceof 守卫**（类实例，查原型链）：

```typescript
function logValue(x: Date | string) {
    if (x instanceof Date) {
        console.log(x.toUTCString());   // Date 分支
    } else {
        console.log(x.toUpperCase());   // string 分支
    }
}
// 只能用于类，不能用于接口/类型别名（它们编译后不存在）
```

**类型谓词（自定义类型守卫）**：

```typescript
function isFish(pet: Fish | Bird): pet is Fish {     // 返回类型是谓词
    return (pet as Fish).swim !== undefined;
}

let pet: Fish | Bird = getSmallPet();
if (isFish(pet)) {
    pet.swim();       // if 分支收窄为 Fish
} else {
    pet.fly();        // else 分支自动收窄为 Bird
}
// 还能用于数组过滤：
const zoo: (Fish | Bird)[] = [/* ... */];
const underWater: Fish[] = zoo.filter(isFish);
```

谓词语法 `参数名 is 类型` 是 TS 识别守卫的标志；`参数名` 必须是当前函数的某个参数。

**控制流分析**：TS 会沿着所有可能路径追踪类型——return/throw 之后的代码不可达，类型自动从联合中剔除：

```typescript
function padLeft2(padding: number | string, input: string) {
    if (typeof padding === "number") {
        return " ".repeat(padding) + input;   // 此路径 return 了
    }
    // 这里 padding 只剩 string（number 已被排除）
    return padding + input;
}
```

**never 穷尽检查**——收窄的终极应用：让 switch 漏分支在编译期报错：

```typescript
interface Circle { kind: "circle"; radius: number; }
interface Square { kind: "square"; sideLength: number; }
type Shape2 = Circle | Square;

function getArea2(shape: Shape2): number {
    switch (shape.kind) {
        case "circle": return Math.PI * shape.radius ** 2;
        case "square": return shape.sideLength ** 2;
        default:
            // 如果所有分支都处理完，shape 收窄为 never
            const _exhaustiveCheck: never = shape;
            return _exhaustiveCheck;
    }
}
// 将来给 Shape2 新增 "triangle" 而忘了加 case，
// default 里的 shape 不再是 never，编译立即报错——缺陷消灭在编译期
```

### 4.5 类型断言

断言是"告诉编译器：我比你清楚"。它**只影响编译期，不产生任何运行时转换**：

```typescript
// as 语法（推荐）
const str: any = "hello";
const len: number = (str as string).length;

// 尖括号语法（.tsx 中与 JSX 冲突，统一用 as）
const len2: number = (<string>str).length;

// 场景 1：any/unknown 恢复具体类型
const response: any = { name: "Alice", age: 25 };
const user = response as { name: string; age: number };

// 场景 2：收窄联合（守卫覆盖不了时）
// 场景 3：DOM（浏览器）document.querySelector 返回 Element | null
```

⚠ **断言不是转换**：`"42" as number` 编译能过，运行时它还是字符串：

```typescript
const strNum: any = "42";
const wrong = strNum as number;
console.log(typeof wrong);   // "string"！
console.log(wrong + 1);      // "421"（字符串拼接）
const real: number = Number(strNum);   // 真正的转换用 Number()/String() 等
```

**非空断言 `!`**：断言值不为 null/undefined（跳过空值检查，慎用）：

```typescript
function printLength(str?: string) {
    console.log(str!.length);   // 若 str 真是 undefined，运行时照炸不误
}
// 能用可选链/判空就不要用 !
```

**as const**（阶段三已见）：字面量收紧 + 深度只读，是安全断言，放心用。

**双重断言**（危险，最后手段）：

```typescript
const n = 42;
// n as string;                     // 错误：number 与 string 无重叠
const s = n as unknown as string;   // 借道 unknown 绕过检查
// 运行时 s 还是 42。仅在处理第三方库类型错误时使用，并写注释说明原因
```

| 断言方式 | 语法 | 风险 |
|---|---|---|
| as 断言 | `v as T` | 低（有重叠检查） |
| 非空断言 | `v!` | 中（跳过空检查） |
| 常量断言 | `v as const` | 无（安全） |
| 双重断言 | `v as unknown as T` | 高（完全绕过） |

### 4.6 可选链与空值合并

**可选链 `?.`**：链上任何一环为 null/undefined，整个表达式短路返回 undefined（不抛错）：

```typescript
interface UserProfile { address?: { city?: string }; }

const user: UserProfile = { };
// 传统：user && user.address && user.address.city
const city = user?.address?.city;          // undefined（不报错）

// 数组元素
const users = [{ name: "Alice" }];
users?.[0]?.name;        // "Alice"
users?.[9]?.name;        // undefined

// 方法调用：方法不存在则返回 undefined
const obj2 = { greet() { return "hi"; } };
obj2.greet?.();          // "hi"
obj2.sayHello?.();       // undefined（而不是 TypeError）
```

**空值合并 `??`**：左侧是 null/undefined 时才取右侧（与 `||` 的区别：`||` 会误伤 0、""、false）：

```typescript
const count = 0;
count || 10;    // 10（0 是假值，被替换——常见 bug）
count ?? 10;    // 0（0 不是 null/undefined，保留）

const name: string | null = null;
const display = name ?? "游客";   // "游客"
```

组合使用是处理嵌套可空数据的标准姿势：`user?.address?.city ?? "未知"`。

### 阶段四练习

- [ ] 定义 `Book` 接口（title、price、isbn?、readonly createdAt），写函数 `discount(book: Book, rate?: number)` 返回折后价（rate 默认 0.9）
- [ ] 用可辨识联合建模"支付方式"：`{ method: "card"; cardNo: string } | { method: "wallet"; phone: number } | { method: "cash" }`，写 `pay()` 函数用 switch 处理并加 never 穷尽检查
- [ ] 实现 `Stack<T>` 不用泛型先写 number 版：类含 push/pop/peek，pop 空栈返回 undefined
- [ ] 写抽象类 `Storage`（抽象方法 get/set），实现 `MemoryStorage` 和 `MapStorage` 两个子类，用多态数组分别调用
- [ ] 自定义类型守卫 `isErrorResponse(v: unknown): v is { code: number; message: string }`，并用它安全处理一个 any 响应
- [ ] 用参数属性把 `User` 类压缩到构造函数一行内（public name、private password、readonly id）
- [ ] 对比实验：写 `A { private x = 1 }` 和 `B { private x = 1 }`，尝试互相赋值，观察报错并解释（提示：阶段六类型兼容性会解释为什么 private 参与比较）
- [ ] 重载练习：实现 `parse(input: string): number[]`（按逗号拆分）和 `parse(input: number): number[]`（返回 `[input]`）两种签名

### 阶段四过关自检

1. 函数类型表达式 `(a: string) => void` 中参数名可以省略吗？调用签名和它的语法区别是什么？
2. 可选参数、默认参数、剩余参数分别怎么写？可选参数为什么必须在最后？
3. 什么是多余属性检查？为什么直接传对象字面量会报错而先赋给变量就不报错？
4. 可辨识联合解决什么问题？判别属性为什么必须是字面量类型？
5. public/private/protected/readonly 各自的可访问范围？参数属性做了什么？
6. TS 的 private 和 JS 的 #private 有什么区别（软私有 vs 硬私有）？
7. implements 会改变类的类型吗？extends 和 implements 的本质区别？
8. typeof / in / instanceof / 类型谓词各适合什么场景？typeof null 返回什么？
9. 什么是控制流分析？never 穷尽检查是怎么在编译期发现漏分支的？
10. `??` 和 `||` 的区别？`v!` 非空断言的风险是什么？

---

## 阶段五：泛型与类型变换（类型体操 · 上）

> **目标**：进入 TS 类型系统的"函数式编程"世界：泛型（类型的参数化）、keyof/typeof/索引访问（类型的读取）、映射类型（类型的遍历变换）、条件类型与 infer（类型的逻辑判断与提取）、模板字面量类型（字符串的类型运算）、递归类型。这是工具类型（阶段六）和一切"类型体操"的地基。

### 5.1 泛型：类型的参数 ⭐⭐

**问题**：写一个"返回数组第一个元素"的函数——

```typescript
// any 版：丢了类型信息，返回值是 any
function first(arr: any[]) { return arr[0]; }

// 每种类型写一遍？不可维护
function firstNumber(arr: number[]) { return arr[0]; }
function firstString(arr: string[]) { return arr[0]; }

// 泛型版：T 是类型参数，调用时确定
function first<T>(arr: T[]): T | undefined {
    return arr[0];
}
const s = first(["a", "b", "c"]);   // s 的类型自动是 string
const n = first([1, 2, 3]);         // number
const u = first([]);                // undefined
```

泛型 = **类型的变量**。它在函数签名里建立"输入和输出类型的关联"：传入 `string[]`，TS 推断 `T = string`，返回值就是 `string`。这就是为什么泛型比 any 强：**any 丢弃信息，泛型传递信息**。

**两种调用方式**：通常靠推断（`first([1,2])`）；推断不了或想放宽时显式指定（`first<string | number>([1, "a"])`）。

**多类型参数与函数组合**：

```typescript
// Input 和 Output 分别从实参和回调推断
function map<Input, Output>(arr: Input[], func: (arg: Input) => Output): Output[] {
    return arr.map(func);
}
const parsed = map(["1", "2", "3"], (n) => parseInt(n));
// parsed: number[] —— Input=string，Output=number，全自动推断
```

**泛型约束 `extends`**：限制类型参数的范围，换取"约束内可用的能力"：

```typescript
// T 至少要有 length 属性
function longest<T extends { length: number }>(a: T, b: T): T {
    return a.length >= b.length ? a : b;
}
longest([1, 2], [1, 2, 3]);     // T = number[]
longest("alice", "bob");        // T = "alice" | "bob"
// longest(10, 100);            // 错误：number 没有 length

// 约束的最经典模式：键必须是对象的键（阶段五 5.3 展开原理）
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
    return obj[key];
}
getProperty({ id: 1, name: "a" }, "name");   // OK
// getProperty({ id: 1 }, "foo");            // 错误：键不存在
```

⚠ **约束的常见误用**——承诺了 T 就要返回 T，不是"满足约束的随便什么东西"：

```typescript
function minimumLength<T extends { length: number }>(obj: T, min: number): T {
    if (obj.length >= min) return obj;
    return { length: min };   // 编译错误！函数承诺返回与入参同类型 T
}                             // 调用方拿到 { length: 6 } 却当数组用 .slice() 会崩
```

**用约束关联两个参数**：

```typescript
// keyof T 模式还可以这样用：值必须匹配对应键的类型
function setField<T, K extends keyof T>(obj: T, key: K, value: T[K]) { obj[key] = value; }
const user = { id: 1, name: "a" };
setField(user, "name", "b");     // OK
setField(user, "name", 2);       // 错误：name 的值必须是 string
setField(user, "id", "x");       // 错误：id 的值必须是 number
```

**泛型默认值**：

```typescript
interface ApiResponse<T = unknown> {
    success: boolean;
    data?: T;
    error?: string;
}
const r1: ApiResponse = { success: true };               // T 默认 unknown
const r2: ApiResponse<string> = { success: true, data: "ok" };
```

**官方泛型设计三准则**（写好泛型函数的检查清单）：

1. **类型参数能不用约束就不用**：`<T>(arr: T[])` 优于 `<T extends any[]>(arr: T)`——前者返回具体元素类型，后者推断成 any。
2. **类型参数越少越好**：不关联两个值的类型参数（如 `Func extends (arg: T) => boolean`）纯属添乱。
3. **类型参数应出现两次以上**：只出现一次的类型参数没有建立任何关联——`greet<Str extends string>(s: Str)` 不如直接 `greet(s: string)`。

**泛型接口与泛型类**：

```typescript
// 泛型接口
interface Pair<T, U> {
    first: T;
    second: U;
}
const pair: Pair<string, number> = { first: "hello", second: 42 };

// 泛型类
class Box<T> {
    constructor(private value: T) {}
    getValue(): T { return this.value; }
}
const stringBox = new Box<string>("TypeScript");
stringBox.getValue();   // 类型 string（实例化时锁定 T）

// 把泛型参数提到接口级别（整个接口共享一个 T）
interface GenericIdentityFn<T> {
    (arg: T): T;
}
const identity: GenericIdentityFn<number> = (arg) => arg;
```

### 5.2 keyof、typeof 与索引访问 ⭐⭐

**keyof**：获取对象类型的**所有键组成的字面量联合**：

```typescript
interface User {
    id: number;
    name: string;
    email: string;
}
type UserKeys = keyof User;        // "id" | "name" | "email"（不是 string！）
let key: UserKeys = "name";
// key = "foo";    // 错误：只能取实际存在的键
```

**typeof**（类型位置的 typeof，与运行时 typeof 完全不同）：从**值**提取它的**类型**：

```typescript
const config = { host: "localhost", port: 3000 };
type Config = typeof config;
// { host: string; port: number } —— 配置对象自动生成类型，改一处自动同步

function getUser() { return { id: 1, name: "Alice" }; }
type User2 = ReturnType<typeof getUser>;   // 与函数返回值联动（5.5 见 ReturnType）
```

**索引访问类型** `T["key"]`：按键取属性的类型：

```typescript
type UserId = User["id"];              // number
type UserName = User["name"];          // string
type IdOrName = User["id" | "name"];   // number | string（联合索引）
type AllValues = User[keyof User];     // number | string（所有值的联合）

// 数组/元组的索引访问：
type Tuple = [string, number, boolean];
type First = Tuple[0];                 // string
type Elem = string[][(number)];        // 数组元素类型（写作 string[][number]）
```

三者组合 = "类型层面的属性读取"：`typeof obj` 拿类型 → `keyof` 拿键 → `T[K]` 拿值类型。

### 5.3 映射类型（Mapped Types）⭐⭐

**基于已有类型遍历变换生成新类型**——`[P in keyof T]` 语法：

```typescript
interface User {
    id: number;
    name: string;
    email: string;
}

// 手写 Partial 的原理：遍历每个键，加 ? 修饰符
type MyPartial<T> = {
    [P in keyof T]?: T[P];       // P 依次是 "id" | "name" | "email"
};                               // T[P] 是对应属性的类型
type PartialUser = MyPartial<User>;
// 等价于 { id?: number; name?: string; email?: string }

const u: PartialUser = { name: "Alice" };   // 只填一部分也合法
```

**修饰符的加减**：

```typescript
// +?（加可选，可省略 +）、-?（去可选）
type MyRequired<T> = { [P in keyof T]-?: T[P] };      // 全部必填
type MyReadonly<T> = { readonly [P in keyof T]: T[P] };   // 全部只读
type Mutable<T> = { -readonly [P in keyof T]: T[P] };     // 去只读
```

**键重映射 as**（TS 4.1+）与**键过滤**：

```typescript
// 重映射：get 属性变成 set 属性（Capitalize 是内置字符串工具类型，5.5 介绍）
type Getters = {
    [P in keyof User as `get${Capitalize<P>}`]: () => User[P];
};
// { getId: () => number; getName: () => string; getEmail: () => string }

// 过滤：重映射为 never 即删除该键
type OnlyStrings<T> = {
    [K in keyof T as T[K] extends string ? K : never]: T[K];
};
interface Mixed { id: number; name: string; email: string; active: boolean; }
type StringProps = OnlyStrings<Mixed>;   // { name: string; email: string }
```

### 5.4 条件类型与 infer ⭐⭐

**条件类型**——类型层面的三元表达式：`T extends U ? X : Y`（T 能赋值给 U 吗）：

```typescript
type IsString<T> = T extends string ? true : false;
type A = IsString<string>;    // true（string extends string）
type B = IsString<number>;    // false

// 实战：类型过滤
type NonNull<T> = T extends null | undefined ? never : T;
type C = NonNull<string | null | undefined>;   // string（never 自动消失，见分布）
```

**分布式条件类型**：裸类型参数遇上联合类型会**自动分配**到每个成员再合并：

```typescript
type ToArray<T> = T extends any ? T[] : never;
type R = ToArray<string | number>;
// 分布过程：ToArray<string> | ToArray<number>
//        = string[] | number[]（注意不是 (string | number)[]）

// 禁用分布：用方括号包住
type ToArrayNonDist<T> = [T] extends [any] ? T[] : never;
type R2 = ToArrayNonDist<string | number>;   // (string | number)[]
```

**infer**——在 extends 子句中声明"待推断的类型变量"，从模式中提取类型（只能在条件类型中使用）：

```typescript
// 提取函数返回值（内置 ReturnType 的原理）
type MyReturnType<T> = T extends (...args: any[]) => infer R ? R : never;
type R1 = MyReturnType<() => { id: number }>;      // { id: number }
type R2b = MyReturnType<() => string>;             // string

// 提取数组元素类型
type ElementOf<T> = T extends (infer E)[] ? E : never;
type E1 = ElementOf<string[]>;      // string
type E2 = ElementOf<number>;        // never（不是数组）

// 提取 Promise 的值类型（递归解包）
type Awaited<T> = T extends Promise<infer V> ? Awaited<V> : T;
type D1 = Awaited<Promise<string>>;         // string
type D2 = Awaited<Promise<Promise<number>>>; // number（递归拆到底）

// 提取函数第一个参数
type FirstParam<T> = T extends (first: infer P, ...rest: any[]) => any ? P : never;
type P1 = FirstParam<(name: string, age: number) => void>;   // string

// infer 是类型体操的"解构赋值"：匹配模式 + 捕获变量
```

### 5.5 模板字面量类型

字符串的**类型层面运算**——用模板语法从联合类型生成新的字面量联合：

```typescript
type World = "world";
type Greeting = `hello ${World}`;         // "hello world"

// 联合交叉相乘
type EmailLocale = "welcome_email" | "email_heading";
type FooterLocale = "footer_title" | "footer_sendoff";
type AllLocaleIDs = `${EmailLocale | FooterLocale}_id`;
// "welcome_email_id" | "email_heading_id" | "footer_title_id" | "footer_sendoff_id"

// 四个内置字符串操作类型（编译器内置，性能好）：
type Upper = Uppercase<"hello">;        // "HELLO"
type Lower = Lowercase<"HELLO">;        // "hello"
type Cap = Capitalize<"hello">;         // "Hello"
type Uncap = Uncapitalize<"Hello">;     // "hello"
```

**经典应用：类型安全的事件系统**——事件名 = 属性名 + "Changed"，回调参数类型 = 对应属性类型：

```typescript
type PropEventSource<Type> = {
    on<Key extends string & keyof Type>(
        eventName: `${Key}Changed`,
        callback: (newValue: Type[Key]) => void
    ): void;
};
// 泛型 on + 模板字面量 + infer 式匹配：
// 调 on("firstNameChanged", cb) 时，TS 从模板反推出 Key = "firstName"，
// 再用 Type[Key] 索引出 string 作为回调参数类型
// declare 声明“这个函数在别处已实现，这里只描述签名”（7.4 详解声明文件）
declare function makeWatched<T>(obj: T): T & PropEventSource<T>;
const person = makeWatched({ firstName: "Saoirse", lastName: "Ronan", age: 26 });

person.on("firstNameChanged", newName => {
    // newName 自动推断为 string
    console.log(newName.toUpperCase());
});
person.on("ageChanged", newAge => { /* newAge: number */ });
// person.on("firstName", () => {});      // 错误：少个 Changed
// person.on("frstNameChanged", () => {});// 错误：拼错属性名
```

### 5.6 递归类型

类型可以引用自身——表达任意深度的结构：

```typescript
// 树形结构
interface TreeNode {
    value: number;
    children?: TreeNode[];          // 递归引用
}

// 嵌套列表：number 或 "包含嵌套列表的数组"
type NestedList = number | NestedList[];

// 深度只读：递归给每一层加 readonly
type DeepReadonly<T> = {
    readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};
interface Nested { a: { b: { c: number } }; }
const obj3: DeepReadonly<Nested> = { a: { b: { c: 1 } } };
// obj3.a.b.c = 2;    // 错误：三层全是只读

// 提取 Promise 深处的值（5.4 的 Awaited 就是递归条件类型）
// 链式结构
interface ListNode { value: string; next: ListNode | null; }
```

### 阶段五练习

- [ ] 实现泛型函数 `groupBy<T>(arr: T[], key: keyof T)`，返回类型用索引签名 `{ [k: string]: T[] }`（提示：值可能是任意类型，可用 String() 归一化键）
- [ ] 手写 `MyPick<T, K extends keyof T>`（遍历 K 而不是 keyof T）并测试
- [ ] 写 `UnwrapPromise<T>`：`Promise<Promise<string>>` → string（递归）
- [ ] 用模板字面量类型生成 `EventName<T>`：`{ click: ...; focus: ... }` → `"onClick" | "onFocus"`
- [ ] 写类型 `KeysOfType<T, U>`：返回 T 中值类型为 U 的键（提示：映射 + never 过滤 + 索引取联合）
- [ ] 实现泛型类 `Result<T, E = Error>`：含 `{ ok: true; value: T } | { ok: false; error: E }` 的工厂函数 `ok()/err()`
- [ ] 用 `typeof + keyof + 索引访问` 三连：从真实常量 `const ENDPOINTS = { user: "/u", order: "/o" }` 提取 `"/u" | "/o"` 类型

### 阶段五过关自检

1. 泛型和 any 的本质区别是什么？"类型参数应出现两次"准则怎么理解？
2. `T extends { length: number }` 约束带来了什么、限制了什么？minimumLength 的错误说明了什么原则？
3. keyof User 的结果是什么类型？`typeof config`（类型位置）做什么？
4. `T["id" | "name"]` 和 `T[keyof T]` 分别得到什么？
5. `[P in keyof T]?: T[P]` 每一部分是什么意思？`-?` 和 `-readonly` 呢？
6. 什么是分布式条件类型？怎么禁用？never 在分布中为什么会"消失"？
7. infer 只能出现在哪里？写一个提取数组元素类型的条件类型。
8. `Capitalize<"hello">` 是什么？模板字面量类型的联合交叉相乘是什么规则？

---
## 阶段六：类型系统深水区（类型体操 · 下）

> **目标**：掌握内置工具类型全家桶（含手写原理）、类型兼容性（结构化类型的规则与例外）、协变逆变、声明合并，完成一组类型体操实战。学完本阶段，你能读懂主流开源库的类型定义，也能为自己的库设计类型。

### 6.1 内置工具类型全家桶 ⭐⭐

工具类型都是泛型类型，原理就是阶段五的映射/条件类型。**逐个手写一遍是最好的练习**。

**对象变换组**：

```typescript
interface Todo {
    title: string;
    description: string;
    completed: boolean;
    createdAt: Date;
}

// Partial<T>：全部可选 —— 更新接口的部分字段
type P = Partial<Todo>;              // { title?: string; ... }
// 原理：{ [K in keyof T]?: T[K] }

// Required<T>：全部必填 —— 与 Partial 相反
type R = Required<Partial<Todo>>;    // 回到全部必填
// 原理：{ [K in keyof T]-?: T[K] }

// Readonly<T>：全部只读 —— 配置、常量
type RO = Readonly<Todo>;
// 原理：{ readonly [K in keyof T]: T[K] }

// Pick<T, K>：挑选部分键
type TodoPreview = Pick<Todo, "title" | "completed">;
// { title: string; completed: boolean }
// 原理：{ [K in K]: T[K] }

// Omit<T, K>：排除部分键（Pick 的反向）
type TodoInfo = Omit<Todo, "description" | "createdAt">;
// { title: string; completed: boolean }
// 原理：Pick<T, Exclude<keyof T, K>>

// Record<K, T>：构造键值字典 —— 键是 K（字面量联合），值是 T
type Role = "admin" | "user" | "guest";
type RolePermissions = Record<Role, string[]>;
const permissions: RolePermissions = {
    admin: ["read", "write", "delete"],   // 必须三个键都给全！
    user: ["read", "write"],
    guest: ["read"],
};
```

**联合类型操作组**：

```typescript
// Exclude<T, U>：从联合 T 中排除可赋值给 U 的成员
type T0 = Exclude<"a" | "b" | "c", "a">;           // "b" | "c"
// 原理：T extends U ? never : T（分布 + never 消失）

// Extract<T, U>：从联合 T 中提取可赋值给 U 的成员
type T1 = Extract<"a" | "b" | "c", "a" | "f">;     // "a"
// 原理：T extends U ? T : never

// NonNullable<T>：剔除 null 和 undefined
type T2 = NonNullable<string | null | undefined>;  // string
// 原理：Exclude<T, null | undefined>
```

**函数推断组**（基于 infer）：

```typescript
function getUser() { return { id: 1, name: "Alice" }; }

// ReturnType<T>：函数返回值类型
type User = ReturnType<typeof getUser>;        // { id: number; name: string }
// 原理：T extends (...args: any[]) => infer R ? R : never

// Parameters<T>：函数参数列表（元组）
type Args = Parameters<(name: string, age: number) => void>;
// [name: string, age: number]
// 原理：T extends (...args: infer P) => any ? P : never

// ConstructorParameters<T>：构造函数参数
type ErrArgs = ConstructorParameters<ErrorConstructor>;   // [message?: string]

// InstanceType<T>：构造函数的实例类型
type DateInst = InstanceType<typeof Date>;     // Date

// Awaited<T>：递归解包 Promise（阶段五手写过原理）
type V = Awaited<Promise<string>>;             // string
type V2 = Awaited<boolean | Promise<number>>;  // boolean | number
```

**速查总表**：

| 工具类型 | 作用 | 一句话原理 |
|---|---|---|
| Partial / Required | 全可选 / 全必填 | 映射 ±? |
| Readonly | 全只读 | 映射 +readonly |
| Pick / Omit | 挑键 / 排键 | 映射 K / Pick+Exclude |
| Record | 键值字典 | 映射 K → T |
| Exclude / Extract | 联合排除 / 提取 | 条件分布 |
| NonNullable | 去空值 | Exclude |
| ReturnType / Parameters | 返回值 / 参数 | infer |
| ConstructorParameters / InstanceType | 构造参数 / 实例 | infer |
| Awaited | 解 Promise | 递归条件 |
| Uppercase 等四件套 | 字符串大小写 | 编译器内置 |

**后端高频组合拳**：

```typescript
// 更新接口：只允许传部分字段，且不允许改 id
type UpdateUserDto = Partial<Omit<User, "id">>;

// 列表查询参数：可选过滤 + 分页
type QueryParams = Partial<Record<"page" | "pageSize", number>> &
                   Partial<Record<"keyword" | "status", string>>;

// API 响应泛型包装（阶段八实战的主类型）
interface ApiResponse<T> {
    success: boolean;
    data?: T;
    error?: string;
}
type UserResponse = ApiResponse<User[]>;
```

### 6.2 类型兼容性：结构化类型的规则 ⭐

TS 采用**结构化类型**：只比较成员形状，不管声明时的名字：

```typescript
interface Pet { name: string; }
class Dog { name: string; }

let pet: Pet;
pet = new Dog();       // 合法！Dog 有 name: string 就行，不需要声明"实现 Pet"

// "x 可赋值给 y"当且仅当 y 需要的成员 x 都有（且类型兼容）
const dog = { name: "Lassie", owner: "Rudd" };   // 多余属性没关系（变量传递）
pet = dog;            // 合法（但对象字面量直接赋值会触发多余属性检查）
```

**函数兼容的方向性**——参数"少的能赋给多的"，返回值"多的能赋给少的"：

```typescript
// 参数：目标函数需要的参数，源函数可以更少（忽略多余参数是 JS 惯例）
let x = (a: number) => 0;
let y = (b: number, s: string) => 0;
y = x;    // OK：x 用更少参数，安全
// x = y; // 错误：y 需要两个参数，x 只提供一个

// 返回值：源函数返回的必须至少覆盖目标需要的
let getX = () => ({ name: "Alice" });
let getY = () => ({ name: "Alice", location: "Seattle" });
getX = getY;    // OK：多返回不碍事
// getY = getX; // 错误：缺 location

// 这解释了内置 API 的宽容：forEach 回调可以只写 (item) 不写 (item, index, arr)
[1, 2, 3].forEach((item) => console.log(item));   // 参数少 → 合法
```

**函数参数双向兼容**：方法参数默认采用双变检查（源或目标任一方向可赋值即可）——这是刻意的"不健全"设计，为了兼容 JS 里常见的事件处理器模式；开启 `strictFunctionTypes` 后独立函数类型改为严格逆变。

**其他兼容规则**：

```typescript
// 枚举与数字互通（数字枚举成员兼容 number；不同枚举之间不兼容）
enum Status { Active, Inactive }
let s: number = Status.Active;    // OK

// 类：只比较实例成员；private/protected 成员必须"来自同一个声明"
class A { private x = 0; sameAs(other: A) { return other.x === this.x; } }  // 跨实例私有访问 OK
class B { private x = 0; }
// let a: A = new B();   // 错误：两个类都有 private x，但声明不同源 → 不兼容
// 这就是阶段四练习里那个实验的答案

// 泛型：类型参数只在使用到时才参与比较
interface Empty<T> { }
let e1: Empty<number>; let e2: Empty<string>;
// e1 = e2;  // OK：T 没被用到，结构一样
```

### 6.3 协变与逆变 ⚠

"方向性"的系统描述（了解概念，读库源码时会遇到）：

```typescript
class Animal { name = "animal"; }
class Dog extends Animal { breed = "lab"; }

// 协变：方向一致。返回 Dog 的函数可以当作"返回 Animal 的函数"用
type AnimalGetter = () => Animal;
type DogGetter = () => Dog;
const getDog: DogGetter = () => new Dog();
const getAnimal: AnimalGetter = getDog;   // ✅ 协变安全：调用方要 Animal，拿到 Dog（是 Animal）没问题

// 逆变：方向相反。接收 Animal 的函数可以当作"接收 Dog 的函数"用
type AnimalHandler = (a: Animal) => void;
type DogHandler = (d: Dog) => void;
const handleAnimal: AnimalHandler = (a) => { console.log(a.name); };
const handleDog: DogHandler = handleAnimal;   // ✅ 逆变安全：传 Dog 进来，Dog 是 Animal，处理没问题
// strictFunctionTypes 开启后，独立函数类型的参数按逆变严格检查
```

记忆法：**返回值协变（只多不少），参数逆变（只宽不窄）**。

### 6.4 声明合并（Declaration Merging）

TS 允许同名声明自动合并——接口合并最常用：

```typescript
interface Box { height: number; }
interface Box { width: number; }
// 合并为 { height: number; width: number }
// 非成员值冲突（同名的同种类成员类型不同）会报错

// 典型场景：给第三方类型"打补丁"
interface Window {
    __APP_CONFIG__?: { apiBase: string };   // 给全局 Window 补一个自定义属性
}
```

命名空间可与类、函数、枚举合并（给它们附加静态成员），模块可做 module augmentation 扩展（阶段七模块部分再提）。**合并是“隐式全局行为”，团队项目中慎用**——补丁应集中放一个文件并注释说明。

### 6.5 Mixin 模式

类只支持单继承，但有时需要给一个类“叠加多个能力”。Mixin 模式：用函数接收一个基类、返回增强后的新类，逐层包裹组合能力：

```typescript
// “构造函数”类型的标准写法（就是 4.1 构造签名的别名形态）
type Constructor<T = {}> = new (...args: any[]) => T;

// 能力 1：时间戳
function Timestamped<TBase extends Constructor>(Base: TBase) {
    return class extends Base {
        timestamp = Date.now();
    };
}
// 能力 2：序列化
function Serializable<TBase extends Constructor>(Base: TBase) {
    return class extends Base {
        serialize(): string {
            return JSON.stringify(this);
        }
    };
}

class User {
    constructor(public name: string) {}
}

// 组合：先叠时间戳，再叠序列化
const TimestampedUser = Timestamped(User);
const FullUser = Serializable(TimestampedUser);

const user = new FullUser("Alice");
user.timestamp;      // 时间戳能力
user.serialize();    // 序列化能力（序列化结果同时含 name 和 timestamp）
```

记忆点：每个 Mixin 只加一个能力、保持独立；组合时层层包裹。与装饰器（7.6）的分工：**给类附加行为**这件事上两者是替代关系——框架生态用装饰器，工具函数层的类组合用 Mixin。

### 6.6 类型体操实战集 ⭐

综合运用阶段五 + 阶段六的全部武器（每题先自己写，再对照）：

```typescript
// 1. DeepPartial：递归可选
type DeepPartial<T> = {
    [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P];
};

// 2. TupleToUnion：元组 → 联合 [string, number] → string | number
type TupleToUnion<T extends any[]> = T[number];

// 3. GetReturnType 的 Promisified 版：函数 Promise 化后的返回值
type Promisified<T> = T extends (...args: infer A) => infer R
    ? (...args: A) => Promise<R>
    : never;

// 4. 深层键路径："a.b.c" 形式（递归模板字面量）
type PathKeys<T> = {
    [K in keyof T & string]: T[K] extends object
        ? K | `${K}.${PathKeys<T[K]>}`
        : K;
}[keyof T & string];
type Paths = PathKeys<{ a: { b: { c: 1 } }; d: 2 }>;
// "a" | "a.b" | "a.b.c" | "d"

// 5. Mutable：递归去只读（对照 DeepReadonly）
type Mutable<T> = {
    -readonly [P in keyof T]: T[P] extends object ? Mutable<T[P]> : T[P];
};

// 6. 函数参数倒序（元组 + 递归 + infer 三合一）
type Reverse<T extends any[]> = T extends [infer F, ...infer R]
    ? [...Reverse<R>, F]
    : [];
type Rev = Reverse<[1, 2, 3]>;    // [3, 2, 1]
```

**度的问题**：类型体操的收益边际递减。业务代码 90% 的场景用接口 + 联合 + 内置工具类型足够；高级变换用在**通用库/工具函数**上。团队代码里出现读不懂的条件类型 = 维护负担。性能上也有代价：巨型条件/递归类型会拖慢编译——简单对象类型用 interface（编译更快），复杂变换才用 type。

### 阶段六练习

- [ ] 不看答案手写：Partial、Required、Readonly、Pick、Omit、Record、Exclude、Extract、NonNullable、ReturnType、Parameters
- [ ] 用工具类型组合定义：`CreateUserDto`（Omit 掉 id 和时间戳）、`UserPatch`（Partial + 不含 id）、`UserSummary`（Pick 前三个字段）
- [ ] 写 `Swap<P>`：把 `{ a: 1; b: 2 }` 的键值反转成 `{ 1: "a"; 2: "b" }`（提示：映射 + as 反向索引）
- [ ] 验证结构化兼容：定义两个"形状相同但无继承关系"的类型互相赋值；再验证"private 同名不同源"不兼容
- [ ] 解释下面代码为什么合法：`const f: (a: number) => void = (a, b) => {}` 反过来为什么不行？
- [ ] 写 `Flatten<T>`：`Flatten<[1, [2, [3]]]>` → `[1, 2, 3]`（递归 + 展开元组）

### 阶段六过关自检

1. Record<Role, string[]> 定义的对象必须包含什么？漏一个键会怎样？
2. Exclude 和 Extract 的原理是什么（用分布条件类型解释）？
3. ReturnType 的完整手写？infer R 在哪里、为什么不能写在外面？
4. 结构化类型系统下，"类没有 implements 某接口"却可以赋值给接口类型，为什么？
5. 函数赋值时参数个数和返回值的兼容方向分别是什么？为什么 forEach 回调可以只写一个参数？
6. 两个都有 `private x` 的类为什么互不兼容？
7. 协变和逆变的记忆口诀？strictFunctionTypes 影响哪一类检查？
8. 声明合并的典型使用场景和风险？
9. Mixin 解决什么问题？`Constructor` 类型（`new (...args: any[]) => T`）描述的是什么？

---

## 阶段七：工程化与语言机制

> **目标**：从"会写类型"到"会做工程"：模块系统（ESM/CJS 及互操作）、import type、模块解析与路径别名、命名空间、声明文件、tsconfig 全解、装饰器（新旧两版）、迭代器与生成器、错误处理模式、性能与迁移。这是 Node.js 工程师面试与实战的高频区。

### 7.1 模块系统 ⭐⭐

**什么是模块**：任何包含顶层 `import` 或 `export` 的文件就是模块；模块内变量私有，必须显式导出才能被外界使用。没有导入导出的文件是"脚本"，其内容挂在全局作用域（想让它变模块又无东西可导，写一行 `export {}`）。

**ES Module 导出导入全形态**：

```typescript
// ===== user.ts =====
export const name = "Alice";                    // 命名导出：变量
export function greet(m: string) { return m; }  // 命名导出：函数
export class User { constructor(public n: string) {} }
export interface Config { host: string }        // 类型也能导出

export default class UserService {}             // 默认导出：每模块最多一个

// ===== main.ts =====
import UserService from "./user";               // 默认导入：不要花括号，名字随便取
import { name, greet } from "./user";           // 命名导入：花括号，名字必须一致
import { greet as sayHello } from "./user";     // 重命名导入（解决冲突）
import * as UserModule from "./user";           // 全部导入塞进命名空间对象
import "./user";                                // 副作用导入：只执行模块，不取值
import UserService2, { name as n2 } from "./user"; // 混合：默认在前，命名在后

// ===== index.ts 聚合（barrel 模式）=====
export { name, greet } from "./user";           // 重新导出
export { default as UserService } from "./user";// 默认导出转命名导出
export * from "./math";                         // 全量转发

// ===== 动态导入（返回 Promise，运行时按需加载）=====
async function load() {
    const math = await import("./math");
    // math.default / math.multiply ...
}
```

**import type**：显式声明"只导入类型"，编译后整体移除（帮助打包工具做安全的 tree-shaking）：

```typescript
import type { User, Config } from "./types";        // 整条语句全是类型
import { createUser, type User } from "./user";     // 混合：值正常导入，类型标 type
```

**CommonJS 语法**（Node.js 传统模块格式，读老代码必备）：

```javascript
// 导出：挂到 module.exports（或 exports）
function absolute(num) { return num < 0 ? -num : num; }
module.exports = { absolute };
// exports.absolute = absolute;   // 等价写法

// 导入
const { absolute } = require("./maths");
```

**CJS 与 ESM 互操作**：`esModuleInterop: true`（推荐常开）允许用 import 语法导入 CJS 模块并支持默认导入；`allowSyntheticDefaultImports` 让无默认导出的 CJS 包也能 `import fs from "fs"`。TypeScript 编译时会把 import 转换成目标模块格式（由 `module` 选项决定）。

### 7.2 模块解析与路径别名

"import 路径"如何映射到磁盘文件，由 `moduleResolution` 决定：

- `node`（Node.js 传统解析：`node_modules` 逐级向上找、支持目录/index、自动补 .ts/.tsx/.d.ts 扩展名）——CommonJS 项目选它
- `node16` / `nodenext`：跟随 Node.js 真实规则（ESM 下相对导入必须写扩展名）
- `bundler`：给 Vite/esbuild 等打包器用

**路径别名**（tsconfig）：

```json
{
    "compilerOptions": {
        "baseUrl": "./src",
        "paths": {
            "@/*": ["./*"],
            "@components/*": ["./components/*"]
        }
    }
}
```

```typescript
import { UserService } from "@/services/userService";   // 替代 ../../services/...
```

⚠ TS 只管"类型层能找到"，运行时能不能找到取决于打包器/运行时是否配置了同样的别名（Node 直接跑需要 tsc-alias 或运行时 loader）。路径别名是**纯编译期映射**。

### 7.3 命名空间与三斜线指令

命名空间是模块普及前的代码组织方式（把全局变量装进命名过的"袋子"）：

```typescript
namespace Drawing {
    export interface IShape { draw(): void; }      // 要 export 才可见
    export class Circle implements IShape {
        draw() { console.log("Circle is drawn"); }
    }
}
new Drawing.Circle().draw();

// 嵌套命名空间
namespace Runoob {
    export namespace invoiceApp {
        export class Invoice {
            calculateDiscount(price: number) { return price * 0.4; }
        }
    }
}
new Runoob.invoiceApp.Invoice().calculateDiscount(500);

// 跨文件命名空间：三斜线指令引用（老式做法）
/// <reference path="IShape.ts" />
/// <reference path="Circle.ts" />
```

**现代结论**：新代码一律用模块，命名空间只在维护老代码、或配合声明合并时出现。两者不要混用（给模块内部再加命名空间纯属多余）。

### 7.4 声明文件（.d.ts）与 declare ⭐⭐

**问题场景**：在 TS 项目里引入一个纯 JS 库：

```typescript
jQuery("#foo");
// error TS2304: Cannot find name 'jQuery'.
```

TS 不认识 JS 库里的全局名。**声明文件**就是"给 JS 库补的类型说明书"：

```typescript
// types/jquery.d.ts —— 声明全局变量与函数
declare var jQuery: (selector: string) => any;
declare function myFunction(param: string): void;
declare namespace MyNamespace {
    function doSomething(): void;
}

// 给无类型的模块"占坑"（快速消错，放弃检查）：
declare module "some-untyped-module";
```

**获取类型定义的优先顺序**（官方指南）：

1. 库自带类型（现代 npm 包在 package.json 标 `"types"` 字段，装了就有）
2. **DefinitelyTyped 社区仓库**：安装 `@types/包名`（如 `npm i -D @types/node`），TS 自动读取 `node_modules/@types`
3. **自己写 .d.ts**（以 jQuery 插件为例，从零创建一个完整声明）：

```typescript
// 第三方 JS 库 CalcThirdPartyJsLib.js 导出了全局对象 Calc：
// Calc.add(a, b)、Calc.multiply(a, b)、Calc.version

// Calc.d.ts —— 为它编写声明文件
declare namespace Calc {
    function add(a: number, b: number): number;
    function multiply(a: number, b: number): number;
    const version: string;
}
```

`.d.ts` 只含类型不产代码；`.ts` 是实现文件。**@types/node、@types/express、@types/jest 是 Node 开发三大件**。

### 7.5 tsconfig.json 全解 ⭐

`npx tsc --init` 生成模板。按职能分域记忆：

**输出控制**：

```jsonc
{
  "compilerOptions": {
    "outDir": "./dist",        // 编译产物目录
    "rootDir": "./src",        // 源码根目录
    "declaration": true,       // 生成 .d.ts（写库必开）
    "sourceMap": true          // 生成 .map（调试时定位回 TS 源码）
  }
}
```

**类型检查（strict 家族）**：

```jsonc
{
  "strict": true,                      // 总开关：下面全部为 true（生产必开）
  // strict 等价于同时开启：
  // "noImplicitAny": true,             // 隐式 any 报错
  // "strictNullChecks": true,          // null/undefined 不再是万物子类型
  // "strictFunctionTypes": true,       // 函数参数逆变检查
  // "strictBindCallApply": true,
  // "strictPropertyInitialization": true,  // 类字段必须初始化
  // "noImplicitThis": true,
  // "alwaysStrict": true
  // 额外推荐：
  // "noUnusedLocals": true,            // 未使用的局部变量报错
  // "noUnusedParameters": true,        // 未使用的参数报错
  // "noImplicitReturns": true          // 漏返回路径报错
  }
}
```

**模块系统**（见 7.1/7.2）：`module`、`moduleResolution`、`esModuleInterop`、`allowSyntheticDefaultImports`、`baseUrl`、`paths`、`isolatedModules`。

**目标版本**：`target`（编译产物的 JS 版本，如 ES2020；影响语法降级——class/箭头函数到 ES5 会变样）、`lib`（可用的内置 API 类型库，Node 项目用 `["ES2020"]`，不需要 DOM）。

**实验性**：`experimentalDecorators`、`emitDecoratorMetadata`（见 7.6 装饰器）。

**枚举 vs 字面量联合的工程选择（兑现阶段三的伏笔）**：需要**运行时对象**时用枚举——反向映射 `Direction[1]`、遍历成员、把 `Color.Red` 当值传递，代价是编译后产生真实代码；只做**类型约束**时优先字面量联合 `type Status = "a" | "b"`——零运行时代码、tree-shaking 友好、与 JSON 数据天然对齐。const 枚举虽会被编译器内联成值（体积最小），但在 isolatedModules（逐文件编译，部分打包器要求开启）下不可使用，新项目慎用。

**Node.js 后端推荐配置**（背下来）：

```jsonc
{
  "compilerOptions": {
    "target": "ES2020",
    "module": "commonjs",
    "lib": ["ES2020"],
    "outDir": "./dist",
    "rootDir": "./src",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,                    // 跳过 .d.ts 检查（加快编译）
    "forceConsistentCasingInFileNames": true,// 文件名大小写一致性（跨平台协作）
    "moduleResolution": "node",
    "declaration": true
  },
  "include": ["src/**/*"],
  "exclude": ["node_modules", "dist"]
}
```

**项目引用与 monorepo（工程进阶）**：

```jsonc
// packages/utils/tsconfig.json —— 子包声明为可组合项目
{
  "extends": "../../tsconfig.base.json",    // 继承基础配置
  "compilerOptions": { "composite": true },
  "references": []                           // 声明依赖的其他项目
}
// 主项目 references 到子包后，可用增量构建：
// npx tsc -b                 // 只重建有变化的项目，大型 monorepo 提速明显
// pnpm workspace + workspace:* 协议管理包间依赖，是当前主流方案
```

### 7.6 装饰器（Decorators）⭐

**版本分水岭（必须先搞清）**：

- **旧版（实验性）**：`experimentalDecorators: true` 启用。**NestJS、TypeORM 等主流后端框架目前用的就是它**，面试与工作都以它为主。
- **新版（标准）**：TypeScript 5.0 起内置 Stage 3 标准装饰器，语法与能力有差异，无需实验开关。生态迁移仍在进行。

以下按**旧版**讲解（后端生态现状）。装饰器 = 一种特殊的函数，用 `@` 附加在类/方法/属性/参数上，在不修改原代码的情况下附加行为（日志、鉴权、路由注册……框架用它实现"约定优于配置"）。

**类装饰器**：接收构造函数：

```typescript
function sealed(target: Function) {
    console.log("装饰器应用于:", target.name);
    Object.seal(target);
    Object.seal(target.prototype);
}

@sealed
class Person {
    constructor(public name: string) {}
}
// 类定义时立即执行（不是实例化时）
```

**装饰器工厂**：接收配置、返回装饰器（最常用形态）：

```typescript
function color(code: string) {              // 工厂：参数化
    return function (target: any, key: string, descriptor: PropertyDescriptor) {
        const original = descriptor.value;
        descriptor.value = function (...args: any[]) {
            const result = original.apply(this, args);
            return `\x1b[${code}m${result}\x1b[0m`;   // ANSI 终端颜色
        };
    };
}
class Logger {
    @color("31") error(msg: string) { return msg; }   // 红色输出
}
```

**方法装饰器**：三个参数（原型、方法名、属性描述符），可改 enumerable/wrap 方法：

```typescript
function enumerable(value: boolean) {
    return function (target: any, propertyKey: string, descriptor: PropertyDescriptor) {
        descriptor.enumerable = value;
    };
}
class Greeter {
    @enumerable(false)
    greet() { return "Hello"; }
}
```

**访问器装饰器**（getter/setter，只能装饰其中一个）、**属性装饰器**（两个参数：原型、属性名，常存元数据）、**参数装饰器**（三个参数：原型、方法名、参数下标）——签名见表：

| 装饰器 | 参数 | 典型用途 |
|---|---|---|
| 类 | (构造函数) | 注册路由/单例/元数据 |
| 方法 | (原型, 方法名, 描述符) | 日志/鉴权/缓存/路由 |
| 访问器 | (原型, 名字, 描述符) | 锁定 configurable |
| 属性 | (原型, 属性名) | 校验规则元数据 |
| 参数 | (原型, 方法名, 参数索引) | 依赖注入标记 |

**执行顺序**：多个装饰器自下而上应用、工厂从上到下求值；参数装饰器先于方法装饰器。叠加 `@A @B` 时像函数复合 `A(B(target))`。

配合 `emitDecoratorMetadata: true` 可为装饰器生成类型元数据——NestJS 依赖注入的基础。

### 7.7 迭代器与生成器

**可迭代协议**：实现 `[Symbol.iterator]()` 方法的对象可用 `for...of`（数组、字符串、Map、Set 内置支持）：

```typescript
for (const x of [1, 2, 3]) { }          // 值
for (const k in { a: 1 }) { }           // 键（for...in 是另一套：枚举属性名）

// 自定义可迭代对象
class Range implements Iterable<number> {
    constructor(private start: number, private end: number) {}
    [Symbol.iterator]() {
        let current = this.start;
        const end = this.end;
        return {
            next(): IteratorResult<number> {
                return current <= end
                    ? { value: current++, done: false }
                    : { value: undefined, done: true };
            },
        };
    }
}
for (const n of new Range(1, 5)) console.log(n);   // 1 2 3 4 5
```

**生成器**：`function*` + `yield`，惰性产出序列：

```typescript
function* counter(): Generator<number, void, void> {
    // Generator<yield 的类型, return 的类型, next() 参数的类型>
    yield 1;
    yield 2;
    yield 3;
}
const it = counter();
it.next();   // { value: 1, done: false }
it.next();   // { value: 2, done: false }
it.next();   // { value: 3, done: false }
it.next();   // { value: undefined, done: true }

// 无限生成器（惰性 = 不会撑爆内存）+ 提前终止
function* naturalNumbers() {
    let i = 1;
    while (true) yield i++;
}
function* take<T>(n: number, gen: Generator<T>): Generator<T> {
    for (const v of gen) {
        if (n-- <= 0) return;
        yield v;
    }
}
for (const v of take(5, naturalNumbers())) console.log(v);  // 1~5

// yield* 委托给另一个生成器
function* combined() { yield* counter(); yield* counter(); }
```

Node 后端场景：流式处理大数据、分页拉取、惰性管道。

### 7.8 错误处理的工程模式 ⭐

**自定义错误类**（区分错误类型）：

```typescript
class AppError extends Error {
    constructor(message: string, public code: string = "APP_ERROR") {
        super(message);
        this.name = "AppError";
        Object.setPrototypeOf(this, AppError.prototype);
        // ↑ 编译到 ES5 时 extends Error 会断原型链，必须手动接回
        // （target ES2015+ 则不需要）
    }
}
class NotFoundError extends AppError {
    constructor(public resource: string) {
        super(`${resource} 不存在`, "NOT_FOUND");
    }
}
throw new NotFoundError("用户");
```

**catch 到的是 unknown**（strict 模式）：必须收窄后才能用：

```typescript
try {
    risky();
} catch (err) {                     // err: unknown（不是 any）
    if (err instanceof Error) {
        console.log(err.message);   // 收窄后可用
    } else {
        console.log(String(err));   // throw 任何值的兜底
    }
}
```

**Result 模式**——用类型系统强迫调用方处理失败（Go 风格）：

```typescript
type Result<T, E = string> =
    | { ok: true; value: T }
    | { ok: false; error: E };

function parseInt2(input: string): Result<number> {
    const n = Number.parseInt(input, 10);
    return Number.isNaN(n)
        ? { ok: false, error: `"${input}" 不是数字` }
        : { ok: true, value: n };
}
const r = parseInt2("42");
if (r.ok) {
    console.log(r.value + 1);   // true 分支才能拿到 value
} else {
    console.log(r.error);       // false 分支才能拿到 error
}
```

**async 错误处理**：async 函数内用 try/catch 包裹 await（阶段二 2.8 已练）。未捕获的 reject 会让 Promise 无人处理——错误最终如何被上层框架收集，取决于框架的错误传递机制；阶段八的 Express 用“同步 throw + 错误中间件”落地同一套 AppError 体系。

### 7.9 从 JS 迁移与性能

**渐进迁移四步法**：

1. `allowJs: true` 让 TS 编译器先"看"JS 文件（不检查）
2. JS 文件顶部加 `// @ts-check` 开启轻量检查；`// @ts-ignore` 压制单行错误
3. 用 JSDoc 注释补类型（不改文件后缀就能获得类型提示）：`/** @param {number} a */`
4. 逐文件改后缀 .ts，分阶段开严格检查（先 noImplicitAny，再 strictNullChecks，最后 strict）

**性能清单**：能推断就不标注（减少重复检查）；any 是性能与安全的双输；接口 vs 类型别名——简单对象用 interface 编译更快；巨型联合/深递归类型是编译耗时大户；`skipLibCheck` 跳过第三方 .d.ts 检查；monorepo 用项目引用增量构建；具名导出比默认导出更利于打包 tree-shaking。

### 阶段七练习

- [ ] 建一个双文件项目：`math.ts`（含命名导出 multiply 和默认导出 add）+ `main.ts`（五种导入方式各用一遍：默认/命名/重命名/命名空间/动态导入）
- [ ] 给一个假想的 `legacy-library`（无类型包）写完整 .d.ts（declare module 形式，含一个函数、一个常量、一个接口）
- [ ] 抄写并逐项注释"Node 后端推荐 tsconfig"，说明每项作用
- [ ] 写方法装饰器 `@log`：调用时打印 `方法名(参数)` 与返回值、耗时
- [ ] 写生成器 `fib()` 产斐波那契数列，配合 `take()` 输出前 10 项
- [ ] 把阶段四的支付函数改造成 Result 模式：`pay(amount): Result<{ orderId: string }, PaymentError>`
- [ ] 迁移练习：写一个含明显类型 bug 的 .js 文件，加 `// @ts-check` + JSDoc，观察编辑器提示

### 阶段七过关自检

1. 什么文件是模块、什么是脚本？`export {}` 有什么用？
2. import type 与普通 import 的区别？为什么打包工具喜欢它？
3. esModuleInterop 解决什么问题？CJS 的 module.exports 和 ESM 的 export 是什么关系？
4. 获取第三方库类型的三个途径（按优先级）？
5. strict 开关包含哪些子检查？strictNullChecks 和 noImplicitAny 各防止什么？
6. target 和 lib 分别控制什么？Node 项目 lib 要包含 DOM 吗？
7. 旧版装饰器怎么开启？五类装饰器的参数签名分别是什么？装饰器工厂解决什么问题？
8. for...of 和 for...in 的区别？Generator<T, R, N> 三个类型参数各是什么？
9. catch 的 err 在 strict 模式下是什么类型？怎么安全使用？
10. Result 模式相比 throw 的优势是什么？

---

## 阶段八：Node.js 后端实战

> **目标**：把前面所有知识组装成一个真实的 Node.js + Express + TypeScript 项目：工程初始化、类型分层设计（实体/DTO/泛型响应）、服务层、REST API、错误处理中间件、单元测试、开发脚本。完成本阶段，你具备初级 TS 后端工程师的工程能力。

### 8.1 项目初始化

```bash
mkdir task-api && cd task-api
npm init -y

# 安装开发依赖：
# typescript     —— 编译器
# @types/node    —— Node.js API 的类型定义（fs/http/process...）
# ts-node        —— 免编译直接运行 .ts（开发用）
# nodemon        —— 文件变化自动重启
npm install -D typescript @types/node ts-node nodemon

# 安装运行时依赖：Express 及其类型
npm install express
npm install -D @types/express

npx tsc --init    # 生成 tsconfig.json，改成 7.5 节的 Node 推荐配置
```

**目录结构**（类型与逻辑分离、服务与路由分离）：

```
task-api/
├── src/
│   ├── types/            # 类型定义层
│   │   └── index.ts
│   ├── services/         # 业务逻辑层
│   │   └── taskService.ts
│   ├── middleware/       # 中间件（错误处理等）
│   └── index.ts          # 入口：Express 应用与路由
├── __tests__/            # 测试
├── dist/                 # 编译产物（git 忽略）
├── package.json
└── tsconfig.json
```

**package.json 脚本**：

```json
{
    "scripts": {
        "build": "tsc",
        "start": "node dist/index.js",
        "dev": "nodemon --exec ts-node src/index.ts"
    }
}
```

### 8.2 类型分层设计 ⭐

后端类型设计的核心思想：**不同边界用不同类型，用工具类型推导而不是手抄**。

```typescript
// ===== src/types/index.ts =====

// 1. 实体类型：数据库/内存中的完整形状
export type TaskStatus = "pending" | "in-progress" | "completed";
export type TaskPriority = "low" | "medium" | "high";

export interface Task {
    id: string;
    title: string;
    description?: string;
    status: TaskStatus;
    priority: TaskPriority;
    createdAt: string;      // ISO 字符串（JSON 可序列化，比 Date 实用）
    updatedAt: string;
    tags?: string[];
}

// 2. DTO（Data Transfer Object）：各 API 边界的输入形状——全部用工具类型从实体推导！
export interface CreateTaskInput {
    title: string;
    description?: string;
    priority: TaskPriority;
    dueDate?: string;
    tags?: string[];
}

// 更新接口 = 实体去掉不可变字段后的 Partial
export type UpdateTaskInput = Partial<Omit<Task, "id" | "createdAt" | "updatedAt">>;

// 3. 查询过滤类型
export interface TaskFilter {
    status?: TaskStatus;
    priority?: TaskPriority;
    search?: string;
}

// 4. 泛型 API 响应：一套包装，处处复用
export interface ApiResponse<T> {
    success: boolean;
    data?: T;
    error?: string;
}

// 5. 分页
export interface PaginationMeta {
    total: number;
    page: number;
    pageSize: number;
    totalPages: number;
}
export interface PaginatedResponse<T> {
    items: T[];
    meta: PaginationMeta;
}
```

体会这个设计的收益：实体加一个字段，所有 DTO/响应类型自动同步；`UpdateTaskInput` 根本不允许调用方传 id（`Omit` 掉了）——**非法请求在类型层面就不存在**。

### 8.3 服务层：纯业务逻辑

```typescript
// ===== src/services/taskService.ts =====
import { Task, CreateTaskInput, UpdateTaskInput, TaskFilter } from "../types";

function generateId(): string {
    return Date.now().toString(36) + Math.random().toString(36).slice(2, 8);
}

class TaskService {
    private tasks: Task[] = [];          // 内存存储（毕设项目可替换为数据库）

    getAll(filter?: TaskFilter): Task[] {
        let result = [...this.tasks];    // 拷贝，防外部篡改内部状态
        if (filter?.status)   result = result.filter(t => t.status === filter.status);
        if (filter?.priority) result = result.filter(t => t.priority === filter.priority);
        if (filter?.search) {
            const q = filter.search.toLowerCase();
            result = result.filter(t =>
                t.title.toLowerCase().includes(q) ||
                t.description?.toLowerCase().includes(q)
            );
        }
        return result;
    }

    getById(id: string): Task | undefined {
        return this.tasks.find(t => t.id === id);
    }

    create(input: CreateTaskInput): Task {
        const now = new Date().toISOString();
        const task: Task = {
            id: generateId(),
            title: input.title,
            description: input.description,
            status: "pending",             // 初始状态由服务决定，不信任客户端
            priority: input.priority,
            createdAt: now,
            updatedAt: now,
            tags: input.tags,
        };
        this.tasks.push(task);
        return task;
    }

    update(id: string, input: UpdateTaskInput): Task | null {
        const index = this.tasks.findIndex(t => t.id === id);
        if (index === -1) return null;
        const updated: Task = {
            ...this.tasks[index],          // 展开保留原字段
            ...input,                      // 覆盖要改的字段
            updatedAt: new Date().toISOString(),
        };
        this.tasks[index] = updated;
        return updated;
    }

    delete(id: string): boolean {
        const index = this.tasks.findIndex(t => t.id === id);
        if (index === -1) return false;
        this.tasks.splice(index, 1);
        return true;
    }
}

// 导出单例：路由层共享同一份数据
export const taskService = new TaskService();
```

注意服务层**不依赖 Express**——它是纯 TS，可独立测试（8.5 正是这么测的）。

### 8.4 Express REST API

```typescript
// ===== src/index.ts =====
import express, { Request, Response, NextFunction } from "express";
import { taskService } from "./services/taskService";
import { CreateTaskInput, Task, ApiResponse } from "./types";
import { AppError, NotFoundError } from "./errors";

const app = express();
app.use(express.json());                  // 解析 JSON 请求体

// 类型化的请求处理器：Request 泛型参数标注 params/body
interface IdParams { id: string; }

app.get("/api/tasks", (req: Request, res: Response) => {
    const tasks = taskService.getAll(req.query as Partial<TaskFilter>);
    const response: ApiResponse<Task[]> = { success: true, data: tasks };
    res.json(response);
});

app.get("/api/tasks/:id", (req: Request<IdParams>, res: Response) => {
    const task = taskService.getById(req.params.id);
    if (!task) throw new NotFoundError("任务");       // 统一抛给错误中间件
    res.json({ success: true, data: task });
});

app.post("/api/tasks", (req: Request<{}, unknown, CreateTaskInput>, res: Response) => {
    // 第三个泛型参数是 body 类型：编辑器里 req.body. 会有完整提示
    const task = taskService.create(req.body);
    res.status(201).json({ success: true, data: task });
});

app.put("/api/tasks/:id", (req: Request<IdParams, unknown, UpdateTaskInput>, res: Response) => {
    const task = taskService.update(req.params.id, req.body);
    if (!task) throw new NotFoundError("任务");
    res.json({ success: true, data: task });
});

app.delete("/api/tasks/:id", (req: Request<IdParams>, res: Response) => {
    const ok = taskService.delete(req.params.id);
    if (!ok) throw new NotFoundError("任务");
    res.status(204).end();
});

// 错误处理中间件：四参数签名，必须放最后
app.use((err: Error, req: Request, res: Response, next: NextFunction) => {
    if (err instanceof AppError) {
        res.status(err.statusCode).json({ success: false, error: err.message });
    } else {
        console.error("未预期错误:", err);
        res.status(500).json({ success: false, error: "服务器内部错误" });
    }
});

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`服务器运行在 http://localhost:${PORT}`);
});
```

配套的自定义错误（7.8 模式的落地）：

```typescript
// ===== src/errors.ts =====
export class AppError extends Error {
    constructor(message: string, public statusCode: number = 400) {
        super(message);
        Object.setPrototypeOf(this, AppError.prototype);
    }
}
export class NotFoundError extends AppError {
    constructor(resource: string) {
        super(`${resource} 不存在`, 404);
    }
}
```

运行开发：`npm run dev`（nodemon + ts-node，改代码自动重启）。用 curl 测试：

```bash
curl -X POST http://localhost:3000/api/tasks \
     -H "Content-Type: application/json" \
     -d '{"title":"学习 TS","priority":"high"}'
```

### 8.5 单元测试（Jest + ts-jest）

```bash
npm install --save-dev jest ts-jest @types/jest
npx ts-jest config:init      # 生成 jest.config.js
```

```javascript
// jest.config.js
module.exports = {
    preset: "ts-jest",
    testEnvironment: "node",
    testMatch: ["**/__tests__/**/*.ts"],
};
```

**测试服务层**（AAA 三段式：准备 Arrange / 执行 Act / 断言 Assert）：

```typescript
// ===== __tests__/taskService.test.ts =====
import { TaskService } from "../src/services/taskService";

describe("TaskService", () => {
    let service: TaskService;

    beforeEach(() => {
        service = new TaskService();   // 每个测试独立实例，互不污染
    });

    describe("create", () => {
        it("应创建任务且初始状态为 pending", () => {
            const task = service.create({ title: "T1", priority: "high" });
            expect(task.id).toBeDefined();
            expect(task.status).toBe("pending");
        });

        it("应自动生成递增不重复的 id", () => {
            const t1 = service.create({ title: "A", priority: "low" });
            const t2 = service.create({ title: "B", priority: "low" });
            expect(t1.id).not.toBe(t2.id);
        });
    });

    describe("update", () => {
        it("更新存在的任务并刷新 updatedAt", () => {
            const t = service.create({ title: "T", priority: "low" });
            const updated = service.update(t.id, { title: "新标题" });
            expect(updated?.title).toBe("新标题");
            expect(updated?.updatedAt >= t.updatedAt).toBe(true);
        });

        it("更新不存在的任务返回 null", () => {
            expect(service.update("不存在", { title: "x" })).toBeNull();
        });
    });

    describe("getAll 过滤", () => {
        it("按 status 过滤", () => {
            service.create({ title: "A", priority: "low" });
            const t = service.create({ title: "B", priority: "low" });
            service.update(t.id, { status: "completed" });
            expect(service.getAll({ status: "completed" })).toHaveLength(1);
        });
    });
});
```

运行：`npx jest`。测试代码同样全程有类型检查——`service.create({ title: 1 })` 在测试里就直接编译报错。

### 阶段八练习

- [ ] 从零搭建本项目（不要照抄：自己敲每一条命令、每一行代码）
- [ ] 增加分页查询 `GET /api/tasks?page=1&pageSize=10`，返回 `PaginatedResponse<Task>`
- [ ] 写一个校验中间件：POST 请求体缺失 title 或 priority 非法时返回 400 + 具体错误信息
- [ ] 给 update 的 404、删除的 204、非法 body 的 400 各写一个集成验证（curl 或测试）
- [ ] 补齐 TaskService 的测试到 90% 以上分支覆盖（getAll 三种过滤组合、delete 成败）
- [ ] 把 `taskService` 中的内存数组换成 `Map<string, Task>`，对外接口不变，测试全部照跑通过——体会“类型契约隔离实现变化”。（Map 是 JS 内置键值容器：`new Map()` 创建，`set(k, v)` 存、`get(k)` 取——无则 undefined、`has(k)` 判存在、`delete(k)` 删、`values()` 遍历所有值、`Array.from(m.values())` 转数组）

### 阶段八过关自检

1. @types/node、@types/express 为什么是开发依赖而不是运行时依赖？
2. 实体类型与 DTO 为什么要分开？UpdateTaskInput 是怎么"从类型层面禁止修改 id"的？
3. ApiResponse<T> 泛型包装带来什么好处？
4. 服务层不 import Express 有什么好处？
5. Express 错误处理中间件的签名有什么特征？它为什么必须注册在最后？
6. Request 泛型参数（如 `Request<IdParams, any, CreateTaskInput>`）分别对应什么？
7. 为什么测试里用 `new TaskService()` 而不是导入单例？
8. ts-node 与 tsc + node 两种运行方式分别适合什么场景？

---

## 阶段九：毕业项目

> 三个项目递进：A 巩固工程全流程；B 换一种程序形态（CLI + 文件 IO + 生成器/泛型管道）；C 触达真实工程形态（装饰器风格的路由注册，衔接 NestJS）。**至少完成 A 和 B**。

### 毕业项目 A：任务管理 REST API 完整版（2 天）

在阶段八骨架上扩展为一个完整服务，验收标准：

**功能需求**：

1. 任务 CRUD（已有）+ 以下新能力：
   - `GET /api/tasks?status=&priority=&search=&page=&pageSize=` 组合过滤分页
   - `PATCH /api/tasks/:id/status` 状态流转端点（pending → in-progress → completed，非法流转返回 422）
   - 标签管理：`PUT /api/tasks/:id/tags`
2. 统计端点 `GET /api/stats`：各状态数量、各优先级数量、完成率（用 reduce）
3. 数据持久化到 JSON 文件：启动时加载、变更后保存，用 Node 的异步文件接口（@types/node 提供完整类型）：

```typescript
import { readFile, writeFile } from "fs/promises";

async function loadTasks(): Promise<Task[]> {
    try {
        const raw = await readFile(DATA_FILE, "utf-8");
        return JSON.parse(raw) as Task[];   // 读出的 JSON 断言回实体类型
    } catch {
        return [];                          // 文件不存在时返回空表
    }
}
async function saveTasks(tasks: Task[]): Promise<void> {
    await writeFile(DATA_FILE, JSON.stringify(tasks, null, 2), "utf-8");
}
```
4. 完整错误体系：AppError / NotFoundError(404) / ValidationError(400) / 状态流转错误(422)，全部走错误中间件
5. 启动时可用环境变量 `PORT`、`DATA_FILE`（`process.env` 读取，给默认值）

**类型约束（体现全程所学）**：

- 实体 / CreateInput / UpdateInput / Filter / 分页全部工具类型化
- 状态流转表：`const TRANSITIONS: Record<TaskStatus, TaskStatus[]>` —— 状态机配置即类型
- 全部响应走 `ApiResponse<T>` / `PaginatedResponse<T>`

**质量要求**：服务层测试 ≥ 15 个用例；`npm run build && npm start` 零错误跑通；用 curl 完整演示一遍 CRUD + 错误路径。

### 毕业项目 B：CLI 工具 `todo-cli`（1.5 天）

做一个可安装的命令行待办工具，综合考查：Node API、生成器、泛型、模块拆分、装饰器日志。

```
用法：
  todo add "买牛奶" --priority high
  todo list [--status pending]
  todo done <id>
  todo stats
```

**技术要点**：

1. 参数解析自己写（`process.argv` 切片 + reduce 到结构化对象——这是 `Record` 的实战）
2. 数据存 `~/.todo.json`，`fs.promises` 读写，全部 async
3. 查询结果做**惰性管道**：`tasks.filter(...)` 改造成生成器链（`function* filtered()`、`function* paginated()`），体会 `for...of` 消费
4. 泛型 `Repository<T extends { id: string }>`：一个类同时管理 tasks 和 tags 两套数据
5. 服务方法加 `@log` 装饰器（7.6 手写过）：自动记录命令调用与耗时
6. package.json 配 `"bin"` 字段 + `npm link` 实现全局命令

### 毕业项目 C（进阶选做）：装饰器路由 —— 迷你 NestJS 风格（1.5 天）

用**旧版装饰器**复刻 NestJS 的核心机制，直通真实框架的心脏：

```typescript
@Controller("/tasks")
class TaskController {
    @Get("/:id")
    getOne(req: Request, res: Response) { /* ... */ }

    @Post("/")
    @Use(logMiddleware)
    create(req: Request, res: Response) { /* ... */ }
}
registerRoutes(app, TaskController);   // 扫描元数据自动注册路由
```

实现思路：方法装饰器把"HTTP 方法 + 路径 + 处理器"写入 `Reflect.defineMetadata`（或一个简单的 `Map`）；`registerRoutes` 遍历控制器原型读出元数据调用 `app.get/post(...)`。完成后重写阶段八的路由层，对比两种风格的优劣。

### 毕业验收清单

- [ ] 项目 A/B 可 `git clone && npm install && npm run build` 后直接运行
- [ ] 无 any（`grep -rn ": any" src/` 为空或仅有注释说明的特殊豁免）
- [ ] strict 模式零错误编译
- [ ] 每个公开函数/类有类型完整的签名（参数与返回值）
- [ ] 错误路径全部结构化（不是 console.log 完事）
- [ ] 服务层测试通过
- [ ] 能向别人讲清楚：类型分层怎么设计的、为什么这么分

---

## 学习原则

1. **两遍学习法**：第一遍走完本路线、做完项目——目标是"会用、敢用"；第二遍**精读官方 Handbook**——此时你已有全局地图，重读会不断产生"原来如此"的连接。官方 Handbook 推荐精读顺序：Basics → Everyday Types → Narrowing → Object Types → More on Functions → Classes → Modules → Type Manipulation 全部小节 → Reference 里的 Enums / Utility Types / Type Compatibility / Decorators。
2. **编译器是老师**：把每个例子里的错误故意制造一遍，读完整错误信息（TypeScript 的错误提示带有"为什么"的推导链），这是最快的进阶方式。
3. **类型为语义服务**：先想清楚数据的合法状态（可辨识联合）、非法操作（Omit 掉的字段），再落类型；而不是给现有代码"刷类型注解"。
4. **少用逃生舱**：any、`as` 双重断言、`@ts-ignore` 每用一次写一行注释说明原因；代码评审时它们是重点关照对象。
5. **练习不超纲 → 主动超纲**：路线内练习只用已学知识；做毕业项目时主动查证（编辑器提示、编译错误）解决新问题，这正是第二遍学习的开始。

## 附录：延伸资料

| 资料 | 定位 |
|---|---|
| TypeScript 官方 Handbook | 类型系统的权威叙述，第二遍学习的主线 |
| TypeScript 官方 tsc 配置手册 | tsconfig 每个选项的官方释义 |
| TypeScript Playground | 官方在线运行/分享 TS 代码，验证类型行为的首选 |
| Type Challenges | 类型体操题库（easy → extreme），阶段六后的健身房 |
| TypeScript Deep Dive | 社区经典进阶书，深挖语义细节 |
| Effective TypeScript（书） | 62 条 TS 使用法则，工程实践向 |
| DefinitelyTyped（GitHub 仓库） | @types 包的家，学写高质量 .d.ts 的范本 |
| MDN Web Docs | JS 语言特性与运行时 API 的权威文档 |
| NestJS 官方文档 | 毕业项目 C 之后的后端框架主路线（装饰器 + DI 生态） |
| Express 官方文档 | 阶段八所用 Web 框架的完整 API |
| Jest 官方文档 | 测试框架完整能力（mock、快照、覆盖率） |
| zod | 运行时数据校验库，弥补"类型编译后擦除"的缺口 |
| pnpm 官方文档 | workspace/monorepo 管理 |
| 2ality（博客） | ECMAScript 新特性的深度解析 |

## 后续深入主题（超出本路线范围，学完后按需展开）

- **前端方向**：React + TypeScript（组件 Props 类型、自定义 Hook 类型）、Vue 3 + TypeScript（组合式 API 类型）、Next.js/Nuxt 全栈框架
- **后端框架**：NestJS（依赖注入、模块系统、装饰器路由——毕业项目 C 的完全体）、Fastify、中间件/拦截器/管道设计
- **运行时校验**：zod / io-ts——"外部数据不可信"问题的系统解法（schema 即类型）
- **模块与构建**：ESM/CJS 双模式发布、tsup/esbuild 打包库、Tree Shaking 深入
- **类型进阶**：Type Challenges hard/extreme、变型注解（variance annotations）、协变逆变的编译器实现、Symbol 深入（unique symbol 与内置 well-known Symbol）
- **工程化进阶**：Monorepo 工具链、渐进迁移大型 JS 存量项目、TypeScript 编译器 API（写代码生成/Lint 工具）
- **新运行时**：Bun、Deno（原生 TS 支持，免编译直接跑 .ts）
- **语言演进**：Explicit Resource Management（using 声明）、装饰器标准版生态迁移、Go 重写的新一代编译器（TypeScript 7 计划，10 倍性能目标，语法不变）
