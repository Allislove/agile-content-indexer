package content

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

	DATA := `{
		"nombre": "Andres",
		"lastName": "Romaña"
	}`

	type Element struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	/*data := `Message-ID: = "<>"
	Date: = ""
	From: = ""
	To: = ""
	Subject: = ""
	Version: = ""
	Content-Type: = ""
	Content-Transfer-Encoding: = ""
	X-From: = ""
	X-To: = ""
	X-cc: = ""
	X-bcc: = ""
	X-Folder: = ""
	X-Origin: = ""
	X-FileName: = ""
	`

	SliceDATA := make([]Element, 0, 4)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		keyVal := strings.Split(line, "=")
		SliceDATA = append(SliceDATA, Element{Name: keyVal[0], Value: keyVal[1]})
		// SliceDATA[keyVal[0]] = keyVal[1] // creamos k:value pair json
	}
	byData, err := json.Marshal(SliceDATA)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", byData) */

	// req, err := http.NewRequest("POST", "http://localhost:4080/api/clients/_doc", bytes.NewReader([]byte(byData)))
	req, err := http.NewRequest("POST", "http://localhost:4080/api/clients/_doc", strings.NewReader(DATA))

	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
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
