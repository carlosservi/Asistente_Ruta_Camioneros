package route

import (
	"encoding/json"
	"io/ioutil"
)

func ReadJsonRestAreas(filePath string) ([]RestArea, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var restAreas []RestArea
	err = json.Unmarshal(data, &restAreas)
	if err != nil {
		return nil, err
	}
	return restAreas, nil
}

func ReadJsonRoutes(filePath string) ([]Route, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var routes []Route
	err = json.Unmarshal(data, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}
