package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Pallinder/go-randomdata"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func strpos(haystack string, needle string) bool {
	if strings.Index(haystack, needle) != -1 {
		return true
	} else {
		return false
	}
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func save(token string, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(token + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var client = &http.Client{}
	question := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("[+] Spotify Mass Account Creator - By: GidhanB.A")
	fmt.Print("[+] Butuh Berapa: ")
	question.Scan()
	qty, _ := strconv.Atoi(question.Text())
	if qty > 100 {
		fmt.Println("[+] Max 100 acc sekali run!")
		question.Scan()
		os.Exit(1)
	}
	fmt.Print("[+] Domain: ")
	question.Scan()
	dom := question.Text()
	fmt.Print("[+] Pass: ")
	question.Scan()
	pass := question.Text()

	var wg sync.WaitGroup
	wg.Add(qty)
	fmt.Println("[+] Running ...")

	for i := 0; i < qty; i++ {
		go func(i int) {
			defer wg.Done()
			email := strings.ToLower(randomdata.SillyName()) + strconv.Itoa(randomInt(10, 99)) + "@" + dom
			req, err := http.NewRequest("POST", "https://spclient.wg.spotify.com:443/signup/public/v1/account/", strings.NewReader("iagree=true&birth_day=12&platform=Android-ARM&creation_point=client_mobile&password="+pass+"&key=142b583129b2df829de3656f9eb484e6&birth_year=2000&email="+email+"&gender=male&app_version=849800892&birth_month=12&password_repeat="+pass))
			check(err)
			req.Header.Set("User-Agent", "Spotify/8.4.98 Android/25 (ASUS_X00HD)")
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Connection", "Keep-Alive")
			res, err := client.Do(req)
			check(err)
			defer res.Body.Close()
			respData, err := ioutil.ReadAll(res.Body)
			respString := string(respData)
			res1 := strings.Split(respString, `"username":"`)
			res2 := strings.Split(res1[1], `"`)
			result := res2[0]
			if len(result) == 25 {
				fmt.Printf("[+] Email: %s - Username: %s\n", email, result)
				save(email+"|"+pass, "akun.txt")
			} else if strpos(result, "try again") == true {
				fmt.Printf("[+] %s\n", result)
			} else {
				fmt.Printf("[+] %s\n", respString)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("[+] Finished")
	fmt.Scanln()
}
