package handlers

import (
	"EuprvaSaobracajna/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type SaobracjanaHandler interface {
	Init(r *mux.Router)
}

type saobracjanaHandler struct {
	saobracjanaService service.SaobracajnaService
}

func NewSaobracajnaHandler(ss service.SaobracajnaService) SaobracjanaHandler {
	return saobracjanaHandler{saobracjanaService: ss}
}

func (s saobracjanaHandler) Init(r *mux.Router) {
	r.StrictSlash(false)
	r.HandleFunc("/saobracajna/Stanice", s.GetStanice).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Gradjanin/Nalozi/{jmbg}", s.GetGradjaninNaloge).Methods("GET", "OPTIONS")

	http.Handle("/", r)
}

func (s saobracjanaHandler) GetStanice(w http.ResponseWriter, r *http.Request) {
	stanice, err := s.saobracjanaService.GetStanice()
	if err != nil {
		jsonResponse(err.Error(), w, http.StatusNotFound)
		return
	}
	jsonResponse(stanice, w, http.StatusOK)
}

func (s saobracjanaHandler) GetGradjaninNaloge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]
	nalozi, err := s.saobracjanaService.GetGradjaninPrekrsajneNaloge(jmbg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	jsonResponse(nalozi, w, http.StatusOK)
}

func extractBearerToken(authHeader string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(authHeader, prefix) {
		return authHeader[len(prefix):]
	}
	return authHeader
}

func jsonResponse(object interface{}, w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(object)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status != 0 {
		w.WriteHeader(status)
	}
	_, err = w.Write(resp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
