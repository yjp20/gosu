package crypto

import (
  "crypto/rand"
  "base64"
)

func GenerateRandomBytes(n int) ([]byte, error) {
  b := make([]byte, n)
  _, err := rand.Read(b)
  if err != nil {
    return nil, err
  }
  return b, nil
}

func GenerateRandomString(n int) (string, error) {
  b, err := GenerateRandomBytes(n)
  return base64.URLEncoding.EncodeToString(b), err
}
