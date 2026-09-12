# Docker 学习路线（macOS · 零基础 → 运维进阶 · 3 周）

> **总时长**：约 21 天（3 周），每天 1.5~2 小时；若每天可投入 3 小时以上，可压缩至 2 周。
> **版本基线**：按 2025 年的 Docker Desktop for Mac（4.4x）撰写，Docker Engine 27+，Docker Compose 采用 V2 语法（`docker compose` 子命令形式，不带连字符）。以 Apple Silicon（M 系列芯片）Mac 为主环境，Intel Mac 完全兼容。
> **内容构成**：知识底座来自菜鸟教程 Docker 全部 24 个页面（教程、安装、容器/镜像/网络、Dockerfile、Compose、Swarm、命令大全等），并深度融合官方文档核心章节（Docker 概念系列、Dockerfile 构建最佳实践、卷与存储、容器网络、Compose 快速上手、Swarm 教程）。
> **适用读者**：零基础（不需要任何 Linux、运维或编程经验），目标为运维进阶——从安装 Docker 一路学到 Swarm 集群编排，能独立容器化一套完整的 Web 应用并部署到集群。
> **阅读约定**：⭐ 核心必会；⚠ 高频考点/易错点；`$` 开头的行表示在终端里输入的命令，`$` 本身不用输入。

## 全局地图

| 阶段 | 主题 | 定位 | 建议天数 |
|---|---|---|---|
| 一 | 认识 Docker 与 macOS 环境搭建 | 建立概念、装好工具 | 2 天 |
| 二 | 第一个容器与生命周期管理 | 会跑、会管单个容器 | 2 天 |
| 三 | 镜像管理与仓库 ⭐ | 理解镜像分层，会拉会推 | 2 天 |
| 四 | Dockerfile 定制镜像 ⭐⚠ | 核心中的核心 | 3 天 |
| 五 | 数据持久化：卷与绑定挂载 ⭐ | 数据不丢 | 2 天 |
| 六 | 容器网络 ⭐ | 端口映射与容器互联 | 2 天 |
| 七 | 容器化常见服务实战 | Nginx/MySQL/Redis 一把梭 | 1 天 |
| 八 | Docker Compose 多容器编排 ⭐⚠ | 一条命令起全套 | 3 天 |
| 九 | Swarm 集群与 Docker Machine | 运维进阶 | 2 天 |
| 十 | 毕业项目 | 融合全部所学 | 3 天 |

**学习主线**：会用（一~三）→ 理解（四~六）→ 工程化（七~八）→ 集群化（九）→ 实战（十）。

---

## 阶段一：认识 Docker 与 macOS 环境搭建

### 目标

- 说清楚 Docker 是什么、解决什么问题、和虚拟机的区别
- 理解镜像 / 容器 / 仓库三大核心概念和 Docker 的客户端-服务端架构
- 在你的 Mac 上装好 Docker Desktop，跑通第一个容器
- 掌握本路线全程要用的终端命令最小集

### 1.1 Docker 是什么

Docker 是一个开源的**应用容器引擎**，用 Go 语言开发。它可以把你的应用连同全部依赖（库、配置、运行环境）打包成一个轻量级、可移植的**容器**，然后发布到任何装了 Docker 的机器上运行。

它解决的问题，一句话：**"在我机器上能跑，在你机器上跑不了"**。开发、测试、生产环境不一致是运维和开发最经典的痛点，容器把环境本身变成了可分发的软件制品。

Docker 的典型应用场景（也是你后续学习的方向标）：

- **微服务架构**：每个服务独立容器化，便于管理和扩展
- **CI/CD 流水线**：自动化构建、测试、部署，环境完全一致
- **开发环境标准化**：新成员一条命令启动全套依赖（数据库、消息队列），不用再手装 MySQL、Redis
- **云原生基础**：Kubernetes 等编排工具以容器为基本单位管理集群

核心优势：

- **跨平台一致性**：容器在哪跑都一样
- **资源高效**：容器直接共享宿主机内核，不像虚拟机那样虚拟出一整套硬件和操作系统，内存和 CPU 开销极低
- **快速部署**：秒级启动
- **隔离性**：每个容器有独立的文件系统、网络和进程空间，互不干扰

### 1.2 容器 vs 虚拟机 ⚠

这是面试和实际选型的高频问题，必须彻底分清：

| 对比项 | 容器 | 虚拟机 |
|---|---|---|
| 本质 | 一个**隔离的进程** + 运行它所需的全部文件 | 一台**完整的计算机**：自己的内核、驱动、整个操作系统 |
| 内核 | 所有容器**共享宿主机内核** | 每个 VM 有自己的内核 |
| 启动速度 | 秒级 | 分钟级 |
| 体积 | MB 级 | GB 级 |
| 隔离强度 | 进程级（较轻） | 硬件级（更强） |
| 一台机器能跑多少 | 几十上百个 | 几个到十几个 |

一个关键事实：**Docker 容器必须运行在 Linux 内核之上**（它依赖 Linux 内核的 namespace、cgroup 等特性）。macOS 和 Windows 不能原生运行 Linux 容器，所以 Docker 在这两个系统上的做法是：先跑一个隐藏的 Linux 虚拟机，容器实际运行在这个虚拟机里（见 1.4 架构图）。你会看到"容器比虚拟机轻"和"Mac 上的 Docker 里有个虚拟机"同时成立——前者说的是容器 vs 每个应用装一套完整 OS，后者是 Docker Desktop 的实现细节，两者不矛盾。实际生产中也常常见到二者结合：云上的机器本身是虚拟机，一台 VM 上再跑几十个容器，提高资源利用率。

容器的四个重要特性（官方总结）：

- **自包含**：容器带上运行所需的一切，不依赖宿主机预装的东西
- **隔离**：对宿主机和其他容器影响极小
- **独立**：删掉一个容器不影响其他容器
- **可移植**：开发机上怎么跑，数据中心和云上就怎么跑

### 1.3 三大核心概念与整体架构 ⭐

**概念一：镜像——只读模板。** 类比"菜谱"。镜像是一个打包好的只读文件，包含运行应用所需的代码、运行时、库和配置。镜像由**分层**组成，每层代表一组文件系统的增删改（分层机制阶段三、四会深入，这里先记住"镜像像千层饼，一层一层叠出来的"）。

**概念二：容器——镜像的运行实例。** 类比"按菜谱做出的一盘菜"。同一个镜像（菜谱）可以做出很多个容器（菜）。容器有自己的文件系统、进程和网络，是动态的、可启动可销毁的。

**概念三：仓库——存储和分发镜像的地方。** 类比"美食广场"。最著名的是 Docker Hub（公共仓库），企业内部常用 Harbor 等私有仓库。`docker pull` 从仓库下载镜像，`docker push` 把镜像传上去。

**架构：客户端-服务端模式。**

```
你敲的命令            Docker 的真正干活的进程         存取镜像
┌─────────┐  REST API ┌──────────────┐          ┌──────────────┐
│ docker  │ ────────► │ dockerd 守护  │ ◄──────► │ Docker Hub / │
│ 客户端  │  (本地套接字)│ 进程          │  pull/push│ 私有仓库      │
└─────────┘           └──────────────┘          └──────────────┘
```

- **Docker 客户端**（`docker` 命令）只负责把你敲的命令通过 REST API 发给守护进程，它自己不干活
- **Docker 守护进程**接收请求，真正完成构建、运行、分发容器的工作
- Docker Desktop 是"全家桶"：Linux 虚拟机 + Docker Engine（dockerd）+ docker CLI + Docker Compose + 可选的 Kubernetes，外加一个图形界面（GUI）

在 macOS 上的实际结构：

```
macOS（你的 Mac）
 └── Docker Desktop
      └── Linux 虚拟机（对你透明，平时感觉不到）
           └── Docker Engine
                ├── 镜像
                ├── 容器
                └── 卷
```

理解这张图能避免很多新手误解：你看到的容器端口、文件、进程，都是从这个 Linux 虚拟机"映射"到 macOS 上的结果。

### 1.4 在 macOS 上安装 Docker Desktop

**第 1 步：确认机器。** 点左上角苹果菜单 →「关于本机」，看芯片是 Apple（M 系列）还是 Intel。两者都能用 Docker，下载对应的安装包即可。内存建议 8GB 以上（最低 4GB）。

**第 2 步：下载安装。** 到 Docker 官网的 Docker Desktop 下载页，选择 **Mac with Apple chip**（或 **Mac with Intel chip**），下载 `.dmg` 文件。双击打开，把 Docker 鲸鱼图标拖进 Applications 文件夹，即完成安装。

**第 3 步：启动。** 从启动台打开 Docker。首次启动会要求接受服务条款，然后等待状态栏（屏幕右上角）出现鲸鱼图标且不再动画——这表示 Linux 虚拟机和 Docker Engine 已就绪。

**第 4 步：验证。** 打开终端（Terminal，macOS 自带，后面所有命令都在这里敲）：

```
$ docker version
```

能同时看到 `Client:` 和 `Server:` 两段版本信息，说明客户端和守护进程都正常。再跑第一个容器：

```
$ docker run hello-world
```

输出 `Hello from Docker!` 即成功。这个小小的命令背后发生的事：Docker 发现本地没有 `hello-world` 镜像 → 从 Docker Hub 下载 → 用它创建容器 → 容器执行打印动作后退出。第一次跑通它，你就完成了完整的"拉取镜像 → 运行容器"闭环。

**配置国内镜像加速（可选但推荐）。** 国内网络直接从 Docker Hub 拉取镜像可能很慢或失败。Docker Desktop 的配置方式：点击状态栏鲸鱼图标 → **Settings** → **Docker Engine**，在 JSON 配置中加入 `registry-mirrors` 字段：

```json
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com"
  ]
}
```

点 **Apply & Restart**。验证是否生效：

```
$ docker info
```

在输出中看到 `Registry Mirrors:` 下面列出你配置的地址即可。注意：国内镜像源可用性时常变化，某个源拉不动就换一个或删掉这条配置；各大云服务商（阿里云等）也提供专属加速地址。若你的网络可以直连 Docker Hub，则无需配置。

### 1.5 终端命令最小集（零基础必读）

后续所有操作都在终端完成，先把这十几个命令用熟（右上到左下逐个试一遍就是最好的练习）：

| 命令 | 作用 | 示例 |
|---|---|---|
| `pwd` | 显示当前所在目录 | `pwd` |
| `ls` | 列出目录内容 | `ls -l`（详细列表） |
| `cd` | 切换目录 | `cd ~/docker-lab`（~ 代表家目录） |
| `mkdir` | 创建目录 | `mkdir ~/docker-lab` |
| `touch` | 创建空文件 | `touch Dockerfile` |
| `cat` | 查看文件内容 | `cat Dockerfile` |
| `echo` | 输出文本/写文件 | `echo hi > a.txt` |
| `curl` | 发 HTTP 请求 | `curl localhost:8080` |
| `code .` 或 `open .` | 用编辑器/访达打开当前目录 | `open .` |
| `Ctrl + C` | 终止当前前台命令 | — |
| `Tab` 键 | 自动补全命令和路径 | — |
| `↑` / `↓` 方向键 | 翻阅历史命令 | — |

建议现在就执行：

```
$ mkdir -p ~/docker-lab && cd ~/docker-lab
```

以后每学一个新主题，就在这个目录下建一个子目录，保持练习代码有条理。

### 练习

- [ ] 用自己的话向想象中的面试官解释：容器和虚拟机的三个区别
- [ ] 画出（或默写）"macOS → Docker Desktop → Linux 虚拟机 → Docker Engine"的层次结构
- [ ] 安装 Docker Desktop，跑通 `docker run hello-world`
- [ ] 配置镜像加速并用 `docker info` 验证生效
- [ ] 在 `~/docker-lab` 下创建目录，练习 `pwd / ls / cd / mkdir / cat / echo` 各至少一次

### 过关自检

1. Docker 解决的核心问题是什么？它打包的"东西"里包含哪些内容？
2. 容器和虚拟机最本质的区别是什么（提示：内核）？为什么容器启动快、占用小？
3. 镜像、容器、仓库三者的关系，用什么类比记忆？
4. 你敲的 `docker run` 命令，是谁最终执行的？客户端和守护进程是怎么通信的？
5. 为什么 macOS 上运行 Linux 容器需要一个隐藏的 Linux 虚拟机？

---

## 阶段二：第一个容器与生命周期管理

### 目标

- 熟练使用 `docker run` 的常用参数（前台/后台/交互/命名/环境变量）
- 掌握容器从创建到删除的完整生命周期命令
- 学会查看容器日志、进入容器、排查容器问题
- 理解 `exec` 和 `attach` 的关键区别

### 2.1 docker run：一切从它开始 ⭐

`docker run 镜像 命令` 的含义：以某镜像为模板创建一个新容器，并在容器里执行指定命令。若本地没有该镜像，Docker 会自动去 Docker Hub 下载。

**最简形式（跑完即退）：**

```
$ docker run ubuntu /bin/echo "Hello world"
Hello world
```

Docker 用 `ubuntu` 镜像创建容器 → 容器内执行 `/bin/echo "Hello world"` → 输出结果 → 命令结束，容器随之停止（注意：停止 ≠ 删除，后面会讲）。

**交互式容器（进去操作）：**

```
$ docker run -it ubuntu /bin/bash
root@0123ce188bd8:/#
```

- `-i`：保持标准输入打开，让你能向容器输入
- `-t`：为容器分配一个伪终端
- 组合成 `-it`，就得到了一个可以交互的容器内 Shell。提示符变成 `root@容器ID:/#`，说明你现在"在容器里"，是个极简的 Ubuntu 环境

在容器里随便看看（这两条是 Linux 命令，混个脸熟即可）：

```
root@0123ce188bd8:/# cat /proc/version
root@0123ce188bd8:/# ls
```

输入 `exit` 或按 `Ctrl + D` 退出容器。⚠ 注意：对 `-it` 方式启动的容器，退出 Shell 就是退出容器的主进程，容器会**停止**。

**后台运行（生产环境常态）：**

```
$ docker run -d ubuntu /bin/sh -c "while true; do echo hello world; sleep 1; done"
2b1b7a428627c51ab8810d541d759f072b4fc75487eed05812646b8534a2fe63
```

`-d` 让容器在后台运行，终端立刻返回，打印出的长字符串是**容器 ID**（全局唯一）。这个容器每秒打印一次 hello world，但你看不到输出——日志要用 `docker logs` 查看（2.3 节）。

**给容器起名字（强烈建议）：**

```
$ docker run -itd --name ubuntu-test ubuntu /bin/bash
```

`--name` 指定容器名。不给名字时 Docker 会随机生成一个（如 `wizardly_chandrasekhar`），敲起来很痛苦；起个好记的名字，后续所有命令都可以用名字代替容器 ID。

**其他常用参数（先记两个，阶段六、七还有更多）：**

- `-e KEY=VALUE`：给容器注入环境变量，例如 `-e MYSQL_ROOT_PASSWORD=123456`
- `--rm`：容器停止后自动删除它自己，适合跑一次性任务

### 2.2 查看容器：ps 家族 ⭐

```
$ docker ps
```

列出**正在运行**的容器。输出的关键列：

| 列 | 含义 |
|---|---|
| CONTAINER ID | 容器 ID（前 12 位，可只用前几位唯一缩写） |
| IMAGE | 使用的镜像 |
| COMMAND | 容器启动时执行的命令 |
| CREATED / STATUS | 创建时间 / 当前状态 |
| PORTS | 端口映射信息（阶段六详解） |
| NAMES | 容器名 |

⚠ 容器共有 7 种状态：`created`（已创建未启动）、`running`（运行中）、`paused`（暂停）、`restarting`（重启中）、`exited`（已停止）、`removing`（迁移中）、`dead`（死亡）。`docker ps` 只显示 running，想看**所有**容器（包括已停止的）：

```
$ docker ps -a
```

只看最近创建的一个：

```
$ docker ps -l
```

### 2.3 启动、停止、重启、删除

```
$ docker stop ubuntu-test      # 停止容器（给容器发终止信号，温和停止）
$ docker start ubuntu-test     # 启动一个已停止的容器
$ docker restart ubuntu-test   # 重启容器
$ docker rm ubuntu-test        # 删除容器（容器必须是停止状态）
$ docker rm -f ubuntu-test     # 强制删除（运行中也会被先停止再删除）
```

⚠ `docker rm` 删的是**容器**，不是镜像；删除运行中的容器会报错 `You cannot remove a running container`，先 stop 或加 `-f`。

清理所有处于终止状态的容器（批量大扫除）：

```
$ docker container prune
```

### 2.4 看日志与排查问题 ⭐

**看日志**（容器内程序打印到标准输出的内容）：

```
$ docker logs ubuntu-test
$ docker logs -f ubuntu-test     # 像 tail -f 一样持续跟踪新输出
```

**看容器内运行的进程：**

```
$ docker top ubuntu-test
```

**看容器底层详细信息**（返回一大段 JSON，记录配置和状态）：

```
$ docker inspect ubuntu-test
```

**看容器实时资源占用**（类似系统的 top，按 `Ctrl + C` 退出）：

```
$ docker stats
```

排查一个"容器起来又立刻退出"的问题，标准三板斧就是：`docker ps -a` 看状态 → `docker logs` 看报错 → `docker inspect` 看配置。

### 2.5 进入容器：exec vs attach ⚠

后台容器想"进去看看"时，有两条路，区别是致命的：

**方式一（推荐）：exec —— 在容器里新开一个进程**

```
$ docker exec -it ubuntu-test /bin/bash
```

这会在容器里启动一个新的 bash，你进入这个 bash；输入 `exit` 退出时，只是关掉了这个 bash，**容器继续运行**。

**方式二（了解即可）：attach —— 附着到容器的主进程**

```
$ docker attach ubuntu-test
```

它把你直接接到容器主进程（PID 1）的标准输入输出上；此时输入 `exit` 会**连容器一起停止**。

⚠ 记忆口诀：**exec 进去是客，走了家还在；attach 进去是主，走了家就塌**。日常一律用 `docker exec -it`。

### 2.6 容器与镜像的导入导出（快照）

把容器文件系统导出成 tar 包（备份/搬运用）：

```
$ docker export ubuntu-test > ubuntu.tar
```

把 tar 快照导入为一个新镜像：

```
$ cat ubuntu.tar | docker import - test/ubuntu:v1
```

⚠ `export/import` 是**容器快照**，会丢弃历史层和元数据。对应地，阶段三会讲 `save/load`——它导出的是**完整镜像**（含分层信息）。两者别混：搬运用 save/load，简单快照用 export/import。

### 练习

- [ ] 分别用前台、交互、后台三种方式各启动一个 ubuntu 容器，观察提示符和行为差异
- [ ] 后台启动一个循环打印的容器，用 `docker logs -f` 跟踪输出 10 秒后 `Ctrl + C` 退出
- [ ] 用 `--name` 启动容器 `mybox`，完成一轮完整生命周期：stop → ps -a 确认 → start → exec 进入 → exit → 确认容器仍在运行 → rm
- [ ] 故意运行一个会立刻报错的容器（如 `docker run ubuntu /bin/nosuchcmd`），用三板斧排查它为什么退出
- [ ] 把一个容器 export 成 tar，再 import 成镜像 `my/ubuntu:snap`，用 `docker images` 确认

### 过关自检

1. `-d`、`-it`、`--name`、`--rm`、`-e` 各是什么作用？
2. 容器有哪些状态？`docker ps` 和 `docker ps -a` 的区别？
3. `docker rm` 一个运行中的容器会怎样？怎么解决？
4. `exec` 和 `attach` 的区别是什么？为什么日常推荐 exec？
5. 容器里的程序往哪里打印日志才能被 `docker logs` 看到？（提示：标准输出）
6. `export/import` 和 `save/load` 分别作用于什么对象？

---

## 阶段三：镜像管理与仓库 ⭐

### 目标

- 熟练拉取、查看、搜索、删除镜像，理解 `REPOSITORY:TAG` 命名
- 深入理解镜像分层机制（为阶段四构建缓存打地基）
- 掌握镜像完整命名的结构，会打标签、登录、推送镜像到 Docker Hub
- 理解 registry 与 repository 的区别

### 3.1 镜像列表与命名规则

```
$ docker images
REPOSITORY          TAG        IMAGE ID       CREATED        SIZE
ubuntu              latest     a04dc4851cbc   3 weeks ago    78.1MB
nginx               latest     6f8d099c3adc   12 days ago    182.7MB
mysql               8.4        f2e8d6c772c0   3 weeks ago    573MB
hello-world         latest     690ed74de00f   6 months ago   13.3kB
```

- **REPOSITORY**：镜像仓库名（如 `ubuntu`），同一仓库可以有多个版本
- **TAG**：版本标签，`ubuntu:15.10`、`ubuntu:14.04` 是同一仓库的不同版本
- ⚠ 不写 TAG 时默认使用 `latest`：`docker run ubuntu` 等价于 `docker run ubuntu:latest`。生产环境建议始终显式写版本号，`latest` 不保证是稳定版，且不可复现
- **IMAGE ID**：镜像的唯一 ID，删除镜像时可以用它

### 3.2 镜像的完整名字（比你想的长）⚠

`docker pull nginx` 里的 `nginx` 只是简写。完整镜像名的结构是：

```
[HOST[:PORT]/]PATH[:TAG]
```

- `HOST`：镜像仓库服务器地址，不写默认是 `docker.io`（Docker Hub）
- `PATH`：对 Docker Hub 来说是 `[命名空间/]仓库名`；不写命名空间时默认 `library`（官方镜像专用命名空间）
- `TAG`：不写默认 `latest`

几个等价对照：

| 你敲的 | 实际含义 |
|---|---|
| `nginx` | `docker.io/library/nginx:latest`（Docker Hub 官方镜像） |
| `docker/welcome-to-docker` | `docker.io/docker/welcome-to-docker:latest`（docker 组织的镜像） |
| `ghcr.io/dockersamples/app:v1` | GitHub Container Registry 上的镜像 |

**registry vs repository（两个容易混的词）**：registry 是存镜像的**整个服务**（如 Docker Hub、阿里云镜像仓库、公司私有 Harbor）；repository 是 registry 里**一个仓库目录**（如 `library/nginx`），里面装着同一软件的多个版本（多个 tag）。关系：registry（图书馆）→ repository（一个书架）→ 某个 tag 的镜像（架上的一本书）。

选镜像的信任等级（从高到低）：**Docker Official Images**（官方维护、library 命名空间，如 nginx/mysql/redis）> **Verified Publisher**（Docker 认证的商业发布者）> **Docker-Sponsored Open Source**（Docker 赞助的开源项目）> 普通用户镜像（谨慎使用）。

### 3.3 拉取与搜索

```
$ docker pull ubuntu:24.04
24.04: Pulling from library/ubuntu
6599cadaf950: Pull complete
23eda618d451: Pull complete
...
Status: Downloaded newer image for ubuntu:24.04
```

⚠ 注意输出里的一行行 `xxx: Pull complete`——每一行就是**一层**在下载。这是 3.5 节分层机制的直观证据。

搜索镜像（也可以直接上 Docker Hub 网站搜，信息更全）：

```
$ docker search httpd
NAME                 DESCRIPTION                    STARS     OFFICIAL
httpd                The Apache HTTP Server         4600      [OK]
centurylink/httpd    Apache httpd 1.x               25
```

- **OFFICIAL** 为 `[OK]` 表示官方镜像，优先选它
- **STARS** 类似点赞数，反映受欢迎程度

### 3.4 删除镜像与清理

```
$ docker rmi hello-world        # 删除镜像（ IMAGE ID 或 REPOSITORY:TAG 均可）
$ docker rmi -f hello-world     # 强制删除
$ docker image prune -a         # 删除所有没有被容器使用的镜像（危险，先想清楚）
```

⚠ 有容器（哪怕是停止的）基于该镜像时，`rmi` 会报错；先删容器再删镜像。

### 3.5 镜像分层机制 ⭐（阶段四的地基，务必吃透）

**官方定义：镜像是只读的、由多层组成的；每层代表一组文件系统的变更（增加、删除、修改文件）。**

想象一个 Python 应用的镜像，从下往上叠了 5 层：

1. 第 1 层：基础 Linux 发行版（基础命令、apt 包管理器）
2. 第 2 层：安装 Python 运行时和 pip
3. 第 3 层：复制进 requirements.txt
4. 第 4 层：执行 `pip install` 安装依赖
5. 第 5 层：复制进你的源代码

分层的两大好处：

- **复用**：你再做一个 Python 应用，第 1、2 层完全一样，直接共用，不用重新下载和存储——构建更快、磁盘更省、分发带宽更少
- **不可变 + 增量**：镜像一旦构建，每层内容不可更改（immutable），要改就在上面叠新层

**容器运行时发生了什么**：Docker 用联合文件系统把各层"叠"成一个统一视图。运行容器时，在最上面额外加一个**可写层**——容器里改文件、写文件都发生在这层；镜像本身的层永远只读。这就是为什么同一个镜像能同时跑 N 个容器而互不干扰：每个容器有自己的可写层，底下共享同一套只读层。

```
容器 A ──► [可写层 A]
容器 B ──► [可写层 B]        ←── 每个容器独享
              ↓ 共享 ↓
        [层5: 源代码]
        [层4: pip 依赖]
        [层3: requirements]
        [层2: Python 运行时]
        [层1: 基础系统]        ←── 镜像的只读层
```

用命令亲眼验证分层：

```
$ docker image history nginx
```

输出从上到下就是镜像的每一层：什么时候建的、用什么命令创建的、占了多少空间。

### 3.6 commit：把容器变成镜像（救急用）

手头一个容器，改了点东西（比如装了个工具），想固化下来：

```
$ docker run -it --name box ubuntu /bin/bash
root@box:/# apt update && apt install -y curl     # 在容器里做修改
root@box:/# exit

$ docker commit -m "add curl" -a "yourname" box yourname/ubuntu:v2
sha256:70bf1840fd7c...

$ docker run -it yourname/ubuntu:v2 curl --version   # 新镜像里 curl 可用
```

- `-m`：提交说明；`-a`：作者；`box`：源容器；`yourname/ubuntu:v2`：目标镜像名

⚠ **定位要摆正**：commit 是"手工快照"，无法复现（一个月后没人知道这镜像里装了什么）、会塞入大量无关变更。它适合调试救急和临时封装；**正式制作镜像一律用 Dockerfile**（阶段四），这是行业铁律。

### 3.7 给镜像打标签、登录、推送 ⭐

**打标签**（给已有镜像起个别名，常用于发布前规范命名）：

```
$ docker tag ubuntu:24.04 yourname/ubuntu:mytag
```

两个 tag 指向同一个 IMAGE ID——打标签不复制镜像，只是加了个名字。

**登录 Docker Hub**（先去 hub.docker.com 注册免费账号）：

```
$ docker login          # 交互输入用户名密码；或直接在 Docker Desktop 里登录
```

**推送**（镜像名必须是 `你的用户名/仓库名` 格式，否则没权限推）：

```
$ docker push yourname/ubuntu:mytag
```

推完在 Docker Hub 网页上就能看到。`docker pull` 是下载，`docker push` 是上传，`docker login/logout` 管理身份。

### 3.8 镜像的离线搬运：save / load

```
$ docker save -o nginx.tar nginx:latest     # 把镜像（含全部分层）存成 tar 文件
$ docker load -i nginx.tar                  # 在另一台机器上加载回来
```

适用于内网环境、离线交付。和阶段二讲的 `export/import` 对照记忆：**save/load 镜像（完整、带分层）；export/import 容器（快照、丢历史）**。

### 练习

- [ ] 拉取 `ubuntu:24.04`，数一数下载输出里有几层；用 `docker image history` 对照层信息
- [ ] `docker search redis`，找出官方镜像并确认 OFFICIAL 标记
- [ ] 启动容器装一个软件（如 `curl`），用 commit 固化为 `mylab/ubuntu:curl`，再从新镜像起容器验证
- [ ] 给你的镜像再打一个 `mylab/ubuntu:curl-v2` 标签，用 `docker images` 观察两个 tag 的 IMAGE ID
- [ ] 注册 Docker Hub，`docker login` 后把 `你的用户名/test:1.0` 推上去，在网页确认
- [ ] `docker save` 导出你的镜像，`docker rmi` 删掉它，再 `docker load` 恢复

### 过关自检

1. `docker run nginx` 的完整镜像名是什么？`library` 命名空间代表什么？
2. registry、repository、tag 三者的关系？
3. 为什么两个不同的 Python 应用镜像可以共用磁盘空间？分层机制除了省空间还有什么好处？
4. 容器运行时的"可写层"是什么？它和镜像只读层的关系？
5. commit 制作的镜像有什么致命缺陷？为什么生产必须用 Dockerfile？
6. `save/load` 与 `export/import` 的区别？

---

## 阶段四：Dockerfile 定制镜像 ⭐⚠

### 目标

- 独立编写 Dockerfile：掌握全部常用指令及其最佳用法
- 理解构建上下文、构建缓存机制，能优化构建速度
- 掌握多阶段构建，做出小而安全的镜像
- 遵循官方最佳实践写出生产级 Dockerfile

### 4.1 Dockerfile 是什么

Dockerfile 是一个纯文本文件（无扩展名），里面写了一条条**指令**，告诉 Docker 如何一步步构建镜像。对比阶段三的 `docker commit`：commit 是"改完容器拍快照"，Dockerfile 是"把构建过程写成可版本管理、可复现的脚本"——这是制作镜像的正道。

第一个 Dockerfile（定制一个改了首页的 nginx）：

```dockerfile
FROM nginx
RUN echo '这是本地构建的 nginx 镜像' > /usr/share/nginx/html/index.html
```

在 `~/docker-lab/nginx-demo/` 下保存为 `Dockerfile`（注意无扩展名），然后构建：

```
$ docker build -t nginx:v3 .
```

- `-t nginx:v3`：目标镜像名和标签
- 末尾的 `.`：**构建上下文**路径（4.2 节详解）
- 构建过程会逐条执行指令，每条指令生成一层（输出里的 `Step 1/2`、`CACHED` 等）

构建完验证（`-p 8080:80` 先混个脸熟：把宿主机的 8080 端口转发到容器的 80 端口，外部才能访问到容器里的服务，正式详解在阶段六）：

```
$ docker run -d -p 8080:80 nginx:v3
$ curl localhost:8080
这是本地构建的 nginx 镜像
```

### 4.2 构建上下文：那个点是什么 ⚠

回顾架构：docker 客户端（C）和 Docker 引擎（S）是分离的，**构建实际发生在引擎里**，引擎看不到你 Mac 上的文件。所以 `docker build .` 会先把上下文路径（`.` 即当前目录）下的**所有内容打包发给引擎**，Dockerfile 里的 `COPY`、`ADD` 就是从这个包里取文件。

两个直接推论：

1. ⚠ **上下文目录里别放无关文件**（node_modules、.git、大日志等都会被打包传输，拖慢构建）——解法是 `.dockerignore`（4.6 节）
2. Dockerfile 里引用的源路径都是相对于上下文的，构建命令在哪个目录跑、传的哪个路径，要有意识

### 4.3 核心指令详解 ⭐

以下按"写一个典型 Dockerfile 的顺序"讲解。

**FROM —— 指定基础镜像（必须是第一条有效指令）**

```dockerfile
FROM python:3.12-slim
```

你的镜像是在基础镜像之上叠加自己的层。选基础镜像的原则（官方最佳实践）：

- 优先官方镜像（Docker Official Images）
- 尽量小：`alpine` 系（小于 10MB 的完整 Linux 发行版）或 `xx-slim` 系（精简版）；小镜像不仅下载快，攻击面也小
- **固定版本**：写 `python:3.12-slim` 而不是 `python:latest`；更严格可锁定到摘要 `FROM alpine:3.21@sha256:a8560...`（digest 固定，即使发布者更新 tag 你拿到的也永远是同一个镜像，供应链安全）
- 构建时可加 `--pull` 强制拉取最新基础镜像：`docker build --pull -t myapp .`

**WORKDIR —— 设置工作目录**

```dockerfile
WORKDIR /app
```

后续指令（RUN、COPY、CMD）都在这个目录下执行；目录不存在会自动创建。⚠ 用绝对路径；不要写 `RUN cd /app && do-sth` 这种，用 WORKDIR 更清晰可靠。

**COPY —— 把文件复制进镜像**

```dockerfile
COPY requirements.txt ./
COPY src ./src
COPY hom* /mydir/          # 支持通配符
```

从**构建上下文**复制到镜像内；目标目录不存在会自动创建。可加 `--chown=user:group` 指定文件属主。

**ADD —— 加强版 COPY（少用）**

ADD 额外能做两件事：源是本地 **tar 包时自动解压**；能从**远程 URL** 下载文件。⚠ 官方建议：**满足不了 COPY 时才考虑 ADD**——COPY 语义清晰、行为可预期。远程下载更推荐在 RUN 里用 curl/wget（构建缓存更精确）。

**RUN —— 构建时执行命令 ⭐**

```dockerfile
RUN pip install --no-cache-dir -r requirements.txt
```

两种格式：

```dockerfile
RUN apt-get install -y curl                              # shell 格式：交给 /bin/sh -c 执行
RUN ["executable", "param1", "param2"]                   # exec 格式：直接执行，不经 shell
```

⚠ **每条 RUN 生成一层，层的陷阱与合并**。错误示范（3 层、层里残留中间垃圾）：

```dockerfile
RUN yum -y install wget
RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz"
RUN tar -xvf redis.tar.gz
```

正确示范（1 层，`&&` 串联，顺手清理缓存）：

```dockerfile
RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
    && tar -xvf redis.tar.gz \
    && rm redis.tar.gz
```

apt 系（Debian/Ubuntu 基础镜像）的标准写法（官方最佳实践，逐条理解）：

```dockerfile
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*
```

- `apt-get update` 和 `install` **必须在同一条 RUN 里**：分开的话，缓存命中 update 层会导致 install 装到过期的包索引（官方称为缓存陷阱）
- `--no-install-recommends`：不装推荐包，减小体积
- `rm -rf /var/lib/apt/lists/*`：清理 apt 缓存，不留在层里
- 多个包**按字母排序**，好维护、好审查

管道命令注意：`RUN wget -O - url | sh` 只看管道最后一个命令的退出码，前面失败也可能"成功"。加 `set -o pipefail`：

```dockerfile
RUN set -o pipefail && wget -O - https://example.com/install.sh | sh
```

**ENV —— 环境变量（构建时和运行时都在）**

```dockerfile
ENV FLASK_APP=app.py
ENV PATH=/usr/local/nginx/bin:$PATH
```

后续指令可直接用 `$FLASK_APP` 引用；容器运行时这些变量也存在，程序能读到。

**ARG —— 构建参数（只在构建期间存在）**

```dockerfile
ARG VERSION=3.12
FROM python:${VERSION}-slim
```

和 ENV 的区别：ARG 构建完就消失，**不会**留在镜像里被运行时读到；构建时可用 `docker build --build-arg VERSION=3.13 .` 覆盖。适合传版本号这类"只构建需要"的参数；⚠ 绝不要用 ARG 传密码。

**EXPOSE —— 声明端口（文档性质）**

```dockerfile
EXPOSE 80
EXPOSE 443
```

作用：告诉使用者"这个镜像的服务监听这些端口"。⚠ 它**不会**自动发布端口！外部访问仍需 `docker run -p`；`-P`（大写，随机映射）时才会自动映射所有 EXPOSE 的端口。

**VOLUME —— 声明匿名挂载点**

```dockerfile
VOLUME ["/data"]
```

用户忘记 `-v` 挂载时，Docker 会自动把这个路径挂到匿名卷，起保护数据的作用（阶段五展开）。数据库类镜像应给数据目录加 VOLUME。

**USER —— 切换执行用户（安全必做）**

```dockerfile
RUN groupadd -r app && useradd --no-log-init -r -g app app
USER app
```

默认容器以 root 运行，被攻破就是宿主机 VM 里的 root。官方最佳实践：能不用 root 就不用；创建专门的用户再 `USER` 切换过去。

**CMD —— 容器默认启动命令 ⭐⚠**

```dockerfile
CMD ["python", "app.py"]
```

容器启动时默认执行什么。三种格式中**只推荐 exec 数组格式**（`CMD ["可执行文件", "参数"]`），行为最明确。关键规则：

- CMD 在 `docker run` 时执行；RUN 在 `docker build` 时执行（时间点不同）
- 多个 CMD 只有**最后一个**生效
- ⚠ `docker run 镜像 附加命令` 会**整体覆盖** CMD：`docker run myapp /bin/bash` 时 app.py 根本不会跑

**ENTRYPOINT —— 入口点（不容易被覆盖）⭐⚠**

```dockerfile
ENTRYPOINT ["nginx", "-c"]
CMD ["/etc/nginx/nginx.conf"]
```

与 CMD 的配合是最经典的考点：ENTRYPOINT 定"执行什么程序"（定参），CMD 给"默认参数"（变参）：

| 命令 | 实际执行 |
|---|---|
| `docker run nginx:test` | `nginx -c /etc/nginx/nginx.conf`（用默认 CMD） |
| `docker run nginx:test -c /etc/nginx/new.conf` | `nginx -c /etc/nginx/new.conf`（CMD 被替换） |

- `docker run` 的命令行参数会**追加给 ENTRYPOINT**（而不是覆盖它）
- 想强行覆盖 ENTRYPOINT 本身：`docker run --entrypoint 其他程序 镜像`
- 多个 ENTRYPOINT 只有最后一个生效
- 官方典型用法：工具类镜像 `ENTRYPOINT ["s3cmd"]` + `CMD ["--help"]`——不传参显示帮助，传参直接执行
- 服务类镜像的通用模式：入口点脚本（如官方 postgres 的 `docker-entrypoint.sh`）负责初始化（建用户、初始化数据目录），最后 `exec` 真正的服务进程，使其成为 PID 1 正确接收信号

**LABEL —— 元数据（替代已弃用的 MAINTAINER）**

```dockerfile
LABEL org.opencontainers.image.authors="you@example.com"
LABEL version="1.0" description="练习镜像"
```

**HEALTHCHECK —— 健康检查**

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD curl -f http://localhost/ || exit 1
```

Docker 定期执行该命令判断容器健康状态（`docker ps` 的 STATUS 列会显示 healthy/unhealthy）。`HEALTHCHECK NONE` 可屏蔽基础镜像自带的检查。

**ONBUILD —— 留给下一代的指令**

本镜像构建时**不执行**；当别的 Dockerfile `FROM` 它时才触发。用于制作"语言基础镜像"（如 ruby:onbuild 自动 COPY 用户代码）。初学阶段了解即可。

**SHELL / STOPSIGNAL**：前者改变 RUN/CMD 的默认 shell（Windows 镜像常用）；后者改变停止容器时发送的信号（默认 SIGTERM）。了解即可。

### 4.4 指令速查表

| 指令 | 作用 | 易错点 |
|---|---|---|
| FROM | 基础镜像 | 要固定版本；越小越好 |
| WORKDIR | 工作目录 | 用绝对路径 |
| COPY | 复制文件进镜像 | 优先用它而不是 ADD |
| ADD | 复制+解压/下载 | 仅 tar 解压或远程 URL 场景 |
| RUN | 构建时执行命令 | 合并成一层；apt 要 update+install 同层 |
| ENV | 环境变量（构建+运行） | 敏感信息别放这 |
| ARG | 构建参数 | 构建后消失；别传密码 |
| EXPOSE | 声明端口 | 不等于发布端口 |
| VOLUME | 匿名卷挂载点 | 防忘挂载数据丢失 |
| USER | 切换用户 | 用户需先创建；尽量非 root |
| CMD | 默认启动命令+默认参数 | 可被 run 参数覆盖；仅最后一个生效 |
| ENTRYPOINT | 主命令入口 | 参数追加而非覆盖；配 CMD 用 |
| LABEL | 元数据 | 替代 MAINTAINER |
| HEALTHCHECK | 健康检查 | 无状态服务可不配 |
| ONBUILD | 给子镜像的钩子 | 制作基础镜像时用 |

### 4.5 构建缓存：为什么你的构建忽快忽慢 ⭐⚠

构建时 Docker 逐条执行指令，**每条都先问：这条之前执行过且输入没变吗？** 是 → 直接复用缓存层（输出 `CACHED`）；否 → 执行并重建，**且其后所有层全部失效重建**。

缓存失效的触发条件：

1. RUN 的命令文本有任何改动
2. COPY/ADD 涉及的文件内容或属性（权限等）变了
3. 前面任何一层失效，后面全部连锁失效

推论（官方核心示例，务必亲手体会）：下面这个 Dockerfile，任何一行代码改动都会让 `yarn install` 重新执行——因为 `COPY . .` 把全部代码拷进来，代码一变这层就失效，后面跟着失效：

```dockerfile
FROM node:22-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "./src/index.js"]
```

优化：**变化少的放前面，变化多的放后面**；先只拷依赖清单，装完依赖，再拷代码：

```dockerfile
FROM node:22-alpine
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --production
COPY . .
CMD ["node", "./src/index.js"]
```

这样改业务代码时，前 4 步全命中缓存，只需重拷代码——构建从 20 秒降到 1 秒。**Python 项目同理**：先 `COPY requirements.txt` → `pip install` → 再 `COPY . .`。这个"依赖先行"模式是写 Dockerfile 的肌肉记忆。

两个构建旗标：

- `docker build --no-cache -t app .`：忽略全部缓存，从头重建（怀疑缓存有问题、要拉最新依赖时用）
- `docker build --pull -t app .`：强制检查基础镜像更新（和 --no-cache 是两回事，可组合）

### 4.6 .dockerignore

放在构建上下文根目录，语法类似 .gitignore，排除不需要发给引擎的文件：

```
node_modules
*.log
.git
.env
Dockerfile
```

作用：上下文更小、构建更快、（重要）防止 `.env` 这类密钥文件被打进镜像。Python 项目至少排除 `__pycache__`、`.git`、`node_modules`。

### 4.7 多阶段构建 ⭐

问题：编译型/带构建工具链的应用，构建环境（JDK、编译器、node_modules）会被全部打进最终镜像——官方实测一个 Spring Boot 应用单阶段 880MB。多阶段构建让你"用一个肥镜像编译，只把产物搬进一个瘦镜像"：

```dockerfile
# ---- 第一阶段：构建环境，起名 builder ----
FROM eclipse-temurin:21-jdk-jammy AS builder
WORKDIR /opt/app
COPY .mvn/ .mvn
COPY mvnw pom.xml ./
RUN ./mvnw dependency:go-offline
COPY ./src ./src
RUN ./mvnw clean install          # 在这里编译出 JAR

# ---- 第二阶段：精简运行时 ----
FROM eclipse-temurin:21-jre-jammy AS final
WORKDIR /opt/app
EXPOSE 8080
COPY --from=builder /opt/app/target/*.jar /opt/app/app.jar   # 只搬产物！
ENTRYPOINT ["java", "-jar", "/opt/app/app.jar"]
```

要点：

- 多个 `FROM` 把构建分成多个阶段，`AS 名字` 给阶段命名
- `COPY --from=builder` 从别的阶段复制文件（也能 `COPY --from=python:3.12 ...` 从镜像复制）
- 最终镜像默认只构建**最后一个阶段**；`docker build --target builder .` 可单独构建中间阶段（调试用）
- 上例最终镜像 428MB，几乎砍半；解释型语言同样适用：一个阶段构建/压缩前端产物，一个阶段用 nginx 只托管静态文件

### 4.8 官方最佳实践清单（写 Dockerfile 前默诵一遍）

1. 使用多阶段构建，构建环境与运行环境分离
2. 选小而可信的基础镜像（官方镜像、alpine/slim 变体）
3. 固定基础镜像版本（tag 或 digest），用 `--pull` 定期获取更新，在 CI 中自动重建
4. 用 `.dockerignore` 排除无关文件
5. 容器应当"用完即弃"（ephemeral）：随时可停、可删、可重建
6. 一个容器只做一件事；应用拆成 web/数据库/缓存多个容器（阶段八 Compose 的哲学基础）
7. 不装不必要的包（少一个包就少一分体积和漏洞）
8. 依赖先行排列指令，充分利用构建缓存
9. RUN 命令合并、清理包缓存、多行参数排序
10. 以非 root 用户运行
11. 敏感信息不进镜像（用运行时环境变量/挂载传入）

### 练习

- [ ] 写 Dockerfile 定制 nginx 首页，构建为 `nginx:mine` 并运行验证
- [ ] 用 Ubuntu 基础镜像写一个含以下内容的镜像：工作目录 /app、环境变量 `APP_MODE=lab`、装好 curl、非 root 用户运行、默认命令打印环境变量。提示：默认命令可用 `CMD ["sh", "-c", "env"]`
- [ ] 故意写出"代码一改依赖就重装"的 Dockerfile，再用"依赖先行"重写，分别构建两次对比 CACHED 行数
- [ ] 给上面的项目加 .dockerignore，验证上下文传输体积变小
- [ ] 写一个 ENTRYPOINT+CMD 组合的镜像：`ENTRYPOINT ["echo"]`、`CMD ["hello"]`，分别用不带参数和带参数 `world` 两种方式 run，观察输出
- [ ] 挑战：为下面这个最小 Python 应用写生产级 Dockerfile（提示：python:3.12-slim 基础镜像、先拷 requirements.txt、非 root、exec 格式 CMD）：

```python
# app.py —— 一个极简 Web 服务，先不用管语法，照抄即可
from http.server import HTTPServer, BaseHTTPRequestHandler

class H(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b"hello from my image")

HTTPServer(("0.0.0.0", 8000), H).serve_forever()
```

requirements.txt 内容为空文件即可。

### 过关自检

1. `docker build -t app .` 最后那个点的含义？为什么上下文里不要放无关文件？
2. CMD 和 RUN 的执行时间点分别是什么？CMD 被什么覆盖？
3. ENTRYPOINT 和 CMD 如何配合？`docker run 镜像 参数` 时参数给了谁？
4. 构建缓存何时失效？"依赖先行"为什么能加速构建？
5. COPY 和 ADD 怎么选？EXPOSE 会不会自动发布端口？
6. 多阶段构建解决什么问题？`COPY --from` 从哪复制？
7. 说出至少 5 条官方 Dockerfile 最佳实践。

---

## 阶段五：数据持久化——卷与绑定挂载 ⭐

### 目标

- 理解容器文件系统的"临时性"，知道哪些数据需要持久化
- 熟练使用命名卷：创建、挂载、查看、备份、删除
- 熟练使用绑定挂载共享宿主机文件
- 能正确在"卷 vs 绑定挂载"之间做选择

### 5.1 为什么需要持久化

容器 = 镜像只读层 + 自己的可写层。你在容器里写的所有文件都落在可写层，**容器一删，可写层随之销毁**，数据全没。数据库容器重启变成空库，这是不可接受的。

Docker 提供两大持久化机制：

- **卷**：Docker 全权管理的存储，存在于容器之外（存在 Linux VM 里由 Docker 管理的目录），容器删了它还在
- **绑定挂载**：把宿主机上的一个目录/文件直接"接"进容器，容器内外实时同步

### 5.2 卷（Volume）⭐

**创建与挂载：**

```
$ docker volume create mydata
$ docker run -d --name web -p 8080:80 -v mydata:/usr/share/nginx/html nginx
```

`-v 卷名:容器内路径`。即使卷不存在，`docker run` 也会自动创建。此后容器往 `/usr/share/nginx/html` 写的任何内容实际存储在卷里；删掉容器、换新容器再挂同一个卷，数据都在。

**管理命令：**

```
$ docker volume ls                 # 列出所有卷
$ docker volume inspect mydata     # 查看卷详情（存储路径等）
$ docker volume rm mydata          # 删除卷（须无容器使用）
$ docker volume prune              # 删除所有未使用的卷（危险操作）
```

⚠ 卷有独立于容器的生命周期；`docker rm` 容器**不会**删卷。匿名卷（不写卷名，如 `-v /data`，或 Dockerfile 里的 `VOLUME` 指令自动创建）也会一直残留——`docker run --rm` 时匿名卷会随容器删除，命名卷不会。定期 `docker volume ls` 清点，用 `prune` 大扫除。

**两种挂载语法：**

```
$ docker run -v mydata:/data nginx                       # -v：简短常用
$ docker run --mount type=volume,src=mydata,dst=/data nginx   # --mount：更显式
```

`--mount` 支持全部高级选项（只读、子路径、卷驱动），Docker 官方更推荐它；日常简单场景 `-v` 足够。加只读：

```
$ docker run -v mydata:/data:ro nginx
$ docker run --mount type=volume,src=mydata,dst=/data,readonly nginx
```

**空卷的"预填充"行为（冷知识但有用）**：把一个**空**卷挂到容器内**已有文件**的目录上时，Docker 会先把该目录原有内容**复制进卷**（例如给 nginx 镜像的 html 目录挂新卷，会自动获得默认页面）；挂**非空**卷时，卷内容会遮住目录原内容（类似 U 盘插到 /mnt 上）。

### 5.3 绑定挂载⭐

把 Mac 上的目录直接映射进容器：

```
$ mkdir -p ~/docker-lab/site
$ echo '<h1>Hi Docker</h1>' > ~/docker-lab/site/index.html
$ docker run -d --name web -p 8080:80 -v ~/docker-lab/site:/usr/local/apache2/htdocs httpd:2.4
$ curl localhost:8080
<h1>Hi Docker</h1>
```

现在直接用编辑器改 Mac 上的 `index.html`，刷新浏览器立即生效——因为容器看到的就是这个目录本身。这就是**开发环境代码热更新**的基础（阶段八 Compose watch 会用到）。

`--mount` 等价写法：

```
$ docker run -d --name web -p 8080:80 \
  --mount type=bind,source=$HOME/docker-lab/site,target=/usr/local/apache2/htdocs \
  httpd:2.4
```

⚠ 两种语法一个重要差异：宿主路径不存在时，`-v` 会自动创建一个空目录，`--mount` 直接报错——后者更能暴露手误，这也是官方推荐 --mount 的原因之一。

权限后缀：`:ro` 只读（容器改不了宿主文件，防误删）、`:rw` 可读写（默认）。

### 5.4 卷 vs 绑定挂载怎么选 ⚠

| 维度 | 卷 | 绑定挂载 |
|---|---|---|
| 由谁管理 | Docker 引擎（存在它自己的地盘） | 你（普通宿主目录） |
| 典型用途 | 数据库文件、应用数据、日志 | 开发时共享源码/配置文件 |
| 备份迁移 | 容易（Docker 管理，不受宿主目录结构影响） | 依赖宿主目录结构 |
| 跨平台 | 好（路径由 Docker 处理） | 受宿主路径差异影响 |
| 性能 | 稳定 | macOS/Windows 上有 VM 中转，大目录可能慢 |

口诀：**存数据用卷，传代码用挂载**。数据库用绑定挂载是常见的反模式（权限、性能、路径耦合都会踩坑）。

### 5.5 数据库持久化实战（MySQL）

经典组合应用——阶段三到五的总复习：

```
$ docker run -d --name mysql-lab \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=lab123456 \
  -v mysql-data:/var/lib/mysql \
  mysql:8.4
```

- `-e MYSQL_ROOT_PASSWORD`：MySQL 官方镜像要求通过此变量设置 root 密码
- `-v mysql-data:/var/lib/mysql`：MySQL 的数据目录挂到命名卷

验证数据真的持久化了：

```
$ docker exec -it mysql-lab mysql -uroot -plab123456 -e "CREATE DATABASE demo;"
$ docker rm -f mysql-lab                    # 删容器
$ docker run -d --name mysql-lab2 -v mysql-data:/var/lib/mysql mysql:8.4
$ docker exec -it mysql-lab2 mysql -uroot -plab123456 -e "SHOW DATABASES;"
```

注意第二次启动**没传密码**——密码变量只在首次初始化空库时使用，数据卷里已有初始化好的数据库。`SHOW DATABASES` 能看到 `demo`，证明数据跨容器存活。（想从 Mac 直连：宿主机上装个客户端连 `127.0.0.1:3306` 即可。）

### 5.6 卷的备份与恢复

卷不能直接用 `docker cp`（它在 Docker 管理的目录里），标准做法是"借一个临时容器挂载卷 + tar 打包"（`$PWD` 是值为当前目录绝对路径的环境变量，相当于阶段一 `pwd` 命令的变量版）：

```
$ docker run --rm \
  -v mysql-data:/data:ro \
  -v $PWD:/backup \
  alpine tar czf /backup/mysql-data-backup.tar.gz -C /data .
```

这条命令挂了两个东西：要备份的卷（只读）和当前目录；用 alpine 迷你镜像里的 tar 把卷内容打成包落到 Mac 当前目录。恢复则反向：

```
$ docker volume create mysql-restored
$ docker run --rm \
  -v mysql-restored:/data \
  -v $PWD:/backup:ro \
  alpine tar xzf /backup/mysql-data-backup.tar.gz -C /data
```

### 练习

- [ ] 创建卷 `lab-vol`，用 `-v` 和 `--mount` 两种方式各挂载一次，`docker volume inspect` 对比
- [ ] 起一个 nginx 容器把卷挂到 html 目录，进容器改 index.html，删容器后挂同一卷起新容器，验证内容还在
- [ ] 用绑定挂载把一个本地网页目录共享给 nginx/httpd，修改宿主文件验证实时生效
- [ ] 完成 5.5 的 MySQL 持久化全流程（建库 → 删容器 → 恢复验证）
- [ ] 给卷挂载加 `:ro`，进容器尝试写文件，观察报错
- [ ] 用临时容器把 `mysql-data` 备份成 tar 包，再恢复到新卷

### 过关自检

1. 容器删除后，可写层的数据去哪了？卷为什么能活下来？
2. `-v` 三段式 `name:/path:ro` 各段含义？`--mount` 等价写法是什么？
3. 命名卷和匿名卷的区别？`--rm` 对它们分别做什么？
4. 数据库该用卷还是绑定挂载？为什么？
5. 把空卷挂到容器内有文件的目录，会发生什么？
6. 怎么备份一个卷？为什么不能直接 cp？

---

## 阶段六：容器网络 ⭐

### 目标

- 掌握端口发布的所有姿势：`-p`、`-P`、绑定指定 IP、UDP
- 理解默认桥接网络 vs 自定义网络的关键差异（DNS 解析）
- 能创建自定义网络并让多个容器互联、按名互访
- 了解 Docker 网络驱动家族

### 6.1 容器网络的默认行为

容器默认就能上网（拉包、ping 外网都行，通过宿主机网络做地址伪装转发），但**外部默认访问不到容器里的服务**——隔离是双向的。让外部能访问，靠**发布端口**；让容器之间互相访问，靠**网络**。

### 6.2 发布端口：-p 的全部姿势 ⭐⚠

发布端口 = 在宿主机上开一个口子，把流量转发进容器，语法骨架 `-p 宿主端口:容器端口`：

```
$ docker run -d -p 8080:80 nginx          # 宿主 8080 → 容器 80
$ docker run -d -p 127.0.0.1:8080:80 nginx  # 只绑定本机回环，外部机器访问不到（三段式）
$ docker run -d -p 8080:80/udp some-app   # 发布 UDP 端口（默认是 TCP）
$ docker run -d -p 80 nginx               # 省略宿主端口：Docker 随机挑一个高位端口
$ docker run -d -P nginx                  # 大写 P：发布镜像 EXPOSE 的所有端口（随机宿主端口）
```

⚠ 两个易错点：

1. `-p` 发布后**默认监听所有网卡**（`0.0.0.0`），意味着局域网内其他机器都能访问。数据库这类敏感服务要养成 `-p 127.0.0.1:3306:3306` 只绑本机的习惯
2. `-P`（大写）依赖镜像里的 EXPOSE 声明；`-p`（小写）不需要 EXPOSE 也能发布

查看映射结果：

```
$ docker ps                    # PORTS 列：0.0.0.0:8080->80/tcp
$ docker port 容器名 80        # 快捷查看某个容器端口的映射：80/tcp -> 0.0.0.0:8080
```

宿主端口冲突（两个容器都想占 8080）会直接报错，换一个宿主端口即可，容器内端口不用动——这也是"一机多实例"的标准玩法（两个 MySQL 容器分别映射 3306、3307）。

### 6.3 网络驱动一览

Docker 内置多种网络驱动，用 `docker network create -d 驱动名` 指定：

| 驱动 | 用途 |
|---|---|
| `bridge`（默认） | 桥接网络：容器经 NAT 上网，最常用 |
| `host` | 直接共用宿主机网络栈，无隔离（端口不用映射，但失去隔离） |
| `none` | 完全无网络（最严格隔离） |
| `overlay` | 跨多台主机的容器互联（Swarm 集群用，阶段九见） |
| `ipvlan` / `macvlan` | 容器直接出现在物理网络里，有自己的 MAC/IP（进阶） |

零基础阶段重点掌握 bridge 即可，host/none 知道存在和用途，overlay 留给阶段九。

### 6.4 默认桥 vs 自定义网络 ⚠（本阶段最重要的知识点）

所有不带 `--network` 的容器都连到**默认桥接网络**（`docker network ls` 里的 `bridge`）。它有个大缺陷：**容器之间只能用 IP 互访，不能用名字**——而容器的 IP 每次重建都可能变。

自定义网络解决了这个问题：**同一自定义网络里的容器，互相可以用容器名当主机名访问**（内置 DNS）。这也是官方明确不建议生产用默认桥的原因之一；另外默认桥上"所有容器挤一张网"，隔离性差，自定义网络则按需组网。

```
$ docker network create -d bridge lab-net      # 创建自定义桥接网络
$ docker network ls                            # 查看：bridge / host / none / lab-net ...
$ docker run -itd --name c1 --network lab-net ubuntu
$ docker run -itd --name c2 --network lab-net ubuntu
```

验证 c1 能按名字 ping 通 c2（ubuntu 精简镜像没有 ping，先装）：

```
$ docker exec -it c1 bash
root@c1:/# apt update && apt install -y iputils-ping
root@c1:/# ping c2
64 bytes from 172.18.0.3: icmp_seq=1 ttl=64 time=0.2 ms
```

`ping c2` 能通——c2 的名字被网络内置 DNS 解析成了它的 IP。把 c2 删了重建，IP 变了，`ping c2` 依然通。这就是容器互联的正确姿势，也是阶段八 Compose 里 `redis`、`db` 这种服务名能当主机名用的原理。

网络管理命令全家福：

```
$ docker network create lab-net               # 创建
$ docker network ls                           # 列表
$ docker network inspect lab-net              # 详情（看哪些容器在里面、子网、网关）
$ docker network connect lab-net 容器名        # 把运行中的容器再接入一个网络
$ docker network disconnect lab-net 容器名     # 断开
$ docker network rm lab-net                   # 删除（须无容器连接）
```

### 6.5 容器 DNS 配置

- 单个容器定制：`--hostname 主机名`（改容器内 hostname）、`--dns=8.8.8.8`（指定 DNS 服务器）、`--dns-search=example.com`（搜索域）
- 全局定制（Linux 引擎场景）：编辑 daemon.json 的 `"dns": [...]` 后重启（macOS 上在 Docker Desktop 的 Settings → Docker Engine 里改）

排查容器网络问题的顺序：`docker inspect` 看 IP 和网络 → 进容器 ping/curl 目标名 → 检查双方是否在同一个自定义网络。

### 练习

- [ ] 起一个 nginx 分别用 `-p 8080:80`、`-p 127.0.0.1:8081:80`、`-p 8082:80/udp` 发布，`docker ps` 和 `docker port` 观察差异
- [ ] 用 `-p` 随机端口发布一个 nginx，用 `docker port` 找到端口号并访问成功
- [ ] 创建网络 `lab-net`，起 c1、c2 两个容器加入，互 ping 容器名验证 DNS
- [ ] 把 c2 删除重建（同名同网络），在 c1 里再 ping c2，体会"IP 会变名字不变"
- [ ] 起第三个容器**不**加 `--network`，在 c1 里 ping 它的名字，验证默认桥上名字不通
- [ ] `docker network inspect lab-net`，找到两个容器的 IP

### 过关自检

1. `-p 127.0.0.1:8080:80` 三段各是什么？不写 IP 默认监听什么？有什么安全隐患？
2. `-p` 和 `-P` 的区别？`-P` 依赖镜像里的什么？
3. 默认桥和自定义网络的两大差异？
4. 为什么容器互联推荐"容器名"而不是 IP？
5. overlay 驱动是为哪种场景准备的？
6. `--hostname` 和 `--dns` 分别修改容器的什么？想给所有容器统一配 DNS 应该改哪里？
7. 两个容器互相 ping 不通，你的排查步骤是什么？

---

## 阶段七：容器化常见服务实战

### 目标

- 掌握"容器化一个现成服务"的通用套路：搜 → 拉 → 跑 → 验证
- 独立部署 Nginx、MySQL、Redis 并能验证服务可用
- 会用资源限制参数约束容器，用 stats 监控

### 7.1 通用四步套路

任何官方镜像服务，都是同一套流程：

1. **搜**：`docker search nginx`（或 Docker Hub 网页），认准 OFFICIAL
2. **拉**：`docker pull nginx:1.27`（固定版本）
3. **跑**：`docker run -d --name 名字 -p 宿主:容器 [-e 变量] [-v 卷:路径] 镜像`
4. **验**：`docker ps` 看状态 → `curl` 或 `docker exec` 进去连服务

背后的原则（官方最佳实践）：**一个容器只做一件事，并把它做好**。Web、数据库、缓存各自独立容器，而不是全塞一个容器里——独立扩缩、独立更新、独立排障。

### 7.2 Nginx ⭐

```
$ docker run --name nginx-test -p 8080:80 -d nginx:1.27
$ curl localhost:8080       # 看到 Welcome to nginx! 页面源码
```

进阶玩法（结合阶段五、六）：挂载自定义配置和静态目录——

```
$ docker run --name nginx-site \
  -p 8080:80 \
  -v $PWD/nginx.conf:/etc/nginx/conf.d/default.conf:ro \
  -v $PWD/html:/usr/share/nginx/html:ro \
  -d nginx:1.27
```

改 Mac 上的 `html/index.html`，站点即时更新。

### 7.3 MySQL ⭐

阶段五 5.5 已完整做过，核心参数回顾：

```
$ docker run -itd --name mysql-test \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -v mysql-data:/var/lib/mysql \
  mysql:8.4
```

进入容器内客户端验证：

```
$ docker exec -it mysql-test mysql -uroot -p123456
mysql> SHOW DATABASES;
```

常用环境变量：`MYSQL_ROOT_PASSWORD`（必填其一）、`MYSQL_DATABASE`（自动建库）、`MYSQL_USER`/`MYSQL_PASSWORD`（自动建用户）。

### 7.4 Redis ⭐

```
$ docker run -itd --name redis-test -p 6379:6379 redis:7.4
```

用容器内 redis-cli 验证（set 一个键再读出来）：

```
$ docker exec -it redis-test redis-cli
127.0.0.1:6379> SET docker lab
OK
127.0.0.1:6379> GET docker
"lab"
127.0.0.1:6379> exit
```

数据持久化：Redis 默认内存库，重启丢数据；要持久化就挂卷到 `/data`（配置 `--appendonly yes` 开 AOF）：

```
$ docker run -itd --name redis-test -p 6379:6379 -v redis-data:/data redis:7.4 --appendonly yes
```

注意镜像名后面直接跟的是**传给 redis-server 的参数**——这正是阶段四 ENTRYPOINT 设计的活例子：redis 镜像 ENTRYPOINT 是 `redis-server`，你追加的参数成为它的启动参数。

### 7.5 其他常见服务速查

套路完全相同，只列关键命令：

```
# Tomcat（8080 端口）
$ docker run -itd --name tomcat-test -p 8888:8080 tomcat:10

# MongoDB（27017）
$ docker run -d --name mongo-test -p 27017:27017 -v mongo-data:/data/db mongo:7

# Apache httpd（80）
$ docker run -p 8081:80 -v $PWD/www/:/usr/local/apache2/htdocs/ -d httpd:2.4

# Python（跑一次性脚本，绑定挂载代码目录）
$ docker run --rm -v $PWD/myapp:/usr/src/myapp -w /usr/src/myapp python:3.12 python hello.py
```

（`-w` 是 `docker run` 的工作目录参数，效果类似 Dockerfile 的 WORKDIR。）

### 7.6 资源限制与监控

防止某个容器吃光宿主机资源（官方推荐做法）：

```
$ docker run -d --name limited-app \
  --memory="512m" \
  --cpus="0.5" \
  nginx:1.27
```

- `--memory`：内存上限（超出可能被 OOM 杀掉）
- `--cpus`：CPU 配额（0.5 = 半个核）

实时监控所有容器的 CPU/内存/网络 IO：

```
$ docker stats
```

结合 `docker top`（进程）和 `docker inspect`（配置），就构成了容器运维的日常观测三板斧。运行中容器还能用 `docker update --memory=1g 容器名` 动态调整限制。

### 练习

- [ ] 用四步套路部署 Tomcat 并访问到它的默认页面
- [ ] 部署 MySQL（带自动建库 `MYSQL_DATABASE=demo`）并用 exec + mysql 客户端验证
- [ ] 部署 Redis：set/get 验证；加上数据卷和 AOF，删除容器重建后验证键还在
- [ ] 用绑定挂载方式让 nginx 提供你自己的静态页面
- [ ] 起一个 `--memory=256m --cpus=0.25` 的容器，用 `docker stats` 观察限额
- [ ] 一台机器上同时跑两个 MySQL 实例（提示：不同宿主端口）

### 过关自检

1. 容器化一个官方镜像服务的四步是什么？
2. MySQL 容器必传的环境变量是哪个？第二次启动为什么不用传？
3. `docker run redis ... --appendonly yes` 末尾的参数是怎么生效的（哪个指令的功劳）？
4. 为什么一个容器只跑一个服务？
5. 怎么限制容器只用 1GB 内存、1 个核？怎么实时看资源占用？

---

## 阶段八：Docker Compose 多容器编排 ⭐⚠

### 目标

- 理解 Compose 解决的问题和声明式的管理思想
- 能用 `compose.yaml` 定义完整的多服务应用并一键启停
- 熟练掌握 Compose 的核心配置指令和常用命令
- 完成一个 Flask + Redis 的多容器项目

### 8.1 为什么需要 Compose

想象不用 Compose 启动"Web + Redis + MySQL"：三条长长的 `docker run`（各自的端口、变量、卷、网络参数），启动顺序靠脑子记，清理要逐个删，换台机器还得重敲一遍——配置散落、易错、不可复现。

Compose 用**一个 YAML 文件声明全部内容**，然后：

```
$ docker compose up -d      # 一键全部启动
$ docker compose down       # 一键全部拆除
```

它的工作方式是**声明式**的：你描述"期望状态"，Compose 负责达成；文件改了再 `up`，它只应用变化的部分，不会推倒重来。

**Dockerfile 和 Compose 文件的分工**（高频混淆点）：Dockerfile 负责**构建镜像**（怎么造零件），Compose 文件负责**运行容器**（怎么组装整机）；Compose 文件里的服务经常通过 `build` 指向一个 Dockerfile。

三步工作流：写 Dockerfile 定义应用环境 → 写 `compose.yaml` 定义服务 → `docker compose up` 启动一切。

**版本说明**：老教程里的 `docker-compose`（带连字符、独立程序、Compose V1）已过时；现在是 `docker compose`（空格、Docker 子命令、Compose V2），Docker Desktop 自带，无需单独安装。文件名推荐 `compose.yaml`（老名字 docker-compose.yml 仍兼容）。老教程里 `compose.yaml` 顶部的 `version: "3"` 字段在现代 Compose 中已废弃，可直接不写。

### 8.2 最小可用的 Compose 文件

```yaml
services:
  app:
    image: docker/welcome-to-docker
    ports:
      - "8080:80"
```

一个 `services` 顶级元素，下面每个服务名（app）对应一个容器配置。运行：

```
$ docker compose up -d
```

Compose 自动做了三件事：创建一个专属网络（项目名_default）→ 创建容器（命名"项目名-服务名-序号"）→ 按配置启动。拆除：

```
$ docker compose down
```

⚠ 注意：默认 down **不删卷**（怕你误删数据）；确认不要数据时用 `docker compose down --volumes`。

### 8.3 完整实战：Flask + Redis 计数器 ⭐

这是贯穿"官方 Compose 教程 + 菜鸟教程"的经典项目，完整走一遍。目录结构：

```
composetest/
├── app.py
├── requirements.txt
├── Dockerfile
└── compose.yaml
```

**app.py**（Python Web 应用：每次访问计数 +1，计数存在 Redis；照抄即可，注释解释每段作用）：

```python
import time
import redis
from flask import Flask

app = Flask(__name__)
# 关键点：主机名写的是服务名 "redis"！
# Compose 会把同项目的 redis 服务解析到对应容器（阶段六自定义网络 DNS 的兑现）
cache = redis.Redis(host='redis', port=6379)

def get_hit_count():
    retries = 5
    while True:
        try:
            return cache.incr('hits')        # Redis 自增计数
        except redis.exceptions.ConnectionError:
            if retries == 0:
                raise
            retries -= 1
            time.sleep(0.5)                  # Redis 没就绪时稍等重试

@app.route('/')
def hello():
    count = get_hit_count()
    return f'Hello World! 我被访问了 {count} 次。\n'
```

**requirements.txt**：

```
flask
redis
```

**Dockerfile**（阶段四最佳实践全套应用：固定版本、依赖先行、非 root）：

```dockerfile
FROM python:3.12-alpine
WORKDIR /code
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
RUN adduser -D appuser
USER appuser
CMD ["python", "app.py"]
```

**compose.yaml**：

```yaml
services:
  web:
    build: .                  # 用当前目录的 Dockerfile 构建镜像
    ports:
      - "8000:5000"           # 宿主 8000 → 容器 5000（Flask 默认端口）
    depends_on:
      - redis
  redis:
    image: redis:7.4-alpine   # 直接用官方镜像
```

启动并验证：

```
$ docker compose up -d --build
$ docker compose ps           # 查看项目内的容器
$ curl localhost:8000
Hello World! 我被访问了 1 次。
$ curl localhost:8000
Hello World! 我被访问了 2 次。
```

`--build` 表示启动前先（重新）构建镜像。多条命令断点体会：

- web 服务没有写 `networks`，但 Compose 自动把 web 和 redis 放进了同一项目网络——**服务名即主机名**，app.py 里的 `host='redis'` 因此能通
- `depends_on` 控制**启动顺序**（先起 redis 再起 web），⚠ 但它只等容器"启动"，**不等服务"就绪"**——Redis 慢一步起来时靠 app.py 里的重试逻辑兜底；更优雅的做法是健康检查（8.5 节）

### 8.4 服务配置指令详解 ⭐

逐个过 `services.<服务名>` 下最常用的指令：

**image / build —— 镜像从哪来（二选一）**

```yaml
services:
  web:
    build:
      context: ./dir              # 构建上下文路径
      dockerfile: Dockerfile-alt  # 指定其他 Dockerfile 文件名
      args:                       # 传给 ARG 的构建参数
        buildno: 1
  redis:
    image: redis:7.4              # 直接用现成镜像
```

**ports / expose —— 端口**

```yaml
    ports:
      - "8000:5000"               # 发布到宿主机（同 -p）
      - "127.0.0.1:8001:5000"     # 三段式也支持
    expose:
      - "3000"                    # 只暴露给同网络服务，不映射宿主机
```

**volumes —— 挂载**

```yaml
services:
  db:
    image: mysql:8.4
    volumes:
      - mysql-data:/var/lib/mysql         # 命名卷（在下面顶级 volumes 声明）
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql:ro   # 绑定挂载+只读
volumes:
  mysql-data:      # 顶级元素：声明命名卷
```

**environment / env_file —— 环境变量**

```yaml
    environment:
      MYSQL_ROOT_PASSWORD: secret
      DEBUG: "true"            # 布尔值要加引号，防 YAML 解析成 true/false 类型
    env_file:
      - .env                   # 从文件批量读入（.env 别提交进 git！）
```

**depends_on —— 依赖顺序**

```yaml
  web:
    depends_on:
      - db
      - redis
```

up 时按序创建（db、redis 先于 web）；stop 时反序。进阶形态见 8.5。

**restart —— 重启策略**

```yaml
    restart: always            # 容器退出总是重启（生产服务常用）
```

四种取值：`no`（默认，不重启）/ `always` / `on-failure`（非零退出才重启）/ `unless-stopped`（总是重启，除非你手动停过）。⚠ 这是单机 Docker 的策略；Swarm 集群里改用 deploy 下的 restart_policy（阶段九）。

**container_name / hostname**

```yaml
    container_name: my-web     # 固定容器名（不用 Compose 默认的 项目名-服务名-1）
```

**command / entrypoint —— 覆盖镜像默认命令**

```yaml
    command: ["redis-server", "--appendonly", "yes"]     # 覆盖 CMD
    entrypoint: /code/entrypoint.sh                       # 覆盖 ENTRYPOINT
```

（7.4 里给 redis 追加参数的做法，在 Compose 里就写成 command。）

**networks —— 加入网络**

```yaml
services:
  web:
    networks:
      - front-net
  db:
    networks:
      - back-net
networks:
  front-net:
  back-net:
```

不写则全部进默认项目网络；写了可以隔离出前后端多张网（db 对 web 可见，两个 web 之间隔离等拓扑自由编排）。

**healthcheck —— 健康检查（同 Dockerfile 语法）**

```yaml
    healthcheck:
      test: ["CMD", "redis-cli", "ping"]
      interval: 10s
      timeout: 3s
      retries: 3
      start_period: 5s
```

**deploy —— 集群部署配置（普通 compose up 会忽略它，Swarm 阶段兑现）**

```yaml
    deploy:
      replicas: 3
      resources:
        limits:
          cpus: '0.50'
          memory: 512M
      restart_policy:
        condition: on-failure
        delay: 5s
        max_attempts: 3
      update_config:
        parallelism: 1        # 滚动更新时每次更新 1 个
        delay: 10s
```

其他速览：`dns`（自定义 DNS）、`labels`（标签）、`logging`（日志驱动与轮转配置）、`secrets`（敏感数据注入）、`stop_grace_period`（停止宽限期，默认 10s）、`tmpfs`（临时文件系统挂载）。

### 8.5 depends_on + healthcheck：真正的就绪等待 ⚠

`depends_on` 默认只等"容器启动"。让 web 等 db"健康后再起"：

```yaml
  web:
    depends_on:
      db:
        condition: service_healthy     # 等健康检查通过
  db:
    image: mysql:8.4
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 5s
      retries: 10
```

这是消除"数据库还没就绪、应用先崩"竞态的标准解法。

### 8.6 Compose 常用命令 ⭐

| 命令 | 作用 |
|---|---|
| `docker compose up -d` | 创建并后台启动全部服务 |
| `docker compose up -d --build` | 启动前重新构建镜像 |
| `docker compose down` | 停止并删除容器、网络（保留卷） |
| `docker compose down --volumes` | 连卷一起删 |
| `docker compose ps` | 查看项目容器状态 |
| `docker compose logs -f` | 跟踪全部服务日志（加服务名只看单个） |
| `docker compose build` | 只构建镜像不启动 |
| `docker compose restart` | 重启服务 |
| `docker compose exec web sh` | 进入某服务容器 |
| `docker compose run web 命令` | 在服务容器里跑一次性命令 |
| `docker compose pull` | 拉取服务镜像 |
| `docker compose config` | 校验并显示最终合并的配置 |

两个进阶小件（了解）：

**watch 模式**——`docker compose watch` + 服务里配 `develop.watch`，改代码自动同步进容器并重启，替代手工 rebuild（阶段五绑定挂载"热更新"的全自动版，以下示例仍以 8.3 的 Flask 项目为例）：

```yaml
services:
  web:
    build: .
    develop:
      watch:
        - action: sync+restart      # 同步该文件并重启服务
          path: ./app.py
          target: /code/app.py      # 同步到容器内的 WORKDIR 下
```

**profiles**——给服务标 `profiles: [debug]`，只有 `--profile debug` 时才启动，实现"可选服务"。

### 练习

- [ ] 完成 8.3 Flask+Redis 全项目，curl 多次验证计数递增，`down` 后 `up -d` 再 curl 观察计数是否延续（想想为什么）
- [ ] 给 redis 服务加命名卷持久化 + `command: redis-server --appendonly yes`，down 再 up 后计数延续
- [ ] 把项目扩展成三服务：加一个 MySQL（healthcheck + depends_on condition），app 改为连接 MySQL 读取一句问候语存取（可拆两步：先只验证 web 能等服务健康后启动）
- [ ] 用 `docker compose logs -f` 观察服务启动顺序，验证 depends_on 生效
- [ ] `docker compose config` 查看你文件的最终解析结果，和手写的对比
- [ ] 挑战：给 web 服务配 `develop.watch`（同步 ./ 到容器 /code 并重建），改一句返回文案验证热更新

### 过关自检

1. Compose 解决什么问题？声明式管理的含义？
2. Dockerfile 和 compose.yaml 的分工？
3. `docker compose up -d` 自动创建了哪些东西（网络/容器命名规则）？down 默认删什么、留什么？
4. 服务之间互相访问用什么主机名？原理是什么（哪个阶段的知识）？
5. depends_on 等不等于"等对方就绪"？正确姿势是什么？
6. volumes 顶级元素和 volumes 服务级配置是什么关系？
7. `restart: always` 和 Swarm 的 restart_policy 是什么关系？

---

## 阶段九：Swarm 集群与 Docker Machine（运维进阶）

### 目标

- 理解 Swarm 的角色模型（manager/worker）与声明式服务思想
- 单节点上完成 Swarm 全套核心操作：初始化、部署、扩缩、滚动升级
- 了解多节点集群的搭建方式，掌握 node 管理
- 用 stack 把 Compose 文件部署到集群；明确 Docker Machine 的历史定位

### 9.1 Swarm 是什么

Docker Swarm 是 Docker **内置**的集群管理工具（装了 Docker 就有，不用额外装东西）：把一群 Docker 主机池变成**一台虚拟的 Docker 主机**来用。你在 manager 上声明"我要 3 个副本的 web 服务"，Swarm 自动调度到各节点、持续维持这个状态（挂一个自动补一个）。

集群两种角色：

- **manager（管理节点）**：负责整个集群的管理——配置、服务调度、状态维护。⚠ 管理操作只在 manager 上执行
- **worker（工作节点）**：接收 manager 派发的任务（task），运行容器

Swarm 的核心特性（理解后你说的每句都有出处）：声明式服务模型（说结果不说过程）、期望状态协调（实际状态自动向期望收敛）、服务扩缩容、多主机网络（overlay）、内置服务发现与负载均衡、默认 TLS 加密节点通信、滚动更新。

**和 Compose 的关系**：Compose 管**一台主机**上的多容器；Swarm 管**多台主机**上的容器集群；Swarm 的 stack 直接用 Compose 文件格式部署（`deploy:` 指令到这时才真正生效——阶段八埋的伏笔在此兑现）。

### 9.2 单节点 Swarm：Docker Desktop 就能练 ⭐

你的 Mac 上就能初始化一个"单节点集群"，把全部核心命令练熟：

**初始化：**

```
$ docker swarm init --advertise-addr 127.0.0.1
Swarm initialized: current node (xxx) is now a manager.
```

`--advertise-addr` 告诉其他节点用什么地址找 manager（单机练习写 127.0.0.1）。初始化的这台就是 manager。输出里还会给一条 `docker swarm join --token ...` 命令——那是给别的机器加入集群用的，单机练习用不上，多节点时见 9.4。

**查看节点：**

```
$ docker node ls
ID                            HOSTNAME     STATUS    AVAILABILITY    MANAGER STATUS
xxx *                         moby         Ready     Active          Leader
```

带 `*` 的是当前所在节点；单节点集群里它身兼 manager+worker。

**部署一个服务：**

```
$ docker service create --name helloworld --replicas 1 alpine ping docker.com
```

`docker service create` 和 `docker run` 长得像，但语义不同：run 是"起一个容器"，service 是"声明一个服务，请保持 1 个副本运行"。

**查看服务与任务：**

```
$ docker service ls                  # 服务列表（REPLICAS 列：1/1 = 期望1个，运行1个）
$ docker service ps helloworld       # 任务分布在哪个节点、什么状态
$ docker service inspect --pretty helloworld   # 服务详情（可读格式）
```

**扩缩容：**

```
$ docker service scale helloworld=3
$ docker service ps helloworld       # 现在 3 个任务在跑
```

**删除服务：**

```
$ docker service rm helloworld
```

**滚动升级（Swarm 最出彩的能力）⭐：**

```
$ docker service create --replicas 3 --name redis --update-delay 10s redis:7.2
$ docker service update --image redis:7.4 redis
$ docker service ps redis       # 观察任务逐个替换：新版本起一个、旧版本停一个，间隔 10s
```

`--update-delay` 是批次间隔；整个过程服务不中断，出问题还会自动回滚暂停（可在 update_config 里细调，见阶段八 8.4 的 deploy 段）。

**退出 Swarm 模式**（练习完恢复普通单机）：

```
$ docker swarm leave --force
```

### 9.3 overlay 网络：跨主机的容器互联

多台主机上的容器怎么互相访问？靠 overlay（叠加）网络：manager 创建 overlay 网络后，不同物理机上加入该网络的服务容器就能像在同一个局域网一样按服务名互访（流量在主机间加密隧道里走）。操作上和阶段六一样是"创建网络 + 服务加入"，只是驱动换了：

```
$ docker network create -d overlay my-overlay
$ docker service create --network my-overlay --name web ...   # 语法见 service create 文档
```

初学阶段理解到"overlay = Swarm 版的 bridge，跨主机版的自定义网络"即可。

### 9.4 多节点集群怎么搭（实验环境）

真集群需要多台 Linux 机器，两条可行路径（任选）：

**路径 A：本机虚拟机。** 用 VirtualBox（免费）装两台最小化 Ubuntu 虚拟机（各自 2GB 内存即可），每台里面装 Docker Engine（Linux 原生安装，无需 Desktop）。Ubuntu 上的安装最简方式（在每台虚拟机里执行）：

```
$ curl -fsSL https://get.docker.com -o get-docker.sh
$ sudo sh get-docker.sh
$ sudo systemctl enable --now docker
```

（国内可先配置镜像加速：编辑 `/etc/docker/daemon.json` 写入 `{"registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]}` 后 `sudo systemctl restart docker`。）

**路径 B：云主机。** 各云厂商最低配的 2~3 台 Linux 主机，安装方式同上，比虚拟机更接近生产。

组网步骤（manager 机器上 init，输出的 join 命令到两台 worker 上执行）：

```
# manager 机（假设 IP 192.168.99.107）
$ docker swarm init --advertise-addr 192.168.99.107
# 它会输出：
#   docker swarm join --token SWMTKN-1-xxxx 192.168.99.107:2377

# worker 机 ×2（原样执行上面输出的命令）
$ docker swarm join --token SWMTKN-1-xxxx 192.168.99.107:2377

# 回到 manager 验证
$ docker node ls      # 应看到 3 个节点：1 个 Leader manager + 2 个 worker
```

节点管理（都在 manager 上执行）：

```
$ docker node update --availability drain worker1   # 排空：不再派新任务，已有任务迁走（维护机器时用）
$ docker node update --availability active worker1  # 恢复
$ docker node ls                                     # AVAILABILITY 列变 Drain/Active
```

drain 是运维日常：某台机器要打补丁，先 drain 让服务迁到别的节点，机器空了再维护，完事 active 回来。

### 9.5 Stack：用 Compose 文件部署集群应用 ⭐

单机时代你写 compose.yaml；集群时代同一个文件（补上 deploy 配置）直接部署到 Swarm，这叫 **stack（应用栈）**。把阶段八的 Flask+Redis 升级成集群版 compose.yaml：

```yaml
services:
  web:
    image: 你的用户名/composetest:1.0      # ⚠ Swarm 各节点要能拉到镜像：先 build+push 到仓库
    deploy:
      replicas: 3                          # 3 个副本
      restart_policy:
        condition: on-failure
      update_config:
        parallelism: 1
        delay: 10s
      resources:
        limits:
          cpus: '0.50'
          memory: 256M
    ports:
      - "8000:5000"
  redis:
    image: redis:7.4-alpine
    deploy:
      replicas: 1
```

部署与 管理（在 manager 上）：

```
$ docker stack deploy -c compose.yaml myapp     # 部署名为 myapp 的应用栈
$ docker stack ls                                # 列出栈
$ docker stack services myapp                    # 栈内服务状态
$ docker stack ps myapp                          # 栈内任务
$ docker stack rm myapp                          # 拆除
```

注意两个差异：① web 改用 `image` 引用已推送的镜像而不是 `build`（集群节点要各自拉镜像，本地构建的镜像别的机器看不到）；② 端口发布到 8000 后，**访问任意一台集群机器的 8000 都能打到服务**——Swarm 的入口负载均衡会把请求路由到某个副本，这是"集群变成一台虚拟主机"的直接体现。之后 `docker service scale myapp_web=5` 动态扩容、改镜像 tag 后重新 `stack deploy` 即完成滚动升级。

### 9.6 Docker Machine：了解它的历史定位

菜鸟教程专门有一章 Docker Machine，这里给你**当前（2025）的正确认知**：它已被淘汰，不再推荐学习和使用。它诞生于 Docker 早期（1.x~18.x 时代），用于"在虚拟机/云主机上批量创建并管理装好 Docker 的主机"。如今：

| 对比项 | Docker Machine | Docker Desktop |
|---|---|---|
| 定位 | Docker 主机批量创建工具 | 本地一体化开发环境 |
| 本地实现 | 依赖 VirtualBox 等外部虚拟化软件 | 内置轻量 Linux VM |
| macOS/Windows 体验 | 间接、复杂 | 原生、开箱即用 |
| 维护状态 | 基本停止更新 | 持续活跃 |
| 当前推荐度 | ❌ 不推荐 | ✅ 强烈推荐 |

它的遗产只是命令形态可作了解（`docker-machine create/ls/ip/ssh/start/stop`），多主机管理职能已被 Terraform 等基础设施工具和 Kubernetes 生态接棒。老教程用 VirtualBox+Machine 搭 Swarm 的做法，用 9.4 的"VirtualBox 虚拟机 + 原生 Docker Engine"替代即可。

### 练习

- [ ] 单节点 `docker swarm init`，部署 helloworld 服务，`service ls/ps/inspect` 各看一遍
- [ ] `scale helloworld=3` 再缩回 1，观察 ps 输出变化
- [ ] 完成一次 redis 7.2→7.4 滚动升级，`service ps` 里找出新旧任务交替的痕迹
- [ ] 把阶段八项目构建推送，写出带 deploy 的 compose.yaml，`stack deploy` 部署，浏览器访问验证
- [ ] `stack ps` 观察 3 个 web 副本分布；`service scale` 扩到 5 再访问
- [ ] （有条件）搭 3 节点集群：node ls 确认拓扑；drain 一个 worker，观察任务迁移；恢复
- [ ] 练完 `docker swarm leave --force` 恢复单机模式，`docker info` 确认

### 过关自检

1. manager 和 worker 各自的职责？管理命令在哪类节点执行？
2. `docker run` 和 `docker service create` 的语义差异？"期望状态协调"什么意思？
3. 滚动升级期间服务会中断吗？靠什么机制保证？
4. overlay 网络和 bridge 网络的关系？解决什么问题？
5. 部署 stack 时为什么服务要用 image 而不是 build？
6. drain 一个节点时发生了什么？什么场景需要 drain？
7. Docker Machine 为什么被淘汰？它的职能被谁接替？

---

## 阶段十：毕业项目

三个项目递进：项目一是主体（单机全栈应用），项目二把它搬进集群，项目三完成发布闭环。做完这三个，你就具备了"拿到一个应用，从 Dockerfile 写到集群部署"的完整能力。

### 项目一：容器化全栈留言板（单机版）⭐

**目标形态**：一条命令拉起整套服务，数据可持久化、服务有健康检查、镜像符合生产规范。

**技术栈**：Python Flask（Web）+ MySQL（存留言）+ Redis（计数/缓存）+ Nginx（反向代理入口）。

**目录结构**：

```
guestbook/
├── app/
│   ├── app.py
│   ├── requirements.txt
│   └── Dockerfile
├── nginx/
│   └── nginx.conf
├── .env
└── compose.yaml
```

**第 1 步：应用代码**（`app/app.py`，可直接使用；逻辑：访问计数存 Redis，留言读写 MySQL）：

```python
import os
import time
import redis
import pymysql
from flask import Flask, request

app = Flask(__name__)

# 服务名当主机名（Compose 网络内置 DNS）
cache = redis.Redis(host='redis', port=6379)

def db():
    return pymysql.connect(
        host='db', user=os.environ['MYSQL_USER'],
        password=os.environ['MYSQL_PASSWORD'],
        database=os.environ['MYSQL_DATABASE'])

def ensure_table():
    for i in range(10):                          # 等数据库就绪
        try:
            with db().cursor() as c:
                c.execute("""CREATE TABLE IF NOT EXISTS msgs (
                             id INT AUTO_INCREMENT PRIMARY KEY,
                             body VARCHAR(200))""")
            db().commit()
            return
        except pymysql.err.OperationalError:
            time.sleep(2)
    raise RuntimeError("db not ready")

@app.route('/')
def index():
    ensure_table()
    visits = cache.incr('visits')
    with db().cursor() as c:
        c.execute("SELECT body FROM msgs ORDER BY id DESC LIMIT 10")
        rows = c.fetchall()
    msgs = "<br>".join(r[0] for r in rows)
    return f"<h1>访客计数：{visits}</h1><form method='post'><input name='body'><button>留言</button></form><div>{msgs}</div>"

@app.route('/', methods=['POST'])
def add():
    body = request.form.get('body', '')[:200]
    if body:
        with db().cursor() as c:
            c.execute("INSERT INTO msgs (body) VALUES (%s)", (body,))
        db().commit()
    return index()
```

`app/requirements.txt`：

```
flask
redis
pymysql
```

**第 2 步：生产级 Dockerfile**（`app/Dockerfile`，阶段四全部要点的集中演练）：

```dockerfile
FROM python:3.12-slim AS builder
WORKDIR /install
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install/deps -r requirements.txt

FROM python:3.12-slim
WORKDIR /app
# 只把依赖从构建阶段搬过来，再拷源码（依赖先行 + 多阶段）
COPY --from=builder /install/deps /usr/local
COPY app.py .
RUN useradd -r guest
USER guest
EXPOSE 5000
CMD ["python", "app.py"]
```

同级加 `.dockerignore`（内容：`__pycache__`、`*.pyc`、`.git`）。

**第 3 步：Nginx 反向代理配置**（`nginx/nginx.conf`）：

```nginx
server {
    listen 80;
    location / {
        proxy_pass http://web:5000;    # upstream 用服务名
        proxy_set_header Host $host;
    }
}
```

**第 4 步：compose.yaml**（阶段八全部指令的综合运用）：

```yaml
services:
  web:
    build: ./app
    environment:
      MYSQL_USER: guest
      MYSQL_PASSWORD: ${MYSQL_PASSWORD}   # 变量插值：启动时从项目根目录 .env 读取
      MYSQL_DATABASE: guestbook
    depends_on:
      db:
        condition: service_healthy
      redis:
        condition: service_started
    restart: unless-stopped
  db:
    image: mysql:8.4
    environment:
      MYSQL_ROOT_PASSWORD_FILE: /run/secrets/db_root_pass
      MYSQL_USER: guest
      MYSQL_PASSWORD_FILE: /run/secrets/db_pass
      MYSQL_DATABASE: guestbook
    volumes:
      - db-data:/var/lib/mysql
    secrets:
      - db_pass
      - db_root_pass
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      interval: 5s
      retries: 12
    restart: unless-stopped
  redis:
    image: redis:7.4-alpine
    command: ["redis-server", "--appendonly", "yes"]
    volumes:
      - redis-data:/data
    restart: unless-stopped
  nginx:
    image: nginx:1.27
    ports:
      - "8080:80"
    volumes:
      - ./nginx/nginx.conf:/etc/nginx/conf.d/default.conf:ro
    depends_on:
      - web
    restart: unless-stopped

volumes:
  db-data:
  redis-data:

secrets:
  db_pass:
    file: ./.secrets/db_pass
  db_root_pass:
    file: ./.secrets/db_root_pass
```

**敏感信息管理说明（两种方式对照记忆）**：

- **db 服务走 secrets**：`./.secrets/db_pass`、`./.secrets/db_root_pass` 是两个各写着一个密码的文本文件（`mkdir -p .secrets && echo '你的密码' > .secrets/db_pass`）。MySQL 镜像原生支持 `*_FILE` 变量从文件读密码，Docker 把 secrets 内容挂到容器的 `/run/secrets/` 下，密码全程不以明文出现在配置里
- **web 服务走 .env 插值**：`${MYSQL_PASSWORD}` 是变量插值——Compose 启动时自动读取项目根目录的 `.env` 文件（内容就一行：`MYSQL_PASSWORD=你的密码`）替换进配置，最终以普通环境变量注入容器，被 app.py 的 `os.environ['MYSQL_PASSWORD']` 读到（`.env` 机制与 8.4 的 `env_file` 指令是两回事：前者做文本替换、后者注入环境变量，此处用前者）

⚠ 无论哪种方式，`.env` 和 `.secrets/` 都必须加进 `.gitignore`，绝不提交进仓库；明文密码打进镜像或配置是生产大忌。

**第 5 步：验收清单**

```
$ docker compose up -d --build
```

- [ ] 浏览器打开 `http://localhost:8080`，看到计数器
- [ ] 提交两条留言，刷新可见
- [ ] `docker compose ps` 四个服务全部运行（db 显示 healthy）
- [ ] `docker compose down && docker compose up -d`，留言还在（卷持久化生效）
- [ ] `docker rm -f` 强杀 db 容器，等几秒刷新页面恢复（restart 策略 + 数据在卷里）
- [ ] 改一行 app.py 后 `docker compose up -d --build`，只重建 web 服务
- [ ] `docker compose logs -f web` 能看到访问日志
- [ ] `docker stats` 里能看到各容器资源占用

### 项目二：把留言板搬进 Swarm（集群版）

**改造要点**（对应阶段九 9.5）：

1. web 镜像 build 后 push 到 Docker Hub：`docker build -t 你的用户名/guestbook:1.0 ./app && docker push ...`
2. compose.yaml 里 web 的 `build: ./app` 换成 `image: 你的用户名/guestbook:1.0`，并加 `deploy:`（web 副本 3、滚动更新参数、资源限制）
3. secrets 在 Swarm 里由集群分发（同样用 file 来源即可）
4. manager 上 `docker stack deploy -c compose.yaml gb`

**验收清单**：

- [ ] `docker stack ps gb` 显示 3 个 web 副本
- [ ] 连续刷新页面，访客计数连续增长而响应来自不同副本（可在响应中加副本主机名观察）
- [ ] 构建 `:2.0` 镜像（改一行页面标题），更新 compose 后重新 `stack deploy`，观察滚动升级不中断
- [ ] `docker service scale gb_web=1` 缩容后再扩回 3
- [ ] （有条件）drain 一个 worker 节点，验证副本自动迁移、服务不中断

### 项目三（可选）：发布与自动化

- [ ] 给 guestbook 镜像打 `1.0` 和 `latest` 两个 tag 都推到 Docker Hub
- [ ] 项目根目录写 README：架构图（文字版）、启动命令、各文件作用
- [ ] 有 GitHub 账号的话：建仓库托管代码，配置 Actions 在 push 时自动 `docker build`（CI 中自动构建镜像，官方最佳实践"在 CI 中构建和测试镜像"的落地）

### 过关自检（毕业总考，全答上方能出师）

1. 从零部署这套留言板到一台新服务器，你的完整操作序列是什么？
2. 每个服务为什么单独一个容器？哪个环节体现了"期望状态协调"？
3. 数据安全由哪些机制保证（卷/secrets/重启策略/健康检查各扮演什么角色）？
4. 单机 Compose 和 Swarm stack 部署同一份应用，配置上最大的区别是什么？
5. 如果 web 服务响应变慢，你的排查路径是什么？（提示：stats → logs → inspect → exec）

---

## 学习原则

1. **两遍学习法**：第一遍走完本路线、做完项目，建立"能干活"的肌肉记忆；第二遍回头**精读官方文档**（Docker 概念系列、Dockerfile 最佳实践、Compose 与 Swarm 手册），把"会用"升级为"知其所以然"。第一遍不求全懂，遇到卡点标记后继续走，第二遍逐个击破。
2. **命令记不住是常态**：Docker 命令体系庞大，记住常用的 20%，其余靠 `docker --help`、`docker 子命令 --help` 和速查表（附录 A）。会查比硬背重要。
3. **每个知识点都动手**：本路线的练习全部基于 Docker Desktop 单机可完成（多节点部分除外）。容器、卷、网络都是"删了重来零成本"的东西——**大胆破坏**，破坏后重建是最好的学习。
4. **坚持用版本号**：从一开始就写 `mysql:8.4` 而不是 `mysql`，这个习惯会为未来的你省掉无数诡异问题。
5. **遇到问题的排查三板斧**：`docker ps -a`（状态）→ `docker logs`（日志）→ `docker inspect`（配置）。九成问题出在这三处可见的信息里。

## 附录 A：Docker 命令速查表

**容器生命周期**：`run`（创建并启动）· `start/stop/restart` · `kill`（强杀）· `rm` · `pause/unpause` · `create`（只建不启）· `exec` · `rename`

**容器操作**：`ps` · `inspect` · `top` · `attach` · `logs` · `events` · `wait` · `export` · `port` · `stats` · `update`（改资源限制）· `diff`（文件系统变更）· `cp`（容器与宿主机互拷文件）

**rootfs**：`commit`（容器→镜像）· `cp` · `diff`

**镜像仓库**：`login/logout` · `pull` · `push` · `search`

**本地镜像**：`images` · `rmi` · `tag` · `build` · `history` · `save/load` · `import`

**信息**：`info` · `version`

**网络**：`network ls/create/rm/inspect/connect/disconnect`

**卷**：`volume ls/create/rm/inspect/prune`

**Compose**：`up/down/ps/logs/build/restart/exec/run/pull/ls/config`（含 `-d`、`--build`、`--volumes` 常用旗标）

**Swarm/服务**：`swarm init/join/leave` · `node ls/update` · `service create/ls/ps/inspect/scale/update/rm` · `stack deploy/ls/services/ps/rm`

## 附录 B：macOS 终端操作补充

- 打开终端：聚焦搜索（`Cmd + Space`）输入 Terminal 回车
- `.dockerignore`/`Dockerfile` 这类"无扩展名文件"可用任意编辑器创建（VS Code、TextEdit 的纯文本模式均可）；VS Code 装 Docker 官方扩展后还能图形化管理容器和获得 Dockerfile 语法提示
- macOS 的 `$HOME` 即 `/Users/你的用户名`，文中 `~/docker-lab` 等价于 `/Users/你的用户名/docker-lab`
- Apple Silicon Mac 拉取仅 x86 的老镜像时会提示平台不匹配，Docker Desktop 默认用 Rosetta 转译运行（性能略降）；选镜像时优先选多架构镜像

## 附录 C：延伸资料（学完第二遍后按需深入）

| 资料 | 定位 |
|---|---|
| Docker 官方文档 Get Started（docker-concepts 概念系列） | 本路线概念部分的源头，图文视频并茂，适合第二遍精读 |
| Dockerfile 最佳实践（官方 Building best practices） | 阶段四的深度扩展，含供应链安全与 CI 建议 |
| Docker 官方镜像仓库（Docker Hub） | 找镜像、看镜像文档（每个官方镜像的 Usage 都值得读） |
| 菜鸟教程 Docker 命令大全 | 网页版速查，配合附录 A 使用 |
| 《Docker——从入门到实践》（开源书） | 中文进阶读物，覆盖原理与生态 |
| Docker Compose 官方 How-tos | watch 模式、多 Compose 文件、环境变量专题 |
| Docker Swarm 官方教程（swarm-tutorial） | 阶段九的多节点完整版教程 |

## 附录 D：后续深入主题（超出本路线范围，学有余力再战）

- **Kubernetes**：容器编排的事实标准，Swarm 之后的必然方向（调度、滚动发布、自愈、存储/网络抽象）
- **容器运行时与生态**：containerd、runc、OCI 标准；Docker 只是生态的一层皮
- **镜像安全**：镜像漏洞扫描（Docker Scout 等）、最小化基础镜像、distroless、镜像签名
- **构建进阶**：BuildKit 高级特性（并行构建、缓存后端、secret 挂载）、buildx 多平台构建
- **网络进阶**：macvlan/ipvlan、容器网络原理、自定义网桥参数调优
- **存储进阶**：tmpfs 挂载、卷驱动（NFS/云盘）、存储驱动原理
- **监控与日志**：容器指标采集、日志驱动、日志聚合方案
- **Docker Engine 生产部署**：Linux 上以守护进程方式部署、daemon.json 全量配置、rootless 模式、容器安全加固
- **DevOps 集成**：CI/CD 流水线中自动构建与推送镜像、镜像版本策略（Git SHA tag）
