package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	pokemon, err := GetPokemonFromPokeApi("pikachu")
	assert.NoError(t, err)

	expected := models.PokeApiPokemonResponse{}
	util.ReadTestSample(t, "pokeapi_response.json", &expected)

	assert.Equal(t, expected, pokemon)
}
