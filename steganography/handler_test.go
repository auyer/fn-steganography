package function

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

var testBase64Image = `iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAYAAADED76LAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAFUlEQVQY02P8DwQMeAATAwEwPBQAABtuBAy91jkOAAAAAElFTkSuQmCC`

var testImageURL = "https://github.com/auyer/steganography/raw/master/examples/stegosaurus.png"

// TestEncodeDecodeBase64 tests the default usage of the .
func TestEncodeDecodeBase64(t *testing.T) {
	message := "hello stego"
	body := []byte(fmt.Sprintf(`
	{	
		"message" : "%s" ,
		"image" : "%s",
		"encode" : true
	}`, message, testBase64Image))

	encodedImageData := Handle(body)
	encodedImageData = strings.TrimPrefix(encodedImageData, "data:image/png;base64,")

	body = []byte(fmt.Sprintf(`
	{	
		"image" : "%s",
		"encode" : false
	}`, encodedImageData))

	response := Handle(body)

	if string(response) != message {
		log.Println("messages do no match")
		t.FailNow()
	}
}

func TestEncodeDecodeBase64WithMeta(t *testing.T) {
	message := "hello stego"
	body := []byte(fmt.Sprintf(`
	{	
		"message" : "%s" ,
		"image" : "%s",
		"encode" : true
	}`, message, "data:image/png;base64,"+testBase64Image))

	encodedImageData := Handle(body)
	encodedImageData = strings.TrimPrefix(encodedImageData, "data:image/png;base64,")

	body = []byte(fmt.Sprintf(`
	{	
		"image" : "%s",
		"encode" : false
	}`, encodedImageData))

	response := Handle(body)

	if string(response) != message {
		log.Println("messages do no match")
		t.FailNow()
	}
}

func TestEncodeDecodeURL(t *testing.T) {
	message := "hello stego"
	body := []byte(fmt.Sprintf(`
	{	
		"message" : "%s" ,
		"image" : "%s",
		"encode" : true
	}`, message, testImageURL))

	encodedImageData := Handle(body)
	encodedImageData = strings.TrimPrefix(encodedImageData, "data:image/png;base64,")

	body = []byte(fmt.Sprintf(`
	{	
		"image" : "%s",
		"encode" : false
	}`, encodedImageData))

	response := Handle(body)

	if string(response) != message {
		log.Println("messages do no match")
		t.FailNow()
	}
}

func TestMissingImage(t *testing.T) {
	message := "hello steganography"
	body := []byte(fmt.Sprintf(`
	{	
		"message" : "%s" ,
		"encode" : true
	}`, message))

	response := Handle(body)
	// fmt.Println(response)
	if response != `{"error": "failed decoding image from base64. image: unknown format"}` {
		log.Println("error not caught: image missing")
		t.FailNow()
	}
}

func TestLargeMessage(t *testing.T) {
	message := "hello steganography"
	body := []byte(fmt.Sprintf(`
	{	
		"message" : "%s" ,
		"image" : "%s",
		"encode" : true
	}`, message, testBase64Image))

	response := Handle(body)
	// fmt.Println(response)
	if response != `{"error": "failed encoding message to image. message too large for image"}` {
		log.Println("error not caught: message too large")
		t.FailNow()
	}
}
func TestMalformedJson(t *testing.T) {
	body := []byte(`
	{	
		"message" : "fail" ,`)

	response := Handle(body)
	// fmt.Println(response)
	if response != `{"error": "bad body. unexpected end of JSON input"}` {
		log.Println("error not caught: malformed json")
		t.FailNow()
	}
}

func TestFailedHTTP(t *testing.T) {
	body := []byte(fmt.Sprintf(`
	{	
		"image" : "%s",
		"encode" : true
	}`, "http://localhost/image.png"))

	response := Handle(body)
	fmt.Println(response)
	if response != `{"error" : "Unable to download image file from URI: http://localhost/image.png"}` {
		log.Println("error not caught: failed http connection")
		t.FailNow()
	}
}

func TestBadImageField(t *testing.T) {
	body := []byte(fmt.Sprintf(`
	{	
		"image" : "%s",
		"encode" : true
	}`, "1/image.png"))

	response := Handle(body)
	fmt.Println(response)
	if response != `{"error": "failed to decode base64 image"}` {
		log.Println("error not caught: no match")
		t.FailNow()
	}
}

// func TestBadB64Field(t *testing.T) {
// 	body := []byte(fmt.Sprintf(`
// 	{
// 		"image" : "%s",
// 		"encode" : true
// 	}`, "5555"))

// 	response := Handle(body)
// 	fmt.Println(response)
// 	if response != `{"error": "image field didnt match a URL or base64 image"}` {
// 		log.Println("error not caught: malformed json")
// 		t.FailNow()
// 	}
// }
