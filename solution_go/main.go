package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var problems Problems
	err := json.NewDecoder(os.Stdin).Decode(&problems)
	if err != nil {
		log.Fatal(err)
	}

	// pprint(problems)

	solutions := new(Solutions)
	problems.SolveProblem1(solutions)
	problems.SolveProblem2(solutions)
	problems.SolveProblem3(solutions)
	problems.SolveProblem4(solutions)
	problems.SolveProblem5(solutions)
	problems.SolveProblem6(solutions)
	problems.SolveProblem7(solutions)
	problems.SolveProblem8(solutions)
	problems.SolveProblem9(solutions)
	problems.SolveProblem10(solutions)
	problems.SolveProblem11(solutions)
	problems.SolveProblem12(solutions)
	problems.SolveProblem13(solutions)
	problems.SolveProblem14(solutions)
	problems.SolveProblem15(solutions)
	problems.SolveProblem16(solutions)

	err = json.NewEncoder(os.Stdout).Encode(solutions)
	if err != nil {
		log.Fatal(err)
	}
}
