package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"
	hierr "github.com/reconquest/hierr-go"
	"github.com/stretchr/testify/assert"
)

const (
	FILENAME = "id_rsa"
	CONTENT  = "SSH private key"
	CHECKSUM = "a8e71ecf216c1ff259d8b5dcd19ba59012ede3c770d0734e88dc496c79757f33"
)

func TestOutCommand(t *testing.T) {
	var response outResponse

	request := bytes.NewBufferString(fmt.Sprintf(`{
		"source": {
			"filename": "%s",
			"content": "%s"
		},
		"params": {}
	}`, FILENAME, CONTENT))

	io, err := mockio.NewMockIO(request)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	err = outCommand(io.In, io.Out)
	if !assert.NoError(t, err) {
		return
	}

	io.Out.Seek(0, 0)
	err = json.NewDecoder(io.Out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, CHECKSUM, response.Version.Checksum)
}

func TestOutCommandWithMalformedJSON(t *testing.T) {
	request := bytes.NewBufferString("{")

	io, err := mockio.NewMockIO(request)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	err = outCommand(io.In, io.Out)
	if assert.Error(t, err) {
		herr := err.(hierr.Error)
		assert.Equal(t, "unable to parse JSON from standard input", herr.GetMessage())
	}
}
