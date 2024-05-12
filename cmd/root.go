/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"schemToTypes/file"
	"schemToTypes/parser"
	"schemToTypes/ui"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "schemToTypes",
	Short: "convert json schema or yaml schema to go struct types or typescript types",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		programLang, filename, output, name := ui.Form()

		fmt.Print("\033[H\033[2J")
		var lang parser.TypeOption
		if programLang == "Golang" {
			lang = parser.Golang
		} else {
			lang = parser.TypeScript
		}

		fmt.Println("generating go struct")
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

		message := "*schemToTypes is generating the go struct for the schema file: " + filename + ".* \n" + " >The struct name will be: " + name + "\n" + "....." + "\n"

		data, err := file.Open(filename)
		if err != nil {
			errorMessage := "Error opening file: " + filename + "\n" + err.Error()
			fmt.Println(errorStyle.Render(errorMessage))
			return
		}
		extension := strings.Split(filename, ".")[len(strings.Split(filename, "."))-1]

		if extension != "yml" && extension != "json" && extension != "yaml" {
			fmt.Println(errorStyle.Render("Error: Invalid file \n The file must be a yaml or json file"))
			return
		}
		text, err := parser.ProcessRequest(data, extension, lang, name)
		if err != nil {
			fmt.Println("Error processing request")
			fmt.Print(err)
			return
		}
		endTime := time.Now()
		in := "# Ⓢ ⓒ ⓗ ⓔ ⓜ Ⓣ ⓞ Ⓣ ⓨ ⓟ ⓔ ⓢ\n" + message + " # Generated Go Struct \n" + "```go \n" + text + "\n" + "```" + "\n" + "---" + "\n Time taken: " + endTime.Sub(startTime).String() + "⏰ \n" + "- [x] Parse file \n" + "- [x] Generated struct\n" + "- [x] Copy to clipboard\n"
		if output == "c" {
			parser.SendToClipboard(text)
			in += "### *** Paste the structs in your editor and Happy coding ***\n"
		} else {
			file.SaveFile("", name, file.Go, text)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.schemToTypes.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("file", "f", "", "input file path")
}
