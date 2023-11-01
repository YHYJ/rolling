/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 11:42:27

Description: 子命令`version`功能函数
*/

package general

import "fmt"

// 程序信息
const (
	Name    string = "Rolling"
	Version string = "v0.5.2"
	Project string = "github.com/yhyj/rolling"
)

// 编译信息
var (
	GitCommitHash string = "unknown"
	BuildTime     string = "unknown"
	BuildBy       string = "unknown"
)

func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", Version)
	if !only {
		programInfo = fmt.Sprintf("%s version: %s\nGit commit hash: %s\nBuilt on: %s\nBuilt by: %s\n", Name, Version, GitCommitHash, BuildTime, BuildBy)
	}
	return programInfo
}