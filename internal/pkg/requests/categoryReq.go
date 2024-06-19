package requests

import "dxj/internal/pkg/dto"

type CategoryListReq struct {
	Name          string `json:"name"` // 搜索名称
	dto.PageParam        //必须带包名导入分页参数
}
