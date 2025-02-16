package local

import (
	"fmt"
	"github.com/kubefirst/kubefirst/internal/reports"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func NewCommandConnect() *cobra.Command {
	connectCmd := &cobra.Command{
		Use:   "connect",
		Short: "connect will open all Kubefirst services port forwards",
		Long: "connect opens all Kubefirst service ports for local connection, it makes the services available to" +
			"allow local request to the deployed services",
		RunE: runConnect,
	}

	return connectCmd
}

func runConnect(cmd *cobra.Command, args []string) error {
	log.Info().Msg("opening Port Forward for console...")

	// style UI with local URLs
	fmt.Println(reports.StyleMessage(reports.LocalConnectSummary()))

	log.Info().Msg("Kubefirst port forward done")
	log.Info().Msg("hanging port forwards until ctrl+c is called")

	return nil
}
