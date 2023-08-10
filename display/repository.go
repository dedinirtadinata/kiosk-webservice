package display

import "github.com/dedinirtadinata/kiosk-webservice/display/model"

/**
 * Created by S DEDI NIRTADINATA on 10/08/23
 */

type Repository interface {
	CreateData(data model.DisplayModel) (int, error)
	GetDataById(Id int) (*model.DisplayModel, error)
	Update(Id int, dataUpdated model.DisplayModel) error
	Delete(Id int) error
	GetAllData() ([]*model.DisplayModel, error)
	GetByCondition(query interface{}, args ...interface{}) ([]*model.DisplayModel, error)
}
