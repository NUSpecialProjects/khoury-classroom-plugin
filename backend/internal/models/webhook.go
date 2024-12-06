package models

// Structs for responses from webhook events

type WebHookPRComment struct {
	Comment WebHookComment `json:"comment"`
	//Repository string  `json:"repository.path"`
}

type WebHookComment struct {
	AuthorAssociation string `json:"author_association"`
	Author            string `json:"user.name"`
	Body              string `json:"body"`
}
