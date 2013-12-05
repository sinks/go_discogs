package go_discogs

import (
    "net/url"
    "net/http"
    "io/ioutil"
)

type Requester interface {
    Get(string, url.Values) ([]byte, error)
}

type HTTPRequester struct  {
    UserAgent               string
    Client                  *http.Client
    ApiURL                  *url.URL
}

func NewHTTPRequester(userAgent string) *HTTPRequester {
    return &HTTPRequester{UserAgent: userAgent,
                          Client: http.DefaultClient,
                          ApiURL: &url.URL{Scheme: "http",
                                           Host: "api.discogs.com",
                                          },
                         }
}

// Resource should 
func (r *HTTPRequester) Get(resource string, param url.Values) ([]byte, error) {

    // Create url
    rURL := r.ApiURL.ResolveReference(&url.URL{Path: resource,
                                             RawQuery: param.Encode()})

    // Generate request
    request, err := http.NewRequest("GET", rURL.String(), nil)
    if err != nil {
        return nil, err
    }
    request.Header.Add("UserAgent", r.UserAgent)

    // Perform request
    response, err := r.Client.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    // Read body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}

