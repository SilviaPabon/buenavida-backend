package utils

import(
  "fmt"
  "errors"
  "testing"
  "github.com/stretchr/testify/require"
)

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
