package xml

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/MantasSilanskas/Tomita-EMOS-XML-Parcer/internal/structs"
)

func ReadXMLFile(fileName string) (structs.Shop, error) {
	// Open our xmlFile
	xmlFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Printf("Failed to open %s", fileName)
		return structs.Shop{}, err
	}

	fmt.Printf("Successfully opened %s \n", fileName)
	fmt.Println("------------------------------------------------------------")

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Failed to read xml file")
		return structs.Shop{}, err
	}

	// we initialize our Shop
	var results structs.Shop

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'results' which we defined above
	err = xml.Unmarshal(byteValue, &results)
	if err != nil {
		fmt.Println(err)
	}

	return results, nil
}
