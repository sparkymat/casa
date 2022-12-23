package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"git.orion.home/oxhead/casa/config/configiface"
	"git.orion.home/oxhead/casa/database/dbiface"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreateApp(cfg configiface.ConfigAPI, db dbiface.DatabaseAPI) echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.FormValue("title")
		description := c.FormValue("description")
		url := c.FormValue("url")

		if title == "" || url == "" {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Title and url needed for adding app")
		}

		imageFile, err := c.FormFile("image_file")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "File upload failed")
		}

		fileUID := ulid.Make().String()
		fileExt := filepath.Ext(imageFile.Filename)

		destinationFolderPath := path.Join(cfg.DataPath(), "uploads")
		//nolint:gomnd
		if err = os.MkdirAll(destinationFolderPath, os.FileMode(0o750)); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "File saving failed")
		}

		destinationFilePath := path.Join(destinationFolderPath, fmt.Sprintf("%s%s", fileUID, fileExt))

		src, err := imageFile.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "File saving failed")
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(destinationFilePath)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "File saving failed")
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "File saving failed")
		}

		_, err = db.CreateCatalogItem(c.Request().Context(), title, url, description, fmt.Sprintf("/uploads/%v%v", fileUID, fileExt))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create catalog item")
		}

		//nolint:wrapcheck
		return c.Redirect(http.StatusSeeOther, "/apps")
	}
}
