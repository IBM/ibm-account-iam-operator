package common

import (
	"os"
	"strings"
)

var (
	ImageList = []string{"MCSP_UTILS_IMAGE", "ACCOUNT_IAM_APP_IMAGE", "IM_CONFIG_JOB_IMAGE"}
)

func ReplaceImages(resource string) (result string) {
	result = resource
	for _, imageName := range ImageList {
		result = strings.ReplaceAll(result, imageName, GetImage(imageName))
	}
	return result
}

func GetImage(imageName string) string {
	image, _ := os.LookupEnv(imageName)
	return image
}
