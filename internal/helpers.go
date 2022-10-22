package internal

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (i *ShopItem) ToShopItemOrdered() ShopItemOrdered {
	return ShopItemOrdered{
		Product:           i.Product,
		ProductNo:         i.ProductNo,
		ProductNo2:        i.ProductNo2,
		Ean:               i.Ean,
		Unit:              i.Unit,
		Pack:              i.Pack,
		Weight:            i.Weight,
		Manufacturer:      i.Manufacturer,
		PriceVat:          i.PriceVat,
		PriceEShopVat:     i.PriceEShopVat,
		PriceVoc:          i.PriceVoc,
		PriceEShopPackVat: i.PriceEShopPackVat,
		PricePackVat:      i.PricePackVat,
		Price:             i.Price,
		PricePack:         i.PricePack,
		Vat:               i.Vat,
		Dues:              i.Dues,
		DiscountGroup:     i.DiscountGroup,
		Category:          i.Category,
		CategoryID:        i.CategoryID,
		EtimTrida:         i.EtimTrida,
		Description:       i.Description,
		DescriptionBasic:  i.DescriptionBasic,
		Related:           i.Related,
		Action:            i.Action,
		Novelty:           i.Novelty,
		Sale:              i.Sale,
		Params:            i.Params,
		URL:               i.URL,

		ImagesURLs: combineURLs(i),

		AttachManual:     i.AttachManual,
		AttachConformity: i.AttachConformity,
	}
}

func combineURLs(i *ShopItem) []string {
	list := []string{}

	if i.IMGURL != "" {
		list = append(list, i.IMGURL)
	}

	if i.IMGURLAlternative1 != "" {
		list = append(list, i.IMGURLAlternative1)
	}

	if i.IMGURLAlternative2 != "" {
		list = append(list, i.IMGURLAlternative2)
	}

	if i.IMGURLAlternative2 != "" {
		list = append(list, i.IMGURLAlternative2)
	}

	if i.IMGURLAlternative3 != "" {
		list = append(list, i.IMGURLAlternative3)
	}

	if i.IMGURLAlternative4 != "" {
		list = append(list, i.IMGURLAlternative4)
	}

	if i.IMGURLAlternative5 != "" {
		list = append(list, i.IMGURLAlternative5)
	}

	if i.IMGURLAlternative6 != "" {
		list = append(list, i.IMGURLAlternative6)
	}

	if i.IMGURLAlternative7 != "" {
		list = append(list, i.IMGURLAlternative7)
	}

	if i.IMGURLAlternative8 != "" {
		list = append(list, i.IMGURLAlternative8)
	}

	if i.IMGURLAlternative9 != "" {
		list = append(list, i.IMGURLAlternative9)
	}

	if i.IMGURLAlternative10 != "" {
		list = append(list, i.IMGURLAlternative10)
	}

	if i.IMGURLAlternative11 != "" {
		list = append(list, i.IMGURLAlternative11)
	}

	if i.IMGURLAlternative12 != "" {
		list = append(list, i.IMGURLAlternative12)
	}

	if i.IMGURLAlternative13 != "" {
		list = append(list, i.IMGURLAlternative13)
	}

	if i.IMGURLAlternative14 != "" {
		list = append(list, i.IMGURLAlternative14)
	}

	if i.IMGURLAlternative15 != "" {
		list = append(list, i.IMGURLAlternative15)
	}

	if i.IMGURLAlternative16 != "" {
		list = append(list, i.IMGURLAlternative16)
	}

	if i.IMGURLAlternative17 != "" {
		list = append(list, i.IMGURLAlternative17)
	}

	if i.IMGURLAlternative18 != "" {
		list = append(list, i.IMGURLAlternative18)
	}

	if i.IMGURLAlternative19 != "" {
		list = append(list, i.IMGURLAlternative19)
	}

	if i.IMGURLAlternative20 != "" {
		list = append(list, i.IMGURLAlternative20)
	}

	if i.IMGURLPack1 != "" {
		list = append(list, i.IMGURLPack1)
	}

	if i.IMGURLPack2 != "" {
		list = append(list, i.IMGURLPack2)
	}

	if i.IMGURLPack3 != "" {
		list = append(list, i.IMGURLPack3)
	}

	if i.IMGURLPack4 != "" {
		list = append(list, i.IMGURLPack4)
	}

	if i.IMGURLPack5 != "" {
		list = append(list, i.IMGURLPack5)
	}

	if i.IMGURLPack6 != "" {
		list = append(list, i.IMGURLPack6)
	}

	if i.IMGURLPack7 != "" {
		list = append(list, i.IMGURLPack7)
	}

	if i.IMGURLPack8 != "" {
		list = append(list, i.IMGURLPack8)
	}

	if i.IMGURLPack9 != "" {
		list = append(list, i.IMGURLPack9)
	}

	if i.IMGURLPack10 != "" {
		list = append(list, i.IMGURLPack10)
	}

	return list
}

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

func DownloadAdditionalInfo(f ShopItemOrdered) error {
	if f.AttachManual != "" {
		err := DownloadFile(f.ProductNo2, fmt.Sprintf("%s_%s.pdf", f.ProductNo2,
			"manual"), f.AttachManual)
		if err != nil {
			return err
		}
	}

	if f.AttachConformity != "" {
		err := DownloadFile(f.ProductNo2, fmt.Sprintf("%s_%s.pdf", f.ProductNo2,
			"conformity"), f.AttachConformity)
		if err != nil {
			return err
		}
	}

	return nil
}
