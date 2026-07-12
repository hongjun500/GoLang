package import_b

import (
	"fmt"

	"github.com/hongjun500/GoLang/chapter01/import_a"
)

func ImportB() {
	fmt.Println("this package import_b")
	import_a.ImportA()
}
