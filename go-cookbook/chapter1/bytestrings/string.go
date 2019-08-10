package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString 展示了一系列在字符串中进行查询的方法
func SearchString() {
	s := "this is a test"

	// 返回 true 表明包含子串
	fmt.Println(strings.Contains(s, "this"))

	// 返回 true 表明包含子串中的任何一字符a或b或c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// 返回 true 表明以该子串开头
	fmt.Println(strings.HasPrefix(s, "this"))

	// 返回 true 表明以该子串结尾
	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString 展示了一系列修改字符串的方法
func ModifyString() {
	s := "simple string"

	// 输出 [simple string]
	fmt.Println(strings.Split(s, " "))

	// 输出 "Simple String"
	fmt.Println(strings.Title(s))

	// 输出 "simple string" 会移除头部和尾部的空白
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader 演示了如何快速创建一个字符串的io.Reader接口
func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	// 在标准输出上打印 s
	io.Copy(os.Stdout, r)
}
