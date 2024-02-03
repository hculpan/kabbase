package secure

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/hculpan/kabbase/pkg/entities"
	"github.com/o1egl/paseto"
)

func GenerateToken(data entities.User, secretKey []byte, duration time.Duration) (string, error) {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Create a new V2 PASETO token
	pasetoV2 := paseto.NewV2()

	// The key must be exactly 32 bytes long for PASETO V2
	if len(secretKey) != chacha20poly1305.KeySize {
		return "", errors.New("secretKey must be 32 bytes long")
	}

	var token paseto.JSONToken
	token.Set("data", string(jsonData))         // Storing the marshaled JSON data in the token
	token.Expiration = time.Now().Add(duration) // Set the expiration time

	// Generate the token
	return pasetoV2.Encrypt(secretKey, &token, nil)
}

func ValidateToken(token string, secretKey []byte) (string, error) {
	// Ensure the secret key length is correct
	if len(secretKey) != chacha20poly1305.KeySize {
		return "", errors.New("secretKey must be 32 bytes long")
	}

	var newJsonToken paseto.JSONToken
	var newFooter string

	// Create a new V2 PASETO parser
	pasetoV2 := paseto.NewV2()

	// Decrypt and validate the token
	err := pasetoV2.Decrypt(token, secretKey, &newJsonToken, &newFooter)
	if err != nil {
		return "", err
	}

	// Check if the token has expired
	fmt.Printf("Expiration: %s\n", newJsonToken.Expiration.String())
	if newJsonToken.Expiration.Before(time.Now()) {
		return "", errors.New("token has expired")
	}

	// Marshal the token payload back to JSON string
	payload, err := json.Marshal(newJsonToken)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}
