package models

type PRComment struct {
	Comment Comment `json:"comment"`
	//Repository string  `json:"repository.path"`
}

type Comment struct {
	AuthorAssociation string `json:"author_association"`
	Author            string `json:"user.name"`
	Body              string `json:"body"`
}
