package forms

import (
  "net/url"
  "fmt"

  "github.com/badoux/checkmail"
)

type Form struct {
  url.Values
  Errors map[string]string
}

func New(data url.Values) *Form {
  return &Form {
    data,
    map[string]string{},
  }
}

func (f *Form) Required(fields ...string) {
  for _, field := range fields {
    value := f.Get(field)
    if len(value) == 0 {
      f.Errors[field] = fmt.Sprintf("%s cannot be empty", field)
    }
  }
}

func (f *Form) MustLength(field string, min int, max int) {
  value := f.Get(field)
  if len(value) < min {
    f.Errors[field] = fmt.Sprintf("%s must be longer than %d characters", field, min)
  }
  if len(value) > max {
    f.Errors[field] = fmt.Sprintf("%s cannot be longer than %d characters", field, max)
  }
}

func (f *Form) MustConfirm(original string, confirm string) {
  o := f.Get(original)
  c := f.Get(confirm)
  if o != c {
    f.Errors[confirm] = fmt.Sprintf("%s must be equal to %s", confirm, original)
  }
}

func (f *Form) MustEmail(field string) {
  var err error
  value := f.Get(field)

  err = checkmail.ValidateFormat(value)
  if err != nil {
    f.Errors[field] = fmt.Sprintf("email error: %s", err)
  }
}

func (f *Form) Valid() bool {
  return len(f.Errors) == 0
}
