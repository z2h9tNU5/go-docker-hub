package list

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	Count    *int     `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Creator             *int       `json:"creator"`
	Id                  *int       `json:"id"`
	Images              []Image    `json:"images"`
	LastUpdated         time.Time  `json:"last_updated"`
	LastUpdater         *int       `json:"last_updater"`
	LastUpdaterUsername *string    `json:"last_updater_username"`
	Name                *string    `json:"name"`
	Repository          *int       `json:"repository"`
	FullSize            *int       `json:"full_size"`
	V2                  *bool      `json:"v2"`
	TagStatus           *string    `json:"tag_status"`
	TagLastPulled       *time.Time `json:"tag_last_pulled"`
	TagLastPushed       *time.Time `json:"tag_last_pushed"`
	MediaType           *string    `json:"media_type"`
	ContentType         *string    `json:"content_type"`
}

type Image struct {
	Architecture *string    `json:"architecture"`
	Features     *string    `json:"features"`
	Variant      *string    `json:"variant"`
	Digest       *string    `json:"digest"`
	Os           *string    `json:"os"`
	OsFeatures   *string    `json:"os_features"`
	OsVersion    *string    `json:"os_version"`
	Size         *int       `json:"size"`
	Status       *string    `json:"status"`
	LastPulled   *time.Time `json:"last_pulled"`
	LastPushed   *time.Time `json:"last_pushed"`
}

func List(img string) (*Response, error) {
	url := "https://registry.hub.docker.com/v2/repositories/library/" + img + "/tags/"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("response read error: %w", err)
	}

	response := new(Response)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, fmt.Errorf("response unmarshal error: %w", err)
	}

	return response, nil
}
