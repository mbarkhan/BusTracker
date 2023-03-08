package services

import "githubmbarkhanBusTracker/auth/models"

type CompanyServiceInt interface {
	Create(*models.Company) (models.Company, error)
	Read(int) (models.Company, error)
	Update(*models.Company) (models.Company, error)
	Delete(int) error
	List() ([]models.Company, error)
}

type companyService struct {
	cmpService CompanyServiceInt
}

func (r *companyService) Create(cmp *models.Company) (models.Company, error) {
	return r.cmpService.Create(cmp)
}

func (r *companyService) Delete(cmp int) error {
	return r.cmpService.Delete(cmp)
}

func (r *companyService) List() ([]models.Company, error) {
	return r.cmpService.List()
}

func (r *companyService) Read(cmpId int) (models.Company, error) {
	return r.cmpService.Read(cmpId)
}

func (r *companyService) Update(cmp *models.Company) (models.Company, error) {
	return r.cmpService.Update(cmp)
}

func NewCompanyService(icmp CompanyServiceInt) CompanyServiceInt {
	return &companyService{cmpService: icmp}
}
