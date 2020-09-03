package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Login struct {
	Token        string `json:"token"`
	UserSettings struct {
		Locale string `json:"locale"`
		Theme  string `json:"theme"`
	} `json:"user_settings"`
}

func main() {
	//Get Credentials
	usern := input("Enter your discord Email: ")
	passw := input("Enter your discord Password: ")
	ck := input("Enter your discord Captcha Key(refer to doc on how to get this): ")
	fmt.Println("Logging In...")
	// Login Request
	login := (`{"email":"` + usern + `","password":"` + passw + `","undelete":false,"captcha_key":"` + ck + `","login_source":null,"gift_code_sku_id":null}`)
	re, _ := http.NewRequest("POST", "https://discord.com/api/v6/auth/login", bytes.NewBuffer([]byte(login)))
	re.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
	re.Header.Set("Accept-Language", "en-US")
	re.Header.Set("Content-Type", "application/json")
	re.Header.Set("Content-Length", "539")
	re.Header.Set("Accept-Encoding", "gzip, deflate")
	re.Header.Set("Referer", "https://discord.com/login?redirect_to=%2Fchannels%2F%40me")
	re.Header.Set("DNT", "1")
	re.Header.Set("X-Fingerprint", "727327584492453979.qqBr0PtKLIV0jh2Sbs_xbEKk5jU")
	re.Header.Set("Accept", "*/*")
	re.Header.Set("X-Super-Properties", "eyJvcyI6IkxpbnV4IiwiYnJvd3NlciI6IkZpcmVmb3giLCJkZXZpY2UiOiIiLCJicm93c2VyX3VzZXJfYWdlbnQiOiJNb3ppbGxhLzUuMCAoWDExOyBMaW51eCB4ODZfNjQ7IHJ2OjY4LjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvNjguMCIsImJyb3dzZXJfdmVyc2lvbiI6IjY4LjAiLCJvc192ZXJzaW9uIjoiIiwicmVmZXJyZXIiOiJodHRwczovL2Rpc2NvcmQuY29tLyIsInJlZmVycmluZ19kb21haW4iOiJkaXNjb3JkLmNvbSIsInJlZmVycmVyX2N1cnJlbnQiOiIiLCJyZWZlcnJpbmdfZG9tYWluX2N1cnJlbnQiOiIiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfYnVpbGRfbnVtYmVyIjo2MjMzMCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")
	re.Header.Set("Accept", "*/*")
	re.Header.Set("Cookie", "__cfduid=d3452cd191c78d7ae65192ec9587dd1251593463237; locale=en-US; __cfruid=7f62496b400df453c509c514e36e7150c14947a2-1593463263")
	//Get login Token
	client := &http.Client{}
	resp, _ := client.Do(re)
	if resp.StatusCode != 200 {
		panic("Couldn't sign-in! Exiting out...")
	} else {
		fmt.Println("Logged in!")
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	newStr := buf.String()
	var logi Login
	tgz := strings.NewReader(newStr)
	gz, _ := gzip.NewReader(tgz)
	defer gz.Close()
	gzs, _ := ioutil.ReadAll(gz)
	json.Unmarshal(gzs, &logi)
	fmt.Println("Your Key is: " + logi.Token)
	key := logi.Token

	//Get Spam Data
	sn := input("Enter Server Number (Refer to doc to find): ")
	rn := input("Enter Room Number (Refer to doc to find): ")
	mes := input("Enter Message: ")
	fmt.Println("Setting up Spam...")
	//Set up Spam
	var rot int
	for rot <= runtime.NumCPU() {
		//Spam
		fmt.Println("Starting spam on core " + strconv.Itoa(rot))
		go spam(rn, mes, sn, key)
		rot++
	}
	input("Press Enter to stop")
	fmt.Println("Stopping...")

}

//Text Spam Function
func spam(rn string, mess string, sn string, key string) {
	var x int
	for {
		ra := rand.Intn(999999999999999-100000000000000) + 1000000000000000000
		ras := strconv.Itoa(ra)
		nonce := "727" + ras
		bd := []byte(`{
    "content":"` + mess + `",
    "nonce":"` + nonce + `",
    "tts":false
}`)
		re, _ := http.NewRequest("POST", "https://discord.com/api/v6/channels/"+rn+"/messages", bytes.NewBuffer(bd))
		re.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
		re.Header.Set("Language", "en-US")
		re.Header.Set("Accept-Encoding", "gzip, deflate")
		re.Header.Set("Referer", "https://discord.com/channels/"+sn+"/"+rn)
		re.Header.Set("Content-Type", "application/json")
		re.Header.Set("Authorization", key)
		re.Header.Set("X-Super-Properties", "eyJvcyI6IkxpbnV4IiwiYnJvd3NlciI6IkZpcmVmb3giLCJkZXZpY2UiOiIiLCJicm93c2VyX3VzZXJfYWdlbnQiOiJNb3ppbGxhLzUuMCAoWDExOyBMaW51eCB4ODZfNjQ7IHJ2OjY4LjApIEdlY2tvLzIwMTAwMTAxIEZpcmVmb3gvNjguMCIsImJyb3dzZXJfdmVyc2lvbiI6IjY4LjAiLCJvc192ZXJzaW9uIjoiIiwicmVmZXJyZXIiOiJodHRwczovL2Rpc2NvcmQuY29tLyIsInJlZmVycmluZ19kb21haW4iOiJkaXNjb3JkLmNvbSIsInJlZmVycmVyX2N1cnJlbnQiOiIiLCJyZWZlcnJpbmdfZG9tYWluX2N1cnJlbnQiOiIiLCJyZWxlYXNlX2NoYW5uZWwiOiJzdGFibGUiLCJjbGllbnRfYnVpbGRfbnVtYmVyIjo2MjMzMCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")
		re.Header.Set("Content-Length", "59")
		re.Header.Set("DNT", "1")
		re.Header.Set("Accept", "*/*")
		re.Header.Set("Cookie", "__cfduid=d3452cd191c78d7ae65192ec9587dd1251593463237; locale=en-US; __cfruid=7f62496b400df453c509c514e36e7150c14947a2-1593463263")
		client := &http.Client{}
		client.Do(re)
		x++
	}
}

//Easly Get User Input (Me being lazy)
func input(t string) string {
	fmt.Print(t)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}
