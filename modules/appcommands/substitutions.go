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

// --------------------- BEACON ------------------------------------

func (a *app) beaconCreateSusbstitutions(runnerID string) {
	if a.ProjectID != "beacon" {
		return
	}

	runtimeArgs := a.CreateCmd.getStringToStringFromArgStoreOrDie("runtime-args")
	if len(runtimeArgs) == 0 {
		runtimeArgs["DiscoveryAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("discovery-addr")
		runtimeArgs["HeartbeatAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("heartbeat-addr")
		runtimeArgs["BootstrapAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("bootstrap-addr")
		runtimeArgs["KeystorePath"] = util.ExpandTilde(a.CreateCmd.getStringFromArgStoreOrDie("keystore-path"))
		runtimeArgs["KeystorePassPath"] = util.ExpandTilde(a.CreateCmd.getStringFromArgStoreOrDie("keystore-pass-path"))

		a.CreateCmd.ArgStore["runtime-args"] = runtimeArgs
	}
}

// --------------------- GATEWAY_DOT -------------------------------

func (a *app) gatewayDotCreateSubstitutions(runnerID string) {
	if a.ProjectID != "gateway_dot" {
		return
	}

	runtimeArgs := a.CreateCmd.getStringToStringFromArgStoreOrDie("runtime-args")
	if len(runtimeArgs) == 0 {

		runtimeArgs["GatewayKeystorePath"] = util.ExpandTilde(a.CreateCmd.getStringFromArgStoreOrDie("gateway-keystore-path"))
		runtimeArgs["GatewayListenPort"] = a.CreateCmd.getStringFromArgStoreOrDie("gateway-listen-port")
		runtimeArgs["BridgeDiscoveryAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("bridge-discovery-addr")
		runtimeArgs["BridgePubsubAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("bridge-pubsub-addr")
		runtimeArgs["BridgeBootstrapAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("bridge-bootstrap-addr")
		runtimeArgs["BridgeListenAddr"] = a.CreateCmd.getStringFromArgStoreOrDie("bridge-listen-addr")
		runtimeArgs["BridgeKeystorePath"] = util.ExpandTilde(a.CreateCmd.getStringFromArgStoreOrDie("bridge-keystore-path"))
		runtimeArgs["BridgeKeystorePassPath"] = util.ExpandTilde(a.CreateCmd.getStringFromArgStoreOrDie("bridge-keystore-pass-path"))
		runtimeArgs["BridgeContracts"] = a.CreateCmd.getStringFromArgStoreOrDie("discovery-addr")

		a.CreateCmd.ArgStore["runtime-args"] = runtimeArgs
	}
}
