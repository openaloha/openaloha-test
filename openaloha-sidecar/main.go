package main

import (
	"flag"
	"fmt"
	"openaloha.io/openaloha/openaloha-sidecar/config"
	"os"

	"openaloha.io/openaloha/openaloha-sidecar/constant"
	"openaloha.io/openaloha/openaloha-sidecar/sync"
)

func main() {
	// parse config from environment variables
	config, err := parseConfig()
	if err != nil {
		fmt.Println("parse config error: ", err)
		return
	}

	// init workspace
	initWorkspace(config.Workspace)

	// start sync job to sync code
	if err := startSync(config); err != nil {
		fmt.Println("start sync job error: ", err)
		return
	}
}

// parse config from environment variables
func parseConfig() (config.Config, error) {
	// init config
	var config config.Config

	// parse config
	// common config
	flag.StringVar(&config.Workspace, "workspace", constant.DEFAULT_WORKSPACE, "workspace")
	// sync config
	flag.StringVar(&config.Sync.Type, "sync.type", constant.SYNC_TYPE_GIT, " sync type")
	flag.StringVar(&config.Sync.Git.Url, "sync.git.url", "", "sync url for git")
	flag.StringVar(&config.Sync.Git.Branch, "sync.git.branch", "", "git branch")
	flag.StringVar(&config.Sync.Git.SyncInterval, "sync.git.syncInterval", "", "git syncInterval")

	// TODO:validate config

	flag.Parse()

	return config, nil
}

// init workspace
func initWorkspace(workspace string) error {
	// clear workspace
	if err := os.RemoveAll(workspace); err != nil {
		return err
	}

	// create workspace
	if err := os.MkdirAll(workspace, 0755); err != nil {
		return err
	}

	return nil
}

// start sync job to sync code
func startSync(config config.Config) error {
	syncFacade := &sync.SyncFacade{
		Config: config,
	}
	return syncFacade.Sync()
}
