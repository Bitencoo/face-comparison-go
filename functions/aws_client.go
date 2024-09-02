package functions

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

// initializeRekognitionClient initializes the AWS Rekognition client
func InitializeRekognitionClient(ctx context.Context) (*rekognition.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}
	return rekognition.NewFromConfig(cfg), nil
}
