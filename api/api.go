package api

import (
	"../conf"
	"../mysql"
	_ "github.com/go-sql-driver/mysql"
)

type PersonService struct {
	Config     *conf.Config
	Repository *mysql.PersonRepository
}

func NewPersonService(config *conf.Config, repository *mysql.PersonRepository) *PersonService {
	return &PersonService{Config: config, Repository: repository}
}

func (service *PersonService) FindAll() []*mysql.Person {
	if service.Config.Enabled {
		return service.Repository.FindAll()
	}

	return []*mysql.Person{}
}
