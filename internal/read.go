package internal

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadXMLFile(fileName string) (Shop, error) {
	// Open our xmlFile
	xmlFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Printf("Failed to open %s", fileName)
		return Shop{}, err
	}

	fmt.Printf("Succesfully opened %s \n", fileName)

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Failed to read xml file")
		return Shop{}, err
	}

	// we initialize our Shop
	var results Shop

	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'results' which we defined above
	err = xml.Unmarshal(byteValue, &results)
	if err != nil {
		fmt.Println(err)
	}

	return results, nil
}
