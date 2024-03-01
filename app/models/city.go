package models

type City struct {
	Name string
}

var Cities = map[string]City{
	"Jerusalem":     {Name: "Jerusalem"},
	"Tel Aviv":      {Name: "Tel Aviv"},
	"Haifa":         {Name: "Haifa"},
	"Rishon LeZion": {Name: "Rishon LeZion"},
	"Petah Tikva":   {Name: "Petah Tikva"},
	"Ashdod":        {Name: "Ashdod"},
	"Netanya":       {Name: "Netanya"},
	"Beer Sheva":    {Name: "Beer Sheva"},
	"Holon":         {Name: "Holon"},
}
