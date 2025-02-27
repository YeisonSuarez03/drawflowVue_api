package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YeisonSuarez03/drawflowVue_api/drawflow"
)

func LogHttpError(description string, w http.ResponseWriter, errorObtained error, status int) {
	jsonError := drawflow.JsonError{
		Error: errorObtained.Error(), 
		Description: description, 
		Status: status,
	}
	jsonErrorEncoded, err := json.Marshal(jsonError)
	if err != nil{
		fmt.Println("Couldn't parse error to json string: ", err)
		http.Error(w, errorObtained.Error(), status)
		return
	}
	http.Error(w, string(jsonErrorEncoded), status)

}