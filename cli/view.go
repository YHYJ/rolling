/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-15 10:24:11

Description: 子命令 'view' 的实现
*/

package cli

import (
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/yhyj/rolling/general"
)

// SystemInfo 获取系统信息
func SystemInfo() {
	// 检索的 Pacman 日志文件
	var fileName = "/var/log/pacman.log"

	// 最终输出数据小数点后保留的位数
	var precision = 2

	// 获取系统时区
	local, _ := time.LoadLocation("Asia/Shanghai")

	// 获取系统安装时间
	lineText := general.ReadFileLine(fileName, 1)                                           // Pacman 日志文件第一行内容，记录有系统安装时间
	startTimeStrTZ := strings.Split(strings.Split(lineText, "[")[1], "]")[0]                // 从日志文件中截取 2023-03-10T10:49:09+0800 格式的时间字符串
	startTime, _ := time.ParseInLocation("2006-01-02T15:04:05Z0700", startTimeStrTZ, local) // 解析日志文件中截取的时间字符串
	startTimeStr := startTime.Format("2006-01-02 15:04")                                    // 格式化时间

	// 获取当前时间
	currentTimeStr := time.Now().Format("2006-01-02 15:04") // 当前时间

	// 获取初始内核版本
	keyText := general.ReadFileKey(fileName, "installed linux ")                                         // Pacman 日志文件中记录有初始内核版本的行
	firstKernel := strings.Replace(strings.Split(strings.Split(keyText, " (")[1], ")")[0], ".", "-", -1) // 初始内核版本

	// 获取当前内核版本
	unameArgs := []string{"-r"}
	latestKernel, err := general.RunCommandGetResult("uname", unameArgs) // 当前内核版本
	if err != nil {
		color.Danger.Println(err)
	}

	// 计算系统安装天数
	startTimeStamp := startTime.Unix()
	currentTime, _ := time.ParseInLocation("2006-01-02 15:04", currentTimeStr, local)
	currentTimeStamp := currentTime.Unix()
	systemDays := int((currentTimeStamp - startTimeStamp) / 86400) // 系统安装天数

	// 获取系统更新相关数据
	systemUpdateCount, err := general.GetSystemUpdateCount(fileName) // 系统更新次数
	if err != nil {
		color.Danger.Println(err)
	}
	systemUpdateMean := float32(systemUpdateCount) / float32(systemDays) // 系统更新频率

	// 获取内核更新相关数据
	kernelUpdateCount := general.ReadFileCount(fileName, "upgraded linux ") // 内核更新次数
	kernelUpdateMean := float32(systemDays) / float32(kernelUpdateCount)    // 内核更新频率

	// float32 类型数据保留2位小数
	systemUpdateMean, kernelUpdateMean = general.RoundFloat32(systemUpdateMean, precision), general.RoundFloat32(kernelUpdateMean, precision)

	// 从 systemDays 和 systemUpdateCount 中选出最大值作为缩进长度
	member := make([]interface{}, 0, 5)
	member = append(member, systemDays, systemUpdateCount, systemUpdateMean, kernelUpdateCount, kernelUpdateMean)
	length := general.FindFakeMaxLength(member)

	// 获取吉祥物
	repoArgs := []string{""}
	mascot, err := general.RunCommandGetResult("repo-elephant", repoArgs)
	if err != nil {
		color.Danger.Println(err)
	}

	// 输出
	titleFormat := "%27v %-2v %-27v\n"
	dataFormat := "%23v %-2v %v %v\n"
	color.Printf(titleFormat, general.FgCyanText("[", startTimeStr, "]"), "--", general.FgCyanText("[", currentTimeStr, "]"))
	color.Printf(titleFormat, general.FgMagentaText(firstKernel), "--", general.FgMagentaText(latestKernel))
	color.Printf(
		dataFormat,
		general.PrimaryText("系统使用时长"),
		"--",
		general.FgYellowText(color.Sprintf("%-*.*v", length, precision+1, systemDays)),
		general.SecondaryText("天"),
	)
	color.Printf(
		dataFormat,
		general.PrimaryText("系统更新次数"),
		"--",
		general.FgYellowText(color.Sprintf("%-*.*v", length, precision+1, systemUpdateCount)),
		general.SecondaryText("次"),
	)
	color.Printf(
		dataFormat,
		general.PrimaryText("系统更新频率"),
		"--",
		general.FgYellowText(color.Sprintf("%-*.*v", length, precision+1, systemUpdateMean)),
		general.SecondaryText("次/天"),
	)
	color.Printf(
		dataFormat,
		general.PrimaryText("内核更新次数"),
		"--",
		general.FgYellowText(color.Sprintf("%-*.*v", length, precision+1, kernelUpdateCount)),
		general.SecondaryText("次"),
	)
	color.Printf(
		dataFormat,
		general.PrimaryText("内核更新频率"),
		"--",
		general.FgYellowText(color.Sprintf("%-*.*v", length, precision+1, kernelUpdateMean)),
		general.SecondaryText("天/次"),
	)
	color.Println(general.SuccessText(mascot))
}
