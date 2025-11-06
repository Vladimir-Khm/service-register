package models

type UpdateServiceDTO struct {
	ServiceName *string `json:"serviceName"`
}

type UpdateMethodDTO struct {
	MethodName *string  `json:"methodName"`
	IsPrivate  *bool    `json:"isPrivate"`
	Price      *float64 `json:"price"`
}

type UpdateArgumentDTO struct {
	ArgumentNumber *int32  `json:"argumentNumber"`
	ArgumentName   *string `json:"argumentName"`
	ArgumentType   *string `json:"argumentType"`
	IsRequired     *bool   `json:"isRequired"`
}
