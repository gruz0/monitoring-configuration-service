package customer

import (
	"github.com/gruz0/monitoring-configuration-service/internal/domain"
	"github.com/gruz0/monitoring-configuration-service/internal/plugin"
	"github.com/gruz0/monitoring-configuration-service/internal/site"
)

type Customer struct {
	ID    int
	Email string
	Sites []*site.Site
}

type Repository struct{}

func New(id int, email string, sites []*site.Site) *Customer {
	return &Customer{
		ID:    id,
		Email: email,
		Sites: sites,
	}
}

func (r *Repository) FindAllActive() []*Customer {
	plugin1 := plugin.New(1, "http.http_status200")
	plugin2 := plugin.New(2, "http.http_redirect")
	plugin3 := plugin.New(3, "content.content_exists")

	domain1 := domain.New(1, "domain1.tld", true)
	domain2 := domain.New(2, "domain2.tld", false)
	domain3 := domain.New(3, "domain3.tld", true)
	domain4 := domain.New(4, "domain4.tld", false)

	site1 := site.New(1, domain1, []*plugin.Plugin{plugin1, plugin2})
	site2 := site.New(2, domain2, []*plugin.Plugin{plugin2, plugin3})
	site3 := site.New(3, domain3, []*plugin.Plugin{plugin1, plugin3})
	site4 := site.New(4, domain4, []*plugin.Plugin{plugin1, plugin2, plugin3})

	return []*Customer{
		New(1, "me@domain1.tld", []*site.Site{site1}),
		New(2, "me@domain2.tld", []*site.Site{site2}),
		New(3, "me@domain3.tld", []*site.Site{site3, site4}),
	}
}

func NewCustomerRepository() *Repository {
	return &Repository{}
}
