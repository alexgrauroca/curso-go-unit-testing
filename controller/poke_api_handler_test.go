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
