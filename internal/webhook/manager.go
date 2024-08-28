package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/kon3gor/newbiebot/internal/models"
	"github.com/kon3gor/selo"
)

type Repo interface {
	GetHooks(owner string, repo string) ([]models.Hook, error)
	SaveHook(h models.Hook) error
}

type WebhookManager struct {
	c      Config
	repo   Repo
	client *http.Client
}

func NewManager(c Config) *WebhookManager {
	return &WebhookManager{
		c:      c,
		repo:   selo.Get[Repo](),
		client: http.DefaultClient,
	}
}

func (m WebhookManager) Hook(owner, repo, url string) error {
	//note do I need sync.Pool for this?
	h := models.Hook{
		Owner: owner,
		Repo:  repo,
		URL:   url,
	}
	return m.repo.SaveHook(h)
}

func (m WebhookManager) Broadcast(e Event) error {
	hooks, err := m.repo.GetHooks(e.Owner, e.Repo)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(hooks))
	for _, hook := range hooks {
		go func(h models.Hook) {
			defer wg.Done()
			m.notify(h, e)
		}(hook)
	}

	wg.Wait()

	return nil
}

func (m WebhookManager) notify(hook models.Hook, e Event) {
	body, err := json.Marshal(e.Payload)
	if err != nil {
		panic(err) //todo: replace with channels or just ignore it since i'm a bad developer
	}

	req, err := http.NewRequest(http.MethodPost, hook.URL, bytes.NewReader(body))
	if err != nil {
		panic(err) //todo: replace with channels or just ignore it since i'm a bad developer
	}

	resp, err := m.client.Do(req)
	if err != nil {
		panic(err) //todo: replace with channels or just ignore it since i'm a bad developer
	}

	if resp.StatusCode != http.StatusOK {
		//todo: add logs
	}
}
