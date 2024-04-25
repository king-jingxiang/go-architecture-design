package _1

import (
	"fmt"
	"testing"
)

type logFunc func(msg string)

func PrintMsg(msg string) {
	fmt.Println(msg)
}

func decoratorFunc(fun logFunc) logFunc {
	logFun := func(msg string) {
		fmt.Println("---开始记录日志---")
		fun(msg)
		fmt.Println("---日志记录结束---")
	}
	return logFun
}
func TestDecorator_Func(t *testing.T) {
	decoratorPrint := decoratorFunc(PrintMsg)
	decoratorPrint("hello")
}
