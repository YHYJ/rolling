/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 11:42:27

Description: 子命令`version`功能函数
*/

package function

import "fmt"

// 程序信息
const (
	Name    = "Rolling"
	Version = "v0.3.9"
	Path    = "github.com/yhyj/rolling"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", Version)
	if !only {
		programInfo = fmt.Sprintf("%s version %s\n", Name, Version)
	}
	return programInfo
}
