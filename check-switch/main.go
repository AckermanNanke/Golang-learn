package checkSwitch

import (
	"fmt"
	"runtime"
)

/*
*	switch循环
 */
func RunSwitch() {
	fmt.Println("\n", "RunSwitch=============================")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	case "windows.":
		fmt.Printf("windows")
		fallthrough
	default:
		fmt.Printf("%s.\n", os)
	}
	fmt.Println("RunSwitch=============================", "\n")
}
