package lastfm

import "testing"

func TestGeoGetTopArtists(t *testing.T) {
	res, err := client.GeoGetTopArtists("usa")

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Artist); count != 50 {
		t.Fatalf("Got %d recent tracks, wanted 50\n", count)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Country", "United States"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	if len(res.Artist) == 0 {
		t.Error("geo.gettopartists returned an empty array")
		return
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}

func TestGeoGetTopTracks(t *testing.T) {
	res, err := client.GeoGetTopTracks("usa")

	if err != nil {
		t.Error(err)
	}

	if count := len(res.Track); count != 50 {
		t.Fatalf("Got %d recent tracks, wanted 50\n", count)
	}

	var tests = []struct {
		test     string
		expected string
	}{
		{"Country", "United States"},
		{"Page", "1"},
		{"PerPage", "50"},
	}

	for _, test := range tests {
		if output := getStringField(res.Attr, test.test); output != test.expected {
			t.Errorf("Test Failed: %s expected, received: %s", test.expected, output)
		}
	}
}
