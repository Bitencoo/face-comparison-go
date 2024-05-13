
# Face Detection and Comparison using AWS Rekognition

This Go program demonstrates the use of AWS Rekognition to detect and compare faces in images. It utilizes the AWS SDK for Go v2 to interact with AWS Rekognition, a powerful service that enables image and video analysis. The script loads two images from the filesystem, checks each for faces, and then compares these faces to determine a similarity score.

## Features

- **Face Detection**: Validates that each input image contains a face.
- **Face Comparison**: Compares two faces to determine how similar they are.

## Prerequisites

- Go 1.x or higher
- AWS account and AWS CLI configured
- AWS IAM permissions sufficient to access AWS Rekognition

## Setup and Installation

1. **Install Go** (if not already installed):
   ```bash
   https://golang.org/doc/install
   ```

2. **Clone the repository** (or just copy the script):
   ```bash
   git clone https://yourrepository.git
   cd your-repository-folder
   ```

3. **Set up AWS CLI and configure your AWS credentials** (if not already done):
   ```bash
   aws configure
   ```
   This will require your AWS Access Key ID, Secret Access Key, and default region.

4. **Install required Go packages**:
   ```bash
   go get github.com/aws/aws-sdk-go-v2/aws
   go get github.com/aws/aws-sdk-go-v2/config
   go get github.com/aws/aws-sdk-go-v2/service/rekognition
   ```

## Usage

1. **Prepare your images**:
   Place the images you want to compare in the same directory as the script, or modify the `loadImage` function to point to their paths.

2. **Run the script**:
   ```bash
   go run main.go
   ```

   The program will load the images, check each for faces, and compare the faces if detected. The output will indicate whether faces were detected and, if so, the similarity percentage between the faces.

## Important Functions

- `loadImage(filename string) []byte`: Reads an image file and converts it into a byte slice for processing.
- `detectFace(ctx context.Context, image []byte, client *rekognition.Client) error`: Detects faces in the provided image using AWS Rekognition.
- `compareFaces(ctx context.Context, sourceImage []byte, targetImage []byte, client *rekognition.Client) (float32, error)`: Compares two faces and returns the similarity score.

## Error Handling

Errors are handled using a custom function `ifErrorPanic(err error)` which will stop the program and display an error message if an operation fails. This approach is straightforward and suitable for demonstration purposes but may be adapted for more nuanced error handling in a production environment.

