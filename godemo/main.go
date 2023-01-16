package main

import (
	"checkDefer"
	"checkInterface"
	"checkReflect"
	"checkSlice"
	"checkStruct"
	"checkSwitch"
	"runGoroutine"
)

func main() {
	checkSwitch.RunSwitch()
	checkDefer.RunDefer()
	checkSlice.GetStructList()
	checkSlice.CreateArr()
	checkStruct.StructCheck()
	checkInterface.RunInterface()
	checkReflect.UseReflectType()
	checkReflect.UseReflectValue()
	runGoroutine.Run()
}
