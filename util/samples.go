package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadTestSample(t *testing.T, sampleFile string, jsonResponse any) {
	// Samples are located into {repo}/samples folder
	body, err := ioutil.ReadFile(fmt.Sprintf("../samples/%s", sampleFile))
	assert.NoError(t, err, fmt.Sprintf("while reading file %s", sampleFile))

	err = json.Unmarshal(body, jsonResponse)
	assert.NoError(t, err, fmt.Sprintf("while doing unmarshal of file %s", sampleFile))
}
