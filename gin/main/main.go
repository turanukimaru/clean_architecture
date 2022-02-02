// go get -u github.com/gin-gonic/gin
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/turanukimaru/ca/usecase/dummycalc"
	"net/http"
	"strconv"
)

type Adder struct {
	A int
	B int
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/add/:a/:b", func(c *gin.Context) {
		a, errA := strconv.Atoi(c.Param("a"))
		if errA != nil {

		}
		b, errB := strconv.Atoi(c.Param("b"))
		if errB != nil {

		}
		sum := dummycalc.Add(a, b)
		message := fmt.Sprintf("%d + %d is %d", a, b, sum)
		c.String(http.StatusOK, message)
	})
	router.POST("/add", func(c *gin.Context) {
		var a Adder
		if err := c.ShouldBind(&a); err != nil {
			message := "Bad Parameters"
			c.String(http.StatusBadRequest, message)
			return
		}
		fmt.Printf("A:%s", c.Param("a"))
		fmt.Printf("B:%s", c.Param("b"))
		sum := dummycalc.Add(a.A, a.B)
		message := fmt.Sprintf("%d + %d is %d", a.A, a.B, sum)
		c.String(http.StatusOK, message)

	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
