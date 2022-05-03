package gcp

import (
	"fmt"
	"net/http"

	"github.com/maito1201/hatebucord"
)

func PostHatebuCommand(w http.ResponseWriter, _ *http.Request) {
	err := hatebucord.PostHatebu()
	if err != nil {
		http.Error(w, fmt.Sprintf("InternalError: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Success")
}
