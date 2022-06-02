package users

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/fakovacic/users-service/internal/users/errors"
)

const (
	defaultLimit int = 10
)

type Meta struct {
	Page    int       `json:"page"`
	Pages   int       `json:"pages"`
	Count   int64     `json:"count"`
	Limit   int       `json:"limit"`
	Filters []Filters `json:"filters"`
}

type Filters struct {
	Field  string   `json:"field"`
	Values []string `json:"values"`
}

func NewMeta(sortField string) *Meta {
	return &Meta{
		Page:  1,
		Limit: defaultLimit,
	}
}

func (d *Meta) ParseRequest(r *http.Request) error {
	vals := r.URL.Query()

	page := vals.Get("page")
	if page != "" {
		iPage, err := strconv.Atoi(page)
		if err != nil {
			return errors.Wrapf(err, "page not valid format")
		}

		d.Page = iPage
	}

	limit := vals.Get("limit")
	if limit != "" {
		iLimit, err := strconv.Atoi(limit)
		if err != nil {
			return errors.Wrapf(err, "limit not valid format")
		}

		d.Limit = iLimit
	}

	for key, values := range vals {
		if strings.HasPrefix(key, "filters.") {
			key = strings.ReplaceAll(key, "filters.", "")

			if len(values) == 0 {
				continue
			}

			if strings.Join(values, ",") == "" {
				continue
			}

			d.Filters = append(d.Filters, Filters{
				Field:  key,
				Values: values,
			})
		}
	}

	return nil
}

func (d *Meta) CalcPages() {
	d.Pages = (int(d.Count) / d.Limit) + 1
}
