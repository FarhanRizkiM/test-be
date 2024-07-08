package controller

import (
	"encoding/json"
	"fmt"
	"github.com/FarhanRizkiM/test-be/helper"
	"github.com/FarhanRizkiM/test-be/model"
	"net/http"
)

func HomeMakmur(w http.ResponseWriter, r *http.Request) {
	Response := fmt.Sprintf("Makmur AI chooy %s", "8080")
	response, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, "Internal server error: JSON marshaling failed", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	return
}

func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Message = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
}