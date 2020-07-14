package fuzz

// abbreviates the string to the targetLength including ellipsis
func AbbreviateString(input string, targetLength int) string {
	bnoden := input
	if len(input) > targetLength {
		if targetLength > 3 {
			targetLength -= 3
		}
		bnoden = input[0:targetLength] + "..."
	}
	return bnoden
}
