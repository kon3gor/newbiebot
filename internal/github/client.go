package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://api.github.com"

	acceptHeaderKey          = "Accept"
	acceptHeaderValue        = "application/vnd.github+json"
	authorizationHeaderKey   = "Authorization"
	authorizationHeaderValue = "Bearer %s"
)

type GithubClient struct {
	c      Config
	client *http.Client
}

func NewClient(c Config) GithubClient {
	return GithubClient{
		c:      c,
		client: http.DefaultClient,
	}
}

func (gc *GithubClient) GetGoodFirstIssues(owner, repo string) ([]Issue, error) {
	path, err := url.JoinPath(baseURL, "repos", owner, repo, "issues")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	gc.fillCommonHeaders(req)

	q := req.URL.Query()
	q.Add("labels", "good first issue")
	req.URL.RawQuery = q.Encode()

	res, err := gc.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data []Issue
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Println(string(b))
		log.Println(res.Status)
		return nil, err
	}

	return data, nil
}

func (gc *GithubClient) fillCommonHeaders(req *http.Request) {
	req.Header.Add(acceptHeaderKey, acceptHeaderValue)
	req.Header.Add(authorizationHeaderKey, fmt.Sprintf(authorizationHeaderValue, gc.c.Token))
}
