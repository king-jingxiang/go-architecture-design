package io_copy

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

const (
	src = "source.file"      // 你的源文件路径
	dst = "destination.file" // 你的目标文件路径
)

func TestProfiler(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	Test_copyWithIOCopy(t)

	ch := make(chan int)

	<-ch

}

func Test_copyWithIOCopy(t *testing.T) {
	// 测试io.Copy
	start := time.Now()
	if err := copyWithIOCopy(src, dst); err != nil {
		fmt.Println("Error copying with io.Copy:", err)
	}
	fmt.Println("Time taken with io.Copy:", time.Since(start))
}

func Test_copyWithBuffer(t *testing.T) {
	// 测试带缓冲区的拷贝
	start := time.Now()
	if err := copyWithBuffer(src, dst, 4096); err != nil {
		fmt.Println("Error copying with buffer:", err)
	}
	fmt.Println("Time taken with buffer:", time.Since(start))

}

func Test_copyWithIOCopyBuffer(t *testing.T) {

	// 测试io.CopyBuffer
	start := time.Now()
	if err := copyWithIOCopyBuffer(src, dst, 4096); err != nil {
		fmt.Println("Error copying with io.CopyBuffer:", err)
	}
	fmt.Println("Time taken with io.CopyBuffer:", time.Since(start))
}

func Test_copyWithReadAndWrite(t *testing.T) {
	// 测试直接读取和写入
	start := time.Now()
	if err := copyWithReadAndWrite(src, dst); err != nil {
		fmt.Println("Error copying with ReadAndWrite:", err)
	}
	fmt.Println("Time taken with ReadAndWrite:", time.Since(start))
}
