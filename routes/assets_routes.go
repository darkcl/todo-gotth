package routes

import (
	"fmt"
	"net/http"
	"os"
	"todo-gotth/assets"

	"github.com/labstack/echo/v4"
)

func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		fmt.Println("using live mode")
		return http.FS(os.DirFS("./assets"))
	}

	fmt.Println("using embed mode")
	return http.FS(assets.Assets)
}

func SetupAssetsRoutes(e *echo.Echo) {
	var isDevelopment = os.Getenv("GO_ENV") != "production"

	assetHandler := http.FileServer(getFileSystem(isDevelopment))
	e.GET("/assets/*", echo.WrapHandler(http.StripPrefix("/assets/", assetHandler)))
}
