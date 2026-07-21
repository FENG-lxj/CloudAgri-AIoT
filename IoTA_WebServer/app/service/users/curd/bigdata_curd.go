package curd

import "goskeleton/app/model"

func CreateAEDataCurdFactory() *AEDataCurd {
	return &AEDataCurd{model.CreateInfluxDBAEDatasFactory()}
}

func CreateASDataCurdFactory() *ASDataCurd {
	return &ASDataCurd{model.CreateInfluxDBASDatasFactory()}
}

func CreateDevManageCurdFactory() *DevManageCurd {
	return &DevManageCurd{model.CreateDevManageFactory("")}
}

type AEDataCurd struct {
	AEDataModels *model.AEModels
}

type ASDataCurd struct {
	ASDataModels *model.ASModels
}

type DevManageCurd struct {
	DevManageModel *model.DevManageModel
}

func (a *AEDataCurd) GetAEDatas(AEID int32, StartTime string) *model.AEModels {
	return a.AEDataModels.GetAEdatas(AEID, StartTime)
}

func (a *ASDataCurd) GetASDatas(ASID int32, StartTime string) *model.ASModels {
	return a.ASDataModels.GetASdatas(ASID, StartTime)
}

func (a *DevManageCurd) GetDevInfo(Types string) (model.DevManageModels, error) {
	return a.DevManageModel.GetDevInfo(Types)
}

func (a *DevManageCurd) UpDev(Types string, DevID int32, Longitude, Dimension float64, Cropper string) error {
	return a.DevManageModel.UpDev(Types, DevID, Longitude, Dimension, Cropper)
}

func (a *DevManageCurd) RemoveDev(Types string, DevID int32) error {
	return a.DevManageModel.RemoveDev(Types, DevID)
}
