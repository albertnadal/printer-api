package models

import (
	"time"
)

type PowerHistoryRegister struct {
  DateStart time.Time `json:"dateStart"`
  DateEnd time.Time `json:"dateEnd"`
  Power int32 `json:"power"`
}

type TariffHistoryRegister struct {
  DateStart time.Time `json:"dateStart"`
  DateEnd time.Time `json:"dateEnd"`
  TariffId string `json:"tariffId"`
}

type DeviceRegister struct {
  DateStart time.Time `json:"dateStart"`
  DateEnd time.Time `json:"dateEnd"`
  DeviceId string `json:"deviceId"`
}

type Contract struct {
  ContractId string `json:"contractId"`
  CompanyId int64 `json:"companyId"`
  TariffId string `json:"tariffId"`
  PayerId string `json:"payerId"`
  OwnerId string `json:"ownerId"`
  MeteringPointId string `json:"meteringPointId"`
  DateStart time.Time `json:"dateStart"`
  DateEnd time.Time `json:"dateEnd"`
  Power int32 `json:"power"`
  Version int32 `json:"version"`
  ExperimentalGroupUser bool `json:"experimentalGroupUser"`
  ClimaticZone string `json:"climaticZone"`
  ActivityCode string `json:"activityCode"`
  Status struct {
		Invalid bool `json:"invalid"`
	} `json:"status"`
  Customer struct {
    CustomerId string `json:"customerId"`
    BuildingData struct {
      BuildingHeatingSourceDhw string `json:"buildingHeatingSourceDhw"`
      BuildingType string `json:"buildingType"`
      DwellingArea int32 `json:"dwellingArea"`
      BuildingHeatingSource string `json:"buildingHeatingSource"`
      BuildingConstructionYear int32 `json:"buildingConstructionYear"`
    } `json:"buildingData"`
    Address struct {
      PostalCode string `json:"postalCode"`
      City string `json:"city"`
      CityCode string `json:"cityCode"`
      ProvinceCode string `json:"provinceCode"`
      CountryCode string `json:"countryCode"`
    } `json:"address"`
  } `json:"customer"`
  Power_ struct {
    DateStart time.Time `json:"dateStart"`
    DateEnd time.Time `json:"dateEnd"`
    Power int32 `json:"power"`
  } `json:"power_"`
  Tariff_ struct {
    DateStart time.Time `json:"dateStart"`
    DateEnd time.Time `json:"dateEnd"`
    TariffId string `json:"tariffId"`
  } `json:"tariff_"`
  Report struct {
    InitialMonth int32 `json:"initialMonth"`
    Language string `json:"language"`
  } `json:"report"`
  PowerHistory []PowerHistoryRegister `json:"powerHistory"`
  Devices []DeviceRegister `json:"devices"`
  TariffHistory []TariffHistoryRegister `json:"tariffHistory"`
  UnderscoreVersion int32 `json:"_version"`
  UnderscoreEtag string `json:"_etag"`
  UnderscoreCreated time.Time `json:"_created"`
  UnderscoreUpdated time.Time `json:"_updated"`
}
