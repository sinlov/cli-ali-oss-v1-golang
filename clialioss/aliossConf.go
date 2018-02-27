package clialioss

type AliossConf struct {
	Client        `json:"client"`
	DefaultBucket `json:"defaultBucket"`
}

type Client struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
}

type DefaultBucket struct {
	Name string `json:"name"`
}
