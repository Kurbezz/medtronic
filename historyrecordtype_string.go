// Code generated by "stringer -type HistoryRecordType"; DO NOT EDIT.

package medtronic

import "strconv"

const _HistoryRecordType_name = "BolusPrimeAlarmDailyTotalBasalProfileBeforeBasalProfileAfterBGCaptureSensorAlarmClearAlarmChangeBasalPatternTempBasalDurationChangeTimeNewTimeLowBatteryBatteryChangeSetAutoOffSuspendPumpResumePumpSelfTestRewindClearSettingsEnableChildBlockMaxBolusEnableRemoteMaxBasalEnableBolusWizardUnknown2EBolusWizard512UnabsorbedInsulin512ChangeBGReminderSetAlarmClockTimeTempBasalRateLowReservoirAlarmClockChangeMeterIDSensorStatusEnableMeterBGReceivedMealMarkerExerciseMarkerInsulinMarkerOtherMarkerChangeBolusWizardSetupSensorSetupSensor51ChangeGlucoseUnitsBolusWizardSetupBolusWizardUnabsorbedInsulinSaveSettingsEnableVariableBolusChangeEasyBolusEnableBGReminderEnableAlarmClockChangeTempBasalTypeChangeAlarmTypeChangeTimeFormatChangeReservoirWarningEnableBolusReminderSetBolusReminderTimeDeleteBolusReminderTimeBolusReminderDeleteAlarmClockTimeDailyTotal515DailyTotal522DailyTotal523ChangeCarbUnitsBasalProfileStartConnectOtherDevicesChangeOtherDeviceChangeMarriageDeleteOtherDeviceEnableCaptureEvent"

var _HistoryRecordType_map = map[HistoryRecordType]string{
	1:   _HistoryRecordType_name[0:5],
	3:   _HistoryRecordType_name[5:10],
	6:   _HistoryRecordType_name[10:15],
	7:   _HistoryRecordType_name[15:25],
	8:   _HistoryRecordType_name[25:43],
	9:   _HistoryRecordType_name[43:60],
	10:  _HistoryRecordType_name[60:69],
	11:  _HistoryRecordType_name[69:80],
	12:  _HistoryRecordType_name[80:90],
	20:  _HistoryRecordType_name[90:108],
	22:  _HistoryRecordType_name[108:125],
	23:  _HistoryRecordType_name[125:135],
	24:  _HistoryRecordType_name[135:142],
	25:  _HistoryRecordType_name[142:152],
	26:  _HistoryRecordType_name[152:165],
	27:  _HistoryRecordType_name[165:175],
	30:  _HistoryRecordType_name[175:186],
	31:  _HistoryRecordType_name[186:196],
	32:  _HistoryRecordType_name[196:204],
	33:  _HistoryRecordType_name[204:210],
	34:  _HistoryRecordType_name[210:223],
	35:  _HistoryRecordType_name[223:239],
	36:  _HistoryRecordType_name[239:247],
	38:  _HistoryRecordType_name[247:259],
	44:  _HistoryRecordType_name[259:267],
	45:  _HistoryRecordType_name[267:284],
	46:  _HistoryRecordType_name[284:293],
	47:  _HistoryRecordType_name[293:307],
	48:  _HistoryRecordType_name[307:327],
	49:  _HistoryRecordType_name[327:343],
	50:  _HistoryRecordType_name[343:360],
	51:  _HistoryRecordType_name[360:373],
	52:  _HistoryRecordType_name[373:385],
	53:  _HistoryRecordType_name[385:395],
	54:  _HistoryRecordType_name[395:408],
	59:  _HistoryRecordType_name[408:420],
	60:  _HistoryRecordType_name[420:431],
	63:  _HistoryRecordType_name[431:441],
	64:  _HistoryRecordType_name[441:451],
	65:  _HistoryRecordType_name[451:465],
	66:  _HistoryRecordType_name[465:478],
	67:  _HistoryRecordType_name[478:489],
	79:  _HistoryRecordType_name[489:511],
	80:  _HistoryRecordType_name[511:522],
	81:  _HistoryRecordType_name[522:530],
	86:  _HistoryRecordType_name[530:548],
	90:  _HistoryRecordType_name[548:564],
	91:  _HistoryRecordType_name[564:575],
	92:  _HistoryRecordType_name[575:592],
	93:  _HistoryRecordType_name[592:604],
	94:  _HistoryRecordType_name[604:623],
	95:  _HistoryRecordType_name[623:638],
	96:  _HistoryRecordType_name[638:654],
	97:  _HistoryRecordType_name[654:670],
	98:  _HistoryRecordType_name[670:689],
	99:  _HistoryRecordType_name[689:704],
	100: _HistoryRecordType_name[704:720],
	101: _HistoryRecordType_name[720:742],
	102: _HistoryRecordType_name[742:761],
	103: _HistoryRecordType_name[761:781],
	104: _HistoryRecordType_name[781:804],
	105: _HistoryRecordType_name[804:817],
	106: _HistoryRecordType_name[817:837],
	108: _HistoryRecordType_name[837:850],
	109: _HistoryRecordType_name[850:863],
	110: _HistoryRecordType_name[863:876],
	111: _HistoryRecordType_name[876:891],
	123: _HistoryRecordType_name[891:908],
	124: _HistoryRecordType_name[908:927],
	125: _HistoryRecordType_name[927:944],
	129: _HistoryRecordType_name[944:958],
	130: _HistoryRecordType_name[958:975],
	131: _HistoryRecordType_name[975:993],
}

func (i HistoryRecordType) String() string {
	if str, ok := _HistoryRecordType_map[i]; ok {
		return str
	}
	return "HistoryRecordType(" + strconv.FormatInt(int64(i), 10) + ")"
}
