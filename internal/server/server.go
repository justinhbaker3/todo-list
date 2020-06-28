package server

import (
	"net/http"

	"github.com/justinhbaker3/todo-list/internal/list"

	"github.com/gin-gonic/gin"
)

// Store is used to store and retrieve lists
type Store interface {
	Get(title string) (*list.List, error)
	Upsert(list *list.List)
	Delete(title string)
}

type Server struct {
	store Store
}

func New(store Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) SetupRoutes() *gin.Engine {
	r := gin.New()

	r.GET("/ping", s.Ping)
	r.GET("/list/:title", s.GetList)
	r.POST("/list", s.PostList)
	r.POST("/list/:title", s.PostItem)
	// todo: add more paths

	return r
}

// Ping functions as a healthcheck
func (s *Server) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// GetList returns a list
func (s *Server) GetList(c *gin.Context) {
	title := c.Param("title")
	l, err := s.store.Get(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, l)
}

type PostListRequest struct {
	Title string `json:"title"`
}

// PostList creates a new list
func (s *Server) PostList(c *gin.Context) {
	var req PostListRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}
	s.store.Upsert(list.NewList(req.Title))
	c.Status(http.StatusOK)
}

// TODO: Add delete list

type PostItemRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// PostItem adds an item to an existing list
func (s *Server) PostItem(c *gin.Context) {
	title := c.Param("title")
	l, err := s.store.Get(title)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	var req PostItemRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
		return
	}

	err = l.Append(list.NewItem(req.Title, req.Description))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	s.store.Upsert(l)
	c.JSON(http.StatusOK, l)
}

// TODO: add delete item
