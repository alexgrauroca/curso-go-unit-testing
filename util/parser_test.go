package util

import (
	"catching-pokemons/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserPokemonSuccess(t *testing.T) {
	expected := models.Pokemon{}
	response := models.PokeApiPokemonResponse{}

	// From util/samples.go
	ReadTestSampleJson(t, "api_response.json", &expected)
	ReadTestSampleJson(t, "pokeapi_response.json", &response)

	parsedPokemon, err := ParsePokemon(response)
	assert.NoError(t, err)
	assert.Equal(t, expected, parsedPokemon)
}

func TestParserPokemonErrNotFoundPokemonType(t *testing.T) {
	response := models.PokeApiPokemonResponse{}

	// From util/samples.go
	ReadTestSampleJson(t, "pokeapi_response_type_not_found.json", &response)

	_, err := ParsePokemon(response)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrNotFoundPokemonType.Error())
}

func TestParserPokemonErrNotFoundPokemonTypeName(t *testing.T) {
	response := models.PokeApiPokemonResponse{}

	// From util/samples.go
	ReadTestSampleJson(t, "pokeapi_response_type_name_not_found.json", &response)

	_, err := ParsePokemon(response)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrNotFoundPokemonTypeName.Error())
}
