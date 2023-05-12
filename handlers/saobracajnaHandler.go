package handlers

import (
	"EuprvaSaobracajna/data"
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
	r.HandleFunc("/saobracajna/KradjaVozila", s.IsAuthorized(s.PostKradjaVozila, false)).Methods("POST", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Nalozi", s.IsAuthorized(s.PostNalog, true)).Methods("POST", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Sud/Nalozi", s.IsAuthorized(s.PostSudskiNalog, true)).Methods("POST", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Nalozi/{jmbg}", s.IsAuthorized(s.GetPolicajacNalozi, true)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Nalozi/NotIzvrseni/{jmbg}", s.IsAuthorized(s.GetPolicajacNeIzvrseniNalozi, true)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Sud/Nalozi/{jmbg}", s.IsAuthorized(s.GetPolicajacSudskiNalozi, true)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Sud/Nalozi/Status/{id}", s.IsAuthorized(s.SetSudNalogStatus, false)).Methods("POST", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Provera/Sud/{jmbg}", s.IsAuthorized(s.ProveraOsobeSud, true)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Provera/VozackaDozvola/Mup/{brojVozacke}", s.IsAuthorized(s.ProveraOsobeMup, true)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Policajac/Provera/SaobracjanaDozvola/Mup/{tablica}", s.IsAuthorized(s.ProveraVozilaMup, true)).Methods("GET", "OPTIONS")
	r.HandleFunc("/saobracajna/Nalog/{id}", s.IsAuthorized(s.GetPdfNalog, true)).Methods("GET", "OPTIONS")

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

func (s saobracjanaHandler) PostKradjaVozila(w http.ResponseWriter, r *http.Request) {
	var prijava data.PrijavaKradjeVozila
	err := json.NewDecoder(r.Body).Decode(&prijava)
	if err != nil {
		http.Error(w, "Problem with decoding JSON", http.StatusBadRequest)
		return
	}

	err = s.saobracjanaService.SendKradjaPrijava(prijava)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse("Kradja reported successfully", w, http.StatusOK)
}

func (s saobracjanaHandler) PostNalog(w http.ResponseWriter, r *http.Request) {
	var nalog data.PrekrsajniNalog
	err := json.NewDecoder(r.Body).Decode(&nalog)
	if err != nil {
		http.Error(w, "Problem with decoding JSON", http.StatusBadRequest)
		return
	}
	savedNalog, err := s.saobracjanaService.SaveNalog(nalog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(savedNalog, w, http.StatusOK)
}

func (s saobracjanaHandler) PostSudskiNalog(w http.ResponseWriter, r *http.Request) {
	var nalog data.SudskiNalog
	err := json.NewDecoder(r.Body).Decode(&nalog)
	if err != nil {
		http.Error(w, "Problem with decoding JSON", http.StatusBadRequest)
		return
	}
	savedNalog, err := s.saobracjanaService.SaveSudskiNalog(nalog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonResponse(savedNalog, w, http.StatusOK)
}

func (s saobracjanaHandler) GetPolicajacNalozi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]
	nalozi, err := s.saobracjanaService.GetPolcajacPrekrsajneNaloge(jmbg)
	if err != nil {
		http.Error(w, "There has been error with getting nalozi", http.StatusNotFound)
		return
	}
	jsonResponse(nalozi, w, http.StatusOK)
}

func (s saobracjanaHandler) GetPolicajacNeIzvrseniNalozi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]
	nalozi, err := s.saobracjanaService.GetPolicajacNeIzvrseniNalozi(jmbg)
	if err != nil {
		http.Error(w, "There has been error with getting nalozi", http.StatusNotFound)
		return
	}
	jsonResponse(nalozi, w, http.StatusOK)
}

func (s saobracjanaHandler) GetPolicajacSudskiNalozi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]
	nalozi, err := s.saobracjanaService.GetPolcajacSudskeNaloge(jmbg)
	if err != nil {
		http.Error(w, "There has been error with getting nalozi", http.StatusNotFound)
		return
	}
	jsonResponse(nalozi, w, http.StatusOK)
}

func (s saobracjanaHandler) GetPdfNalog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fileDto, err := s.saobracjanaService.GetPdfNalog(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(fileDto, w, http.StatusOK)
}

func (s saobracjanaHandler) ProveraOsobeSud(w http.ResponseWriter, r *http.Request) {

}

func (s saobracjanaHandler) SetSudNalogStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var status data.SudStatusDTO
	err := json.NewDecoder(r.Body).Decode(&status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.saobracjanaService.UpdateSudNalogStatus(id, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResponse(nil, w, http.StatusOK)
}

func (s saobracjanaHandler) ProveraOsobeMup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brojVozacke := vars["brojVozacke"]
	vozacka, err := s.saobracjanaService.GetVozacka(brojVozacke)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	jsonResponse(vozacka, w, http.StatusOK)
}

func (s saobracjanaHandler) ProveraVozilaMup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tablica := vars["tablica"]
	saobracjana, err := s.saobracjanaService.GetSaobracjana(tablica)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	jsonResponse(saobracjana, w, http.StatusOK)
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
