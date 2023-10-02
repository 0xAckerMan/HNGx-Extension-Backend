package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const uploadDir = "./uploads/"

func healthcheck(c *fiber.Ctx) error {
    status := map[string]any {
        "status" : "Active",
        "Environment" : "Dev",
        "Version" : "1.0.0",
    }

    return c.Status(fiber.StatusOK).JSON(status)
}

func main(){
    app := fiber.New(fiber.Config{
    // Set the maximum request body size to a suitable value (e.g., 32MB).
    // Adjust this value based on your application's requirements.
    BodyLimit: 32 * 1024 * 1024, // 32MB
    })
    app.Use(logger.New())
    app.Get("/api/health", healthcheck)
    app.Post("/upload", uploadHandler)
    app.Get("/videos/:filename", videoHandler)
    app.Post("/uploader/:videoFilename", streamUploadHandler)
    app.Get("/videos", listVideosHandler)
    fmt.Printf("Server listening on port %s\n", "3000")
    log.Fatal(app.Listen(":80"))
}


func uploadHandler(c *fiber.Ctx) error {
    // Extract video filename from the request header
    videoFilename := c.Get("X-Video-Filename")
    if videoFilename == "" {
        return c.Status(http.StatusBadRequest).SendString("Missing video filename")
    }

    // Create the file on the server or append to an existing file
    videoPath := filepath.Join(uploadDir, videoFilename)
    file, err := os.OpenFile(videoPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Failed to create/open the file on the server")
    }
    defer file.Close()

    // Get the request body as []byte
    videoData := c.Request().Body()

    // Create an io.Reader from the []byte using bytes.NewReader
    videoDataReader := bytes.NewReader(videoData)

    // Copy the video data from the io.Reader to the file
    _, err = io.Copy(file, videoDataReader)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Failed to write video data to the file")
    }

    return c.Status(http.StatusCreated).SendString(fmt.Sprintf("Video data chunk received successfully for filename: %s", videoFilename))
}

type customReader struct {
    src    io.Reader
    buffer []byte
}

func (r *customReader) Read(p []byte) (n int, err error) {
    // Read from the underlying source and copy to the provided buffer.
    n, err = r.src.Read(r.buffer)
    if err != nil && err != io.EOF {
        return n, err
    }

    // Copy the read data to the destination slice.
    copy(p, r.buffer[:n])

    return n, err
}

func streamUploadHandler(c *fiber.Ctx) error {
    // Extract video filename from the URL params
    videoFilename := c.Params("videoFilename")
    if videoFilename == "" {
        return c.Status(http.StatusBadRequest).SendString("Missing video filename")
    }

    // Create the file on the server or append to an existing file
    videoPath := filepath.Join(uploadDir, videoFilename)
    file, err := os.OpenFile(videoPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Failed to create/open the file on the server")
    }
    defer file.Close()

    // Get the request body as []byte
    videoData := c.Request().Body()

    // Create an io.Reader from the []byte using bytes.NewReader
    videoDataReader := bytes.NewReader(videoData)

    // Copy the video data from the io.Reader to the file
    _, err = io.Copy(file, videoDataReader)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Failed to write video data to the file")
    }

    return c.Status(http.StatusCreated).SendString(fmt.Sprintf("Video data chunk received successfully for filename: %s", videoFilename))
}



func videoHandler(c *fiber.Ctx) error {
	// Extract the video filename from the URL params
	videoFilename := c.Params("filename")

	if videoFilename == "" {
		return c.Status(http.StatusBadRequest).SendString("Invalid video filename")
	}

	// Serve the video file
	videoPath := filepath.Join(uploadDir, videoFilename)
	return c.SendFile(videoPath)
}

func listVideosHandler(c *fiber.Ctx) error {
    // List all video files in the uploads directory
    videoFiles, err := ioutil.ReadDir(uploadDir)
    if err != nil {
        return c.Status(http.StatusInternalServerError).SendString("Failed to list videos")
    }

    // Extract video filenames and send them as a JSON response
    var videoList []string
    for _, fileInfo := range videoFiles {
        videoList = append(videoList, fileInfo.Name())
    }

    return c.JSON(videoList)
}

