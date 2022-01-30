package main

import (
	"flag"
	"fmt"
	"sort/bubblesort"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	/*flag.Parse()

	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}

	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Read values:", values)

	}*/
	arr := make([]int, 1)
	arr = []int{5, 3, 24, 13}
	bubblesort.BubbleSort(arr)
	fmt.Println(arr)
}

/*
func readValues(infile string) (value []int, err error) {
	file, err := os.Open(infile)
	// 打开的文件有错误
	if err != nil {
		fmt.Println("Failed to open th input file ", infile)
		return nil, err
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values := make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		// 转换字符串数组为字符串
		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil{
			err = err1
			return
		}

		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
	}
	defer file.Close()

	for _, value := range values {
		// 将数字转成字符串
		str := strconv.Itoa(value)
		_, err := file.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}*/
