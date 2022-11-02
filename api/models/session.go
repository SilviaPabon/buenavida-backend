package models

import(
  // "fmt"
  "time"
  "context"
  "github.com/SilviaPabon/buenavida-backend/configs"
)

var redis = configs.ConnectToRedis() 

func SaveRefreshTokenOnRedis(refreshToken, owner string) error {
  ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
  defer cancel()

  err := redis.Set(ctx, owner, refreshToken, 12 * time.Hour).Err()
  return err
}
