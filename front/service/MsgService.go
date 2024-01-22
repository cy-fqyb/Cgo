package service

import "Cgo/front/models"

type msgService struct{}

var MsgService = msgService{}

// 获取用户未读消息数量
func (msgService) GetMsgNum(user_id string) ([]models.MsgNum, error) {
	results, err := MsgDao.GetMsgNum(user_id)
	if err != nil {
		return nil, err
	}
	return results, nil
}
