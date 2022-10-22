package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Open our xmlFile
	xmlFile, err := os.Open("productGeneralFeed.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened productGeneralFeed.xml")

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var results Shop
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'results' which we defined above
	err = xml.Unmarshal(byteValue, &results)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range results.Items {
		item := v.toShopItemOrdered()

		if err := os.Mkdir(item.ProductNo2, os.ModePerm); err != nil {
			fmt.Println(err)
		}

		if len(item.ImagesURLs) == 0 {
			continue
		}

		for i, g := range item.ImagesURLs {
			downloadFile(v.ProductNo2, fmt.Sprintf("%s_%d.jpg", v.ProductNo2, i), g)
		}

	}
}

func downloadFile(folder, fileName, url string) error {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	filePath := filepath.Join(folder, fileName)

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

//

type Shop struct {
	XMLName xml.Name   `xml:"SHOP"`
	Items   []ShopItem `xml:"SHOPITEM"`
}

type ShopItem struct {
	Product           string   `xml:"PRODUCT"`
	ProductNo         string   `xml:"PRODUCTNO"`
	ProductNo2        string   `xml:"PRODUCTNO2"`
	Ean               string   `xml:"EAN"`
	Unit              string   `xml:"UNIT"`
	Pack              string   `xml:"PACK"`
	Weight            string   `xml:"WEIGHT"`
	Manufacturer      string   `xml:"MANUFACTURER"`
	PriceVat          float64  `xml:"PRICE_VAT"`
	PriceEShopVat     float64  `xml:"PRICE_ESHOP_VAT"`
	PriceVoc          float64  `xml:"PRICE_VOC"`
	PriceEShopPackVat float64  `xml:"PRICE_ESHOP_PACK_VAT"`
	PricePackVat      float64  `xml:"PRICE_PACK_VAT"`
	Price             float64  `xml:"PRICE"`
	PricePack         float64  `xml:"PRICE_PACK"`
	Vat               float64  `xml:"VAT"`
	Dues              float64  `xml:"DUES"`
	DiscountGroup     string   `xml:"DISCOUNT_GROUP"`
	Category          string   `xml:"CATEGORY"`
	CategoryID        string   `xml:"CATEGORY_ID"`
	EtimTrida         string   `xml:"ETIM_TRIDA"`
	Description       string   `xml:"DESCRIPTION"`
	DescriptionBasic  string   `xml:"DESCRIPTION_BASIC"`
	Related           []string `xml:"RELATED"`
	Action            string   `xml:"ACTION"`
	Novelty           string   `xml:"NOVELTY"`
	Sale              string   `xml:"SALE"`
	Params            []Param  `xml:"PARAMS"`
	URL               string   `xml:"URL"`

	IMGURL string `xml:"IMGURL"`

	IMGURLAlternative1  string `xml:"IMGURL_ALTERNATIVE_01"`
	IMGURLAlternative2  string `xml:"IMGURL_ALTERNATIVE_02"`
	IMGURLAlternative3  string `xml:"IMGURL_ALTERNATIVE_03"`
	IMGURLAlternative4  string `xml:"IMGURL_ALTERNATIVE_04"`
	IMGURLAlternative5  string `xml:"IMGURL_ALTERNATIVE_05"`
	IMGURLAlternative6  string `xml:"IMGURL_ALTERNATIVE_06"`
	IMGURLAlternative7  string `xml:"IMGURL_ALTERNATIVE_07"`
	IMGURLAlternative8  string `xml:"IMGURL_ALTERNATIVE_08"`
	IMGURLAlternative9  string `xml:"IMGURL_ALTERNATIVE_09"`
	IMGURLAlternative10 string `xml:"IMGURL_ALTERNATIVE_10"`
	IMGURLAlternative11 string `xml:"IMGURL_ALTERNATIVE_11"`
	IMGURLAlternative12 string `xml:"IMGURL_ALTERNATIVE_12"`
	IMGURLAlternative13 string `xml:"IMGURL_ALTERNATIVE_13"`
	IMGURLAlternative14 string `xml:"IMGURL_ALTERNATIVE_14"`
	IMGURLAlternative15 string `xml:"IMGURL_ALTERNATIVE_15"`
	IMGURLAlternative16 string `xml:"IMGURL_ALTERNATIVE_16"`
	IMGURLAlternative17 string `xml:"IMGURL_ALTERNATIVE_17"`
	IMGURLAlternative18 string `xml:"IMGURL_ALTERNATIVE_18"`
	IMGURLAlternative19 string `xml:"IMGURL_ALTERNATIVE_19"`
	IMGURLAlternative20 string `xml:"IMGURL_ALTERNATIVE_20"`

	IMGURLPack1 string `xml:"IMGURL_PACK_01"`
	IMGURLPack2 string `xml:"IMGURL_PACK_02"`

	AttachManual     string `xml:"ATTACH_MANUAL"`
	AttachConformity string `xml:"ATTACH_CONFORMITY"`
}

type Param struct {
	ParamName string `xml:"PARAM_NAME"`
	Val       string `xml:"VAL"`
	Order     string `xml:"ORDER"`
}

type ShopItemOrdered struct {
	Product           string
	ProductNo         string
	ProductNo2        string
	Ean               string
	Unit              string
	Pack              string
	Weight            string
	Manufacturer      string
	PriceVat          float64
	PriceEShopVat     float64
	PriceVoc          float64
	PriceEShopPackVat float64
	PricePackVat      float64
	Price             float64
	PricePack         float64
	Vat               float64
	Dues              float64
	DiscountGroup     string
	Category          string
	CategoryID        string
	EtimTrida         string
	Description       string
	DescriptionBasic  string
	Related           []string
	Action            string
	Novelty           string
	Sale              string
	Params            []Param
	URL               string

	ImagesURLs []string

	AttachManual     string
	AttachConformity string
}

func (i *ShopItem) toShopItemOrdered() ShopItemOrdered {
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

	return list
}
