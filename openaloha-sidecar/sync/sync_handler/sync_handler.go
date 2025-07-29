package synchandler

// SyncHandler is the interface for the sync handler
type SyncHandler interface {
	// Init is the method to initialize code
	Init() error
	// Refresh is the method to refresh code
	Refresh() error
	// check support the syncType
	IsSupport(syncType string) bool
}