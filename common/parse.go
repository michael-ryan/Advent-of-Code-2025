package common

import (
	"fmt"
	"io"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/pkg/errors"
)

// ParseInput will parse the data in a given file into a defined struct.
// T should be a struct type that represents the grammar tagged
// so as to be compatible with the participle library.
// See here for tagging documentation: https://github.com/alecthomas/participle/blob/master/README.md
func ParseInput[T any](input io.Reader) (*T, error) {
	parser, err := participle.Build[T]()
	if err != nil {
		return nil, fmt.Errorf("could not build parser: %w", err)
	}

	parsed, err := parser.Parse("test.txt", input)
	if err != nil {
		return nil, fmt.Errorf("could not parse input: %w", err)
	}

	return parsed, nil
}

func GetInputFileReader(dayNumber int) io.Reader {
	filename := fmt.Sprintf("day%v/input.txt", dayNumber)
	reader, err := os.Open(filename)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("could not open input file %v", filename)))
	}
	return reader
}
