package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "11111111111111111Hello World from Go! 👋"
	resp["language"] = "222222222222go"
	resp["cloud"] = "333333333333333Hosted on Vercel! ▲"
	resp["github"] = "https://github.com/riccardogiorato/template-go-vercel/blob/main/api/json.go"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
	} else {
		w.Write(jsonResp)
	}
}
