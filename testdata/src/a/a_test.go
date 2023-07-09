package a

// Used just for brevity of testcode.
// This linter does not care about the type of a testcase item.
type testcase struct{}

// testcases defined with `:=`.
func shortVarDecl() {
	{
		testcases := []testcase{} // want "use map for testcases"
		_ = testcases
	}
	{
		testcases := map[string]testcase{} // ok
		_ = testcases

	}
}

// testcases defined with `var`.
func varDecl() {
	{
		var testcases []testcase // want "use map for testcases"
		_ = testcases
	}
	{
		var testcases map[string]testcase // ok
		_ = testcases
	}
}
