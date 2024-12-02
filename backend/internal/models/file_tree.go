package models

import "github.com/google/go-github/github"

type FileTreeNode struct {
	Status FileStatus       `json:"status"`
	Entry  github.TreeEntry `json:"entry"`
}

type FileStatus struct {
	Status string      `json:"status"`
	Diff   []LineRange `json:"diff"`
}

type LineRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
