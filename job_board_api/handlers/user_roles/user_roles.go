package user_roles

import (
	"fmt"
	"net/http"
	"github.com/job_board/jbdb"
)

type Handler struct {
	db *jbdb.JobBoardDB
}

// Add user_roles-related handlers here
func New (jbdb *jbdb.JobBoardDB) *Handler {
	return &Handler{
		db: jbdb,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Would return user roles")
}
