package internal_test

import (
	"fmt"
	"testing"

	"github.com/JulianH99/clone/internal"
)

func TestReadConfig(t *testing.T) {
	config := internal.GetConfig()
	fmt.Println(config)
}
