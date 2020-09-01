package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize()

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS activities
(
	id SERIAL,
	name TEXT NOT NULL,
	slots INTEGER,
	CONSTRAINT activities_pkey PRIMARY KEY (id)
)`

func clearTable() {
	a.DB.Exec("DELETE FROM activities")
	a.DB.Exec("ALTER SEQUENCE activities_id_seq RESTART WITH 1")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/activities", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentActivity(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/activity/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Event not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Event not found'. Got '%s'", m["error"])
	}
}

func TestCreateActivity(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"name":"test activity", "slots": 5}`)
	req, _ := http.NewRequest("POST", "/activity", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "test activity" {
		t.Errorf("Expected activity name to be 'test activity'. Got '%v'", m["name"])
	}

	if m["slots"] != 5.0 {
		t.Errorf("Expected activity slots to be '5'. Got '%v'", m["slots"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected activity ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetActivity(t *testing.T) {
	clearTable()
	addActivities(1)

	req, _ := http.NewRequest("GET", "/activity/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addActivities(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO activities(name, slots) VALUES($1, $2)", "Activity "+strconv.Itoa(i), (i+1.0)*3)
	}
}

func TestUpdateActivity(t *testing.T) {

	clearTable()
	addActivities(1)

	req, _ := http.NewRequest("GET", "/activity/1", nil)
	response := executeRequest(req)
	var originalActivity map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalActivity)

	var jsonStr = []byte(`{"name":"test activity - updated name", "slots": 151}`)
	req, _ = http.NewRequest("PUT", "/activity/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != originalActivity["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalActivity["id"], m["id"])
	}

	if m["name"] == originalActivity["name"] {
		t.Errorf("Expected the name to change from '%v' to '%v'. Got '%v'", originalActivity["name"], m["name"], m["name"])
	}

	if m["slots"] == originalActivity["slots"] {
		t.Errorf("Expected the slots to change from '%v' to '%v'. Got '%v'", originalActivity["slots"], m["slots"], m["slots"])
	}
}

func TestDeleteActivity(t *testing.T) {
	clearTable()
	addActivities(1)

	req, _ := http.NewRequest("GET", "/activity/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/activity/1", nil)
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/activity/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}
