package function

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"

	// image/x packages allow decoding format x
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/auyer/steganography"
)

var urlRegEx, _ = regexp.Compile(`\b((http|https):\/\/?)[^\s()<>]+(?:\([\w\d]+\)|([^[:punct:]\s]|\/?))`)

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

func getImage(inputURL string) (string, error) {
	res, err := http.Get(inputURL)
	if err != nil {
		return fmt.Sprintf(`{"error" : "Unable to download image file from URI: %s"}`, inputURL), err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf(`{"error" : "Unable to read response body: %s"}`, err), err
	}
	return string(data), nil
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
	data.Image = strings.TrimSpace(data.Image)
	if urlRegEx.Match([]byte(data.Image)) {
		data.Image, err = getImage(data.Image)
		if err != nil {
			return data.Image
		}
		return encodeDecode(data)
	}
	data.Image = strings.TrimPrefix(data.Image, "data:image/png;base64,")
	data.Image = strings.TrimPrefix(data.Image, "data:image/jpeg;base64,")
	data.Image = strings.TrimPrefix(data.Image, "data:image/jpg;base64,")
	decodedImg, err := base64.StdEncoding.DecodeString(data.Image)
	if err != nil {
		return `{"error": "failed to decode base64 image"}`
	}
	data.Image = string(decodedImg)
	return encodeDecode(data)
}

func encodeDecode(data requestData) string {

	img, _, err := image.Decode(strings.NewReader(data.Image))
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
		return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buff.Bytes())
	}
	// DECODING MODE
	message := steganography.Decode(steganography.GetMessageSizeFromImage(img), img)
	return string(message)
}
