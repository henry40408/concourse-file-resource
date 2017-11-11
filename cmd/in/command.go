package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/spf13/afero"

	cs "github.com/henry40408/concourse-file-resource/internal/checksum"
	"github.com/henry40408/concourse-file-resource/internal/models"

	"github.com/reconquest/hierr-go"
)

type inRequest struct {
	Source  models.Source  `json:"source"`
	Version models.Version `json:"version"`
}

type inResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}

func inCommand(fs afero.Fs, args []string, stdin io.Reader, stdout io.Writer) error {
	var request inRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from standard input")
	}

	destination, err := validateDestination(fs, args)
	if err != nil {
		return err
	}

	err = putFileInDestination(fs, destination, &request)
	if err != nil {
		return err
	}

	err = respondWithNewVersion(stdout, &request)
	if err != nil {
		return err
	}

	return nil
}

func validateDestination(fs afero.Fs, args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("need at least one argument as destination")
	}

	// NOTE first argument is the path to the program
	destination := args[1]

	destStat, err := fs.Stat(destination)
	if err != nil {
		return "", hierr.Errorf(err, "unable to access destination")
	}

	if !destStat.IsDir() {
		return "", fmt.Errorf("destination is not a directory: %s", destination)
	}

	return destination, nil
}

func putFileInDestination(fs afero.Fs, destination string, request *inRequest) error {
	fullpath := path.Join(destination, request.Source.FileName)

	file, err := fs.OpenFile(fullpath, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		return hierr.Errorf(err, "unable to create file in destination directory")
	}

	file.Write([]byte(request.Source.Content))

	return nil
}

func respondWithNewVersion(stdout io.Writer, request *inRequest) error {
	filename := bytes.NewBufferString(request.Source.FileName)
	content := bytes.NewBufferString(request.Source.Content)

	checksum, err := cs.Calculate(filename, content)
	if err != nil {
		return hierr.Errorf(err, "unable to calculate checksum of filename and content")
	}

	response := inResponse{
		Version:  models.Version{Checksum: checksum},
		Metadata: make([]models.Metadata, 0),
	}

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "unable to dump JSON to standard output")
	}

	return nil
}
