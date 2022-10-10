package service

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/cemayan/earthquake_collector_mini/internal/types/earthquake"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"net/http"
)

type XMLService interface {
	Fetch(url string) []byte
	Parse(data []byte) []earthquake.Earthquake
}

type XMLSvc struct {
}

func (X XMLSvc) Fetch(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("There was an error downloading the file!")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Errorf("Read body: %v", err)
	}
	return data
}

func (X XMLSvc) Parse(data []byte) []earthquake.Earthquake {

	eqList := &earthquake.Eqlist{}

	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&eqList)

	if err != nil {
		fmt.Errorf("there was an error decoding the type!")
	}

	return eqList.Eqlist
}

func NewXMLService() XMLService {
	return &XMLSvc{}
}
