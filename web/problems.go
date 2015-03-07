package web

import (
  "../api"
  "encoding/json"
  "fmt"
  "net/http"
)

func problemsHandler(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Path[len(ProblemsPath):]
  p, _ := api.FetchProblem(id)
  data, _ := json.Marshal(p)
  fmt.Fprintf(w, string(data))
}
