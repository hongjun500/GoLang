// @author hongjun500
// @date 2023/8/11 14:34
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package import_b

import (
	"fmt"

	"github.com/hongjun500/GoLang-master/chapter01/import_a"
)

func ImportB() {
	fmt.Println("this package import_b")
	import_a.ImportA()
}
