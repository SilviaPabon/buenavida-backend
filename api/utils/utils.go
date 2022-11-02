package utils

import(
  "time"
  "golang.org/x/crypto/bcrypt"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
  "github.com/golang-jwt/jwt/v4"
)

// HashPassword Bcrypt hash password
func HashPassword(passBytes []byte) ([]byte, error) {
  hash, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)

  if err != nil {
    return []byte{}, err
  }

  return hash, nil

}

// ComparePasswords Bcrypt compare password with its possible has
func ComparePasswords(hash, plain []byte) bool {
  err := bcrypt.CompareHashAndPassword(hash, plain)

  if err != nil{
    return false
  }

  return true
}

// CreateAccessToken jwt create short-live access token
func CreateJWTAccessToken(user *interfaces.User) (string, error){
  claims := interfaces.JWTCustomClaims{
    jwt.RegisteredClaims{
      // Jwt default claims
      // IMPORTANT: This should be replaced to 15-30 minutes 
      // When refresh token functionality is ready
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
      IssuedAt: jwt.NewNumericDate(time.Now()),
      NotBefore: jwt.NewNumericDate(time.Now()),
      Issuer: "Buenavida", 
      Subject: user.Email,
    },
    user.ID, 
    user.Email,
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  signed, err := token.SignedString(configs.GetJWTSecret())

  if err != nil {
    return "", err
  }

  return signed, nil

}

// CreateJWTRefreshToken jwt create "long"-live access token
func CreateJWTRefreshToken(user *interfaces.User) (string, error){
  claims := interfaces.JWTCustomClaims{
    jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
      IssuedAt: jwt.NewNumericDate(time.Now()),
      NotBefore: jwt.NewNumericDate(time.Now()),
      Issuer: "Buenavida", 
      Subject: user.Email,
    }, 
    user.ID, 
    user.Email,
  }
  
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  signed, err := token.SignedString(configs.GetJWTSecret())

  if err != nil {
    return "", err
  }

  return signed, nil
}
