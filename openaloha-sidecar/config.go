package main

// Config is the configuration for the sidecar
type Config struct {
	// Sync is the configuration for the sync service
	Sync SyncConfig	`json:"sync"`
	// Run is the configuration for the run service
	Run RunConfig `json:"run"`
}

// SyncConfig is the configuration for the sync service
type SyncConfig struct {
	// Type is the type of sync service to use
	Type string `json:"type"`
	// Workspace is the workspace to sync
	Workspace string `json:"workspace"`
	// Git is the configuration for the git sync service
	Git GitConfig `json:"git"`
}

// GitConfig is the configuration for the git sync service
type GitConfig struct {
	// Repo is the repository to sync
	Repo string `json:"repo"`
	// Branch is the branch to sync
	Branch string `json:"branch"`
	// SyncInterval is the interval to sync
	SyncInterval string `json:"sync_interval"`
}

// RunConfig is the configuration for the run service
type RunConfig struct {

}