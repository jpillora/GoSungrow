package api

import (
	"GoSungro/Only"
	"errors"
	"fmt"
)


type AreaStruct struct {
	ApiRoot   interface{}	// *web.Web
	Name      AreaName
	EndPoints TypeEndPoints
	Error     error
}


func (as AreaStruct) Exists(name EndPointName) error {
	var err error
	for range Only.Once {
		if _, ok := as.EndPoints[name]; !ok {
			err = errors.New("unknown endpoint")
			break
		}
	}

	return err
}

func (as *AreaStruct) SortEndPoints() []EndPointName {
	return as.EndPoints.SortEndPoints()
}

func (as *AreaStruct) GetEndPoint(name EndPointName) EndPoint {
	return as.EndPoints.GetEndPoint(name)
}

func (as AreaStruct) ListEndpoints() {
	for range Only.Once {
		fmt.Printf("Listing all endpoints from area '%s':\n", as.Name)
		as.EndPoints.ListEndpoints()

		// table := tablewriter.NewWriter(os.Stdout)
		// table.SetHeader([]string{"Area", "EndPoint", "Url"})
		// table.SetBorder(true)
		// for _, endpoint := range as.SortEndPoints() {
		// 	table.Append([]string{
		// 		string(as.Name),
		// 		string(as.EndPoints[endpoint].GetName()),
		// 		as.EndPoints[endpoint].GetUrl().String(),
		// 	})
		// }
		// table.Render()
	}
}

func (as AreaStruct) CountEndpoints() int {
	return as.EndPoints.CountEndpoints()
}
