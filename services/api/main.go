package main

import (
	"MockExamLab/internal/cmd"
)

// @title           					Swagger Example API
// @version         					1.0
// @description     					This is a sample server celler server.
// @termsOfService  					http://swagger.io/terms/
// @contact.name   						API Support
// @contact.url    						http://www.swagger.io/support
// @contact.email  						support@swagger.io
// @license.name  						Apache 2.0
// @license.url   						http://www.apache.org/licenses/LICENSE-2.0.html
// @host      							mel-api.go7.ir
// @schemes 							http https
// @BasePath  							/api/v1
// @securityDefinitions.apikey 			ApiKeyAuth
// @in 									header
// @name 								Authorization
func main() {
	//localhost:8080
	//mel-api.go7.ir
	cmd.Execute()
}
