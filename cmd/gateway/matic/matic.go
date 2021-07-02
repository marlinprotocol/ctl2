package matic

import (
	"github.com/marlinprotocol/ctl2/cmd/gateway/matic/bor"
	"github.com/spf13/cobra"
)

var MaticCmd = &cobra.Command{
	Use:   "matic",
	Short: "run matic chain gateways",
	Long:  `Allows controlling gateways (+bridges) for matic blockchain`,
}

func init() {
	MaticCmd.AddCommand(bor.BorCmd)
}
