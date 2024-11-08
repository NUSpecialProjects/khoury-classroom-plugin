package models

import "github.com/google/go-github/github"

type FileTreeNode struct {
	Status FileStatus
	Entry  github.TreeEntry
}

type FileStatus struct {
	Status string
	Diff   []LineRange
}

type LineRange struct {
	Start int
	End   int
}
