package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/YeisonSuarez03/drawflowVue_api/drawflow"
	"github.com/YeisonSuarez03/drawflowVue_api/drawflow/helpers"
)


func GetPrograms(w http.ResponseWriter, r *http.Request) {

	dg, cancel := helpers.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	const q = `query getPrograms(){
		programs(func: type(Program)) {
			uid
			name
			description
		}
	}`

	resp, err := dg.NewTxn().Query(ctx, q)
	if err != nil {
		helpers.LogHttpError("Error trying to execute query", w, err, http.StatusInternalServerError)
		return
	}

	var respJson drawflow.JsonPrograms
	json.Unmarshal(resp.Json, &respJson)
	if err != nil {
		helpers.LogHttpError("Error trying to unmarshal data from query", w, err, http.StatusInternalServerError)
		return	
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respJson)
}