package dto

type WsDto struct {
	ToId uint64 `json:"to_id"`
	Msg  string `json:"msg"`
	Type string `json:"type"`
}
