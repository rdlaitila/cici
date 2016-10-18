package hateoas

// Link represents a link object in a hateoas response
type Link struct {
	// Rel represents a relation identifier
	Rel string

	// Href represents a relation location
	Href string
}
