package tpi

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stratadigm/tpi_data"
	"google.golang.org/appengine"
	"net/http"
	_ "net/http/httptest"
	"strconv"
	"testing"
	_ "time"
)

var (
	id  = int64(0)
	g1  = &tpi_data.User{Name: "Roger", Email: "jan@uary.com", Password: "allurbase", Confirmed: false, Rep: 0}
	g2  = &tpi_data.User{Name: "Bob", Email: "jan@uary.com", Password: "allurbase", Confirmed: false, Rep: 0}
	uri = "http://192.168.0.9:8080" //"https://2.thalipriceindex.appspot.com"
)

func TestCreateUser(t *testing.T) {

	var err error
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err := http.NewRequest("POST", uri+"/user", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Response: %v", resp.Status)
	}
	g2 := &tpi_data.User{}
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(g2)
	if err != nil || g1.Name != g2.Name {
		t.Errorf("Response: %v", g2)
	}
	id = g2.Id

}

func TestRetrieveUser(t *testing.T) {

	//Retrieve
	var err error
	req, err := http.NewRequest("GET", uri+"/user/"+string(id), nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	u1 := &tpi_data.User{}
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(u1)
	if err != nil || u1.Name != g1.Name {
		t.Errorf("Response: %v", u1)
	}

}

func TestUpdateUser(t *testing.T) {

	var err error
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g2)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err := http.NewRequest("PUT", uri+"/user"+"/10000002", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Response: %v", resp.Status)
	}
	g2 := &tpi_data.User{}
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(g2)
	if err != nil || g1.Name == g2.Name {
		t.Errorf("Response: %v", g2)
	}

}

func TestDeleteUser(t *testing.T) {

	var err error
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err := http.NewRequest("DELETE", uri+"/user"+"/10000002", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Response: %v", resp.Status)
	}

}

func TestCRUDVenue(t *testing.T) {

	var err error
	g1 := &tpi_data.Venue{Name: "Udupi Uphara", Location: appengine.GeoPoint{Lat: 13.9, Lng: 75.4}}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	//Create
	req, err := http.NewRequest("POST", "https://thalipriceindex.appspot.com/create/venue", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	//req.Header.Set("X-Custom-Header", "")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Response: %v", resp.Status)
	}

	//Retrieve

	//g2 := &User{}
	//dec := json.NewDecoder(resp.Body)
	//err = json.Decode(g2)

	//Update

	//Delete

}

func TestLogin(t *testing.T) {

	var err error
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	//Login
	req, err := http.NewRequest("POST", "https://thalipriceindex.appspot.com/auth_token", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	tok := &tpi_data.AuthToken{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(tok)
	if err != nil {
		t.Errorf("Json resp err: %v", err)
	} //else {
	//	t.Errorf("Json token: %v", tok)
	//}
	resp.Body.Close()

	//Test token
	//req, err = http.NewRequest("GET", "http://192.168.0.9:8080/hello", nil)
	req, err = http.NewRequest("GET", "https://thalipriceindex.appspot.com/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", string(tok))) // Authorized
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "abcd")) // Unauthorized
	//req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	hw := make([]byte, int(resp.ContentLength))
	bdec := bufio.NewReader(resp.Body)
	_, err = bdec.Read(hw)
	if err != nil {
		t.Errorf("Read resp err: %v", err)
	}

}

func TestRefresh(t *testing.T) {

	var err error
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err := http.NewRequest("PUT", uri+"/auth_token", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	cl := resp.Header.Get("Content-Length")
	icl, err := strconv.Atoi(cl)
	if err != nil {
		t.Errorf("Content Length err: %v", err)
	}
	tok := make([]byte, icl)
	dec := bufio.NewReader(resp.Body)
	_, err = dec.Read(tok)
	if err != nil {
		t.Errorf("Json resp err: %v", err)
	} else {
		t.Errorf("Json token: %v", string(tok))
	}

}
