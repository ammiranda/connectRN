package rest_api

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ammiranda/connectRN/pkg/image_service"
	"github.com/ammiranda/connectRN/pkg/rest_api/models/request"
	"github.com/ammiranda/connectRN/pkg/rest_api/models/response"
	"github.com/ammiranda/connectRN/pkg/user_service"
	"github.com/stretchr/testify/require"
)

func TestUsersRoute_Success(t *testing.T) {
	u := user_service.NewService()
	i := image_service.NewService()
	r := NewRouter(u, i)
	w := httptest.NewRecorder()

	users := request.UserRequestBody{
		request.User{
			Name:      "Alex",
			UserID:    1,
			DOB:       "1987-02-22",
			CreatedOn: 1642612034,
		},
	}

	json_payload, err := json.Marshal(users)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/users", bytes.NewBuffer(json_payload))
	require.NoError(t, err)
	r.ServeHTTP(w, req)

	resp := response.UserResponse{}
	require.Equal(t, 201, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "Alex", resp[0].Name)
	require.Equal(t, 1, resp[0].UserID)
	require.Equal(t, "Sunday", resp[0].DOBDayOfWeek)
	require.Equal(t, "2022-01-19T12:07:14-05:00", resp[0].CreatedOn)
}

func TestImageRoute_Success(t *testing.T) {
	u := user_service.NewService()
	i := image_service.NewService()
	r := NewRouter(u, i)
	w := httptest.NewRecorder()

	img := image.NewRGBA(image.Rect(0, 0, 1000, 300))
	outputFile, err := os.Create("test.jpg")
	require.NoError(t, err)

	defer outputFile.Close()

	jpeg.Encode(outputFile, img, &jpeg.Options{})

	var requestBody bytes.Buffer
	multiPartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multiPartWriter.CreateFormFile("image", "test.jpg")
	require.NoError(t, err)

	multiPartWriter.Close()

	_, err = io.Copy(fileWriter, outputFile)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/images", &requestBody)
	require.NoError(t, err)
	req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	r.ServeHTTP(w, req)

	require.Equal(t, 200, w.Code)
}
