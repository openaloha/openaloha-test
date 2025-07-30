package synchandler

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"openaloha.io/openaloha/openaloha-sidecar/config"
	"openaloha.io/openaloha/openaloha-sidecar/constant"
	syncfunc "openaloha.io/openaloha/openaloha-sidecar/sync/sync_func"
)

// GitSyncHandler is the handler for the git sync service
type GitSyncHandler struct {
}

// Init is the method to initialize code
func (h *GitSyncHandler) Init(workspace string, syncConfig config.SyncConfig, initFunc syncfunc.InitFunc) error {
	// git clone
	if err := h.GitClone(workspace, syncConfig.Git.Url, syncConfig.Git.Branch); err != nil {
		return err
	}

	// execute init func
	initFunc()

	return nil
}

// Refresh is the method to refresh code
func (h *GitSyncHandler) Refresh(workspace string, syncConfig config.SyncConfig, refreshFunc syncfunc.RefreshFunc) error {
	// parse sync interval
	duration, err := time.ParseDuration(syncConfig.Git.SyncInterval)
	if err != nil {
		return err
	}

	// git pull from repo
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("[%s] 执行任务, 当前时间: %s\n", 
				"git sync handler", 
				time.Now().Format("2006-01-02 15:04:05"))
				// git pull
				if err := h.GitPull(workspace, syncConfig.Git.Branch); err != nil {
					fmt.Println("git pull error, err", err)
				}

				// execute refresh func
				refreshFunc()
		}
	}
	
	return nil
}

// IsSupport is the method to check support the syncType
func (h *GitSyncHandler) IsSupport(syncType string) bool {
	return constant.SYNC_TYPE_GIT == syncType
}

// GitClone is the method to git clone
func (h *GitSyncHandler) GitClone(workspace string, url string, branch string) error {
	// git clone from repo
	if _, err := git.PlainClone(workspace, &git.CloneOptions{
		URL:      url,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		Progress: os.Stdout,
	}); err != nil {
		return err
	}
	return nil
}

// GitPull is the method to git pull
func (h *GitSyncHandler) GitPull(workspace string, branch string) error {
	// open git repo
	repo, err := git.PlainOpen(workspace)
	if err != nil{
		return err
	}

	// Get the working directory for the repository
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// git pull from repo
	if err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress: os.Stdout,
	}); err != nil {
		return err
	}

	return nil
}
