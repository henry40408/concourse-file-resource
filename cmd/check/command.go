package main

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/henry40408/concourse-file-resource/internal/checksum"
	"github.com/henry40408/concourse-file-resource/internal/models"

	"github.com/reconquest/hierr-go"
)

type checkRequest struct {
	Source  models.Source
	Version models.Version
}

type checkResponse []models.Version

func checkCommand(stdin io.Reader, stdout io.Writer) error {
	var request checkRequest

	err := json.NewDecoder(stdin).Decode(&request)
	if err != nil {
		return hierr.Errorf(err, "unable to parse JSON from standard input")
	}

	filename := bytes.NewBufferString(request.Source.FileName)
	content := bytes.NewBufferString(request.Source.Content)

	checksum, err := checksum.Calculate(filename, content)
	if err != nil {
		return hierr.Errorf(err, "unable to calculate checksum of filename and content")
	}

	response := checkResponse{}
	if (models.Version{}) != request.Version {
		response = append(response, request.Version)
	}
	response = append(response, models.Version{Checksum: checksum})

	err = json.NewEncoder(stdout).Encode(&response)
	if err != nil {
		return hierr.Errorf(err, "unable to dump JSON to standard output")
	}

	return nil
}
