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




func (buf *Buffer) ReadFile(filePath *string, processor func(string) error) error{
	var bufPool = sync.Pool{
    New: func() interface{} {
        b := make([]byte, buf.size)
        return &b
    },
}

	file, err := os.Open(*filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, buf.size)

	bufferPtr := bufPool.Get().(*[]byte)
	buffer := *bufferPtr
	defer bufPool.Put(bufferPtr)

	var leftover string
	for {
		n, err := reader.Read(buffer)
		
		if err != nil && err != io.EOF {
			return err
		}

		chunk := leftover + string(buffer[:n])
		lines := strings.Split(chunk, "\n")
		leftover = lines[len(lines)-1]

		for _, line := range lines[:len(lines)-1] {
			if line == "" {
				continue
			}
			if err := processor(line); err != nil {
				return err
			}
		}

		if err == io.EOF {
			if leftover != "" {
				if err := processor(leftover); err != nil {
					return err
				}
			}
			break
		}
	}
	return nil
}

	