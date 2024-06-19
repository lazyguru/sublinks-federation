package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sublinks/sublinks-federation/internal/activitypub"
	"sublinks/sublinks-federation/internal/model"

	"github.com/gorilla/mux"
)

type Outbox struct {
	Context      *Context            `json:"@context"`
	Type         string              `json:"type"`
	Id           string              `json:"id"`
	OrderedItems []*activitypub.Page `json:"orderedItems"`
	TotalItems   int                 `json:"totalItems"`
}

func (server *Server) SetupApubRoutes() {
	server.Logger.Debug("Setting up Apub routes")
	server.Router.HandleFunc("/{type}/{id}/inbox", server.getInboxHandler).Methods("GET")
	server.Router.HandleFunc("/{type}/{id}/inbox", server.postInboxHandler).Methods("POST")
	server.Router.HandleFunc("/{type}/{id}/outbox", server.getOutboxHandler).Methods("GET")
	server.Router.HandleFunc("/{type}/{id}/outbox", server.postOutboxHandler).Methods("POST")
}

func (server *Server) getInboxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/activity+json")
}

func (server *Server) postInboxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/activity+json")
}

func (server *Server) getOutboxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/activity+json")
	var outbox Outbox
	vars := mux.Vars(r)
	switch vars["type"] {
	case "u":
		user := server.ServiceManager.UserService().GetByUsername(vars["id"])
		if user == nil {
			server.Logger.Error(fmt.Sprintf("User %s not found", vars["id"]), nil)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		outboxItems := getPages(user.Posts)
		outbox = Outbox{
			Context:      GetContext(),
			Type:         "OrderedCollection",
			Id:           fmt.Sprintf("%s/outbox", user.Id),
			OrderedItems: outboxItems,
			TotalItems:   len(outboxItems),
		}
		w.WriteHeader(http.StatusOK)
	case "c":
		group := server.ServiceManager.CommunityService().GetByUsername(vars["id"])
		if group == nil {
			server.Logger.Error(fmt.Sprintf("Community %s not found", vars["id"]), nil)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		outboxItems := getPages(group.Posts)
		outbox = Outbox{
			Context:      GetContext(),
			Type:         "OrderedCollection",
			Id:           fmt.Sprintf("%s/outbox", group.Id),
			OrderedItems: outboxItems,
			TotalItems:   len(outboxItems),
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if outbox.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	content, _ := json.MarshalIndent(outbox, "", "  ")
	_, err := w.Write(content)
	if err != nil {
		server.Logger.Error("Error writing response", err)
	}
}

func getPages(posts []model.Post) []*activitypub.Page {
	var pages []*activitypub.Page
	for _, post := range posts {
		postLd := activitypub.ConvertPostToPage(&post)
		postLd.Context = activitypub.GetContext()
		pages = append(pages, postLd)
	}
	return pages
}

func (server *Server) postOutboxHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/activity+json")
}

type Context []interface{}

func GetContext() *Context {
	return &Context{
		"https://join-lemmy.org/context.json",
		"https://www.w3.org/ns/activitystreams",
	}
}
