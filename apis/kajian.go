package apis


import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
)

func GetKajiansApi(c *gin.Context){
  c.String(http.StatusOK, "All Kajian")
}

func GetKajianByID(c *gin.Context){
  kid := c.Param("id")
  id, err := strconv.Atoi(kid)

  if err != nil {
    invalidRoute(c)
    return
  }

  msg := fmt.Sprintf("Kajian by %d", id)
  c.String(http.StatusOK, msg)
}
