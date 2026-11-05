package store

import "fmt"

type Store map[string]string

func (s Store) Set(key, value string) {
	s[key] = value
	fmt.Printf("Key: %s, Value: %s set", key, value)
}

func (s Store) Get(key string) (string, error) {
	var returnValue string

	if value, ok := s[key]; !ok {
		fmt.Printf("Error getting the value: %s", key)
		return "", fmt.Errorf("No value found for that key")
	} else {
		returnValue = value
	}

	return returnValue, nil
}

func (s Store) Delete(key string) {
	delete(s, key)
}
