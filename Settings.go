package main

// Errors.
const (
	ErrCLAFolderNotSet        = "Folder is not set"
	ErrCLAFileWithNamesNotSet = "File with Names is not set"
)

type Settings struct {
	FolderPath        string
	FileWithNamesPath string
}
