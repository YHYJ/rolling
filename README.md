<h1 align="center">Rolling</h1>

<!-- File: README.md -->
<!-- Author: YJ -->
<!-- Email: yj1516268@outlook.com -->
<!-- Created Time: 2023-02-20 16:24:15 -->

---

<p align="center">
  <a href="https://github.com/YHYJ/rolling/actions/workflows/release.yml"><img src="https://github.com/YHYJ/rolling/actions/workflows/release.yml/badge.svg" alt="Go build and release by GoReleaser"></a>
</p>

---

## Table of Contents

<!-- vim-markdown-toc GFM -->

* [Install](#install)
  * [一键安装](#一键安装)
* [Usage](#usage)
* [Compile](#compile)
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)

<!-- vim-markdown-toc -->

---

<!----------------------------------->
<!--            _ _ _              -->
<!--  _ __ ___ | | (_)_ __   __ _  -->
<!-- | '__/ _ \| | | | '_ \ / _` | -->
<!-- | | | (_) | | | | | | | (_| | -->
<!-- |_|  \___/|_|_|_|_| |_|\__, | -->
<!--                        |___/  -->
<!----------------------------------->

---

适用于 Arch Linux 的系统安装信息统计工具

## Install

### 一键安装

```bash
curl -fsSL https://raw.githubusercontent.com/YHYJ/rolling/main/install.sh | sudo bash -s
```

## Usage

- `view`子命令

  该子命令用于查看系统安装和更新信息

- `version`子命令

  查看程序版本信息

- `help`子命令

  查看程序帮助信息

## Compile

### 当前平台

```bash
go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/rolling/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/rolling/general.BuildTime=`date +%s` -X github.com/yhyj/rolling/general.BuildBy=$USER" -o build/rolling main.go
```

### 交叉编译

使用命令`go tool dist list`查看支持的平台

#### Linux

```bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/rolling/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/rolling/general.BuildTime=`date +%s` -X github.com/yhyj/rolling/general.BuildBy=$USER" -o build/rolling main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是 x86_64 则 GOARCH=amd64
> - 结果是 aarch64 则 GOARCH=arm64
