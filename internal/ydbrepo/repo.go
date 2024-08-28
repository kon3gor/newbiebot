package ydbrepo

import "github.com/kon3gor/newbiebot/internal/models"

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) SaveHook(h models.Hook) error {
	return nil
}

func (r *Repo) GetHooks(owner, repo string) ([]models.Hook, error) {
	return nil, nil
}
