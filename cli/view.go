/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-15 10:24:11

Description: 子命令`view`功能函数
*/

package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/yhyj/rolling/general"
)

// SystemInfo 获取系统信息
func SystemInfo() {
	// 检索的 Pacman 日志文件
	var fileName = "/var/log/pacman.log"

	// 获取系统安装时间和当前时间
	lineText := general.ReadFileLine(fileName, 1)
	startTimeStrTZ := strings.Split(strings.Split(lineText, "[")[1], "]")[0] // 2023-03-10T10:49:09+0800
	currentTimeStr := time.Now().Format("2006-01-02 15:04")

	// 获取初始和当前内核版本
	keyText := general.ReadFileKey(fileName, "installed linux ")
	firstKernel := strings.Replace(strings.Split(strings.Split(keyText, " (")[1], ")")[0], ".", "-", -1)
	unameArgs := []string{"-r"}
	latestKernel, err := general.RunCommandGetResult("uname", unameArgs)
	if err != nil {
		fmt.Printf(general.ErrorBaseFormat, err)
	}

	// 计算系统安装天数
	local, _ := time.LoadLocation("Asia/Shanghai")
	startTime, _ := time.ParseInLocation("2006-01-02T15:04:05Z0700", startTimeStrTZ, local)
	startTimeStr := startTime.Format("2006-01-02 15:04")
	startTimeStamp := startTime.Unix()
	currentTime, _ := time.ParseInLocation("2006-01-02 15:04", currentTimeStr, local)
	currentTimeStamp := currentTime.Unix()
	systemDays := (currentTimeStamp - startTimeStamp) / 86400

	// 获取系统/内核更新相关数据
	systemUpdateCount := general.ReadFileCount(fileName, "system_checkupdate.hook")
	systemUpdateMean := float32(systemUpdateCount) / float32(systemDays)
	kernelUpdateCount := general.ReadFileCount(fileName, "upgraded linux ")
	kernelUpdateMean := float32(systemDays) / float32(kernelUpdateCount)

	// 获取吉祥物
	repoArgs := []string{""}
	mascot, err := general.RunCommandGetResult("repo-elephant", repoArgs)
	if err != nil {
		fmt.Printf(general.ErrorBaseFormat, err)
	}

	// 输出
	dataFormat1 := "\x1b[36m[%16v]\x1b[0m %-2v \x1b[36m[%-16v]\x1b[0m\n"
	dataFormat2 := "\x1b[35m%18v\x1b[0m %-2v \x1b[35m%-18v\x1b[0m\n"
	dataFormat3 := "\x1b[37m%12v\x1b[0m %-2v \x1b[37m%-4.3v\x1b[0m \x1b[34m%v\x1b[0m\n"
	fmt.Printf(dataFormat1, startTimeStr, "--", currentTimeStr)
	fmt.Printf(dataFormat2, firstKernel, "--", latestKernel)
	fmt.Printf(dataFormat3, "系统使用时长", "--", systemDays, "天")
	fmt.Printf(dataFormat3, "系统更新次数", "--", systemUpdateCount, "次")
	fmt.Printf(dataFormat3, "系统更新频率", "--", systemUpdateMean, "次/天")
	fmt.Printf(dataFormat3, "内核更新次数", "--", kernelUpdateCount, "次")
	fmt.Printf(dataFormat3, "内核更新频率", "--", kernelUpdateMean, "天/次")
	fmt.Println(mascot)
}
