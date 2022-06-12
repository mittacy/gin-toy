package tpl

import (
	"github.com/mittacy/gin-toy/tools/gotoy/internal/tpl/api"
	"github.com/mittacy/gin-toy/tools/gotoy/internal/tpl/data"
	"github.com/mittacy/gin-toy/tools/gotoy/internal/tpl/model"
	"github.com/mittacy/gin-toy/tools/gotoy/internal/tpl/service"
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
	CmdTpl.AddCommand(data.CmdData)
	CmdTpl.AddCommand(service.CmdService)
	CmdTpl.AddCommand(api.CmdApi)
}

func run(cmd *cobra.Command, args []string) {}
