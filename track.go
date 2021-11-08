package lastfm

type Tracks struct {
	Track []Track    `json:"track"`
	Attr  ArtistAttr `json:"@attr"`
}

type Track struct {
	Name       string      `json:"name"`
	Playcount  string      `json:"playcount"`
	Listeners  string      `json:"listeners"`
	MBID       string      `json:"mbid"`
	URL        string      `json:"url"`
	Streamable string      `json:"streamable"`
	Artist     AlbumArtist `json:"artist"`
	Image      []Image     `json:"image"`
	Attr       TrackRank   `json:"@attr"`
}

type StreamableTrack struct {
	FullTrack string `json:"fulltrack"`
	Text      string `json:"#text"`
}

type TrackAlbum struct {
	Streamable StreamableTrack `json:"streamable"`
	Duration   int             `json:"duration"`
	URL        string          `json:"url"`
	Name       string          `json:"name"`
	Attr       TrackRankInt    `json:"@attr"`
	Artist     AlbumArtist     `json:"artist"`
}

type TrackRank struct {
	Rank string `json:"rank"`
}

type TrackRankInt struct {
	Rank int `json:"rank"`
}

func (c *Client) TrackGetInfo(track, artist string) (string, error) {
	// http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json
	lastfmURL := c.getNoAuthURL("method.track.getInfo", "track."+track, "artist."+artist)
	return lastfmURL, nil
}