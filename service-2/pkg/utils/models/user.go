package models

type AllUserDetails struct {
	Count uint64   `json:"count"`
	Names []string `json:"names"`
}

type MethodType struct {
	Method   int `json:"method" binding:"min=1,max=2"` //  only accepting 1 or 2
	WaitTime int `json:"waitTime" `
}
