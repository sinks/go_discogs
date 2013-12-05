package go_discogs

import (
    "net/url"
    "reflect"
)

// Search parameters
// http://www.discogs.com/developers/resources/database/search-endpoint.html
type SearchParameters struct {
    Q               string      `param:"q"`
    Type            string      `param:"type"`
    Title           string      `param:"title"`
    ReleaseTitle    string      `param:"release_title"`
    Credit          string      `param:"credit"`
    Artist          string      `param:"artist"`
    Anv             string      `param:"anv"`
    Label           string      `param:"label"`
    Genre           string      `param:"genre"`
    Style           string      `param:"style"`
    Country         string      `param:"country"`
    Year            string      `param:"year"`
    Format          string      `param:"format"`
    CatalogNumber   string      `param:"catno"`
    Barcode         string      `param:"barcode"`
    Track           string      `param:"track"`
    Submitter       string      `param:"submitter"`
    Contributer     string      `param:"contributer"`
}

type PaginationParameters struct {
    Page            string      `param:"page"`
    PerPage         string      `param:"per_page"`
}

func StructToValues(param interface{}) url.Values {
    // get the query parameters
    params := url.Values{}

    // loop over fields
    t := reflect.TypeOf(&param).Elem()
    v := reflect.ValueOf(&param).Elem()
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        tag := field.Tag.Get("param")
        value := url.QueryEscape(v.Field(i).String())
        if value != "" {
            params.Add(tag, value)
        }
    }
    return params
}
