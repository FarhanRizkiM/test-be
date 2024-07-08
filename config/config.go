package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/FarhanRizkiM/backend-ai/helper"
	"os"
)

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Mauls",
	AppName:       "Mauls ai",
}
var IPport, netstring = helper.GetAddress()

var PrivateKey = os.Getenv("PRIVATEKEY")
var PublicKey = os.Getenv("PUBLICKEY")