package order

import (
	"net/http"
	"order/api/configs"
	"order/api/internal/jwt"
	"order/api/internal/model"
	"order/api/pkg/middleware"
	"order/api/pkg/request"
	"order/api/pkg/response"
	"strconv"
)

type OrderHandler struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	Config          *configs.Config
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps) {
	orderHandler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
		Config:          deps.Config,
	}

	router.Handle("POST /order", middleware.IsAuth(orderHandler.CreateOrder(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuth(orderHandler.GetOrderById(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuth(orderHandler.GetMyOrders(), deps.Config))
}

func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[CreateOrderRequest](w, r)
		if err != nil {
			return
		}

		userID, ok := r.Context().Value(jwt.UserIDKey).(uint)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		order := model.Order{
			UserID: userID,
		}

		createdOrder, err := h.OrderRepository.Create(&order, payload.ProductIDs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response.Response(w, http.StatusCreated, orderToResponse(createdOrder))
	}
}

func (h *OrderHandler) GetOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(jwt.UserIDKey).(uint)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		orderID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, "invalid order id", http.StatusBadRequest)
			return
		}

		order, err := h.OrderRepository.GetById(uint(orderID))
		if err != nil {
			http.Error(w, "order not found", http.StatusNotFound)
			return
		}

		if order.UserID != userID {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		response.Response(w, http.StatusOK, order)
	}
}
func (h *OrderHandler) GetMyOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(jwt.UserIDKey).(uint)
		if !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		orders, err := h.OrderRepository.GetByUserId(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		out := make([]GetOrderResponse, len(orders))
		for i, o := range orders {
			out[i] = toDTO(o)
		}

		response.Response(w, http.StatusOK, out)
	}
}
