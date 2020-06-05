package main

import (
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/forgoer/openssl"
	"io/ioutil"
	"os"
)

type NaviConnections struct {
	XMLName    xml.Name         `xml:"Connections"`
	Connection []NaviConnection `xml:"Connection"`
}

type NaviConnection struct {
	ConnectionName  string `xml:"ConnectionName,attr"`
	ConnType        string `xml:"ConnType,attr"`
	ServiceProvider string `xml:"ServiceProvider,attr"`
	Host            string `xml:"Host,attr"`
	Port            string `xml:"Port,attr"`
	UserName        string `xml:"UserName,attr"`
	Password        string `xml:"Password,attr"`
	SSH_Host        string `xml:"SSH_Host,attr"`
	SSH_UserName    string `xml:"SSH_UserName,attr"`
	SSH_Password    string `xml:"SSH_Password,attr"`
}

func decodepwd(pwd string) (string, error) {
	aesKey := []byte("libcckeylibcckey")
	aesIV := []byte("libcciv libcciv ")
	passDecode, err := hex.DecodeString(pwd)
	if err != nil {
		return "", err
	}
	dec, err := openssl.AesCBCDecrypt(passDecode, aesKey, aesIV, openssl.PKCS5_PADDING)
	if err != nil {
		return "", err
	}
	return string(dec), nil
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "./navi.ncx", "navicat export file")
	flag.Parse()
	file, err := os.Open(filePath) // For read access.

	if err != nil {
		fmt.Printf("file can't open. error: %s", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	navis := NaviConnections{}
	err = xml.Unmarshal(data, &navis)
	if err != nil {
		fmt.Printf("xml parser error: %v", err)
		return
	}
	connectionData := []map[string]string{}
	for _, e := range navis.Connection {
		output := make(map[string]string)

		output["ConnectionName"] = e.ConnectionName
		output["ConnType"] = e.ConnType
		output["ServiceProvider"] = e.ServiceProvider
		output["Host"] = e.Host
		output["Port"] = e.Port
		output["UserName"] = e.UserName
		//decode password
		output["Password"] = ""
		output["Password"], err = decodepwd(e.Password)
		if err != nil {
			fmt.Printf("Password parser error: %s", err)
			return
		}
		output["SSH_Host"] = e.SSH_Host
		output["SSH_UserName"] = e.SSH_UserName
		output["SSH_Password"] = ""
		if e.SSH_Password != "" {
			output["SSH_Password"], err = decodepwd(e.SSH_Password)
			if err != nil {
				fmt.Printf("Password parser error: %s", err)
				return
			}
		}

		connectionData = append(connectionData, output)
	}

	jsonStr, err := json.Marshal(connectionData)

	if err != nil {
		fmt.Println(" ToJson err: ", err)
	}

	fmt.Println(string(jsonStr))

}
