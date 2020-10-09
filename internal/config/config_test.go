package config

import (
	"fmt"
	"golang-web-template/pkg/utils"
	"testing"
)

func TestGetSourcePath(t *testing.T){
	path := utils.GetSourcePath()
	fmt.Println(path)
}
