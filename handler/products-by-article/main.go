package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vados1/superkid-serverless-go/model"
)

type Product struct {
	CategoryName   string   `json:"categoryName"`
	Description    string   `json:"description"`
	Name           string   `json:"name"`
	Article        string   `json:"article"`
	Price          int      `json:"price"`
	PhotoLinkShort string   `json:"photoLinkShort"`
	PhotoLinkLong  string   `json:"photoLinkLong"`
	Sizes          []string `json:"sizes"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	article := request.PathParameters["article"]

	var result Product
	for _, prod := range model.Products {
		for _, subProd := range prod.SubProducts {
			if article == subProd.Article {
				var sizes []string
				for _, size := range prod.Sizes {
					sizes = append(sizes, size.Size)
				}

				result = Product{
					CategoryName:   prod.Category.Name,
					Description:    prod.Category.Description,
					Name:           subProd.Name,
					Article:        subProd.Article,
					Price:          subProd.Price,
					PhotoLinkShort: subProd.PhotoLinkShort,
					PhotoLinkLong:  subProd.PhotoLinkLong,
					Sizes:          sizes,
				}
			}
		}
	}

	var buf bytes.Buffer
	body, err := json.Marshal(result)
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
