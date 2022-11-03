package utils

import(
  "fmt"
  "errors"
  "testing"
  "github.com/golang-jwt/jwt/v4"
  "github.com/stretchr/testify/require"
  "github.com/SilviaPabon/buenavida-backend/configs"
  "github.com/SilviaPabon/buenavida-backend/interfaces"
)

// getTestUser helper function to create an example user
func getTestUser() interfaces.User {
  user := interfaces.User{
    ID: 1, 
    Firstname: "Foo", 
    Lastname: "Bar", 
    Email: "foo@bar.com", 
    Password: "secret",
  }

  return user
}

// TestHashPasswordSuccess 
func TestHashPasswordSuccess(t *testing.T){
  c := require.New(t)
  
  // Create hash
  hash, err := HashPassword([]byte("testingpassword"))
  c.NoError(err)

  // Verify hash length
  expectedLength := 60
  c.Equalf(expectedLength, len(string(hash)), fmt.Sprintf("Expexted %d as hash length but got %d", expectedLength, len(string(hash))))

  // Verify hash and passwords are valid
  areEqual := ComparePasswords(hash, []byte("testingpassword"))
  c.Equalf(true, areEqual, fmt.Sprintf("Exptected passwords to be equals"))

  // Verify hash with different password
  areEqual = ComparePasswords(hash, []byte("testingpassword2"))
  c.Equalf(false, areEqual, fmt.Sprintf("Expected passwords to be different"))
}

// TestHashPasswordFail
func TestHashPasswordFail(t *testing.T){
  c := require.New(t)

  // Mock bcrypt function
  originalFunc := bcryptGenerateFromPassword
  bcryptGenerateFromPassword = func([]byte, int) ([]byte, error) {
    // Intentional error
    return []byte{}, errors.New("Oops...") 
  }

  hash, err := HashPassword([]byte("testingpassword"))
  c.NotNilf(err, fmt.Sprintf("Expexted an error but got nil"))
  c.Equalf(0, len(hash), fmt.Sprintf("Expected an empty hash"))

  // Return fucntion to its original value
  bcryptGenerateFromPassword = originalFunc
}

// TestCreateAccessTokenSuccess 
func TestCreateAccessTokenSuccess(t *testing.T){
  c := require.New(t)

  // Create a testing user (interface)
  user := getTestUser()

  // Test function
  signedString, _, err := CreateJWTAccessToken(&user)
  c.NoError(err)
  c.NotEqualf(signedString, "", fmt.Sprintf("Exptected signed string not to be empty"))

  // Get token claims
  claims := &interfaces.JWTCustomClaims{}

  _, err = jwt.ParseWithClaims(signedString, claims, func(token *jwt.Token) (interface{}, error){
    return configs.GetJWTSecret(), nil
  })

  c.NoError(err)
  c.Equalf(user.ID, claims.ID, fmt.Sprintf("Expected token ID to be: %d but got: %d", user.ID, claims.ID))
  c.Equalf(user.Email, claims.Email, fmt.Sprintf("Expected token Email to be: %s but got: %s", user.Email, claims.Email))
  c.Equalf(user.Email, claims.RegisteredClaims.Subject, fmt.Sprintf("Expected %s to be token subject but found %s", user.Email, claims.RegisteredClaims.Subject))

}

// TestCreateRefreshTokenSuccess
func TestCreateRefreshTokenSuccess(t *testing.T){
  c := require.New(t)
  user := getTestUser()

  // Test function
  signedString, _, err := CreateJWTRefreshToken(&user)
  c.NoError(err)
  c.NotEqualf(signedString, "", fmt.Sprintf("Expected signed string not to be empty"))

  // Get token claims
  claims := &interfaces.JWTCustomClaims{}

  _, err = jwt.ParseWithClaims(signedString, claims, func(token *jwt.Token) (interface{}, error){
    return configs.GetJWTSecret(), nil
  })

  c.NoError(err)
  c.Equalf(user.ID, claims.ID, fmt.Sprintf("Expected token ID to be: %d but got: %d", user.ID, claims.ID))
  c.Equalf(user.Email, claims.Email, fmt.Sprintf("Expected token Email to be: %s but got: %s", user.Email, claims.Email))
  c.Equalf(user.Email, claims.RegisteredClaims.Subject, fmt.Sprintf("Expected %s to be token subject but found %s", user.Email, claims.RegisteredClaims.Subject))
}


