package node

import (
	"cici/api"
	"cici/hateoas"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getNode(response http.ResponseWriter, request *http.Request) {
	var err error

	path := strings.Replace(
		request.URL.Query().Get("path"),
		"\\",
		"/",
		-1,
	)

	if path == "" {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: http.StatusBadRequest,
			Message:    "path must be supplied",
		})
		return
	}

	depthi := 0
	depth := request.URL.Query().Get("depth")

	if depth != "" {
		depthi, err = strconv.Atoi(depth)
		if err != nil {
			api.WriteErrorResult(request, response, &api.ErrorResult{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			})
			return
		}
	}

	fileinfo, err := os.Stat(path)

	if err != nil {
		api.WriteErrorResult(request, response, &api.ErrorResult{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	parentnode := &nodeInfo{
		Name:     fileinfo.Name(),
		Path:     path,
		ModTime:  fileinfo.ModTime(),
		Mode:     fileinfo.Mode().String(),
		Size:     fileinfo.Size(),
		IsDir:    fileinfo.IsDir(),
		Children: make([]*nodeInfo, 0),
		Links: []*hateoas.Link{
			&hateoas.Link{
				Rel:  "self",
				Href: resourceBase + "?path=" + path,
			},
		},
	}

	if !fileinfo.IsDir() {
		api.WriteSuccessResult(request, response, &api.Result{
			StatusCode: 200,
			Data:       parentnode,
		})
		return
	}

	walkret := make(chan *nodeInfo)

	go walk(parentnode.Path, 1, depthi, parentnode, walkret)

	api.WriteSuccessResult(request, response, &api.Result{
		StatusCode: 200,
		Data:       <-walkret,
	})
}

func walk(currpath string, currdepth int, maxdepth int, parent *nodeInfo, ret chan *nodeInfo) {
	if currdepth > maxdepth || maxdepth < 1 || !parent.IsDir {
		ret <- parent
		return
	}

	files, _ := ioutil.ReadDir(currpath)
	for _, file := range files {
		fullpath := strings.Replace(strings.Join([]string{currpath, file.Name()}, "/"), "//", "/", -1)

		node := &nodeInfo{
			Name:     file.Name(),
			Path:     fullpath,
			ModTime:  file.ModTime(),
			Mode:     file.Mode().String(),
			Size:     file.Size(),
			IsDir:    file.IsDir(),
			Children: make([]*nodeInfo, 0),
			Links: []*hateoas.Link{
				&hateoas.Link{
					Rel:  "self",
					Href: resourceBase + "?path=" + fullpath,
				},
			},
		}

		if !node.IsDir {
			parent.Children = append(parent.Children, node)
		} else {
			retr := make(chan *nodeInfo)
			go walk(node.Path, currdepth+1, maxdepth, node, retr)
			parent.Children = append(parent.Children, <-retr)
		}
	}

	ret <- parent
}
