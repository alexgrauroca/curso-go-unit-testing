package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "pikachu"
	body := util.ReadTestSample(t, "pokeapi_response.json")
	httpmock.RegisterResponder("GET", GetPokemonFromPokeApiUrl(id), httpmock.NewBytesResponder(200, body))

	pokemon, err := GetPokemonFromPokeApi(id)
	assert.NoError(t, err)

	expected := models.PokeApiPokemonResponse{}
	util.ReadTestSampleJson(t, "pokeapi_response.json", &expected)

	assert.Equal(t, expected, pokemon)
}

func TestGetPokemonFromPokeApiNotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "id-not-found"
	httpmock.RegisterResponder("GET", GetPokemonFromPokeApiUrl(id), httpmock.NewStringResponder(404, ""))

	_, err := GetPokemonFromPokeApi(id)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrPokemonNotFound.Error())
}

func TestGetPokemonFromPokeApiFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "pikachu"
	httpmock.RegisterResponder("GET", GetPokemonFromPokeApiUrl(id), httpmock.NewStringResponder(500, ""))

	_, err := GetPokemonFromPokeApi(id)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrPokeApiFailure.Error())
}
