package go_discogs

// TODO: add a custom unmarshaller
//       handle adding unknown fields to an Extra field

type Resource struct {
    Id                          uint64              `json:"id"`
    ResourceURL                 string              `json:"resource_url"`
}

type PaginationURLs struct {
    Last                        string              `json:"last"`
    Next                        string              `json:"next"`
}

type Pagination struct {
    PerPage                     int                 `json:"per_page"`
    Items                       int                 `json:"items"`
    Page                        int                 `json:"page"`
    URLs                        PaginationURLs      `json:"urls"`
    Pages                       int                 `json:"pages"`
}

type Image struct {
    URI                         string              `json:"uri"`
    ResourceURL                 string              `json:"resource_url"`
    Height                      int                 `json:"height"`
    Width                       int                 `json:"width"`
    Type                        string              `json:"type"`
    URI150                      string              `json:"uri150"`
}
