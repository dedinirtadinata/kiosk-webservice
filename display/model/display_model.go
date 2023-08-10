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
	Deleted       int    `db:"dp_deleted" json:"-"`
}

func (u *DisplayModel) TableName() string {
	return "display"
}

type RequestCreate struct {
	Title         string `db:"dp_title" json:"title"`
	Description   string `db:"dp_description" json:"description"`
	BackgroundURL string `db:"dp_bgimage" json:"bg_url"`
	CardImageURL  string `db:"dp_cardimage" json:"card_url"`
	Url           string `db:"dp_url" json:"url"`
}
