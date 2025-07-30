package synchandler

import "openaloha.io/openaloha/openaloha-sidecar/config"

// SyncHandler is the interface for the sync handler
type SyncHandler interface {
	// Init is the method to initialize code
	Init(workspace string, syncConfig config.SyncConfig) error

	// Refresh is the method to refresh code
	Refresh(syncConfig config.SyncConfig) error

	// check support the syncType
	IsSupport(syncType string) bool
}
