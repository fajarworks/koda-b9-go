package pkg

import (
	"fmt"
	"io"
	"os"
)

func OpenAndReadFile(path string) (string, error) {
	defer func() {

		if err := recover(); err != nil {
			fmt.Println("panic recovered :", err)
		}
	}()

	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	cnt, err := io.ReadAll(file)
	if err != nil {
		panic("cannot read content of file")
	}

	// fmt.Println(string(cnt))

	return string(cnt), nil
}
