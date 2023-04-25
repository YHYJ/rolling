# README

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-02-20 16:24:15 -->

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Usage](#usage)
* [Compile](#compile)

<!-- vim-markdown-toc -->

---

<!-- Object info -->

---

适用于Arch Linux的系统安装信息统计工具

## Usage

- `view`子命令

    该子命令用于查看系统安装和更新信息

- `version`子命令

    查看程序版本信息

- `help`

    查看程序帮助信息

## Compile

- 编译当前平台可执行文件：

```bash
go build main.go
```

- **交叉编译**指定平台可执行文件：

```bash
# 适用于Linux AArch64平台
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build main.go
```

```bash
# 适用于macOS amd64平台
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
```

```bash
# 适用于Windows amd64平台
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```
