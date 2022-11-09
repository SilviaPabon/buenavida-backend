package models

import (
	// "fmt"
	"context"
	"errors"
	"time"

	"github.com/SilviaPabon/buenavida-backend/configs"
	"github.com/google/uuid"
)

var redis = configs.ConnectToRedis()

// SaveRefreshTokenOnRedis save or replace the user refresh-token on the "white list"
func SaveRefreshTokenOnRedis(refreshTokenIdentifier uuid.UUID, owner string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert uuid to string
	identifierStr := refreshTokenIdentifier.String()

	if identifierStr == "" {
		return errors.New("Unable to parse UUID")
	}

	err := redis.Set(ctx, owner, identifierStr, 6*time.Hour).Err()
	return err
}

// GetRefreshTokenFromRedis get user refresh token from the user mail
func GetRefreshTokenFromRedis(owner string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, err := redis.Get(ctx, owner).Result()
	return token, err
}

// DeleteRegreshTokenFromRedis Deletes a token from the user email
func DeleteRefreshTokenFromRedis(owner string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := redis.Del(ctx, owner).Err()
	return err
}
