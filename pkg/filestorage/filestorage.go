package filestorage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const (
	host         = "http://localhost"
	path         = "files"
	apiversion   = "v1"
	port         = 80
	maxfilesizeM = 50
)

// Get a common url and path for http request.
func getRequestURI() string {
	return fmt.Sprintf("%s:%d/%s/%s", host, port, apiversion, path)
}

func request(method, url string, body io.Reader, headers map[string]string) (resBody []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, fmt.Errorf("could not create request: %s", err)
	}

	for k, val := range headers {
		req.Header.Set(k, val)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error making http request: %s", err)
	}

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, fmt.Errorf("could not read response body: %s", err)
	}

	return resBody, res.StatusCode, nil
}

// Make a GET request and prints all the file names.
func GetFileList() (fileList []string, err error) {
	resBody, _, err := request(http.MethodGet, getRequestURI(), nil, nil)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resBody, &fileList)
	if err != nil {
		return nil, fmt.Errorf("couldn't fetch a file list: %s", err)
	}

	return fileList, err
}

// Make a DELETE request for deleting a file from the server by its name.
func DeleteFile(fileName string) {
	requestURL := fmt.Sprintf("%s/%s", getRequestURI(), fileName)
	_, statusCode, err := request(http.MethodDelete, requestURL, nil, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch statusCode {
	case http.StatusBadRequest:
		fmt.Printf("Unknown file name. Please check the request and try again")
	case http.StatusNotFound:
		fmt.Printf("A file %s does not exist. Please check the file name and try again.\n", fileName)
	case http.StatusNoContent:
		fmt.Printf("file %s deleted successfully\n", fileName)
	default:
		fmt.Printf("Unknown status of operation: %d\n", statusCode)
	}
}

// Make a POST request for uploading a file to the server.
func UploadFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Couldn't read a file %s\n", filePath)
		return
	}

	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		fmt.Println("Error occured while uploading a file")
		return
	}

	io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		fmt.Println("Error occured while uploading a file")
		return
	}

	headers := map[string]string{}
	headers["Content-Type"] = writer.FormDataContentType()

	_, statusCode, err := request(http.MethodPost, getRequestURI(), body, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch statusCode {
	case http.StatusBadRequest:
		fmt.Println("Bad request. Please check the file and try again")
	case http.StatusUnprocessableEntity:
		fmt.Println("This type of file doesn't allowed.")
	case http.StatusForbidden:
		fmt.Println("File with the same name already exists in the storage. You can rename the file and try to upload again.")
	case http.StatusRequestEntityTooLarge:
		fmt.Printf("File is too large. Allowed file size is %dM.\n", maxfilesizeM)
	case http.StatusCreated:
		fmt.Println("file uploaded successfully.")
	default:
		fmt.Printf("Unknown status of operation: %d.\n", statusCode)
	}
}
