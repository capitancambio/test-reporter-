package reporter

import (
	"fmt"
	"time"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

//GoTestCompatibleReporter is useful for testing purposes
type GoTestCompatibleReporter struct {
	Config config.GinkgoConfigType
}

func NewGoTestCompatibleReporter() *GoTestCompatibleReporter {
	return &GoTestCompatibleReporter{}
}

func createTestName(texts []string) string {
	return texts[len(texts)-2]
}

func (reporter *GoTestCompatibleReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
}

func (reporter *GoTestCompatibleReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
	//reporter.BeforeSuiteSummary = setupSummary
}

func (reporter *GoTestCompatibleReporter) SpecWillRun(specSummary *types.SpecSummary) {
	fmt.Println("=== RUN  ", createTestName(specSummary.ComponentTexts))
}

func (reporter *GoTestCompatibleReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	var header string = "--- "

	postfix := "\n" + specSummary.CapturedOutput
	if specSummary.Passed() {
		header += "PASS"
	} else if specSummary.Failed() {
		header += "FAIL"
	} else if specSummary.Panicked() {
		header += "PANIC"
	} else if specSummary.Skipped() {
		header += "SKIP"
	} else if specSummary.Pending() {
		header += "PENDING"
	}
	fmt.Printf("%s: %s (%0.2f)s\n",
		header,
		createTestName(specSummary.ComponentTexts),
		(float64(specSummary.RunTime) / float64(time.Second)))
	if postfix != "" {
		fmt.Println(postfix)
	}
}

func (reporter *GoTestCompatibleReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (reporter *GoTestCompatibleReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
}