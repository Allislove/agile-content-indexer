package content

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Content indexer
func ContentIndexer(message string) string {
	r := gin.Default()
	r.GET("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": message,
		})
	})
	r.Run() // serve on 0.0.0.0:8080
	return message
}

// Pushing data to zin
func PushData(message string) string {
	f, err := os.Open("enron_mail_20110402")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	file, err := os.Open("../enron_mail_20110402")
	fmt.Println(file, "********FILE")
	// data := "../enron_mail_20110402"

	req, err := http.NewRequest("POST", "http://localhost:4080/api/clients/_multi", io.MultiReader(f))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#900")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
	return message
}
