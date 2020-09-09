package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gosuri/uilive"
	log "github.com/sirupsen/logrus"
)

// LoadSeedsFromFile generates a slice of strings for every URL in a file
func LoadSeedsFromFile(path string) (seeds []string, err error) {
	var totalCount, validCount int
	writer := uilive.New()
	writer.Start()

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return seeds, err
	}
	defer file.Close()

	// Initialize scanner
	scanner := bufio.NewScanner(file)
	log.WithFields(log.Fields{"path": path}).Info("Start reading input list")
	for scanner.Scan() {
		totalCount++
		valid := govalidator.IsURL(scanner.Text())
		if valid == false {
			continue
		}

		if len(scanner.Text()) > 0 {
			seeds = append(seeds, scanner.Text())
			validCount++
		}

		fmt.Fprintf(writer, "\t   Reading input list.. Found %d valid URLs out of %d URLs read so far.\n", validCount, totalCount)
		writer.Flush()
	}
	writer.Stop()

	if err := scanner.Err(); err != nil {
		return seeds, err
	}

	if len(seeds) == 0 {
		return seeds, errors.New("no valid URL in the seed list")
	}

	log.WithFields(log.Fields{"path": path}).Info("Found ", validCount, " valid URLs in the input file, now sending to Isodos..")

	return seeds, nil
}
