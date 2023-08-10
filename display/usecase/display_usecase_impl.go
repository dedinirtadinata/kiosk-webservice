package display

import (
	"github.com/dedinirtadinata/kiosk-webservice/display"
	"github.com/dedinirtadinata/kiosk-webservice/display/model"
)

/**
 * Created by S DEDI NIRTADINATA on 10/08/23
 */

type displayUsecase struct {
	repo display.Repository
}

func (d displayUsecase) CreateData(title string, description string, bgImage string, cardImage string, url string) (int, error) {
	return d.repo.CreateData(model.DisplayModel{
		Title:         title,
		Description:   description,
		BackgroundURL: bgImage,
		CardImageURL:  cardImage,
		Url:           url,
	})
}

func (d displayUsecase) GetDataById(Id int) (*model.DisplayModel, error) {
	return d.repo.GetDataById(Id)
}

func (d displayUsecase) Update(Id int, title string, description string, bgImage string, cardImage string, url string) error {
	return d.repo.Update(Id, model.DisplayModel{
		Title:         title,
		Description:   description,
		BackgroundURL: bgImage,
		CardImageURL:  cardImage,
		Url:           url,
	})
}

func (d displayUsecase) Delete(Id int) error {
	return d.repo.Delete(Id)
}

func (d displayUsecase) GetAllData() ([]*model.DisplayModel, error) {
	return d.repo.GetAllData()
}

func NewdisplayUsecase(r display.Repository) display.Usecase {
	return &displayUsecase{
		repo: r,
	}
}
