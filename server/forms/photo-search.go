package forms

import (
	"time"
)

type PhotoSearchForm struct {
	Query         string    `form:"q"`
	Tags          string    `form:"tags"`
	Category      string    `form:"category"`
	Country       string    `form:"country"`
	CameraID      int       `form:"camera_id"`
	Order         string    `form:"order"`
	Count         int       `form:"count" binding:"required"`
	Offset        int       `form:"offset"`
	Before        time.Time `form:"before" time_format:"2006-01-02"`
	After         time.Time `form:"after" time_format:"2006-01-02"`
	FavoritesOnly bool      `form:"category"`
}
