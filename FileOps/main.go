package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	fileName, pathName string = "./sampleFiles/Sample.txt", "./sampleFiles"
	anotherFile        string = "Sample2.txt"
)

func main() {

	fmt.Println("File Operations: ", pathName)

	if _, err := os.Stat(pathName); errors.Is(err, os.ErrNotExist) {

		err := os.Mkdir("sampleFiles", 0755)

		if err != nil {
			log.Fatalln("Unable to create a Directory: ", err)
		}
	}

	createFile(fileName)

	writeFile(fileName)

	fileInfo(fileName)

	// readFile(fileName)

	// readTotalFileIntoMemory(fileName)

	// writeTotalFileFromMemory(fileName, []byte("Sample"))

	readAndWriteBytes(fileName, anotherFile)

	readLineByLine(fileName)

	deleteFile(fileName, "")
}

func createFile(fileName string) {
	file, err := os.Create(fileName)

	if err != nil {
		log.Fatalln("Unable to Create a new File", err)
	}
	defer file.Close()
}

func writeFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalln("Error while opening to write to a file: ", err)
	}

	defer file.Close()

	num, err := file.WriteString(`Lorem ipsum dolor sit amet consectetur adipisicing elit. 
	Debitis ipsa sunt assumenda delectus quas vero itaque cumque reiciendis similique 
	tenetur animi obcaecati voluptatibus sint culpa nostrum aliquam, vitae odit ipsam?`)

	if err != nil {
		log.Fatalln("Error while writing to a file: ", err)
	}
	fmt.Println("Number of Bytes written: ", num)
}

func fileInfo(fileName string) {
	file, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln("Error while getting info of the file: ", err)
	}
	fmt.Println("Name of the file: ", file.Name())
	fmt.Println("Size of the file in Bytes: ", file.Size())
	fmt.Println("Permission Mode of the file: ", file.Mode())
	fmt.Println("Last Modification Time of the file: ", file.ModTime())
	fmt.Println("Is it a directory: ", file.IsDir())
	fmt.Println(file.Sys())
}

func readAndWriteBytes(fileName, anotherFile string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)

	if err != nil {
		log.Fatalln("Error while opening the file: ", err)
	}

	defer file.Close()

	file2, err := os.OpenFile(anotherFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)

	if err != nil {
		log.Fatalln("Error while opening the 2 file: ", err)
	}

	defer file2.Close()

	buff := make([]byte, 128)

	// reader := bufio.NewReader(file)

	for {

		n, err := file.Read(buff)

		fmt.Println("N is : ", n)
		if err == io.EOF || n == 0 {
			break
		}

		if err != nil {
			log.Fatalln("Error while reading file in bytes: ", err)
		}

		// fmt.Println("Bytes data is: ", string(buff[0:n]))

		if _, err = file2.Write(buff[0:n]); err != nil {
			log.Fatalln("Error while writing bytes to a file: ", err)
		}

	}
}

func readLineByLine(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0755)

	if err != nil {
		log.Fatalln("Error while opening the file: ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// To read word by word
	// scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		lineText := scanner.Text()
		fmt.Println(lineText)
	}

	if err = scanner.Err(); err != nil {
		log.Fatalln("Error while reading file: ", err)
	}
}

func deleteFile(fileName, pathName string) {
	err := os.Remove(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	if pathName != "" {
		err = os.RemoveAll(pathName)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

// func readFile(fileName string) {
// 	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

// 	if err != nil {
// 		log.Fatalln("Error while opening the file", err)
// 	}

// 	defer file.Close()
// }

// func readTotalFileIntoMemory(fileName string) {

// 	fileBytes, err := ioutil.ReadFile(fileName)

// 	if err != nil {
// 		log.Fatalln("Error while reading the file", err)
// 	}

// 	fmt.Println(string(fileBytes))

// }

// func writeTotalFileFromMemory(fileName string, bytesData []byte) {

// 	err := ioutil.WriteFile(fileName, bytesData, 0644)

// 	if err != nil {
// 		log.Fatalln("Error while writing the file", err)
// 	}

// }
