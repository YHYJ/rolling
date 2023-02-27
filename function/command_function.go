/*
File: command_function.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-22 14:15:50

Description: 执行系统命令的函数
*/

package function

import (
	"fmt"
	"os/exec"
	"strings"
)

var (
	_      string
	err    error
	result string
	flag   bool
)

// 运行指定命令并获取命令输出
func RunCommandGetResult(command string, args []string) string {
	_, err = exec.LookPath(command)
	if err == nil {
		// 定义命令
		cmd := exec.Command(command, args...)
		// 执行命令并获取命令输出
		output, _ := cmd.Output()
		// 类型转换
		result = strings.TrimRight(string(output), "\n")
	} else {
		result = fmt.Sprintf("%v: %v\n", "Command not found", command)
	}

	return result
}
