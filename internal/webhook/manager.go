package webhook

import (
	"errors"
	"sync"
)

type Repo interface {
	GetHooks(owner string, repo string) ([]Hook, error)
	SaveHook(h Hook) error
}

type WebhookManager struct {
	c    Config
	repo Repo
}

func NewManager(c Config, repo Repo) WebhookManager {
	return WebhookManager{
		c:    c,
		repo: repo,
	}
}

func (m WebhookManager) Hook(owner, repo, url string) error {
	h := Hook{owner, repo, url}
	return m.repo.SaveHook(h)
}

func (m WebhookManager) Broadcast(owner, repo string) error {
	hooks, err := m.repo.GetHooks(owner, repo)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(hooks))
	for _, hook := range hooks {
		go func() {
			defer wg.Done()
			m.notify(hook)
		}()
	}

	return nil
}

func (m WebhookManager) notify(hook Hook) {
	//todo: make request here
}
