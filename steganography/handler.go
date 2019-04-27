package function

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"log"
	"strings"

	"github.com/auyer/steganography"
)

// requestData structure holds the Input Data.
/*
Message string : the message that will be enbedded in the image
Image string : Base64 encoded image
*/
type requestData struct {
	Message string `json:"message"`
	Image   string `json:"image"`
	Encode  bool   `json:"encode"`
}

// Handle a serverless request
func Handle(req []byte) string {
	// Reading the body
	data := requestData{}
	err := json.Unmarshal(req, &data)
	// If the body is not correct
	if err != nil {
		log.Println(fmt.Sprintf("error: bad body. %s ", err.Error()))
		return fmt.Sprintf(`{"error": "bad body. %s"}`, err.Error())
	}

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data.Image))
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Println(fmt.Sprintf("error: failed decoding image from base64. %s", err.Error()))
		return fmt.Sprintf(`{"error": "failed decoding image from base64. %s"}`, err.Error())
	}

	if data.Encode {

		buff := new(bytes.Buffer)

		err = steganography.Encode(buff, img, []byte(data.Message))

		if err != nil {
			log.Println(fmt.Sprintf("error: failed encoding message to image. %s ", err.Error()))
			return fmt.Sprintf(`{"error": "failed encoding message to image. %s"}`, err.Error())
		}
		return fmt.Sprintf(`{"image": "%s"}`, base64.StdEncoding.EncodeToString(buff.Bytes()))
	}
	// DECODING MODE
	message := steganography.Decode(steganography.GetMessageSizeFromImage(img), img)
	return fmt.Sprintf(`{"message": "%s"}`, string(message))

}
