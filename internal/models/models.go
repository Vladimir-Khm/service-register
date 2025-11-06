package models

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Address        string   `json:"address"`
	UserTelegramID int64    `json:"ID"`
	Roles          []string `json:"roles"`

	jwt.StandardClaims
}

type ServiceModel struct {
	ID          uint   `json:"ID" gorm:"primaryKey;autoIncrement:true"`
	ServiceName string `json:"serviceName" gorm:"unique;not null"`

	Methods []Method `json:"methods" gorm:"foreignKey:ServiceModelID;constraint:OnDelete:CASCADE;"`
}

type Method struct {
	ID         uint    `json:"ID" gorm:"primaryKey;autoIncrement:true"`
	MethodName string  `json:"methodName" gorm:"not null"`
	Price      float64 `json:"price"`
	IsPrivate  bool    `json:"isPrivate"`

	ServiceModelID uint `json:"serviceModelID"`

	Arguments    []Argument   `json:"arguments" gorm:"foreignKey:MethodID;constraint:OnDelete:CASCADE;"`
	ServiceModel ServiceModel `json:"-" gorm:"foreignKey:ServiceModelID;references:ID"`
}

type Argument struct {
	ID             uint   `json:"ID" gorm:"primaryKey;autoIncrement:true"`
	ArgumentNumber int32  `json:"argumentNumber"`
	ArgumentName   string `json:"argumentName" gorm:"not null"`
	ArgumentType   string `json:"argumentType"`
	IsRequired     bool   `json:"isRequired"`

	MethodID uint   `json:"methodID"`
	Method   Method `json:"-" gorm:"foreignKey:MethodID;references:ID"`
}
