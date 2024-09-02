package main

import (
	"context"
	"face-comparison-go/functions"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Initialize AWS Rekognition client
	client, err := functions.InitializeRekognitionClient(ctx)
	functions.IfErrorExit(err)

	// Load images to be compared
	sourceImage := functions.LoadImage("a.png")
	targetImage := functions.LoadImage("b.png")

	// Perform face detection and comparison
	err = functions.DetectFace(ctx, sourceImage, client)
	functions.IfErrorExit(err)

	err = functions.DetectFace(ctx, targetImage, client)
	functions.IfErrorExit(err)

	similarity, err := functions.CompareFaces(ctx, sourceImage, targetImage, client)
	functions.IfErrorExit(err)

	fmt.Printf("Faces Matched! Similarity Between Faces: %.2f%% \n", similarity)
}
