package checkInterface

import (
	"fmt"
)

type runSelf interface {
	log()
}

type runSelfStruct struct {
	name string
	age  int
}

func (r runSelfStruct) log() {
	fmt.Println("自执行接口", r)
}

/*
*	switch循环
 */
func RunInterface() {
	var runM runSelf = runSelfStruct{
		"我自己", 24,
	}
	runM.log()
	fmt.Println("输出runSelfStruct", runM)
}
