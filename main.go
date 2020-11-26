package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func areAnagrams(word string, checked string) bool {
	if len(word) != len(checked) {
		return false
	}

	hash := make(map[string]int)

	for _, char := range word {
		charCount := hash[string(char)]
		if charCount == 0 {
			hash[string(char)] = 1
		} else {
			hash[string(char)] = charCount + 1
		}
	}

	for _, char := range checked {
		charCount := hash[string(char)]
		if charCount == 0 {
			hash[string(char)] = 1
		} else {
			hash[string(char)] = charCount + 1
		}
	}

	// if the strings are anagram,
	// then the counts will be even
	for _, value := range hash {
		if value % 2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	words := []string{};

	r := gin.Default()

	r.POST("/load", func(c *gin.Context) {
		err := c.BindJSON(&words);
		if (err != nil) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"words": words,
			})
		}
	})

	r.GET("/get", func(c *gin.Context) {
		checkedWord := c.Query("word")
		if (checkedWord == "") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Word is empty",
			})
		} else {
			result := []string{};
			for _, word := range words {
				if areAnagrams(word, checkedWord) {
					result = append(result, word)
				}
			}
			c.JSON(http.StatusOK, result)
		}
	})

	r.Run()
}