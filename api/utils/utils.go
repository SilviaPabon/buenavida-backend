package utils

import(
  "golang.org/x/crypto/bcrypt"
)

func HashPassword(passBytes []byte) ([]byte, error) {
  hash, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)

  if err != nil {
    return []byte{}, err
  }

  return hash, nil

}
