package smartcharge

import (
	"net/http"
	"strconv"
)

type LiveDataResult struct {
	Result LiveData
}

type LiveData struct {
	Currency                string  `json:"Currency"`
	IsSuspended             bool    `json:"isSuspended"`
	TotalWh                 float64 `json:"TotalWh"`
	LiveKWH                 float64 `json:"LiveKWH"`
	LiveAmps                float64 `json:"LiveAmps"`
	LiveKW                  float64 `json:"LiveKW"`
	LiveVolts               float64 `json:"LiveVolts"`
	LiveAmps_L1             float64 `json:"LiveAmps_L3"`
	LiveAmps_L2             float64 `json:"LiveAmps_L3"`
	LiveAmps_L3             float64 `json:"LiveAmps_L3"`
	CurrentCost             float64 `json:"CurrentCost"`
	CurrentCostCharging     float64 `json:"CurrentCostCharging"`
	CurrentCostChargingTime float64 `json:"CurrentCostChargingTime"`
	CurrentCostOccupied     float64 `json:"CurrentCostOccupied"`
	StartupCost             float64 `json:"StartupCost"`
	IsFinished              bool    `json:"IsFinished"`
}

func (s *SessionService) GetLiveData(sessionId int) (*LiveDataResult, *http.Response, error) {

	reqUrl := "v2/ServiceSessions/LiveData/" + strconv.Itoa(sessionId)
	req, err := s.client.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	liveData := &LiveDataResult{}
	resp, err := s.client.Do(req, liveData)
	if err != nil {
		return nil, resp, err
	}

	return liveData, resp, err
}

type SessionService struct {
	client *Client
}

type SessionsResult struct {
	Result []Session
}

type SessionResult struct {
	Result Session
}

type Session struct {
	SessionId        int     `json:"PK_ServiceSessionID"`
	SessionStart     string  `json:"SessionStart"`
	SessionEnd       string  `json:"SessionEnd"`
	DurationCharging int     `json:"DurationCharging"`
	ChargingBoxID    int     `json:"ChargingBoxID"`
	ChargingPointID  int     `json:"ChargingPointID"`
	CustomerID       int     `json:"FK_CustomerID"`
	TotalkWh         float64 `json:"TotalkWh"`
	LivekW           float64 `json:"LivekW"`
	PointName        string  `json:"PointName"`
	MaxKWH           float64 `json:"MaxKWH"`
	StartMethod      string  `json:"StartMethod"`
}

func (s *SessionService) GetSession(sessionId int) (*SessionResult, *http.Response, error) {

	reqUrl := "v2/ServiceSessions/" + strconv.Itoa(sessionId)
	req, err := s.client.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	session := &SessionResult{}
	resp, err := s.client.Do(req, session)
	if err != nil {
		return nil, resp, err
	}

	return session, resp, err
}

func (s *SessionService) GetActiveSessions(userId int) (*SessionsResult, *http.Response, error) {

	reqUrl := "v2/ServiceSessions/Active/" + strconv.Itoa(userId)
	req, err := s.client.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	sessions := &SessionsResult{}
	resp, err := s.client.Do(req, sessions)
	if err != nil {
		return nil, resp, err
	}

	return sessions, resp, err
}

type MeterValuesResult struct {
	Result MeterValues
}

type MeterValues struct {
	Items     []Item `json:"Items"`
	TotalLogs int    `json:"TotalLogs"`
}

type Item struct {
	Date string  `json:"Date"`
	KWh  float64 `json:"kWh"`
	Amps float64 `json:"Amps"`
	Kw   float64 `json:"kW"`
}

// Defaults to fetch last 60 values. /<num> can be added behind sessionId to fetch more
func (s *SessionService) GetMeterValues(sessionId int) (*MeterValuesResult, *http.Response, error) {

	reqUrl := "v2/Metervalues/SessionResult/" + strconv.Itoa(sessionId)
	req, err := s.client.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, nil, err
	}

	meterValues := &MeterValuesResult{}
	resp, err := s.client.Do(req, meterValues)
	if err != nil {
		return nil, resp, err
	}

	return meterValues, resp, err
}
