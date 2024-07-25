// Package rest provides a simple and convenient way to make RESTful HTTP requests.
//
// It defines a Requester interface with methods corresponding to each HTTP method.
// The Rest struct is an implementation of the Requester, which also provides a base
// URL and a set of base options that can be augmented with further options for each request.
//
// Example usage:
//
//	baseURL, _ := url.Parse("https://api.example.com")
//	client := rest.NewRest(baseURL, do.WithHeader("Authorization", "Bearer token"))
//	response, err := client.GET(ctx, do.WithPath("/users"))
//	if err != nil {
//	    // handle error
//	}
//	// use response
package rest
