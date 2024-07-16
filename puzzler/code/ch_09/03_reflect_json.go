package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	ID      string   `json:"id"`
	Genders []string `json:"gender"`
	Age     int      `json:"aga"`
	Name    string   `json:"name"`
}

type ConfigList struct {
	Configs []Config `json:"configs"`
	Config
	Email string `json:"email"`
}

func main() {
	configList := ConfigList{
		Configs: []Config{
			{
				ID:      "1",
				Genders: []string{"asfd", "asfd"},
				Age:     20,
				Name:    "Tom",
			},
			{
				ID:      "2",
				Genders: []string{"asfd", "asfd", "asfd"},
				Age:     30,
				Name:    "Jack",
			},
		},
		Config: Config{
			ID:      "3",
			Genders: []string{"asfd", "asfd", "asfd", "as"},
			Age:     40,
			Name:    "Rose",
		},
		Email: "nsddd.top",
	}
	fmt.Println(configList)

	// iter
	for _, conf := range configList.Configs {
		fmt.Println(conf)
	}

	// marshal
	jsonStr, err := json.Marshal(configList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))

	var cl = &ConfigList{}

	if err := json.Unmarshal([]byte(jsonStr), cl); err != nil {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	} else {
		fmt.Println(cl)
	}

}
