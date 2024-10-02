package github

type Github interface {
	Ping() error
	ListRepos() []string
}
