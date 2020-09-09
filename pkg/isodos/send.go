package isodos

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/dustin/go-humanize"
	log "github.com/sirupsen/logrus"
)

// Response contains the response data of a request to Isodos
type Response struct {
	Digest    string `json:"digest"`
	Submitter string `json:"submitter"`
	UUID      string `json:"uuid"`
}

// URL is a single JSON object that compose the
// NDJSON payload for the batch endpoint
type URL struct {
	URL string `json:"url"`
}

var batchEndpoint = "/api/batch/"

// Send proceed to send a slice of URLs to Isodos
func (c *Client) Send(seeds []string, logging bool) (resp *Response, err error) {
	var payload string
	var seedsCount = len(seeds)

	endpoint, err := url.Parse(c.IsodosURL)
	if err != nil {
		return resp, err
	}
	endpoint.Path = path.Join(endpoint.Path, batchEndpoint+c.Project)

	// Generate the NDJSON payload from the seeds
	if logging {
		log.WithFields(log.Fields{
			"seeds-count": seedsCount,
		}).Info("Generating NDJSON payload")
	}
	for key, seed := range seeds {
		line := URL{URL: seed}
		lineJSON, err := json.Marshal(line)
		if err != nil {
			return resp, err
		}

		if key == len(seeds)-1 {
			payload = payload + string(lineJSON)
		} else {
			payload = payload + string(lineJSON) + "\n"
		}
	}
	payloadSize := humanize.Bytes(uint64(binary.Size([]byte(payload))))
	if logging {
		log.WithFields(log.Fields{
			"seeds-count":  seedsCount,
			"payload-size": payloadSize,
		}).Info("NDJSON payload generated")
	}

	// Create request
	request, err := http.NewRequest("POST", endpoint.String(), bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return resp, err
	}
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	request.Header.Add("Authorization", "LOW "+c.S3Key+":"+c.S3Secret)

	// Execute request
	if logging {
		log.WithFields(log.Fields{
			"seeds-count":  seedsCount,
			"payload-size": payloadSize,
		}).Info("Executing request")
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return resp, err
	}
	if logging {
		log.WithFields(log.Fields{
			"seeds-count":  seedsCount,
			"payload-size": payloadSize,
			"status":       response.Status,
		}).Info("Request executed")
	}

	// Check response status
	if response.StatusCode != http.StatusCreated {
		return resp, errors.New("HTTP request error, status code: " + strconv.Itoa(response.StatusCode))
	}

	// Decode body to Response
	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
