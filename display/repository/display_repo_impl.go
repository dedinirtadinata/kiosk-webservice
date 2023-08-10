package display

/**
 * Created by S DEDI NIRTADINATA on 10/08/23
 */

import (
	"github.com/dedinirtadinata/kiosk-webservice/display"
	"github.com/dedinirtadinata/kiosk-webservice/display/model"
	"gorm.io/gorm"
)

type displayRepository struct {
	DB *gorm.DB
}

func (u displayRepository) CreateData(model model.DisplayModel) (int, error) {
	result := u.DB.Create(&model) // pass pointer of data to Create

	if result.Error != nil {
		return 0, result.Error
	}
	return model.ID, nil // returns error
}

func (u displayRepository) GetDataById(Id int) (*model.DisplayModel, error) {
	var result model.DisplayModel
	r := u.DB.First(&result, "id = ?", Id)

	if r.Error != nil {
		return nil, r.Error
	}
	return &result, nil
}

func (u displayRepository) Update(Id int, dataUpdated model.DisplayModel) error {
	var model model.DisplayModel
	u.DB.First(&model, Id)
	result := u.DB.Model(&model).UpdateColumns(dataUpdated)
	return result.Error
}

func (u displayRepository) Delete(Id int) error {
	var model model.DisplayModel
	u.DB.First(&model, Id)
	result := u.DB.Model(&model).Select("deleted").Updates(map[string]interface{}{"dp_deleted": 1})
	return result.Error
}

func (d displayRepository) GetAllData() (result []*model.DisplayModel, err error) {
	return d.GetByCondition("deleted", 0)
}

func (u displayRepository) GetByCondition(query interface{}, args ...interface{}) (result []*model.DisplayModel, err error) {
	r := u.DB.Where(query, args...).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return result, nil
}

func NewdisplayRepository(db *gorm.DB) display.Repository {
	return &displayRepository{
		DB: db,
	}
}
