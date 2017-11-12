package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/henry40408/concourse-ssh-resource/pkg/mockio"
	hierr "github.com/reconquest/hierr-go"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

const (
	DESTINATION = "/tmp"
	FILENAME    = "id_rsa"
	CONTENT     = "SSH private key"
)

func TestInCommand(t *testing.T) {
	fs := afero.NewMemMapFs()

	// NOTE make directory before access it
	fs.Mkdir(DESTINATION, 0755)

	args := []string{"in", DESTINATION}

	request := bytes.NewBuffer([]byte(fmt.Sprintf(`{
		"source": {
			"filename": "%s",
			"content": "%s"
		},
		"version": {}
	}`, FILENAME, CONTENT)))

	io, err := mockio.NewMockIO(request)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	err = inCommand(fs, args, io.In, io.Out)
	if !assert.NoError(t, err) {
		return
	}

	fullpath := path.Join(DESTINATION, FILENAME)
	file, err := fs.OpenFile(fullpath, os.O_RDONLY, 0644)
	if !assert.NoError(t, err) {
		return
	}

	content, err := ioutil.ReadAll(file)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, []byte(CONTENT), content)
}

func TestInCommandWithMalformedJSON(t *testing.T) {
	reader := bytes.NewBuffer([]byte(`{`))
	io, err := mockio.NewMockIO(reader)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	err = inCommand(afero.NewMemMapFs(), []string{}, io.In, io.Out)
	if assert.Error(t, err) {
		herr := err.(hierr.Error)
		assert.Equal(t, "unable to parse JSON from standard input", herr.GetMessage())
	}
}

func TestInCommandWithInsufficientArguments(t *testing.T) {
	reader := bytes.NewBuffer([]byte(`{}`))
	io, err := mockio.NewMockIO(reader)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	args := []string{"in"}
	err = inCommand(afero.NewMemMapFs(), args, io.In, io.Out)
	if assert.Error(t, err) {
		assert.Equal(t, "need at least one argument as destination", err.Error())
	}
}

func TestInCommandWithNonExistedDestination(t *testing.T) {
	reader := bytes.NewBuffer([]byte(`{}`))
	io, err := mockio.NewMockIO(reader)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	args := []string{"in", "/pmt"}
	err = inCommand(afero.NewMemMapFs(), args, io.In, io.Out)
	if assert.Error(t, err) {
		herr := err.(hierr.Error)
		assert.Equal(t, "unable to access destination", herr.GetMessage())
	}
}

func TestInCommandWithFileAsDestination(t *testing.T) {
	destination := "/pmt"

	fs := afero.NewMemMapFs()
	fs.Create(destination)

	reader := bytes.NewBuffer([]byte(`{}`))
	io, err := mockio.NewMockIO(reader)
	defer io.Cleanup()
	if !assert.NoError(t, err) {
		return
	}

	args := []string{"in", destination}
	err = inCommand(fs, args, io.In, io.Out)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Sprintf("destination is not a directory: %s", destination), err.Error())
	}
}
