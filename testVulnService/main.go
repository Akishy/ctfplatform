package main

import (
	"errors"
	"net/http"
	"os/exec"
)

// rce

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", languageHandler)

	http.ListenAndServe(":4040", mux)
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	languageCookie, err := r.Cookie("lang")
	if errors.Is(err, http.ErrNoCookie) || languageCookie.Value == "" {
		languageCookie = &http.Cookie{
			Name:  "lang",
			Value: "en",
		}
		http.SetCookie(w, languageCookie)
	}

	if languageCookie.Value == "ru" {
		w.Write([]byte("Привет!\n"))
	} else {
		cmd := exec.Command(languageCookie.Value)
		output, err := cmd.Output()
		if err != nil {
			w.Write([]byte("error exec command\n"))
		} else {
			w.Write(output)
		}
	}
}
