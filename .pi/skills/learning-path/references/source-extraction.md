# 素材抓取与提取指南

生成学习路线前的素材核实工具箱。原则：**一切以抓到的真实内容为准**。

## 通用流程

### 1. 可达性验证
```bash
code=$(curl -s -o /dev/null -w "%{http_code}" --max-time 20 -H "User-Agent: Mozilla/5.0" "https://example.com/doc/page")
```

### 2. 全文抓取
```bash
curl -s -L --max-time 40 -H "User-Agent: Mozilla/5.0" "https://example.com/doc/page" -o /tmp/page.html
```
注意事项：
- **带 `-L` 跟随重定向**：很多文档站旧链接 301（如 go.dev 的 `.html` 后缀会 301 到无后缀地址）
- **带 User-Agent**：部分站点对无 UA 请求返回异常内容
- 网络慢时用后台并行 `&` + `wait` 批量抓，并设 timeout

### 3. 目录（TOC）提取
优先从静态 HTML 提取章节链接：
```bash
grep -oE 'href="[^"]*"' /tmp/page.html | grep -iE '关键词' | sort -u
```
再逐页抓 `<title>` 核实章节真实名称，**不要凭 URL 猜标题**。

### 4. HTML → 文本提炼
保留代码块与标题层级、剔除导航噪音的提取脚本（python3，无第三方依赖）：
```python
import re, os, html as h

def extract(raw):
    # 定位正文区（按站点调整：runoob 用 id=content/maincontent，go.dev 用 id=article-body）
    m = re.search(r'<div[^>]*id="article-body"[^>]*>(.*)', raw, re.S) or \
        re.search(r'<div[^>]*id="maincontent"[^>]*>(.*)', raw, re.S)
    body = m.group(1) if m else raw
    body = re.sub(r'<script.*?</script>|<style.*?</style>', '', body, flags=re.S)
    body = re.sub(r'<pre[^>]*>', '\n```\n', body)          # 代码块 → 围栏
    body = re.sub(r'</pre>', '\n```\n', body)
    body = re.sub(r'<(h[1-6])[^>]*>', lambda m: '\n' + '#'*int(m.group(1)[1]) + ' ', body)
    body = re.sub(r'<br\s*/?>', '\n', body)
    body = re.sub(r'<li[^>]*>', '\n- ', body)
    body = re.sub(r'<(p|div|tr|td|th)[^>]*>', '\n', body)
    text = h.unescape(re.sub(r'<[^>]+>', '', body))
    return re.sub(r'\n{3,}', '\n\n', text).strip()
```

## 常见坑（真实案例）

| 现象 | 原因与解法 |
|---|---|
| 章节目录提取不全 | 侧边栏由 JS 渲染，不在静态 HTML 里。改为在正文区 grep 章节链接（`href="/xx/xx.html"`），再抓各页 title 核实 |
| 教程只有第一部分 | 分页教程。在当前页 grep 同前缀链接（如 `/doc/tutorial/`）发现全部子页，逐一抓取 |
| 抓下来只有 60 字节 | 301 重定向页。去掉 `.html` 后缀重试 + `-L` |
| 正文混满导航菜单 | 正文 div 定位失败导致全页提取。查该站正文容器的 id/class，用正则先截取正文区 |
| 内容提取后仍是导航噪音为主 | 从 `Documentation`/正文标题锚点截断后再解析 |

## 已验证站点速查表

| 站点 | 技巧 |
|---|---|
| runoob.com | 目录不在静态 HTML；grep 正文 href 提取全部章节页；正文容器 `id=content` / `id=maincontent`；页面自带在线运行环境链接 |
| go.dev | URL 不带 `.html` 后缀；分页教程通过 grep `/doc/tutorial/` 发现全部子页；正文容器 `id=article-body`；教程页验证清单：getting-started、create-module（7 个分页）、generics、database-access、web-service-gin、fuzz |

新站点按"可达性 → 目录提取 → 单页正文提取 → 批量抓取"四步试跑一页成功后再批量。
