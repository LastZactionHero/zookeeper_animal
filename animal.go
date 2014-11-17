package main

import(
  "fmt"
  "os"
  "io/ioutil"
  "bytes"
  "mime/multipart"
  "net/http"
  "os/exec"
)

const url string = "http://localhost:8080/post_photo"
const filepath string = "./zookeeper.jpg"

func body() (*bytes.Buffer, string) {
  file, _ := os.Open(filepath)
  fileContents, _ := ioutil.ReadAll(file)
  fi, _ := file.Stat()
  file.Close()

  body := new(bytes.Buffer)
  writer := multipart.NewWriter(body)
  part, _ := writer.CreateFormFile("file", fi.Name())

  part.Write(fileContents)
  writer.Close()

  return body, writer.FormDataContentType()
}

func postPhoto() bool {
  body, contentType := body()
  request, _ := http.NewRequest("POST", url, body)
  request.Header.Add("Content-Type", contentType)

  client := &http.Client{}
  response, err := client.Do(request)
  if err != nil {
    return false
  }

  return response.StatusCode == 200
}

func main() {
  cmd := exec.Command("ls")
  out, err := cmd.Output()

  if err != nil {

  } else {
    fmt.Println(string(out))
  }
  
  if postPhoto() {
    fmt.Println("Photo Post: success")
  } else {
    fmt.Println("Photo Post: failure")
  }
}