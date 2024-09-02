package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

// detectFace detects faces in the provided image
func DetectFace(ctx context.Context, image []byte, client *rekognition.Client) error {
	input := &rekognition.DetectFacesInput{
		Image: &types.Image{
			Bytes: image,
		},
		Attributes: []types.Attribute{types.AttributeDefault},
	}

	result, err := client.DetectFaces(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to detect faces: %w", err)
	}

	if len(result.FaceDetails) == 0 {
		return fmt.Errorf("no face detected")
	}

	fmt.Println("Face detected successfully!")
	return nil
}

// compareFaces compares two faces and returns the similarity score
func CompareFaces(ctx context.Context, sourceImage, targetImage []byte, client *rekognition.Client) (float32, error) {
	input := &rekognition.CompareFacesInput{
		SourceImage: &types.Image{
			Bytes: sourceImage,
		},
		TargetImage: &types.Image{
			Bytes: targetImage,
		},
		SimilarityThreshold: aws.Float32(70.0),
	}

	result, err := client.CompareFaces(ctx, input)
	if err != nil {
		return 0, fmt.Errorf("error comparing faces: %w", err)
	}

	if len(result.FaceMatches) == 0 {
		return 0, fmt.Errorf("no matched faces")
	}

	return *result.FaceMatches[0].Similarity, nil
}

// ifErrorExit logs an error and exits if the error is not nil
func IfErrorExit(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
