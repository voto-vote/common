package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPagination checks if pagination numbers are always correct
func TestPagination(t *testing.T) {

	// Test pagination with parameters in different scenarios
	// Normal scenario
	t.Log("Normal scenario")
	p := GeneratePagination(78, 20, 1, 20)

	assert.Equal(t, p.Total, 78, "Not the correct total value")
	assert.Equal(t, p.PerPage, 20, "Not the correct per_page value")
	assert.Equal(t, p.CurrentPage, 1, "Not the correct current_page value")
	assert.Equal(t, p.LastPage, 4, "Not the correct last_page value")
	assert.Equal(t, p.From, 0, "Not the correct from value")
	assert.Equal(t, p.To, 19, "Not the correct to value")

	t.Log("Less items than per_page scenario")
	// Less items than per_page scenario
	p = GeneratePagination(10, 20, 1, 10)

	assert.Equal(t, 10, p.Total, "Not the correct total value")
	assert.Equal(t, 20, p.PerPage, "Not the correct per_page value")
	assert.Equal(t, 1, p.CurrentPage, "Not the correct current_page value")
	assert.Equal(t, 1, p.LastPage, "Not the correct last_page value")
	assert.Equal(t, 0, p.From, "Not the correct from value")
	assert.Equal(t, 9, p.To, "Not the correct to value")

	t.Log("Invalid page scenario")
	// Invalid page scenario
	p = GeneratePagination(10, 20, 2, 10)

	assert.Equal(t, 10, p.Total, "Not the correct total value")
	assert.Equal(t, 20, p.PerPage, "Not the correct per_page value")
	assert.Equal(t, 1, p.CurrentPage, "Not the correct current_page value")
	assert.Equal(t, 1, p.LastPage, "Not the correct last_page value")
	assert.Equal(t, 0, p.From, "Not the correct from value")
	assert.Equal(t, 9, p.To, "Not the correct to value")

	t.Log("Last page scenario 1")
	// Last page scenario 1
	p = GeneratePagination(21, 30, 1, 10)

	assert.Equal(t, 21, p.Total, "Not the correct total value")
	assert.Equal(t, 30, p.PerPage, "Not the correct per_page value")
	assert.Equal(t, 1, p.CurrentPage, "Not the correct current_page value")
	assert.Equal(t, 1, p.LastPage, "Not the correct last_page value")
	assert.Equal(t, 0, p.From, "Not the correct from value")
	assert.Equal(t, 20, p.To, "Not the correct to value")

	t.Log("Last page scenario 2")
	// Last page scenario 2
	p = GeneratePagination(78, 20, 4, 18)

	assert.Equal(t, 78, p.Total, "Not the correct total value")
	assert.Equal(t, 20, p.PerPage, "Not the correct per_page value")
	assert.Equal(t, 4, p.CurrentPage, "Not the correct current_page value")
	assert.Equal(t, 4, p.LastPage, "Not the correct last_page value")
	assert.Equal(t, 60, p.From, "Not the correct from value")
	assert.Equal(t, 77, p.To, "Not the correct to value")

}

func TestMetadataLinks(t *testing.T) {

	// Test pagination with parameters in different scenarios
	// Normal scenario
	t.Log("Normal scenario")
	l := GenerateMetadata("/test", 78, 20, 1)

	assert.Contains(t, l.Self, "1", "Self link is not correct")
	assert.Contains(t, l.Next, "2", "Next link is not correct")
	assert.Contains(t, l.Prev, "1", "Prev link is not correct")

	t.Log("Last page scenario")
	l = GenerateMetadata("/test", 78, 20, 4)

	assert.Contains(t, l.Self, "4", "Self link is not correct")
	assert.Contains(t, l.Next, "4", "Next link is not correct")
	assert.Contains(t, l.Prev, "3", "Prev link is not correct")

	t.Log("Invalid page scenario")
	l = GenerateMetadata("/test", 78, 20, 10)

	assert.Contains(t, l.Self, "4", "Self link is not correct")
	assert.Contains(t, l.Next, "4", "Next link is not correct")
	assert.Contains(t, l.Prev, "3", "Prev link is not correct")

}
