/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"schemToTypes/file"
	"schemToTypes/parser"
	"strings"

	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		filepath, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println("Error getting file path")
			fmt.Print(err)
		}
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("Error getting name")
			fmt.Print(err)
		}
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println("Error getting output")
			fmt.Print(err)

		}
		extension := strings.Split(filepath, ".")[len(strings.Split(filepath, "."))-1]
		fmt.Println("Extension: ", extension)

		if extension != "yml" && extension != "json" {
			fmt.Println("The file must be a yaml or json file")
			return
		}

		fmt.Println("SchemToTypes is generating the go struct for the schema file: ", filepath)
		fmt.Println("The struct name will be: ", name)
		fmt.Println(".....")
		data, err := file.Open(filepath)
		if err != nil {
			fmt.Println("Error opening file")
			fmt.Print(err)
		}
		text, err := parser.ProcessRequest(data, extension, parser.Golang, name)
		if err != nil {
			fmt.Println("Error processing request")
			fmt.Print(err)
			return
		}
		if output == "c" {
			parser.SendToClipboard(text)
		} else {
			file.SaveGoFile("", name, text)
		}
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	goCmd.Flags().StringP("file", "f", "", "A file to parse yaml or json schema")
	goCmd.Flags().StringP("name", "n", "", "The name of the struct")
	goCmd.Flags().StringP("output", "o", "c", "The output file path or c for clipboard")

}
