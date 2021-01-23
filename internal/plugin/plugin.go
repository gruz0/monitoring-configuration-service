package plugin

type Plugin struct {
	ID   int
	Name string
}

type Repository struct{}

func New(id int, name string) *Plugin {
	return &Plugin{
		ID:   id,
		Name: name,
	}
}

func (r *Repository) FindAll() []*Plugin {
	return []*Plugin{}
}

func NewPluginRepository() *Repository {
	return &Repository{}
}
