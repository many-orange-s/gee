package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Myrecover() HandlerFunc {
	return func(c *Context) {
		//就是匿名函数立马执行 里面的括号就是往里面要穿的参数
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Println("%s \n\n", trace(message))
				c.String(http.StatusOK, "Internal Server Error")
			}
		}()
	}
}
