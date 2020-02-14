// Code generated by "stringer -type HistoryRecordType"; DO NOT EDIT.

package medtronic

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Bolus-1]
	_ = x[Prime-3]
	_ = x[Alarm-6]
	_ = x[DailyTotal-7]
	_ = x[BasalProfileBefore-8]
	_ = x[BasalProfileAfter-9]
	_ = x[BGCapture-10]
	_ = x[SensorAlarm-11]
	_ = x[ClearAlarm-12]
	_ = x[ChangeBasalPattern-20]
	_ = x[TempBasalDuration-22]
	_ = x[ChangeTime-23]
	_ = x[NewTime-24]
	_ = x[LowBattery-25]
	_ = x[BatteryChange-26]
	_ = x[SetAutoOff-27]
	_ = x[PrepareInsulinChange-28]
	_ = x[SuspendPump-30]
	_ = x[ResumePump-31]
	_ = x[SelfTest-32]
	_ = x[Rewind-33]
	_ = x[ClearSettings-34]
	_ = x[EnableChildBlock-35]
	_ = x[MaxBolus-36]
	_ = x[EnableRemote-38]
	_ = x[MaxBasal-44]
	_ = x[EnableBolusWizard-45]
	_ = x[Unknown2E-46]
	_ = x[BolusWizard512-47]
	_ = x[UnabsorbedInsulin512-48]
	_ = x[ChangeBGReminder-49]
	_ = x[SetAlarmClockTime-50]
	_ = x[TempBasalRate-51]
	_ = x[LowReservoir-52]
	_ = x[AlarmClock-53]
	_ = x[ChangeMeterID-54]
	_ = x[BGReceived512-57]
	_ = x[ConfirmInsulinChange-58]
	_ = x[SensorStatus-59]
	_ = x[EnableMeter-60]
	_ = x[BGReceived-63]
	_ = x[MealMarker-64]
	_ = x[ExerciseMarker-65]
	_ = x[InsulinMarker-66]
	_ = x[OtherMarker-67]
	_ = x[EnableSensorAutoCal-68]
	_ = x[ChangeBolusWizardSetup-79]
	_ = x[SensorSetup-80]
	_ = x[Sensor51-81]
	_ = x[Sensor52-82]
	_ = x[ChangeSensorAlarm-83]
	_ = x[Sensor54-84]
	_ = x[Sensor55-85]
	_ = x[ChangeSensorAlert-86]
	_ = x[ChangeBolusStep-87]
	_ = x[BolusWizardSetup-90]
	_ = x[BolusWizard-91]
	_ = x[UnabsorbedInsulin-92]
	_ = x[SaveSettings-93]
	_ = x[EnableVariableBolus-94]
	_ = x[ChangeEasyBolus-95]
	_ = x[EnableBGReminder-96]
	_ = x[EnableAlarmClock-97]
	_ = x[ChangeTempBasalType-98]
	_ = x[ChangeAlarmType-99]
	_ = x[ChangeTimeFormat-100]
	_ = x[ChangeReservoirWarning-101]
	_ = x[EnableBolusReminder-102]
	_ = x[SetBolusReminderTime-103]
	_ = x[DeleteBolusReminderTime-104]
	_ = x[BolusReminder-105]
	_ = x[DeleteAlarmClockTime-106]
	_ = x[DailyTotal515-108]
	_ = x[DailyTotal522-109]
	_ = x[DailyTotal523-110]
	_ = x[ChangeCarbUnits-111]
	_ = x[BasalProfileStart-123]
	_ = x[ConnectOtherDevices-124]
	_ = x[ChangeOtherDevice-125]
	_ = x[ChangeMarriage-129]
	_ = x[DeleteOtherDevice-130]
	_ = x[EnableCaptureEvent-131]
}

const _HistoryRecordType_name = "BolusPrimeAlarmDailyTotalBasalProfileBeforeBasalProfileAfterBGCaptureSensorAlarmClearAlarmChangeBasalPatternTempBasalDurationChangeTimeNewTimeLowBatteryBatteryChangeSetAutoOffPrepareInsulinChangeSuspendPumpResumePumpSelfTestRewindClearSettingsEnableChildBlockMaxBolusEnableRemoteMaxBasalEnableBolusWizardUnknown2EBolusWizard512UnabsorbedInsulin512ChangeBGReminderSetAlarmClockTimeTempBasalRateLowReservoirAlarmClockChangeMeterIDBGReceived512ConfirmInsulinChangeSensorStatusEnableMeterBGReceivedMealMarkerExerciseMarkerInsulinMarkerOtherMarkerEnableSensorAutoCalChangeBolusWizardSetupSensorSetupSensor51Sensor52ChangeSensorAlarmSensor54Sensor55ChangeSensorAlertChangeBolusStepBolusWizardSetupBolusWizardUnabsorbedInsulinSaveSettingsEnableVariableBolusChangeEasyBolusEnableBGReminderEnableAlarmClockChangeTempBasalTypeChangeAlarmTypeChangeTimeFormatChangeReservoirWarningEnableBolusReminderSetBolusReminderTimeDeleteBolusReminderTimeBolusReminderDeleteAlarmClockTimeDailyTotal515DailyTotal522DailyTotal523ChangeCarbUnitsBasalProfileStartConnectOtherDevicesChangeOtherDeviceChangeMarriageDeleteOtherDeviceEnableCaptureEvent"

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
	28:  _HistoryRecordType_name[175:195],
	30:  _HistoryRecordType_name[195:206],
	31:  _HistoryRecordType_name[206:216],
	32:  _HistoryRecordType_name[216:224],
	33:  _HistoryRecordType_name[224:230],
	34:  _HistoryRecordType_name[230:243],
	35:  _HistoryRecordType_name[243:259],
	36:  _HistoryRecordType_name[259:267],
	38:  _HistoryRecordType_name[267:279],
	44:  _HistoryRecordType_name[279:287],
	45:  _HistoryRecordType_name[287:304],
	46:  _HistoryRecordType_name[304:313],
	47:  _HistoryRecordType_name[313:327],
	48:  _HistoryRecordType_name[327:347],
	49:  _HistoryRecordType_name[347:363],
	50:  _HistoryRecordType_name[363:380],
	51:  _HistoryRecordType_name[380:393],
	52:  _HistoryRecordType_name[393:405],
	53:  _HistoryRecordType_name[405:415],
	54:  _HistoryRecordType_name[415:428],
	57:  _HistoryRecordType_name[428:441],
	58:  _HistoryRecordType_name[441:461],
	59:  _HistoryRecordType_name[461:473],
	60:  _HistoryRecordType_name[473:484],
	63:  _HistoryRecordType_name[484:494],
	64:  _HistoryRecordType_name[494:504],
	65:  _HistoryRecordType_name[504:518],
	66:  _HistoryRecordType_name[518:531],
	67:  _HistoryRecordType_name[531:542],
	68:  _HistoryRecordType_name[542:561],
	79:  _HistoryRecordType_name[561:583],
	80:  _HistoryRecordType_name[583:594],
	81:  _HistoryRecordType_name[594:602],
	82:  _HistoryRecordType_name[602:610],
	83:  _HistoryRecordType_name[610:627],
	84:  _HistoryRecordType_name[627:635],
	85:  _HistoryRecordType_name[635:643],
	86:  _HistoryRecordType_name[643:660],
	87:  _HistoryRecordType_name[660:675],
	90:  _HistoryRecordType_name[675:691],
	91:  _HistoryRecordType_name[691:702],
	92:  _HistoryRecordType_name[702:719],
	93:  _HistoryRecordType_name[719:731],
	94:  _HistoryRecordType_name[731:750],
	95:  _HistoryRecordType_name[750:765],
	96:  _HistoryRecordType_name[765:781],
	97:  _HistoryRecordType_name[781:797],
	98:  _HistoryRecordType_name[797:816],
	99:  _HistoryRecordType_name[816:831],
	100: _HistoryRecordType_name[831:847],
	101: _HistoryRecordType_name[847:869],
	102: _HistoryRecordType_name[869:888],
	103: _HistoryRecordType_name[888:908],
	104: _HistoryRecordType_name[908:931],
	105: _HistoryRecordType_name[931:944],
	106: _HistoryRecordType_name[944:964],
	108: _HistoryRecordType_name[964:977],
	109: _HistoryRecordType_name[977:990],
	110: _HistoryRecordType_name[990:1003],
	111: _HistoryRecordType_name[1003:1018],
	123: _HistoryRecordType_name[1018:1035],
	124: _HistoryRecordType_name[1035:1054],
	125: _HistoryRecordType_name[1054:1071],
	129: _HistoryRecordType_name[1071:1085],
	130: _HistoryRecordType_name[1085:1102],
	131: _HistoryRecordType_name[1102:1120],
}

func (i HistoryRecordType) String() string {
	if str, ok := _HistoryRecordType_map[i]; ok {
		return str
	}
	return "HistoryRecordType(" + strconv.FormatInt(int64(i), 10) + ")"
}
