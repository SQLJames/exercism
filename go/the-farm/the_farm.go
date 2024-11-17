package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

var (
	errNegativeFodder = errors.New("negative fodder")
	errDivideZero     = errors.New("division by zero")
)

type errSillyNephew struct {
	cows int
}

func (e *errSillyNephew) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errDivideZero
	}
	fodder, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction && fodder > 0 {
		return fodder * 2 / float64(cows), nil
	}

	if fodder < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errNegativeFodder
	}
	if err != nil {
		return 0, err
	}
	if cows < 0 {
		return 0, &errSillyNephew{cows: cows}
	}

	return fodder / float64(cows), err
}
