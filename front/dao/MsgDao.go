package dao

import (
	"Cgo/front/models"
	"Cgo/global"
)

type msgDao struct{}

var MsgDao = msgDao{}

// 获取用户未读消息数量
// select u.id, u.name, count(*) as msg_count
// from msg as m
//
//	inner join users as u on m.from_id = u.id
//
// WHERE
//
//	m.create_time > DATE_SUB(NOW(), INTERVAL 3 DAY)
//
// group by
//
//	u.id,
//	u.name
//
// order by msg_count desc;
func (msgDao) GetMsgNum(user_id string) ([]models.MsgNum, error) {
	var results []models.MsgNum
	if err := global.DB.Table("msg as m").
		Select("u.id, u.name, count(*) as msg_count, u.avatar").
		Joins("left join users as u on m.from_id = u.id").
		Where("m.to_id = ? and m.is_read = 0", user_id).
		Group("u.id, u.name").
		Order("msg_count desc").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}
