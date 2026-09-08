package pkg

import (
	"io"
	"os"
)

func OpenAndReadFile(path string) (string, error) {

	defer func() {
		recover()
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
