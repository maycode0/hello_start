# hello_start

> 一套**自包含、零外部链接依赖**的技术学习路线合集——从零基础到工程实践，读完每一份路线不需要中途去查任何外部资料。

## 📚 路线索引

| 文档 | 定位 | 前置 | 总时长 |
|---|---|---|---|
| [`typescript_start/TypeScript学习路线.md`](./typescript_start/TypeScript学习路线.md) | 主线：零基础 → TS 类型系统全量（含类型体操、装饰器）→ Node.js 后端工程实战 | 无（JS 亦零基础） | 约 31 天 × 4~5h |
| [`typescript_start/TypeScript前端学习路线.md`](./typescript_start/TypeScript前端学习路线.md) | 专项：React 方向——DOM/lib.dom、.tsx 规则、组件/Hooks 类型化、前端工程化 | 主线阶段一~七 | 约 10 天 × 4~5h |
| [`golang_start/GO学习路线.md`](./golang_start/GO学习路线.md) | Go 语言：零基础 → 并发编程 → 数据库 / Gin Web 服务与毕业实战 | 无 | 约 6~7 周 × 1~2h |
| [`docker_start/Docker学习路线.md`](./docker_start/Docker学习路线.md) | Docker：macOS 零基础 → Dockerfile / Compose / Swarm 集群运维进阶 | 无 | 约 21 天 × 1.5~2h |

## 🗺️ 路线关系

```
TypeScript 主线（31 天）
├─ 阶段一~七：JS 底座 → 类型系统 → 工程化   ←── 前端专项的前置
│
├─ 后端方向：阶段八~九（Express REST API / CLI / 装饰器路由）
│                          └─ 与前端专项的项目 B 合流：
│                             React 前端 ↔ Express 后端 ↔ ApiResponse<T> 共享类型契约
│
└─ 前端方向：TypeScript 前端专项（10 天，React + Vite）

Go 学习路线 ──────── 独立成线（可与任意 TS 路线并行或先后）
Docker 学习路线 ──── 独立成线（后端方向的天然搭配：容器化部署）
```

**按目标选择**：

- **前端工程师**：TS 主线阶段一~七 → TS 前端专项 →（项目 B 需补主线阶段八）
- **Node.js 后端工程师**：TS 主线全程 → Docker 路线
- **Go 后端工程师**：Go 路线 → Docker 路线
- **全栈**：TS 主线全程 → 前端专项（项目 B 收官）

## ✨ 每份路线的共同特点

- **自包含**：正文零链接、零"参见官网"式指引，所有知识点与代码示例全量内嵌
- **真实素材融合**：知识底座来自菜鸟教程全部相关章节，深度融合官方文档（TypeScript Handbook / Go 官方教程 / Docker 官方文档）——所有素材均实际抓取核实，不凭记忆编造
- **统一阶段结构**：每个阶段 = 目标 → 知识提炼（概念 + 最小可运行代码 + 易错点）→ 练习（勾选清单）→ 过关自检（能回答才进入下一阶段）
- **伏笔与兑现**：前阶段埋"后面会讲"的钩子，后阶段必须真正兑现——保证无"先用后教"
- **毕业项目**：每份路线以 2~3 个融合全部所学的实战项目收尾，至少一个衔接真实工程形态
- **两遍学习法**：第一遍走路线做项目（会用），第二遍精读官方文档（知所以然）
- **一致性审查**：交付前按结构 / 内容 / 声明 / 范围四级清单审查，问题分级（🔴/🟡/🟢）修复留痕

## 📖 阅读约定

| 符号 | 含义 |
|---|---|
| ⭐ | 核心必会 |
| ⚠ | 高频考点 / 易错点 |
| 💡 | 伏笔（后文兑现）或实践建议 |
| `- [ ]` | 练习清单（完成后勾选） |

## 📁 目录结构

```
.
├── typescript_start/
│   ├── TypeScript学习路线.md        # TS 主线（Node.js 后端方向）
│   └── TypeScript前端学习路线.md    # TS 前端专项（React）
├── docker_start/
│   └── Docker学习路线.md            # Docker（macOS 运维进阶）
├── golang_start/
│   └── GO学习路线.md                # Go 语言
└── .pi/skills/learning-path/    # 生成本合集的学习路线技能定义
```

## 🔧 工程规范

- 提交遵循 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/)：`<type>(<scope>): <描述>`（`docs` / `chore` 等，见提交历史）
- 文档大改后同步更新头部时长/版本声明，保持“声明与事实一致”

## 🙏 致谢

本合集的学习路线由 [pi](https://github.com/earendil-works/pi)（AI agent 工具链）与 [GLM](https://github.com/zai-org/GLM-5)（智谱大语言模型）协作生成：素材抓取核实、内容编排与一致性审查由 pi 驱动 GLM 完成，方向决策与最终把关由作者负责。相关提交以 Co-authored-by 尾注署名。

## 📌 内容说明

各路线的版本基线、素材构成、适用读者等详细信息见每份文档的头部声明区。路线内容基于公开教程与官方文档整理提炼，供个人学习使用。
