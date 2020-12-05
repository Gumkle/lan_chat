package main

import "github.com/Gumkle/consoler/consoler"

func setupHosts(self *Host, logger *consoler.Logger) []Host {
	potentialHostIp := askForKnownHost()
	if potentialHostIp != nil {
		foundHost, err := connectToHost(potentialHostIp)
		if err != nil {
			logger.PrintError("Podany host nie odpowiada")
			setupHosts(nil, nil)
		}
	} else {
		foundHost, err := searchLanForHost()
	}
	hosts := askForHosts(foundHost)
	return hosts
}
