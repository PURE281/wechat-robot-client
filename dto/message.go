package dto

type MessageCommonRequest struct {
	MessageID int64 `form:"message_id" json:"message_id" binding:"required"`
}

type SendMessageCommonRequest struct {
	ToWxid string `form:"to_wxid" json:"to_wxid" binding:"required"`
}

type SendTextMessageRequest struct {
	SendMessageCommonRequest
	Content string   `form:"content" json:"content" binding:"required"`
	At      []string `form:"at" json:"at"`
}

type SendMusicMessageRequest struct {
	SendMessageCommonRequest
	Song string `form:"song" json:"song" binding:"required"`
}

type TextMessageItem struct {
	Nickname  string `json:"nickname"`
	Message   string `json:"message"`
	CreatedAt int64  `json:"created_at"`
}

type SendFileMessageRequest struct {
	ToWxid      string `form:"to_wxid" json:"to_wxid" binding:"required"`
	Filename    string `form:"filename" json:"filename" binding:"required"`
	FileHash    string `form:"file_hash" json:"file_hash" binding:"required"`
	FileSize    int64  `form:"file_size" json:"file_size" binding:"required"`
	ChunkIndex  int    `form:"chunk_index" json:"chunk_index"`
	TotalChunks int    `form:"total_chunks" json:"total_chunks" binding:"required"`
}
