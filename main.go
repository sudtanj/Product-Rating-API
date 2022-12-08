package main

import (
	"bufio"
	"fmt"
	"os"
	"product_rating/database"
	"product_rating/review"
	"product_rating/string_utils"
	"strconv"
	"strings"
)

type Env struct {
	DatabasePath string
}

var env = Env{
	DatabasePath: "./data/reviews.json",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		scanner.Scan()

		text := scanner.Text()

		if strings.Compare("application review:summary", text) == 0 {
			result := review.CalculateSummary(database.GetProductReviewData(env.DatabasePath))
			fmt.Println(string_utils.ToJSONPrettier[review.DefaultResponse](*result))
			continue
		}

		if strings.Contains(text, "application review:product") {
			splitCommand := strings.Split(text, " ")
			if len(splitCommand) != 3 {
				fmt.Println("Invalid command for review product. valid example \"application review:product 1 !\" ")
				continue
			}
			productId, err := strconv.Atoi(splitCommand[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			result := review.CalculateProduct(database.GetProductReviewData(env.DatabasePath), productId)
			fmt.Println(string_utils.ToJSONPrettier[review.DefaultResponse](*result))
			continue
		}

		if strings.Compare("exit", text) == 0 {
			break
		}

		fmt.Println("Invalid syntax!")
	}
}
