package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
)

type FileIdentity struct {
  Filename string
  Size int
}

func IsImage(fileName string) bool {
  extension := strings.ToLower(strings.Split(fileName, ".")[1])
  return extension == "jpg" || extension == "png" || extension == "jpeg"
}

func GetImageExtension(fileName string) string {
  return strings.ToLower(strings.Split(fileName, ".")[1])
}

func GenerateImageName(userId uuid.UUID, extension string) string {
  currentTime := time.Now().UnixNano()
  newName := fmt.Sprintf("%s%s%s%s%s", userId, "-", strconv.Itoa(int(currentTime)), ".", extension)
  return newName
}

func DecodedImage(file dto.FileHandler, id uuid.UUID) (FileIdentity, error) {

  fileIdentity := FileIdentity{}

  // Handler image Size max 1mb
  const maxSize = 1048576
  if file.FileHeader.Size > maxSize {
    return fileIdentity, errors.New("File to large")
  }

  // Handler is type image available or not
  if !IsImage(file.FileHeader.Filename) {
    return fileIdentity, errors.New("Image type is not available, only jpg, png & jpeg")
  }

  imageExtension := GetImageExtension(file.FileHeader.Filename)
  newImageName := GenerateImageName(id , imageExtension)

  fileIdentity.Filename = newImageName
  fileIdentity.Size = int(file.FileHeader.Size)

  return fileIdentity, nil
} 
