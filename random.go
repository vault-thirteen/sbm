package sbm

import (
	"github.com/vault-thirteen/random"
)

// Creates a Pair of random Values, which have the Sum eqeal to the Sum
// specified.
func createRandomValuePair(
	valueSum uint,
) (valueLeft uint, valueRight uint, err error) {

	valueLeft, err = random.Uint(0, valueSum)
	if err != nil {
		return
	}

	valueRight = valueSum - valueLeft
	return
}