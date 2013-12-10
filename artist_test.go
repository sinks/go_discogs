
package go_discogs

import (
    "testing"
    "strings"
    "net/url"
    "encoding/json"
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

var artist207958 = []byte(`
{
    "profile": "New York based avant-garde collective from Baltimore, Maryland.",
    "releases_url": "http://api.discogs.com/artists/207958/releases",
    "name": "Animal Collective",
    "uri": "http://www.discogs.com/artist/Animal+Collective",
    "members": [
    {
        "active": true,
        "resource_url": "http://api.discogs.com/artists/2083551",
        "id": 2083551,
        "name": "Brian Weitz"
    },
    {
        "active": true,
        "resource_url": "http://api.discogs.com/artists/323620",
        "id": 323620,
        "name": "Dave Portner"
    },
    {
        "active": true,
        "resource_url": "http://api.discogs.com/artists/563743",
        "id": 563743,
        "name": "Josh Dibb"
    },
    {
        "active": true,
        "resource_url": "http://api.discogs.com/artists/985780",
        "id": 985780,
        "name": "Noah Lennox"
    }
    ],
    "urls": [
    "http://www.myspace.com/animalcollective",
    "http://www.myanimalhome.net/",
    "http://en.wikipedia.org/wiki/Animal_Collective"
    ],
    "images": [
    {
        "uri": "http://api.discogs.com/image/A-207958-1126891746.jpeg",
        "height": 340,
        "width": 600,
        "resource_url": "http://api.discogs.com/image/A-207958-1126891746.jpeg",
        "type": "secondary",
        "uri150": "http://api.discogs.com/image/A-150-207958-1126891746.jpeg"
    },
    {
        "uri": "http://api.discogs.com/image/A-207958-1139700349.jpeg",
        "height": 535,
        "width": 350,
        "resource_url": "http://api.discogs.com/image/A-207958-1139700349.jpeg",
        "type": "secondary",
        "uri150": "http://api.discogs.com/image/A-150-207958-1139700349.jpeg"
    },
    {
        "uri": "http://api.discogs.com/image/A-207958-1147965696.jpeg",
        "height": 400,
        "width": 600,
        "resource_url": "http://api.discogs.com/image/A-207958-1147965696.jpeg",
        "type": "secondary",
        "uri150": "http://api.discogs.com/image/A-150-207958-1147965696.jpeg"
    },
    {
        "uri": "http://api.discogs.com/image/A-207958-1187027106.jpeg",
        "height": 301,
        "width": 360,
        "resource_url": "http://api.discogs.com/image/A-207958-1187027106.jpeg",
        "type": "primary",
        "uri150": "http://api.discogs.com/image/A-150-207958-1187027106.jpeg"
    },
    {
        "uri": "http://api.discogs.com/image/A-207958-1187027124.jpeg",
        "height": 396,
        "width": 600,
        "resource_url": "http://api.discogs.com/image/A-207958-1187027124.jpeg",
        "type": "secondary",
        "uri150": "http://api.discogs.com/image/A-150-207958-1187027124.jpeg"
    }
    ],
    "resource_url": "http://api.discogs.com/artists/207958",
    "id": 207958,
    "data_quality": "Needs Vote",
    "namevariations": [
    "AC",
    "Avey Tare & Panda Bear",
    "Avey Tare, Panda Bear & Geologist"
    ]
}
`)

type FakeRequester struct {

}

func (f *FakeRequester) Get(resource string, param url.Values) ([]byte, error) {
    if strings.HasSuffix(resource, "1") {
        return artist1, nil
    }
    if strings.HasSuffix(resource, "207958") {
        return artist207958, nil
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


    result, _ = artist.Get(207958)

    if result.Profile != "New York based avant-garde collective from Baltimore, Maryland." {
        t.Fatalf("Profile: %s", result.Profile)
    }

    if len(result.Members) != 4 {
        t.Fatalf("Members length: %d", len(result.Members))
    }

    if result.Members[0].Active != true {
        t.Fatal("Members Active 0")
    }

    if result.Members[0].ResourceURL != "http://api.discogs.com/artists/2083551" {
        t.Fatalf("Member ResourceURL 0: %s", result.Members[0].ResourceURL)
    }

    if result.Members[0].Id != 2083551 {
        t.Fatalf("Member Id 0: %d", result.Members[0].Id)
    }

    if result.Members[0].Name != "Brian Weitz" {
        t.Fatalf("Member Name 0: %s", result.Members[0].Name)
    }

    if result.Members[3].Active != true {
        t.Fatal("Members Active 3")
    }

    if result.Members[3].ResourceURL != "http://api.discogs.com/artists/985780" {
        t.Fatalf("Member ResourceURL 3: %s", result.Members[3].ResourceURL)
    }

    if result.Members[3].Id != 985780 {
        t.Fatalf("Member Id 3: %d", result.Members[3].Id)
    }

    if result.Members[3].Name != "Noah Lennox" {
        t.Fatalf("Member Name 3: %s", result.Members[3].Name)
    }

    if len(result.URLs) != 3 {
        t.Fatalf("URLs length: %d", len(result.URLs))
    }

    if result.URLs[0] != "http://www.myspace.com/animalcollective" {
        t.Fatalf("URLs 0: %s", result.URLs[0])
    }

    if result.URLs[2] != "http://en.wikipedia.org/wiki/Animal_Collective" {
        t.Fatalf("URLs 2: %s", result.URLs[2])
    }

    if len(result.Images) != 5 {
        t.Fatalf("Images length: %d", len(result.Images))
    }

    if result.Images[4].ResourceURL != "http://api.discogs.com/image/A-207958-1187027124.jpeg" {
        t.Fatalf("Images Resource URL 4: %s", result.Images[4].ResourceURL)
    }

    if result.Id != 207958 {
        t.Fatalf("Id: %d", result.Id)
    }
}

var ArtistReleases67156 = []byte(`
{
"pagination": {
"per_page": 5,
"items": 410,
"page": 1,
"urls": {
"last": "http://api.discogs.com/artists/67156/releases?per_page=5&page=82",
"next": "http://api.discogs.com/artists/67156/releases?per_page=5&page=2"
},
"pages": 82
},
"releases": [
{
"thumb": "http://api.discogs.com/image/R-150-1568350-1365965467-8602.jpeg",
"artist": "Flaming Lips, The",
"main_release": 1568350,
"title": "The Flaming Lips",
"role": "Main",
"year": 1984,
"resource_url": "http://api.discogs.com/masters/114232",
"type": "master",
"id": 114232
},
{
"thumb": "http://api.discogs.com/image/R-150-724026-1280542367.jpeg",
"artist": "Flaming Lips, The",
"main_release": 724026,
"title": "Hear It Is",
"role": "Main",
"year": 1986,
"resource_url": "http://api.discogs.com/masters/114228",
"type": "master",
"id": 114228
},
{
"thumb": "http://api.discogs.com/image/R-150-1417309-1261044250.jpeg",
"artist": "Flaming Lips, The",
"main_release": 1417309,
"title": "Oh My Gawd!!!...",
"role": "Main",
"year": 1987,
"resource_url": "http://api.discogs.com/masters/114248",
"type": "master",
"id": 114248
},
{
"status": "Accepted",
"thumb": "http://api.discogs.com/image/R-150-2085863-1280566884.jpeg",
"title": "Untitled",
"format": "Flexi, 7\", S/Sided, Promo",
"label": "The Bob Magazine, Eva-Tone Soundsheets",
"role": "Main",
"year": 1988,
"resource_url": "http://api.discogs.com/releases/2085863",
"artist": "Fleshtones* / Steve Kilbey / Flaming Lips*",
"type": "release",
"id": 2085863
},
{
"thumb": "http://api.discogs.com/image/R-150-988293-1181182245.jpeg",
"artist": "Flaming Lips*",
"main_release": 988293,
"title": "Drug Machine",
"role": "Main",
"year": 1989,
"resource_url": "http://api.discogs.com/masters/146472",
"type": "master",
"id": 146472
}
]
}
`)

func TestArtistReleases(t *testing.T) {
    test := []byte(`{"releases": [{"type": "release", "name": "feels", "id": 45}, {"id": 37, "type": "master", "name": "masterfeels"}]}`)
    json.Unmarshal(test, &ArtistReleasesResult{})

    var artistRelease ArtistReleasesResult
    json.Unmarshal(ArtistReleases67156, &artistRelease)

    if artistRelease.Pagination.PerPage != 5 {
        t.Fatal("Pagination Per Page")
    }
    if artistRelease.Pagination.Items != 410 {
        t.Fatal("Pagination Items")
    }
    if artistRelease.Pagination.Page != 1 {
        t.Fatal("Pagination Page")
    }
    if artistRelease.Releases[0].Id != 2085863 {
        t.Fatalf("Releases Id 0: %d", artistRelease.Releases[0].Id)
    }
    if artistRelease.Masters[0].Id != 114232 {
        t.Fatalf("Masters Id 0: %d", artistRelease.Masters[0].Id)
    }
}
