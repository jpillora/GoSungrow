package energyTrend

import (
	"fmt"

	"github.com/jpillora/GoSungrow/iSolarCloud/api/apiReflect"
)

const Url = "/v1/powerStationService/energyTrend"
const Disabled = false

type RequestData struct {
	// DeviceType string `json:"device_type" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	Echartunit string `json:"echartunit"`
	EndTime    string `json:"endTime"`
	EnergyMap  struct {
		ValStr string `json:"valStr"`
	} `json:"energyMap"`
	Energyunit string `json:"energyunit"`
	PowerMap   struct {
		Dates  []interface{} `json:"dates"`
		Units  string        `json:"units"`
		ValStr string        `json:"valStr"`
	} `json:"powerMap"`
}

func (e *ResultData) IsValid() error {
	var err error
	//switch {
	//case e.Dummy == "":
	//	break
	//default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	//}
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}
