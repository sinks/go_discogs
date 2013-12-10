package go_discogs

import (
    "encoding/json"
    "github.com/mitchellh/mapstructure"
)

type ArtistResult struct {
    Resource                                        `json:",squash"` // squash from mapstructure
    URI                         string              `json:"uri"`
    ReleasesURL                 string              `json:"releases_url"`
    Name                        string              `json:"name"`
    RealName                    string              `json:"realname"`
    NameVariations              []string            `json:"namevariations"`
    Aliases                     []Alias             `json:"aliases"`
    Images                      []Image             `json:"images"`
    Members                     []Member            `json:"members"`
    Profile                     string              `json:"profile"`
    URLs                        []string            `json:"urls"`
    DataQuality                 string              `json:"data_quality"`
}

type ArtistReleasesResult struct {
    Pagination                  Pagination          `json:"pagination"`
    Releases                    []ArtistRelease     `json:"releases"`
    Masters                     []ArtistMaster      `json:"releases"`
}

type ArtistRelease struct {
    Resource                                        `json:",squash"` // squash from mapstructure
    Status                      string              `json:"status"`
    Thumb                       string              `json:"thumb"`
    Title                       string              `json:"title"`
    Format                      string              `json:"format"`
    Label                       string              `json:"label"`
    Role                        string              `json:"role"`
    Year                        int                 `json:"year"`
    TrackInfo                   string              `json:"trackinfo"`
    Artist                      string              `json:"artist"`
    Type                        string              `json:"type"`
}


type ArtistMaster struct {
    Resource                                        `json:",squash"` // squash from mapstructure
    Thumb                       string              `json:"thumb"`
    MainRelease                 int                 `json:"main_release"`
    Title                       string              `json:"title"`
    Role                        string              `json:"role"`
    Year                        int                 `json:"year"`
    Artist                      string              `json:"artist"`
    Type                        string              `json:"type"`
}

type Alias struct {
    Resource                                        `json:",squash"` // squash from mapstructure
    Name                        string              `json:"name"`
}

type Member struct {
    Resource                                        `json:",squash"` // squash from mapstructure
    Name                        string              `json:"name"`
    Active                      bool                `json:"active"`
}


func (arr *ArtistReleasesResult) UnmarshalJSON(src []byte) error {
    var m map[string]interface{}

    json.Unmarshal(src, &m)
    // get Master types
    releases := m["releases"]
    t := releases.([]interface{})

    // Populate Masters and Releases
    for _, v := range t {
        switch vv := v.(type) {
        case map[string]interface{}:
            if vv["type"] == "master" {
                var master ArtistMaster
                var md mapstructure.Metadata

                config := &mapstructure.DecoderConfig{TagName: "json",
                                                      Result: &master,
                                                      Metadata: &md}
                decoder, err := mapstructure.NewDecoder(config)
                if err != nil {
                    return err
                }
                err = decoder.Decode(vv)
                arr.Masters = append(arr.Masters, master)
            }
            if vv["type"] == "release" {
                var release ArtistRelease
                config := &mapstructure.DecoderConfig{TagName: "json",
                                                      Result: &release}
                decoder, err := mapstructure.NewDecoder(config)
                if err != nil {
                    return err
                }
                err = decoder.Decode(vv)
                arr.Releases = append(arr.Releases, release)
            }
        }
    }

    // Populate pagination
    pagination := m["pagination"]
    var page Pagination
    config := &mapstructure.DecoderConfig{TagName: "json", Result: &page}
    decoder, _ := mapstructure.NewDecoder(config)
    decoder.Decode(pagination)
    arr.Pagination = page

    // get Release Types
    return nil
}

