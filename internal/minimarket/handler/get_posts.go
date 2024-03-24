package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nevano11/minimarket/internal/minimarket/model"
	log "github.com/sirupsen/logrus"
)

func (h *HttpHandler) getPosts(ctx *gin.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, getPostsMaxDuration)
	defer cancel()

	var token *string
	t := ctx.GetHeader("token")
	if len(t) > 0 {
		token = &t
	}

	filter := model.PostFilter{}

	if len(ctx.Query("page-number")) > 0 {
		number, err := strconv.Atoi(ctx.Query("page-number"))
		if err == nil {
			filter.PageNumber = &number
		}
	}
	if len(ctx.Query("page-size")) > 0 {
		number, err := strconv.Atoi(ctx.Query("page-size"))
		if err == nil {
			filter.PageSize = &number
		}
	}
	if len(ctx.Query("price-order")) > 0 {
		number, err := strconv.Atoi(ctx.Query("price-order"))
		if err == nil {
			filter.PriceOrder = model.OrderType(number)
		}
	}
	if len(ctx.Query("date-order")) > 0 {
		number, err := strconv.Atoi(ctx.Query("date-order"))
		if err == nil {
			filter.DateOrder = model.OrderType(number)
		}
	}
	if len(ctx.Query("max-price")) > 0 {
		number, err := strconv.Atoi(ctx.Query("max-price"))
		if err == nil {
			filter.MaxPrice = &number
		}
	}
	if len(ctx.Query("min-price")) > 0 {
		number, err := strconv.Atoi(ctx.Query("min-price"))
		if err == nil {
			filter.MinPrice = &number
		}
	}

	posts, err := h.postService.GetPosts(ctxWithTimeout, filter, token)
	if err != nil {
		log.Errorf("failed to get posts: %s", err.Error())
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
