package queryDeviceList

import (
	"fmt"
	"strconv"

	"github.com/jpillora/GoSungrow/Only"
	"github.com/jpillora/GoSungrow/iSolarCloud/api"
	"github.com/jpillora/GoSungrow/iSolarCloud/api/apiReflect"
	"github.com/jpillora/GoSungrow/iSolarCloud/api/output"
)

const Url = "/v1/devService/queryDeviceList"
const Disabled = false

type RequestData struct {
	PsId string `json:"ps_id" required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}

type ResultData struct {
	DevCountByStatusMap struct {
		FaultCount   int64 `json:"fault_count" PointId:"fault_count" PointType:""`
		OfflineCount int64 `json:"offline_count" PointId:"offline_count" PointType:""`
		RunCount     int64 `json:"run_count" PointId:"run_count" PointType:""`
		WarningCount int64 `json:"warning_count" PointId:"warning_count" PointType:""`
	} `json:"dev_count_by_status_map"`
	DevCountByTypeMap struct {
		One4 int64 `json:"14"`
		Two2 int64 `json:"22"`
	} `json:"dev_count_by_type_map"`
	DevTypeDefinition struct {
		One    string `json:"1"`
		One0   string `json:"10"`
		One1   string `json:"11"`
		One2   string `json:"12"`
		One3   string `json:"13"`
		One4   string `json:"14"`
		One5   string `json:"15"`
		One6   string `json:"16"`
		One7   string `json:"17"`
		One8   string `json:"18"`
		One9   string `json:"19"`
		Two0   string `json:"20"`
		Two1   string `json:"21"`
		Two2   string `json:"22"`
		Two3   string `json:"23"`
		Two4   string `json:"24"`
		Two5   string `json:"25"`
		Two6   string `json:"26"`
		Two8   string `json:"28"`
		Two9   string `json:"29"`
		Three  string `json:"3"`
		Three0 string `json:"30"`
		Three1 string `json:"31"`
		Three2 string `json:"32"`
		Three3 string `json:"33"`
		Three4 string `json:"34"`
		Three5 string `json:"35"`
		Three6 string `json:"36"`
		Three7 string `json:"37"`
		Three8 string `json:"38"`
		Three9 string `json:"39"`
		Four   string `json:"4"`
		Four0  string `json:"40"`
		Four1  string `json:"41"`
		Four2  string `json:"42"`
		Four3  string `json:"43"`
		Four4  string `json:"44"`
		Four5  string `json:"45"`
		Four6  string `json:"46"`
		Four7  string `json:"47"`
		Four8  string `json:"48"`
		Five   string `json:"5"`
		Five0  string `json:"50"`
		Six    string `json:"6"`
		Seven  string `json:"7"`
		Eight  string `json:"8"`
		Nine   string `json:"9"`
		Nine9  string `json:"99"`
	} `json:"dev_type_definition"`
	PageList []struct {
		AlarmCount              int64       `json:"alarm_count" PointId:"alarm_count" PointType:""`
		ChnnlID                 int64       `json:"chnnl_id" PointId:"channel_id" PointType:""`
		CommandStatus           int64       `json:"command_status" PointId:"command_status" PointType:""`
		ComponentAmount         int64       `json:"component_amount" PointId:"component_amount" PointType:""`
		DataFlag                int64       `json:"data_flag" PointId:"data_flag" PointType:""`
		DataFlagDetail          int64       `json:"data_flag_detail" PointId:"data_flag_detail" PointType:""`
		DeviceArea              string      `json:"device_area" PointId:"device_area" PointType:""`
		DeviceAreaName          string      `json:"device_area_name" PointId:"device_area_name" PointType:""`
		DeviceCode              int64       `json:"device_code" PointId:"device_code" PointType:""`
		DeviceID                int64       `json:"device_id" PointId:"device_id" PointType:""`
		DeviceModelCode         string      `json:"device_model_code" PointId:"device_model_code" PointType:""`
		DeviceModelID           string      `json:"device_model_id" PointId:"device_model_id" PointType:""`
		DeviceName              string      `json:"device_name" PointId:"device_name" PointType:""`
		DeviceStatus            int64       `json:"device_status" PointId:"device_status" PointType:""`
		DeviceType              int64       `json:"device_type" PointId:"device_type" PointType:""`
		FaultCount              int64       `json:"fault_count" PointId:"fault_count" PointType:""`
		FaultStatus             string      `json:"fault_status" PointId:"fault_status" PointType:""`
		FunctionEnum            string      `json:"function_enum" PointId:"function_enum" PointType:""`
		InstallerAlarmCount     int64       `json:"installer_alarm_count" PointId:"installer_alarm_count" PointType:""`
		InstallerDevFaultStatus int64       `json:"installer_dev_fault_status" PointId:"installer_dev_fault_status" PointType:""`
		InstallerFaultCount     int64       `json:"installer_fault_count" PointId:"installer_fault_count" PointType:""`
		InverterModelType       int64       `json:"inverter_model_type" PointId:"inverter_model_type" PointType:""`
		IsDeveloper             string      `json:"is_developer" PointId:"is_developer" PointType:""`
		IsG2point5Module        int64       `json:"is_g2point5_module" PointId:"is_g2point5_module" PointType:""`
		IsInit                  int64       `json:"is_init" PointId:"is_init" PointType:""`
		IsSecond                int64       `json:"is_second" PointId:"is_second" PointType:""`
		IsSupportParamset       int64       `json:"is_support_paramset" PointId:"is_support_paramset" PointType:""`
		NodeTimestamps          interface{} `json:"node_timestamps" PointId:"node_timestamps" PointType:""`
		OwnerAlarmCount         int64       `json:"owner_alarm_count" PointId:"owner_alarm_count" PointType:""`
		OwnerDevFaultStatus     int64       `json:"owner_dev_fault_status" PointId:"owner_dev_fault_status" PointType:""`
		OwnerFaultCount         int64       `json:"owner_fault_count" PointId:"owner_fault_count" PointType:""`
		PointData               PointData   `json:"point_data"`
		Points                  interface{} `json:"points" PointId:"points" PointType:""`
		PsTimezoneInfo          struct {
			IsDst    string `json:"is_dst"`
			TimeZone string `json:"time_zone"`
		} `json:"psTimezoneInfo"`
		PsID          int64       `json:"ps_id" PointId:"ps_id" PointType:""`
		PsKey         string      `json:"ps_key" PointId:"ps_key" PointType:""`
		RelState      int64       `json:"rel_state" PointId:"rel_state" PointType:""`
		Sn            string      `json:"sn" PointId:"sn" PointType:""`
		StringAmount  int64       `json:"string_amount" PointId:"string_amount" PointType:""`
		TypeName      string      `json:"type_name" PointId:"type_name" PointType:""`
		UnitName      interface{} `json:"unit_name" PointId:"unit_name" PointType:""`
		UUID          string      `json:"uuid" PointId:"uuid" PointType:""`
		UUIDIndexCode string      `json:"uuid_index_code" PointId:"uuid_index_code" PointType:""`
	} `json:"pageList"`
	RowCount int64 `json:"rowCount"`
}

type PointData []PointStruct
type PointStruct struct {
	CodeID                 int64  `json:"code_id"`
	CodeIDOrderID          string `json:"code_id_order_id"`
	CodeName               string `json:"code_name"`
	DevPointLastUpdateTime string `json:"dev_point_last_update_time"`
	IsPlatformDefaultUnit  int64  `json:"is_platform_default_unit"`
	IsShow                 int64  `json:"is_show"`
	OrderID                int64  `json:"order_id"`
	OrderNum               int64  `json:"order_num"`
	PointGroupID           int64  `json:"point_group_id"`
	PointGroupIDOrderID    string `json:"point_group_id_order_id"`
	PointGroupName         string `json:"point_group_name"`
	PointID                int64  `json:"point_id"`
	PointName              string `json:"point_name"`
	PointSign              string `json:"point_sign"`
	Relate                 int64  `json:"relate"`
	TimeStamp              string `json:"time_stamp"`
	Unit                   string `json:"unit"`
	ValIsFixd              string `json:"val_is_fixd"`
	ValidSize              int64  `json:"valid_size"`
	Value                  string `json:"value"`
	ValueDescription       string `json:"value_description"`
}

// type VirtualPointStruct struct {
// 	api.DataEntry
// 	ValueFloat float64
// }

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	//	break
	// default:
	//	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

// type DecodeResultData ResultData
//
// func (e *ResultData) UnmarshalJSON(data []byte) error {
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
// }

func (e *ResultData) GetDataByName(name string) PointData {
	var ret PointData
	for range Only.Once {
		i := len(e.PageList)
		if i == 0 {
			break
		}
		for _, p := range e.PageList {
			if p.DeviceName != name {
				continue
			}
			ret = p.PointData
			break
		}
	}
	return ret
}

func (e *EndPoint) GetDataTable() output.Table {
	var table output.Table
	for range Only.Once {
		table = output.NewTable()
		e.Error = table.SetTitle("")
		if e.Error != nil {
			break
		}

		_ = table.SetHeader(
			"Date",
			"PointStruct Id",
			"PointStruct Group Name",
			"Description",
			"Value",
			"Unit",
		)

		for _, d := range e.Response.ResultData.PageList {
			for _, p := range d.PointData {
				p.Value, p.Unit = api.DivideByThousandIfRequired(p.Value, p.Unit)
				// gp := api.GetPointInt("", p.PointID)
				// if gp != nil {
				// 	_ = table.AddRow(
				// 		api.NewDateTime(p.TimeStamp).PrintFull(),
				// 		api.NameDevicePointInt(d.PsKey, p.PointID),
				// 		p.PointGroupName,
				// 		p.PointName,
				// 		gp.Description,
				// 		p.Value,
				// 		p.Unit,
				// 		gp.Unit,
				// 	)
				// 	continue
				// }

				_ = table.AddRow(
					api.NewDateTime(p.TimeStamp).PrintFull(),
					api.NameDevicePointInt(d.PsKey, p.PointID),
					p.PointGroupName,
					p.PointName,
					p.Value,
					p.Unit,
				)
			}
		}

		table.InitGraph(output.GraphRequest{
			Title:        "",
			TimeColumn:   output.SetInteger(1),
			SearchColumn: output.SetInteger(2),
			NameColumn:   output.SetInteger(4),
			ValueColumn:  output.SetInteger(5),
			UnitsColumn:  output.SetInteger(6),
			SearchString: output.SetString(""),
			MinLeftAxis:  output.SetFloat(0),
			MaxLeftAxis:  output.SetFloat(0),
		})

	}
	return table
}

type DataEntry api.DataEntry
type EntryMap api.DataMap

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		// // Used for virtual entries.
		// // 0 - sungrow_battery_charging_power
		// var PVPowerToBattery VirtualPointStruct
		//
		// // sensor.sungrow_battery_discharging_power
		// var BatteryPowerToLoad VirtualPointStruct
		//
		// // 0 - sensor.sungrow_total_export_active_power
		// var PVPowerToGrid VirtualPointStruct
		//
		// // sensor.sungrow_purchased_power
		// var GridPowerToLoad VirtualPointStruct
		//
		// // 0 - sensor.sungrow_daily_battery_charging_energy_from_pv
		// var YieldBatteryCharge VirtualPointStruct
		// // var DailyBatteryChargingEnergy VirtualPointStruct
		//
		// // sensor.sungrow_daily_battery_discharging_energy
		// var DailyBatteryDischargingEnergy VirtualPointStruct
		//
		// // 0 - sensor.sungrow_daily_feed_in_energy_pv
		// var YieldFeedIn VirtualPointStruct
		//
		// // sensor.sungrow_daily_purchased_energy
		// var DailyPurchasedEnergy VirtualPointStruct
		//
		// var PVPower VirtualPointStruct
		//
		// var LoadPower VirtualPointStruct
		//
		// var YieldSelfConsumption VirtualPointStruct
		// // var DailyFeedInEnergy VirtualPointStruct
		// var TotalPvYield VirtualPointStruct
		//
		// var DailyTotalLoad VirtualPointStruct
		//
		// var TotalEnergyConsumption VirtualPointStruct

		entries.StructToPoints("", e.Response.ResultData.DevCountByStatusMap)
		for _, d := range e.Response.ResultData.PageList {
			entries.StructToPoints("", d)
		}

		for _, d := range e.Response.ResultData.PageList {
			for _, p := range d.PointData {
				pid := api.SetPoint(strconv.FormatInt(p.PointID, 10))
				uv := api.CreateUnitValue(p.Value, p.Unit)
				entries.AddUnitValue("", d.PsKey, pid, p.PointName, api.NewDateTime(p.TimeStamp), uv)

				// vt := api.GetPointInt(d.PsKey, p.PointID)
				// if !vt.Valid {
				// 	vt = &api.Point {
				// 		PsKey: d.PsKey,
				// 		Id:    api.PointToName(strconv.FormatInt(p.PointID, 10)),
				// 		Name:  p.PointName,
				// 		Unit:  uv.Unit,
				// 		Type:  "PointTypeInstant",
				// 	}
				// }
				//
				// entries.Add(pid, api.DataEntry {
				// 	Date:       api.NewDateTime(p.TimeStamp),
				// 	Id:         api.NameDevicePointInt(d.PsKey, p.PointID),
				// 	GroupName:  p.PointGroupName,
				// 	Name:       p.PointName,
				// 	Value:      uv.Value,
				// 	ValueFloat: uv.ValueFloat,
				// 	Unit:       uv.Unit,
				// 	Point:      vt,
				// 	Index:      0,
				// })

				// Handle virtual results.
				// switch pid {
				// 	case "13126":
				// 		// BatteryChargingPower
				// 		entries["PVPowerToBattery"] = entries[pid]
				// 	case "13150":
				// 		// BatteryDischargingPower
				// 		entries["BatteryPowerToLoad"] = entries[pid]
				// 	case "13121":
				// 		// TotalExportActivePower
				// 		entries["PVPowerToGrid"] = entries[pid]
				// 	case "13149":
				// 		// PurchasedPower
				// 		entries["GridPowerToLoad"] = entries[pid]
				// 	case "13003":
				// 		// TotalDcPower
				// 		entries["PVPower"] = addVirtualAlias(entries[pid], "pv_power", "PV Power")
				// 	case "13119":
				// 		// TotalLoadActivePower
				// 		entries["LoadPower"] = addVirtualAlias(entries[pid], "load_power", "Load Power")
				//
				// 		// addVirtualAlias(entries[pid], "FOO", "FOO")
				//
				// 	case "13112":
				// 		// Daily PV Yield
				// 		entries["DailyPvEnergy"] = addVirtualAlias(entries["DailyPvEnergy"], "daily_pv_energy", "Daily PV Energy")
				// 	case "13174":
				// 		// DailyBatteryChargingEnergyFromPv
				// 		entries["YieldBatteryCharge"] = addVirtualAlias(entries[pid], "pv_battery_charge", "PV Battery Charge")
				// 	case "13029":
				// 		// DailyBatteryDischargingEnergy
				// 		entries["DailyBatteryDischargingEnergy"] = entries[pid]
				// 	case "13122":
				// 		// entries["DailyFeedInEnergy"] = addVirtualAlias(entries[pid], "pv_feed_in", "PV Feed In")
				// 		// @TODO - This may differ from DailyFeedInEnergyPv
				// 	case "13173":
				// 		// DailyFeedInEnergyPv
				// 		entries["YieldFeedIn"] = addVirtualAlias(entries[pid], "pv_feed_in", "PV Feed In")
				// 	case "13147":
				// 		// DailyPurchasedEnergy
				// 		entries["DailyPurchasedEnergy"] = addVirtualAlias(entries[pid], "daily_purchased_energy", "Daily Purchased Energy")
				//
				// 	case "13116":
				// 		// DailyLoadEnergyConsumptionFromPv
				// 		entries["YieldSelfConsumption"] = addVirtualAlias(entries[pid], "pv_self_consumption", "PV Self Consumption")
				// 	case "13134":
				// 		// TotalPvYield
				// 		entries["TotalPvYield"] = addVirtualAlias(entries[pid], "pv_total_yield", "PV Total Yield")
				//
				// 	case "13199":
				// 		// Daily Load Energy Consumption
				// 		entries["DailyTotalLoad"] = addVirtualAlias(entries[pid], "daily_total_energy", "Daily Total Energy")
				//
				// 	case "13130":
				// 		// Total Load Energy Consumption
				// 		entries["TotalEnergyConsumption"] = addVirtualAlias(entries[pid], "total_energy_consumption", "Total Energy Consumption"
				// }
			}
		}

		if len(entries.Entries) == 0 {
			break
		}

		// TotalDcPower
		entries.FromRefAddAlias("p13003", api.VirtualPsId, "pv_power", "")
		// BatteryChargingPower
		entries.FromRefAddAlias("p13126", api.VirtualPsId, "pv_power_to_battery", "")
		// BatteryDischargingPower
		entries.FromRefAddAlias("p13150", api.VirtualPsId, "battery_power_to_load", "")
		// TotalExportActivePower
		entries.FromRefAddAlias("p13121", api.VirtualPsId, "pv_power_to_grid", "")

		// TotalLoadActivePower
		entries.FromRefAddAlias("p13119", api.VirtualPsId, "load_power", "")

		// PurchasedPower
		entries.FromRefAddAlias("p13149", api.VirtualPsId, "grid_power_to_load", "")

		// Daily PV Yield
		entries.FromRefAddAlias("p13112", api.VirtualPsId, "daily_pv_energy", "")
		// DailyPvEnergy := entries.getFloatValue("DailyTotalLoad") - entries.getFloatValue("DailyPurchasedEnergy")
		// DailyBatteryChargingEnergyFromPv
		entries.FromRefAddAlias("p13174", api.VirtualPsId, "pv_battery_charge", "")
		// DailyBatteryDischargingEnergy
		entries.FromRefAddAlias("p13029", api.VirtualPsId, "battery_discharge", "")

		// @TODO - This may differ from DailyFeedInEnergyPv
		// entries["DailyFeedInEnergy"] = entries.AddVirtualAliasFromRef("13122", "pv_feed_in", "PV Feed In")

		// DailyFeedInEnergyPv
		entries.FromRefAddAlias("p13173", api.VirtualPsId, "pv_feed_in", "")
		// DailyPurchasedEnergy
		entries.FromRefAddAlias("p13147", api.VirtualPsId, "daily_purchased_energy", "")
		// DailyLoadEnergyConsumptionFromPv
		entries.FromRefAddAlias("p13116", api.VirtualPsId, "pv_self_consumption", "")
		// TotalPvYield
		entries.FromRefAddAlias("p13134", api.VirtualPsId, "pv_total_yield", "")
		// Daily Load Energy Consumption
		entries.FromRefAddAlias("p13199", api.VirtualPsId, "daily_total_energy", "")
		// Total Load Energy Consumption
		entries.FromRefAddAlias("p13130", api.VirtualPsId, "total_energy_consumption", "")
		// entries.AddPointFromRef(api.Point{ Id:"p13130" }, api.Point{ PsKey:api.VirtualPsId, Id:"total_energy_consumption" })

		// entries.CopyEntry("p13130").CreateAlias()
		// 		entries.GetEntry(api.Point{PsKey:psId, Id:"total_income", Unit:p.TotalIncome.Unit, Type:api.PointTypeTotal}, now, p.TotalIncome.Value)

		/*
			PVPower				- TotalDcPower
			PVPowerToBattery	- BatteryChargingPower
			PVPowerToLoad		- TotalDcPower - BatteryChargingPower - TotalExportActivePower
			PVPowerToGrid		- TotalExportActivePower

			LoadPower			- TotalLoadActivePower
			BatteryPowerToLoad	- BatteryDischargingPower
			BatteryPowerToGrid	- ?

			GridPower			- lowerUpper(PVPowerToGrid, GridPowerToLoad)
			GridPowerToLoad		- PurchasedPower
			GridPowerToBattery	- ?

			YieldSelfConsumption	- DailyLoadEnergyConsumptionFromPv
			YieldBatteryCharge		- DailyBatteryChargingEnergyFromPv
			YieldFeedIn				- DailyFeedInEnergyPv
		*/

		// Add virtual entries.
		// ts := ret.Entries[0].Date
		// var value float64

		entries.FromRefAddFloat("pv_self_consumption",
			api.VirtualPsId, "pv_daily_yield", "",
			entries.GetFloatValue("pv_self_consumption")+entries.GetFloatValue("pv_battery_charge")+entries.GetFloatValue("pv_feed_in"))

		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId, "pv_self_consumption_percent", "",
			entries.GetPercent("pv_self_consumption", "daily_pv_energy"))
		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId, "pv_battery_charge_percent", "",
			entries.GetPercent("pv_battery_charge", "daily_pv_energy"))
		entries.FromRefAddFloat("daily_pv_energy",
			api.VirtualPsId, "pv_feed_in_percent", "",
			entries.GetPercent("pv_feed_in", "daily_pv_energy"))

		// @TODO - Add this calculation.
		DailyPvEnergy := entries.GetFloatValue("daily_total_energy") - entries.GetFloatValue("daily_purchased_energy")
		fmt.Sprintf("%f", DailyPvEnergy)
		entries.FromRefAddFloat("daily_total_energy",
			api.VirtualPsId, "daily_pv_energy_percent", "",
			api.GetPercent(DailyPvEnergy, entries.GetValue("daily_total_energy")))
		entries.FromRefAddFloat("daily_total_energy",
			api.VirtualPsId, "daily_purchased_energy_percent", "",
			entries.GetPercent("daily_purchased_energy", "daily_total_energy"))

		entries.FromRefAddFloat("pv_power",
			api.VirtualPsId, "pv_power_to_load", "",
			entries.GetFloatValue("pv_power")-entries.GetFloatValue("pv_power_to_battery")-entries.GetFloatValue("pv_power_to_grid"))

		// Battery
		entries.FromRefAddFloat("pv_power_to_battery",
			api.VirtualPsId, "battery_power", "",
			entries.LowerUpper("pv_power_to_battery", "battery_power_to_load"))
		entries.FromRefAddFloat("pv_power_to_battery",
			api.VirtualPsId, "battery_power_to_grid", "",
			0.0)

		// Grid
		entries.FromRefAddFloat("grid_power_to_load",
			api.VirtualPsId, "grid_power", "",
			entries.LowerUpper("pv_power_to_grid", "grid_power_to_load"))
		entries.FromRefAddFloat("grid_power_to_load",
			api.VirtualPsId, "grid_power_to_battery", "",
			0.0)

		entries.FromRefAddState("pv_power", api.VirtualPsId, "pv_power_active", "")
		entries.FromRefAddState("battery_power", api.VirtualPsId, "battery_power_active", "")
		entries.FromRefAddState("grid_power", api.VirtualPsId, "grid_power_active", "")
		entries.FromRefAddState("load_power", api.VirtualPsId, "load_power_active", "")

		entries.FromRefAddState("pv_power_to_battery", api.VirtualPsId, "pv_power_to_battery_active", "")
		entries.FromRefAddState("pv_power_to_load", api.VirtualPsId, "pv_power_to_load_active", "")
		entries.FromRefAddState("pv_power_to_grid", api.VirtualPsId, "pv_power_to_grid_active", "")

		entries.FromRefAddState("battery_power_to_load", api.VirtualPsId, "battery_power_to_load_active", "")
		entries.FromRefAddState("battery_power_to_grid", api.VirtualPsId, "battery_power_to_grid_active", "")

		entries.FromRefAddState("grid_power_to_load", api.VirtualPsId, "grid_power_to_load_active", "")
		entries.FromRefAddState("grid_power_to_battery", api.VirtualPsId, "grid_power_to_battery_active", "")

		entries.FromRefAddFloat("pv_battery_charge",
			api.VirtualPsId, "battery_energy", "",
			entries.LowerUpper("pv_battery_charge", "battery_discharge"))

		entries.FromRefAddFloat("pv_feed_in",
			api.VirtualPsId, "grid_energy", "",
			entries.LowerUpper("pv_feed_in", "daily_purchased_energy"))

		// for _, pid := range entries.Order {
		// 	// entries[pid].Index = i
		// 	ret.Entries = append(ret.Entries, entries.Entries[pid])
		// }
	}

	return entries
}

// func (ref *EntryMap) getFloatValue(entry string) float64 {
// 	return (*ref)[entry].ValueFloat
// }
//
// func lowerUpper(lower api.DataEntry, upper api.DataEntry) float64 {
// 	if lower.ValueFloat > 0 {
// 		return 0 - lower.ValueFloat
// 	}
// 	return upper.ValueFloat
// }
//
// func getPercent(value api.DataEntry, max api.DataEntry) float64 {
// 	if max.ValueFloat == 0 {
// 		return 0
// 	}
// 	return (value.ValueFloat / max.ValueFloat) * 100
// }

// func addState(now api.DateTime, point string, name string, state bool, index int) api.DataEntry {
// 	return add(now, "virtual", point, name, api.UnitValue{ Value: fmt.Sprintf("%v", state), Unit: "binary"}, index)
// }
//
// func addValue(now api.DateTime, point string, name string, value string, unit string, index int) api.DataEntry {
// 	return add(now, "virtual", point, name, api.UnitValue{ Value: value, Unit: unit}, index)
//
// 	// vt := api.GetPoint(psId, point)
// 	// if !vt.Valid {
// 	// 	vt = &api.PointStruct{
// 	// 		PsKey:       psId,
// 	// 		Id:          point,
// 	// 		Description: name,
// 	// 		Unit:        "",
// 	// 		Type:        "PointTypeInstant",
// 	// 	}
// 	// }
// 	// return api.DataEntry {
// 	// 	Date:           now,
// 	// 	PointId:        api.NameDevicePoint(psId, point),
// 	// 	PointGroupName: "Summary",
// 	// 	PointName:      name,
// 	// 	Value:          value,
// 	// 	Unit:           "",
// 	// 	ValueType:      vt,
// 	// 	Index:          index,
// 	// }
// }
//
// func addIntValue(now api.DateTime, point string, name string, value int64, unit string, index int) api.DataEntry {
// 	return add(now, "virtual", point, name, api.UnitValue{ Value: strconv.FormatInt(value, 10), Unit: unit }, index)
// }
//
// func addFloatValue(now api.DateTime, point string, name string, value float64, unit string, index int) api.DataEntry {
// 	return add(now, "virtual", point, name, api.UnitValue{ Value: api.Float64ToString(value), Unit: unit }, index)
// }
//
// func addFloatValue(ref api.DataEntry, psId string, point string, name string, index int) api.DataEntry {
// 	ref.PointId = psId
// 	ref.PointName = point
// 	return add(now, "virtual", point, name, api.UnitValue{ Value: api.Float64ToString(value), Unit: unit }, index)
// }
//
// func addVirtualState(ref api.DataEntry, point string, name string) api.DataEntry {
//
// 	return api.DataEntry {
// 		Date:           ref.Date,
// 		PointId:        api.NameDevicePoint("virtual", point),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          fmt.Sprintf("%v", isActive(ref.ValueFloat)),
// 		ValueFloat:     0,
// 		Unit:           "binary",
// 		ValueType:      &api.Point {
// 			PsKey:       "virtual",
// 			Id:          point,
// 			Description: name,
// 			Unit:        "binary",
// 		},
// 		Index:          0,
// 	}
// }
//
// func addVirtualValue(ref api.DataEntry, point string, name string, value float64) api.DataEntry {
// 	return api.DataEntry {
// 		Date:           ref.Date,
// 		PointId:        api.NameDevicePoint("virtual", point),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          api.Float64ToString(value),
// 		ValueFloat:     value,
// 		Unit:           ref.Unit,
// 		ValueType:      &api.Point {
// 			PsKey:       "virtual",
// 			Id:          point,
// 			Description: name,
// 			Unit:        ref.Unit,
// 		},
// 		Index:          0,
// 	}
// }
//
// func addVirtualAlias(ref api.DataEntry, point string, name string) api.DataEntry {
// 	ref.PointId = api.NameDevicePoint("virtual", point)
// 	ref.PointGroupName = "Virtual"
// 	ref.PointName = name
// 	ref.ValueType.PsKey = "virtual"
// 	ref.ValueType.Description = name
// 	ref.ValueType.Id = point
// 	ref.Index = 0
//
// 	return ref
// }
//
// func add(now api.DateTime, psId string, point string, name string, value api.UnitValue) api.DataEntry {
// 	vt := api.GetPoint(psId, point)
// 	if !vt.Valid {
// 		vt = &api.Point{
// 			PsKey:       psId,
// 			Id:          point,
// 			Description: name,
// 			Unit:        value.Unit,
// 			Type:        "PointTypeInstant",
// 		}
// 	}
// 	return api.DataEntry {
// 		Date:           now,
// 		PointId:        api.NameDevicePoint(psId, point),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          value.Value,
// 		Unit:           value.Unit,
// 		ValueType:      vt,
// 		Index:          0,
// 	}
// }
//
// func addState(now api.DateTime, point string, name string, state bool, index int) api.DataEntry {
// 	return api.DataEntry {
// 		Date:           now,
// 		PointId:        api.NameDevicePoint("virtual", point),
// 		PointGroupName: "Virtual",
// 		PointName:      name,
// 		Value:          fmt.Sprintf("%v", state),
// 		Unit:           "binary",
// 		ValueType: &api.PointStruct{
// 			PsKey:       "virtual",
// 			Id:          point,
// 			Description: name,
// 			Unit:        "binary",
// 			Type:        "PointTypeInstant",
// 		},
// 		Index: index,
// 	}
// }
//
// func isActive(value float64) bool {
// 	if (value > 0.01) || (value < -0.01) {
// 		return true
// 	}
// 	return false
// }
