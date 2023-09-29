package service

import "github.com/Seunghoon-Oh/cloud-ml-experiments-manager/data"

func GetExps() []string {
	return data.GetRedisData()
}

func CreateExp() string {
	// TODO: Create Kubernetes Object
	return data.CreateRedisData()
}
