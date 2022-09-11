package getPowerStatistics

import (
	"fmt"

	"github.com/jpillora/GoSungrow/iSolarCloud/api"
	"github.com/jpillora/GoSungrow/iSolarCloud/api/apiReflect"
)

const Url = "/v1/powerStationService/getPowerStatistics"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("ps_id: Can be fetched from getPsList")
	return ret
}

type ResultData struct {
	PRVlaue        string        `json:"PRVlaue"`
	City           interface{}   `json:"city"`
	DayPower       api.UnitValue `json:"dayPower"`
	DesignCapacity api.UnitValue `json:"design_capacity"`
	EqVlaue        string        `json:"eqVlaue"`
	NowCapacity    api.UnitValue `json:"nowCapacity"`
	PsName         string        `json:"ps_name"`
	PsShortName    string        `json:"ps_short_name"`
	Status1        string        `json:"status1"`
	Status2        string        `json:"status2"`
	Status3        string        `json:"status3"`
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
