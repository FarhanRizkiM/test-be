package testbe

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/FarhanRizkiM/test-be/url"
)

func init() {
	functions.HTTP("Mauls", url.Web)
}