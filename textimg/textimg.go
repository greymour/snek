package textimg

import (
	"fmt"
	"os"
)

type TextImg struct {
	rawdata string
	data    [][]string
}

type TextImgFormatter func(text string) [][]string

func New(filePath string, formatter TextImgFormatter) *TextImg {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Could not read file from path `%s`", filePath)
	}
	fileStr := string(file)
	data := formatter(fileStr)
	return &TextImg{fileStr, data}
}

func (ti *TextImg) Draw(parent [][]string) [][]string {
	for i := 0; i < len(parent); i++ {
		parent[i] = ti.data[i]
	}
	return parent
}
