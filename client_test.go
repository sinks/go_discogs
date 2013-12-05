
package go_discogs

import (
    "testing"
)

// TODO: mock the response
func TestClientArtistDatabase(t *testing.T) {
    client := NewClient("MrOrganiser/0.1 +http://github.com/sinks/")
    result, err := client.Artist.Get(1)

    if err != nil {
        t.Fatalf("Artist get error %s", err)
    }

    if result.Name != "Persuader, The" {
        t.Fatal("Artist name not correct")
    }

    if result.Id != 1 {
        t.Fatal("Artist id does not match")
    }

    if result.ResourceURL != "http://api.discogs.com/artists/1" {
        t.Fatal("Artist resource url does not match")
    }

    // TODO: Add all fields for images
    // TODO: Test multi images
    if result.Images[0].Width != 600 {
        t.Fatal("Artist image width does not match")
    }


}
