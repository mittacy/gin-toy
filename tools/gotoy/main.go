package main

import (
	"github.com/mittacy/gin-toy/tools/gotoy/internal/project"
	"github.com/spf13/cobra"
	"log"
)

const version = "v0.0.1"

var rootCmd = &cobra.Command{
	Use:     "gotoy",
	Short:   "gotoy: An elegant toolkit for Gin.",
	Long:    `gotoy: An elegant toolkit for Gin.`,
	Version: version,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
