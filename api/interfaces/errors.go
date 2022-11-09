package interfaces

import (
	"errors"
)

var ErrAlreadyInFavorites = errors.New("Product is already on favorites")

var ErrNotInFavorites = errors.New("Product is not in favorites")
