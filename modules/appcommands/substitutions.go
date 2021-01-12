// THIS IS MESSY SYSTEM. DO NOT USE IT IF YOU DO NOT ABSOLUTELY HAVE TO

package appcommands

import "github.com/marlinprotocol/ctl2/modules/util"

// ----------------- RELAY ETH -------------------------------------

func (a *app) relayEthCreateSubstitutions(runnerID string) {
	if a.ProjectID != "relay_eth" {
		return
	}

	runtimeArgs := a.CreateCmd.getStringToStringFromArgStoreOrDie("runtime-args")
	if len(runtimeArgs) == 0 {
		runtimeArgs["DiscoveryAddrs"] = a.CreateCmd.getStringFromArgStoreOrDie("discovery-addrs")
		runtimeArgs["HeartbeatAddrs"] = a.CreateCmd.getStringFromArgStoreOrDie("heartbeat-addrs")
		runtimeArgs["DataDir"] = util.ExpandTilde(a.CreateCmd.getStringFromArgStoreOrDie("datadir"))
		runtimeArgs["DiscoveryPort"] = a.CreateCmd.getStringFromArgStoreOrDie("discovery-port")
		runtimeArgs["PubsubPort"] = a.CreateCmd.getStringFromArgStoreOrDie("pubsub-port")
		runtimeArgs["Address"] = a.CreateCmd.getStringFromArgStoreOrDie("address")
		runtimeArgs["Name"] = a.CreateCmd.getStringFromArgStoreOrDie("name")
		runtimeArgs["AbciVersion"] = a.CreateCmd.getStringFromArgStoreOrDie("abci-version")
		runtimeArgs["SyncMode"] = a.CreateCmd.getStringFromArgStoreOrDie("sync-mode")

		a.CreateCmd.ArgStore["runtime-args"] = runtimeArgs
	}
}