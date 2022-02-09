package image_service

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type ImageService interface {
	GenerateImage([]byte, string) ([]byte, error)
}

type service struct{}

func NewService() ImageService {
	return &service{}
}

func (s *service) GenerateImage(j []byte, fileName string) ([]byte, error) {
	img, err := jpeg.Decode(bytes.NewReader(j))
	if err != nil {
		return nil, err
	}

	var m image.Image
	bounds := img.Bounds()
	if bounds.Dx() > bounds.Dy() {
		m = resize.Resize(256, 0, img, resize.Lanczos3)
	} else {
		m = resize.Resize(0, 256, img, resize.Lanczos3)
	}

	out, err := os.Create(fmt.Sprintf("%s.png", fileName))
	if err != nil {
		return nil, err
	}

	defer out.Close()

	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, m); err != nil {
		return nil, err
	}

	if _, err := io.Copy(buf, out); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
