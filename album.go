package lastfm

import "fmt"

// Album struct
type Album struct {
	Name      string      `json:"name"`
	Playcount int         `json:"playcount"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Artist    AlbumArtist `json:"artist"`
	Image     []Image     `json:"image"`
	Attr      ArtistAttr  `json:"@attr"`
}

// SimpleAlbum ...
type SimpleAlbum struct {
	Name       string  `json:"name"`
	Artist     string  `json:"artist"`
	URL        string  `json:"url"`
	Image      []Image `json:"image"`
	Streamable string  `json:"streamable"`
	MBID       string  `json:"mbid"`
}

// FullAlbum ...
type FullAlbum struct {
	Artist    string      `json:"artist"`
	MBID      string      `json:"mbid"`
	Tags      Tags        `json:"tags"`
	Playcount string      `json:"playcount"`
	Image     []Image     `json:"image"`
	Tracks    AlbumTracks `json:"tracks"`
	URL       string      `json:"url"`
	Name      string      `json:"name"`
	Listeners string      `json:"listeners"`
	Wiki      Wiki        `json:"wiki"`
}

// AlbumRanked ...
type AlbumRanked struct {
	Name      string      `json:"name"`
	Playcount string      `json:"playcount"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Artist    AlbumArtist `json:"artist"`
	Image     []Image     `json:"image"`
	Attr      Rank        `json:"@attr"`
}

// WeeklyAlbumChart ...
type WeeklyAlbumChart struct {
	Artist    AlbumRecent `json:"artist"`
	MBID      string      `json:"mbid"`
	URL       string      `json:"url"`
	Name      string      `json:"name"`
	Attr      Rank        `json:"@attr"`
	Playcount string      `json:"playcount"`
}

// AlbumRecent ...
type AlbumRecent struct {
	MBID string `json:"mbid"`
	Text string `json:"#text"`
}

// AlbumTracks ...
type AlbumTracks struct {
	Track []TrackAlbum `json:"track"`
}

// AlbumTopTags ...
type AlbumTopTags struct {
	Tag  []TagWithCount `json:"tag"`
	Attr AlbumAttr      `json:"@attr"`
}

// AlbumAttr ...
type AlbumAttr struct {
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

// AlbumSearchRes ...
type AlbumSearchRes struct {
	Query             OpenSearchQuery `json:"opensearch:Query"`
	QueryTotalResults string          `json:"opensearch:totalResults"`
	QueryStartIndex   string          `json:"opensearch:startIndex"`
	QueryItemsPerPage string          `json:"opensearch:itemsPerPage"`
	AlbumMatches      AlbumMatches    `json:"albummatches"`
	Attr              SearchAttr      `json:"@attr"`
}

// AlbumMatches ...
type AlbumMatches struct {
	Album []SimpleAlbum `json:"album"`
}

/*
Get the metadata and tracklist for an album on Last.fm using the album name or a musicbrainz id.

artist (Required (unless mbid)] : The artist name

album (Required (unless mbid)] : The album name

mbid (Optional) : The musicbrainz id for the album

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

username (Optional) : The username for the context of the request. If supplied, the user's playcount for this album is included in the response.

lang (Optional) : The language to return the biography in, expressed as an ISO 639 alpha-2 code.

api_key (Required) : A Last.fm API key.
*/
func (c *Client) AlbumGetInfo(album, artist string, opts ...RequestOption) (FullAlbum, error) {
	// http://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=YOUR_API_KEY&artist=Cher&album=Believe&format=json
	lastfmURL := fmt.Sprintf("%s&method=album.getinfo", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist), AlbumOpt(album))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var albumInfo struct {
		FullAlbum FullAlbum `json:"album"`
	}

	err := c.get(lastfmURL, &albumInfo)

	if err != nil {
		return FullAlbum{}, err
	}

	return albumInfo.FullAlbum, nil
}

/*
Get the top tags for an album on Last.fm, ordered by popularity.

artist (Required (unless mbid)] : The artist name

album (Required (unless mbid)] : The album name

autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.

mbid (Optional) : The musicbrainz id for the album

api_key (Required) : A Last.fm API key.
*/
func (c *Client) AlbumGetTopTags(album, artist string, opts ...RequestOption) (AlbumTopTags, error) {
	// http://ws.audioscrobbler.com/2.0/?method=album.gettoptags&artist=radiohead&album=the%20bends&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=album.gettoptags", c.baseApiURL)

	opts = append(opts, ArtistOpt(artist), AlbumOpt(album))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var topTags struct {
		TopTags AlbumTopTags `json:"toptags"`
	}

	err := c.get(lastfmURL, &topTags)

	if err != nil {
		return AlbumTopTags{}, err
	}

	return topTags.TopTags, nil
}

/*
Search for an album by name. Returns album matches sorted by relevance.

limit (Optional) : The number of results to fetch per page. Defaults to 30.

page (Optional) : The page number to fetch. Defaults to first page.

album (Required) : The album name

api_key (Required) : A Last.fm API key.
*/
func (c *Client) AlbumSearch(album string, opts ...RequestOption) (AlbumSearchRes, error) {
	// http://ws.audioscrobbler.com/2.0/?method=album.search&album=believe&api_key=YOUR_API_KEY&format=json
	lastfmURL := fmt.Sprintf("%s&method=album.search", c.baseApiURL)

	opts = append(opts, AlbumOpt(album))

	values := processOptions(opts...).urlParams

	if query := values.Encode(); query != "" {
		lastfmURL += "&" + query
	}

	var searchRes struct {
		SearchResults AlbumSearchRes `json:"results"`
	}

	err := c.get(lastfmURL, &searchRes)

	if err != nil {
		return AlbumSearchRes{}, err
	}
	return searchRes.SearchResults, nil
}
