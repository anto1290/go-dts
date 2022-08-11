package listcontroller

import (
	"net/http"

	"github.com/anto1290/go-dts/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var lists []models.List

	models.DB.Find(&lists)
	c.JSON(http.StatusOK, gin.H{
		"lists": lists,
	})
}
func Get(c *gin.Context) {
	var list models.List
	id := c.Param("id")
	if err := models.DB.First(&list, "id='"+id+"'").Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "record not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})

}
func Create(c *gin.Context) {
	var list models.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	models.DB.Create(&list)
	c.JSON(http.StatusOK, gin.H{
		"list": list,
	})
}
func Update(c *gin.Context) {
	var list models.List
	id := c.Param("id")
	if err := c.ShouldBindJSON(&list); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if models.DB.Model(&list).Where("id = ?", id).Updates(&list).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "record not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "record updated"})
}
func Delete(c *gin.Context) {
	var list models.List
	id := c.Param("id")
	if models.DB.Delete(&list, "id = ?", id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "record not deleted",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "record deleted"})

}
