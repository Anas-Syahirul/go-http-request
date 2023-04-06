package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int	`json:"water"`
	Wind int	`json:"wind"`
}

func main() {
	
	for {
		data := Data{
			Water: rand.Intn(100),
			Wind: rand.Intn(100),
		}
		
		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil{
			log.Fatalln(err)
		}
	
		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}
	
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		
		fmt.Println(string(body))
		dataFinal := Data{}
		err = json.Unmarshal(body,&dataFinal)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("status water : %s\nstatus wind : %s\n", checkWaterStatus(dataFinal.Water), checkWindStatus(dataFinal.Wind))
		time.Sleep(15 * time.Second)
	}
	
}

func checkWaterStatus(waterValue int) string{
	if waterValue <= 5{
		return "aman"
	} else if waterValue >= 6 && waterValue <=8{
		return "siaga"
	} else {
		return "bahaya"
	}
}

func checkWindStatus(windValue int) string{
	if windValue <= 6{
		return "aman"
	} else if windValue >= 7 && windValue <= 15{
		return "siaga"
	} else {
		return "bahaya"
	}
}

