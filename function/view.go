/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-15 10:24:11

Description: 子命令`view`功能函数
*/

package function

import (
	"fmt"
	"strings"
	"time"
)

func SystemInfo() {
	// 检索的Pacman日志文件
	var fileName = "/var/log/pacman.log"

	// 获取系统安装时间和当前时间
	var lineText = ReadFileLine(fileName, 1)
	var startTimeStrTZ = strings.Split(strings.Split(lineText, "[")[1], "]")[0] // 2023-03-10T10:49:09+0800
	var currentTimeStr = time.Now().Format("2006-01-02 15:04")

	// 获取初始和当前内核版本
	var keyText = ReadFileKey(fileName, "installed linux ")
	var firstKernel = strings.Split(strings.Split(keyText, " (")[1], ")")[0]
	var unameArgs = []string{"-r"}
	var latestKernel = RunCommandGetResult("uname", unameArgs)

	// 计算系统安装天数
	var local, _ = time.LoadLocation("Asia/Shanghai")
	var startTime, _ = time.ParseInLocation("2006-01-02T15:04:05Z0700", startTimeStrTZ, local)
	var startTimeStr = startTime.Format("2006-01-02 15:04")
	var startTimeStamp = startTime.Unix()
	var currentTime, _ = time.ParseInLocation("2006-01-02 15:04", currentTimeStr, local)
	var currentTimeStamp = currentTime.Unix()
	var systemDays = (currentTimeStamp - startTimeStamp) / 86400

	// 获取系统/内核更新相关数据
	var systemUpdateCount = ReadFileCount(fileName, "starting full system upgrade")
	var systemUpdateMean = float32(systemUpdateCount) / float32(systemDays)
	var kernelUpdateCount = ReadFileCount(fileName, "upgraded linux ")
	var kernelUpdateMean = float32(systemDays) / float32(kernelUpdateCount)

	// 获取吉祥物
	var repoArgs = []string{""}
	var mascot = RunCommandGetResult("repo-elephant", repoArgs)

	// 输出
	fmt.Printf("\033[36m[%16v]\033[0m %-2v \033[36m[%-16v]\033[0m\n", startTimeStr, "--", currentTimeStr)
	fmt.Printf("\033[35m%18v\033[0m %-2v \033[35m%-18v\033[0m\n", firstKernel, "--", latestKernel)
	fmt.Printf("\033[37m%12v\033[0m %-2v \033[37m%-4.3v\033[0m \033[34m%v\033[0m\n", "系统使用时长", "--", systemDays, "天")
	fmt.Printf("\033[37m%12v\033[0m %-2v \033[37m%-4.3v\033[0m \033[34m%v\033[0m\n", "系统更新次数", "--", systemUpdateCount, "次")
	fmt.Printf("\033[37m%12v\033[0m %-2v \033[37m%-4.3v\033[0m \033[34m%v\033[0m\n", "系统更新频率", "--", systemUpdateMean, "次/天")
	fmt.Printf("\033[37m%12v\033[0m %-2v \033[37m%-4.3v\033[0m \033[34m%v\033[0m\n", "内核更新次数", "--", kernelUpdateCount, "次")
	fmt.Printf("\033[37m%12v\033[0m %-2v \033[37m%-4.3v\033[0m \033[34m%v\033[0m\n", "内核更新频率", "--", kernelUpdateMean, "天/次")
	fmt.Println(mascot)
}
