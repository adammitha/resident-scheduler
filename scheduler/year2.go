package scheduler

// Solution is the 'node' in the tree we are searching for our optimal solution
type Solution struct {
	Filled  []Day // List of days with slots filled
	Pending []Day // List of days with slots not yet filled
}

var fnForSol func(Solution) Solution
var fnForLOS func([]Solution) Solution

func nextSolns(sol Solution) []Solution {
	return []Solution{
		sol,
	}
}

// FillSchedule produces a []Day with all of the call slots filled
func FillSchedule(startDate string, residents []Resident) Solution {
	fnForSol = func(sol Solution) Solution {
		if len(sol.Pending) == 0 {
			return sol
		}
		return fnForLOS(nextSolns(sol))
	}

	fnForLOS = func(los []Solution) Solution {
		return los[0]
	}

	return fnForSol(Solution{
		Filled:  make([]Day, 28),
		Pending: make([]Day, 28)},
	)
}
