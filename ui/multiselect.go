package ui

import "github.com/charmbracelet/huh"

func SelectOptions(options []string) (string, error) {
	var selectOption string

	var huhOptions []huh.Option[string]
	for _, option := range options {
		huhOptions = append(huhOptions, huh.NewOption(option, option))
	}
	huh.NewOption("Option 1", "Option 1")
	var s = huh.NewSelect[string]().Title("Select a value").Description("Need to selected a option").Options(huhOptions...).Value(&selectOption)

	if err := s.Run(); err != nil {
		return "", err
	}
	return selectOption, nil
}
