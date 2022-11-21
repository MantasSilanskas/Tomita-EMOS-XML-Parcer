package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MantasSilanskas/Tomita-EMOS-XML-Parcer/internal/utils"
	"github.com/MantasSilanskas/Tomita-EMOS-XML-Parcer/internal/xml"
)

func main() {
	var (
		fileName             string
		successful, errCount int
		unsuccessful         []string
	)

	fmt.Print("Enter xml filename: ")
	fmt.Scanf("%s", &fileName)
	fmt.Println("------------------------------------------------------------")

	shop, err := xml.ReadXMLFile(fileName)
	if err != nil {
		fmt.Printf("failed to read xml file error: %s \n", err)
		return
	} else {
		fmt.Printf("Successfully extracted data from %s starting download of the pictures. \n", fileName)
	}

	fmt.Println("------------------------------------------------------------")

	if err := os.Mkdir("products", os.ModePerm); err != nil {
		if err.Error() == "mkdir products: file exists" {
			fmt.Println("Products dict exists already")
		}
	}

	logsFile, err := utils.CreateLogsFile(fmt.Sprintf("%s_logs.txt", time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		fmt.Println("failed to create logs file")
	}

	for _, item := range shop.Items {
		errCount = 0

		orderedFile := item.ToShopItemOrdered()

		if len(orderedFile.ImagesURLs) == 0 {
			continue
		}

		path := filepath.Join("products", item.ProductNo2)

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			if err.Error() == "mkdir products: file exists" {
				fmt.Printf("%s dict exists already \n", path)
			}
		}

		for i, g := range orderedFile.ImagesURLs {
			fileName = fmt.Sprintf("%s_%d.jpg", item.ProductNo2, i)

			if i == 0 {
				fileName = fmt.Sprintf("%s.jpg", item.ProductNo2)
			}

			err = utils.DownloadFile(item.ProductNo2, fileName, g)
			if err != nil {
				errCount++
			}
		}

		if item.AttachManual != "" {
			err = utils.DownloadFile(item.ProductNo2, fmt.Sprintf("%s_%s.pdf",
				item.ProductNo2, "manual"), item.AttachManual)
			if err != nil {
				errCount++
			}
		}

		if item.AttachConformity != "" {
			err = utils.DownloadFile(item.ProductNo2, fmt.Sprintf("%s_%s.pdf",
				item.ProductNo2, "conformity"), item.AttachConformity)
			if err != nil {
				errCount++
			}
		}

		err = utils.CreateInformation(orderedFile, item.ProductNo2, fmt.Sprintf("%s_info.txt",
			item.ProductNo2))
		if err != nil {
			fmt.Println("failed to to create information file")
			errCount++
		}

		if errCount == 0 {
			successful++

			logsFile.WriteString(fmt.Sprintf(
				"Succesfully downloaded pictures and information of %s \n", orderedFile.ProductNo2))
		} else {
			unsuccessful = append(unsuccessful, orderedFile.ProductNo2)
			fmt.Printf("Failed to download all pictures and information of %s product \n",
				orderedFile.ProductNo2)

			logsFile.WriteString(fmt.Sprintf(
				"Failed to download pictures and information of %s \n", orderedFile.ProductNo2))
		}

		if successful%10 == 0 {
			fmt.Printf("Pictures and information of %d has been downloaded succesfully. \n", successful)
		}
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("List of products unsuccessful of which data failed to download data", unsuccessful)
}
