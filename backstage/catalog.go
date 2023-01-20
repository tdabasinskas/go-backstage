package backstage

// catalogService handles communication with Backstage catalog API.
type catalogService struct {
	service

	// Entities handles communication with the Backstage entities endpoints in Backstage Catalog API.
	Entities *entityService

	// APIs handles communication with the API related methods of the Backstage Catalog API.
	APIs *apiService

	// Components handles communication with the Component related methods of the Backstage Catalog API.
	Components *componentService

	// Domains handles communication with the Domain related methods of the Backstage Catalog API.
	Domains *domainService

	// Groups handles communication with the Group related methods of the Backstage Catalog API.
	Groups *groupService

	// Locations handles communication with the Location related methods of the Backstage Catalog API.
	Locations *locationService

	// Resources handles communication with the Resource related methods of the Backstage Catalog API.
	Resources *resourceService

	// Systems handles communication with the System related methods of the Backstage Catalog API.
	Systems *systemService

	// Users handles communication with the User related methods of the Backstage Catalog API.
	Users *userService
}

const (
	// catalogApiPath is the path to the catalog API.
	catalogApiPath = "/catalog"
)

// newCatalogService returns a new instance of catalogService.
func newCatalogService(c *Client) *catalogService {
	p, _ := c.BaseURL.Parse(catalogApiPath)

	s := &catalogService{
		service: service{
			client:  c,
			apiPath: p,
		},
	}
	s.Entities = newEntityService(s)
	s.APIs = newApiService(s.Entities)
	s.Components = newComponentService(s.Entities)
	s.Domains = newDomainService(s.Entities)
	s.Groups = newGroupService(s.Entities)
	s.Locations = newLocationService(s.Entities)
	s.Resources = newResourceService(s.Entities)
	s.Systems = newSystemService(s.Entities)
	s.Users = newUserService(s.Entities)

	return s
}
