package users

import (
	"encoding/json"
	"net/http"

	"github.com/openmultiplayer/web/server/src/web"
)

type queryParams struct {
	Sort string `qstring:"sort"`
	Max  int    `qstring:"max"`
	Skip int    `qstring:"skip"`
}

func (s *service) list(w http.ResponseWriter, r *http.Request) {
	var p queryParams
	if !web.ParseQuery(w, r, &p) {
		return
	}

	if p.Sort == "" {
		p.Sort = "desc"
	}
	if p.Max == 0 {
		p.Max = 50
	} else if p.Max > 100 {
		p.Max = 100
	}

	users, err := s.repo.GetUsers(r.Context(), p.Sort, p.Max, p.Skip)
	if err != nil {
		web.StatusInternalServerError(w, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}
