package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"game_character/api/models"
	"game_character/api/responses"
	"game_character/api/utils/formaterror"

	"github.com/gorilla/mux"
)

func (server *Server) CreateItem(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	item := models.Item{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	item.Prepare()
	err = item.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	
	itemCreated, err := item.SaveItem(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, itemCreated.ID))
	responses.JSON(w, http.StatusCreated, itemCreated)
}

func (server *Server) GetItems(w http.ResponseWriter, r *http.Request) {

	item := models.Item{}

	items, err := item.FindAllItems(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, items)
}

func (server *Server) GetItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	item := models.Item{}

	itemReceived, err := item.FindItemByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, itemReceived)
}

func (server *Server) UpdateItem(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the item id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Check if the item exist
	item := models.Item{}
	err = server.DB.Debug().Model(models.Item{}).Where("id = ?", pid).Take(&item).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Item not found"))
		return
	}

	// Read the data itemed
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	itemUpdate := models.Item{}
	err = json.Unmarshal(body, &itemUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemUpdate.Prepare()
	err = itemUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	itemUpdate.ID = item.ID 

	if itemUpdate.Character_code != item.Character_code {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Not Allowed to change character code"))
		return
	}

	itemUpdated, err := itemUpdate.UpdateAItem(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, itemUpdated)
}
