package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//Creating new Rekognition Client
	client := rekognition.NewFromConfig(cfg)

	//Loading the images to be compared
	sourceImage := loadImage("a.png")
	err = detectFace(context.TODO(), sourceImage, client)
	ifErrorPanic(err)

	targetImage := loadImage("b.png")
	err = detectFace(context.TODO(), targetImage, client)
	ifErrorPanic(err)
	
	res, err := compareFaces(context.TODO(), sourceImage, targetImage, client)
	ifErrorPanic(err)

	fmt.Println("Faces Matched!")
	fmt.Printf("Similarity Between Faces: %.2f%% \n", res)
}


// loadImage loads an image file and returns the byte slice
func loadImage(filename string) []byte {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Error opening file: ", err)
        return nil
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        fmt.Println("Error getting file info: ", err)
        return nil
    }

    data := make([]byte, stat.Size())
    _, err = file.Read(data)
    if err != nil {
        fmt.Println("Error reading file: ", err)
        return nil
    }

    return data
}

func detectFace(ctx context.Context, image []byte, client *rekognition.Client) error {
	input := rekognition.DetectFacesInput{
		Image : &types.Image{
			Bytes: image,
		},
		Attributes: []types.Attribute{types.AttributeDefault},
	}

	result, err := client.DetectFaces(ctx, &input)

	if err != nil {
		err = fmt.Errorf("Failed to detect faces!\nCause of error: %w", err)
		return err
	}

	if(len(result.FaceDetails) == 0) {
		err = fmt.Errorf("No face was detected!")
		return err
	}

	fmt.Println("Face detected Succesfully!")

	return nil
}

func compareFaces(ctx context.Context, sourceImage []byte, targetImage []byte, client *rekognition.Client) (float32, error) {
	input := &rekognition.CompareFacesInput{
        SourceImage: &types.Image{
            Bytes: sourceImage,
        },
        TargetImage: &types.Image{
            Bytes: targetImage,
        },
        SimilarityThreshold: aws.Float32(70.0),
    }

	result, err := client.CompareFaces(context.TODO(), input)

	if err != nil {
		err = fmt.Errorf("Error Comparing Faces.\nCause: %w", err)
	}

	fmt.Println("Success!")

	if(len(result.UnmatchedFaces) > 0) {
		err = fmt.Errorf("Unmatched Faces")
	}

	if(len(result.FaceMatches) > 0) {
		return *result.FaceMatches[0].Similarity, nil 
	}

	return 0, err
}

func ifErrorPanic(err error) {
	if(err != nil) {
		panic(err)
	}
}