package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	IgnoreBlockStart	=	"GITIGNORE<<<"
	IgnoreBlockEnd		=	"GITIGNORE>>>"
	IgnoreSingleLine	=	"GITIGNORE"
	GitReplace		=	"GITREPLACE"
)

func main() {

	flag := fmt.Sprintf(`(`)
	for i := 1; i < len(os.Args) - 1; i ++ {
		flag = fmt.Sprintf(`%v%v|`, flag, os.Args[i])
	}

	flag = fmt.Sprintf(`%v%v)`, flag, os.Args[len(os.Args)-1])

	lines := SplitLines(GetPipe())

	ignoreStartIndex := -1
	ignoreEndtIndex := -1
	for index := 0; index < len(lines); index ++ {
		line := lines[index]
		reg := regexp.MustCompile(fmt.Sprintf(`(\s*)%v%v`, flag, IgnoreBlockStart))
		if reg.MatchString(line) {
			ignoreStartIndex = index
		} else {
			reg = regexp.MustCompile(fmt.Sprintf(`%v%v$`, flag, IgnoreSingleLine))
			if reg.MatchString(line) {
				lines = append(lines[:index], lines[index+1:]...)
				index --
			}
			reg = regexp.MustCompile(fmt.Sprintf(`^(\s*)%v%v(\s+)with`, flag, GitReplace))
			if reg.MatchString(line) {
				lines = append(lines[:index+1], lines[index+2:]...)
				replacedReg := regexp.MustCompile(fmt.Sprintf(`%v%v(\s+)with(\s+)`, flag, GitReplace))
				lines[index] = replacedReg.ReplaceAllString(line, "")
			}
		}
		reg = regexp.MustCompile(fmt.Sprintf(`(\s*)%v%v`, flag, IgnoreBlockEnd))
		if reg.MatchString(line) {
			ignoreEndtIndex = index
			lines = append(lines[:ignoreStartIndex], lines[ignoreEndtIndex+1:]...)
			index = index - (ignoreEndtIndex - ignoreStartIndex + 1)
		}
	}

	for index := 0; index < len(lines) - 1; index ++ {
		fmt.Println(lines[index])
	}
	if lines[len(lines)-1] != "" {
		fmt.Printf(lines[len(lines)-1])
	}
}

// SplitLines :  按行分割字符串
func SplitLines(str string) []string {
	str = strings.Replace(str, "\r", "\n", -1)
	return strings.Split(str, "\n")
}

// GetPipe :  从管道读取
func GetPipe() string {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("Please input from pipe")
	}
	r := bufio.NewReader(os.Stdin)
	str := ""
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		str += string(buf[:n])
	}
	return str
}

