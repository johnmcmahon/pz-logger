// Copyright 2016, RadiantBlue Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logger

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venicegeo/pz-gocommon/gocommon"
)

type Server struct {
	service *Service
	Routes  []piazza.RouteData
}

func (server *Server) handleGetRoot(c *gin.Context) {
	resp := server.service.GetRoot()
	c.IndentedJSON(resp.StatusCode, resp)
}

func (server *Server) handlePostMessage(c *gin.Context) {
	var mssg Message
	err := c.BindJSON(&mssg)
	if err != nil {
		resp := &piazza.JsonResponse{StatusCode: http.StatusBadRequest, Message: err.Error()}
		c.IndentedJSON(resp.StatusCode, resp)
	}
	resp := server.service.PostMessage(&mssg)
	c.IndentedJSON(resp.StatusCode, resp)
}

func (server *Server) handleGetStats(c *gin.Context) {
	resp := server.service.GetStats()
	c.IndentedJSON(resp.StatusCode, resp)
}

func (server *Server) handleGetMessage(c *gin.Context) {
	params := piazza.NewQueryParams(c.Request)
	resp := server.service.GetMessage(params)
	c.IndentedJSON(resp.StatusCode, resp)
}

func (server *Server) Init(service *Service) {
	server.service = service

	server.Routes = []piazza.RouteData{
		{"GET", "/", server.handleGetRoot},
		{"GET", "/message", server.handleGetMessage},
		{"POST", "/message", server.handlePostMessage},
		{"GET", "/admin/stats", server.handleGetStats},
	}
}