package main

import (
	"fmt"
	"net/http"
)

func (app *application) createKeypairHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create keypairs success")
}


func (app *application) getKeypairHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get keypairs success")
}

func (app *application) signHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sign request...")
}

