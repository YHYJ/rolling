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
  * [当前平台](#当前平台)
  * [交叉编译](#交叉编译)
    * [Linux](#linux)
    * [macOS](#macos)
    * [Windows](#windows)

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

适用于Arch Linux的系统安装信息统计工具

## Usage

- `view`子命令

    该子命令用于查看系统安装和更新信息

- `version`子命令

    查看程序版本信息

- `help`

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
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

#### macOS

```bash
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -X github.com/yhyj/rolling/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/rolling/general.BuildTime=`date +%s` -X github.com/yhyj/rolling/general.BuildBy=$USER" -o build/rolling main.go
```

> 使用`uname -m`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64

#### Windows

```powershell
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -gcflags="-trimpath" -ldflags="-s -w -H windowsgui -X github.com/yhyj/rolling/general.GitCommitHash=`git rev-parse HEAD` -X github.com/yhyj/rolling/general.BuildTime=`date +%s` -X github.com/yhyj/rolling/general.BuildBy=$USER" -o build/rolling.exe main.go
```

> 使用`echo %PROCESSOR_ARCHITECTURE%`确定硬件架构
>
> - 结果是x86_64则GOARCH=amd64
> - 结果是aarch64则GOARCH=arm64
