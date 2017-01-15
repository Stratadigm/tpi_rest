package tpi

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stratadigm/tpi_data"
	"google.golang.org/appengine"
	"image/jpeg"
	"net/http"
	_ "net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	id        = int64(0)
	g1        = &tpi_data.User{Name: "Roger", Email: "jan@uary.com", Password: "allurbase", Confirmed: false, Rep: 0}
	g2        = &tpi_data.User{Name: "Bob", Email: "jan@uary.com", Password: "allurbase", Confirmed: false, Rep: 0}
	uri       = "http://192.168.0.9:8080" //"https://2.thalipriceindex.appspot.com"
	prod      = "http://2.thalipriceindex.appspot.com"
	blacklist = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODQ0MTA5MDYsIkVtYWlsIjoiamFuQHVhcnkuY29tIiwiVHlwZSI6IkNvbnRyaWJ1dG9yIn0.g6_5-viyOsjCHKu7QFEiI8FkAEVXBSKD0itEAh2C1okW4OSBG0EFCPOuMqXF0voHtFid3UglaZZ44z07CNnATgXfDlVO1ZLNBfebM2xL995i07oaHz3FBAKgdCggm3bBPVIsfUt_8NwqdJuFu9BOKPq4e16yzS-fXAT2PBn8PLU0kxafSnCi8ZDO6r2XvpuriSPRosY1W0oLrQXw8c4fDbI4kNblzWYYzxe4V2zfKFY6ZOEDwMZM96uKamhi_y1D_vWJJQXdh6hrNgu_e9FVSFsSXUDXvACzXQbJ2viJZq5GAfVOkzGq-FJ42ewJ5mOZfjcxJ0-ru1Iro69X9o2JMA"
)

//Go 1.7
/*func TestUser(t *testing.T) {

	var err error
	var buf bytes.Buffer
	g3 := &tpi_data.User{}
	t.Run("Create", func(t *testing.T) {
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
		jd := json.NewDecoder(resp.Body)
		err = jd.Decode(g3)
		if err != nil || g1.Name != g3.Name {
			t.Errorf("Response: %v %v", g1.Name, g3.Name)
		}
		id = g3.Id
	})

	time.Sleep(2 * time.Second) // need to sleep to let datastore persist

	t.Run("Retrieve", func(t *testing.T) {
		req, err := http.NewRequest("GET", uri+"/user/"+strconv.FormatInt(id, 10), nil)
		if err != nil {
			t.Errorf("Request : %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Client do request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Response: %v", resp.Status)
		}
		jd := json.NewDecoder(resp.Body)
		err = jd.Decode(g3)
		if err != nil || g3.Name != g1.Name {
			t.Errorf("Response: %v %v", g1.Name, g3.Name)
		}
	})
	t.Run("Update", func(t *testing.T) {
		enc := json.NewEncoder(&buf)
		err = enc.Encode(g2)
		if err != nil {
			t.Errorf("Encode json : %v", err)
		}
		req, err := http.NewRequest("PUT", uri+"/user"+"/"+strconv.FormatInt(id, 10), &buf)
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
		jd := json.NewDecoder(resp.Body)
		err = jd.Decode(g3)
		if err != nil || g2.Name != g3.Name {
			t.Errorf("Response: %v %v", g2.Name, g3.Name)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		v := url.Values{}
		v.Set("email", g2.Email)
		v.Set("fullname", g2.Name)
		req, err := http.NewRequest("DELETE", uri+"/user"+"/"+strconv.FormatInt(id, 10)+"?"+v.Encode(), nil)
		if err != nil {
			t.Errorf("Request : %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("Client do request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusAccepted {
			t.Errorf("Response: %v", resp.Status)
		}
	})
}*/

func TestCRUDUser(t *testing.T) {

	var err error
	var buf bytes.Buffer
	g3 := &tpi_data.User{}
	//Create
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
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(g3)
	if err != nil || g1.Name != g3.Name {
		t.Errorf("Response: %v %v", g1.Name, g3.Name)
	}
	id = g3.Id

	time.Sleep(2 * time.Second) // need to sleep to let datastore persist

	//Retrieve
	req, err = http.NewRequest("GET", uri+"/user/"+strconv.FormatInt(id, 10), nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	jd = json.NewDecoder(resp.Body)
	err = jd.Decode(g3)
	if err != nil || g3.Name != g1.Name {
		t.Errorf("Response: %v %v", g1.Name, g3.Name)
	}

	//Update
	enc = json.NewEncoder(&buf)
	err = enc.Encode(g2)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err = http.NewRequest("PUT", uri+"/user"+"/"+strconv.FormatInt(id, 10), &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Response: %v", resp.Status)
	}
	jd = json.NewDecoder(resp.Body)
	err = jd.Decode(g3)
	if err != nil || g2.Name != g3.Name {
		t.Errorf("Response: %v %v", g2.Name, g3.Name)
	}
	//Delete
	v := url.Values{}
	v.Set("email", g2.Email)
	v.Set("fullname", g2.Name)
	req, err = http.NewRequest("DELETE", uri+"/user"+"/"+strconv.FormatInt(id, 10)+"?"+v.Encode(), nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Response: %v", resp.Status)
	}

}

func TestDeleteUser(t *testing.T) {

	var err error
	v := url.Values{}
	v.Set("email", g1.Email)
	v.Set("fullname", g1.Name)
	req, err := http.NewRequest("DELETE", uri+"/user"+"/10000062?"+v.Encode(), nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}

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
	g3 := &tpi_data.User{}

	//Create credentials
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
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(g3)
	if err != nil || g1.Name != g3.Name {
		t.Errorf("Response: %v %v", g1.Name, g3.Name)
	}
	id = g3.Id

	time.Sleep(2 * time.Second) // need to sleep to let datastore persist

	//Login with created credentials
	enc = json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err = http.NewRequest("POST", uri+"/auth_token", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	tok := &tpi_data.AuthToken{}
	jd = json.NewDecoder(resp.Body)
	err = jd.Decode(tok)
	if err != nil {
		t.Errorf("Json resp err: %v", err)
	} //else {
	//	t.Errorf("Json token: %v", tok)
	//}
	resp.Body.Close()

	//Test token receieved
	req, err = http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	hw := make([]byte, int(resp.ContentLength))
	bdec := bufio.NewReader(resp.Body)
	_, err = bdec.Read(hw)
	if err != nil {
		t.Errorf("Read resp err: %v", err)
	}
	resp.Body.Close()

	//Logout
	req, err = http.NewRequest("DELETE", uri+"/auth_token", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	resp.Body.Close()

	//Test token again after logout - should fail
	req, err = http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Response: %v", resp.Status)
	}

	//Delete the user
	v := url.Values{}
	v.Set("email", g1.Email)
	v.Set("fullname", g1.Name)
	req, err = http.NewRequest("DELETE", uri+"/user"+"/"+strconv.FormatInt(id, 10)+"?"+v.Encode(), nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Response: %v", resp.Status)
	}

}

func TestBlacklist(t *testing.T) {

	var err error
	//Test token
	req, err := http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", blacklist)) // Unauthorized

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Response: %v", resp.Status)
	}

}

func TestRefresh(t *testing.T) {

	var err error
	var buf bytes.Buffer
	g3 := &tpi_data.User{}

	//Create credentials
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
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(g3)
	if err != nil || g1.Name != g3.Name {
		t.Errorf("Response: %v %v", g1.Name, g3.Name)
	}
	id = g3.Id

	time.Sleep(2 * time.Second) // need to sleep to let datastore persist

	//Login with created credentials
	enc = json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		t.Errorf("Encode json : %v", err)
	}
	req, err = http.NewRequest("POST", uri+"/auth_token", &buf)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	tok := &tpi_data.AuthToken{}
	jd = json.NewDecoder(resp.Body)
	err = jd.Decode(tok)
	if err != nil {
		t.Errorf("Json resp err: %v", err)
	} //else {
	//	t.Errorf("Json token: %v", tok)
	//}
	resp.Body.Close()

	//Test token receieved - assumed 10s validity
	req, err = http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	hw := make([]byte, int(resp.ContentLength))
	bdec := bufio.NewReader(resp.Body)
	_, err = bdec.Read(hw)
	if err != nil {
		t.Errorf("Read resp err: %v", err)
	}
	resp.Body.Close()

	time.Sleep(12 * time.Second)

	//Test token again after expiry - should fail since we waited 10s
	req, err = http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("Response: %v", resp.Status)
	}

	//Refresh token
	req, err = http.NewRequest("PUT", uri+"/auth_token", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Expired

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Response: %v", resp.Status)
	}
	tok = &tpi_data.AuthToken{}
	jd = json.NewDecoder(resp.Body)
	err = jd.Decode(tok)
	if err != nil {
		t.Errorf("Json resp err: %v", err)
	}
	resp.Body.Close()

	//Test refreshed token 	- should succeed
	req, err = http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}

	//Delete the user
	v := url.Values{}
	v.Set("email", g1.Email)
	v.Set("fullname", g1.Name)
	req, err = http.NewRequest("DELETE", uri+"/user"+"/"+strconv.FormatInt(id, 10)+"?"+v.Encode(), nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("Response: %v", resp.Status)
	}

}

func TestRetrieveImage(t *testing.T) {

	//Retrieve
	var err error
	req, err := http.NewRequest("GET", prod+"/thali/1000017/image", nil)
	if err != nil {
		t.Errorf("Request : %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response: %v", resp.Status)
	}
	img, err := jpeg.Decode(resp.Body)
	if err != nil {
		t.Errorf("Response: %v", err.Error())
	}
	w, err := os.Create("/home/sridhar/dev/go/src/github.com/stratadigm/test.jpg")
	if err != nil {
		t.Logf("%v", err)
	}
	defer w.Close()
	if err = jpeg.Encode(w, img, nil); err != nil {
		t.Errorf("Failed to write: %v\n", err.Error())
	}

}
