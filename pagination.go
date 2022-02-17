package common

import "strconv"

// ProcessPaginationInput parses the user input for limit and page query parameters
func ProcessPaginationInput(l string, p string) (int, int, error) {

	// Convert to integer to process further
	intL, err := strconv.Atoi(l)
	if err != nil {
		intL = DEFAULT_ITEMS
	}
	intP, err := strconv.Atoi(p)
	if err != nil {
		intP = DEFAULT_ITEMS
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
