package controllers

import (
	"arbif/src/config"
	"arbif/src/database"
	"arbif/src/models"
	"arbif/src/repositories"
	"arbif/src/responses"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateLead(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var lead models.Lead
	if err = json.Unmarshal(requestBody, &lead); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = lead.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Conectar()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewLeadsRepository(db)
	if err = repository.Create(lead); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	binData, err := json.Marshal(lead)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	payload := bytes.NewBuffer(binData)
	url := config.BinURL
	contentType := "application/json"
	resp, err := http.Post(url, contentType, payload)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

	var simulation models.Simulation
	simulation.LeadId = lead.Id
	simulation.CreditValue = lead.CreditValue
	simulation.TotalInstallment = lead.TotalInstallment

	if err = simulation.Prepare(); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, simulation)
}
