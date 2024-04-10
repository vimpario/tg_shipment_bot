package api_tracker_service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTrackInfo(urlApi string) []byte {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, urlApi, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("X-Api-Key", "WK-wuo-3PK6gJNMLx_5PXhll0dtBAaWX")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Resp success")

	return body
}
