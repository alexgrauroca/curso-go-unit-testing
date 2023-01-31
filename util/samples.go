package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ReadTestSample from samples folder and returning it directly without tranforming it
func ReadTestSample(t *testing.T, sampleFile string) []byte {
	// Samples are located into {repo}/samples folder
	body, err := ioutil.ReadFile(fmt.Sprintf("../samples/%s", sampleFile))
	assert.NoError(t, err, fmt.Sprintf("while reading file %s", sampleFile))

	return body
}

// ReadTestSampleJson from samples folder and returning it unmarshaled.
// jsonResponse contains the sample file content unmarshaled.
func ReadTestSampleJson(t *testing.T, sampleFile string, jsonResponse any) {
	body := ReadTestSample(t, sampleFile)
	err := json.Unmarshal(body, jsonResponse)
	assert.NoError(t, err, fmt.Sprintf("while doing unmarshal of file %s", sampleFile))
}
