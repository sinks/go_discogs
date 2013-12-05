package go_discogs

import (
    "net/http"
    "net/url"
    "encoding/json"
)

var (
    releasePath = "/releases/"
)

type Rating struct {
    Count                   int             `json:"count"`
    Average                 float32         `json:"average"`
}

type Community struct {
    Status                  string          `json:"status"`
    Have                    string          `json:"have"`
    Want                    string          `json:"wamt"`
    Rating                  Rating          `json:"rating"`
    Contributors            []User          `json:"contributors"`
    Submitter               User            `json:"submitter"`
}

type User struct {
    Username                string          `json:"username"`
    ResourceUrl             string          `json:"resource_url"`
}

type Image struct {
    Uri                     string          `json:"uri"`
    Height                  int             `json:"height"`
    Width                   int             `json:"width"`
    ResourceUrl             string          `json:"respurce_url"`
    Type                    string          `json:"type"`
    Uri150                  string          `json:"uri150"`
}

type Track struct {
    Duration                string          `json:"duration"`
    Position                string          `json:"position"`
    Title                   string          `json:"title"`
    Artists                 []Artist        `json:"artists"`
    ExtraArtists            []Artist        `json:"extraartists"`
}

type Artist struct {
    Id                      int             `json:"id"`
    Join                    string          `json:"join"`
    Name                    string          `json:"name"`
    Anv                     string          `json:"anv"`
    Tracks                  string          `json:"tracks"`
    Role                    string          `json:"role"`
    ResourceUrl             string          `json:"resource_url"`
}

type Video struct {
    Duration                int             `json:"duration"`
    Description             string          `json:"description"`
    Embed                   bool            `json:"embed"`
    Uri                     string          `json:"uri"`
    Title                   string          `json:"title"`
}

type Format struct {
    Descriptions            []string        `json:"descriptions"`
    Name                    string          `json:"name"`
    Quantity                string          `json:"qty"`
}

type Series struct {
    Id                      int             `json:"id"`
    ResourceUrl             string          `json:"resource_url"`
    EntityType              string          `json:"entity_type"`
    CatalogNumber           string          `json:"catno"`
    Name                    string          `json:"name"`
}

type Identifier struct {
    Type                    string          `json:"type"`
    Value                   string          `json:"value"`
}

type Label struct {
    Id                      int             `json:"id"`
    ResourceUrl             string          `json:"resource_url"`
    EntityType              string          `json:"entity_type"`
    EntityTypeName          string          `json:"entity_type_name"`
    CatalogNumber           string          `json:"catno"`
    Name                    string          `json:"name"`
}

type Company struct {
    Id                      int             `json:"id"`
    EntityType              string          `json:"entity_type"`
    EntityTypeName          string          `json:"entity_type_name"`
    CatalogNumber           string          `json:"catno"`
    Name                    string          `json:"name"`
    ResourceUrl             string          `json:"resource_url"`
}

type Release struct {
    Id                      int             `json:"id"`
    Title                   string          `json:"title"`
    ResourceUrl             string          `json:"resource_url"`
    Uri                     string          `json:"uri"`
    Status                  string          `json:"status"`
    DataQuality             string          `json:"data_quality"`
    MasterId                int             `json:"master_id"`
    MasterUrl               string          `json:"master_url`
    Country                 string          `json:"country"`
    Year                    int             `json:"year"`
    Released                string          `json:"released"`
    ReleasedFormatted       string          `json:"release_formatted"`
    Notes                   string          `json:"notes"`
    Styles                  []string        `json:"styles"`
    Genres                  []string        `json:"genres"`
    EstimatedWeight         int             `json:"estimated_weight:`
    FormatQuantity          int             `json:"format_quantity"`
    Community               Community       `json:"community"`
    Labels                  []Label         `json:"labels"`
    Companies               []Company       `json:"companies"`
    ExtraArtists            []Artist        `json:"extraartists"`
    Videos                  []Video         `json:"videos"`
    Artists                 []Artist        `json:"artists"`
    Formats                 []Format        `json:"formats"`
    Images                  []Image         `json:"images"`
    Identifiers             []Identifier    `json:"identifiers"`
    Tracklist               []Track         `json:"tracklist"`
    Series                  []Series        `json:"series"`
    Thumb                   string          `json:"thumb"`
}

func Fetch(id int) Release {
    // set the host and path
    searchUrl, err := url.Parse(apiUrl + releasePath + id)
    if err != nil {
    }

    // issue the get request
    resp, err := http.Get(searchUrl.String())
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var release Release{}
    err = json.Unmarshal(body, &release)

    return release
}

