package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// GetEnhanceUploadedFile retrieves an uploaded file from an HTTP request.
// It takes an HTTP request (r) and the name of the form field containing the file (formName) as input.
// The function returns a strings.Reader containing the file's content, the filename, the multipart.FileHeader,
// and an error if any occurs.

// Step-by-Step Process:
// 1. Initialize variables: fileStr (to store the file's content) and file (a strings.Reader to hold the file content).
// 2. Log the start of the function.
// 3. Attempt to retrieve the file data and header using r.FormFile(formName).
// 4. If an error occurs during retrieval, log the error and return an empty file, empty fileStr, the header, and the error.
// 5. If the file data is successfully retrieved, read its content into fileStr.
// 6. Create a strings.Reader (file) from fileStr to facilitate further use of the file's content.
// 7. Log the end of the function.
// 8. Return the file, fileStr, the multipart.FileHeader, and nil as the error.
func GetFileDetails(r *http.Request, formName string) (*strings.Reader, string, *multipart.FileHeader, error) {
	log.Println("GetFileDetails(+)")

	fileStr := ""
	var file *strings.Reader

	// Attempt to retrieve the file data and header using r.FormFile(formName)
	fileBody, header, lErr := r.FormFile(formName)

	if lErr != nil {
		// If an error occurs during retrieval, return an empty file, empty fileStr, the header, and the error
		return file, fileStr, header, fmt.Errorf("GetFileDetails:001" + lErr.Error())
	} else {
		// If the file data is successfully retrieved, read its content into fileStr
		datas, _ := io.ReadAll(fileBody)
		fileStr = string(datas)

		// Create a strings.Reader (file) from fileStr to facilitate further use of the file's content
		file = strings.NewReader(fileStr)

		log.Println("GetFileDetails(-)")
		// Log the end of the function
		return file, fileStr, header, nil
	}
}

// ReadCSV reads the contents of a CSV file from a strings.Reader and returns the data as a 2D slice of strings.

// Step 1: Initialize a 2D slice to store the CSV data
// Step 2: Create a CSV reader for the input string reader
// Step 3: Read the CSV file row by row
// Step 4: Check for the end of the file
// Step 5: Append each row to the 2D slice
// Step 6: Return the 2D slice containing the CSV data and no error
func ReadCSV(r *http.Request, pFile string, pDelimeter rune) ([][]string, error) {
	var lRecord [][]string
	lFile, _, _, lErr := GetFileDetails(r, pFile)
	if lErr != nil {
		return lRecord, fmt.Errorf("ReadCSV:001" + lErr.Error())
	} else {
		// Step 2: Create a CSV reader for the input string reader
		lRows := csv.NewReader(lFile)
		lRows.Comma = pDelimeter
		// Step 3: Read the CSV file row by row
		for {
			// Step 4: Read a row from the CSV
			lRecordRow, lErr := lRows.Read()

			// Step 4: Check for the end of the file
			if lErr == io.EOF {
				break // Exit the loop when we reach the end of the file
			} else {
				// Step 5: Append the read row to the 2D slice
				lRecord = append(lRecord, lRecordRow)
			}
		}
	}
	// Step 6: Return the 2D slice containing the CSV data and no error
	return lRecord, nil
}

// ReadText reads the contents of a text file from a strings.Reader and returns the data as a 2D slice of strings.

// Step 1: Initialize a 2D slice to store the text data
// Step 2: Create a CSV reader for the input string reader (assuming it's CSV-formatted text)
// Step 3: Read the text file row by row
// Step 4: Check for the end of the file
// Step 5: Append each row to the 2D slice
// Step 6: Return the 2D slice containing the text data and no error
func ReadText(r *http.Request, pFile string, pDelimeter rune) ([][]string, error) {
	var lRecord [][]string
	lFile, _, _, lErr := GetFileDetails(r, pFile)
	if lErr != nil {
		return lRecord, fmt.Errorf("ReadText:001" + lErr.Error())
	} else {
		// Step 2: Create a CSV reader for the input string reader (assuming it's CSV-formatted text)
		lRows := csv.NewReader(lFile)
		lRows.Comma = pDelimeter
		// Step 3: Read the text file row by row
		for {
			// Step 4: Read a row from the text
			lRecordRow, lErr := lRows.Read()

			// Step 4: Check for the end of the file
			if lErr == io.EOF {
				break // Exit the loop when we reach the end of the file
			} else {
				// Step 5: Append the read row to the 2D slice
				lRecord = append(lRecord, lRecordRow)
			}
		}
	}
	// Step 6: Return the 2D slice containing the text data and no error
	return lRecord, nil
}

// ReadXlsxFile reads an uploaded XLSX file from an HTTP request, extracts its contents,
// and performs specific operations on the data.
// It takes an HTTP request (r), the name of the uploaded file (pFile) as inputs.
// The function first retrieves the uploaded file and saves it to the server. It then opens
// the XLSX file, reads the data from a specified tab, and stores it in a 2D array (record).
// After processing the data, the function removes the temporary file created during the operation.
// Any encountered errors are returned as error messages.

// File Creation and Reading:(Step-by-Step Process)
// 1. The uploaded file is retrieved from the HTTP request.
// 2. It is saved to a server location with a unique file name based on the provided Header.Filename.
// 3. The content of the uploaded file is read, and a temporary file is created on the server.
// 4. The XLSX file is opened using excelize, and data is extracted from a specified tab (e.g., "TabName").
// 5. The data is stored in a 2D array (record) for further processing.

func ReadXlsxFile(r *http.Request, pFile string) error {
	log.Println("ReadXlsFile +")
	var record [][]string
	lFile, _, Header, lErr := GetFileDetails(r, pFile)
	if lErr != nil {
		return fmt.Errorf("ReadXlsFile:001" + lErr.Error())
	} else {
		path := "./"
		fileName := path + Header.Filename
		//Creating file's in specific server path
		out, lErr := os.Create(fileName)
		if lErr != nil {
			return fmt.Errorf("ReadXlsFile : 002" + lErr.Error())
		} else {
			datas, _ := io.ReadAll(lFile)
			fileStr := string(datas)

			file := strings.NewReader(fileStr)

			_, lErr = io.Copy(out, file) // Copy the file's content to the server file
			if lErr != nil {
				return fmt.Errorf("ReadXlsFile : 003" + lErr.Error())
			} else {
				lNewFile, lErr := excelize.OpenFile(fileName)
				if lErr != nil {
					return fmt.Errorf("ReadXlsFile : 004" + lErr.Error())
				} else {
					rows, lErr := lNewFile.GetRows("TabName") // Specify the tab name in the XLSX file
					if lErr != nil {
						return fmt.Errorf("ReadXlsFile : 005" + lErr.Error())
					} else {
						for _, row := range rows {
							record = append(record, row)
						}
					}

					// Your condition to filter records can be applied here.

					// This method is used to remove the temporary file created during the operation.
					lErr = os.Remove(fileName)
					if lErr != nil {
						return fmt.Errorf("ReadXlsFile : 005-" + lErr.Error())
					}
				}
			}
		}
	}
	log.Println("ReadXlsFile-")
	return nil
}
