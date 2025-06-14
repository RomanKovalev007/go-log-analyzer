package reader

import (
	"bufio"
	"io"
	"os"
	"strings"
)


type Buffer struct{
	size int
}
func NewReader(size int) *Buffer{
	return &Buffer{size: size}
}

func (buf *Buffer) ReadFile(filePath *string, processor func(string) error) error{
	file, err := os.Open(*filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, buf.size)
	var leftover string
	for {
		buffer := make([]byte, buf.size)
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

	