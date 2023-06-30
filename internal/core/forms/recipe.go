package forms

type RecipeInput struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}

type PartialRecipeInput struct {
	IsActive  bool `json:"is_active"`
	IsRemoved bool `json:"is_removed"`
}
