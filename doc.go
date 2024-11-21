/*
Package backstage provides a client for the Backstage API.

# Usage

Add the package to your project as following:

	import "github.com/datolabs-io/go-backstage/v3"

Once imported, create a new Backstage API client to access different parts of Backstage API:

	client, err := backstage.NewClient(baseURL, "default", nil)

If you want to use a custom HTTP client (for example, to handle authentication, retries or different timeouts), you can pass it as the
third argument:

	httpClient := &http.Client{}
	client, err := backstage.NewClient(baseURL, "default", httpClient)

The client than can be used to access different parts of the API, e.g. get the list of entities, sorted in specific order:

	entities, response, err := c.Catalog.Entities.s.List(context.Background(), &ListEntityOptions{
	        Filters: []string{},
	        Fields:  []string{},
	        Order:   []ListEntityOrder{{ Direction: OrderDescending, Field: "metadata.name" },
	    },
	})
*/
package backstage
