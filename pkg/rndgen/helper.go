/*
Copyright 2019 Vladislav Dmitriyev.
*/

package rndgen

import (
	"fmt"
)

func buildError(methodName, errorString string) error {
	return fmt.Errorf("error occured during '%s' method execution: %s", methodName, errorString)
}

func generateBytes(length int) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := rander.Read(bytes)
	if err != nil {
		return nil, buildError("generateBytes", err.Error())
	}

	return bytes, nil
}

func computeBitMask(charsPower int) (byte, error) {
	if charsPower == 0 || charsPower > 256 {
		return 0, buildError("computeBitMask", "chars length must be greater than 0 and less than or equal to 256")
	}

	var bitLength byte
	var bitMask byte
	for bits := charsPower - 1; bits != 0; {
		bits = bits >> 1
		bitLength++
	}
	bitMask = 1<<bitLength - 1

	return bitMask, nil
}

// generate method returns random values with specified length using alphabet in argument.
// The algorithm was found at https://stackoverflow.com/a/35615565
func generate(alphabet string, length int) ([]byte, error) {
	alphabetLength := len(alphabet)

	// Compute bitMask
	bitMask, err := computeBitMask(alphabetLength)
	if err != nil {
		return nil, err
	}

	// Compute bufferSize
	bufferSize := length + length/3

	// Create random string
	result := make([]byte, length)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			// Random byte buffer is empty, get a new one
			randomBytes, err = generateBytes(bufferSize)
			if err != nil {
				return nil, err
			}
		}
		// Mask bytes to get an index into the character slice
		if idx := int(randomBytes[j%length] & bitMask); idx < alphabetLength {
			result[i] = alphabet[idx]
			i++
		}
	}

	return result, nil
}
