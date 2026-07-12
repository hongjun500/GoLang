// Package import_a 和 package import_b 互相调用，会导致死循环，程序无法运行,下面是错误示范
package import_a

import (
	"fmt"

	"github.com/hongjun500/GoLang/chapter01/import_b"
)

func ImportA() {
	fmt.Println("this package import_a")
	import_b.ImportB()
}
