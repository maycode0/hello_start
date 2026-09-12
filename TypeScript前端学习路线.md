# TypeScript 前端学习路线：React 方向专项

> **总时长**：约 10 天（2 周），每天 4~5 小时，合计约 40~50 小时
> **版本基线**：TypeScript 5.x、React 18、Vite 5.x、现代浏览器（Chrome/Edge/Firefox）
> **前置要求**：已完成《TypeScript 学习路线（Node.js 后端版）》**阶段一~七**——即：掌握 JS 语言核心（函数/对象/数组/异步/解构）、TS 类型系统全部（类型注解/接口/类/类型收窄/泛型/条件类型/工具类型）、模块系统与 tsconfig/声明文件等工程化。本份不再重复语言内容，遇到时直接使用
> **内容构成**：菜鸟教程 React 教程核心章节（组件/JSX/Props/State/事件/表单/Hooks/refs/memo/自定义 Hook/Vite 安装）+ 菜鸟 TypeScript 教程的 React 实战与综合项目实战两章 + TypeScript 官方 Handbook 的 JSX 参考章 + 官方类型声明章（lib.dom 部分相关内容）+ 菜鸟 JavaScript 教程的 DOM/事件概念章节，五源融合
> **读者设定**：TS 已过关、React 零基础、有一点 HTML 基础（边做边补）的前端方向学习者
> **React 版本约定**：全程使用**函数组件 + Hooks**（现代标准写法）；类组件仅在必要处理解性提及

---

## 总览：这份专项路线补什么

第一份路线教会了你"TypeScript 这门语言"和"Node.js 后端工程"。这份专项补上剩下的半壁江山——**浏览器端的 TS 开发**。三条主线：

1. **环境的换轨**：代码从 Node.js 运行时搬到浏览器——你要认识 DOM（浏览器给 JS 的 API 世界）、新的 tsconfig 形态（`jsx`、`lib: DOM`、`noEmit`）、新的构建工具（Vite 接管编译）。
2. **React 从零到能用**：组件、Props、State、事件、Hooks（含自定义 Hook）——全部以 TypeScript 视角讲授：每个概念先看 JS 写法，再看"加上类型后的正确姿势"。
3. **类型系统的前端战场**：Props 接口设计、事件类型、`useState` 泛型、受控表单、样式映射、请求层 `ApiResponse<T>`——第一份学的泛型/工具类型/可辨识联合，在这里全部找到真实用武之地。

**阶段地图**：

| 阶段 | 主题 | 天数 |
|---|---|---|
| 一 | 换轨浏览器：DOM、事件与 TS 前端工程（Vite + .tsx） | 2 天 |
| 二 | React 核心：组件、Props、State、事件与表单（全程 TS） | 2 天 |
| 三 | Hooks 深入：useEffect/useRef/性能三件套与自定义 Hook | 2 天 |
| 四 | 前端工程化：tsconfig 深解、类型生态、请求层与样式类型 | 2 天 |
| 五 | 毕业项目：任务管理前端（本地版 + 对接后端 API 版） | 2 天 |

---

## 阶段一：换轨浏览器——DOM、事件与 TS 前端工程

> **目标**：理解浏览器端的 JS 世界（DOM 与事件），搞清 TS 如何为浏览器 API 提供类型，用 Vite 建起第一个 React + TS 项目，掌握 .tsx 文件的特殊类型规则。

### 1.1 浏览器里的 JavaScript：DOM 与事件 ⭐

后端代码跑在 Node.js 里，前端代码跑在浏览器里。浏览器给 JS 提供的最重要 API 是 **DOM（Document Object Model，文档对象模型）**：页面加载时，浏览器把 HTML 解析成一棵**对象树**，JS 通过操作这棵树来读写页面。

```javascript
// DOM 树：document（根）→ html → body → div/p/button ...
// 查找元素的三种经典方式：
const byId = document.getElementById("intro");       // 按 id
const byTag = document.getElementsByTagName("p");     // 按标签名
const byClass = document.getElementsByClassName("x"); // 按类名
// 现代写法（TS 时代的主力）：
const el = document.querySelector("#intro");          // CSS 选择器，找第一个
const list = document.querySelectorAll(".item");      // 找全部

// 找到元素后可以：
el.innerHTML = "新内容";       // 改 HTML 内容
el.style.color = "red";        // 改样式
el.setAttribute("href", "#");  // 改属性
```

**事件**：浏览器里"发生的事情"——用户点击、输入、页面加载完成。事件发生时执行对应的处理代码：

| 常见事件 | 触发时机 |
|---|---|
| onclick | 用户点击元素 |
| onchange | 元素内容改变（输入框失焦后） |
| onmouseover / onmouseout | 鼠标移入 / 移出 |
| onkeydown | 按下键盘按键 |
| onload | 页面加载完成 |

原生写法是 HTML 属性挂事件（`<button onclick="fn()">`）或 `addEventListener`。**React 用自己的事件系统**（阶段二详讲），原生方式只做理解用。

**命令式的痛**（React 存在的理由）：原生 DOM 开发是"命令式"——你告诉浏览器**怎么做**，一步步操作 DOM：

```javascript
// 命令式：手动查元素、读当前值、算新值、写回去
const button = document.getElementById("myButton");
button.addEventListener("click", function () {
    const counter = document.getElementById("counter");
    const currentValue = parseInt(counter.textContent);
    counter.textContent = currentValue + 1;
});
```

状态散落在 DOM 里，逻辑与视图搅在一起——这正是 React 要解决的问题。

### 1.2 TS 如何认识 DOM：lib.dom.d.ts ⭐

`document`、`window`、`HTMLElement` 这些名字，TS 从哪里知道它们的类型？答案是**内置声明文件**：TS 自带一套描述浏览器 API 的类型库（文件名形如 `lib.dom.d.ts`），是否加载由 tsconfig 的 `lib` 选项控制。

```jsonc
// 前端 tsconfig 必须包含 DOM 库（第一份 Node 路线特意排除了它）：
{ "compilerOptions": { "lib": ["ES2020", "DOM", "DOM.Iterable"] } }
```

`DOM` 提供全部浏览器 API 类型，`DOM.Iterable` 补充 DOM 集合的迭代支持（`for...of` NodeList 等）。

**HTMLElement 家族与类型收窄**——DOM API 返回值大多可空、大多宽泛，正好复用第一份学的收窄与断言：

```typescript
// querySelector 返回 Element | null：先判空
const btn = document.querySelector("#submit");
if (btn) {
    (btn as HTMLButtonElement).disabled = true;   // Element → 具体元素类型
}

// 更精确的选择器 + 断言一步到位：
const input = document.querySelector("#username") as HTMLInputElement;
console.log(input.value);        // HTMLInputElement 才有 value 属性

// 事件对象的类型：event.target 是 EventTarget | null，先断言再用
document.querySelector("#form")?.addEventListener("submit", (e) => {
    e.preventDefault();
    const target = e.target as HTMLFormElement;
    console.log(target.id);
});
```

元素类型速查（都在 lib.dom 里）：`HTMLElement`（基类）、`HTMLInputElement`（输入框）、`HTMLButtonElement`、`HTMLFormElement`、`HTMLSelectElement`（下拉框）、`HTMLDivElement` 等——"HTML + 标签名 + Element"命名规律。

### 1.3 Vite + React + TypeScript 工程 ⭐

Vite 是当前 React 官方推荐的构建工具：开发时秒级启动、热更新（改代码页面即时刷新，不用手动重载）；生产构建基于 Rollup。开箱支持 JSX、TypeScript、CSS Modules。

```bash
# 一条命令创建 React + TS 项目（--template react-ts 是关键）：
npm create vite@latest my-app -- --template react-ts

cd my-app
npm install       # 安装依赖
npm run dev       # 启动开发服务器 → 浏览器打开 http://localhost:5173
```

**项目结构**（对照 Node 项目，差异一目了然）：

```
my-app/
├── public/             # 静态资源（原样复制进构建产物）
├── src/
│   ├── assets/         # 图片等资源
│   ├── App.tsx         # 主组件（.tsx = 含 JSX 的 TS 文件）
│   ├── main.tsx        # 入口：把 App 挂载到页面
│   ├── index.css       # 全局样式
│   └── vite-env.d.ts   # Vite 环境的类型声明
├── index.html          # 页面入口模板（React 挂载点在这里）
├── package.json
├── tsconfig.json       # TS 配置（阶段四深解）
└── vite.config.ts      # Vite 配置
```

**入口文件 main.tsx**——React 应用与页面的连接点：

```tsx
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";

ReactDOM.createRoot(document.getElementById("root")!).render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);
// createRoot：React 18 挂载 API
// getElementById 返回 HTMLElement | null，这里用非空断言 !（index.html 里必有 #root）
// StrictMode：开发模式下的严格检查（重复渲染暴露副作用问题）
```

**双工具分工**（对比第一份的 Node 工作流）：

| 职责 | Node 项目（第一份） | 前端项目（本份） |
|---|---|---|
| 类型检查 | tsc | tsc（`noEmit: true` 只检查不产出） |
| 代码转译/打包 | tsc 编译出 .js | **Vite**（esbuild 转译 + Rollup 打包） |
| 运行 | node dist/index.js | 浏览器（dev 服务器 / `npm run build` 产物部署） |

`npm run build` 产出 `dist/` 静态文件，可部署到 Vercel、Netlify 等任意静态托管。

### 1.4 .tsx 文件与 JSX 的类型规则 ⭐

**JSX** 是 JavaScript 的语法扩展：在 TS 里写类似 HTML 的标签来描述 UI。官方对 JSX 的 TS 支持有明确约定：

1. **文件必须用 `.tsx` 扩展名**（.ts 文件里的 JSX 会报错）；
2. tsconfig 开启 `jsx` 选项（常用 `"jsx": "react-jsx"`，阶段四详解五种模式）。

**`.tsx` 里的断言只能用 `as`**——尖括号断言 `<string>x` 与 JSX 标签语法冲突，在 .tsx 中被禁用：

```tsx
const foo = bar as Foo;    // ✅ .tsx 唯一合法写法
// const foo = <Foo>bar;   // ❌ 在 .tsx 中被解析为 JSX 标签，报错
```

（第一份路线里"统一用 as"的建议，在 .tsx 里升级为硬性规则。）

**大小写约定**——JSX 标签首字母决定类型检查方式：

```tsx
<div />        // 小写开头：内置元素（intrinsic）→ 查 JSX.IntrinsicElements 接口
<MyComponent /> // 大写开头：值元素 → 按作用域内的组件（函数/类）解析

// 内置元素的类型清单由 React 的类型库注册（div、span、button 等全量登记），
// 所以 <div className="x"> 的每个属性都有类型检查——拼错属性名直接报错。
// 自定义组件没在作用域内导入时，<MyComponent /> 直接编译报错。
```

**属性检查与 children**：内置元素按注册的属性类型检查（`className` 缺失/拼写错会报错；无效标识符属性如 `data-*` 豁免）；组件标签按其 **Props 第一个参数**检查属性；标签的嵌套内容进入特殊的 `children` 属性（2.3 节展开）。

> 💡 伏笔：`JSX.IntrinsicElements` 这份"HTML 标签字典"是谁写进来的？阶段四讲 `@types/react` 类型生态时兑现。

### 阶段一练习

- [ ] 用 `npm create vite@latest` 创建 react-ts 项目并跑起来，把 App.tsx 标题改成自己的名字，体验保存即热更新
- [ ] 阅读 main.tsx，回答：`getElementById` 后面的 `!` 是什么？为什么这里可以放心用？
- [ ] 在项目里新建 `dom.ts`（普通 .ts），写一段原生 DOM 代码：`querySelector` 一个按钮，判空后断言为 HTMLButtonElement，`addEventListener` 点击时 `console.log(btn.disabled)`，在 index.html 里临时加个 `<button id="btn">测试</button>` 验证
- [ ] 对比实验：把同一段代码里的断言换成 `as unknown as HTMLDivElement`，观察 `.disabled` 报什么错，体会 IntrinsicElements/元素类型的检查强度
- [ ] 试着在 .tsx 里写 `<string>someValue` 断言，记录编译器报错信息

### 阶段一过关自检

1. DOM 是什么？`querySelector` 返回什么类型？为什么用之前必须处理 null？
2. `lib` 选项里的 "DOM" 和 "DOM.Iterable" 各提供什么？Node 项目为什么不要包含它们？
3. `.tsx` 和 `.ts` 的区别？.tsx 里写类型断言必须用什么语法，为什么？
4. `<div />` 和 `<Div />` 在类型检查上走的是两条什么路径？谁决定 `<div>` 有哪些合法属性？
5. Vite 工程里 tsc 和 Vite 各负责什么？`noEmit` 的意义？
6. React 应用是怎么"挂"到页面上的？描述 main.tsx 每一行的作用。

---

## 阶段二：React 核心——组件、Props、State、事件与表单

> **目标**：掌握 React 四大件（组件/Props/State/事件）与 JSX 全部语法规则，每个概念都落到"TS 怎么写才类型安全"。学完能独立写出类型完整的交互组件与受控表单。

### 2.1 React 思维：声明式与单向数据流 ⭐

对比 1.1 的命令式代码，React 是**声明式**——你描述“UI 是状态的函数”，不手动操作 DOM（示例中的 `useState` 是 React 的状态 Hook，先把它当作“带更新函数的可变变量”，2.4 详解）：

```tsx
// 声明式：告诉计算机要什么结果
function Counter() {
    const [count, setCount] = useState(0);
    return <button onClick={() => setCount(count + 1)}>{count}</button>;
}
// count 变 → React 自动重渲染 → 按钮文字自动更新
// 你全程没有碰 document
```

三大心智模型：

1. **组件化**：UI 拆成独立、可复用、可组合的组件（搭积木），每个组件封装自己的结构/样式/状态；
2. **单向数据流**：数据只能从父组件流向子组件（props 向下），子组件要改数据只能**调用父组件传下来的回调**（事件向上）：

```tsx
function Parent() {
    const [message, setMessage] = useState("Hello");
    return (
        <Child
            message={message}            // 数据向下
            onUpdate={setMessage}        // 回调向上
        />
    );
}
function Child({ message, onUpdate }: { message: string; onUpdate: (v: string) => void }) {
    return <button onClick={() => onUpdate("Updated!")}>{message}</button>;
}
```

3. **UI = f(state)**：状态是唯一真相，改状态（而不是改 DOM），React 负责算出最小 DOM 更新（虚拟 DOM diff）。

### 2.2 JSX 语法规则 ⭐

JSX 不是 HTML，规则比 HTML 严格：

```tsx
function Demo() {
    const name = "runoob";
    const items = [
        { id: 1, name: "苹果" },
        { id: 2, name: "香蕉" },
    ];
    return (
        <>
            <h1 className="title">你好，{name}！</h1>
            {/* ① 表达式嵌入：{} 里放任意 JS 表达式（不能放 if/for 语句） */}
            <p>计算：{2 + 3 * 5}</p>

            {/* ② 条件渲染：三元 或 && */}
            {name ? <p>欢迎，{name}</p> : <p>请登录</p>}
            {items.length > 0 && <p>共 {items.length} 项</p>}

            {/* ③ 列表渲染：map()，每个元素必须带唯一稳定的 key */}
            <ul>
                {items.map(item => (
                    <li key={item.id}>{item.name}</li>
                ))}
            </ul>

            {/* ④ 标签必须闭合；class → className；for → htmlFor */}
            <br />
            <label htmlFor="inp">名字</label>
            <input id="inp" className="input" />
        </>
    );
}
```

**规则清单**：

- **所有标签必须闭合**（`<br />`、`<img ... />`、`<div></div>`）
- **class → className，for → htmlFor**（class/for 是 JS 保留字；写错直接类型报错）
- **自定义组件必须大写开头**（小写被当 HTML 标签解析）
- **return 只能有一个根节点**——用 Fragment（空标签 `<>...</>`）包裹多个兄弟，不产生多余 DOM
- **注释写 `{/* ... */}`**
- **key 的纪律**：列表渲染必须给 key；用数据里的唯一 ID，**不要用数组下标**（顺序变化时引发状态错乱）；key 缺失 React 会警告

**样式两种基础写法**：

```tsx
// 内联样式：对象 + 驼峰属性名（动态值方便，静态样式难维护）
<div style={{ backgroundColor: "#f0f0f0", padding: 20, borderRadius: 8 }}>
    <h3 style={{ color: "blue" }}>标题</h3>
</div>
// 注意：数字会自动加 px；backgroundColor 是驼峰不是 background-color

// CSS Modules（Vite 原生支持，作用域隔离，推荐）：
// Card.module.css 里写 .container { ... }
import styles from "./Card.module.css";
<div className={styles.container}>
    <h3 className={styles.title}>标题</h3>
</div>
// 类名自动哈希（如 container__abc123），全局永不冲突
```

### 2.3 函数组件与 Props 类型 ⭐⭐

组件 = 返回 JSX 的函数。三种等价写法（团队统一即可）：

```tsx
// 方式 1：函数声明
function Welcome(props: WelcomeProps) {
    return <h1>Hello, {props.name}</h1>;
}
// 方式 2：箭头函数
const Welcome2 = (props: WelcomeProps) => <h1>Hello, {props.name}</h1>;
// 方式 3：解构参数（最常用，配合默认值）
const Welcome3 = ({ name, age = 18 }: WelcomeProps) => (
    <h1>Hello, {name}, {age}</h1>
);
```

**Props 接口——组件的公开契约**（第一份学的接口设计，在前端的主战场）：

```tsx
interface ButtonProps {
    text: string;                          // 必填
    onClick: () => void;                   // 回调也是类型的一部分
    disabled?: boolean;                    // 可选
    variant?: "primary" | "secondary";     // 字面量联合：变体收窄成枚举
}

// React.FC<Props>：React 提供的函数组件类型（含 children 等内置能力）
const Button: React.FC<ButtonProps> = ({
    text,
    onClick,
    disabled = false,
    variant = "primary",
}) => {
    return (
        <button className={`btn btn-${variant}`} onClick={onClick} disabled={disabled}>
            {text}
        </button>
    );
};

// 使用处全程类型检查：
<Button text="保存" onClick={() => save()} />;
// <Button text="保存" />;                  // ❌ 缺 onClick
// <Button text="保存" onClick={save} variant="danger" />;  // ❌ variant 不在联合内
```

**children：特殊的 Prop**——标签之间的嵌套内容：

```tsx
interface CardProps {
    title?: string;
    children: React.ReactNode;     // React.ReactNode：能放进 JSX 的一切（元素/字符串/数组…）
}
function Card({ title, children }: CardProps) {
    return (
        <div className="card">
            {title && <h2>{title}</h2>}
            <div className="card-body">{children}</div>
        </div>
    );
}
// 使用：
<Card title="我的卡片">
    <p>内容</p>
    <button>点我</button>
</Card>
```

**Props 的铁律：只读**——永远不要修改 props（`props.name = "x"` 是大忌），需要"变"就创建新值或用 state。对象/数组/函数/JSX 都能当 props 传，解构、默认值、展开全用第一份学的 JS 语法。

**组合模式**：小组件拼成大组件——

```tsx
function Avatar({ src, alt }: { src: string; alt: string }) {
    return <img src={src} alt={alt} className="avatar" />;
}
interface User { name: string; email: string; avatar: string; }
function UserCard({ user, onEdit, onDelete }: {
    user: User;
    onEdit: (user: User) => void;
    onDelete: (id: string) => void;
}) {
    return (
        <div className="user-card">
            <Avatar src={user.avatar} alt={user.name} />
            <h3>{user.name}</h3>
            <p>{user.email}</p>
            <button onClick={() => onEdit(user)}>编辑</button>
            <button onClick={() => onDelete(user.name)}>删除</button>
        </div>
    );
}
```

### 2.4 useState：状态的类型 ⭐

```tsx
import { useState } from "react";

function Counter() {
    // 基本类型：有初始值时自动推断（count: number）
    const [count, setCount] = useState(0);

    // 复杂/可空类型：显式泛型
    const [user, setUser] = useState<{ name: string; age: number } | null>(null);

    // 联合/字面量：显式泛型 + as const 初始值
    const [status, setStatus] = useState<"idle" | "loading" | "done">("idle");

    return (
        <div>
            <p>计数: {count}</p>
            <button onClick={() => setCount(count + 1)}>+1</button>
            {/* 函数式更新：基于上一个状态计算，避免批处理下的过期值 */}
            <button onClick={() => setCount(c => c + 1)}>函数式+1</button>

            {user && <p>{user.name}, {user.age}</p>}
            <button onClick={() => setUser({ name: "Alice", age: 25 })}>设置用户</button>
        </div>
    );
}
```

**类型要点**：

- 初始值能确定类型时靠推断（`useState(0)` → number）；**初始值是 null/undefined/空数组时必须显式泛型**：`useState<User[]>([])`、`useState<string | null>(null)`
- 更新对象状态用**展开保留 + 覆盖**（第一份 2.6 的对象展开在这里天天用）：`setUser({ ...user, age: user.age + 1 })`
- `setCount` 的参数类型自动锁定为 state 的类型——传错立即报错

### 2.5 事件处理与事件类型 ⭐

React 事件与原生事件两个差异：**命名驼峰**（onClick 不是 onclick）、**传函数不传字符串**（`onClick={handleClick}` 不是 `onclick="handleClick()"`）。

```tsx
function ActionLink() {
    function handleClick(e: React.MouseEvent<HTMLAnchorElement>) {
        e.preventDefault();      // 阻止默认行为必须显式调用（不能 return false）
        console.log("链接被点击");
    }
    return <a href="#" onClick={handleClick}>点我</a>;
}
```

**事件类型速查**（React 提供完整类型族，按需标注）：

| 事件类型 | 场景 | 常用属性 |
|---|---|---|
| `React.MouseEvent<T>` | onClick / onMouseOver | clientX、clientY |
| `React.ChangeEvent<T>` | onChange（input/select） | `e.target.value` |
| `React.FormEvent<T>` | onSubmit | preventDefault |
| `React.KeyboardEvent<T>` | onKeyDown | key、ctrlKey |
| `React.FocusEvent<T>` | onFocus / onBlur | relatedTarget |

泛型参数 `T` 是事件源元素类型（如 `ChangeEvent<HTMLInputElement>` 里 `e.target` 就是 HTMLInputElement，`.value` 有类型）。**技巧**：内联箭头函数时事件类型自动推断，不用手写——

```tsx
<input onChange={(e) => console.log(e.target.value)} />
// e 自动推断为 ChangeEvent<HTMLInputElement>，e.target.value 是 string
// 把处理器提出来单独定义时，才需要显式标注事件类型
```

### 2.6 受控表单 ⭐

**受控组件**：表单元素的值由 React state 驱动（`value={x}` + `onChange={setX}`），数据单向流动，输入即验证、提交时数据已在 state 里。完整 TS 表单三件套（input / select / submit）：

```tsx
import { useState } from "react";

type TaskPriority = "low" | "medium" | "high";   // 字面量联合当枚举用

interface TaskFormProps {
    onSubmit: (input: { title: string; priority: TaskPriority }) => void;
}

function TaskForm({ onSubmit }: TaskFormProps) {
    const [title, setTitle] = useState("");
    const [priority, setPriority] = useState<TaskPriority>("medium");

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();                     // 阻止页面刷新
        if (!title.trim()) return;              // 简单校验
        onSubmit({ title: title.trim(), priority });
        setTitle("");                           // 重置表单
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="任务标题"
            />
            <select
                value={priority}
                onChange={(e) => setPriority(e.target.value as TaskPriority)}
            >
                <option value="low">低</option>
                <option value="medium">中</option>
                <option value="high">高</option>
            </select>
            <button type="submit">添加</button>
        </form>
    );
}
```

**`e.target.value as TaskPriority` 模式**：`select` 的 value 永远是 string，但选项集合是封闭的——断言回字面量联合是安全且标准的做法（3.5 节系统化）。

### 阶段二练习

- [ ] 写 `Badge` 组件：Props 含 `count: number`、`max?: number`（超出显示 `max+`），红色圆角徽章；用内联样式实现
- [ ] 写 `TodoList` 组件：state 是 `string[]`，渲染列表（map + key），含输入框与"添加"按钮（受控输入），点击条目删除
- [ ] 把上面的删除改为"点击切换完成状态"：state 改为 `{ text: string; done: boolean }[]`，用交叉类型/接口定义条目，已完成条目加删除线样式
- [ ] 写 `PriceTag` 组件：Props 为 `price: number | string`（数字或 "¥12.50" 格式），在组件内用类型收窄统一显示格式——体会联合类型 Props 的设计
- [ ] 复刻 2.6 的 TaskForm，额外加"描述"多行文本框（textarea，同样受控）和空标题时的红色错误提示（条件渲染）

### 阶段二过关自检

1. 命令式和声明式的区别？单向数据流里子组件如何影响父组件？
2. JSX 与 HTML 的五条关键差异（闭合/className/大写/单根/表达式）？
3. 列表渲染为什么必须 key？为什么不该用下标当 key？
4. React.FC\<Props\> 是什么？Props 接口里回调函数怎么定义类型？
5. children 是什么？类型用什么？
6. 什么时候 useState 必须显式泛型？函数式更新解决什么问题？
7. `onChange={(e) => ...}` 里 e 的类型是谁推断的？独立声明的处理器怎么标注？
8. 什么是受控组件？`value` 和 `onChange` 各承担什么角色？

---
## 阶段三：Hooks 深入——useEffect、useRef、性能三件套与自定义 Hook

> **目标**：掌握函数组件的核心进阶能力：useEffect（副作用与数据获取）、useRef（DOM 引用）、React.memo/useMemo/useCallback（性能优化），并学会用自定义 Hook 封装可复用逻辑——每个都带完整类型写法。这是 React 与 TS 结合力最强的部分。

### 3.1 useEffect：副作用与数据获取 ⭐⭐

useEffect 处理"渲染之外的事"：请求、订阅、定时器、手动改 DOM 标题。签名：`useEffect(副作用函数, 依赖数组)`。

```tsx
import { useState, useEffect } from "react";

interface User { id: number; name: string; }

function DataFetcher() {
    const [users, setUsers] = useState<User[]>([]);          // 空数组必须显式泛型
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        fetch("/api/users")
            .then(res => res.json())
            .then((data: User[]) => {                        // 响应断言为实体类型
                setUsers(data);
                setLoading(false);
            })
            .catch((err: unknown) => {
                setError(err instanceof Error ? err.message : "加载失败");
                setLoading(false);                            // unknown 收窄（第一份 7.8）
            });
    }, []);   // 空依赖：仅在组件挂载时执行一次

    // 可辨识联合 + JSX 条件渲染：三种互斥状态分别渲染
    if (loading) return <div>加载中...</div>;
    if (error) return <div>错误: {error}</div>;
    return (
        <ul>
            {users.map(user => <li key={user.id}>{user.name}</li>)}
        </ul>
    );
}
```

**依赖数组三种形态**：

| 写法 | 执行时机 |
|---|---|
| 不传 | 每次渲染后都执行（几乎不用） |
| `[]` | 仅挂载时一次（初始加载） |
| `[a, b]` | 挂载时 + a/b 变化时（监听过滤条件等） |

**清理函数**：副作用函数返回一个函数，组件卸载时执行（定时器/订阅必须清理）：

```tsx
useEffect(() => {
    const timer = setInterval(() => setSeconds(s => s + 1), 1000);
    return () => clearInterval(timer);   // 卸载时清理，防内存泄漏
}, []);
```

**loading/error/data 三件套建模**（可辨识联合的优雅版——把三个 useState 合成一个状态机）：

```tsx
type FetchState<T> =
    | { status: "loading" }
    | { status: "error"; message: string }
    | { status: "success"; data: T };      // 泛型携带数据类型

function useUserList(): FetchState<User[]> {
    const [state, setState] = useState<FetchState<User[]>>({ status: "loading" });
    // 渲染处 switch (state.status) 收窄：每个分支访问自己的专属字段
}
// 比三个布尔 state 更安全：不可能出现 loading 和 error 同时 true 的非法状态
```

### 3.2 useRef：DOM 引用与可变容器 ⭐

两种用途：

```tsx
import { useRef, useEffect } from "react";

function FocusDemo() {
    // 用途 1：引用 DOM 元素——泛型写元素类型，初始值 null
    const inputRef = useRef<HTMLInputElement>(null);

    useEffect(() => {
        // ref.current 可能为 null（挂载前），用 ?. 或判断
        inputRef.current?.focus();          // 挂载后自动聚焦
    }, []);

    return <input ref={inputRef} />;        // ref 属性挂到元素上
    // 挂载后 inputRef.current 就是那个 HTMLInputElement
    // 读值：inputRef.current?.value；赋值也要判空：inputRef.current!.value = "x"
}
```

```tsx
// 用途 2：跨渲染保存"不触发重渲染"的可变值（如定时器 id、上一次的值）
const timerRef = useRef<number | undefined>(undefined);
useEffect(() => {
    timerRef.current = setInterval(() => {}, 1000);
    return () => clearInterval(timerRef.current);
}, []);
// 改 ref.current 不会引发重渲染——这是它与 state 的本质区别
```

**纪律**：不要在渲染期间读写 ref（渲染必须纯粹）；操作 DOM 前先处理 null。

### 3.3 性能三件套：React.memo / useMemo / useCallback ⭐

三者分工（面试高频）：

```tsx
import { memo, useMemo, useCallback, useState } from "react";

// ① React.memo：记忆整个组件——props 浅比较没变就跳过重渲染
interface ItemProps { count: number; onClick: () => void; }
const Child = memo(({ count, onClick }: ItemProps) => {
    console.log("Child 渲染了");
    return <button onClick={onClick}>{count}</button>;
});

function App() {
    const [count, setCount] = useState(0);
    const [text, setText] = useState("hello");

    // 问题：App 每次渲染都会新建 increment 函数 → 传给 Child 的 props 变了 → memo 失效
    // ② useCallback：记忆函数——依赖不变时返回同一个引用
    const increment = useCallback(() => setCount(c => c + 1), []);
    // ③ useMemo：记忆计算结果——依赖不变时不重算
    const doubled = useMemo(() => count * 2, [count]);

    return (
        <>
            <Child count={count} onClick={increment} />
            <input value={text} onChange={e => setText(e.target.value)} />
            {/* 改 text 时：App 重渲染，但 Child 的 props 没变（memo）+ increment 没变（useCallback）→ Child 跳过渲染 */}
            <p>双倍：{doubled}</p>
        </>
    );
}
```

| 工具 | 记忆对象 | 典型场景 |
|---|---|---|
| React.memo | 组件 | 大列表项、复杂子组件 |
| useMemo | 值/计算结果 | 昂贵计算、引用稳定性（传给 memo 子组件的对象/数组） |
| useCallback | 函数 | 传给 memo 子组件的回调 |

⚠ **性能优化的纪律**：先测量再优化；普通组件直接写，慢了再上三件套。React.memo 是**浅比较**（嵌套对象每次都算"变了"）；滥用 useMemo/useCallback 自身也有记忆开销。类型签名：`useCallback(fn, deps)` 返回与 fn 同类型；`useMemo(() => value, deps)` 返回 factory 的返回类型——都自动推断。

### 3.4 自定义 Hook：封装可复用逻辑 ⭐⭐

命名以 `use` 开头的函数即是 Hook（React 靠此约定启用特殊能力）。**规则**：只在函数顶层调用其他 Hook（不能在 if/循环里）；自定义 Hook 内部可以自由使用 useState/useEffect 等。

**完整案例：useTasks**（综合项目实战的核心抽象——把"任务数据 + 增删改 + loading/error"整体封装，组件只剩渲染）：

```tsx
// src/hooks/useTasks.ts
import { useState, useEffect, useCallback } from "react";
import { TaskService } from "../services/taskService";
import type { Task, CreateTaskInput, UpdateTaskInput, TaskFilter } from "../types";

// Hook 的返回值先定义成接口——这是 Hook 的"公开契约"
interface UseTasksReturn {
    tasks: Task[];
    loading: boolean;
    error: string | null;
    filter: TaskFilter;
    createTask: (input: CreateTaskInput) => void;
    updateStatus: (id: string, status: Task["status"]) => void;
    deleteTask: (id: string) => void;
    setFilter: (filter: TaskFilter) => void;
    refresh: () => void;
}

export function useTasks(): UseTasksReturn {
    const [tasks, setTasks] = useState<Task[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [filter, setFilter] = useState<TaskFilter>({});

    const loadTasks = useCallback(() => {
        setLoading(true);
        setError(null);
        try {
            setTasks(TaskService.getAll(filter));       // 同步版 service（异步版见 4.3）
        } catch (err) {
            setError(err instanceof Error ? err.message : "加载失败");
        } finally {
            setLoading(false);
        }
    }, [filter]);   // 过滤条件变化 → loadTasks 引用更新

    useEffect(() => {
        loadTasks();     // 挂载时 + filter 变化时重新加载
    }, [loadTasks]);

    const createTask = useCallback((input: CreateTaskInput) => {
        TaskService.create(input);
        loadTasks();     // 操作后刷新
    }, [loadTasks]);

    const deleteTask = useCallback((id: string) => {
        TaskService.delete(id);
        loadTasks();
    }, [loadTasks]);

    const updateStatus = useCallback((id: string, status: Task["status"]) => {
        TaskService.update(id, { status });
        loadTasks();
    }, [loadTasks]);

    return { tasks, loading, error, filter, createTask, updateStatus, deleteTask, setFilter, refresh: loadTasks };
}

// 组件里使用——三行拥有全部能力：
function TaskList() {
    const { tasks, loading, error, deleteTask } = useTasks();
    if (loading) return <div>加载中...</div>;
    if (error) return <div>错误: {error}</div>;
    return <>{tasks.map(t => <div key={t.id}>{t.title} <button onClick={() => deleteTask(t.id)}>删</button></div>)}</>;
}
```

注意 `Task["status"]`——索引访问类型直接复用实体字段类型，Task 的 status 一改，Hook 签名自动同步（第一份 5.2 的实战兑现）。

**自定义 Hook 的设计准则**：返回接口显式定义（调用方有完整提示）；内部 useCallback 锁定函数引用（配合调用方的 memo）；每个 Hook 只管一个领域（useTasks 不管 UI，useFilter 不管数据）。

### 3.5 组件里的 TS 模式集 ⭐

**① Record 映射状态/样式**——字面量联合 → 类名/文案/图标的完整映射（编译期保证全覆盖）：

```tsx
type TaskStatus = "pending" | "in-progress" | "completed";

const statusStyles: Record<TaskStatus, string> = {
    "pending": "status-pending",
    "in-progress": "status-progress",     // 漏写任何一个键 → 编译报错
    "completed": "status-completed",
};
const statusLabels: Record<TaskStatus, string> = {
    "pending": "待处理", "in-progress": "进行中", "completed": "已完成",
};

function StatusBadge({ status }: { status: TaskStatus }) {
    return <span className={statusStyles[status]}>{statusLabels[status]}</span>;
    // 新增状态时 Record 强制你补全两个映射——switch 漏分支的绝缘升级版
}
```

**② 事件值断言模式**：封闭选项的 `e.target.value` 断言回字面量联合（2.6 已见，此处明确适用边界）——仅当选项集合由代码完全控制时安全；来自用户自由输入的值（input 文本）绝不能断言，该校验就校验。

**③ 回调 Props 与状态提升**：两个兄弟组件共享状态时，把状态提升到共同父组件，通过 props 下发数据 + 回调修改（2.1 单向数据流的直接应用）。

**④ 派生数据用 useMemo**：过滤/排序结果不存 state（双份真相会失同步），渲染时从源数据派生：

```tsx
const visibleTasks = useMemo(
    () => tasks.filter(t => t.status === filter.status),
    [tasks, filter]
);
```

### 阶段三练习

- [ ] 写 `useClock()` Hook：每秒更新当前时间（useEffect + 清理定时器），在组件里渲染 HH:mm:ss
- [ ] 把 3.1 的三 useState 版 DataFetcher 重构成 `FetchState<T>` 可辨识联合版，渲染处用 switch 收窄
- [ ] 写 `useLocalStorage<T>(key: string, initial: T)`：状态持久化到 localStorage（读取用 try/catch + JSON.parse，写入 JSON.stringify）——泛型 Hook 初体验（提示：Hook 本身就是函数，把第一份 5.1 的泛型函数写法套在 Hook 上即可，内部可直接复用 `useState<T>(initial)`）
- [ ] 用 useRef 实现自动聚焦 + "第二次点击才触发"的按钮（ref 记录点击次数，不引发渲染）
- [ ] 做一个 memo 对比实验：父组件含 input + 计数器，子组件不包 memo / 包 memo / memo + useCallback 三种组合，观察 console.log 渲染次数
- [ ] 为阶段二的 TodoList 写 `useTodos()` 自定义 Hook：返回类型接口显式定义，组件瘦身到只负责渲染

### 阶段三过关自检

1. useEffect 依赖数组三种形态的执行时机？清理函数什么时候执行、解决什么问题？
2. FetchState\<T\> 可辨识联合比三个布尔状态好在哪里？
3. useRef 的两种用途？为什么改 ref.current 不触发重渲染？访问 DOM 前必须处理什么？
4. React.memo / useMemo / useCallback 各记忆什么？为什么"传新函数引用"会让 memo 失效？
5. 自定义 Hook 的命名和调用规则？返回类型为什么建议显式定义接口？
6. `Task["status"]` 在 Hook 签名里起什么作用？
7. Record\<TaskStatus, string\> 样式映射比 switch/对象字面量裸写安全在哪？
8. 派生数据为什么不该存进 state？

---

## 阶段四：前端工程化——tsconfig 深解、类型生态、请求层与样式类型

> **目标**：把前端工程的"配置面"吃透：前端 tsconfig 与 Node 版的逐项差异、jsx 五种模式、@types/react 类型生态、路径别名、类型化请求层封装、CSS Modules 与资源的类型声明。学完你能独立配置和诊断一个 Vite + React + TS 项目。

### 4.1 前端 tsconfig 深解 ⭐

Vite react-ts 模板的推荐配置（对照第一份 7.5 的 Node 配置，逐项看差异）：

```jsonc
{
    "compilerOptions": {
        "target": "ES2020",
        "useDefineForClassFields": true,      // 类字段用 defineProperty 语义（现代目标的对齐行为）
        "lib": ["ES2020", "DOM", "DOM.Iterable"],  // ← 关键差异：加 DOM 家族
        "module": "ESNext",                   // ← ESM（Node 版是 commonjs）
        "skipLibCheck": true,

        "moduleResolution": "bundler",        // ← 打包器解析策略（Node 版用 node）
        "allowImportingTsExtensions": true,   // 允许 import 带 .ts 扩展名（配合 noEmit）
        "resolveJsonModule": true,            // 允许 import JSON 文件
        "isolatedModules": true,              // 每文件独立编译（Vite/esbuild 的前提约束）

        "noEmit": true,                       // ← 关键差异：tsc 只做类型检查，不产出 JS
        "jsx": "react-jsx",                   // ← JSX 编译模式

        "strict": true,                       // 与 Node 版一致：严格全家桶
        "noUnusedLocals": true,
        "noUnusedParameters": true,
        "noFallthroughCasesInSwitch": true    // switch 必须 break（防穿透）
    },
    "include": ["src"]
}
```

**jsx 五种模式**（编译器怎么处理 JSX）：

| 模式 | 产物 | 用途 |
|---|---|---|
| `preserve` | 保留 JSX（.jsx） | 交给 Babel 等后续处理 |
| `react` | `React.createElement(...)`（.js） | 经典运行时（需手动 import React） |
| `react-jsx` ⭐ | 自动注入 `jsx(...)` 调用 | **现代默认**：无需 import React |
| `react-jsxdev` | 同上（开发版） | 开发模式 |
| `react-native` | 保留 JSX（.js） | React Native |

**为什么 noEmit**：Vite（esbuild）负责转译和打包，速度远快于 tsc；tsc 退居"纯类型检查器"（`npx tsc --noEmit` 或编辑器实时检查）。这也是 `allowImportingTsExtensions` 存在的原因——反正不产出文件，import 时写全扩展名更明确。

**isolateModules 的连锁影响**（呼应第一份 7.5 的伏笔）：逐文件编译下，`export const x = 42` 被单独看时无法知道它是类型还是值——所以类型重导出必须写 `export type { X }`；const enum 不可用。这是打包器时代的 TS 纪律。

### 4.2 React 类型生态与路径别名 ⭐

**@types/react 全家桶**（兑现阶段一的伏笔）：

- `react` 包自带/配套类型：`@types/react`（核心：React.FC、React.ReactNode、事件类型族、**JSX.IntrinsicElements 的全量注册**——`<div>`、`<button>` 每个标签的属性类型都在这里）、`@types/react-dom`（ReactDOM API）
- Vite 模板已内置安装；老项目手动：`npm i -D @types/react @types/react-dom`
- 排查"某个 JSX 属性是什么类型"的方法：编辑器里 Cmd+点击属性名，跳进 .d.ts 看声明——`onClick?: MouseEventHandler<HTMLButtonElement>` 一目了然

**路径别名（tsconfig 与 Vite 双配置）**——tsconfig 只管类型层，运行时靠 Vite 解析，两边要配平：

```jsonc
// tsconfig.json（类型层）
{
    "compilerOptions": {
        "baseUrl": ".",
        "paths": { "@/*": ["src/*"] }
    }
}
```

```typescript
// vite.config.ts（运行时层）
import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "./src"),
        },
    },
});
```

配置后 `import { useTasks } from "@/hooks/useTasks"` 替代层层 `../../`。

### 4.3 类型化请求层 ⭐

**统一响应类型**（第一份 8.2 的 ApiResponse 在前端复活——前后端共用同一套类型约定）：

```typescript
// src/types/api.ts
export interface ApiResponse<T> {
    success: boolean;
    data?: T;
    error?: string;
}
```

**泛型请求封装**——fetch + 泛型 + unknown 安全链路：

```typescript
// src/api/client.ts
export async function request<T>(url: string, init?: RequestInit): Promise<T> {
    const res = await fetch(url, {
        headers: { "Content-Type": "application/json" },
        ...init,
    });
    const body: unknown = await res.json();          // JSON.parse 产物一律先按 unknown
    if (!res.ok) {
        // 收窄 unknown：可能是我们定义的错误结构，也可能是任何东西
        const message =
            typeof body === "object" && body !== null && "error" in body
                ? String((body as ApiResponse<never>).error)
                : `请求失败（${res.status}）`;
        throw new Error(message);
    }
    return (body as ApiResponse<T>).data as T;       // 断言回实体类型（信任后端契约）
}

// 业务 API：每个端点一个类型化函数
export const taskApi = {
    list:   () => request<Task[]>("/api/tasks"),
    create: (input: CreateTaskInput) =>
        request<Task>("/api/tasks", { method: "POST", body: JSON.stringify(input) }),
};
```

**纪律**：断言 = "信任后端契约"的显式声明；契约不可信时（第三方 API），在 `request` 内部插入运行时校验（见后续深入的 zod）。异步版 useTasks 把 `TaskService.getAll` 换成 `taskApi.list().then(setTasks)` 即接入。

### 4.4 样式与资源的类型

**CSS Modules 的类型声明**——`import styles from "./x.module.css"` 默认无类型报错，两种处理：

```typescript
// 方式 1：Vite 模板自带的 src/vite-env.d.ts 已声明（打开看一眼）：
/// <reference types="vite/client" />
// vite/client 内置了 *.module.css 等资源的类型：styles 的类型是
// { readonly [key: string]: string }——任意键到类名的映射

// 方式 2：需要精确键提示时，自己补声明（第一份 7.4 的 declare module 实战）：
declare module "*.module.css" {
    const classes: { readonly [key: string]: string };
    export default classes;
}
// 更进一步可配合工具生成精确类型（见后续深入）
```

**其他资源导入**：`import logo from "./logo.svg"`——同样由 `vite/client` 声明（字符串 URL）。**className 组合**的工具函数与条件类名：

```tsx
// 模板字符串组合（轻量场景够用）
<div className={`card ${selected ? "card-active" : ""} ${className ?? ""}`}>

// 样式 Props 用字面量联合收窄（variant 模式，2.3 已见）：
// variant?: "primary" | "secondary" | "danger"
```

### 4.5 TS + React 踩坑清单 ⚠

1. **`useState(0)` 之后存了别的东西**：state 类型由初始值锁定，`setCount("3")` 编译错——需要宽类型就显式泛型。
2. **children 忘了进 Props 接口**：组件收了嵌套内容但接口没写 children → 类型不匹配；React.FC 自带 children，解构写法需手写。
3. **事件类型选错**：`onSubmit` 用了 `MouseEvent`、`onChange` 用了 `FormEvent`——属性访问直接报错；内联写法让推断替你选。
4. **对象 props 每次渲染都是新引用**：`<Child config={{ a: 1 }} />` 让 memo 失效——useMemo 固定引用或把配置提出组件外（模块级常量）。
5. **可选链渲染出的假值**：`{count && <Badge/>}` 在 count 为 0 时会渲染出字面 `0`——用 `{count > 0 && ...}`。
6. **`key` 传进组件却没在 Props 里**：key 是 React 保留 prop，不会出现在 props 对象里；需要传标识就另起名（如 `taskId`）。
7. **在 .tsx 里写泛型箭头组件的歧义**：`<T>(props: P<T>) => ...` 的 `<T>` 会被当 JSX 标签——写函数声明 `function List<T>() {}` 或加逗号 `<T,>` 消歧（进阶场景，了解即可）。

### 阶段四练习

- [ ] 新建项目对照 4.1 逐项注释 tsconfig，回答：与 Node 版相比哪 6 项不同、为什么
- [ ] 配置双端路径别名：tsconfig paths + vite.config.ts alias，用 `@/` 导入一个组件验证两处都能解析
- [ ] 实现 4.3 的 request\<T\> + taskApi，用 `npx tsc --noEmit` 验证零错误
- [ ] 给项目加一个 `.module.css` 文件和一张 svg 导入，故意删掉 vite-env.d.ts 的 reference 行，观察报错再恢复
- [ ] 复现踩坑清单第 1、4、5 条（各写一个错误示例 + 修复示例）

### 阶段四过关自检

1. 前端 tsconfig 与 Node 版至少 6 处差异？各自为什么？
2. `jsx: "react-jsx"` 模式下还需要 `import React` 吗？noEmit 意味着 tsc 在项目里还干什么？
3. isolatedModules 对类型导出和 const enum 有什么约束？
4. `JSX.IntrinsicElements` 的注册来自哪个包？怎么查一个 JSX 属性的确切类型？
5. 路径别名为什么必须 tsconfig 和 vite.config.ts 两边都配？各管哪一层？
6. request\<T\> 里 `res.json()` 的结果为什么先声明为 unknown？两处断言分别"信任"了什么？
7. `*.module.css` 的类型声明来自哪里？没有它 import 会发生什么？
8. `useState` 初始值锁死类型后想存联合类型怎么办？内联对象 props 为什么会让 memo 失效，两种解法是什么？

---

## 阶段五：毕业项目——任务管理前端

> 两个版本递进：A 版本地数据（综合运用组件/Hooks/类型设计）；B 版对接第一份路线阶段八的 Express API（前后端合流，两份路线的会师点）。**至少完成 A，强烈建议完成 B**。

### 项目 A：任务管理应用·本地版（1.5 天）

**需求**（与第一份阶段八的后端任务 API 同一套领域模型——B 版无缝衔接）：

1. 类型定义（`src/types/task.ts`）：`TaskStatus`/`TaskPriority` 字面量联合；`Task` 实体；`CreateTaskInput`；`TaskFilter`（status/priority/search）
2. 服务层（`src/services/taskService.ts`）：内存数组 CRUD + 组合过滤（核心思路：`Task[]` 数组 + `find/findIndex` 定位 + `filter` 链式过滤，即 3.4 中 loadTasks 调用的那些方法，自行实现即可；已学过第一份阶段八的读者可直接迁移 8.3 的实现——同一份代码两端复用）
3. 自定义 Hook（`src/hooks/useTasks.ts`）：按 3.4 的完整版实现，含 loading/error/filter
4. 组件树：
   - `App`：布局 + 组合
   - `TaskForm`：受控表单（标题/描述/优先级，2.6 模式）
   - `TaskFilter`：状态/优先级过滤器（下拉 + 搜索框，调用 setFilter）
   - `TaskList` + `TaskItem`：列表渲染 + `Record<TaskStatus, string>` 样式徽章 + 状态切换下拉（`as TaskStatus`）+ 删除按钮
5. 样式：CSS Modules，至少三个模块文件（form/list/item）

**验收**：

- [ ] `npx tsc --noEmit` 零错误，strict 开启，业务代码无 any
- [ ] 全部 Props/Hook 返回值显式接口；事件内联时靠推断、提出时显式标注（两种都出现）
- [ ] 新增一个任务状态 `archived`：编译器强制你补全 Record 映射与状态流转——体验类型系统"推着你走完全部分支"
- [ ] 大列表（造 200 条数据）下用 memo + useCallback 优化 TaskItem，能解释每处优化的必要性

### 项目 B：对接 Express 后端·前后端合流版（0.5~1 天，选做强烈推荐）

**前提**：一个可用的任务 API。首选：启动第一份路线阶段八/毕业项目 A 的 Express 服务（`task-api`，端口 3000）；若尚未完成后端路线，先回头补做阶段八骨架（约半天）即可获得该服务。

**改造步骤**：

1. 请求层：实现 4.3 的 `request<T>` + `taskApi`；Vite 开发代理打通跨域（`vite.config.ts` 里 `server: { proxy: { "/api": "http://localhost:3000" } }`——TS 代码里的 `/api/tasks` 请求会被代理到后端）
2. useTasks 异步化：service 调用替换为 `taskApi.list/create/delete`，操作后 `refresh`；loading/error 用 3.1 的 `FetchState<T>` 状态机建模
3. 错误展示：后端 404/422 错误（第一份的自定义错误体系）在前端渲染为用户可读的提示
4. **类型对齐实验**：把后端 `src/types/index.ts` 与前端 `src/types/task.ts` 并排对比——同一领域模型的两端实现；体会"类型即契约"（进阶：抽成共享包，见后续深入 monorepo）

**验收**：

- [ ] 前端增删改查全部走 HTTP，刷新页面数据仍在（后端持久化）
- [ ] 制造一次后端 404（删不存在的任务），前端展示结构化错误信息而非崩溃
- [ ] 能画出完整链路图：用户点击 → 事件处理器 → taskApi → fetch → Express 路由 → TaskService → 响应 → ApiResponse<T> → setState → 重渲染

### 毕业总结清单

- [ ] 两个项目的代码能向别人讲清：类型分层怎么设计、Hook 边界在哪、为什么这样拆组件
- [ ] 能不看资料写出：Props 接口、事件处理、useState 泛型、useEffect 依赖数组、自定义 Hook 骨架
- [ ] 遇到新类型错误（编辑器红线）能独立读错误信息定位——这是本路线最重要的毕业能力

---

## 学习原则

1. **两遍学习法**：第一遍走完本路线、做完项目——"会用、敢写"；第二遍精读官方 React 文档的 TypeScript 指南与官方 Handbook 的 JSX 章，带着实战问题重读，连接自然生长。
2. **编辑器是主教练**：React + TS 的类型提示极其丰富——悬浮看推断、Cmd+点击看声明（跳进 @types/react 读源码是最好的进阶读物）、故意写错看报错。
3. **类型为组件契约服务**：Props 接口 = 组件的 API 文档；字面量联合 + Record 映射 = 穷尽所有状态；先设计类型再写组件，和后端"先设计 DTO 再写路由"是同一个思维。
4. **两份路线互为参照**：后端版学的泛型/工具类型/可辨识联合，在每个前端章节都有对应战场（ApiResponse ↔ 请求层、Record ↔ 样式映射、状态机 ↔ loading/error）——主动建立这些连接。

## 附录：延伸资料

| 资料 | 定位 |
|---|---|
| React 官方文档 | 组件/Hooks 权威指南（含 TypeScript 使用指南章节） |
| TypeScript 官方 Handbook 的 JSX 章 | JSX 类型系统的机制原理（本路线阶段一的深度版） |
| @types/react 类型定义源码 | 所有事件类型/组件类型的第一手出处 |
| Vite 官方文档 | 构建配置、代理、插件生态 |
| React TypeScript Cheatsheets | 社区维护的 React+TS 速查表（进阶模式大全） |
| Tailwind CSS 官方文档 | 原子化 CSS 方案（本路线 CSS Modules 之后的进阶） |
| zod 官方文档 | 运行时校验：让"信任后端契约"的断言变成可验证 |
| React Router 官方文档 | SPA 路由（多页应用必经之路） |
| Vitest 官方文档 | Vite 原生测试框架（组件测试） |
| Zustand / Redux Toolkit 文档 | 跨组件状态管理（props 传递太深时） |
| TanStack Query 文档 | 服务端状态管理（把 useTasks 这类 Hook 工业化） |

## 后续深入主题（超出本路线范围，学完后按需展开）

- **路由与多页**：React Router（路由参数/嵌套路由/守卫）、Next.js 全栈框架（SSR/SSG）
- **状态管理**：Context + useReducer（中大型应用）、Zustand、Redux Toolkit
- **服务端状态**：TanStack Query（缓存/重试/失效——useTasks 的工业级替代）
- **组件库**：Ant Design / shadcn/ui 的 TS 用法、受控与非受控组件设计、compound components 模式
- **泛型组件**：`function List<T>(props: ListProps<T>)` 与 tsx 里的 `<T,>` 消歧、render props 模式的类型
- **运行时校验**：zod + React Hook Form（表单校验的终极形态：schema 推导 TS 类型，一份定义两端生效）
- **测试**：Vitest + React Testing Library（组件测试）、Playwright（E2E）
- **样式进阶**：Tailwind CSS、CSS-in-JS 的类型权衡
- **性能与工程**：React.lazy + Suspense 代码分割、Profiler 定位渲染瓶颈、pnpm monorepo 共享前后端类型包（毕业项目 B 的延伸）
- **周边生态**：Vue 3 + TS（另一套主流方案，概念互通）、React Native（跨端）
