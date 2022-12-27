package detail

import "os"

func GetHostname() (string, error) {
	hostname, err := os.Hostname
	return hostname, err
}

//func Hostname() (name string, err error)