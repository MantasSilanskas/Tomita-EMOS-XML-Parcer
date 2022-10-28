package internal

import "encoding/xml"

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

	IMGURLPack1  string `xml:"IMGURL_PACK_01"`
	IMGURLPack2  string `xml:"IMGURL_PACK_02"`
	IMGURLPack3  string `xml:"IMGURL_PACK_03"`
	IMGURLPack4  string `xml:"IMGURL_PACK_04"`
	IMGURLPack5  string `xml:"IMGURL_PACK_05"`
	IMGURLPack6  string `xml:"IMGURL_PACK_06"`
	IMGURLPack7  string `xml:"IMGURL_PACK_07"`
	IMGURLPack8  string `xml:"IMGURL_PACK_08"`
	IMGURLPack9  string `xml:"IMGURL_PACK_09"`
	IMGURLPack10 string `xml:"IMGURL_PACK_10"`

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
