package utils

import(
  "fmt"
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
