package function

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

var testBase64Image = `iVBORw0KGgoAAAANSUhEUgAAAAgAAAAICAYAAADED76LAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAFUlEQVQY02P8DwQMeAATAwEwPBQAABtuBAy91jkOAAAAAElFTkSuQmCC`

// TestEncodeDecode tests the default usage of the linker.
func TestEncodeDecode(t *testing.T) {
	message := "hello stego"
	body := []byte(fmt.Sprintf(`
	{	
		"message" : "%s" ,
		"image" : "%s",
		"encode" : true
	}`, message, testBase64Image))

	encodedImageData := Handle(body)

	data := requestData{}
	json.Unmarshal([]byte(encodedImageData), &data)

	body = []byte(fmt.Sprintf(`
	{	
		"image" : "%s",
		"encode" : false
	}`, data.Image))

	response := Handle(body)

	json.Unmarshal([]byte(response), &data)

	if data.Message != message {
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
	fmt.Println(response)
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
	fmt.Println(response)
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
	fmt.Println(response)
	if response != `{"error": "bad body. unexpected end of JSON input"}` {
		log.Println("error not caught: malformed json")
		t.FailNow()
	}
}
