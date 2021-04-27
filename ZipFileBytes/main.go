package main
import (
	"archive/zip"
  "bytes"
  "io/ioutil"
  "log"
)
func converter(inBuf []byte) ([]byte,error) {
 
  buf := new(bytes.Buffer)

  // Create a new zip archive.
  zipWriter := zip.NewWriter(buf)
 
	zipFile, err := zipWriter.Create("image.png")

	  if err != nil {
    return nil,err
	  }
 
	  _, err = zipFile.Write(inBuf)

	  if err != nil {
    return nil,err
	  }

    err = zipWriter.Close()

    if err != nil {
     return nil,err
	  }

	  return buf.Bytes(),nil
  }
 
func main() {

	fileBytes, err := ioutil.ReadFile("image.png")

	if err != nil {
    log.Fatal(err)
	}

	// passing fileBytes([]bytes) to zip function
	result,err := converter(fileBytes)
 
  if err != nil {
     log.Fatal(err)
      }
    
err = ioutil.WriteFile("sample.zip", result, 0777)

if err != nil {
   log.Fatal(err)
    }

}