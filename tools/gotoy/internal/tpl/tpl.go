package tpl

import (
	"github.com/mittacy/gin-toy/tools/gotoy/internal/tpl/model"
	"github.com/spf13/cobra"
)

// CmdTpl represents the proto command.
var CmdTpl = &cobra.Command{
	Use:   "tpl",
	Short: "Generate the template files",
	Long:  "Generate the template files.",
	Run:   run,
}

func init() {
	CmdTpl.AddCommand(model.CmdModel)
}

func run(cmd *cobra.Command, args []string) {}
