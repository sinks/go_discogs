package go_discogs


type ArtistResult struct {
    Id                          int                 `json:"id"`
    ResourceURL                 string              `json:"resource_url"`
    Name                        string              `json:"name"`
    RealName                    string              `json:"realname"`
    ReleasesURL                 string              `json:"releases_url"`
    Images                      []Image             `json:"images"`
}


type ArtistRelease struct {
    Status                      string              `json:"status"`
    Thumb                       string              `json:"thumb"`
    Title                       string              `json:"title"`
    Format                      string              `json:"format"`
    Label                       string              `json:"label"`
    Role                        string              `json:"role"`
    Year                        int                 `json:"year"`
    TrackInfo                   string              `json:"trackinfo"`
    ResourceURL                 string              `json:"resourceURL"`
    Artist                      string              `json:"artist"`
    Type                        string              `json:"type"`
    Id                          int                 `json:"id"`
}


type ArtistMaster struct {
    Thumb                       string              `json:"thumb"`
    MainRelease                 int                 `json:"main_release"`
    Title                       string              `json:"title"`
    Role                        string              `json:"role"`
    Year                        int                 `json:"year"`
    ResourceURL                 string              `json:"resourceURL"`
    Artist                      string              `json:"artist"`
    Type                        string              `json:"type"`
    Id                          int                 `json:"id"`

}

type ArtistReleasesResult struct {
    Pagination                  Pagination          `json:"pagination"`
    Releases                    []ArtistRelease     `json:"releases"`
    Master                      []ArtistMaster      `json:"releases"`
}

func (arr *ArtistReleasesResult) UnmarshalJSON(src []byte) error {
    var m map[string]interface{}

    json.Unmarshal(src, &m)

    // get Master types
    for m["releases"] 
    // get Release Types

}
