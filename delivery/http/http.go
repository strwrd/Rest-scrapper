package http

import (
	"context"
	"net/http"
	"time"

	"github.com/strwrd/rest-scrapper/model"

	"github.com/labstack/echo"
	"github.com/strwrd/rest-scrapper/usecase"
)

// Delivery ...
type Delivery interface {
	Start() error
	Stop() error
}

// handler ...
type handler struct {
	server  *echo.Echo
	usecase usecase.Usecase
}

// NewHandler ...
func NewHandler(usecase usecase.Usecase) Delivery {
	// Creating echo server
	server := echo.New()

	// Return handler object
	return &handler{
		server,
		usecase,
	}
}

// Start server listening port
func (h *handler) Start() error {

	// GET /archieves
	h.server.GET("/archieves", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Do usecase
		res, err := h.usecase.GetAllArchieve(ctx)
		if err != nil {
			e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			return c.JSON(e.Code, e)
		}
		return c.JSON(http.StatusOK, res)
	})

	// GET /journals  || GET /journals?archieveId=:val
	h.server.GET("/journals", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Checking query parameter
		if c.QueryParam("archieveId") != "" {
			// Do usecase
			res, err := h.usecase.GetJournalsByArchieveID(ctx, c.QueryParam("archieveId"))
			if err != nil {
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			return c.JSON(http.StatusOK, res)
		}

		// Do usecase
		res, err := h.usecase.GetAllJournal(ctx)
		if err != nil {
			e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			return c.JSON(e.Code, e)
		}
		return c.JSON(http.StatusOK, res)
	})

	// GET /archieve?archieveId=:val || GET /archieve?code=:val
	h.server.GET("/archieve", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Checking query parameter
		if c.QueryParam("archieveId") != "" {
			// Do usecase
			res, err := h.usecase.GetArchieveByArchieveID(ctx, c.QueryParam("archieveId"))
			if err != nil {
				if err == model.ErrDataNotFound {
					e := echo.NewHTTPError(http.StatusNotFound, err.Error())
					return c.JSON(e.Code, e)
				}
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			return c.JSON(http.StatusOK, res)
		} else if c.QueryParam("code") != "" {
			// Do usecase
			res, err := h.usecase.GetArchieveByCode(ctx, c.QueryParam("code"))
			if err != nil {
				if err == model.ErrDataNotFound {
					e := echo.NewHTTPError(http.StatusNotFound, err.Error())
					return c.JSON(e.Code, e)
				}
				e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				return c.JSON(e.Code, e)
			}
			return c.JSON(http.StatusOK, res)
		} else {
			return c.JSON(http.StatusNotAcceptable, "unknown query parameter")
		}
	})

	// GET /journal/:id
	h.server.GET("/journal/:id", func(c echo.Context) error {
		// Create timeout request
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		// Do usecase
		res, err := h.usecase.GetJournalByJournalID(ctx, c.Param("id"))
		if err != nil {
			if err == model.ErrDataNotFound {
				e := echo.NewHTTPError(http.StatusNotFound, err.Error())
				return c.JSON(e.Code, e)
			}
			e := echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			return c.JSON(e.Code, e)
		}
		return c.JSON(http.StatusOK, res)
	})

	return h.server.Start(":8080")
}

// Shutdown() ...
func (h *handler) Stop() error {
	// Create timeout process
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return h.server.Shutdown(ctx)
}
