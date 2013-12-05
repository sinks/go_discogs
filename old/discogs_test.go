
package go_discogs

import (
    "testing"
)

func TestSearch(t *testing.T) {
    search := DiscogsSearch{}
    search.Parameters.Artist = "Animal Collective"
    search.Parameters.ReleaseTitle = "Feels"
    search.Parameters.Type = "release"

    search.Search()
}
