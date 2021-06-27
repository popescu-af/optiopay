package service

import (
	"net/http"

	"github.com/popescu-af/optiopay/services/main-svc/internal/logic"
)

func writeErrorToHTTPResponse(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}

	switch err.(type) {
	case *logic.NotFoundError:
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write([]byte(err.Error()))
}
