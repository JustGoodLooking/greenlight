package main

import (
	"fmt"
	"net/http"
)

func (app *application) createKeypairHandler(w http.ResponseWriter, r *http.Request) {
	keypair, pri, err := app.models.Keypair.New(123, "123", "123")

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.keyStore.Set(keypair.ID, pri)


	err = app.writeJSON(w, http.StatusOK, envelope{"movie": keypair}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}


func (app *application) getKeypairHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get keypairs success")
}

func (app *application) signHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sign request...")
}

