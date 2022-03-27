package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func getAllPrograms(w http.ResponseWriter, r *http.Request) {
	dgClient := newClient()
	txn := dgClient.NewTxn()
	programs , err := txn.Query(context.Background(), allProgramsQuery)
	if hasError(err, w, "Something went wrong.", http.StatusInternalServerError) {
		return
	}
	
	w.Write(programs.Json)
}

func saveProgram(w http.ResponseWriter, r *http.Request) {
	var body Program
	_ = json.NewDecoder(r.Body).Decode(&body)
	p := Program { Code: body.Code }
	program, err := json.Marshal(p)
	if hasError(err, w, "Invalid request.", http.StatusBadRequest) {
		return
	}

	dgClient := newClient()
	txn := dgClient.NewTxn()
	mutBuyers := &api.Mutation{
		CommitNow: true,
		SetJson: program,
	}
	result, err := txn.Mutate(context.Background(), mutBuyers)
	if hasError(err, w, "Something went wrong.", http.StatusInternalServerError) {
		return
	}

	w.Write(result.Json)
}

func getProgramById(w http.ResponseWriter, r *http.Request) {
	programId := chi.URLParam(r, "uid")
	programByIdQuery := getProgramByIdQuery(programId)
	dgClient := newClient()
	txn := dgClient.NewTxn()

	program, err := txn.Query(context.Background(), programByIdQuery)
	if hasError(err, w, "Couldn't find program", http.StatusNotFound) {
		return
	}

	w.Write(program.Json)
}

func executeProgram(w http.ResponseWriter, r *http.Request) {
	var program Program
	_ = json.NewDecoder(r.Body).Decode(&program)
	code := []byte(program.Code)

	err := os.WriteFile(PYTHON_PROGRAM_FILE, code, 0644)
	if hasError(err, w, "Couldn't write python file", http.StatusInternalServerError) {
		return
	}

	err = godotenv.Load(".env")
	if hasError(err, w, "Couldn't load .env file", http.StatusInternalServerError) {
		return
	}

	pythonExe := os.Getenv("PYTHON_PATH")
	cmd := exec.Command(pythonExe, fmt.Sprintf("./%s", PYTHON_PROGRAM_FILE))
	result, err := cmd.Output()
	if hasError(err, w, "Invalid Python syntax", http.StatusBadRequest) {
		return
	}

	json.NewEncoder(w).Encode(string(result))
}