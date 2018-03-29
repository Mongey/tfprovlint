package provparse

// Provider represents the data for the provider.
type Provider struct {
	Name        string
	Resources   []Resource
	DataSources []Resource
}

func findResource(resources []Resource, name string) *Resource {
	for _, r := range resources {
		if r.Name == name {
			return &r
		}
	}

	return nil
}

// Resource looks up a resource within the provider by name and returns nil if not found.
func (p *Provider) Resource(name string) *Resource {
	return findResource(p.Resources, name)
}

// DataSource looks up a data source in the provider by name and returns nil if not found.
func (p *Provider) DataSource(name string) *Resource {
	return findResource(p.DataSources, name)
}

// Resource represents the data for a resource or data source of the provider.
type Resource struct {
	Provider         string // azurerm
	Name             string // azurerm_image
	NameSuffix       string // image
	Type             string // data vs resource?
	ShortDescription string // Get information about an Image
	Description      string // Use this data source to access information about an Image.
	// TODO: +Example usage, etc?
	// TODO: resource category
	Attributes []Attribute
}

func findAttribute(atts []Attribute, name string) *Attribute {
	for _, att := range atts {
		if att.Name == name {
			return &att
		}
	}

	return nil
}

// Attribute returns an attribute of the resource by name or nil if not found.
func (r *Resource) Attribute(name string) *Attribute {
	return findAttribute(r.Attributes, name)
}

// Attribute returns a child attribute by name or nil if not found.
func (a *Attribute) Attribute(name string) *Attribute {
	return findAttribute(a.Attributes, name)
}

// Attribute represents a data element of the resource schema.
type Attribute struct {
	Name        string
	Description string
	Optional    bool
	Required    bool
	Computed    bool

	Attributes []Attribute
	Min        int
	Max        int
}