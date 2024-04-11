package file

import (
	"fmt"
	"os"
	"strings"
)

type Extension string

const (
	Go Extension = ".go"
	Ts Extension = ".ts"
)

func SaveFile(filepath string, name string, extension Extension, text string) {
	// get current directory
	if filepath == "" {
		filepath, _ = os.Getwd()
	}

	if strings.Contains(filepath, "/") {
		pathsArray := strings.Split(text, "/")
		filepath = pathsArray[len(pathsArray)-1]
	}
	// try to save the file to the path with the name .go and the text inside it
	file, err := os.Create(filepath + "/" + name + string(extension))
	if err != nil {
		fmt.Println("Error creating file")
		fmt.Print(err)
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error writing to file")
		fmt.Print(err)
	}

}
