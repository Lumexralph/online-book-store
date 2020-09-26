// Package impl has the ImageUploadService grpc server implementation.
package services

import (
	"bytes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"productservice/internal/domain"
	repo "productservice/internal/repository"
)

// ProductImageService is an implementation of ProductImageServiceServer gRPC server.
type ProductImageService struct {
	store repo.ImageRepository // product image store.
}

// NewProductImageService creates a new product image service.
func NewProductImageService(r repo.ImageRepository, ) *ProductImageService {
	return &ProductImageService{store: r}
}

// (1 << 20) * 1 == 1 << 20 (1MB)
const maxImageSize = 5 << 20 // 5MB

// UploadImage method will store a product image.
func (p *ProductImageService) UploadImage(stream domain.ProductImageService_UploadImageServer) error {
	// First receive the first request, which contains the metadata information of the image and product id.
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}
	// extract the metadata
	productID := req.GetInfo().GetProductId()
	imageFormat := req.GetInfo().GetImageFormat()
	log.Printf("uploadImage: upload image format %q, productId %d\n", imageFormat, productID)

	// TODO: find the productID if it exists, if mot, return an error

	// Start receiving the image chunks data and store it in-memory buffer/queue
	var imageData bytes.Buffer
	imageSize := 0

	for {
		log.Println("waiting to receive remaining data: ")

		req, err := stream.Recv()
		if err == io.EOF { // means we are done reading all the data
			log.Println("finished reading data.")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive image data: %v", err)
		}

		// get every chunk of the data till there is mo more data
		imageChunk := req.GetImageChunk()
		imageChunkSize := len(imageChunk)
		log.Printf("uploadImage: received %d size data chunk", imageChunkSize)
		imageSize += imageChunkSize
		if imageSize > maxImageSize {
			return status.Errorf(codes.InvalidArgument, "image size %d, should not be larger than 5MB", imageSize)
		}
		_, err = imageData.Write(imageChunk)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to create image chunk: %v", err)
		}
	}

	// save the collected image data
	imageURL, err := p.store.SaveImage(productID, imageData, imageFormat)
	if err != nil {
		return status.Errorf(codes.Internal, "save-image: failed to save image to the store")
	}
	// return a single response
	res := &domain.UploadImageResponse{ImageUrl: imageURL}
	err = stream.SendAndClose(res)
	if err != nil {
		return status.Errorf(codes.Unknown, "upload-image: failed to send response: %v", err)
	}

	log.Printf("upload-image: product image saved id: %d, imageURL: %q\n", productID, imageURL)
	return nil
}
