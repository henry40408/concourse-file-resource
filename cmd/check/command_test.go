package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/icrowley/fake"

	hierr "github.com/reconquest/hierr-go"
	"github.com/stretchr/testify/assert"
)

const (
	FILENAME = "id_rsa"
	CONTENT  = "SSH private key"
	CHECKSUM = "a8e71ecf216c1ff259d8b5dcd19ba59012ede3c770d0734e88dc496c79757f33"
)

func TestCheckCommand(t *testing.T) {
	var response checkResponse

	in := bytes.NewBufferString(fmt.Sprintf(`{
		"source": {
			"filename": "%s",
			"content": "%s"
		}
	}`, FILENAME, strings.Replace(CONTENT, "\n", "\\n", -1)))
	out := bytes.NewBuffer([]byte{})
	err := checkCommand(in, out)
	if !assert.NoError(t, err) {
		return
	}

	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	if assert.Equal(t, 1, len(response)) {
		assert.Equal(t, CHECKSUM, response[0].Checksum)
	}
}

func TestCheckCommandWithVersion(t *testing.T) {
	var response checkResponse

	randomString := fake.WordsN(3)
	in := bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"source": {
			"filename": "%s",
			"content": "%s"
		},
		"version": { "sha256sum": "%s" }
	}`, FILENAME, CONTENT, randomString)))
	out := bytes.NewBuffer([]byte{})
	err := checkCommand(in, out)
	if !assert.NoError(t, err) {
		return
	}

	err = json.NewDecoder(out).Decode(&response)
	if !assert.NoError(t, err) {
		return
	}

	if assert.Equal(t, 2, len(response)) {
		assert.Equal(t, randomString, response[0].Checksum)
		assert.Equal(t, CHECKSUM, response[1].Checksum)
	}
}

func TestCheckCommandWithMalformedJSON(t *testing.T) {
	in := bytes.NewBufferString(`{`)
	out := bytes.NewBuffer([]byte{})
	err := checkCommand(in, out)
	herr := err.(hierr.Error)
	assert.Equal(t, herr.GetMessage(), "unable to parse JSON from standard input")
}
