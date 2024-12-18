package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func AppStart() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening/creating log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	engine := gin.Default()

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	port := 7100
	apiURL := fmt.Sprintf("http://localhost:%d/test", port)

	go func() {
		time.Sleep(1 * time.Second)

		fmt.Println("Opening browser to:", apiURL)
		if err := openBrowser(apiURL); err != nil {
			fmt.Printf("Failed to open browser: %v\n", err)
		}
	}()

	err = engine.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return exec.Command(cmd, args...).Start()
}
