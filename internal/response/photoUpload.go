package response

type BatchUploadResponse struct {
	BatchId string `json:"batch_id"`
	Status string  `json:"status"`
	Uploaded []UploadedPhotoResponse `json:"uploaded"`
	Failed []FailedPhotoResponse   `json:"failed"`
}

type FailedPhotoResponse struct {
    Filename string   `json:"filename"`
    Errors map[string]string `json:"error"`
}

type UploadedPhotoResponse struct {
    PhotoID  int64  `json:"photo_id"`
    Filename string `json:"filename"`
    Status   string `json:"status"` // e.g., "processing"
}