package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserPokemonSuccess(t *testing.T) {
	expected := models.Pokemon{}
	response := models.PokeApiPokemonResponse{}

	readSample(t, "api_response.json", &expected)
	readSample(t, "pokeapi_response.json", &response)

	parsedPokemon, err := ParsePokemon(response)
	assert.NoError(t, err)
	assert.Equal(t, expected, parsedPokemon)
}

func TestParserPokemonErrNotFoundPokemonType(t *testing.T) {
	response := models.PokeApiPokemonResponse{}

	readSample(t, "pokeapi_response_type_not_found.json", &response)

	_, err := ParsePokemon(response)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrNotFoundPokemonType.Error())
}

func TestParserPokemonErrNotFoundPokemonTypeName(t *testing.T) {
	response := models.PokeApiPokemonResponse{}

	readSample(t, "pokeapi_response_type_name_not_found.json", &response)

	_, err := ParsePokemon(response)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrNotFoundPokemonTypeName.Error())
}

func readSample(t *testing.T, sampleFile string, jsonResponse any) {
	body, err := ioutil.ReadFile(fmt.Sprintf("samples/%s", sampleFile))
	assert.NoError(t, err, fmt.Sprintf("while reading file %s", sampleFile))

	err = json.Unmarshal(body, jsonResponse)
	assert.NoError(t, err, fmt.Sprintf("while doing unmarshal of file %s", sampleFile))
}
