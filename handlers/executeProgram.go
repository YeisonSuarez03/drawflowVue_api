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

func RunPythonFile(cmd *exec.Cmd, w http.ResponseWriter) (string, string, error){
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr, err
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
	
	outStr, errStr, err := RunPythonFile(exec.Command("python", pythonFileName), w)
	os.Remove(pythonFileName)

	if err != nil{
		helpers.LogHttpError(errStr, w, err, http.StatusInternalServerError)
		return
	}

	jsonFileResponse := FileOutput{Output: outStr, Error: errStr}
	json.NewEncoder(w).Encode(jsonFileResponse)

}