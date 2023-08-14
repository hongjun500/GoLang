package logging

import (
	"fmt"
	"log"
	"os"

	"time"
)

/**
 * @author hongjun500
 * @date 2022/9/12
 * @tool ThinkPadX1隐士
 * Created with GoLand 2022.2
 * Description: file.go
 */

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	//返回文件信息结构描述文件。如果出现错误，会返回*PathError
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	/*
		O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
		O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
		O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
		O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
		O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
		O_SYNC   int = syscall.O_SYNC   // 同步IO
		O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时
	*/
	//调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。如果出现错误，则为*PathError。
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

func mkDir() {
	//返回与当前目录对应的根路径名
	dir, _ := os.Getwd()
	//创建对应的目录以及所需的子目录
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
