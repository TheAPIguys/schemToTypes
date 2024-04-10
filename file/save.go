package file

import (
	"fmt"
	"os"
	"strings"
)

func SaveGoFile(filepath string, name string, text string) {
	// get current directory
	if filepath == "" {
		filepath, _ = os.Getwd()
	}

	if strings.Contains(filepath, "/") {
		pathsArray := strings.Split(text, "/")
		filepath = pathsArray[len(pathsArray)-1]
	}
	// try to save the file to the path with the name .go and the text inside it
	file, err := os.Create(filepath + "/" + name + ".go")
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
