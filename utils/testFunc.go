package utils

import "github.com/always-waiting/mypkg/pkg/base"

func Test(a int) int {
	return a
}

func TestA(a int) int {
	base.PrintCallFunc()
	return a
}
