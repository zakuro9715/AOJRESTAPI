package web

import (
	"encoding/json"
	"fmt"
	"github.com/zakuro9715/AOJRESTAPI/api"
	"net/http"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(UsersPath):]
	u, _ := api.FetchUser(id)
	data, _ := json.Marshal(u)
	fmt.Fprint(w, string(data))
}
