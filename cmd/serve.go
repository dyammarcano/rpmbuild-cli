package cmd

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/static"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: serveFunc,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveFunc(cmd *cobra.Command, args []string) error {
	//rootPath, err := currentDirectory(args)
	//if err != nil {
	//	return err
	//}

	router := gin.Default()

	router.GET("/assets/*filepath", func(c *gin.Context) {
		path := c.Param("filepath")
		data, err := static.GetRsouce("assets" + path)
		if err != nil {
			c.String(http.StatusNotFound, "Resource file not found")
			return
		}

		c.Data(http.StatusOK, geMimeType(path), data)
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		data, err := static.GetRsouce("favicon.ico")
		if err != nil {
			c.String(http.StatusNotFound, "favicon file not found")
			return
		}

		c.Data(http.StatusOK, http.DetectContentType(data), data)
	})

	router.GET("/", func(c *gin.Context) {
		data, err := static.GetRsouce("index.html")
		if err != nil {
			c.String(http.StatusNotFound, "File not found")
			return
		}

		c.Data(http.StatusOK, http.DetectContentType(data), data)
	})

	// Define your API routes here
	router.GET("/api/data", func(c *gin.Context) {
		// Handle your API logic here
		c.JSON(http.StatusOK, gin.H{"message": "API endpoint hit"})
	})

	return router.Run(":8080")
}

func geMimeType(path string) string {
	switch {
	case strings.HasSuffix(path, ".css"):
		return "text/css; charset=utf-8"
	case strings.HasSuffix(path, ".html"):
		return "text/html; charset=utf-8"
	case strings.HasSuffix(path, ".js"):
		return "application/javascript; charset=utf-8"
	case strings.HasSuffix(path, ".png"):
		return "image/png"
	case strings.HasSuffix(path, ".svg"):
		return "image/svg+xml; charset=utf-8"
	default:
		return "text/plain; charset=utf-8"
	}
}
