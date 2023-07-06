package tracker

import (
	"fmt"
	"sync"

	"github.com/Checkmarx/kics/internal/constants"
	"github.com/Checkmarx/kics/pkg/model"
)

// CITracker contains information of how many queries were loaded and executed
// and how many files were found and executed

var (
	trackerMu sync.Mutex
)

type CITracker struct {
	ExecutingQueries   int
	ExecutedQueries    int
	FoundFiles         int
	FailedSimilarityID int
	LoadedQueries      int
	ParsedFiles        int
	ScanSecrets        int
	ScanPaths          int
	lines              int
	FoundCountLines    int
	ParsedCountLines   int
	IgnoreCountLines   int
	Version            model.Version
}

// NewTracker will create a new instance of a tracker with the number of lines to display in results output
// number of lines can not be smaller than 1
func NewTracker(previewLines int) (*CITracker, error) {
	if previewLines < constants.MinimumPreviewLines || previewLines > constants.MaximumPreviewLines {
		return &CITracker{},
			fmt.Errorf("output lines minimum is %v and maximum is %v", constants.MinimumPreviewLines, constants.MaximumPreviewLines)
	}
	return &CITracker{
		lines: previewLines,
	}, nil
}

// GetOutputLines returns the number of lines to display in results output
func (c *CITracker) GetOutputLines() int {
	return c.lines
}

// TrackQueryLoad adds a loaded query
func (c *CITracker) TrackQueryLoad(queryAggregation int) {
	c.LoadedQueries += queryAggregation
}

// TrackQueryExecuting adds a executing queries
func (c *CITracker) TrackQueryExecuting(queryAggregation int) {
	c.ExecutingQueries += queryAggregation
}

// TrackQueryExecution adds a query executed
func (c *CITracker) TrackQueryExecution(queryAggregation int) {
	trackerMu.Lock()
	defer trackerMu.Unlock()
	c.ExecutedQueries += queryAggregation
}

// TrackFileFound adds a found file to be scanned
func (c *CITracker) TrackFileFound() {
	c.FoundFiles++
}

// TrackFileParse adds a successful parsed file to be scanned
func (c *CITracker) TrackFileParse() {
	c.ParsedFiles++
}

// FailedDetectLine - queries that fail to detect line are counted as failed to execute queries
func (c *CITracker) FailedDetectLine() {
	c.ExecutedQueries--
}

// FailedComputeSimilarityID - queries that failed to compute similarity ID
func (c *CITracker) FailedComputeSimilarityID() {
	c.FailedSimilarityID++
}

// TrackScanSecret - add to secrets scanned
func (c *CITracker) TrackScanSecret() {
	c.ScanSecrets++
}

// TrackScanPath - paths to preform scan
func (c *CITracker) TrackScanPath() {
	c.ScanPaths++
}

// TrackVersion - information if current version is latest
func (c *CITracker) TrackVersion(retrievedVersion model.Version) {
	c.Version = retrievedVersion
}

// TrackFileFoundCountLines - information about the lines of the scanned files
func (c *CITracker) TrackFileFoundCountLines(countLines int) {
	c.FoundCountLines += countLines
}

// TrackFileParseCountLines - information about the lines of the parsed files
func (c *CITracker) TrackFileParseCountLines(countLines int) {
	c.ParsedCountLines += countLines
}

// TrackFileIgnoreCountLines - information about the lines ignored of the parsed files
func (c *CITracker) TrackFileIgnoreCountLines(countLines int) {
	c.IgnoreCountLines += countLines
}
