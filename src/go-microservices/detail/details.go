package detail

import "os"

func GetHostname() (string, error) {
	hostname, err := os.Hostname
	if err != nil {
		return nil, err
	}
	
	return hostname, err
}

//func Hostname() (name string, err error)