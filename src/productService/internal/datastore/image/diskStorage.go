package image

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"os"
	"sync"
)

// DiskStore creates an interface to save image files to the disk,
// and store its information in memory.
type DiskStore struct {
	mutex       sync.RWMutex // help us handle concurrent situations.
	imageFolder string       // the directory to sore the images.
	images      map[string]*ImageInfo
}

type ImageInfo struct {
	ProductID uint64
	Path      string
	Format    string
}

// NewDiskStore will initialize a new disk store.
func NewDiskStore(imageFolder string) *DiskStore {
	return &DiskStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

// SaveImage generates a new random UUID for the image ID, make the path to store the image
// by joining the image folder, image ID, and image format.
func(ds *DiskStore) SaveImage(productID uint64, imageData bytes.Buffer, imageFormat string) (string, error) {
	imageID := uuid.New().String()
	imagePath := fmt.Sprintf("%s/%s%s", ds.imageFolder, imageID, imageFormat)
	// create the file to store the image data
	f, err := os.Create(imagePath)
	defer f.Close()
	if err != nil {
		return "", fmt.Errorf("SaveImage: cannot create image file: %w", err)
	}
	// write all the image data to the file
	_, err = imageData.WriteTo(f)
	if err != nil {
		return "", fmt.Errorf("SaveImage: cannot write image data to file: %w", err)
	}

	// there might be contest for accessing the images map
	ds.mutex.Lock()
	defer ds.mutex.Unlock()
	// create a new entry for the image
	ds.images[imageID] = &ImageInfo{
		ProductID: productID,
		Path: imagePath,
		Format: imageFormat,
	}
	return imagePath, nil
}
