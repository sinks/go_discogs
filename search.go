
package go_discogs

const searchPath = "/database/search"

type SearchDatabase struct {
    Requester           Requester
}

/*
func (sd *SearchDatabase) Search(s SearchParameters, p PaginationParameters) (SearchResult, error) {
    // turn struct into url.Values (map[string]string)

    jsonResult, err := Requester.Get(searchPath, parameters)
    if err != nil {
        return nil, err
    }

    var result SearchResult
    json.Unmarshal(jsonResult, &result)
    
    
    return result, nil
}
*/
