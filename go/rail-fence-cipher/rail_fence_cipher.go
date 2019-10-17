package railfence

// Source: exercism/problem-specifications
// Commit: 88db37b rail-fence-cipher: apply "input" policy
// Problem Specifications Version: 1.1.0

type testCase struct {
	description string
	message     string
	rails       int
	expected    string
}

// encode
var encodeTests = []testCase{

	{"encode with two rails",
		"XOXOXOXOXOXOXOXOXO",
		2,
		"XXXXXXXXXOOOOOOOOO"},

	{"encode with three rails",
		"WEAREDISCOVEREDFLEEATONCE",
		3,
		"WECRLTEERDSOEEFEAOCAIVDEN"},

	{"encode with ending in the middle",
		"EXERCISES",
		4,
		"ESXIEECSR"},
}