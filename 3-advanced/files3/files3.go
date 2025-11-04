package main
// Demo for files and directories
// error in file.close can be improved
import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	const relPathName = "test/test"
	const fileName = "test3.txt"
	const pathFile = "./" + relPathName + "/" + fileName
	fmt.Println(pathFile)

	makeDir(relPathName)
	createFile(relPathName, pathFile)
	writeFile(pathFile)
	readFile1(pathFile)
	appendFile(pathFile)
	readFile2(pathFile)
	deleteFile(pathFile)
	removePath(relPathName)
}

func FileClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func makeDir(pathName string) {
	// detect if pathdir exists
	if _, err := os.Stat(pathName); os.IsNotExist(err) {
		//mkdir pathName
		errMkdirAll	 := os.MkdirAll(pathName, os.ModePerm)
		if errMkdirAll != nil {
			log.Println("MkdirAll:", pathName, err)
		}
		fmt.Println("==> done creating path", pathName)
	} else {
		fmt.Println("Path" , pathName, "already exists")
	}
}

func createFile(pathName string, pathFileName string) {
	makeDir(pathName)
	var _, err = os.Stat(pathFileName)

	// create file if not exists
	if os.IsNotExist(err) {
		// create file
		var file, err = os.Create(pathFileName)
		if isError(err) {
			log.Println(err)
		}
		fmt.Println("==> done creating file", pathFileName)
		if file.Close() != nil {
			log.Println(err)
		}
	} else {
		fmt.Println("==>", pathFileName, "already exists!")
	}
}

func writeFile(pathFile string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(pathFile, os.O_RDWR, 0644)
	if isError(err) {
		log.Println(err)
	}
	defer FileClose(file)

	// write some text line-by-line to file
	_, err = file.WriteString("Hallo!\n")
	if isError(err) {
		log.Println(err)
	}
	_, err = file.WriteString("FRA-UAS SOA\n")
	if isError(err) {
		log.Println(err)
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		log.Println(err)
	}

	fmt.Println("==> done writing to file")
}

func readFile1(pathFile string) {
	// re-open file
	var file, err = os.OpenFile(pathFile, os.O_RDWR, 0644)
	if isError(err) {
		log.Println(err)
	}
	defer FileClose(file)

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Println(err)
	}

	// read file, line by line
	var inputBuf = make([]byte, stat.Size())
	for {
		_, err = file.Read(inputBuf)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(inputBuf))
}

func readFile2(pathFile string) {
	// read from a file easyly
	inputBuf, err := ioutil.ReadFile(pathFile)
	if err != nil {
		return
	}
	fmt.Println("==> done reading from file")
	fileContentString := string(inputBuf)
	fmt.Println(fileContentString)
}

func deleteFile(pathFile string) {
	// delete file
	var err = os.Remove(pathFile)
	if isError(err) {
		log.Println(err)
	}

	fmt.Println("==> done deleting file")
}

func removePath(path string) {
	// delete Pathfile
	if path != "" && path != "." {
		var err = os.RemoveAll(path)
		if isError(err) {
			log.Println(err)
		}
	}

	fmt.Println("==> done deleting path")
}

func appendFile(pathFile string) {
	//append file
	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer FileClose(file)

	if _, err := file.WriteString("Frankfurt am Main (was appended)\n"); err != nil {
		log.Println(err)
	}

	fmt.Println("==> done appending file")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return err != nil
}
