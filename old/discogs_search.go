package go_discogs

import (
    "net/http"
    "net/url"
    "reflect"
    "fmt"
    "io/ioutil"
    "encoding/json"
)


var (
    searchPath = "/database/search" // api_url /database/search?q=<term1>%20<term2>
)



type Result struct {
    Pagination      Pagination  `json:"pagination"`
}

type Pagination struct {
    PerPage         int         `json:"per_page"`
    Items           int         `json:"pages"`
    Page            int         `json:"page"`
    Urls            PaginationUrls  `json:"urls"`
    Items           int         `param:"items"`
}

type PaginationUrls struct {
    Last            string      `json:"last"`
    Next            string      `json:"next"`
}

type LabelResult struct {
    Id              int         `json:"id"`
    Uri             string      `json:"uri"`
    ResourceUrl     string      `json:"resource_url"`
    Title           string      `json:"title"`
    Type            string      `json:"type"`
    Thumb           string      `json:"thumb"`
}

type MasterResult struct {
    Id              int         `json:"id"`
    Style           []string    `json:"style"`
    Thumb           string      `json:"thumb"`
    Format          []string    `json:"format"`
    Country         string      `json:"country"`
    Title           string      `json:"title"`
    Uri             string      `json:"uri"`
    Label           []string    `json:"label"`
    CatalogNumber   string      `json:"catno"`
    Year            string      `json:"year"`
    Genre           []string    `json:"genre"`
    ResourceUrl     string      `json:"resource_url"`
    Type            string      `json:"type"`
}

type ArtistResult struct {
    Id              int         `json:"id"`
    Uri             string      `json:"uri"`
    ResourceUrl     string      `json:"resource_url"`
    Title           string      `json:"title"`
    Type            string      `json:"type"`
    Thumb           string      `json:"thumb"`
}

type ReleaseResult struct {
    Id              int         `json:"id"`
    Style           []string    `json:"style"`
    Thumb           string      `json:"thumb"`
    Format          []string    `json:"format"`
    Country         string      `json:"country"`
    Title           string      `json:"title"`
    Uri             string      `json:"uri"`
    Label           []string    `json:"label"`
    CatalogNumber   string      `json:"catno"`
    Year            string      `json:"year"`
    Genre           []string    `json:"genre"`
    ResourceUrl     string      `json:"resource_url"`
    Type            string      `json:"type"`
}

type DiscogsSearch struct {
    Parameters      SearchParameters
}

func (d *DiscogsSearch) Search()  {
    // set the host and path
    searchUrl, err := url.Parse(apiUrl + searchPath)
    if err != nil {
    }

    // get the query parameters
    params := url.Values{}

    t := reflect.TypeOf(&d.Parameters).Elem()
    v := reflect.ValueOf(&d.Parameters).Elem()
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        tag := field.Tag.Get("param")
        value := url.QueryEscape(v.Field(i).String())
        if value != "" {
            params.Add(tag, value)
        }
    }
    searchUrl.RawQuery = params.Encode()

    fmt.Println(searchUrl.String())
    // issue the get request
    resp, err := http.Get(searchUrl.String())
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var f interface{}
    err = json.Unmarshal(body, &f)
}

