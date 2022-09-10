package handlers

import (
	"context"
	"drawflow/drawflow"
	"drawflow/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func GetProgramById(w http.ResponseWriter, r *http.Request) {
	dg, cancel := helpers.GetDgraphClient() 
	defer cancel()

	queryId := make(map[string]string)
	queryId["$id"] = chi.URLParam(r, "id")
	fmt.Println("QUERYPARAMSID: ", r.URL.Query().Get("id"))
	ctx := context.Background()

	const q = `query getProgramById($id: string){
		program(func: uid($id)) {
			uid
			name
			description
			drawflow
		}
	}`

	resp, err := dg.NewTxn().QueryWithVars(ctx, q, queryId)
	if err != nil {
		helpers.LogHttpError("Error trying to execute query", w, err, http.StatusInternalServerError)
		return
	}

	var respJson drawflow.JsonProgram
	json.Unmarshal(resp.Json, &respJson)
	if err != nil {
		helpers.LogHttpError("Error trying to unmarshal data from query", w, err, http.StatusInternalServerError)
		return	
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respJson)
}