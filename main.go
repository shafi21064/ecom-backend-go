package main

import "github.com/shafi21064/ecom-go/cmd"

func main() {
	cmd.Serve()

	// jwt, err := util.CreateJwt(
	// 	"secret-one",
	// 	util.Payload{
	// 		Sub:         45,
	// 		Name:        "shafi",
	// 		Email:       "shafi@email.com",
	// 		IsShopOwner: false,
	// 	},
	// )

	// if err != nil {
	// 	print(err)
	// 	return
	// }
	// println(jwt)
}
