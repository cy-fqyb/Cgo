package dao

import (
	"Cgo/front/models"
	"Cgo/global"
)

type msgDao struct{}

var MsgDao = msgDao{}

// 获取用户未读消息数量
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

// 获取用户和指定好友或群组2天内的聊天记录
func (msgDao) GetMsgHistory(from_id, user_id string) ([]models.MsgHistory, error) {
	var results []models.MsgHistory
	if err := global.DB.Table("msg as m").
		Select("m.id, m.from_id, m.to_id, m.msg, m.create_time, m.is_read, u.name, u.avatar").
		Joins("left join users as u on m.from_id = u.id").
		Where("m.create_time > DATE_SUB(NOW(), INTERVAL 2 DAY) and ((m.from_id = ? and m.to_id = ?) or (m.from_id = ? and m.to_id = ?))", from_id, user_id, user_id, from_id).
		Order("m.create_time asc").
		Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// 获取用户好友消息列表
func (msgDao) GetMsgList(user_id string) ([]models.MsgList, error) {
	var results []models.MsgList
	subQuery := global.DB.Table("msg").
		Select("from_id, COUNT(*) AS message_count").
		Where("DATE(create_time) = CURDATE()").
		Group("from_id")

	err := global.DB.Table("users u").
		Select("u.id as friend_id, u.name as friend_name, msg.msg as content, msg.create_time AS time, u.avatar as avatar, daily_message_counts.message_count as msg_count").
		Joins("left join msg on msg.from_id = u.id").
		Joins("LEFT JOIN (?) AS daily_message_counts ON u.id = daily_message_counts.from_id", subQuery).
		Where("msg.to_id = ? AND DATE(msg.create_time) = CURDATE() AND msg.create_time = (SELECT MAX(create_time) FROM msg WHERE msg.from_id = u.id)", user_id).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
