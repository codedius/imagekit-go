package imagekit

import (
	"context"
	"errors"
)

//
// REQUESTS
//

type CreateFolderRequest struct {
	// FolderName of the folder to create.
	//
	// All characters except alphabets and numbers will be replaced by an underscore i.e. _
	FolderName string `json:"folderName"`
	// ParentFolderPath where the new folder should be created.
	//
	// For root use / else containing/folder/
	// Note: If any folder(s) is not present in the parentFolderPath parameter, it will be automatically created. For example,
	// if you pass /product/images/summer, then product, images, and summer folders will be created if they don't already exist.
	ParentFolderPath string `json:"parentFolderPath"`
}

//
// METHODS
//

// CreateFolder will create a new folder.
//
// You can specify the folder name and location of the parent folder where this new folder should be created.
func (s *MediaService) CreateFolder(ctx context.Context, r *CreateFolderRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/folder", nil, requestTypeAPI)
	if err != nil {
		return err
	}

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
