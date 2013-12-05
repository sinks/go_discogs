
package go_discogs

import (
    "testing"
    "strings"
    "net/url"
)

var artist1 = []byte(
`
    {
        "releases_url": "http://api.discogs.com/artists/1/releases",
        "name": "Persuader, The",
        "namevariations": [
        "Persuader",
        "Presuader, The"
        ],
        "uri": "http://www.discogs.com/artist/Persuader%2C+The",
        "images": [
        {
            "uri": "http://api.discogs.com/image/A-1-1138987958.jpeg",
            "height": 450,
            "width": 600,
            "resource_url": "http://api.discogs.com/image/A-1-1138987958.jpeg",
            "type": "secondary",
            "uri150": "http://api.discogs.com/image/A-150-1-1138987958.jpeg"
        }
        ],
        "resource_url": "http://api.discogs.com/artists/1",
        "aliases": [
        {
            "resource_url": "http://api.discogs.com/artists/19541",
            "id": 19541,
            "name": "Dick Track"
        },
        {
            "resource_url": "http://api.discogs.com/artists/278760",
            "id": 278760,
            "name": "Faxid"
        },
        {
            "resource_url": "http://api.discogs.com/artists/16055",
            "id": 16055,
            "name": "Groove Machine"
        },
        {
            "resource_url": "http://api.discogs.com/artists/196957",
            "id": 196957,
            "name": "Janne Me' Amazonen"
        },
        {
            "resource_url": "http://api.discogs.com/artists/239",
            "id": 239,
            "name": "Jesper Dahlbäck"
        },
        {
            "resource_url": "http://api.discogs.com/artists/25227",
            "id": 25227,
            "name": "Lenk"
        },
        {
            "resource_url": "http://api.discogs.com/artists/439150",
            "id": 439150,
            "name": "Pinguin Man, The"
        }
        ],
        "id": 1,
        "data_quality": "Correct",
        "realname": "Jesper Dahlbäck"
    }
`)

type FakeRequester struct {

}

func (f *FakeRequester) Get(resource string, param url.Values) ([]byte, error) {
    if strings.HasSuffix(resource, "1") {
        return artist1, nil
    }

    return nil, nil
}

func TestGet(t *testing.T) {
    artist := ArtistDatabase{Requester: &FakeRequester{}}

    result, _ := artist.Get(1)

    if result.ReleasesURL != "http://api.discogs.com/artists/1/releases" {
        t.Fatal("Releases URL")
    }

    if result.Name != "Persuader, The" {
        t.Fatal("Name")
    }

    if result.NameVariations[0] != "Persuader" {
        t.Fatalf("Name Variation 0: %s", result.NameVariations[0])
    }

    if result.NameVariations[1] != "Presuader, The" {
        t.Fatalf("Name Variation 1: %s", result.NameVariations[1])
    }

    if result.URI != "http://www.discogs.com/artist/Persuader%2C+The" {
        t.Fatal("URI")
    }

    if result.Images[0].URI != "http://api.discogs.com/image/A-1-1138987958.jpeg" {
        t.Fatal("Image URI 0")
    }

    if result.Images[0].Width != 600 {
        t.Fatal("Image Width 0")
    }

    if result.Images[0].Height != 450 {
        t.Fatal("Image Height 0")
    }

    if result.Images[0].ResourceURL != "http://api.discogs.com/image/A-1-1138987958.jpeg" {
        t.Fatal("Image Heigh Resource URL 0")
    }

    if result.Images[0].Type != "secondary" {
        t.Fatal("Image Type 0")
    }

    if result.Images[0].URI150 != "http://api.discogs.com/image/A-150-1-1138987958.jpeg" {
        t.Fatal("Image URI150 0")
    }

    if result.ResourceURL != "http://api.discogs.com/artists/1" {
        t.Fatalf("Resource URL: %s", result.ResourceURL)
    }

    if len(result.Aliases) != 7 {
        t.Fatal("Aliases length")
    }

    if result.Aliases[0].ResourceURL != "http://api.discogs.com/artists/19541" {
        t.Fatalf("Aliases Resource URL 0: %s", result.Aliases[0].ResourceURL)
    }

    if result.Aliases[0].Id != 19541 {
        t.Fatalf("Aliases Id 0: %d", result.Aliases[0].Id)
    }

    if result.Aliases[0].Name != "Dick Track" {
        t.Fatalf("Aliases Name 0: %s", result.Aliases[0].Name)
    }

    if result.Aliases[6].ResourceURL != "http://api.discogs.com/artists/439150" {
        t.Fatalf("Aliases Resource URL 6: %s", result.Aliases[6].ResourceURL)
    }

    if result.Aliases[6].Id != 439150 {
        t.Fatalf("Aliases Id 6: %d", result.Aliases[6].Id)
    }

    if result.Aliases[6].Name != "Pinguin Man, The" {
        t.Fatalf("Aliases Name 6: %s", result.Aliases[6].Name)
    }

    if result.Id != 1 {
        t.Fatal("Artist id does not match")
    }

    if result.DataQuality != "Correct" {
        t.Fatalf("Data Quality: %s", result.DataQuality)
    }
}
