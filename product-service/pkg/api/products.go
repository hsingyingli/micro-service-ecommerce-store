package api

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"product/pkg/db"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type GerProductResponse struct {
	product db.Product
}

func (server *Server) GetProduct(ctx *gin.Context) {
	id, isExist := ctx.GetQuery("id")
	if !isExist || len(id) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "please provide product ID in query"})
		return
	}

	iid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product, err := server.store.GetProduct(ctx, int64(iid))

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rsp := GerProductResponse{
		product: product,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type ListProductsResponse struct {
	products []db.Product
}

func (server *Server) ListProducts(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)

	l := ctx.DefaultQuery("limit", "10")
	o := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(l)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	products, err := server.store.ListProducts(ctx, user.Id, int64(limit), int64(offset))

	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//rsp := ListProductsResponse{
	//products: products,
	//}

	ctx.JSON(http.StatusOK, products)
}

type CreateProductRequest struct {
	Title       string `form:"title" binding:"required"`
	Price       string `form:"price" binding:"required"`
	Amount      string `form:"amount" binding:"required"`
	Description string `form:"description" binding:"required"`
}

func (server *Server) CreateProduct(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)

	var req CreateProductRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get the image data
	image, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer image.Close()

	// get the image type (extension)
	ext := filepath.Ext(file.Filename)
	imageType := strings.TrimPrefix(ext, ".")

	// get the image name
	imageName := strings.TrimSuffix(file.Filename, ext)

	imageData, err := ioutil.ReadAll(image)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	price, err := strconv.ParseInt(req.Price, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	amount, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product, err := server.store.CreateProduct(ctx, db.CreateProductParam{
		UID:         user.Id,
		Title:       req.Title,
		Price:       price,
		Amount:      amount,
		Description: req.Description,
		ImageData:   imageData,
		ImageName:   imageName,
		ImageType:   imageType,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (server *Server) DeleteProductById(ctx *gin.Context) {
	user := ctx.MustGet(authorizationPayloadKey).(*User)
	i := ctx.Query("id")
	id, err := strconv.Atoi(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = server.store.DeleteProductById(ctx, int64(id), user.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
