package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/vault-thirteen/errorz"
	v13file "github.com/vault-thirteen/file"
	v13reader "github.com/vault-thirteen/reader"
)

// Errors.
const (
	ErrfFileDoesNotExist = "File does not exist: %v"
)

// Messages.
const (
	MsgfAllFilesExist = "All Files from the List in '%v' do exist."
)

func main() {
	var settings Settings
	var err error
	settings, err = getSettingsFromCommandLine()
	mustBeNoError(err)
	err = checkFiles(settings)
	mustBeNoError(err)
	var msg = fmt.Sprintf(MsgfAllFilesExist, settings.FileWithNamesPath)
	fmt.Print(msg)
}

func mustBeNoError(
	err error,
) {
	if err != nil {
		log.Fatal(err)
	}
}

func getSettingsFromCommandLine() (settings Settings, err error) {
	var clArgs = os.Args
	if len(clArgs) < 2 {
		err = errors.New(ErrCLAFolderNotSet)
		return
	}
	if len(clArgs) < 3 {
		err = errors.New(ErrCLAFileWithNamesNotSet)
		return
	}
	settings.FolderPath = clArgs[1]
	settings.FileWithNamesPath = clArgs[2]
	return
}

func checkFiles(
	settings Settings,
) (err error) {
	var fileNames []string
	fileNames, err = getFileNames(settings.FileWithNamesPath)
	if err != nil {
		return
	}
	var filePath string
	for _, fileName := range fileNames {
		filePath = filepath.Join(settings.FolderPath, fileName)
		err = ensureThatFileExists(filePath)
		if err != nil {
			return
		}
	}
	return
}

func getFileNames(
	fileWithNamesPath string,
) (fileNames []string, err error) {
	var file *os.File
	file, err = os.Open(fileWithNamesPath)
	if err != nil {
		return
	}
	defer func() {
		var derr error
		derr = file.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	var reader = v13reader.NewReader(file)
	var bytes []byte
	bytes, err = reader.ReadLineEndingWithCRLF()
	for {
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}
		filePath := strings.TrimSpace(string(bytes))
		if len(filePath) > 0 {
			fileNames = append(fileNames, filePath)
		}
		bytes, err = reader.ReadLineEndingWithCRLF()

	}
}

func ensureThatFileExists(
	filePath string,
) (err error) {
	var exists bool
	exists, err = v13file.Exists(filePath)
	if err != nil {
		return
	}
	if !exists {
		err = fmt.Errorf(ErrfFileDoesNotExist, filePath)
		return
	}
	return
}
