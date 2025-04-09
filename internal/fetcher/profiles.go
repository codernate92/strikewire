package fetcher

type ThreatProfile struct {
	Malware        []string `json:"malware"`
	Infrastructure []string `json:"infrastructure"`
	TTPs           []string `json:"ttps"`
	Reports        []string `json:"reports"`
}

var Profiles = map[string]ThreatProfile{
	"APT28": {
		Malware:        []string{"X-Agent", "Sednit", "FancyBearImplants"},
		Infrastructure: []string{"apt28.c2server.net", "185.100.87.84"},
		TTPs:           []string{"T1086", "T1059", "T1071"},
		Reports:        []string{"https://www.mandiant.com/resources/apt28-report"},
	},
	"Lapsus$": {
		Malware:        []string{"None", "Social Engineering Scripts"},
		Infrastructure: []string{"lapsus.fakecdn.net"},
		TTPs:           []string{"T1566", "T1110", "T1078"},
		Reports:        []string{"https://unit42.paloaltonetworks.com/lapsus-group/"},
	},
}
