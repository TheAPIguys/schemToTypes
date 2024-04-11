/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"schemToTypes/file"
	"schemToTypes/parser"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// tsCmd represents the ts command
var tsCmd = &cobra.Command{
	Use:   "ts",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// clear console
		fmt.Print("\033[H\033[2J")

		startTime := time.Now()

		var errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#d1d5db")).
			Background(lipgloss.Color("#7f1d1d")).
			BorderBackground(lipgloss.Color("d1d5db")).
			PaddingTop(2).
			PaddingBottom(2).
			PaddingLeft(10).
			PaddingRight(10).
			AlignHorizontal(lipgloss.Center)

		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(errorStyle.Render("Error getting file path"))
			fmt.Print(errorStyle.Render(fmt.Sprintf("%s", err)))
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(errorStyle.Render("Error getting name"))
			fmt.Println(errorStyle.Render(fmt.Sprintf("%s", err)))
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(errorStyle.Render("Error getting output"))
			fmt.Println(errorStyle.Render(fmt.Sprintf("%s", err)))

		}
		extension := strings.Split(filepath, ".")[len(strings.Split(filepath, "."))-1]

		if extension != "yml" && extension != "json" {
			fmt.Println(errorStyle.Render("Error: Invalid file \n The file must be a yaml or json file"))
			return
		}

		message := "*schemToTypes is generating the go struct for the schema file: " + filepath + ".* \n" + " >The struct name will be: " + name + "\n" + "....." + "\n"

		data, err := file.Open(filepath)
		if err != nil {
			errorMessage := "Error opening file: " + filepath + "\n" + err.Error()
			fmt.Println(errorStyle.Render(errorMessage))
			return
		}
		text, err := parser.ProcessRequest(data, extension, parser.TypeScript, name)
		if err != nil {
			fmt.Println("Error processing request")
			fmt.Print(err)
			return
		}
		endTime := time.Now()
		in := "# Ⓢ ⓒ ⓗ ⓔ ⓜ Ⓣ ⓞ Ⓣ ⓨ ⓟ ⓔ ⓢ\n" + message + " # Generated Go Struct \n" + "```typescript \n" + text + "\n" + "```" + "\n" + "---" + "\n Time taken: " + endTime.Sub(startTime).String() + "⏰ \n" + "- [x] Parse file \n" + "- [x] Generated struct\n" + "- [x] Copy to clipboard\n"
		if output == "c" {
			parser.SendToClipboard(text)
			in += "### *** Paste the structs in your editor and Happy coding ***\n"
		} else {
			file.SaveFile("", name, file.Ts, text)
			in += "### *** The struct has been saved to the current directory ***\n"
		}
		out, err := glamour.Render(in, "dark")
		if err != nil {
			fmt.Println("Error rendering markdown")
			fmt.Print(err)
		}
		fmt.Print(out)
	},
}

func init() {
	rootCmd.AddCommand(tsCmd)

	tsCmd.Flags().StringP("file", "f", "", "A file to parse yaml or json schema")
	tsCmd.Flags().StringP("name", "n", "", "The name of the struct")
	tsCmd.Flags().StringP("output", "o", "c", "The output file path or c for clipboard")
}
