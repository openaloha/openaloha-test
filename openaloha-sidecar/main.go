package main

import "openaloha.io/openaloha/openaloha-sidecar/sync"


func main() {
	// parse config from environment variables


	// init workspace


	// start sync job to sync code

}

// parse config from environment variables
func parseConfig() (Config, error) {
	return nil, nil
}

// init workspace
func initWorkspace(workspace string) error {
	return nil
}

// start sync job to sync code
func startSync() error {
	syncFacade := &sync.SyncFacade{}
	return syncFacade.Sync()
}