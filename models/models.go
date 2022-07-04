package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title             string
	CriticScore 	  int
	CriticSentiment   int
	AudienceScore 	  int
	AudienceSentiment int
	IsEmailSent 	  bool
}
