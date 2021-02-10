package smartcharge

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSessionService_GetLiveData(t *testing.T) {
	jsonBytes, _ := ioutil.ReadFile("testdata/sessions/live-data.json")

	r := *NewMockResponseOkString(string(jsonBytes))
	c := NewMockClient(r)

	session, _, err := c.Session.GetLiveData(123)

	assert.NoError(t, err)
	assert.Equal(t, "NOK", session.Result.Currency)
	assert.Equal(t, false, session.Result.IsFinished)
	assert.Equal(t, 9885.0, session.Result.TotalWh)

}

func TestSessionService_GetSession(t *testing.T) {
	jsonBytes, _ := ioutil.ReadFile("testdata/sessions/session.json")

	r := *NewMockResponseOkString(string(jsonBytes))
	c := NewMockClient(r)

	session, _, err := c.Session.GetSession(123)

	assert.NoError(t, err)
	assert.Equal(t, 1299, session.Result.ChargingBoxID)
	assert.Equal(t, 6.899, session.Result.LivekW)
	assert.Equal(t, "2021-02-11T00:20:41.85", session.Result.SessionStart)
}

func TestSessionService_GetActiveSessions(t *testing.T) {
	jsonBytes, _ := ioutil.ReadFile("testdata/sessions/active.json")

	r := *NewMockResponseOkString(string(jsonBytes))
	c := NewMockClient(r)

	activeSessions, _, err := c.Session.GetActiveSessions(123)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(activeSessions.Result))
	assert.Equal(t, 1299, activeSessions.Result[0].ChargingBoxID)
}

func TestSessionService_GetMeterValues(t *testing.T) {
	jsonBytes, _ := ioutil.ReadFile("testdata/sessions/meter-values-session.json")

	r := *NewMockResponseOkString(string(jsonBytes))
	c := NewMockClient(r)

	values, _, err := c.Session.GetMeterValues(123)

	assert.NoError(t, err)
	assert.Equal(t, 116, values.Result.TotalLogs)
	assert.Equal(t, 60, len(values.Result.Items))
	assert.Equal(t, 9.885, values.Result.Items[0].KWh)
}
