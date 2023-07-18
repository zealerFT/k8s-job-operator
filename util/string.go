package util

import (
	"bufio"
	"strings"
)

// StringToSlice
// 按照操作系统自动选择的行结束符来分割字符串，你可以使用bufio包的Scanner
// bufio.NewScanner使用一种基于换行符的扫描方式，可以自动识别\n，\r\n和\r作为行结束符，适用于跨平台的场景
func StringToSlice(args string) []string {
	scanner := bufio.NewScanner(strings.NewReader(args))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
