package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"greenlight.goodlooking.com/internal/data"
	"greenlight.goodlooking.com/internal/request"
	"greenlight.goodlooking.com/internal/response"
	"greenlight.goodlooking.com/internal/validator"
)

func (app *application) signHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sign request...")
	r.ParseMultipartForm(32 << 20)

	// check total size
	err := r.ParseMultipartForm(MaxUploadSize)
	if err != nil {
		app.requestEntityTooLargeResponse(w, r)
		return
	}

	// parse file
	uploadedPhotos := r.MultipartForm.File["photo"]

	if len(uploadedPhotos) == 0 {
		app.noFileUploaded(w, r)
		return
	}

	if len(uploadedPhotos) > 5 {
		app.tooManyFileUploaded(w, r)
		return
	}

	// 產生batch id
	batchId := "123"
	var uploadedPhotosResp []response.UploadedPhotoResponse
	var failedPhotosResp []response.FailedPhotoResponse

	for _, fileHeader := range uploadedPhotos {
		uploadedPhoto, err := fileHeader.Open()


		if err != nil {
			failedPhotoResp := response.FailedPhotoResponse{
				Filename: fileHeader.Filename,
				Errors: map[string]string{
					"file": "failed to open file: " + err.Error(),
				},
			}
			failedPhotosResp = append(failedPhotosResp, failedPhotoResp)
			continue
		}

		defer uploadedPhoto.Close()

		// 包成photo struct
		photo := &data.Photo{
			ID:           1,
			UserID:       1,
			BatchID:      batchId,
			UploadID:     batchId,
			Filename:     fileHeader.Filename,
			Size:         fileHeader.Size,
			MimeType:     fileHeader.Header.Get("Content-Type"),
			URL:          "",
			ThumbnailURL: "",
			Status:       "processing", // processing, done, failed
			RetryCount:   0,
			ErrorMessage: "",
			CreatedAt:    time.Now(),
			CompletedAt:  nil,
		}

		input := &request.PhotoUploadInput{
			Photo: photo,
			File: uploadedPhoto,
		}

		// validator
		v := validator.New()
		if request.ValidatePhotoInput(v, input); !v.Valid() {

			failedPhotoResp := response.FailedPhotoResponse{
				Filename: fileHeader.Filename,
				Errors: v.Errors,
			}
			// append 進去response fail
			failedPhotosResp = append(failedPhotosResp, failedPhotoResp)
			continue

		}
		input.File.Seek(0, io.SeekStart)
		uploadedPhotoResponse := response.UploadedPhotoResponse{
			PhotoID: 123,
			Filename: fileHeader.Filename,
			Status: "processing",
		}

		uploadedPhotosResp = append(uploadedPhotosResp, uploadedPhotoResponse)

		// 沒問題insert db
		// append 進去uploaded的地方

	}

	resp := response.BatchUploadResponse{
		BatchId: "hahaa",
		Status:   "processing",
		Uploaded: uploadedPhotosResp,
		Failed:   failedPhotosResp,
	}

	app.writeJSON(w, http.StatusOK, envelope{"upload": resp}, nil)
}

func (app *application) createKeypairHandler(w http.ResponseWriter, r *http.Request) {
	// keypair, pri, err := app.models.Keypair.New(123, "123", "123")

	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// app.keyStore.Set(keypair.ID, pri)

	// err = app.writeJSON(w, http.StatusOK, envelope{"movie": keypair}, nil)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }
}

func (app *application) getKeypairHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get keypairs success")
}
