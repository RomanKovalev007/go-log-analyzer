package reader

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
)


type Buffer struct{
	size int
}
func NewReader(size int) *Buffer{
	return &Buffer{size: size}
}

var bufPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 4*1024) // 4KB буфер
    },
}

func (buf *Buffer) ReadFile(filePath *string, processor func(string) error) error{
	file, err := os.Open(*filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := bufPool.Get().([]byte)
	defer bufPool.Put(buf)

	reader := bufio.NewReaderSize(file, buf.size)
	var leftover string
	for {
		
		n, err := reader.Read(buffer)
		
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		chunk := leftover + string(buffer[:n])
		lines := strings.Split(chunk, "\n")
		leftover = lines[len(lines)-1]
		for _,line := range lines{
			if err := processor(line); err != nil{
				return err
			}
		}
		if err == io.EOF {
			if leftover != ""{
				if err := processor(leftover); err != nil{
					return err
				}
			}
			break
		}
		
	}
	return nil
}

	