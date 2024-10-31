package models

import "github.com/google/go-github/github"

type FileTreeNode struct {
	Status string
	Entry  github.TreeEntry
}
