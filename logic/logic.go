package logic

import (
	"auth/security/authenticator"
	"fmt"
	"net/http"
)

func DoSomeLogic(w http.ResponseWriter, r *http.Request) {

	authenticator.Authenticate(w, r)

	fmt.Fprintf(w, "This is protected by authentication")
}
