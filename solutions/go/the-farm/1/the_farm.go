package thefarm

import "fmt"
import "errors"

// TODO: define the 'DivideFood' function
func DivideFood(f FodderCalculator, nbCows int) (float64, error) {
    totalFodder, err := f.FodderAmount(nbCows)

	if err != nil {
        return 0.0, err
    }

    multBy, err := f.FatteningFactor()
    if err != nil {
        return 0.0, err
    }

    return (totalFodder * multBy) / float64(nbCows), nil
}



// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(f FodderCalculator, nbCows int) (float64, error) {
    if nbCows > 0 {
        return DivideFood(f, nbCows)
    } 
	return 0, errors.New("invalid number of cows")
    
}




// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct {
    nbCows int
    mess string
}


func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.nbCows, e.mess)
}


func ValidateNumberOfCows(nbCows int) error {
    if nbCows < 0 {
        return &InvalidCowsError {
            nbCows: nbCows,
            mess: "there are no negative cows",
        }
    } else if nbCows == 0 {
        return &InvalidCowsError {
            nbCows: nbCows,
            mess: "no cows don't need food",
        }
    }
    return nil
}




// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
