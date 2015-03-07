package web

import (
	"../api"
	"encoding/json"
	"fmt"
	"net/http"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Path[len(UsersPath):]
  u, _ := api.FetchUser(id)
  data, _ := json.Marshal(u)
  fmt.Fprint(w, string(data))
}
