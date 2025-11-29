# Advent of Code 2025 Solutions

My solutions to Advent of Code 2025.

I've tried to build a nicer working environment than my attempts in previous years, as it got very tedious to have a main function with 25 lines of:

```Go
day1.Solve()
day2.Solve()
...
```

This ended up with a slightly freaky code generation & dependency injection setup (see `registry/` and `gen.go`), but the end result is that the program will now discover when 

## Running the project

Please don't cheat yourself and use this code to solve the challenges to avoid doing it yourself. I have published this repository for educational purposes, in both the hopes of receiving constructive criticism from others and letting others learn from what I've written.

You will need to supply puzzle input. These are expected in `day<number>/resources/input.txt`. 

From the root directory:
```Bash
go generate ./... # Only needed when the number of day solvers changes
go run .
```

Some solvers may be inefficient and take some time to run. Compute and print only specific days' answer by providing their numbers as a list of arguments. For example, to only compute and solve days 3, 6 and 10, run:
```Bash
go run . 3 6 10
```

## Using this repository as a template for your own solutions

Clone the repo, and jump to a period of time before I implemented my solvers with
```Bash
git checkout clean

# Then consider making this repo your own with:
rm -r .git && git init
```

Optionally, delete the helper code used across multiple days' puzzles in the `common/` directory, though doing so will leave some things broken later in the section. Fixing the errors that causes is left as an exercise to the reader.

### To implement a day's solvers:

```Bash
day=1 # or 2, 3, ...
mkdir -p "day${day}/resources"
touch "day${day}/resources/input.txt"
```

Then paste your puzzle input into the newly-created `input.txt`.

To make a day solver:
```Bash
echo "package day${day}" > "day${day}/day${day}.go"
```

A sample day solver:
```Go
const thisDay = 1 // or 2, 3, ...

func init() { // necessary for automatic day discovery
	registry.Register(thisDay, solve)
}

func solve() *common.Answer {
	reader := common.GetInputFileReader(thisDay)

	parsed, err := common.ParseInput[locations](reader)
	if err != nil {
		panic(err)
	}

	return &common.Answer{
		Part1: part1(parsed),
		Part2: part2(parsed),
	}
}
```

`common.ParseInput` uses [participle](https://github.com/alecthomas/participle) to parse the input file. `locations` in the above snippet is a struct that defines a grammar, which would be used parse the input of [Advent of Code 2024 day 1](https://adventofcode.com/2024/day/1):

```Go
type locations struct {
	Pairs []*pair `parser:"@@+"`
}

type pair struct {
	Left  int `parser:"@Int"`
	Right int `parser:"@Int"`
}
```

The fields of `common.Answer` and thus the return type of `part1`, `part2` functions is `any`. As we aren't going to do any processing with these values beyond printing them, it's really fine to not know their types.

From the project root directory, get this running with:
```Bash
go generate ./... # If you haven't done this since adding/removing a day's solver
go run .
```

#### Testing a day solver

Advent of Code puzzles typically provide a simple puzzle example and worked answer. This can be used as a test case which can greatly help in debugging the more difficult puzzles.

Where `<N>` is the appropriate day number:

Create a `test.txt` file in `day<N>/resources/`. Paste in the puzzle's example input. On day 1 of Advent of Code 2024, this looked like:

```
3   4
4   3
2   5
1   3
3   9
3   3
```

Next, create a `day<N>_test.go` file in the `day<N>/` directory, and populate it with test cases and expected values, for example:
```Go
package dayN

import (
    ...
)

func getParsedInput() *locations {
	filename := path.Join("resources", "test.txt")
	reader, err := os.Open(filename)
	if err != nil {
		panic(errors.Wrap(err, fmt.Sprintf("could not open input file %v", filename)))
	}

	parsed, err := common.ParseInput[locations](reader)
	if err != nil {
		panic(err)
	}

	return parsed
}

func TestPart1(t *testing.T) {
	expected := 11 // from worked example
	actual := part1(getParsedInput())
	if expected != actual {
		t.Errorf("part1() = %v, expected %v", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 31 // from worked example
	actual := part2(getParsedInput())
	if expected != actual {
		t.Errorf("part2() = %v, expected %v", actual, expected)
	}
}
```

Run these tests from the project root with:
```Bash
go test ./...
```

Impement test-driven development by writing these tests first, along with stubs for the `part1` and `part2` functions. Then write your solution and only attempt to run your code on the real puzzle input when the tests pass.
