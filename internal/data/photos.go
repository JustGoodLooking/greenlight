package data

import (
	"time"
)

// CREATE TABLE photo (
//     id BIGSERIAL PRIMARY KEY,
//     user_id BIGINT NOT NULL,
//     batch_id UUID,
//     upload_id UUID,
//     filename TEXT,
//     size BIGINT,
//     mime_type TEXT,
//     url TEXT,
//     thumbnail_url TEXT,
//     status TEXT, -- processing, done, failed
//     retry_count INT DEFAULT 0,
//     error_message TEXT,
//     created_at TIMESTAMP DEFAULT NOW(),
//     completed_at TIMESTAMP
// );

type Photo struct {
	ID           int64      `json:"id"`
	UserID       int64      `json:"user_id"`
	BatchID      string     `json:"batch_id"`
	UploadID     string     `json:"upload_id"`
	Filename     string     `json:"filename"`
	Size         int64      `json:"size"`
	MimeType     string     `json:"mime_type"`
	URL          string     `json:"url"`
	ThumbnailURL string     `json:"thumbnail_url"`
	Status       string     `json:"status"` // processing, done, failed
	RetryCount   int        `json:"retry_count"`
	ErrorMessage string     `json:"error_message"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"` // nullable
}

type UploadResponse struct {
	BatchID  string `json:"batch_id"`
	UploadID string `json:"upload_id"`
	Status   string `json:"status"` // usually "processing"
}

type PhotoStatusResponse struct {
	UploadID     string `json:"upload_id"`
	Status       string `json:"status"`
	URL          string `json:"url,omitempty"`
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

type BatchStatusResponse struct {
	BatchID string                `json:"batch_id"`
	Files   []PhotoStatusResponse `json:"files"`
}
