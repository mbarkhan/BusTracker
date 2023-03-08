package services

import "githubmbarkhanBusTracker/auth/models"

type PathesServiceInt interface {
	Create(*models.Pathes) (models.Pathes, error)
	Read(int) (models.Pathes, error)
	Update(*models.Pathes) (models.Pathes, error)
	Delete(int) error
	List() ([]models.Pathes, error)
}

type PathesService struct {
	pathes PathesServiceInt
}

func (r *PathesService) Create(path *models.Pathes) (models.Pathes, error) {
	return r.pathes.Create(path)
}

func (r *PathesService) Delete(pathID int) error {
	return r.pathes.Delete(pathID)
}

func (r *PathesService) List() ([]models.Pathes, error) {
	return r.pathes.List()
}

func (r *PathesService) Read(cmpId int) (models.Pathes, error) {
	return r.pathes.Read(cmpId)
}

func (r *PathesService) Update(psnPath *models.Pathes) (models.Pathes, error) {
	return r.pathes.Update(psnPath)
}

func NewPathService(ipth PathesServiceInt) PathesServiceInt {
	return &PathesService{pathes: ipth}
}
