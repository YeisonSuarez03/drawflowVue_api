package handlers

import (
	"context"
	"drawflow/drawflow"
	"drawflow/helpers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgraph-io/dgo/v210/protos/api"
)


func PostProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dg, cancel := helpers.GetDgraphClient()
	defer cancel()
	ctx := context.Background()

	fmt.Println(r.Body)

	var p drawflow.Program

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		helpers.LogHttpError("Incorrect body structure", w, err, http.StatusBadRequest)
		return
	}

	fmt.Println(p)
	pb, err := json.Marshal(p)
	if err != nil {
		helpers.LogHttpError("Error trying to marshal databody to a json string", w, err, http.StatusInternalServerError)
		return
	}
	fmt.Println(string(pb))
	txn := dg.NewTxn()

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}
	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		helpers.LogHttpError("Error trying to execute mutation", w, err, http.StatusServiceUnavailable)
		return
	}

	//after all we return the assigned uids
	json.NewEncoder(w).Encode(assigned.Uids)
}
