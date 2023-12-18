package task12

func MakeSetFromSequence(sequence []string) map[string]bool {
	set := make(map[string]bool)
	for _, elem := range sequence {
		set[elem] = true
	}
	return set
}
