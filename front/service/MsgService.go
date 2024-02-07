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

// 获取用户未读消息列表
func (msgService) GetMsgHistory(from_id string, user_id string) ([]models.MsgHistory, error) {
	results, err := MsgDao.GetMsgHistory(from_id, user_id)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// 获取用户好友消息列表
func (msgService) GetMsgList(user_id string) ([]models.MsgList, error) {
	results, err := MsgDao.GetMsgList(user_id)
	if err != nil {
		return nil, err
	}
	return results, nil
}
