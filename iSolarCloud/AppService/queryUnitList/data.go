package queryUnitList

import (
	"fmt"

	"github.com/jpillora/GoSungrow/iSolarCloud/api/apiReflect"
)

const Url = "/v1/userService/queryUnitList"
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

type ResultData []struct {
	IsBasicUnit  int64  `json:"is_basic_unit"`
	TargetUnit   string `json:"target_unit"`
	UnitConverID int64  `json:"unit_conver_id"`
	UnitName     string `json:"unit_name"`
	UnitType     int64  `json:"unit_type"`
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
