package display

import "github.com/dedinirtadinata/kiosk-webservice/display/model"

/**
 * Created by S DEDI NIRTADINATA on 10/08/23
 */

type Usecase interface {
	CreateData(title string, description string, bgImage string, cardImage string, url string) (int, error)
	GetDataById(Id int) (*model.DisplayModel, error)
	Update(Id int, title string, description string, bgImage string, cardImage string, url string) error
	Delete(Id int) error
	GetAllData() ([]*model.DisplayModel, error)
}
