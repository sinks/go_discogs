package go_discogs

import (
    "fmt"
    "encoding/json"
)

type ArtistDatabase struct {
    Requester           Requester
}

const artistPath = "/artists/%d"
const artistReleasePath = "/artists/%d/releases"

func (ad *ArtistDatabase) Get(id int) (*ArtistResult, error) {
    // generate url
    data, err := ad.Requester.Get(fmt.Sprintf(artistPath, id), nil)

    if err != nil {
        return nil, err
    }

    result := ArtistResult{}

    // Convert json body to type
    err = json.Unmarshal(data, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

/*
func (ad *ArtistDatabase) Releases(id int, pagination PaginationParameters) {
    // generate url
    data, err := ad.Requester.Get(fmt.Sprintf(artistReleasePath, id),
                                  StructToValues(pagination))

    if err != nil {
        return nil, err
    }

    result := ArtistiReleassResult{}

    // Convert json body to type
    err = json.Unmarshal(data, &result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}
*/
