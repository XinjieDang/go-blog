package models

import "gorm.io/gorm"

type BlogMenu struct {
	MenuName   string `json:"menu_name"`   // 菜单项名称
	URL        string `json:"url"`         // 菜单项链接地址
	ParentID   *int   `json:"parent_id"`   // 父菜单ID，nil表示顶级菜单
	OrderIndex int    `json:"order_index"` // 排序索引
	IsVisible  bool   `json:"is_visible"`  // 是否可见
	gorm.Model
}
