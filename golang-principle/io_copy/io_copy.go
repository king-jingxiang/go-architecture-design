package io_copy

import (
	"io"
	"os"
)

// 使用io.Copy拷贝文件
func copyWithIOCopy(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

// 通过缓冲区读取和写入数据
func copyWithBuffer(src, dst string, bufSize int) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	buffer := make([]byte, bufSize)
	for {
		n, err := sourceFile.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		_, err = destinationFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}
	return nil
}

// 使用io.CopyBuffer
func copyWithIOCopyBuffer(src, dst string, bufSize int) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	buffer := make([]byte, bufSize)
	_, err = io.CopyBuffer(destinationFile, sourceFile, buffer)
	return err
}

// 直接读取和写入数据
func copyWithReadAndWrite(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	data, err := io.ReadAll(sourceFile)
	if err != nil {
		return err
	}

	_, err = destinationFile.Write(data)
	return err
}
