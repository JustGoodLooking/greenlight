package request

import (
	"mime/multipart"

	"greenlight.goodlooking.com/internal/data"
	"greenlight.goodlooking.com/internal/validator"
)

type PhotoUploadInput struct {
    Photo *data.Photo
    File  multipart.File
}


func ValidatePhotoInput(v *validator.Validator, input *PhotoUploadInput) {
	v.Check(input.Photo.Filename != "", "filename", "must be provided")
	    // 驗證 MimeType
		allowedMimeTypes := map[string]bool{
			"image/jpeg": true,
			"image/png":  true,
			"image/webp": true,
			"image/heic": true,
		}
		v.Check(allowedMimeTypes[input.Photo.MimeType], "mime_type", "must be a supported image type")
	
		// 驗證檔案大小（限制 15 MB 範例）
		maxSize := int64(15 << 20)
		v.Check(input.Photo.Size <= maxSize, "size", "must be less than 10MB")
	
		// 驗證 magic hex
		buffer := make([]byte, 512)
		_, err := input.File.Read(buffer)
		if err != nil {
			v.AddError("file", "unable to read file for validation")
			return
		}

	
		if !validateMagicHex(buffer) {
			v.AddError("file", "unsupported or invalid file format")
		}
}


func validateMagicHex(buffer []byte) bool {
    // JPEG
    if len(buffer) >= 3 && buffer[0] == 0xFF && buffer[1] == 0xD8 && buffer[2] == 0xFF {
        return true
    }
    // PNG
    if len(buffer) >= 8 &&
        buffer[0] == 0x89 && buffer[1] == 0x50 &&
        buffer[2] == 0x4E && buffer[3] == 0x47 &&
        buffer[4] == 0x0D && buffer[5] == 0x0A &&
        buffer[6] == 0x1A && buffer[7] == 0x0A {
        return true
    }
    // WEBP
    if len(buffer) >= 12 &&
        buffer[0] == 'R' && buffer[1] == 'I' && buffer[2] == 'F' && buffer[3] == 'F' &&
        buffer[8] == 'W' && buffer[9] == 'E' && buffer[10] == 'B' && buffer[11] == 'P' {
        return true
    }
    // HEIC (簡易偵測)
    if len(buffer) >= 12 &&
        buffer[4] == 'f' && buffer[5] == 't' && buffer[6] == 'y' && buffer[7] == 'p' &&
        buffer[8] == 'h' && buffer[9] == 'e' && buffer[10] == 'i' && buffer[11] == 'c' {
        return true
    }
    return false
}