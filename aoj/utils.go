package aoj

import (
  "errors"
  "reflect"
  "strings"
)

var ErrTrimFailed = errors.New("Trim failed")

func TrimField(i interface{}, name string) error {
  ip := reflect.ValueOf(i)
  if ip.Kind() != reflect.Ptr {
    return ErrTrimFailed
  }

  ie := ip.Elem()
  if ie.Kind() != reflect.Struct {
    return ErrTrimFailed
  }

  f := ie.FieldByName(name)
  if !(f.CanSet() && f.Kind() == reflect.String) {
    return ErrTrimFailed
  }

  f.SetString(strings.Trim(f.String(), "\n"))
  return nil
}
