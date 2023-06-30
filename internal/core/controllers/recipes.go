package controllers

import (
	"net/http"
	"recipes/core/forms"
	modelsServices "recipes/core/modelServices"
	"recipes/core/models"

	"github.com/gin-gonic/gin"
)

func RecipesList(c *gin.Context) {
	recipes, err := modelsServices.GetAllRecipes()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(*recipes),
		"data":  recipes,
	})
}

func CreateRecipe(c *gin.Context) {
	var recipeInput forms.RecipeInput

	if err := c.ShouldBindJSON(&recipeInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err := modelsServices.CreateRecipe(&recipeInput, c.MustGet("currentUser").(models.User))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Created Recipe",
		"data":    recipe,
	})
}

func UpdateRecipe(c *gin.Context) {

	recipe, err := modelsServices.GetRecipeByID(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	var recipeInput forms.RecipeInput
	if err := c.ShouldBindJSON(&recipeInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err = modelsServices.UpdateRecipe(recipe, &recipeInput, c.MustGet("currentUser").(models.User))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Successfully Updated Recipe",
		"recipe":  recipe,
	})
}

func PartialUpdateRecipe(c *gin.Context) {

	recipe, err := modelsServices.GetRecipeByID(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}
	
	var partialRecipeInput forms.PartialRecipeInput
	if err := c.ShouldBindJSON(forms.PartialRecipeInput{}); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe, err = modelsServices.PartialUpdateRecipe(recipe, &partialRecipeInput, c.MustGet("currentUser").(models.User))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": "Successfull Patch Update Recipe",
		"recipe":  recipe,
	})
}

func RecipeDetails(c *gin.Context) {

	recipe, err := modelsServices.GetRecipeByID(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   recipe,
	})
}
