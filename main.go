package testbe

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/FarhanRizkiM/backend-ai/url"
)

func init() {
	functions.HTTP("Mauls", url.Web)
}