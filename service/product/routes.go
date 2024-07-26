package product

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Parz1val02/ecom/types"
	"github.com/Parz1val02/ecom/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products/{id}", h.handleGetProductByID).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = utils.WriteJSON(w, http.StatusOK, ps)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleGetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error when parsing id"))
		return
	}

	p, err := h.store.GetProductByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return

	}

	err = utils.WriteJSON(w, http.StatusOK, p)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// Obtain json payload
	var payload types.CreateProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	// Create product
	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, map[string]string{"success": fmt.Sprintf("product with name %s created successfully", payload.Name)})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
