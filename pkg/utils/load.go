package utils

import (
	"bufio"
	"errors"
	"os"

	"github.com/asaskevich/govalidator"
)

// LoadSeedsFromFile generates a slice of strings for every URL in a file
func LoadSeedsFromFile(path string) (seeds []string, err error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return seeds, err
	}
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valid := govalidator.IsURL(scanner.Text())
		if valid == false {
			continue
		}

		if len(scanner.Text()) > 0 {
			seeds = append(seeds, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return seeds, err
	}

	if len(seeds) == 0 {
		return seeds, errors.New("no valid URL in the seed list")
	}

	return seeds, nil
}
