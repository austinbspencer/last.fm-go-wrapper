package lastfm

import (
	"errors"

	"github.com/biter777/countries"
)

type GeoTopArtists struct {
	Artist []Artist `json:"artist"`
	Attr   GeoAttr  `json:"@attr"`
}

type GeoAttr struct {
	Country    string `json:"country"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

func (c *Client) GeoGetTopArtists(country string) (string, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json

	// Check if the country is defined by the ISO 3166-1 country names standard
	thisCountry := countries.ByName(country)

	if thisCountry.String() == "Unknown" {
		return "", errors.New("country param invalid")
	}
	lastfmURL := c.getNoAuthURL("method.geo.gettopartists", "country."+thisCountry.String())
	return lastfmURL, nil
}
