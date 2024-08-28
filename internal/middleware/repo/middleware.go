package repo

import (
	"github.com/kon3gor/newbiebot/internal/webhook"
	"github.com/kon3gor/newbiebot/internal/ydbrepo"
	"github.com/kon3gor/selo"
)

type RepoProxy struct {
	repo *ydbrepo.Repo
}

func NewProxyRepo() webhook.Repo {
	return &RepoProxy{
		repo: selo.Get[*ydbrepo.Repo](),
	}
}

func (r *RepoProxy) GetHooks(owner string, repo string) ([]webhook.Hook, error) {
	hooks, err := r.repo.GetHooks(owner, repo)
	if err != nil {
		return nil, err
	}

	result := make([]webhook.Hook, 0, len(hooks))
	for _, hook := range hooks {
		result = append(result, webhook.Hook(hook))
	}

	return result, nil
}
func (r *RepoProxy) SaveHook(h webhook.Hook) error {
	return r.repo.SaveHook(ydbrepo.Hook(h))
}
