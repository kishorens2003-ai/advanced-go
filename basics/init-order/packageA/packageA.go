package packagea

import "fmt"

var Data []int

func init() {
	Data = []int{1, 2, 3}

	fmt.Println("package A intialized")
}
