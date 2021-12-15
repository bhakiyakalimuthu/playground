package main

import (
	"encoding/json"
	"fmt"
)

var data = []byte(`
    {
        "client": {
                "id": "2212fw",
                "name": "Papupapa Hernandez",
                "email": "papupapa@gmail.com",
                "phones": ["554-223-2311", "332-232-2123"]
        }
    }
`)

var data1 = []byte(`
	{
    "i$tems": {
        "it$em": [
            {
                "batt$ers": {
                    "ba$tter": [
                        {
                            "i$d": "1001",
                            "t$ype": "Regular"
                        },
                        {
                            "i$d": "1002",
                            "type": "Chocolate"
                        },
                        {
                            "i$d": "1003",
                            "t$ype": "Blueberry"
                        },
                        {
                            "i$d": "1004",
                            "t$ype": "Devil's Food"
                        }
                    ]
                },
                "$id": "0001",
                "$name": "Cake",
                "$ppu": 0.55,
                "$topping": [
                    {
                        "i$d": "5001",
                        "t$ype": "None"
                    },
                    {
                        "i$d": "5002",
                        "t$ype": "Glazed"
                    },
                    {
                        "i$d": "5005",
                        "t$ype": "Sugar"
                    },
                    {
                        "i$d": "5007",
                        "t$ype": "Powdered Sugar"
                    },
                    {
                        "i$d": "5006",
                        "t$ype": "Chocolate with Sprinkles"
                    },
                    {
                        "i$d": "5003",
                        "t$ype": "Chocolate"
                    },
                    {
                        "i$d": "5004",
                        "t$ype": "Maple"
                    }
                ],
                "ty$pe": "donut"
            }
        ]
    }
}`)

type Info struct {
	Client Client `json:"client"`
}

type Client struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Phones []string `json:"phones"`
}

type LongJson struct {
	ITems struct {
		ItEm []struct {
			BattErs struct {
				BaTter []struct {
					ID   string `json:"i$d"`
					TYpe string `json:"t$ype,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"ba$tter"`
			} `json:"batt$ers"`
			ID      string  `json:"$id"`
			Name    string  `json:"$name"`
			Ppu     float64 `json:"$ppu"`
			Topping []struct {
				ID   string `json:"i$d"`
				TYpe string `json:"t$ype"`
			} `json:"$topping"`
			TyPe string `json:"ty$pe"`
		} `json:"it$em"`
	} `json:"i$tems"`
}

func main() {
	var d Info
	if err := json.Unmarshal(data, &d); err != nil {
		fmt.Printf("failed to unmarshal %v", err)
	}
	fmt.Printf("data unmarshal %+v\n", d)

	var l LongJson

	if err := json.Unmarshal(data1, &l); err != nil {
		fmt.Printf("failed to unmarshal %v", err)
	}
	fmt.Printf("data1 unmarshal %+v\n", l)

	b, _ := json.Marshal(&d)
	fmt.Printf("data %+v", string(b))

	//json.NewEncoder().Encode()
	//json.NewDecoder().Decode()
}
