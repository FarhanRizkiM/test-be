package testing

import (
	"fmt"
	"testing"

	// "github.com/FarhanRizkiM/test-be/controller"
	// "github.com/FarhanRizkiM/test-be/model"
	"github.com/whatsauth/watoken"
)

// func TestCreateNewUserRole(t *testing.T) {
// 	var userdata model.User
// 	userdata.Username = "farhanrizki"
// 	userdata.Password = "test12345"
// 	userdata.PasswordHash = "test12345"
// 	userdata.Email = "rizki@gmail.com"
// 	mconn := controller.SetConnection("MONGOSTRING", "testSeterahBot")
// 	controller.CreateNewUserRole(mconn, "user", userdata)
// }

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("test12345", privateKey)
	fmt.Println(hasil, err)
}
