package imagekit

import (
	"context"
	"errors"
)

//
// REQUESTS
//

type RemoveTagsRequest struct {
	// FileIDs is the list of unique ID of the uploaded files.
	FileIDs []string `json:"fileIds"`
	// Tags is an array of tags to add on these files.
	Tags []string `json:"tags"`
}

//
// METHODS
//

// RemoveTags from multiple files in a single request.
func (s *MediaService) RemoveTags(ctx context.Context, r *RemoveTagsRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/removeTags", nil, requestTypeAPI)
	if err != nil {
		return err
	}

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
