package packageb

import (
	"fmt"
	pkgA "init-order/packageA"
)

var Data []int

func init() {
	Data = append(pkgA.Data, 4, 5, 6)

	fmt.Println("package B intialized")
}
