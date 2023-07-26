package form

type DomainRequest struct {
	Name string `json:"name" binding:"required"`
}
