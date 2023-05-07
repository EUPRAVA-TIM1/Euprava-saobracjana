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
	jwtService         service.JwtService
}

func NewSaobracajnaHandler(ss service.SaobracajnaService, jwtService service.JwtService) SaobracjanaHandler {
	return saobracjanaHandler{saobracjanaService: ss, jwtService: jwtService}
}

func (s saobracjanaHandler) Init(r *mux.Router) {
	r.StrictSlash(false)
	r.HandleFunc("/saobracajna/Stanice", s.IsAuthorized(s.GetStanice, false)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Gradjanin/Nalozi/{jmbg}", s.IsAuthorized(s.GetGradjaninNaloge, false)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Authorise", s.IsAuthorized(s.Authorise, false)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Authorise", s.IsAuthorized(s.Authorise, true)).Methods("GET", "OPTIONS")

	http.Handle("/", r)
}
func (s saobracjanaHandler) Authorise(w http.ResponseWriter, r *http.Request) {
	jsonResponse("OK", w, http.StatusOK)

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

/*IsAuthorized verifies that token is valid and  checks if it can access worker related stuff if isWorker bool is true
 */
func (s saobracjanaHandler) IsAuthorized(next http.HandlerFunc, isWorker bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := extractBearerToken(r.Header.Get("Authorization"))
		if !s.jwtService.Validate(token) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You don't have permission for this task"))
			return
		}
		if isWorker && !s.jwtService.IsAWorker(token) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("You don't have permission for this task"))
			return
		}
		next(w, r)
	}
}

func extractBearerToken(authHeader string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(authHeader, prefix) {
		return authHeader[len(prefix):]
	}
	return authHeader
}
