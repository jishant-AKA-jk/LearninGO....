package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	pointVotcounter := &initialVotes
	return pointVotcounter
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	p := counter
	if counter == nil {
		return 0
	}
	return *p
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var res *ElectionResult
	res = &ElectionResult{Name: candidateName, Votes: votes}
	return res

}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return result.Name + " (" + fmt.Sprint(result.Votes) + ")"
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for name := range results {
		if name == candidate {
			results[name]--
		}
	}

}
