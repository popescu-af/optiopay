package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/popescu-af/saas-y/pkg/log"

	"github.com/popescu-af/optiopay/services/main-svc/pkg/exports"
)

// HTTPWrapper decorates the APIs with from/to HTTP code.
type HTTPWrapper struct {
	api exports.API
}

// NewHTTPWrapper creates an HTTP wrapper for the service API.
func NewHTTPWrapper(api exports.API) *HTTPWrapper {
	return &HTTPWrapper{api: api}
}

func encodeJSONResponse(i interface{}, status *int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if status != nil {
		w.WriteHeader(*status)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(i)
}

func parseIntParameter(param string) (int64, error) {
	return strconv.ParseInt(param, 10, 64)
}

func parseUintParameter(param string) (uint64, error) {
	return strconv.ParseUint(param, 10, 64)
}

func parseFloatParameter(param string) (float64, error) {
	return strconv.ParseFloat(param, 64)
}

// Paths lists the paths that the API serves.
func (h *HTTPWrapper) Paths() Paths {
	return Paths{
		{
			strings.ToUpper("POST"),
			"/add",
			h.AddEmployee,
		},
		{
			strings.ToUpper("POST"),
			"/remove",
			h.RemoveEmployee,
		},
		{
			strings.ToUpper("GET"),
			"/manager",
			h.Manager,
		},
		{
			strings.ToUpper("GET"),
			"/hierarchy",
			h.Hierarchy,
		},
	}
}

// AddEmployee HTTP wrapper.
func (h *HTTPWrapper) AddEmployee(w http.ResponseWriter, r *http.Request) {
	// Body
	body := &exports.AddInfo{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.ErrorCtx("decoding input failed", log.Context{"error": err})
		return
	}

	// Call implementation
	err := h.api.AddEmployee(body)
	if err != nil {
		writeErrorToHTTPResponse(err, w)
		log.ErrorCtx("call to implementation failed", log.Context{"error": err})
		return
	}
}

// RemoveEmployee HTTP wrapper.
func (h *HTTPWrapper) RemoveEmployee(w http.ResponseWriter, r *http.Request) {
	// Body
	body := &exports.RemoveInfo{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.ErrorCtx("decoding input failed", log.Context{"error": err})
		return
	}

	// Call implementation
	err := h.api.RemoveEmployee(body)
	if err != nil {
		writeErrorToHTTPResponse(err, w)
		log.ErrorCtx("call to implementation failed", log.Context{"error": err})
		return
	}
}

// Manager HTTP wrapper.
func (h *HTTPWrapper) Manager(w http.ResponseWriter, r *http.Request) {
	// Header params
	firstEmployee := r.Header.Get("first_employee")

	secondEmployee := r.Header.Get("second_employee")

	// Call implementation
	result, err := h.api.Manager(firstEmployee, secondEmployee)
	if err != nil {
		writeErrorToHTTPResponse(err, w)
		log.ErrorCtx("call to implementation failed", log.Context{"error": err})
		return
	}

	encodeJSONResponse(result, nil, w)
}

// Hierarchy HTTP wrapper.
func (h *HTTPWrapper) Hierarchy(w http.ResponseWriter, r *http.Request) {

	// Call implementation
	result, err := h.api.Hierarchy()
	if err != nil {
		writeErrorToHTTPResponse(err, w)
		log.ErrorCtx("call to implementation failed", log.Context{"error": err})
		return
	}

	encodeJSONResponse(result, nil, w)
}
