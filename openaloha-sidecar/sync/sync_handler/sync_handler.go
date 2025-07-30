package synchandler

import (
	"openaloha.io/openaloha/openaloha-sidecar/config"
	syncfunc "openaloha.io/openaloha/openaloha-sidecar/sync/sync_func"
)

// InitFunc is the function type for initialization
type InitFunc func() error

// RefreshFunc is the function type for refresh
type RefreshFunc func() error

// SyncHandler is the interface for the sync handler
type SyncHandler interface {
	// Init is the method to initialize code
	Init(workspace string, syncConfig config.SyncConfig, initFunc syncfunc.InitFunc) error

	// Refresh is the method to refresh code
	Refresh(workspace string, syncConfig config.SyncConfig, refreshFunc syncfunc.RefreshFunc) error

	// check support the syncType
	IsSupport(syncType string) bool
}
