package ydbrepo

type Repo struct {
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) SaveHook(h Hook) error {
	return nil
}

func (r *Repo) GetHooks(owner, repo string) ([]Hook, error) {
	return nil, nil
}
