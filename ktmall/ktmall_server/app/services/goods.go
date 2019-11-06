package services

import (
	"errors"
	"ktmall/app/models"
	"ktmall/common/serializer"
	"math"

	"github.com/jinzhu/gorm"
)

type GoodsService struct {
	DB *gorm.DB
}

func (g GoodsService) GoodsList(categoryId, pageNo int, keyword string) ([]serializer.Data, error) {
	if categoryId == 0 && keyword == "" {
		return nil, errors.New("categoryId or keyword not found")
	}

	limit := 6
	offset := (pageNo - 1) * limit
	list := make([]*models.GoodsInfo, 0)
	allCount := 0
	maxPage := 1.0
	var query *gorm.DB

	// 按 category_id 搜索
	if categoryId != 0 {
		g.DB.Model(&models.GoodsInfo{}).Where("category_id = ?", categoryId).Count(&allCount)
		maxPage = math.Ceil(float64(allCount) / float64(limit))
		query = g.DB.Where("category_id = ?", categoryId).
			Offset(offset).
			Limit(limit).
			Find(&list).
			Order("id desc")
	} else {
		// 按 keyword 搜索
		keyword = "%%" + keyword + "%%"
		g.DB.Model(&models.GoodsInfo{}).Where("`desc` LIKE ?", keyword).Count(&allCount)
		maxPage = math.Ceil(float64(allCount) / float64(limit))
		query = g.DB.Where("`desc` LIKE ?", keyword).
			Offset(offset).
			Limit(limit).
			Find(&list).
			Order("id desc")
	}

	if err := query.Error; err != nil || len(list) == 0 {
		return []serializer.Data{}, err
	}

	var (
		ids     = make([]uint, len(list))
		results = make([]serializer.Data, len(list))
	)

	for i, v := range list {
		ids[i] = v.ID
		results[i] = v.Serialize()
	}

	skus := make([]*models.GoodsSku, 0)
	skuquery := g.DB.Where("goods_id IN (?)", ids).Find(&skus)
	if err := skuquery.Error; err != nil {
		return nil, err
	}

	for i, v := range list {
		ss := make([]serializer.Data, 0)
		for _, kv := range skus {
			if v.ID == kv.GoodsId {
				ss = append(ss, kv.Serialize())
			}
		}
		results[i]["goodsSku"] = ss
		results[i]["maxPage"] = maxPage
		results[i]["allCount"] = allCount
	}

	return results, nil
}
