package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	read2()
}

// 按字节读取
func read2() {
	f, err := os.OpenFile("./test.txt", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	chunk := make([]byte, 16)
	var bytes []byte

	for {
		n, err := f.Read(chunk)
		if err == io.EOF {
			break // 读取结束
		}
		if err != nil {
			break // 读取出错
		}
		bytes = append(bytes, chunk[:n]...) // 这里需要取每次具体读取到的数据, 不然会有错误数据
	}

	fmt.Println(string(bytes))
}

// 按行读取文件
func read1() {
	f, err := os.OpenFile("./test.txt", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(buf))
	}
}

// 按字节写入文件
func writeByte() {
	f, err := os.OpenFile("./test.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("123456"))
}

// 按字符串写入文件
func write() {
	f, err := os.OpenFile("./test.txt", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Seek(5, io.SeekStart)
	_, err = f.WriteString("123")
	if err != nil {
		panic(err)
	}
}

// 打开一个文件
func open() {
	f, err := os.Open("./test.txt") // 只读打开
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("open ok")
}

// 创建文件
func crate() {
	fd, err := os.Create("./test.txt")
	if err != nil {
		panic("test.txt create err")
	}
	defer fd.Close()
	fmt.Println("create ok")
}
