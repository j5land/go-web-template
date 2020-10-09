package config

import (
	"fmt"
	"go-web-template/pkg/utils"
	"testing"
)

func TestGetSourcePath(t *testing.T){
	path := utils.GetSourcePath()
	fmt.Println(path)
}
