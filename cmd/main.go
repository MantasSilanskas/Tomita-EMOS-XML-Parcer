package main

import (
	"fmt"
	"os"
	"path/filepath"
	"tomitaXmlParser/internal"
)

func main() {
	var fileName string

	fmt.Print("Enter xml filename: ")
	fmt.Scanf("%s", &fileName)

	shop, err := internal.ReadXMLFile(fileName)
	if err != nil {
		fmt.Printf("failed to read xml file")
		return
	}

	if err := os.Mkdir("products", os.ModePerm); err != nil {
		if err.Error() == "mkdir products: file exists" {
			fmt.Println("Products dict exists already")
		}
	}

	for _, item := range shop.Items {
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
			err = internal.DownloadFile(item.ProductNo2, fmt.Sprintf("%s_%d.jpg",
				item.ProductNo2, i), g)
			if err != nil {
				fmt.Println("failed to download: ", item.ProductNo2,
					fmt.Sprintf("%s_%d.jpg", item.ProductNo2, i), g)
			}
		}

		err := internal.DownloadAdditionalInfo(orderedFile)
		if err != nil {
			fmt.Println("failed to download additional information files")
		}

		err = internal.CreateInformation(orderedFile, item.ProductNo2, fmt.Sprintf("%s_info.txt",
			item.ProductNo2))
		if err != nil {
			fmt.Println("failed to to create information file")
		}

		fmt.Printf("Finished download all pictures and information of %s product \n",
			orderedFile.ProductNo2)
	}
}
