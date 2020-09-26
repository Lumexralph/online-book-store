package repository

import (
	"bytes"
)

// ImageRepository interface that any image storage component implementation have to provide.
type ImageRepository interface {
	// SaveImage uploads image file somewhere on the server or on the cloud.
	// It returns the image storage path.
	SaveImage(productID uint64, imageData bytes.Buffer, imageFormat string) (string, error) // will handle the save/upload of our image.
}
