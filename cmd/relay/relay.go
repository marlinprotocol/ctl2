package relay

import (
	"github.com/marlinprotocol/ctl2/cmd/relay/cosmos"
	"github.com/marlinprotocol/ctl2/cmd/relay/eth"
	"github.com/marlinprotocol/ctl2/cmd/relay/iris"
	"github.com/spf13/cobra"
)

var RelayCmd = &cobra.Command{
	Use:   "relay",
	Short: "Run relays of various blockchaing",
	Long:  `Allows controlling relays (+abci) for multiple blockchains`,
}

func init() {
	RelayCmd.AddCommand(eth.EthCmd)
	RelayCmd.AddCommand(cosmos.CosmosCmd)
	RelayCmd.AddCommand(iris.IrisCmd)
}
