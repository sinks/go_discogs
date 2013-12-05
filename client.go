package go_discogs

// client := go_discogs.NewClientDatabase("user_agent")
//
// Searching
// client.Search.Search(params, pagination)
//
// Get an Artist
// client.Artist.Get(1234)
// client.Artist.Get(1234).Releases(pagination) or client.Artist.Releases(1234, pagination)
// 
// Fetch a release
// client.Release.Get(1234)
//
// Fetch a master
// client.Master.Get(1234)
// client.Master.Get(1234).Versions(pagination) or client.Master.Versions(1234, pagination)
//
// Fetch a label
// client.Label.Get(1234)
// client.Label.Get(1234).Releases(pagination) or client.Label.Releases(1234, pagination)
type Client struct {
    Search                  *SearchDatabase
    Artist                  *ArtistDatabase
    Release                 *ReleaseDatabase
    Master                  *MasterDatabase
    Label                   *LabelDatabase
}

/*
type ClientMarketplace struct {
    Inventory
    Listing
    Order
    Fee
    PriceSuggestions
}

type ClientUser struct {
    Identity
    Profile
    Collections
    Wantlist
}
*/

func NewClient(userAgent string) *Client {
    requester := NewRequester(userAgent)
    return &Client{Search: &SearchDatabase{Requester: requester},
                   Artist: &ArtistDatabase{Requester: requester},
                   Release: &ReleaseDatabase{Requester: requester},
                   Master: &MasterDatabase{Requester: requester},
                   Label: &LabelDatabase{Requester: requester},
                  }
}
