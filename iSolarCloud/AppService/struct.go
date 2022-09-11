package AppService

import (
	"fmt"

	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/acceptPsSharing"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/activateEmail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addConfig"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addDeviceRepair"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addDeviceToStructureForHousehold"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addDeviceToStructureForHouseholdByPsIdS"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addFault"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addFaultOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addFaultPlan"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addFaultRepairSteps"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addHouseholdEvaluation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addHouseholdLeaveMessage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addHouseholdOpinionFeedback"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addHouseholdWorkOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addOnDutyInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addOperRule"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addOrDelPsStructure"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addOrderStep"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addPowerStationForHousehold"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addPowerStationInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addReportConfigEmail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addSysAdvancedParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/addSysOrgNew"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/aliPayAppTest"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/auditOperRule"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchAddStationBySn"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchImportSN"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchInsertUserAndOrg"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchModifyDevicesInfoAndPropertis"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchProcessPlantReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchUpdateDeviceSim"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/batchUpdateUserIsAgreeGdpr"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/boundMobilePhone"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/boundUserMail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/caculateDeviceInputDiscrete"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/calculateDeviceDiscrete"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/calculateInitialCompensationData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/cancelDeliverMail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/cancelOrderScan"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/cancelParamSetTask"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/cancelPsSharing"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/cancelRechargeOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/changRechargeOrderToCancel"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/changeHouseholdUser2Installer"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/changeRemoteParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkDealerOrgCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkDevSnIsBelongsToUser"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkInverterResult"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkIsCanDoParamSet"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkIsIvScan"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkOssObjectExist"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkServiceIsConnect"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkTechnicalParameters"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUnitStatus"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUpRechargeDevicePaying"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUserAccountUnique"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUserAccountUniqueAll"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUserInfoUnique"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUserIsExist"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUserListIsExist"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/checkUserPassword"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/cloudDeploymentRecord"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/comfirmParamModel"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/communicationModuleDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/compareValidateCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/componentInfo2Cloud"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/confirmFault"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/confirmIvFault"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/confirmReportConfig"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/createAppkeyInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/createRenewInvoice"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealCommandReply"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealDeletePsFailPsDelete"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealFailRemoteUpgradeSubTasks"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealFailRemoteUpgradeTasks"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealFaultOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealGroupStringDisableOrEnable"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealNumberOfServiceCalls2Mysql"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealParamSettingAfterComplete"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealPsDataSupplement"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealPsReportEmailSend"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealRemoteUpgrade"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealSnElectrifyCheck"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealSysDeviceSimFlowInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/dealSysDeviceSimInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/definiteTimeDealSnExpRemind"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/definiteTimeDealSnStatus"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/delDeviceRepair"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/delOperRule"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/delayCallApiResidueTimes"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteComponent"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteCustomerEmployee"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteDeviceAccount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteDeviceSimById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteElectricitySettlementData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteFaultPlan"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteFirmwareFiles"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteHouseholdEvaluation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteHouseholdLeaveMessage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteHouseholdWorkOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteInverterSnInChnnl"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteModuleLog"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteOnDutyInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteOperateBillFile"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteOssObject"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deletePowerDevicePointById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deletePowerPicture"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deletePowerRobotInfoBySnAndPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deletePowerRobotSweepStrategy"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteProductionData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deletePs"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteRechargeOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteRegularlyConnectionInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteReportConfigEmailAddr"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteSysAdvancedParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteSysOrgNew"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteTemplate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deleteUserInfoAllByUserId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteDeleteTime"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteGetTime"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteInsertTime"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deviceInputDiscreteUpdateTime"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/devicePointsDataFromMySql"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/deviceReplace"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/editDeviceRepair"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/editOperRule"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/energyPovertyAlleviation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/energyTrend"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/exportParamSettingValPDF"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/exportPlantReportPDF"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/faultAutoClose"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/faultCloseRemindOrderHandler"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findCodeValueList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findEmgOrgInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findEnvironmentInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findFromHbaseAndRedis"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findInfoByuuid"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findLossAnalysisList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findOnDutyInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findPsType"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findSingleStationPR"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/findUserPassword"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/genTLSUserSigByUserAccount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/generateRandomPassword"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAPIServiceInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAccessedPermission"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAllDeviceByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAllPowerDeviceSetName"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAllPowerRobotViewInfoByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAllPsIdByOrgIds"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAllUserRemindCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAndOutletsAndUnit"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getApiCallsForAppkeys"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAreaInfoCodeByCounty"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAreaList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getAutoCreatePowerStation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getBackReadValue"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getBatchNewestPointData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCallApiResidueTimes"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getChangedPsListByTime"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getChnnlListByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCloudList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCloudServiceMappingConfig"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCommunicationDeviceConfigInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCommunicationModuleMonitorData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getComponentModelFactory"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getConfigList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getConnectionInfoBySnAndLocalPort"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCountDown"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCountryServiceInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCounty"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCustomerEmployee"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getCustomerList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDataFromHBase"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDataFromHbaseByRowKey"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevInstalledPowerByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevRecord"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevRunRecordList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevSimByList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevSimList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceAccountById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceFaultStatisticsData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceModelInfoList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevicePointMinuteDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevicePoints"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDevicePropertys"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceRepairDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceTechBranchCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceTypeInfoList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDeviceTypeList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getDstInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getElectricitySettlementData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getElectricitySettlementDetailData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getEncryptPublicKey"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFaultCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFaultDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFaultMsgByFaultCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFaultMsgListByPageWithYYYYMM"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFaultMsgListWithYYYYMM"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFaultPlanList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFileOperationRecordOne"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getFormulaFaultAnalyseList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getGroupStringCheckResult"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getGroupStringCheckRule"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHisData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHistoryInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdEvaluation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdLeaveMessage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdOpinionFeedback"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdPsInstallerByUserId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdStoragePsReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdUserInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdWorkOrderInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getHouseholdWorkOrderList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getI18nConfigByType"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getI18nFileInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getI18nInfoByKey"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getI18nVersion"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getIncomeSettingInfos"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInfoFromAMap"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInfomationFromRedis"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInstallInfoList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInstallerInfoByDealerOrgCodeOrId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInvertDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInverterDataCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInverterProcess"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getInverterUuidBytotalId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getIvEchartsData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getIvEchartsDataById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getKpiInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getListMiFromHBase"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getMapInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getMapMiFromHBase"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getMenuAndPrivileges"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getMicrogridEStoragePsReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getModuleLogInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getModuleLogTaskList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getNationProvJSON"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getNeedOpAsynOpRecordList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getNoticeInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getNumberOfServiceCalls"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOSSConfig"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOperRuleDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOperateBillFileId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOperateTicketForDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrCreateNetEaseUserToken"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrderDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrderDataSql2"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrderDatas"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrderDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrderStatistics"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrgIdNameByUserId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrgInfoByDealerOrgCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrgListByName"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrgListByUserId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOrgListForUser"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOssObjectStream"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getOwnerFaultConfigList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPListinfoFromMysql"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getParamSetTemplate4NewProtocol"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getParamSetTemplatePointInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getParamterSettingBase"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPhotoInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPlanedOrNotPsList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPlantReportPDFList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerChargeSettingInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerDeviceModelTechList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerDeviceModelTree"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerDevicePointInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerDeviceSetTaskDetailList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerDeviceSetTaskList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerFormulaFaultAnalyse"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerPictureList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerRobotInfoByRobotSn"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerRobotSweepAttrByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerRobotSweepStrategy"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerRobotSweepStrategyList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerSettingCharges"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerSettingHistoryRecords"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationBasicInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationForHousehold"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationPR"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationTableDataSql"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStationTableDataSqlCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerStatistics"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPowerTrendDayData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPrivateCloudValidityPeriod"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getProvInfoListByNationCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsAuthKey"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsCurveInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsDataSupplementTaskList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsDetailByUserTokens"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsDetailForSinglePage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsDetailWithPsType"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsHealthState"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsInstallerByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsInstallerOrgInfoByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsListByName"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsListForPsDataByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsListStaticData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsUser"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getPsWeatherList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRechargeOrderDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRechargeOrderItemDeviceList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRechargeOrderList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRegionalTree"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRemoteParamSettingList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeDeviceList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeScheduleDetails"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeSubTasksList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRemoteUpgradeTaskList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getReportData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getReportEmailConfigInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getReportExportColumns"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getReportListByUserId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRobotDynamicCleaningView"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRobotNumAndSweepCapacity"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getRuleUnit"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSendReportConfigCron"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSerialNum"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getShieldMapConditionList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSimIdBySnList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSingleIVData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSnChangeRecord"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSnConnectionInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getStationInfoSql"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSungwsConfigCache"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSungwsGlobalConfigCache"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSweepDevParamSetTemplate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSweepRobotDevList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSysMsg"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSysOrgNewList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSysOrgNewOne"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getSysUserById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getTableDataSql"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getTableDataSqlCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getTemplateByInfoType"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getTemplateList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUUIDByUpuuid"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUpTimePoint"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserByInstaller"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserDevOnlineOffineCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserGDPRAttrs"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserHavePowerStationCount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserInfoByUserAccounts"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getUserPsOrderList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getValFromHBase"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getValidateCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getValidateCodeAtRegister"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getWeatherInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getWechatPushConfig"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/getWorkInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/groupStringCheck"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/handleDevByCommunicationSN"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/householdResetPassBySN"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/immediatePayment"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/importExcelData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/incomeStatistics"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/informPush"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/insertEmgOrgInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/insightSynDeviceStructure2Cloud"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/intoDataToHbase"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/ipLocationQuery"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/isHave2GSn"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/judgeDevIsHasInitSetTemplate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/judgeIsSettingMan"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/listOssFiles"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/loadAreaInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/loadPowerStation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/login"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/loginByToken"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/logout"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/mobilePhoneHasBound"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifiedDeviceInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifyEmgOrgStruc"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifyFaultPlan"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifyOnDutyInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifyPassword"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifyPersonalUnitList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/modifyPsUser"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/moduleLogParamSet"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/operateOssFile"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/operationPowerRobotSweepStrategy"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/orgPowerReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/paramSetTryAgain"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/paramSetting"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/planPower"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/powerDevicePointList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/powerTrendChartData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/psForcastInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/psHourPointsValue"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryAllPsIdAndName"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryBatchCreatePsTaskList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryBatchSpeedyAddPowerStationResult"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryCardStatusCTCC"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryChildAccountList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryCompensationRecordData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryCompensationRecordList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryComponent"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryComponentTechnicalParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryCountryGridAndRelation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryCountryList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryCtrlTaskById"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceInfoForApp"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceListByUserId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceListForApp"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceModelTechnical"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDevicePointDayMonthYearDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDevicePointMinuteDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDevicePointsDayMonthYearDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceRealTimeDataByPsKeys"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceRepairList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryDeviceTypeInfoList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryEnvironmentList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFaultList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFaultPlanDetail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFaultRepairSteps"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFaultTypeAndLevelByCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFaultTypeByDevice"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFaultTypeByDevicePage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryFirmwareFilesPage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryInfotoAlert"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryInverterModelList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryInverterVersionList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryM2MCardInfoCMCC"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryM2MCardTermInfoCMCC"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryModelInfoByModelId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryMutiPointDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryNoticeList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryNumberOfRenewalReminders"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOperRules"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOrderList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOrderStep"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOrgGenerationReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOrgInfoList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOrgPowerElecPercent"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryOrgPsCompensationRecordList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryParamSettingTask"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPersonalUnitList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPointDataTopOne"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPowerStationInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsAreaByUserIdAndAreaCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsCompensationRecordList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsDataByDate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsIdList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsListByUserIdAndAreaCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsNameByPsId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsPrByDate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsProfit"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsReportComparativeAnalysisOfPowerGeneration"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPsStructureList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPuuidsByCommandTotalId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryPuuidsByCommandTotalId2"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryRepairRuleList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryReportListForManagementPage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryReportMsg"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/querySharingPs"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/querySysAdvancedParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryTimeBySN"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryTrafficByDateCTCC"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryTrafficCTCC"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUnitList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUnitUuidBytotalId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserBtnPri"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserByUserIds"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserExtensionAttribute"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserForStep"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserProcessPri"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUserWechatBindRel"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/queryUuidByTotalIdAndUuid"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/rechargeOrderSetMeal"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/renewSendReportConfirmEmail"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/reportList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveCustomerEmployee"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveDevSimList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveDeviceAccountBatchData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveEnviromentIncomeInfos"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveEnvironmentCurve"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveFirmwareFile"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveIncomeSettingInfos"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveOrUpdateGroupStringCheckRule"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveParamModel"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/savePowerCharges"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/savePowerDevicePoint"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/savePowerRobotInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/savePowerRobotSweepAttr"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/savePowerSettingCharges"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/savePowerSettingInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveProductionBatchData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveRechargeOrderObj"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveRechargeOrderOtherInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveRepair"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveReportExportColumns"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveSetParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveSysUserMsg"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/saveTemplate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/searchM2MMonthFlowCMCC"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/selectSysTranslationNames"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sendPsTimeZoneInstruction"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/setUpFormulaFaultAnalyse"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/setUserGDPRAttrs"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/settingNotice"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/shareMyPs"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sharePsBySN"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/showInverterByUnit"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/showOnlineUsers"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/showWarning"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/snIsExist"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/snsIsExist"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/speedyAddPowerStation"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/stationDeviceHistoryDataList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/stationUnitsList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/stationsDiscreteData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/stationsIncomeList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/stationsPointReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/stationsYearPlanReport"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sureAndImportSelettlementData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sweepDevParamSet"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sweepDevRunControl"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sweepDevStrategyIssue"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/sysTimeZoneList"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/unLockUser"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/unlockChildAccount"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateCommunicationModuleState"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateDevInstalledPower"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateFault"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateFaultData"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateFaultMsgByFaultCode"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateFaultStatus"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateHouseholdWorkOrder"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateInverterSn2ModuleSn"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateOperateTicketAttachmentId"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateOrderDeviceByCustomerService"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateOwnerFaultConfig"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateParamSettingSysMsg"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePlatformLevelFaultLevel"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePowerDevicePoint"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePowerRobotInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePowerRobotSweepAttr"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePowerStationForHousehold"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePowerStationInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updatePowerUserInfo"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateReportConfigByEmailAddr"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateShareAttr"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateSnIsSureFlag"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateStationPics"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateSysAdvancedParam"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateSysOrgNew"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateTemplate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateUinfoNetEaseUser"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateUserExtensionAttribute"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateUserLanguage"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateUserPosition"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/updateUserUpOrg"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/upgrade"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/upgrate"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/uploadFileToOss"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/userAgreeGdprProtocol"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/userInfoUniqueCheck"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/userMailHasBound"
	"github.com/jpillora/GoSungrow/iSolarCloud/AppService/userRegister"
	"github.com/jpillora/GoSungrow/iSolarCloud/api"
	"github.com/jpillora/GoSungrow/iSolarCloud/api/output"
)

var _ api.Area = (*Area)(nil)

type Area api.AreaStruct

func init() {
	// name := api.GetArea(Area{})
	// fmt.Printf("Name: %s\n", name)
}

func Init(apiRoot api.Web) Area {
	area := Area{
		ApiRoot: apiRoot,
		Name:    api.GetArea(Area{}),
		EndPoints: api.TypeEndPoints{
			api.GetName(login.EndPoint{}):                       login.Init(apiRoot),
			api.GetName(findPsType.EndPoint{}):                  findPsType.Init(apiRoot),
			api.GetName(getPowerDevicePointNames.EndPoint{}):    getPowerDevicePointNames.Init(apiRoot),
			api.GetName(getPowerStatistics.EndPoint{}):          getPowerStatistics.Init(apiRoot),
			api.GetName(getPsDetailWithPsType.EndPoint{}):       getPsDetailWithPsType.Init(apiRoot),
			api.GetName(getPsList.EndPoint{}):                   getPsList.Init(apiRoot),
			api.GetName(getHouseholdStoragePsReport.EndPoint{}): getHouseholdStoragePsReport.Init(apiRoot),
			api.GetName(queryUnitList.EndPoint{}):               queryUnitList.Init(apiRoot),
			api.GetName(queryMutiPointDataList.EndPoint{}):      queryMutiPointDataList.Init(apiRoot),
			api.GetName(communicationModuleDetail.EndPoint{}):   communicationModuleDetail.Init(apiRoot),
			api.GetName(getHistoryInfo.EndPoint{}):              getHistoryInfo.Init(apiRoot), // @TODO - Not working.
			api.GetName(energyTrend.EndPoint{}):                 energyTrend.Init(apiRoot),
			api.GetName(getTemplateList.EndPoint{}):             getTemplateList.Init(apiRoot),
			api.GetName(getDevicePoints.EndPoint{}):             getDevicePoints.Init(apiRoot), // @TODO - Doesn't seem to generate anything.
			api.GetName(getPsDetail.EndPoint{}):                 getPsDetail.Init(apiRoot),
			api.GetName(getPsHealthState.EndPoint{}):            getPsHealthState.Init(apiRoot),
			api.GetName(queryAllPsIdAndName.EndPoint{}):         queryAllPsIdAndName.Init(apiRoot),
			api.GetName(getAllPowerDeviceSetName.EndPoint{}):    getAllPowerDeviceSetName.Init(apiRoot),
			api.GetName(getOrgListByName.EndPoint{}):            getOrgListByName.Init(apiRoot),
			api.GetName(getPsListByName.EndPoint{}):             getPsListByName.Init(apiRoot),
			api.GetName(checkUnitStatus.EndPoint{}):             checkUnitStatus.Init(apiRoot), // @TODO - Returns 404 error. Disabled.
			api.GetName(getDeviceInfo.EndPoint{}):               getDeviceInfo.Init(apiRoot),   // @TODO - Returns null.
			api.GetName(getDeviceList.EndPoint{}):               getDeviceList.Init(apiRoot),
			api.GetName(powerDevicePointList.EndPoint{}):        powerDevicePointList.Init(apiRoot),  // @TODO - CRITICAL - GetByJson all point_id definitions.
			api.GetName(queryDeviceListForApp.EndPoint{}):       queryDeviceListForApp.Init(apiRoot), // @TODO - CRITICAL - Show more detail info on devices.
			api.GetName(queryDeviceListByUserId.EndPoint{}):     queryDeviceListByUserId.Init(apiRoot),
			api.GetName(getPsListStaticData.EndPoint{}):         getPsListStaticData.Init(apiRoot),
			api.GetName(getPsReport.EndPoint{}):                 getPsReport.Init(apiRoot),
			api.GetName(getPsUser.EndPoint{}):                   getPsUser.Init(apiRoot),
			api.GetName(getPsWeatherList.EndPoint{}):            getPsWeatherList.Init(apiRoot),
			api.GetName(getAreaList.EndPoint{}):                 getAreaList.Init(apiRoot),
			api.GetName(getCloudList.EndPoint{}):                getCloudList.Init(apiRoot),
			api.GetName(getConfigList.EndPoint{}):               getConfigList.Init(apiRoot),          // @TODO - Returns 404 error. Disabled.
			api.GetName(getDeviceModelInfoList.EndPoint{}):      getDeviceModelInfoList.Init(apiRoot), // @TODO - CRITICAL - Show all device model info.
			api.GetName(getDeviceTypeList.EndPoint{}):           getDeviceTypeList.Init(apiRoot),
			api.GetName(getDeviceTypeInfoList.EndPoint{}):       getDeviceTypeInfoList.Init(apiRoot), // @TODO - CRITICAL - Show all device types.
			api.GetName(getInvertDataList.EndPoint{}):           getInvertDataList.Init(apiRoot),
			api.GetName(queryUserList.EndPoint{}):               queryUserList.Init(apiRoot),
			api.GetName(getPowerPictureList.EndPoint{}):         getPowerPictureList.Init(apiRoot),
			api.GetName(getUserList.EndPoint{}):                 getUserList.Init(apiRoot),
			api.GetName(getUserPsOrderList.EndPoint{}):          getUserPsOrderList.Init(apiRoot),
			api.GetName(reportList.EndPoint{}):                  reportList.Init(apiRoot),

			api.GetName(getKpiInfo.EndPoint{}): getKpiInfo.Init(apiRoot),
			api.GetName(getMapInfo.EndPoint{}): getMapInfo.Init(apiRoot), // @TODO - No data.

			// Tasks
			api.GetName(getPowerDeviceSetTaskDetailList.EndPoint{}): getPowerDeviceSetTaskDetailList.Init(apiRoot), // @TODO - Needs to be checked.
			api.GetName(getPowerDeviceSetTaskList.EndPoint{}):       getPowerDeviceSetTaskList.Init(apiRoot),
			api.GetName(getPsDataSupplementTaskList.EndPoint{}):     getPsDataSupplementTaskList.Init(apiRoot),
			api.GetName(queryBatchCreatePsTaskList.EndPoint{}):      queryBatchCreatePsTaskList.Init(apiRoot),
			api.GetName(getModuleLogTaskList.EndPoint{}):            getModuleLogTaskList.Init(apiRoot),
			api.GetName(getRemoteUpgradeTaskList.EndPoint{}):        getRemoteUpgradeTaskList.Init(apiRoot),     // @TODO - No data.
			api.GetName(getRemoteUpgradeSubTasksList.EndPoint{}):    getRemoteUpgradeSubTasksList.Init(apiRoot), // @TODO - No data.
			api.GetName(queryParamSettingTask.EndPoint{}):           queryParamSettingTask.Init(apiRoot),
			api.GetName(queryCtrlTaskById.EndPoint{}):               queryCtrlTaskById.Init(apiRoot), // @TODO - Returns 404 error. Disabled.

			// Use the following for all critical data every 5 minutes.
			api.GetName(queryDeviceList.EndPoint{}): queryDeviceList.Init(apiRoot), // @TODO - This gives you ALL info at 5 min intervals.

			// Disabled from here on.

			api.GetName(cancelParamSetTask.EndPoint{}):            cancelParamSetTask.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeSubTasks.EndPoint{}): dealFailRemoteUpgradeSubTasks.Init(apiRoot),
			api.GetName(dealFailRemoteUpgradeTasks.EndPoint{}):    dealFailRemoteUpgradeTasks.Init(apiRoot),

			api.GetName(updateTemplate.EndPoint{}):                  updateTemplate.Init(apiRoot),
			api.GetName(getTemplateByInfoType.EndPoint{}):           getTemplateByInfoType.Init(apiRoot),
			api.GetName(deleteTemplate.EndPoint{}):                  deleteTemplate.Init(apiRoot),
			api.GetName(getParamSetTemplate4NewProtocol.EndPoint{}): getParamSetTemplate4NewProtocol.Init(apiRoot),
			api.GetName(getParamSetTemplatePointInfo.EndPoint{}):    getParamSetTemplatePointInfo.Init(apiRoot),
			api.GetName(getSweepDevParamSetTemplate.EndPoint{}):     getSweepDevParamSetTemplate.Init(apiRoot),
			api.GetName(judgeDevIsHasInitSetTemplate.EndPoint{}):    judgeDevIsHasInitSetTemplate.Init(apiRoot),
			api.GetName(saveTemplate.EndPoint{}):                    saveTemplate.Init(apiRoot),

			api.GetName(queryDevicePointMinuteDataList.EndPoint{}):          queryDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(getDevicePointMinuteDataList.EndPoint{}):            getDevicePointMinuteDataList.Init(apiRoot),
			api.GetName(queryDeviceInfo.EndPoint{}):                         queryDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceInfoForApp.EndPoint{}):                   queryDeviceInfoForApp.Init(apiRoot),
			api.GetName(queryDeviceModelTechnical.EndPoint{}):               queryDeviceModelTechnical.Init(apiRoot),
			api.GetName(queryDevicePointDayMonthYearDataList.EndPoint{}):    queryDevicePointDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDevicePointsDayMonthYearDataList.EndPoint{}):   queryDevicePointsDayMonthYearDataList.Init(apiRoot),
			api.GetName(queryDeviceRealTimeDataByPsKeys.EndPoint{}):         queryDeviceRealTimeDataByPsKeys.Init(apiRoot),
			api.GetName(addDeviceRepair.EndPoint{}):                         addDeviceRepair.Init(apiRoot),
			api.GetName(addDeviceToStructureForHousehold.EndPoint{}):        addDeviceToStructureForHousehold.Init(apiRoot),
			api.GetName(addDeviceToStructureForHouseholdByPsIdS.EndPoint{}): addDeviceToStructureForHouseholdByPsIdS.Init(apiRoot),
			api.GetName(batchModifyDevicesInfoAndPropertis.EndPoint{}):      batchModifyDevicesInfoAndPropertis.Init(apiRoot),
			api.GetName(batchUpdateDeviceSim.EndPoint{}):                    batchUpdateDeviceSim.Init(apiRoot),
			api.GetName(caculateDeviceInputDiscrete.EndPoint{}):             caculateDeviceInputDiscrete.Init(apiRoot),
			api.GetName(calculateDeviceDiscrete.EndPoint{}):                 calculateDeviceDiscrete.Init(apiRoot),
			api.GetName(checkUpRechargeDevicePaying.EndPoint{}):             checkUpRechargeDevicePaying.Init(apiRoot),
			api.GetName(dealSysDeviceSimFlowInfo.EndPoint{}):                dealSysDeviceSimFlowInfo.Init(apiRoot),
			api.GetName(dealSysDeviceSimInfo.EndPoint{}):                    dealSysDeviceSimInfo.Init(apiRoot),
			api.GetName(delDeviceRepair.EndPoint{}):                         delDeviceRepair.Init(apiRoot),
			api.GetName(deleteDeviceAccount.EndPoint{}):                     deleteDeviceAccount.Init(apiRoot),
			api.GetName(deleteDeviceSimById.EndPoint{}):                     deleteDeviceSimById.Init(apiRoot),
			api.GetName(deletePowerDevicePointById.EndPoint{}):              deletePowerDevicePointById.Init(apiRoot),
			api.GetName(editDeviceRepair.EndPoint{}):                        editDeviceRepair.Init(apiRoot),
			api.GetName(getAllDeviceByPsId.EndPoint{}):                      getAllDeviceByPsId.Init(apiRoot),
			api.GetName(getCommunicationDeviceConfigInfo.EndPoint{}):        getCommunicationDeviceConfigInfo.Init(apiRoot),
			api.GetName(getDeviceAccountById.EndPoint{}):                    getDeviceAccountById.Init(apiRoot),
			api.GetName(getDeviceFaultStatisticsData.EndPoint{}):            getDeviceFaultStatisticsData.Init(apiRoot),
			api.GetName(getDevicePropertys.EndPoint{}):                      getDevicePropertys.Init(apiRoot),
			api.GetName(getDeviceRepairDetail.EndPoint{}):                   getDeviceRepairDetail.Init(apiRoot),
			api.GetName(getDeviceTechBranchCount.EndPoint{}):                getDeviceTechBranchCount.Init(apiRoot),
			api.GetName(getPowerDeviceModelTechList.EndPoint{}):             getPowerDeviceModelTechList.Init(apiRoot),
			api.GetName(getPowerDeviceModelTree.EndPoint{}):                 getPowerDeviceModelTree.Init(apiRoot),
			api.GetName(getPowerDevicePointInfo.EndPoint{}):                 getPowerDevicePointInfo.Init(apiRoot),
			api.GetName(getRechargeOrderItemDeviceList.EndPoint{}):          getRechargeOrderItemDeviceList.Init(apiRoot),
			api.GetName(getRemoteUpgradeDeviceList.EndPoint{}):              getRemoteUpgradeDeviceList.Init(apiRoot),
			api.GetName(insightSynDeviceStructure2Cloud.EndPoint{}):         insightSynDeviceStructure2Cloud.Init(apiRoot),
			api.GetName(modifiedDeviceInfo.EndPoint{}):                      modifiedDeviceInfo.Init(apiRoot),
			api.GetName(queryDeviceRepairList.EndPoint{}):                   queryDeviceRepairList.Init(apiRoot),
			api.GetName(queryDeviceTypeInfoList.EndPoint{}):                 queryDeviceTypeInfoList.Init(apiRoot),
			api.GetName(queryFaultTypeByDevice.EndPoint{}):                  queryFaultTypeByDevice.Init(apiRoot),
			api.GetName(queryFaultTypeByDevicePage.EndPoint{}):              queryFaultTypeByDevicePage.Init(apiRoot),
			api.GetName(saveDeviceAccountBatchData.EndPoint{}):              saveDeviceAccountBatchData.Init(apiRoot),
			api.GetName(savePowerDevicePoint.EndPoint{}):                    savePowerDevicePoint.Init(apiRoot),
			api.GetName(stationDeviceHistoryDataList.EndPoint{}):            stationDeviceHistoryDataList.Init(apiRoot),
			api.GetName(updateOrderDeviceByCustomerService.EndPoint{}):      updateOrderDeviceByCustomerService.Init(apiRoot),
			api.GetName(updatePowerDevicePoint.EndPoint{}):                  updatePowerDevicePoint.Init(apiRoot),

			api.GetName(psHourPointsValue.EndPoint{}):                 psHourPointsValue.Init(apiRoot),
			api.GetName(getPsCurveInfo.EndPoint{}):                    getPsCurveInfo.Init(apiRoot),
			api.GetName(exportPlantReportPDF.EndPoint{}):              exportPlantReportPDF.Init(apiRoot),
			api.GetName(getTableDataSql.EndPoint{}):                   getTableDataSql.Init(apiRoot),
			api.GetName(getTableDataSqlCount.EndPoint{}):              getTableDataSqlCount.Init(apiRoot),
			api.GetName(getPListinfoFromMysql.EndPoint{}):             getPListinfoFromMysql.Init(apiRoot),
			api.GetName(getDataFromHBase.EndPoint{}):                  getDataFromHBase.Init(apiRoot),
			api.GetName(getListMiFromHBase.EndPoint{}):                getListMiFromHBase.Init(apiRoot),
			api.GetName(getMapMiFromHBase.EndPoint{}):                 getMapMiFromHBase.Init(apiRoot),
			api.GetName(getValFromHBase.EndPoint{}):                   getValFromHBase.Init(apiRoot),
			api.GetName(getPowerStationTableDataSql.EndPoint{}):       getPowerStationTableDataSql.Init(apiRoot),
			api.GetName(getPowerStationTableDataSqlCount.EndPoint{}):  getPowerStationTableDataSqlCount.Init(apiRoot),
			api.GetName(getDataFromHbaseByRowKey.EndPoint{}):          getDataFromHbaseByRowKey.Init(apiRoot),
			api.GetName(findInfoByuuid.EndPoint{}):                    findInfoByuuid.Init(apiRoot),
			api.GetName(getStationInfoSql.EndPoint{}):                 getStationInfoSql.Init(apiRoot),
			api.GetName(getPowerTrendDayData.EndPoint{}):              getPowerTrendDayData.Init(apiRoot),
			api.GetName(powerTrendChartData.EndPoint{}):               powerTrendChartData.Init(apiRoot),
			api.GetName(getUpTimePoint.EndPoint{}):                    getUpTimePoint.Init(apiRoot),
			api.GetName(getSerialNum.EndPoint{}):                      getSerialNum.Init(apiRoot),
			api.GetName(getPsAuthKey.EndPoint{}):                      getPsAuthKey.Init(apiRoot),
			api.GetName(getApiCallsForAppkeys.EndPoint{}):             getApiCallsForAppkeys.Init(apiRoot),
			api.GetName(exportParamSettingValPDF.EndPoint{}):          exportParamSettingValPDF.Init(apiRoot),
			api.GetName(devicePointsDataFromMySql.EndPoint{}):         devicePointsDataFromMySql.Init(apiRoot),
			api.GetName(acceptPsSharing.EndPoint{}):                   acceptPsSharing.Init(apiRoot),
			api.GetName(activateEmail.EndPoint{}):                     activateEmail.Init(apiRoot),
			api.GetName(addConfig.EndPoint{}):                         addConfig.Init(apiRoot),
			api.GetName(addFault.EndPoint{}):                          addFault.Init(apiRoot),
			api.GetName(addFaultOrder.EndPoint{}):                     addFaultOrder.Init(apiRoot),
			api.GetName(addFaultPlan.EndPoint{}):                      addFaultPlan.Init(apiRoot),
			api.GetName(addFaultRepairSteps.EndPoint{}):               addFaultRepairSteps.Init(apiRoot),
			api.GetName(addHouseholdEvaluation.EndPoint{}):            addHouseholdEvaluation.Init(apiRoot),
			api.GetName(addHouseholdLeaveMessage.EndPoint{}):          addHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(addHouseholdOpinionFeedback.EndPoint{}):       addHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(addHouseholdWorkOrder.EndPoint{}):             addHouseholdWorkOrder.Init(apiRoot),
			api.GetName(addOnDutyInfo.EndPoint{}):                     addOnDutyInfo.Init(apiRoot),
			api.GetName(addOperRule.EndPoint{}):                       addOperRule.Init(apiRoot),
			api.GetName(addOrDelPsStructure.EndPoint{}):               addOrDelPsStructure.Init(apiRoot),
			api.GetName(addOrderStep.EndPoint{}):                      addOrderStep.Init(apiRoot),
			api.GetName(addPowerStationForHousehold.EndPoint{}):       addPowerStationForHousehold.Init(apiRoot),
			api.GetName(addPowerStationInfo.EndPoint{}):               addPowerStationInfo.Init(apiRoot),
			api.GetName(addReportConfigEmail.EndPoint{}):              addReportConfigEmail.Init(apiRoot),
			api.GetName(addSysAdvancedParam.EndPoint{}):               addSysAdvancedParam.Init(apiRoot),
			api.GetName(addSysOrgNew.EndPoint{}):                      addSysOrgNew.Init(apiRoot),
			api.GetName(aliPayAppTest.EndPoint{}):                     aliPayAppTest.Init(apiRoot),
			api.GetName(auditOperRule.EndPoint{}):                     auditOperRule.Init(apiRoot),
			api.GetName(batchAddStationBySn.EndPoint{}):               batchAddStationBySn.Init(apiRoot),
			api.GetName(batchImportSN.EndPoint{}):                     batchImportSN.Init(apiRoot),
			api.GetName(batchInsertUserAndOrg.EndPoint{}):             batchInsertUserAndOrg.Init(apiRoot),
			api.GetName(batchProcessPlantReport.EndPoint{}):           batchProcessPlantReport.Init(apiRoot),
			api.GetName(batchUpdateUserIsAgreeGdpr.EndPoint{}):        batchUpdateUserIsAgreeGdpr.Init(apiRoot),
			api.GetName(boundMobilePhone.EndPoint{}):                  boundMobilePhone.Init(apiRoot),
			api.GetName(boundUserMail.EndPoint{}):                     boundUserMail.Init(apiRoot),
			api.GetName(calculateInitialCompensationData.EndPoint{}):  calculateInitialCompensationData.Init(apiRoot),
			api.GetName(cancelDeliverMail.EndPoint{}):                 cancelDeliverMail.Init(apiRoot),
			api.GetName(cancelOrderScan.EndPoint{}):                   cancelOrderScan.Init(apiRoot),
			api.GetName(cancelPsSharing.EndPoint{}):                   cancelPsSharing.Init(apiRoot),
			api.GetName(cancelRechargeOrder.EndPoint{}):               cancelRechargeOrder.Init(apiRoot),
			api.GetName(changRechargeOrderToCancel.EndPoint{}):        changRechargeOrderToCancel.Init(apiRoot),
			api.GetName(changeHouseholdUser2Installer.EndPoint{}):     changeHouseholdUser2Installer.Init(apiRoot),
			api.GetName(changeRemoteParam.EndPoint{}):                 changeRemoteParam.Init(apiRoot),
			api.GetName(checkDealerOrgCode.EndPoint{}):                checkDealerOrgCode.Init(apiRoot),
			api.GetName(checkDevSnIsBelongsToUser.EndPoint{}):         checkDevSnIsBelongsToUser.Init(apiRoot),
			api.GetName(checkInverterResult.EndPoint{}):               checkInverterResult.Init(apiRoot),
			api.GetName(checkIsCanDoParamSet.EndPoint{}):              checkIsCanDoParamSet.Init(apiRoot),
			api.GetName(checkIsIvScan.EndPoint{}):                     checkIsIvScan.Init(apiRoot),
			api.GetName(checkOssObjectExist.EndPoint{}):               checkOssObjectExist.Init(apiRoot),
			api.GetName(checkServiceIsConnect.EndPoint{}):             checkServiceIsConnect.Init(apiRoot),
			api.GetName(checkTechnicalParameters.EndPoint{}):          checkTechnicalParameters.Init(apiRoot),
			api.GetName(checkUserAccountUnique.EndPoint{}):            checkUserAccountUnique.Init(apiRoot),
			api.GetName(checkUserAccountUniqueAll.EndPoint{}):         checkUserAccountUniqueAll.Init(apiRoot),
			api.GetName(checkUserInfoUnique.EndPoint{}):               checkUserInfoUnique.Init(apiRoot),
			api.GetName(checkUserIsExist.EndPoint{}):                  checkUserIsExist.Init(apiRoot),
			api.GetName(checkUserListIsExist.EndPoint{}):              checkUserListIsExist.Init(apiRoot),
			api.GetName(checkUserPassword.EndPoint{}):                 checkUserPassword.Init(apiRoot),
			api.GetName(cloudDeploymentRecord.EndPoint{}):             cloudDeploymentRecord.Init(apiRoot),
			api.GetName(comfirmParamModel.EndPoint{}):                 comfirmParamModel.Init(apiRoot),
			api.GetName(compareValidateCode.EndPoint{}):               compareValidateCode.Init(apiRoot),
			api.GetName(componentInfo2Cloud.EndPoint{}):               componentInfo2Cloud.Init(apiRoot),
			api.GetName(confirmFault.EndPoint{}):                      confirmFault.Init(apiRoot),
			api.GetName(confirmIvFault.EndPoint{}):                    confirmIvFault.Init(apiRoot),
			api.GetName(confirmReportConfig.EndPoint{}):               confirmReportConfig.Init(apiRoot),
			api.GetName(createAppkeyInfo.EndPoint{}):                  createAppkeyInfo.Init(apiRoot),
			api.GetName(createRenewInvoice.EndPoint{}):                createRenewInvoice.Init(apiRoot),
			api.GetName(dealCommandReply.EndPoint{}):                  dealCommandReply.Init(apiRoot),
			api.GetName(dealDeletePsFailPsDelete.EndPoint{}):          dealDeletePsFailPsDelete.Init(apiRoot),
			api.GetName(dealFaultOrder.EndPoint{}):                    dealFaultOrder.Init(apiRoot),
			api.GetName(dealGroupStringDisableOrEnable.EndPoint{}):    dealGroupStringDisableOrEnable.Init(apiRoot),
			api.GetName(dealNumberOfServiceCalls2Mysql.EndPoint{}):    dealNumberOfServiceCalls2Mysql.Init(apiRoot),
			api.GetName(dealParamSettingAfterComplete.EndPoint{}):     dealParamSettingAfterComplete.Init(apiRoot),
			api.GetName(dealPsDataSupplement.EndPoint{}):              dealPsDataSupplement.Init(apiRoot),
			api.GetName(dealPsReportEmailSend.EndPoint{}):             dealPsReportEmailSend.Init(apiRoot),
			api.GetName(dealRemoteUpgrade.EndPoint{}):                 dealRemoteUpgrade.Init(apiRoot),
			api.GetName(dealSnElectrifyCheck.EndPoint{}):              dealSnElectrifyCheck.Init(apiRoot),
			api.GetName(definiteTimeDealSnExpRemind.EndPoint{}):       definiteTimeDealSnExpRemind.Init(apiRoot),
			api.GetName(definiteTimeDealSnStatus.EndPoint{}):          definiteTimeDealSnStatus.Init(apiRoot),
			api.GetName(delOperRule.EndPoint{}):                       delOperRule.Init(apiRoot),
			api.GetName(delayCallApiResidueTimes.EndPoint{}):          delayCallApiResidueTimes.Init(apiRoot),
			api.GetName(deleteComponent.EndPoint{}):                   deleteComponent.Init(apiRoot),
			api.GetName(deleteCustomerEmployee.EndPoint{}):            deleteCustomerEmployee.Init(apiRoot),
			api.GetName(deleteElectricitySettlementData.EndPoint{}):   deleteElectricitySettlementData.Init(apiRoot),
			api.GetName(deleteFaultPlan.EndPoint{}):                   deleteFaultPlan.Init(apiRoot),
			api.GetName(deleteFirmwareFiles.EndPoint{}):               deleteFirmwareFiles.Init(apiRoot),
			api.GetName(deleteHouseholdEvaluation.EndPoint{}):         deleteHouseholdEvaluation.Init(apiRoot),
			api.GetName(deleteHouseholdLeaveMessage.EndPoint{}):       deleteHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(deleteHouseholdWorkOrder.EndPoint{}):          deleteHouseholdWorkOrder.Init(apiRoot),
			api.GetName(deleteInverterSnInChnnl.EndPoint{}):           deleteInverterSnInChnnl.Init(apiRoot),
			api.GetName(deleteModuleLog.EndPoint{}):                   deleteModuleLog.Init(apiRoot),
			api.GetName(deleteOnDutyInfo.EndPoint{}):                  deleteOnDutyInfo.Init(apiRoot),
			api.GetName(deleteOperateBillFile.EndPoint{}):             deleteOperateBillFile.Init(apiRoot),
			api.GetName(deleteOssObject.EndPoint{}):                   deleteOssObject.Init(apiRoot),
			api.GetName(deletePowerPicture.EndPoint{}):                deletePowerPicture.Init(apiRoot),
			api.GetName(deletePowerRobotInfoBySnAndPsId.EndPoint{}):   deletePowerRobotInfoBySnAndPsId.Init(apiRoot),
			api.GetName(deletePowerRobotSweepStrategy.EndPoint{}):     deletePowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(deleteProductionData.EndPoint{}):              deleteProductionData.Init(apiRoot),
			api.GetName(deletePs.EndPoint{}):                          deletePs.Init(apiRoot),
			api.GetName(deleteRechargeOrder.EndPoint{}):               deleteRechargeOrder.Init(apiRoot),
			api.GetName(deleteRegularlyConnectionInfo.EndPoint{}):     deleteRegularlyConnectionInfo.Init(apiRoot),
			api.GetName(deleteReportConfigEmailAddr.EndPoint{}):       deleteReportConfigEmailAddr.Init(apiRoot),
			api.GetName(deleteSysAdvancedParam.EndPoint{}):            deleteSysAdvancedParam.Init(apiRoot),
			api.GetName(deleteSysOrgNew.EndPoint{}):                   deleteSysOrgNew.Init(apiRoot),
			api.GetName(deleteUserInfoAllByUserId.EndPoint{}):         deleteUserInfoAllByUserId.Init(apiRoot),
			api.GetName(deviceInputDiscreteDeleteTime.EndPoint{}):     deviceInputDiscreteDeleteTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteGetTime.EndPoint{}):        deviceInputDiscreteGetTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteInsertTime.EndPoint{}):     deviceInputDiscreteInsertTime.Init(apiRoot),
			api.GetName(deviceInputDiscreteUpdateTime.EndPoint{}):     deviceInputDiscreteUpdateTime.Init(apiRoot),
			api.GetName(deviceReplace.EndPoint{}):                     deviceReplace.Init(apiRoot),
			api.GetName(editOperRule.EndPoint{}):                      editOperRule.Init(apiRoot),
			api.GetName(energyPovertyAlleviation.EndPoint{}):          energyPovertyAlleviation.Init(apiRoot),
			api.GetName(faultAutoClose.EndPoint{}):                    faultAutoClose.Init(apiRoot),
			api.GetName(faultCloseRemindOrderHandler.EndPoint{}):      faultCloseRemindOrderHandler.Init(apiRoot),
			api.GetName(findCodeValueList.EndPoint{}):                 findCodeValueList.Init(apiRoot),
			api.GetName(findEmgOrgInfo.EndPoint{}):                    findEmgOrgInfo.Init(apiRoot),
			api.GetName(findEnvironmentInfo.EndPoint{}):               findEnvironmentInfo.Init(apiRoot),
			api.GetName(findFromHbaseAndRedis.EndPoint{}):             findFromHbaseAndRedis.Init(apiRoot),
			api.GetName(findLossAnalysisList.EndPoint{}):              findLossAnalysisList.Init(apiRoot),
			api.GetName(findOnDutyInfo.EndPoint{}):                    findOnDutyInfo.Init(apiRoot),
			api.GetName(findSingleStationPR.EndPoint{}):               findSingleStationPR.Init(apiRoot),
			api.GetName(findUserPassword.EndPoint{}):                  findUserPassword.Init(apiRoot),
			api.GetName(genTLSUserSigByUserAccount.EndPoint{}):        genTLSUserSigByUserAccount.Init(apiRoot),
			api.GetName(generateRandomPassword.EndPoint{}):            generateRandomPassword.Init(apiRoot),
			api.GetName(getAPIServiceInfo.EndPoint{}):                 getAPIServiceInfo.Init(apiRoot),
			api.GetName(getAccessedPermission.EndPoint{}):             getAccessedPermission.Init(apiRoot),
			api.GetName(getAllPowerRobotViewInfoByPsId.EndPoint{}):    getAllPowerRobotViewInfoByPsId.Init(apiRoot),
			api.GetName(getAllPsIdByOrgIds.EndPoint{}):                getAllPsIdByOrgIds.Init(apiRoot),
			api.GetName(getAllUserRemindCount.EndPoint{}):             getAllUserRemindCount.Init(apiRoot),
			api.GetName(getAndOutletsAndUnit.EndPoint{}):              getAndOutletsAndUnit.Init(apiRoot),
			api.GetName(getAreaInfoCodeByCounty.EndPoint{}):           getAreaInfoCodeByCounty.Init(apiRoot),
			api.GetName(getAutoCreatePowerStation.EndPoint{}):         getAutoCreatePowerStation.Init(apiRoot),
			api.GetName(getBackReadValue.EndPoint{}):                  getBackReadValue.Init(apiRoot),
			api.GetName(getBatchNewestPointData.EndPoint{}):           getBatchNewestPointData.Init(apiRoot),
			api.GetName(getCallApiResidueTimes.EndPoint{}):            getCallApiResidueTimes.Init(apiRoot),
			api.GetName(getChangedPsListByTime.EndPoint{}):            getChangedPsListByTime.Init(apiRoot),
			api.GetName(getChnnlListByPsId.EndPoint{}):                getChnnlListByPsId.Init(apiRoot),
			api.GetName(getCloudServiceMappingConfig.EndPoint{}):      getCloudServiceMappingConfig.Init(apiRoot),
			api.GetName(getCommunicationModuleMonitorData.EndPoint{}): getCommunicationModuleMonitorData.Init(apiRoot),
			api.GetName(getComponentModelFactory.EndPoint{}):          getComponentModelFactory.Init(apiRoot),

			api.GetName(getConnectionInfoBySnAndLocalPort.EndPoint{}):                 getConnectionInfoBySnAndLocalPort.Init(apiRoot),
			api.GetName(getCountDown.EndPoint{}):                                      getCountDown.Init(apiRoot),
			api.GetName(getCountryServiceInfo.EndPoint{}):                             getCountryServiceInfo.Init(apiRoot),
			api.GetName(getCounty.EndPoint{}):                                         getCounty.Init(apiRoot),
			api.GetName(getCustomerEmployee.EndPoint{}):                               getCustomerEmployee.Init(apiRoot),
			api.GetName(getCustomerList.EndPoint{}):                                   getCustomerList.Init(apiRoot),
			api.GetName(getDevInstalledPowerByPsId.EndPoint{}):                        getDevInstalledPowerByPsId.Init(apiRoot),
			api.GetName(getDevRecord.EndPoint{}):                                      getDevRecord.Init(apiRoot),
			api.GetName(getDevRunRecordList.EndPoint{}):                               getDevRunRecordList.Init(apiRoot),
			api.GetName(getDevSimByList.EndPoint{}):                                   getDevSimByList.Init(apiRoot),
			api.GetName(getDevSimList.EndPoint{}):                                     getDevSimList.Init(apiRoot),
			api.GetName(getDstInfo.EndPoint{}):                                        getDstInfo.Init(apiRoot),
			api.GetName(getElectricitySettlementData.EndPoint{}):                      getElectricitySettlementData.Init(apiRoot),
			api.GetName(getElectricitySettlementDetailData.EndPoint{}):                getElectricitySettlementDetailData.Init(apiRoot),
			api.GetName(getEncryptPublicKey.EndPoint{}):                               getEncryptPublicKey.Init(apiRoot),
			api.GetName(getFaultCount.EndPoint{}):                                     getFaultCount.Init(apiRoot),
			api.GetName(getFaultDetail.EndPoint{}):                                    getFaultDetail.Init(apiRoot),
			api.GetName(getFaultMsgByFaultCode.EndPoint{}):                            getFaultMsgByFaultCode.Init(apiRoot),
			api.GetName(getFaultMsgListByPageWithYYYYMM.EndPoint{}):                   getFaultMsgListByPageWithYYYYMM.Init(apiRoot),
			api.GetName(getFaultMsgListWithYYYYMM.EndPoint{}):                         getFaultMsgListWithYYYYMM.Init(apiRoot),
			api.GetName(getFaultPlanList.EndPoint{}):                                  getFaultPlanList.Init(apiRoot),
			api.GetName(getFileOperationRecordOne.EndPoint{}):                         getFileOperationRecordOne.Init(apiRoot),
			api.GetName(getFormulaFaultAnalyseList.EndPoint{}):                        getFormulaFaultAnalyseList.Init(apiRoot),
			api.GetName(getGroupStringCheckResult.EndPoint{}):                         getGroupStringCheckResult.Init(apiRoot),
			api.GetName(getGroupStringCheckRule.EndPoint{}):                           getGroupStringCheckRule.Init(apiRoot),
			api.GetName(getHisData.EndPoint{}):                                        getHisData.Init(apiRoot),
			api.GetName(getHouseholdEvaluation.EndPoint{}):                            getHouseholdEvaluation.Init(apiRoot),
			api.GetName(getHouseholdLeaveMessage.EndPoint{}):                          getHouseholdLeaveMessage.Init(apiRoot),
			api.GetName(getHouseholdOpinionFeedback.EndPoint{}):                       getHouseholdOpinionFeedback.Init(apiRoot),
			api.GetName(getHouseholdPsInstallerByUserId.EndPoint{}):                   getHouseholdPsInstallerByUserId.Init(apiRoot),
			api.GetName(getHouseholdUserInfo.EndPoint{}):                              getHouseholdUserInfo.Init(apiRoot),
			api.GetName(getHouseholdWorkOrderInfo.EndPoint{}):                         getHouseholdWorkOrderInfo.Init(apiRoot),
			api.GetName(getHouseholdWorkOrderList.EndPoint{}):                         getHouseholdWorkOrderList.Init(apiRoot),
			api.GetName(getI18nConfigByType.EndPoint{}):                               getI18nConfigByType.Init(apiRoot),
			api.GetName(getI18nFileInfo.EndPoint{}):                                   getI18nFileInfo.Init(apiRoot),
			api.GetName(getI18nInfoByKey.EndPoint{}):                                  getI18nInfoByKey.Init(apiRoot),
			api.GetName(getI18nVersion.EndPoint{}):                                    getI18nVersion.Init(apiRoot),
			api.GetName(getIncomeSettingInfos.EndPoint{}):                             getIncomeSettingInfos.Init(apiRoot),
			api.GetName(getInfoFromAMap.EndPoint{}):                                   getInfoFromAMap.Init(apiRoot),
			api.GetName(getInfomationFromRedis.EndPoint{}):                            getInfomationFromRedis.Init(apiRoot),
			api.GetName(getInstallInfoList.EndPoint{}):                                getInstallInfoList.Init(apiRoot),
			api.GetName(getInstallerInfoByDealerOrgCodeOrId.EndPoint{}):               getInstallerInfoByDealerOrgCodeOrId.Init(apiRoot),
			api.GetName(getInverterDataCount.EndPoint{}):                              getInverterDataCount.Init(apiRoot),
			api.GetName(getInverterProcess.EndPoint{}):                                getInverterProcess.Init(apiRoot),
			api.GetName(getInverterUuidBytotalId.EndPoint{}):                          getInverterUuidBytotalId.Init(apiRoot),
			api.GetName(getIvEchartsData.EndPoint{}):                                  getIvEchartsData.Init(apiRoot),
			api.GetName(getIvEchartsDataById.EndPoint{}):                              getIvEchartsDataById.Init(apiRoot),
			api.GetName(getMenuAndPrivileges.EndPoint{}):                              getMenuAndPrivileges.Init(apiRoot),
			api.GetName(getMicrogridEStoragePsReport.EndPoint{}):                      getMicrogridEStoragePsReport.Init(apiRoot),
			api.GetName(getModuleLogInfo.EndPoint{}):                                  getModuleLogInfo.Init(apiRoot),
			api.GetName(getNationProvJSON.EndPoint{}):                                 getNationProvJSON.Init(apiRoot),
			api.GetName(getNeedOpAsynOpRecordList.EndPoint{}):                         getNeedOpAsynOpRecordList.Init(apiRoot),
			api.GetName(getNoticeInfo.EndPoint{}):                                     getNoticeInfo.Init(apiRoot),
			api.GetName(getNumberOfServiceCalls.EndPoint{}):                           getNumberOfServiceCalls.Init(apiRoot),
			api.GetName(getOSSConfig.EndPoint{}):                                      getOSSConfig.Init(apiRoot),
			api.GetName(getOperRuleDetail.EndPoint{}):                                 getOperRuleDetail.Init(apiRoot),
			api.GetName(getOperateBillFileId.EndPoint{}):                              getOperateBillFileId.Init(apiRoot),
			api.GetName(getOperateTicketForDetail.EndPoint{}):                         getOperateTicketForDetail.Init(apiRoot),
			api.GetName(getOrCreateNetEaseUserToken.EndPoint{}):                       getOrCreateNetEaseUserToken.Init(apiRoot),
			api.GetName(getOrderDataList.EndPoint{}):                                  getOrderDataList.Init(apiRoot),
			api.GetName(getOrderDataSql2.EndPoint{}):                                  getOrderDataSql2.Init(apiRoot),
			api.GetName(getOrderDatas.EndPoint{}):                                     getOrderDatas.Init(apiRoot),
			api.GetName(getOrderDetail.EndPoint{}):                                    getOrderDetail.Init(apiRoot),
			api.GetName(getOrderStatistics.EndPoint{}):                                getOrderStatistics.Init(apiRoot),
			api.GetName(getOrgIdNameByUserId.EndPoint{}):                              getOrgIdNameByUserId.Init(apiRoot),
			api.GetName(getOrgInfoByDealerOrgCode.EndPoint{}):                         getOrgInfoByDealerOrgCode.Init(apiRoot),
			api.GetName(getOrgListByUserId.EndPoint{}):                                getOrgListByUserId.Init(apiRoot),
			api.GetName(getOrgListForUser.EndPoint{}):                                 getOrgListForUser.Init(apiRoot),
			api.GetName(getOssObjectStream.EndPoint{}):                                getOssObjectStream.Init(apiRoot),
			api.GetName(getOwnerFaultConfigList.EndPoint{}):                           getOwnerFaultConfigList.Init(apiRoot),
			api.GetName(getParamterSettingBase.EndPoint{}):                            getParamterSettingBase.Init(apiRoot),
			api.GetName(getPhotoInfo.EndPoint{}):                                      getPhotoInfo.Init(apiRoot),
			api.GetName(getPlanedOrNotPsList.EndPoint{}):                              getPlanedOrNotPsList.Init(apiRoot),
			api.GetName(getPlantReportPDFList.EndPoint{}):                             getPlantReportPDFList.Init(apiRoot),
			api.GetName(getPowerChargeSettingInfo.EndPoint{}):                         getPowerChargeSettingInfo.Init(apiRoot),
			api.GetName(getPowerFormulaFaultAnalyse.EndPoint{}):                       getPowerFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(getPowerRobotInfoByRobotSn.EndPoint{}):                        getPowerRobotInfoByRobotSn.Init(apiRoot),
			api.GetName(getPowerRobotSweepAttrByPsId.EndPoint{}):                      getPowerRobotSweepAttrByPsId.Init(apiRoot),
			api.GetName(getPowerRobotSweepStrategy.EndPoint{}):                        getPowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(getPowerRobotSweepStrategyList.EndPoint{}):                    getPowerRobotSweepStrategyList.Init(apiRoot),
			api.GetName(getPowerSettingCharges.EndPoint{}):                            getPowerSettingCharges.Init(apiRoot),
			api.GetName(getPowerSettingHistoryRecords.EndPoint{}):                     getPowerSettingHistoryRecords.Init(apiRoot),
			api.GetName(getPowerStationBasicInfo.EndPoint{}):                          getPowerStationBasicInfo.Init(apiRoot),
			api.GetName(getPowerStationData.EndPoint{}):                               getPowerStationData.Init(apiRoot),
			api.GetName(getPowerStationForHousehold.EndPoint{}):                       getPowerStationForHousehold.Init(apiRoot),
			api.GetName(getPowerStationInfo.EndPoint{}):                               getPowerStationInfo.Init(apiRoot),
			api.GetName(getPowerStationPR.EndPoint{}):                                 getPowerStationPR.Init(apiRoot),
			api.GetName(getPrivateCloudValidityPeriod.EndPoint{}):                     getPrivateCloudValidityPeriod.Init(apiRoot),
			api.GetName(getProvInfoListByNationCode.EndPoint{}):                       getProvInfoListByNationCode.Init(apiRoot),
			api.GetName(getPsDetailByUserTokens.EndPoint{}):                           getPsDetailByUserTokens.Init(apiRoot),
			api.GetName(getPsDetailForSinglePage.EndPoint{}):                          getPsDetailForSinglePage.Init(apiRoot),
			api.GetName(getPsInstallerByPsId.EndPoint{}):                              getPsInstallerByPsId.Init(apiRoot),
			api.GetName(getPsInstallerOrgInfoByPsId.EndPoint{}):                       getPsInstallerOrgInfoByPsId.Init(apiRoot),
			api.GetName(getPsListForPsDataByPsId.EndPoint{}):                          getPsListForPsDataByPsId.Init(apiRoot),
			api.GetName(getRechargeOrderDetail.EndPoint{}):                            getRechargeOrderDetail.Init(apiRoot),
			api.GetName(getRechargeOrderList.EndPoint{}):                              getRechargeOrderList.Init(apiRoot),
			api.GetName(getRegionalTree.EndPoint{}):                                   getRegionalTree.Init(apiRoot),
			api.GetName(getRemoteParamSettingList.EndPoint{}):                         getRemoteParamSettingList.Init(apiRoot),
			api.GetName(getRemoteUpgradeScheduleDetails.EndPoint{}):                   getRemoteUpgradeScheduleDetails.Init(apiRoot),
			api.GetName(getReportData.EndPoint{}):                                     getReportData.Init(apiRoot),
			api.GetName(getReportEmailConfigInfo.EndPoint{}):                          getReportEmailConfigInfo.Init(apiRoot),
			api.GetName(getReportExportColumns.EndPoint{}):                            getReportExportColumns.Init(apiRoot),
			api.GetName(getReportListByUserId.EndPoint{}):                             getReportListByUserId.Init(apiRoot),
			api.GetName(getRobotDynamicCleaningView.EndPoint{}):                       getRobotDynamicCleaningView.Init(apiRoot),
			api.GetName(getRobotNumAndSweepCapacity.EndPoint{}):                       getRobotNumAndSweepCapacity.Init(apiRoot),
			api.GetName(getRuleUnit.EndPoint{}):                                       getRuleUnit.Init(apiRoot),
			api.GetName(getSendReportConfigCron.EndPoint{}):                           getSendReportConfigCron.Init(apiRoot),
			api.GetName(getShieldMapConditionList.EndPoint{}):                         getShieldMapConditionList.Init(apiRoot),
			api.GetName(getSimIdBySnList.EndPoint{}):                                  getSimIdBySnList.Init(apiRoot),
			api.GetName(getSingleIVData.EndPoint{}):                                   getSingleIVData.Init(apiRoot),
			api.GetName(getSnChangeRecord.EndPoint{}):                                 getSnChangeRecord.Init(apiRoot),
			api.GetName(getSnConnectionInfo.EndPoint{}):                               getSnConnectionInfo.Init(apiRoot),
			api.GetName(getSungwsConfigCache.EndPoint{}):                              getSungwsConfigCache.Init(apiRoot),
			api.GetName(getSungwsGlobalConfigCache.EndPoint{}):                        getSungwsGlobalConfigCache.Init(apiRoot),
			api.GetName(getSweepRobotDevList.EndPoint{}):                              getSweepRobotDevList.Init(apiRoot),
			api.GetName(getSysMsg.EndPoint{}):                                         getSysMsg.Init(apiRoot),
			api.GetName(getSysOrgNewList.EndPoint{}):                                  getSysOrgNewList.Init(apiRoot),
			api.GetName(getSysOrgNewOne.EndPoint{}):                                   getSysOrgNewOne.Init(apiRoot),
			api.GetName(getSysUserById.EndPoint{}):                                    getSysUserById.Init(apiRoot),
			api.GetName(getUUIDByUpuuid.EndPoint{}):                                   getUUIDByUpuuid.Init(apiRoot),
			api.GetName(getUserById.EndPoint{}):                                       getUserById.Init(apiRoot),
			api.GetName(getUserByInstaller.EndPoint{}):                                getUserByInstaller.Init(apiRoot),
			api.GetName(getUserDevOnlineOffineCount.EndPoint{}):                       getUserDevOnlineOffineCount.Init(apiRoot),
			api.GetName(getUserGDPRAttrs.EndPoint{}):                                  getUserGDPRAttrs.Init(apiRoot),
			api.GetName(getUserHavePowerStationCount.EndPoint{}):                      getUserHavePowerStationCount.Init(apiRoot),
			api.GetName(getUserInfoByUserAccounts.EndPoint{}):                         getUserInfoByUserAccounts.Init(apiRoot),
			api.GetName(getValidateCode.EndPoint{}):                                   getValidateCode.Init(apiRoot),
			api.GetName(getValidateCodeAtRegister.EndPoint{}):                         getValidateCodeAtRegister.Init(apiRoot),
			api.GetName(getWeatherInfo.EndPoint{}):                                    getWeatherInfo.Init(apiRoot),
			api.GetName(getWechatPushConfig.EndPoint{}):                               getWechatPushConfig.Init(apiRoot),
			api.GetName(getWorkInfo.EndPoint{}):                                       getWorkInfo.Init(apiRoot),
			api.GetName(groupStringCheck.EndPoint{}):                                  groupStringCheck.Init(apiRoot),
			api.GetName(handleDevByCommunicationSN.EndPoint{}):                        handleDevByCommunicationSN.Init(apiRoot),
			api.GetName(householdResetPassBySN.EndPoint{}):                            householdResetPassBySN.Init(apiRoot),
			api.GetName(immediatePayment.EndPoint{}):                                  immediatePayment.Init(apiRoot),
			api.GetName(importExcelData.EndPoint{}):                                   importExcelData.Init(apiRoot),
			api.GetName(incomeStatistics.EndPoint{}):                                  incomeStatistics.Init(apiRoot),
			api.GetName(informPush.EndPoint{}):                                        informPush.Init(apiRoot),
			api.GetName(insertEmgOrgInfo.EndPoint{}):                                  insertEmgOrgInfo.Init(apiRoot),
			api.GetName(intoDataToHbase.EndPoint{}):                                   intoDataToHbase.Init(apiRoot),
			api.GetName(ipLocationQuery.EndPoint{}):                                   ipLocationQuery.Init(apiRoot),
			api.GetName(isHave2GSn.EndPoint{}):                                        isHave2GSn.Init(apiRoot),
			api.GetName(judgeIsSettingMan.EndPoint{}):                                 judgeIsSettingMan.Init(apiRoot),
			api.GetName(listOssFiles.EndPoint{}):                                      listOssFiles.Init(apiRoot),
			api.GetName(loadAreaInfo.EndPoint{}):                                      loadAreaInfo.Init(apiRoot),
			api.GetName(loadPowerStation.EndPoint{}):                                  loadPowerStation.Init(apiRoot),
			api.GetName(loginByToken.EndPoint{}):                                      loginByToken.Init(apiRoot),
			api.GetName(logout.EndPoint{}):                                            logout.Init(apiRoot),
			api.GetName(mobilePhoneHasBound.EndPoint{}):                               mobilePhoneHasBound.Init(apiRoot),
			api.GetName(modifyEmgOrgStruc.EndPoint{}):                                 modifyEmgOrgStruc.Init(apiRoot),
			api.GetName(modifyFaultPlan.EndPoint{}):                                   modifyFaultPlan.Init(apiRoot),
			api.GetName(modifyOnDutyInfo.EndPoint{}):                                  modifyOnDutyInfo.Init(apiRoot),
			api.GetName(modifyPassword.EndPoint{}):                                    modifyPassword.Init(apiRoot),
			api.GetName(modifyPersonalUnitList.EndPoint{}):                            modifyPersonalUnitList.Init(apiRoot),
			api.GetName(modifyPsUser.EndPoint{}):                                      modifyPsUser.Init(apiRoot),
			api.GetName(moduleLogParamSet.EndPoint{}):                                 moduleLogParamSet.Init(apiRoot),
			api.GetName(operateOssFile.EndPoint{}):                                    operateOssFile.Init(apiRoot),
			api.GetName(operationPowerRobotSweepStrategy.EndPoint{}):                  operationPowerRobotSweepStrategy.Init(apiRoot),
			api.GetName(orgPowerReport.EndPoint{}):                                    orgPowerReport.Init(apiRoot),
			api.GetName(paramSetTryAgain.EndPoint{}):                                  paramSetTryAgain.Init(apiRoot),
			api.GetName(paramSetting.EndPoint{}):                                      paramSetting.Init(apiRoot),
			api.GetName(planPower.EndPoint{}):                                         planPower.Init(apiRoot),
			api.GetName(psForcastInfo.EndPoint{}):                                     psForcastInfo.Init(apiRoot),
			api.GetName(queryBatchSpeedyAddPowerStationResult.EndPoint{}):             queryBatchSpeedyAddPowerStationResult.Init(apiRoot),
			api.GetName(queryCardStatusCTCC.EndPoint{}):                               queryCardStatusCTCC.Init(apiRoot),
			api.GetName(queryChildAccountList.EndPoint{}):                             queryChildAccountList.Init(apiRoot),
			api.GetName(queryCompensationRecordData.EndPoint{}):                       queryCompensationRecordData.Init(apiRoot),
			api.GetName(queryCompensationRecordList.EndPoint{}):                       queryCompensationRecordList.Init(apiRoot),
			api.GetName(queryComponent.EndPoint{}):                                    queryComponent.Init(apiRoot),
			api.GetName(queryComponentTechnicalParam.EndPoint{}):                      queryComponentTechnicalParam.Init(apiRoot),
			api.GetName(queryCountryGridAndRelation.EndPoint{}):                       queryCountryGridAndRelation.Init(apiRoot),
			api.GetName(queryCountryList.EndPoint{}):                                  queryCountryList.Init(apiRoot),
			api.GetName(queryEnvironmentList.EndPoint{}):                              queryEnvironmentList.Init(apiRoot),
			api.GetName(queryFaultList.EndPoint{}):                                    queryFaultList.Init(apiRoot),
			api.GetName(queryFaultPlanDetail.EndPoint{}):                              queryFaultPlanDetail.Init(apiRoot),
			api.GetName(queryFaultRepairSteps.EndPoint{}):                             queryFaultRepairSteps.Init(apiRoot),
			api.GetName(queryFaultTypeAndLevelByCode.EndPoint{}):                      queryFaultTypeAndLevelByCode.Init(apiRoot),
			api.GetName(queryFirmwareFilesPage.EndPoint{}):                            queryFirmwareFilesPage.Init(apiRoot),
			api.GetName(queryInfotoAlert.EndPoint{}):                                  queryInfotoAlert.Init(apiRoot),
			api.GetName(queryInverterModelList.EndPoint{}):                            queryInverterModelList.Init(apiRoot),
			api.GetName(queryInverterVersionList.EndPoint{}):                          queryInverterVersionList.Init(apiRoot),
			api.GetName(queryM2MCardInfoCMCC.EndPoint{}):                              queryM2MCardInfoCMCC.Init(apiRoot),
			api.GetName(queryM2MCardTermInfoCMCC.EndPoint{}):                          queryM2MCardTermInfoCMCC.Init(apiRoot),
			api.GetName(queryModelInfoByModelId.EndPoint{}):                           queryModelInfoByModelId.Init(apiRoot),
			api.GetName(queryNoticeList.EndPoint{}):                                   queryNoticeList.Init(apiRoot),
			api.GetName(queryNumberOfRenewalReminders.EndPoint{}):                     queryNumberOfRenewalReminders.Init(apiRoot),
			api.GetName(queryOperRules.EndPoint{}):                                    queryOperRules.Init(apiRoot),
			api.GetName(queryOrderList.EndPoint{}):                                    queryOrderList.Init(apiRoot),
			api.GetName(queryOrderStep.EndPoint{}):                                    queryOrderStep.Init(apiRoot),
			api.GetName(queryOrgGenerationReport.EndPoint{}):                          queryOrgGenerationReport.Init(apiRoot),
			api.GetName(queryOrgInfoList.EndPoint{}):                                  queryOrgInfoList.Init(apiRoot),
			api.GetName(queryOrgPowerElecPercent.EndPoint{}):                          queryOrgPowerElecPercent.Init(apiRoot),
			api.GetName(queryOrgPsCompensationRecordList.EndPoint{}):                  queryOrgPsCompensationRecordList.Init(apiRoot),
			api.GetName(queryPersonalUnitList.EndPoint{}):                             queryPersonalUnitList.Init(apiRoot),
			api.GetName(queryPointDataTopOne.EndPoint{}):                              queryPointDataTopOne.Init(apiRoot),
			api.GetName(queryPowerStationInfo.EndPoint{}):                             queryPowerStationInfo.Init(apiRoot),
			api.GetName(queryPsAreaByUserIdAndAreaCode.EndPoint{}):                    queryPsAreaByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(queryPsCompensationRecordList.EndPoint{}):                     queryPsCompensationRecordList.Init(apiRoot),
			api.GetName(queryPsDataByDate.EndPoint{}):                                 queryPsDataByDate.Init(apiRoot),
			api.GetName(queryPsIdList.EndPoint{}):                                     queryPsIdList.Init(apiRoot),
			api.GetName(queryPsListByUserIdAndAreaCode.EndPoint{}):                    queryPsListByUserIdAndAreaCode.Init(apiRoot),
			api.GetName(queryPsNameByPsId.EndPoint{}):                                 queryPsNameByPsId.Init(apiRoot),
			api.GetName(queryPsPrByDate.EndPoint{}):                                   queryPsPrByDate.Init(apiRoot),
			api.GetName(queryPsProfit.EndPoint{}):                                     queryPsProfit.Init(apiRoot),
			api.GetName(queryPsReportComparativeAnalysisOfPowerGeneration.EndPoint{}): queryPsReportComparativeAnalysisOfPowerGeneration.Init(apiRoot),
			api.GetName(queryPsStructureList.EndPoint{}):                              queryPsStructureList.Init(apiRoot),
			api.GetName(queryPuuidsByCommandTotalId.EndPoint{}):                       queryPuuidsByCommandTotalId.Init(apiRoot),
			api.GetName(queryPuuidsByCommandTotalId2.EndPoint{}):                      queryPuuidsByCommandTotalId2.Init(apiRoot),
			api.GetName(queryRepairRuleList.EndPoint{}):                               queryRepairRuleList.Init(apiRoot),
			api.GetName(queryReportListForManagementPage.EndPoint{}):                  queryReportListForManagementPage.Init(apiRoot),
			api.GetName(queryReportMsg.EndPoint{}):                                    queryReportMsg.Init(apiRoot),
			api.GetName(querySharingPs.EndPoint{}):                                    querySharingPs.Init(apiRoot),
			api.GetName(querySysAdvancedParam.EndPoint{}):                             querySysAdvancedParam.Init(apiRoot),
			api.GetName(queryTimeBySN.EndPoint{}):                                     queryTimeBySN.Init(apiRoot),
			api.GetName(queryTrafficByDateCTCC.EndPoint{}):                            queryTrafficByDateCTCC.Init(apiRoot),
			api.GetName(queryTrafficCTCC.EndPoint{}):                                  queryTrafficCTCC.Init(apiRoot),
			api.GetName(queryUnitUuidBytotalId.EndPoint{}):                            queryUnitUuidBytotalId.Init(apiRoot),
			api.GetName(queryUserBtnPri.EndPoint{}):                                   queryUserBtnPri.Init(apiRoot),
			api.GetName(queryUserByUserIds.EndPoint{}):                                queryUserByUserIds.Init(apiRoot),
			api.GetName(queryUserExtensionAttribute.EndPoint{}):                       queryUserExtensionAttribute.Init(apiRoot),
			api.GetName(queryUserForStep.EndPoint{}):                                  queryUserForStep.Init(apiRoot),
			api.GetName(queryUserProcessPri.EndPoint{}):                               queryUserProcessPri.Init(apiRoot),
			api.GetName(queryUserWechatBindRel.EndPoint{}):                            queryUserWechatBindRel.Init(apiRoot),
			api.GetName(queryUuidByTotalIdAndUuid.EndPoint{}):                         queryUuidByTotalIdAndUuid.Init(apiRoot),
			api.GetName(rechargeOrderSetMeal.EndPoint{}):                              rechargeOrderSetMeal.Init(apiRoot),
			api.GetName(renewSendReportConfirmEmail.EndPoint{}):                       renewSendReportConfirmEmail.Init(apiRoot),
			api.GetName(saveCustomerEmployee.EndPoint{}):                              saveCustomerEmployee.Init(apiRoot),
			api.GetName(saveDevSimList.EndPoint{}):                                    saveDevSimList.Init(apiRoot),
			api.GetName(saveEnviromentIncomeInfos.EndPoint{}):                         saveEnviromentIncomeInfos.Init(apiRoot),
			api.GetName(saveEnvironmentCurve.EndPoint{}):                              saveEnvironmentCurve.Init(apiRoot),
			api.GetName(saveFirmwareFile.EndPoint{}):                                  saveFirmwareFile.Init(apiRoot),
			api.GetName(saveIncomeSettingInfos.EndPoint{}):                            saveIncomeSettingInfos.Init(apiRoot),
			api.GetName(saveOrUpdateGroupStringCheckRule.EndPoint{}):                  saveOrUpdateGroupStringCheckRule.Init(apiRoot),
			api.GetName(saveParamModel.EndPoint{}):                                    saveParamModel.Init(apiRoot),
			api.GetName(savePowerCharges.EndPoint{}):                                  savePowerCharges.Init(apiRoot),
			api.GetName(savePowerRobotInfo.EndPoint{}):                                savePowerRobotInfo.Init(apiRoot),
			api.GetName(savePowerRobotSweepAttr.EndPoint{}):                           savePowerRobotSweepAttr.Init(apiRoot),
			api.GetName(savePowerSettingCharges.EndPoint{}):                           savePowerSettingCharges.Init(apiRoot),
			api.GetName(savePowerSettingInfo.EndPoint{}):                              savePowerSettingInfo.Init(apiRoot),
			api.GetName(saveProductionBatchData.EndPoint{}):                           saveProductionBatchData.Init(apiRoot),
			api.GetName(saveRechargeOrderObj.EndPoint{}):                              saveRechargeOrderObj.Init(apiRoot),
			api.GetName(saveRechargeOrderOtherInfo.EndPoint{}):                        saveRechargeOrderOtherInfo.Init(apiRoot),
			api.GetName(saveRepair.EndPoint{}):                                        saveRepair.Init(apiRoot),
			api.GetName(saveReportExportColumns.EndPoint{}):                           saveReportExportColumns.Init(apiRoot),
			api.GetName(saveSetParam.EndPoint{}):                                      saveSetParam.Init(apiRoot),
			api.GetName(saveSysUserMsg.EndPoint{}):                                    saveSysUserMsg.Init(apiRoot),
			api.GetName(searchM2MMonthFlowCMCC.EndPoint{}):                            searchM2MMonthFlowCMCC.Init(apiRoot),
			api.GetName(selectSysTranslationNames.EndPoint{}):                         selectSysTranslationNames.Init(apiRoot),
			api.GetName(sendPsTimeZoneInstruction.EndPoint{}):                         sendPsTimeZoneInstruction.Init(apiRoot),
			api.GetName(setUpFormulaFaultAnalyse.EndPoint{}):                          setUpFormulaFaultAnalyse.Init(apiRoot),
			api.GetName(setUserGDPRAttrs.EndPoint{}):                                  setUserGDPRAttrs.Init(apiRoot),
			api.GetName(settingNotice.EndPoint{}):                                     settingNotice.Init(apiRoot),
			api.GetName(shareMyPs.EndPoint{}):                                         shareMyPs.Init(apiRoot),
			api.GetName(sharePsBySN.EndPoint{}):                                       sharePsBySN.Init(apiRoot),
			api.GetName(showInverterByUnit.EndPoint{}):                                showInverterByUnit.Init(apiRoot),
			api.GetName(showOnlineUsers.EndPoint{}):                                   showOnlineUsers.Init(apiRoot),
			api.GetName(showWarning.EndPoint{}):                                       showWarning.Init(apiRoot),
			api.GetName(snIsExist.EndPoint{}):                                         snIsExist.Init(apiRoot),
			api.GetName(snsIsExist.EndPoint{}):                                        snsIsExist.Init(apiRoot),
			api.GetName(speedyAddPowerStation.EndPoint{}):                             speedyAddPowerStation.Init(apiRoot),
			api.GetName(stationUnitsList.EndPoint{}):                                  stationUnitsList.Init(apiRoot),
			api.GetName(stationsDiscreteData.EndPoint{}):                              stationsDiscreteData.Init(apiRoot),
			api.GetName(stationsIncomeList.EndPoint{}):                                stationsIncomeList.Init(apiRoot),
			api.GetName(stationsPointReport.EndPoint{}):                               stationsPointReport.Init(apiRoot),
			api.GetName(stationsYearPlanReport.EndPoint{}):                            stationsYearPlanReport.Init(apiRoot),
			api.GetName(sureAndImportSelettlementData.EndPoint{}):                     sureAndImportSelettlementData.Init(apiRoot),
			api.GetName(sweepDevParamSet.EndPoint{}):                                  sweepDevParamSet.Init(apiRoot),
			api.GetName(sweepDevRunControl.EndPoint{}):                                sweepDevRunControl.Init(apiRoot),
			api.GetName(sweepDevStrategyIssue.EndPoint{}):                             sweepDevStrategyIssue.Init(apiRoot),
			api.GetName(sysTimeZoneList.EndPoint{}):                                   sysTimeZoneList.Init(apiRoot),
			api.GetName(unLockUser.EndPoint{}):                                        unLockUser.Init(apiRoot),
			api.GetName(unlockChildAccount.EndPoint{}):                                unlockChildAccount.Init(apiRoot),
			api.GetName(updateCommunicationModuleState.EndPoint{}):                    updateCommunicationModuleState.Init(apiRoot),
			api.GetName(updateDevInstalledPower.EndPoint{}):                           updateDevInstalledPower.Init(apiRoot),
			api.GetName(updateFault.EndPoint{}):                                       updateFault.Init(apiRoot),
			api.GetName(updateFaultData.EndPoint{}):                                   updateFaultData.Init(apiRoot),
			api.GetName(updateFaultMsgByFaultCode.EndPoint{}):                         updateFaultMsgByFaultCode.Init(apiRoot),
			api.GetName(updateFaultStatus.EndPoint{}):                                 updateFaultStatus.Init(apiRoot),
			api.GetName(updateHouseholdWorkOrder.EndPoint{}):                          updateHouseholdWorkOrder.Init(apiRoot),
			api.GetName(updateInverterSn2ModuleSn.EndPoint{}):                         updateInverterSn2ModuleSn.Init(apiRoot),
			api.GetName(updateOperateTicketAttachmentId.EndPoint{}):                   updateOperateTicketAttachmentId.Init(apiRoot),
			api.GetName(updateOwnerFaultConfig.EndPoint{}):                            updateOwnerFaultConfig.Init(apiRoot),
			api.GetName(updateParamSettingSysMsg.EndPoint{}):                          updateParamSettingSysMsg.Init(apiRoot),
			api.GetName(updatePlatformLevelFaultLevel.EndPoint{}):                     updatePlatformLevelFaultLevel.Init(apiRoot),
			api.GetName(updatePowerRobotInfo.EndPoint{}):                              updatePowerRobotInfo.Init(apiRoot),
			api.GetName(updatePowerRobotSweepAttr.EndPoint{}):                         updatePowerRobotSweepAttr.Init(apiRoot),
			api.GetName(updatePowerStationForHousehold.EndPoint{}):                    updatePowerStationForHousehold.Init(apiRoot),
			api.GetName(updatePowerStationInfo.EndPoint{}):                            updatePowerStationInfo.Init(apiRoot),
			api.GetName(updatePowerUserInfo.EndPoint{}):                               updatePowerUserInfo.Init(apiRoot),
			api.GetName(updateReportConfigByEmailAddr.EndPoint{}):                     updateReportConfigByEmailAddr.Init(apiRoot),
			api.GetName(updateShareAttr.EndPoint{}):                                   updateShareAttr.Init(apiRoot),
			api.GetName(updateSnIsSureFlag.EndPoint{}):                                updateSnIsSureFlag.Init(apiRoot),
			api.GetName(updateStationPics.EndPoint{}):                                 updateStationPics.Init(apiRoot),
			api.GetName(updateSysAdvancedParam.EndPoint{}):                            updateSysAdvancedParam.Init(apiRoot),
			api.GetName(updateSysOrgNew.EndPoint{}):                                   updateSysOrgNew.Init(apiRoot),
			api.GetName(updateUinfoNetEaseUser.EndPoint{}):                            updateUinfoNetEaseUser.Init(apiRoot),
			api.GetName(updateUserExtensionAttribute.EndPoint{}):                      updateUserExtensionAttribute.Init(apiRoot),
			api.GetName(updateUserLanguage.EndPoint{}):                                updateUserLanguage.Init(apiRoot),
			api.GetName(updateUserPosition.EndPoint{}):                                updateUserPosition.Init(apiRoot),
			api.GetName(updateUserUpOrg.EndPoint{}):                                   updateUserUpOrg.Init(apiRoot),
			api.GetName(upgrade.EndPoint{}):                                           upgrade.Init(apiRoot),
			api.GetName(upgrate.EndPoint{}):                                           upgrate.Init(apiRoot),
			api.GetName(uploadFileToOss.EndPoint{}):                                   uploadFileToOss.Init(apiRoot),
			api.GetName(userAgreeGdprProtocol.EndPoint{}):                             userAgreeGdprProtocol.Init(apiRoot),
			api.GetName(userInfoUniqueCheck.EndPoint{}):                               userInfoUniqueCheck.Init(apiRoot),
			api.GetName(userMailHasBound.EndPoint{}):                                  userMailHasBound.Init(apiRoot),
			api.GetName(userRegister.EndPoint{}):                                      userRegister.Init(apiRoot),
		},
	}

	return area
}

// ****************************************
// Methods not scoped by api.EndPoint interface type

func GetAreaName() string {
	return string(api.GetArea(Area{}))
}

func (a Area) GetEndPoint(name api.EndPointName) api.EndPoint {
	var ret api.EndPoint
	for _, e := range a.EndPoints {
		// fmt.Printf("endpoint: %v\n", e)
		if e.GetName() == name {
			ret = e
			break
		}
	}
	return ret
}

// ****************************************
// Methods scoped by api.Area interface type

func (a Area) Init(apiRoot *api.Web) api.AreaStruct {
	panic("implement me")
}

func (a Area) GetAreaName() api.AreaName {
	return a.Name
}

func (a Area) GetEndPoints() api.TypeEndPoints {
	for _, endpoint := range a.EndPoints {
		fmt.Printf("endpoint: %v\n", endpoint)
	}
	return a.EndPoints
}

func (a Area) Call(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) SetRequest(name api.EndPointName, ref interface{}) error {
	panic("implement me")
}

func (a Area) GetRequest(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) GetResponse(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) GetData(name api.EndPointName) output.Json {
	panic("implement me")
}

func (a Area) IsValid(name api.EndPointName) error {
	panic("implement me")
}

func (a Area) GetError(name api.EndPointName) error {
	panic("implement me")
}
