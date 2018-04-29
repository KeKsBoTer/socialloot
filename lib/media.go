package lib

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strings"

	"github.com/KeKsBoTer/socialloot/models"
	"github.com/astaxie/beego"
	"github.com/nfnt/resize"
)

type decodeImgFunc = func(io.Reader) (interface{}, error)
type encodeImgFunc = func(io.Writer, interface{}) error

const thumbnailSize = 255

func createMedia(file *string, decode decodeImgFunc, encode encodeImgFunc) (*models.Media, error) {
	reader := strings.NewReader(*file)
	buffer := new(bytes.Buffer)
	img, err := decode(reader)
	if err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot decode image")
	}
	mediaImage := models.Media{
		Type: models.MediaImage,
	}
	if err := encode(buffer, img); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot encode image")
	}
	mediaImage.File = buffer.String()
	buffer.Reset()
	var decImage image.Image
	if i, ok := img.(image.Image); ok {
		decImage = i
	} else if i, ok := img.(*gif.GIF); ok {
		decImage = i.Image[0]
	} else {
		panic("invalid image object")
	}

	// create small png thumbnail
	size := decImage.Bounds().Size()
	var width, height uint
	if size.X > size.Y {
		width = uint(thumbnailSize * float64(size.X) / float64(size.Y))
		height = thumbnailSize
	} else {
		height = uint(thumbnailSize * float64(size.Y) / float64(size.X))
		width = thumbnailSize
	}
	thumbnail := resize.Resize(width, height, decImage, resize.Lanczos3)
	if err := png.Encode(buffer, thumbnail); err != nil {
		beego.Error(err)
		return nil, errors.New("Cannot encode image")
	}
	mediaImage.Thumbnail = buffer.String()
	return &mediaImage, nil
}

func parseImage(file *string, fileType string) (*models.Media, error) {

	var encodeFunc encodeImgFunc
	var decodeFunc decodeImgFunc
	switch fileType {
	case "image/png":
		encodeFunc = func(w io.Writer, img interface{}) error {
			return png.Encode(w, img.(image.Image))
		}
		decodeFunc = func(r io.Reader) (interface{}, error) {
			return png.Decode(r)
		}
	case "image/jpg", "image/jpeg":
		encodeFunc = func(w io.Writer, img interface{}) error {
			return jpeg.Encode(w, img.(image.Image), nil)
		}
		decodeFunc = func(r io.Reader) (interface{}, error) {
			return jpeg.Decode(r)
		}
	case "image/gif":
		encodeFunc = func(w io.Writer, img interface{}) error {
			return gif.EncodeAll(w, img.(*gif.GIF))
		}
		decodeFunc = func(r io.Reader) (interface{}, error) {
			return gif.DecodeAll(r)
		}
	}
	return createMedia(file, decodeFunc, encodeFunc)
}
