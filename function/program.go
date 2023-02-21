/*
File: program.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 11:42:27

Description: 输出程序信息
*/

package function

import "fmt"

// 程序信息
var (
	name    string = "Rolling"
	version string = "v0.1.6"
)

// 输出程序名称
func ProgramInfo() string {
	programInfo := fmt.Sprintf("\033[1m%s\033[0m %s \033[1m%s\033[0m\n", name, "version", version)
	return programInfo
}
