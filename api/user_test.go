package api

import (
	"github.com/zakuro9715/AOJRESTAPI/aoj"
	"testing"
  timeutil "github.com/zakuro9715/AOJRESTAPI/util/time"
  "time"
)

func TestNewUserCore(t *testing.T) {
	user := &aoj.User{
		Id:             "id",
		Name:           "name",
		Affliation:     "affliation",
		RegisterDate:   1262304000000,
    LastSubmitDate: 1262304000000,
	}
  userCore := NewUserCore(user)

	if userCore.Id != "id" {
		t.Errorf("got %v, want %v", userCore.Id, "id")
	}
	if userCore.Name != "name" {
		t.Errorf("got %v, want %v", userCore.Name, "name")
	}
	if userCore.Affliation != "affliation" {
		t.Errorf("got %v, want %v", userCore.Affliation, "affliation")
	}

  expectedTime := time.Date(2010, 01, 01, 0, 0, 0, 0, timeutil.JST)
  if userCore.RegisteredAt != expectedTime {
    t.Errorf("got %v, want %v", userCore.RegisteredAt, expectedTime)
  }
  if userCore.LastSubmittedAt != expectedTime {
    t.Errorf("got %v, want %v", userCore.LastSubmittedAt, expectedTime)
  }
}
