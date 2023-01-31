package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

var (
	testGetPokemonIdBasePath = "/pokemon/{id}"
	testValidId              = "pikachu"
	testNotFoundId           = "not-found-id"
	testInvalidId            = ""
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	body := util.ReadTestSample(t, "pokeapi_response.json")
	httpmock.RegisterResponder("GET", GetPokemonFromPokeApiUrl(testValidId), httpmock.NewBytesResponder(200, body))

	pokemon, err := GetPokemonFromPokeApi(testValidId)
	assert.NoError(t, err)

	expected := models.PokeApiPokemonResponse{}
	util.ReadTestSampleJson(t, "pokeapi_response.json", &expected)

	assert.Equal(t, expected, pokemon)
}

func TestGetPokemonFromPokeApiNotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", GetPokemonFromPokeApiUrl(testNotFoundId), httpmock.NewStringResponder(404, ""))

	_, err := GetPokemonFromPokeApi(testNotFoundId)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrPokemonNotFound.Error())
}

func TestGetPokemonFromPokeApiFailure(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", GetPokemonFromPokeApiUrl(testInvalidId), httpmock.NewStringResponder(500, ""))

	_, err := GetPokemonFromPokeApi(testInvalidId)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrPokeApiFailure.Error())
}

func TestMuxGetPokemonSuccess(t *testing.T) {
	r, err := http.NewRequest("GET", testGetPokemonIdBasePath, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	vars := map[string]string{
		"id": testValidId,
	}

	r = mux.SetURLVars(r, vars)
	GetPokemon(w, r)

	expectedPokemon := models.Pokemon{}
	util.ReadTestSampleJson(t, "api_response.json", &expectedPokemon)

	actualPokemon := models.Pokemon{}
	err = json.Unmarshal(w.Body.Bytes(), &actualPokemon)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expectedPokemon, actualPokemon)
}

func TestMuxGetPokemonNotFound(t *testing.T) {
	r, err := http.NewRequest("GET", testGetPokemonIdBasePath, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	vars := map[string]string{
		"id": testNotFoundId,
	}

	r = mux.SetURLVars(r, vars)
	GetPokemon(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestMuxGetPokemonFailure(t *testing.T) {
	r, err := http.NewRequest("GET", testGetPokemonIdBasePath, nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	vars := map[string]string{
		"id": testInvalidId,
	}

	r = mux.SetURLVars(r, vars)
	GetPokemon(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
