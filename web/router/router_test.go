package router

import (
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	r := new(Router)
	if len(r.mappings) != 0 {
		t.Errorf("Initial mappings is not empty")
	}

	handler := func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}

	r.Register("GET", "/", handler)
	if len(r.mappings) != 1 || r.mappings[0].Pattern != "/" {
		t.Errorf("Faild to register new mapping")
	}

	r.Register("GET", "/", handler)
	if len(r.mappings) != 2 || r.mappings[1].Pattern != "/" {
		t.Errorf("Faild to register duplicated mapping")
	}
}
