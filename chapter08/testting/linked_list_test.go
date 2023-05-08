package testting

import (
	"github.com/hongjun500/GoLang-master/chapter08/collection"
	"log"
	"testing"
)

func TestLinkedList(t *testing.T) {
	javaList := collection.List[string]{
		Val:  "Java",
		Next: nil,
	}
	empty := &collection.List[int64]{}
	javaList.AddElement("Go")
	log.Print(empty)
}
