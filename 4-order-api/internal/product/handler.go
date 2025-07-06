package product

import (
	"net/http"
	"order/api/pkg/request"
	"order/api/pkg/response"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandler struct {
	productRepository *ProductRepository
}

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	productHandler := &ProductHandler{
		productRepository: deps.ProductRepository,
	}

	router.HandleFunc("GET /products", productHandler.GetProducts())
	router.HandleFunc("GET /products/{id}", productHandler.GetProductById())
	router.HandleFunc("POST /products", productHandler.CreateProduct())
	router.HandleFunc("PATCH /products/{id}", productHandler.UpdateProduct())
	router.HandleFunc("DELETE /products/{id}", productHandler.DeleteProduct())

}

func (h *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.productRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Response(w, http.StatusOK, products)
	}
}

func (h *ProductHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[CreateProductRequest](w, r)
		if err != nil {
			return
		}
		product := NewProduct(payload.Name, payload.Description, payload.Images)
		_, err = h.productRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Response(w, http.StatusCreated, product)
	}
}

func (h *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := r.PathValue("id")
		id, err := strconv.ParseUint(stringId, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.productRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Response(w, http.StatusNoContent, nil)
	}
}

func (h *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := r.PathValue("id")
		id, err := strconv.ParseUint(stringId, 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payload, err := request.HandleBody[UpdateProductRequest](w, r)

		if err != nil {
			return
		}

		product := &Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        payload.Name,
			Description: payload.Description,
			Images:      payload.Images,
		}

		product, err = h.productRepository.Update(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Response(w, http.StatusOK, product)
	}
}

func (h *ProductHandler) GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringId := r.PathValue("id")
		id, err := strconv.ParseUint(stringId, 10, 32)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := h.productRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Response(w, http.StatusOK, product)
	}
}
