package services

import (
	"encoding/json"
	"errors"
	"os"
)

type Profile struct{
	Id int
	Name string
	Pet string
	Age int
	Address string
}

// get user by id
func FetchProfile(id int) (*Profile, error){
	filePath := "dataApi/user.json"
	jsonBytes, err := os.ReadFile(filePath)

	if err != nil{
		return nil, err
	}
	
	var profiles []Profile
	if err := json.Unmarshal(jsonBytes, &profiles); err != nil{
		return nil, err
	}

	for _, profile := range profiles{
		if profile.Id == id{
			return &profile, nil
		}
	}
	// return &profile, nil
	return nil, errors.New("Profile not found on this id")
}

// get all user
func FetchAllProfile()([]Profile, error){
	const filePath string = "dataApi/user.json"
	jsonBytes, err := os.ReadFile(filePath)

	if err != nil{
		return nil, err
	}

	var profiles []Profile
	if err := json.Unmarshal(jsonBytes, &profiles); err != nil{
		return nil, err
	}

	return profiles, nil
}