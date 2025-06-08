package api

import "github.com/gin-gonic/gin"

type addGroupMemberRequest struct {
	GroupName string `json:"group_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
}

func (s *Server) addGroupMember(c *gin.Context) {}
