package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"math/rand"
	"vadim.dissa.gmail.com/superkid-lambda-go/model"
)

const numberOfFeaturedSubProducts = 6

type Product struct {
	Name           string `json:"name"`
	Article        string `json:"article"`
	PhotoLinkShort string `json:"photoLinkShort"`
	Price          int    `json:"price"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	article := request.PathParameters["article"]
	subProducts := remove(model.SubProducts, article)
	featuredSubProducts := getFeaturedSubProducts(subProducts)

	var results []Product
	for _, subProduct := range featuredSubProducts {
		results = append(results, Product{
			Name:           subProduct.Name,
			Article:        subProduct.Article,
			PhotoLinkShort: subProduct.PhotoLinkShort,
			Price:          subProduct.Price,
		})
	}

	var buf bytes.Buffer
	body, err := json.Marshal(results)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
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

func remove(subProducts []model.SubProduct, article string) []model.SubProduct {
	var filteredSubProducts []model.SubProduct

	for _, subProduct := range subProducts {
		if subProduct.Article != article {
			filteredSubProducts = append(filteredSubProducts, subProduct)
		}
	}

	return filteredSubProducts
}

func getFeaturedSubProducts(subProducts []model.SubProduct) []model.SubProduct {
	var featuredSubProducts []model.SubProduct

	for i := 0; i < numberOfFeaturedSubProducts; i++ {
		subProduct := subProducts[rand.Intn(len(subProducts))]
		featuredSubProducts = append(featuredSubProducts, subProduct)
		subProducts = remove(subProducts, subProduct.Article)
	}

	return featuredSubProducts
}

func main() {
	lambda.Start(Handler)
}
