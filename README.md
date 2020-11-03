# image-ai

A CLI tool for analyzing the contents of an image. 

**Note:** This repo is currently a work-in-progress so the information may not be complete, accurate, or final.

## Technologies and Packages Used
- [Golang](https://golang.org/) - version >= 1.11 (must support Go modules)
  - [Cobra](https://github.com/spf13/cobra)
  - [AWS SDK](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html)
- [Amazon Rekognition](https://aws.amazon.com/rekognition/)
- [Amazon S3](https://aws.amazon.com/s3/)
- [Amazon Dynamo DB](https://aws.amazon.com/dynamodb/)

## Available Commands

*   `detect-objects` 

    **Description** Detects objects in an image.

    **Example** `./image-ai detect-objects --image-path ~/pictures/photo1.png`

    **Required flags** `--image-path` [relative path to image file]

*   `describe-face`

    **Description** Detects the largest face in the image and returns a description of the face.

    **Example** `./image-ai describe-face --image-path ~/pictures/photo1.png`

    **Required flags** `--image-path` [relative path to image file]

*   `compare-faces`

     **Description** Compares the largest face in the source and the target image to identify if they are of the same person. The target image can be in a collection or supplied by the user.

    **Example** `./image-ai compare-faces --src-image-path ~/pictures/photo1.png --target-image-path ~/pictures/photo2.png`

    **Required flags** `--src-image-path` [relative path to image file]

    **Optional flags** (must select one)
    - `--target-image-path` [relative path to image file]
    - `--collection` [collection name]

*   `collection`

    Use the following commands to manage a collection of images.

    - `list`

        **Description** List collections in the AWS account

        **Example** `./image-ai collection list`

    - `create`

        **Description** Create a new collection

        **Example** `./image-ai collection create [collection name]`

        **Required args**  `--collection-name` [name of collection]

    - `delete`

        **Description** Delete a collection

        **Example** `./image-ai collection delete [collection name]`

        **Required args** `--collection-name` [name of collection]]

    - `add-image`

        **Description** Add an image to a collection

        **Example** `./image-ai collection add-image --collection-name myCollection --image-path ~/pictures/photo1.png`

        **Required args**
        - `--collection-name` [name of collection]
        - `--image-path` [relative path to image file]

    - `delete-image` (TBD - might change)

        **Description** Delete an image from a collection

        **Example** `./image-ai collection delete-image --collection-name myCollection --file-name photo1.png`

       **Required args**
        - `--collection-name` [name of collection]
        - `--file-name` [file name to match the S3 key for the image file]

### Global Flags
`--aws-region` - Default value is `us-west-2`

`--aws-profile` - Default value is `default` 

## Setup
1. Install [git](https://git-scm.com/downloads) and [golang](https://golang.org/doc/install), if not installed already. 
2. Create an AWS account if you don't have one.
3. This program authenticates with your AWS account credentials stored in `~/.aws/credentials`. If you don't have this setup yet, install the AWS CLI and run `aws configure`. 
For more information on configuration, see [AWS CLI configuration documentation](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-config)
4. See [here](https://docs.aws.amazon.com/rekognition/latest/dg/limits.html) for image formats/sizes accepted by Amazon Rekognition

## Running the Program
1. Clone this repository and `cd` into the repo root directory.
2. Run `go build` in your terminal. 
This will create an executable`image-ai` file.
3. Run `./image-ai [additional arugments here]` to run the program.
4. For a faster development iteration, instead of steps 2 and 3, you can also run `go run main.go [additional arugments here]`
