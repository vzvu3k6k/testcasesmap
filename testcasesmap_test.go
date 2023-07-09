package testcasesmap_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/vzvu3k6k/testcasesmap"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, testcasesmap.Analyzer, "a")
}
