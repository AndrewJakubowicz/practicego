package main

import (
	"fmt"
	"github.com/spyr1014/practicego/03_package/stringutil"
)

// This grabs from the package stringutil. All those files are treated like one file because
// they are in the same package.
func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
}
