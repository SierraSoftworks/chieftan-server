package models

import "gopkg.in/mgo.v2/bson"

type ProjectSummary struct {
	ID   bson.ObjectId `json:"id"`
	Name string        `json:"name"`
	URL  string        `json:"url"`
}

type Project struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	URL         string        `json:"url"`
}

func (p *Project) Summary() *ProjectSummary {
	return &ProjectSummary{
		ID:   p.ID,
		Name: p.Name,
		URL:  p.URL,
	}
}
