package cli

type Selection string

const (
	SelectionDatabase Selection = "database"
	SelectionQuit     Selection = "quit"
	SelectionConfig   Selection = "config"
	SelectionEditor   Selection = "editor"
)
