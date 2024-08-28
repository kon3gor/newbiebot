package webhook

type Event struct {
	Owner   string
	Repo    string
	Payload any
}
