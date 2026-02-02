package model

/**
 * Created by S DEDI NIRTADINATA on 10/08/23
 */

type DisplayModel struct {
	ID            int    `db:"dp_id;primaryKey" json:"id"`
	Title         string `db:"dp_title" json:"title"`
	Description   string `db:"dp_description" json:"description"`
	BackgroundURL string `db:"dp_bgimage" json:"bg_url"`
	CardImageURL  string `db:"dp_cardimage" json:"card_url"`
	Url           string `db:"dp_url" json:"url"`
	ApiKey        string `db:"dp_apikey" json:"api_key"` // Bearer token untuk auth (e.g. Grafana)
	Deleted       int    `db:"dp_deleted" json:"-"`
}

func (u *DisplayModel) TableName() string {
	return "display"
}

type RequestCreate struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	BackgroundURL string `json:"bg_url"`
	CardImageURL  string `json:"card_url"`
	Url           string `json:"url"`
	ApiKey        string `json:"api_key"` // Bearer token untuk auth (e.g. Grafana)
}
