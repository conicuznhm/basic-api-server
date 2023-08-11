package services

import (
	"encoding/json"
	"os"
)

type Profile struct{
	Name string
	Pet string
	Age int
	Address string
}

func FetchProfile() (*Profile, error){
	filePath := "dataApi/user.json"
	jsonBytes, err := os.ReadFile(filePath)

	if err != nil{
		return nil, err
	}
	
	var profile Profile
	if err := json.Unmarshal(jsonBytes, &profile); err != nil{
		return nil, err
	}
	return &profile, nil
}