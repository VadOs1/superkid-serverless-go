package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"vadim.dissa.gmail.com/superkid-lambda-go/model"
)

type Product struct {
	CategoryId     int    `json:"categoryId"`
	Name           string `json:"name"`
	Article        string `json:"article"`
	PhotoLinkShort string `json:"photoLinkShort"`
	Price          int    `json:"price"`
}

func Handler() (events.APIGatewayProxyResponse, error) {

	var results []Product
	for _, prod := range model.Products {
		for _, subProd := range prod.SubProducts {
			results = append(results, Product{
				CategoryId:     prod.Category.Id,
				Name:           prod.Category.Name,
				Article:        subProd.Article,
				PhotoLinkShort: subProd.PhotoLinkShort,
				Price:          subProd.Price,
			})
		}
	}

	var buf bytes.Buffer
	body, err := json.Marshal(results)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	json.HTMLEscape(&buf, body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       buf.String(),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
