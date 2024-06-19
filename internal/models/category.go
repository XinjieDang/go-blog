package models

type Category struct {
	BasicModel
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	IsActive  int64  `json:"is_active" gorm:"default:0"`  //0:正常 1:删除
	SortOrder int64  `json:"sort_order" gorm:"default:0"` //排序值
}
