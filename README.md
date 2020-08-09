# image-ai

A CLI tool for analyzing the contents of an image. This repo is currently a WIP so everything may not be complete, accurate, or finalized.

## Technologies and Packages Used
- [Golang](https://golang.org/) - version >= 1.11 (must support Go modules)
  - [Cobra](https://github.com/spf13/cobra)
- [Amazon Rekognition](https://aws.amazon.com/rekognition/)
- [Amazon S3](https://aws.amazon.com/s3/)
- [Amazon Dynamo DB](https://aws.amazon.com/dynamodb/)

## Available Commands
*   `detect-objects`

    Detects objects in an image.

    ex `./image-ai detect-objects --image-path ~/pictures/photo1.png`

    _Required flags_
    - `--image-path` [relative path to image file]

*   `describe-face`

    Detects the largest face in the image and returns a description of the face.

    ex `./image-ai describe-face --image-path ~/pictures/photo1.png`

    _Required flags_
    - `--image-path` [relative path to image file]

*   `compare-faces`

    Takes a source and target image and copares the largest face from each image to see if they are the same person. The target image can be in a collection or supplied by the user.

    ex `./image-ai compare-faces --src-image-path ~/pictures/photo1.png --target-image-path ~/pictures/photo2.png`

    _Required flags_
    - `--src-image-path` [relative path to image file]

    Optional flags (must select one)
    - `--target-image-path` [relative path to image file]
    - `--collection` [collection name]

*   `collection`

    The below commands are for managing a collection of images.

    - `list`

        List collections in the AWS account

        ex. `./image-ai collection list`

    - `create`

        Create a new collection

        ex. `./image-ai collection create [collection name]`

        _Required args_
        - `collection name` name of collection

    - `delete`

        Delete a collection

        ex. `./image-ai collection delete [collection name]`

        _Required args_
        - `collection name` name of collection

    - `add-image`

        Add an image to a collection

        ex. `./image-ai collection add-image --collection-name myCollection --image-path ~/pictures/photo1.png`

        _Required flags_
        - `--collection-name` [name of collection]
        - `--image-path` [relative path to image file]

    - `delete-image` (TBD - might change)

        Delete an image from a collection

        ex. `./image-ai collection delete-image --collection-name myCollection --file-name photo1.png`

        _Required flags_
        - `--collection-name` [name of collection]
        - `--file-name` [file name to match the S3 key for the image file]

## Setup
1. Make sure you have [git](https://git-scm.com/downloads) and [golang](https://golang.org/doc/install) installed. Create an AWS account if you don't have one.
2. This program authenticates with your AWS account credentials stored in `~/.aws/credentials`. If you don't have this setup yet install the AWS CLI and run `aws configure`. More info [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-config)
3. See [here](https://docs.aws.amazon.com/rekognition/latest/dg/limits.html) for image formats/sizes accepted by Amazon Rekognition

## Running the Program
1. Clone this repository and `cd` into the repo root directory
2. Run `go build` in your terminal. This will create an executable`image-ai` folder
3. Run `./image-ai [additional arugments here]` to run the program
4. Instead of steps 3 and 4, you can also run `go run main.go [additional arugments here]` for faster development iteration
