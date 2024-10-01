package main

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

// most used words and times of occurrence in the text
// input:- I am manish manish is good person manish like cricket player manish like music manish is now holyboll player// output// [// {// "word" : "manish",// "count":5// },// {// "word" : "like",// "count":2// }

type tekionRequest struct {
	Input string `json:"input"`
}

type tekionResponse struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func countHandler(ctx *gin.Context) {
	var ij tekionRequest
	err := ctx.BindJSON(&ij)
	if err != nil {
		panic(err)
	}
	words := strings.Split(ij.Input, " ")
	var keys []string
	countMap := make(map[string]int)
	for _, word := range words {
		c := countMap[word]
		if c == 0 {
			keys = append(keys, word)
		}
		countMap[word] = c + 1
	}
	fmt.Println(countMap)
	fmt.Println(keys)
	slices.SortStableFunc(keys, func(i, j string) int {
		result := 0
		if countMap[i] > countMap[j] {
			result = -1
		}
		return result
	})
	fmt.Println(keys)
	var response []tekionResponse
	for i, k := range keys {
		if i > 9 {
			break
		}
		response = append(response, tekionResponse{k, countMap[k]})
	}
	ctx.JSON(http.StatusOK, response)
}
