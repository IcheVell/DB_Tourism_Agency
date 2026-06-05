package handler

import (
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

func writeInternalError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "internal server error",
	})
}

func writeDatabaseError(c echo.Context, err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return writeInternalError(c)
	}

	switch pgErr.Code {
	case "23505":
		return c.JSON(http.StatusConflict, map[string]string{
			"error":      "unique constraint violation",
			"constraint": pgErr.ConstraintName,
		})

	case "23503":
		return c.JSON(http.StatusConflict, map[string]string{
			"error":      "foreign key constraint violation",
			"constraint": pgErr.ConstraintName,
		})

	case "23514":
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":      "check constraint violation",
			"constraint": pgErr.ConstraintName,
		})

	case "23502":
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":  "not null constraint violation",
			"column": pgErr.ColumnName,
		})

	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":      "database error",
			"constraint": pgErr.ConstraintName,
			"code":       pgErr.Code,
		})
	}
}
