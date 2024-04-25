package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
)

func SelectOptions(title string, description string, options []string) (string, error) {
	var selectOption string

	var huhOptions []huh.Option[string]
	for _, option := range options {
		huhOptions = append(huhOptions, huh.NewOption(option, option))
	}

	var s = huh.NewSelect[string]().Title(title).Description(description).Options(huhOptions...).Value(&selectOption)

	if err := s.Run(); err != nil {
		return "", err
	}
	return selectOption, nil
}

func Form() (programLang string, filename string, output string, name string) {
	var b bool
	var listOfFiles []huh.Option[string]

	// get list of files that are yaml or json on current directory
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".yaml") || strings.Contains(file.Name(), ".json") {
			listOfFiles = append(listOfFiles, huh.NewOption(file.Name(), file.Name()))
		}
	}
	fmt.Println(listOfFiles)
	var listOfProgramLang []huh.Option[string]
	listOfProgramLang = append(listOfProgramLang, huh.NewOption("go", "Golang"))
	listOfProgramLang = append(listOfProgramLang, huh.NewOption("ts", "Typescript"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("Welcome to schemToTypes").Description(`
			This tool helps you to convert json schema or yaml schema to go struct types or typescript types. For quick process use the command line interface. As schemToTypes go -f <file path> or schemToTypes ts -f <file path> -o default is to clipboard or filename if you want a file output -n name of the main struct/type`),
		), huh.NewGroup(
			huh.NewNote().Title("Select the language and the schema file"),
			huh.NewSelect[string]().Options(listOfProgramLang...).Value(&programLang).Title("Select a language"),

			huh.NewSelect[string]().Options(listOfFiles...).Value(&filename).Title("Select a schema file"),
			huh.NewInput().Title("Name of the main struct/type").Value(&name),
			huh.NewConfirm().Affirmative("yes").Negative("no").Title("Do you want to output to a clipboard so you can ").Value(&b),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}
	if !b {
		output = "c"
	} else {
		form.Update(huh.NewInput().Title("Name of the output file").Value(&output))
		if err := form.Run(); err != nil {
			log.Fatal(err)
		}

	}

	return programLang, filename, output, name
}
