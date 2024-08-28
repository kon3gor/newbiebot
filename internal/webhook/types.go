package webhook

type Hook struct {
	Owner string
	Repo  string
	URL   string
}

type Event struct {
	Owner   string
	Repo    string
	Payload any
}
