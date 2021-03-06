package tpi

import (
	"bufio"
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	//"github.com/dgrijalva/jwt-go/request"
	"github.com/stratadigm/tpi_auth"
	"github.com/stratadigm/tpi_data"
	//"google.golang.org/appengine"
	"io/ioutil"
	"net/http"
	_ "net/http/httptest"
	"os"
	_ "strconv"
	//"testing"
	"time"
)

const (
	pubKeyPath = "settings/keys/public_key.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	verifyKey *rsa.PublicKey
)

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func init() {

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	fatal(err)

}

func ExampleToken() {

	var err error
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err = enc.Encode(g1)
	if err != nil {
		fmt.Printf("Encode json : %v\n", err)
	}
	//Login
	req, err := http.NewRequest("POST", uri+"/auth_token", &buf)
	if err != nil {
		fmt.Printf("Request : %v \n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Client do request: %v \n", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Response: %v", resp.Status)
	}
	tok := &tpi_data.AuthToken{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(tok)
	if err != nil {
		fmt.Printf("Json resp err: %v", err)
	} //else {
	//	t.Errorf("Json token: %v", tok)
	//}
	defer resp.Body.Close()

	err = ioutil.WriteFile("temp", []byte(tok.Token), os.ModePerm)
	fatal(err)

	token, err := jwt.ParseWithClaims(tok.Token, &tpi_auth.TPIClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	if err != nil {
		fmt.Printf("JWT Parse: %v", err)
	}

	claims := token.Claims.(*tpi_auth.TPIClaims)
	fmt.Println(claims.UserInfo.Email, claims.StandardClaims.ExpiresAt-time.Now().Unix())

	//Test token
	req, err = http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		fmt.Printf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tok.Token)) // Authorized
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", string(tok))) // Authorized
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "abcd")) // Unauthorized
	//req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Client do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Response: %v", resp.Status)
	}
	hw := make([]byte, int(resp.ContentLength))
	bdec := bufio.NewReader(resp.Body)
	_, err = bdec.Read(hw)
	fmt.Println(string(hw))

	//Output:
	//dec@ember.com 478
	//Hello, World!

}

func ExampleTokenFail() {

	var err error

	tok, err := ioutil.ReadFile("temp")
	fatal(err)

	//Test token
	req, err := http.NewRequest("GET", uri+"/hello", nil)
	if err != nil {
		fmt.Printf("Request : %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", string(tok))) // Authorized
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", string(tok))) // Authorized
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", "abcd")) // Unauthorized
	//req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Client do request: %v", err)
	}
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Response: %v", resp.Status)
	}

	//Output:
	//Response: 401 Unauthorized

}
