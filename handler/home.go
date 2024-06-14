package handler

import (
    "net/http"
    "github.com/aidosgal/nextgen/view/home"
)

func HandleHomeShow (w http.ResponseWriter, r *http.Request) error {
    return Render(w, r, home.Index())
}
