package sync

import (
	synchandler "openaloha.io/openaloha/openaloha-sidecar/sync/sync_handler"
)

// SyncFacade is the facade for the sync service
type SyncFacade struct{
	syncHandlers []*synchandler.SyncHandler
}

// Sync is the method to init and refresh code
func (f *SyncFacade) Sync() error {
	// init sync handler

	// get sync handler by sync type


	// init code by sync handler


	// refresh code by sync handler

}

func init() {
	// init sync handler
	append(syncHandler)
}

// get sync handler by sync type
func getSyncHandler(syncType string) (*synchandler.SyncHandler, error) {

	return nil, nil
}
