package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/klearwave/service-info/pkg/db"
	"github.com/klearwave/service-info/pkg/models"
	modelsv0 "github.com/klearwave/service-info/pkg/v0/models"
)

var (
	ErrMissingNodeImageId = errors.New("missing id parameter")
)

// CreateNodeImage defines the service for creating a new node image and storing in a database.
func (service *Service) CreateNodeImage(ctx context.Context, request *modelsv0.NodeImageRequestCreate) (*modelsv0.NodeImageResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	NodeImage := &modelsv0.NodeImage{}
	NodeImage.NodeImageInput = request.Body

	// create the node image
	if err := service.Database.Create(NodeImage); err != nil {
		return nil, err
	}

	return &modelsv0.NodeImageResponse{
		Body: *NodeImage,
	}, nil
}

// GetNodeImage defines the service for retrieving a specific node image from the database.
func (service *Service) GetNodeImage(ctx context.Context, request *modelsv0.NodeImageRequestGet) (*modelsv0.NodeImageResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	response := &modelsv0.NodeImageResponse{Body: modelsv0.NodeImage{}}

	if _, err := service.Database.FindBy("id", request.Id, &response.Body); err != nil {
		return nil, err
	}

	if response.Body.Id == db.MissingDatabaseID {
		response.Status = http.StatusNotFound

		return response, nil
	}

	return response, nil
}

// GetNodeImages defines the service for retrieving all node images from a database.
func (service *Service) GetNodeImages(ctx context.Context, request *modelsv0.NodeImageRequestList) (*modelsv0.NodeImagesResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	NodeImages := []modelsv0.NodeImage{}

	result := service.Database.Connection.Find(&NodeImages)
	if result.Error != nil {
		return nil, result.Error
	}

	return &modelsv0.NodeImagesResponse{
		Body: NodeImages,
	}, nil
}

// Delete defines the service for deleting a node image from a database.
func (service *Service) DeleteNodeImage(ctx context.Context, request *modelsv0.NodeImageRequestDelete) (*models.DeleteResponse, error) {
	service.Database.Lock.Lock()
	defer service.Database.Lock.Unlock()

	NodeImage := modelsv0.NodeImage{}

	if request.Id == db.MissingDatabaseID {
		return nil, ErrMissingNodeImageId
	}

	if _, err := service.Database.FindBy("id", request.Id, &NodeImage); err != nil {
		return nil, err
	}

	if err := service.Database.Delete(NodeImage.Id, NodeImage); err != nil {
		return nil, err
	}

	response := &models.DeleteResponse{}
	response.Body.Message = "Delete Success"

	return response, nil
}
