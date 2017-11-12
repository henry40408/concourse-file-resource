package main

import (
	"bytes"
	"encoding/json"
	"io"

	cs "github.com/henry40408/concourse-file-resource/internal/checksum"
	"github.com/henry40408/concourse-file-resource/internal/models"

	hierr "github.com/reconquest/hierr-go"
)

type outRequest struct {
	Source models.Source `json:"source"`
	Params models.Params `json:"params"`
}

type outResponse struct {
	Version  models.Version    `json:"version"`
	Metadata []models.Metadata `json:"metadata"`
}

func outCommand(stdin io.Reader, stdout io.Writer) error {
	var request outRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from standard input")
	}

	filename := bytes.NewBufferString(request.Source.FileName)
	content := bytes.NewBufferString(request.Source.Content)
	checksum, err := cs.Calculate(filename, content)
	if err != nil {
		return hierr.Errorf(err, "unable to calculate checksum from filename and content")
	}

	response := outResponse{
		Version:  models.Version{Checksum: checksum},
		Metadata: make([]models.Metadata, 0),
	}
	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "unable to dump JSON to standard output")
	}

	return nil
}
