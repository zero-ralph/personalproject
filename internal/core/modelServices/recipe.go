package modelsServices

import (
	"errors"
	"recipes/core/forms"
	"recipes/core/models"
	"time"
)

func GetAllRecipes() (*[]models.Recipe, error) {
	var recipes []models.Recipe

	if err := models.DBConn.Find(&recipes).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return &recipes, nil
}

func GetRecipeByID(recipeId string) (*models.Recipe, error) {
	var recipe models.Recipe

	if err := models.DBConn.Where("id=?", recipeId).First(&recipe).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return &recipe, nil
}

func CreateRecipe(recipeInput *forms.RecipeInput, currentUser models.User) (*models.Recipe, error) {
	recipe := models.Recipe{
		Title:          recipeInput.Title,
		Content:        recipeInput.Content,
		ActiveInactive: models.ActiveInactive{IsActive: recipeInput.IsActive},
		Author:         currentUser.UUIDModel.ID,
	}

	if _, err := recipe.Save(); err != nil {
		return nil, errors.New(err.Error())
	}

	return &recipe, nil
}

// TODO
// Update of the Recipe must be only done by the Author
func UpdateRecipe(recipe *models.Recipe, recipeInput *forms.RecipeInput, currentUser models.User) (*models.Recipe, error) {
	if err := models.DBConn.Model(&recipe).Updates(map[string]interface{}{
		"title":      recipeInput.Title,
		"content":    recipeInput.Content,
		"is_active":  recipeInput.IsActive,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return recipe, nil
}

// TODO
// Update of the Recipe must be only done by the Author
func PartialUpdateRecipe(recipe *models.Recipe, partialRecipeInput *forms.PartialRecipeInput, currentUser models.User) (*models.Recipe, error) {
	if err := models.DBConn.Model(&recipe).Updates(map[string]interface{}{
		"is_active":  partialRecipeInput.IsActive,
		"is_removed": partialRecipeInput.IsRemoved,
	}).Error; err != nil {
		return nil, errors.New(err.Error())
	}
	return recipe, nil
}
