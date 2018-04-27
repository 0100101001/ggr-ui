package config

import (
	"encoding/xml"
)

// Browsers - a set of available browsers
type Browsers struct {
	XMLName  xml.Name  `xml:"urn:config.gridrouter.qatools.ru browsers"`
	Browsers []Browser `xml:"browser"`
}

// Browser - one browser name, e.g. Firefox with all available versions
type Browser struct {
	Name           string    `xml:"name,attr"`
	DefaultVersion string    `xml:"defaultVersion,attr"`
	Versions       []Version `xml:"version"`
}

// Version - concrete browser version
type Version struct {
	Number  string   `xml:"number,attr"`
	Regions []Region `xml:"region"`
}

// Hosts - a list of hosts for browser version
type Hosts []Host

// Region - a datacenter to group hosts
type Region struct {
	Name  string `xml:"name,attr"`
	Hosts Hosts  `xml:"host"`
}

// Host - just a hostname
type Host struct {
	Name     string `xml:"name,attr"`
	Port     int    `xml:"port,attr"`
	Count    int    `xml:"count,attr"`
	Username string `xml:"username,attr,omitempty"`
	Password string `xml:"password,attr,omitempty"`
	VNC      string `xml:"vnc,attr,omitempty"`
	Scheme   string `xml:"scheme,attr,omitempty"`
	region   string
	vncInfo  *vncInfo
}

type vncInfo struct {
	Scheme string
	Host   string
	Port   string
	Path   string
}
