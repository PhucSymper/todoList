package model

type Item struct {
	Id         string `json: "Id" form:"Id" query:"Id"`
	Title      string `json: "Title" form:"Title" query:"Title"`
	IsComplete bool   `json: "IsComplete" form:"IsComplete" query:"isComplete"`
}

// type ListItem struct {
// 	ListItem []Item
// }

type Student struct {
	Id   string `json: "Id" form:"Id" query:"Id"`
	Name string `json: "Name" form:"Name" query:"Name"`
	Date string `json: "Date" form:"Date" query:"Date"`
	Msv  string `json: "Msv" form:"Msv" query:"Msv"`
}
