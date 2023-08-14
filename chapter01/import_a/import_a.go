// @author hongjun500
// @date 2023/8/11 14:34
// @tool ThinkPadX1隐士
// Created with 2022.2.2.IntelliJ IDEA
// Description:

package import_a

import (
	"fmt"

	"github.com/hongjun500/GoLang-master/chapter01/import_b"
)

func ImportA() {
	fmt.Println("this package import_a")
	import_b.ImportB()
}
