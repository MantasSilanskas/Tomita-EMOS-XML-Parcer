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
		return structs.Shop{}, fmt.Errorf("failed to open %s file", fileName)
	}

	fmt.Printf("Successfully opened %s \n", fileName)
	fmt.Println("------------------------------------------------------------")

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return structs.Shop{}, fmt.Errorf("failed to read xml file")
	}

	// we initialize our Shop
	var results structs.Shop

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'results' which we defined above
	err = xml.Unmarshal(byteValue, &results)
	if err != nil {
		return structs.Shop{}, fmt.Errorf("failed to unmarshal data")
	}

	return results, nil
}
