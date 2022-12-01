package promotion

import "github.com/bububa/oceanengine/marketing-api/util"

// BudgetUpdateRequest 更新广告预算 API Request
type BudgetUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Data 批量修改，包含广告id和出价
	Data []BudgetUpdateData `json:"data,omitempty"`
}

// BudgetUpdateData 修改信息
type BudgetUpdateData struct {
	// PromotionID 广告ID，广告ID需要属于广告主
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// Budget 预算，单位“元”，精度：两位小数。
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r BudgetUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
