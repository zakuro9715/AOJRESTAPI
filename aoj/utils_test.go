package aoj

import (
	"testing"
)

func TestTrimField(t *testing.T) {
  tests := []struct{ Value, want string }{
		{"abc", "abc"},
		{" abc ", " abc "},
		{"\nabc\n", "abc"},
		{"\n abc \n", " abc "},
		{" \nabc\n ", " \nabc\n "},
  }
  for _, tt := range tests {
    if err := TrimField(&tt, "Value"); err != nil {
      t.Errorf("err: %v", err.Error())
    }
    if tt.Value != tt.want {
      t.Errorf("got %v, want %v", tt.Value, tt.want)
    }
  }

  if err := TrimField("", ""); err == nil {
    t.Errorf("want err: %v", ErrTrimFailed)
  }
  if err := TrimField(struct{Value string}{}, "Value"); err == nil {
    t.Errorf("want err: %v", ErrTrimFailed)
  }
  if err := TrimField(&struct{value string}{}, "value"); err == nil {
    t.Errorf("want err: %v", ErrTrimFailed)
  }
  if err := TrimField(&struct{Value int}{}, "Value"); err == nil {
    t.Errorf("want err: %v", ErrTrimFailed)
  }
}
