package api

import (
	"net/http"

	db "github.com/akkahshh24/go-dutch/db/sqlc"
	"github.com/akkahshh24/go-dutch/token"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type createGroupRequest struct {
	Name        string   `json:"name" binding:"required,min=3,max=50"`
	Description string   `json:"description" binding:"max=255"`
	Members     []string `json:"members"`
}

func (s *Server) createGroup(c *gin.Context) {
	var req createGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateGroupParams{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   authPayload.UserID,
	}

	group, err := s.store.CreateGroup(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if len(req.Members) > 0 {
		for _, member := range req.Members {
			user, err := s.store.GetUserByName(c.Request.Context(), member)
			if err != nil {
				if err == pgx.ErrNoRows {
					c.JSON(http.StatusNotFound, errorResponse(err))
					return
				}
				c.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			arg := db.AddGroupMemberParams{
				GroupID: group.ID,
				UserID:  user.ID,
			}

			_, err = s.store.AddGroupMember(c.Request.Context(), arg)
			if err != nil {
				c.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	}

	c.JSON(http.StatusCreated,
		successResponseWithData("Group created successfully",
			gin.H{"Name": group.Name}),
	)
}
