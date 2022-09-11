package getTemplateList

import (
	"fmt"

	"github.com/jpillora/GoSungrow/Only"
	"github.com/jpillora/GoSungrow/iSolarCloud/api"
	"github.com/jpillora/GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/jpillora/GoSungrow/iSolarCloud/api/output"
)

const Url = "/v1/devService/getTemplateList"
const Disabled = false

type RequestData struct {
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	PageList []struct {
		TemplateID   int64  `json:"template_id"`
		TemplateName string `json:"template_name"`
		UpdateTime   string `json:"update_time"`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
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

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table

	for range Only.Once {
		table = output.NewTable()
		e.Error = table.SetTitle("")
		if e.Error != nil {
			break
		}

		e.Error = table.SetHeader(
			"Template Id",
			"Template Name",
			"Update On",
		)
		if e.Error != nil {
			break
		}

		for _, p := range e.Response.ResultData.PageList {
			_ = table.AddRow(
				p.TemplateID,
				p.TemplateName,
				api.NewDateTime(p.UpdateTime).PrintFull(),
			)
			if table.Error != nil {
				continue
			}
		}

		// table.InitGraph(output.GraphRequest {
		// 	Title:        "",
		// 	TimeColumn:   output.SetInteger(1),
		// 	SearchColumn: output.SetInteger(2),
		// 	NameColumn:   output.SetInteger(3),
		// 	ValueColumn:  output.SetInteger(4),
		// 	UnitsColumn:  output.SetInteger(5),
		// 	SearchString: output.SetString(""),
		// 	MinLeftAxis:  output.SetFloat(0),
		// 	MaxLeftAxis:  output.SetFloat(0),
		// })
	}

	return table
}
