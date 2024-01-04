package fizzbuzz

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrorNotInteger   = errors.New("invalid parameters: int1, int2, and limit must be integers")
	ErrorZeroParamter = errors.New("invalid parameters: int1, int2, and limit must be greater than 0")
)

type Params struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func GetEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		params, err := validate(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := fizzbuzz(params)

		c.JSON(http.StatusOK, gin.H{"result": res})
	}
}

func fizzbuzz(p Params) string {
	var result strings.Builder

	for i := 1; i <= p.Limit; i++ {
		switch {
		case i%p.Int1 == 0 && i%p.Int2 == 0:
			result.WriteString(p.Str1 + p.Str2)
		case i%p.Int1 == 0:
			result.WriteString(p.Str1)
		case i%p.Int2 == 0:
			result.WriteString(p.Str2)
		default:
			result.WriteString(strconv.Itoa(i))
		}

		if i < p.Limit {
			result.WriteString(",")
		}
	}

	return result.String()
}

func validate(c *gin.Context) (Params, error) {
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(c.Query("int1")) {
		return Params{}, ErrorNotInteger
	}

	if !regexp.MustCompile(`^[0-9]+$`).MatchString(c.Query("int2")) {
		return Params{}, ErrorNotInteger
	}

	if !regexp.MustCompile(`^[0-9]+$`).MatchString(c.Query("limit")) {
		return Params{}, ErrorNotInteger
	}

	int1, err := strconv.Atoi(c.Query("int1"))
	if err != nil {
		return Params{}, err
	}

	int2, err := strconv.Atoi(c.Query("int2"))
	if err != nil {
		return Params{}, err
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return Params{}, err
	}

	if int1 <= 0 || int2 <= 0 || limit <= 0 {
		return Params{}, ErrorZeroParamter
	}

	return Params{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  c.Query("str1"),
		Str2:  c.Query("str2"),
	}, nil
}
