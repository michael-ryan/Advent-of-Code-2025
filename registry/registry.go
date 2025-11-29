package registry

import "mike-pr.com/AdventOfCode2025/common"

type SolverFunc func() *common.Answer

var solvers = make(map[int]SolverFunc)

func Register(day int, fn SolverFunc) {
	solvers[day] = fn
}

func All() map[int]SolverFunc {
	return solvers
}
