package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/MantasSilanskas/Tomita-EMOS-XML-Parcer/internal/structs"
)

func DownloadFile(folder, fileName, url string) error {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	filePath := filepath.Join("products", folder, fileName)

	//Create a empty file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func CreateInformation(f structs.ShopItemOrdered, folder, fileName string) error {
	filePath := filepath.Join("products", folder, fileName)

	//Create a empty file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	file.WriteString(fmt.Sprintf("Product: %s \n", f.Product))
	file.WriteString(fmt.Sprintf("ProductNo: %s \n", f.ProductNo))
	file.WriteString(fmt.Sprintf("ProductNo2: %s \n", f.ProductNo2))
	file.WriteString(fmt.Sprintf("EAN: %s \n", f.Ean))
	file.WriteString(fmt.Sprintf("Unit: %s \n", f.Unit))
	file.WriteString(fmt.Sprintf("Pack: %s \n", f.Pack))
	file.WriteString(fmt.Sprintf("Weight: %s \n", f.Weight))
	file.WriteString(fmt.Sprintf("Manufacturer: %s \n", f.Manufacturer))
	file.WriteString(fmt.Sprintf("Price VAT: %f \n", f.PriceVat))
	file.WriteString(fmt.Sprintf("Price Eshop VAT: %f \n", f.PriceEShopVat))
	file.WriteString(fmt.Sprintf("Price VOC: %f \n", f.PriceVoc))
	file.WriteString(fmt.Sprintf("Price Eshop pack VAT: %f \n", f.PriceEShopPackVat))
	file.WriteString(fmt.Sprintf("Price pack VAT: %f \n", f.PricePackVat))
	file.WriteString(fmt.Sprintf("VAT: %f \n", f.Vat))
	file.WriteString(fmt.Sprintf("DUES: %f \n", f.Dues))
	file.WriteString(fmt.Sprintf("Discount group: %s \n", f.DiscountGroup))
	file.WriteString(fmt.Sprintf("Category: %s \n", f.Category))
	file.WriteString(fmt.Sprintf("Category ID: %s \n", f.CategoryID))
	file.WriteString(fmt.Sprintf("ETIM Trida: %s \n", f.EtimTrida))
	file.WriteString(fmt.Sprintf("Description: %s \n", f.Description))
	file.WriteString(fmt.Sprintf("Description basic: %s \n", f.DescriptionBasic))

	for _, v := range f.Related {
		file.WriteString(fmt.Sprintf("Related: %s \n", v))
	}

	file.WriteString(fmt.Sprintf("Action: %s \n", f.Action))
	file.WriteString(fmt.Sprintf("Novelty: %s \n", f.Novelty))
	file.WriteString(fmt.Sprintf("Sale: %s \n", f.Sale))

	for i, v := range f.Params {
		file.WriteString(fmt.Sprintf("Param %d: \n", i+1))
		file.WriteString(fmt.Sprintf("Param name: %s \n", v.ParamName))
		file.WriteString(fmt.Sprintf("Order: %s \n", v.Order))
		file.WriteString(fmt.Sprintf("Val: %s \n", v.Val))
	}

	file.WriteString(fmt.Sprintf("URL: %s \n", f.URL))

	return nil
}

func CreateLogsFile(fileName string) (*os.File, error) {
	filePath := filepath.Join("products", fileName)

	return os.Create(filePath)
}
