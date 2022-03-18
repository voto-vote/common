package common

import (
	"math"
	"strconv"
)

// Defines maximal amount of returned items within an array
const MAX_ITEMS = 100

// Defines the default amount of returned items within an array
const DEFAULT_ITEMS = 20

// Defines the default amount of returned items within an array
const MIN_ITEMS = 10

// ProcessPaginationInput parses the user input for limit and page query parameters
func ProcessPaginationInput(l string, p string) (int, int, error) {

	intL := 0
	intP := 0
	if len(l) == 0 {
		intL = DEFAULT_ITEMS
	}
	if len(p) == 0 {
		intP = 0
	}

	// Convert to integer to process further
	intL, err := strconv.Atoi(l)
	if err != nil {
		intL = DEFAULT_ITEMS
	}
	intP, err = strconv.Atoi(p)
	if err != nil {
		intP = 0
	}
	// Check if no limit is provided
	if intL == 0 || intL < MIN_ITEMS {
		intL = DEFAULT_ITEMS
	}
	if intL > MAX_ITEMS {
		intL = DEFAULT_ITEMS
	}
	return intL, intP, nil
}

// GenerateMetadata generate a links struct which can be used within MetaData struct to return from API
func GenerateMetadata(path string, total int, l int, p int) Links {

	prevPage := p
	if total <= (p*l + l) {
		prevPage = 0
	} else {
		prevPage = p - 1
	}
	nextPage := p
	if (p*l + l) >= total {
		nextPage = p
	} else {
		nextPage = p + 1
	}
	return Links{
		Self: path + "?page=" + strconv.Itoa(p),
		Prev: path + "?page=" + strconv.Itoa(prevPage),
		Next: path + "?page=" + strconv.Itoa(nextPage),
	}
}

// GeneratePagination generates the pagination struct calculated by the input
func GeneratePagination(total int, per_page int, page int, items int) Pagination {

	lastPage := math.Floor(float64(total) / float64(per_page))
	from := page*per_page - per_page

	if from < 0 {
		from = 1
	}

	if from > total {
		from = -1
		page = -1
	}

	to := from + items
	if to > total {
		to = -1
		page = -1
	}

	return Pagination{
		Total:       total,
		PerPage:     per_page,
		CurrentPage: page,
		LastPage:    int(lastPage),
		From:        from,
		To:          to,
		Links:       Links{},
	}
}
