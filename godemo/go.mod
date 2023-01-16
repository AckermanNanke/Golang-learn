module godemo

go 1.18

require (
	checkDefer v0.0.0
	checkInterface v0.0.0
	checkReflect v0.0.0
	checkSlice v0.0.0
	checkStruct v0.0.0
	checkSwitch v0.0.0
	runGoroutine v0.0.0
)

replace (
	checkDefer => ../check-defer
	checkInterface => ../check-interface
	checkReflect => ../check-reflect
	checkSlice => ../check-slice
	checkStruct => ../check-struct
	checkSwitch => ../check-switch
	runGoroutine => ../run-goroutine
)
