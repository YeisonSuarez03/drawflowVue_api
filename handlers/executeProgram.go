package handlers

import (
	"bytes"
	"drawflow/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type Execute struct {
	TextFile string `json:"textfile,omitempty"`
}

type FileOutput struct {
	Output string `json:"output,omitempty"`
	Error string `json:"error,omitempty"`
}

func RunPythonFile(cmd *exec.Cmd, w http.ResponseWriter) (string, string){
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil{
		helpers.LogHttpError("Error while trying to execute python file", w, err, http.StatusInternalServerError)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr
}

func ExecuteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var p Execute
	
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		helpers.LogHttpError("Incorrect body structure", w, err, http.StatusBadRequest)
		return
	}
	
	pythonFileName := uuid.New().String() + ".py" 
	// pythonFileName := "program.py" 
	fmt.Println(pythonFileName)
	
	b := []byte(p.TextFile)
    err = ioutil.WriteFile(pythonFileName, b, 0777)
    if err != nil {
		helpers.LogHttpError("Error trying to create python file", w, err, http.StatusInternalServerError)
		return
    }
	
	outStr, errStr := RunPythonFile(exec.Command("python", pythonFileName), w)

	jsonFileResponse := FileOutput{Output: outStr, Error: errStr}
	if err != nil{
		helpers.LogHttpError("Error trying to marshal python file to bytes", w, err, http.StatusInternalServerError)
	}
	
	os.Remove(pythonFileName)
	json.NewEncoder(w).Encode(jsonFileResponse)

}