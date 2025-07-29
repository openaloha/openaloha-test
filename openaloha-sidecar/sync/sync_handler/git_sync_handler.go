package synchandler

import "openaloha.io/openaloha/openaloha-sidecar/sync/constant"

// GitSyncHandler is the handler for the git sync service
type GitSyncHandler struct {

}

// Init is the method to initialize code
func (h *GitSyncHandler) Init() error {
	h.GitClone()
	return nil
}

// Refresh is the method to refresh code
func (h *GitSyncHandler) Refresh() error {
	h.GitPull()
	return nil
}

// IsSupport is the method to check support the syncType
func (h *GitSyncHandler) IsSupport(syncType string) bool {
	return constant.SYNC_TYPE_GIT == syncType
}


// GitClone is the method to git clone
func (h *GitSyncHandler) GitClone() error {
	return nil
}

// GitPull is the method to git pull
func (h *GitSyncHandler) GitPull() error {
	return nil
}