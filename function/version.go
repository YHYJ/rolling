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
var (
	name    string = "Rolling"
	version string = "v0.2.3"
)

func ProgramInfo() string {
	programInfo := fmt.Sprintf("\033[1m%s\033[0m %s \033[1m%s\033[0m\n", name, "version", version)
	return programInfo
}
