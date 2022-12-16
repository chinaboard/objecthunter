package objecthunter

import (
	"encoding/xml"
	"fmt"
	"github.com/chinaboard/objecthunter/oss"
	"io"
	"net/http"
	"strings"
)

type ObjectHunter struct {
	endpoint string
}

func NewObjectHunter(endpoint string) *ObjectHunter {
	endpoint = strings.TrimSuffix(endpoint, "/") + "/"
	return &ObjectHunter{endpoint}
}

func (receiver *ObjectHunter) ListKeys(next string, filter *HunterFilter) []string {
	temp, err := receiver.do(next)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var result []string

	if filter.listHandle(*temp) {
		for _, object := range temp.Objects {
			if filter.objectHandle(object) {
				result = append(result, object.Key)
			}
		}
	}

	return result
}

func (receiver *ObjectHunter) do(next string) (*oss.ListObjectsResultV2, error) {
	url := receiver.endpoint + "?list-type=2&max-keys=1000"
	if next != "" {
		url += "&continuation-token=" + next
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	var out oss.ListObjectsResultV2
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
