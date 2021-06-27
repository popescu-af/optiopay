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
	case *logic.AlreadyFoundError:
		w.WriteHeader(http.StatusConflict)
	case *logic.ArgumentError:
		w.WriteHeader(http.StatusBadRequest)
	case *logic.NotFoundError:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(err.Error()))
}
