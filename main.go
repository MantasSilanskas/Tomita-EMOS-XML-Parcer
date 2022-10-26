package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MantasSilanskas/Tomita-EMOS-XML-Parcer/internal"
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

	shop, err := internal.ReadXMLFile(fileName)
	if err != nil {
		fmt.Printf("failed to read xml file")
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

	for _, item := range shop.Items {
		errCount = 0

		orderedFile := item.ToShopItemOrdered()

		if len(orderedFile.ImagesURLs) == 0 {
			continue
		}

		path := filepath.Join("products", item.ProductNo2)

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			if err.Error() == "mkdir products: file exists" {
				fmt.Println("Products dict exists already")
			}
		}

		for i, g := range orderedFile.ImagesURLs {
			fileName = fmt.Sprintf("%s_%d.jpg", item.ProductNo2, i)

			if i == 0 {
				fileName = fmt.Sprintf("%s.jpg", item.ProductNo2)
			}

			err = internal.DownloadFile(item.ProductNo2, fileName, g)
			if err != nil {
				errCount++
			}
		}

		if item.AttachManual != "" {
			err = internal.DownloadAttachManual(item.ProductNo2, fmt.Sprintf("%s_%s.pdf",
				item.ProductNo2, "manual"), item.AttachManual)
			if err != nil {
				errCount++
			}
		}

		if item.AttachConformity != "" {
			err = internal.DownloadAttachConformity(item.ProductNo2, fmt.Sprintf("%s_%s.pdf",
				item.ProductNo2, "conformity"), item.AttachConformity)
			if err != nil {
				errCount++
			}
		}

		err = internal.CreateInformation(orderedFile, item.ProductNo2, fmt.Sprintf("%s_info.txt",
			item.ProductNo2))
		if err != nil {
			fmt.Println("failed to to create information file")
			errCount++
		}

		if errCount == 0 {
			successful++
		} else {
			unsuccessful = append(unsuccessful, orderedFile.ProductNo2)
			fmt.Printf("Failed to download all pictures and information of %s product \n",
				orderedFile.ProductNo2)
		}

		if successful%10 == 0 {
			fmt.Printf("Pictures and information of %d has been downloaded succesfully. \n", successful)
		}

		// fmt.Printf("Finished download all pictures and information of %s product \n",
		// 	orderedFile.ProductNo2)
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("List of products unsuccessful of which data failed to download data", unsuccessful)
}
