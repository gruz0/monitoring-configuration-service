package domain

type Domain struct {
	ID       int
	Name     string
	Verified bool
}

type Repository struct{}

func New(id int, name string, verified bool) *Domain {
	return &Domain{
		ID:       id,
		Name:     name,
		Verified: verified,
	}
}

func (r *Repository) FindAllBySiteID(_ int) []*Domain {
	return []*Domain{}
}

func NewDomainRepository() *Repository {
	return &Repository{}
}
