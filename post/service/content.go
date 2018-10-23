package service

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/lukma/sample-go-clean/common"
	"github.com/lukma/sample-go-clean/post/data"
	"github.com/lukma/sample-go-clean/post/domain"
)

type service struct {
	repository      domain.ContentRepository
	responseHandler func(c *gin.Context, err error, data map[string]interface{})
}

func NewContentService() domain.ContentService {
	repository := data.NewContentRepository()
	responseHandler := common.JsonHandler

	return &service{
		repository:      repository,
		responseHandler: responseHandler,
	}
}

func (service *service) GetContentsHandler(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	query := c.DefaultQuery("query", "")
	sort := c.DefaultQuery("sort", "created_date")
	if sortMode := c.DefaultQuery("sort_mode", "asc"); sortMode != "asc" {
		sort = "-" + sort
	}

	var contents []domain.ContentEntity
	var recordsTotal int
	var recordsFiltered int

	var errData error
	var errRecordsTotal error
	var errRecordsFiltered error

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		contents, errData = service.repository.GetContents(limit, offset, query, sort)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		recordsTotal, errRecordsTotal = service.repository.CountContent("")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		recordsFiltered, errRecordsFiltered = service.repository.CountContent(query)
	}()

	wg.Wait()

	var err error
	if errData != nil {
		err = errData
	} else if errRecordsTotal != nil {
		err = errRecordsTotal
	} else if errRecordsFiltered != nil {
		err = errRecordsFiltered
	}

	service.responseHandler(c, err, gin.H{
		"data":            contents,
		"recordsTotal":    recordsTotal,
		"recordsFiltered": recordsFiltered,
	})
}

func (service *service) GetContentHandler(c *gin.Context) {
	contentID := c.Param("id")
	content, err := service.repository.GetContent(contentID)
	service.responseHandler(c, err, gin.H{"data": content})
}

func (service *service) CreateContentHandler(c *gin.Context) {
	var form domain.CreateContentForm
	err := c.Bind(&form)

	var thumbnail string
	if err == nil {
		thumbnail, err = common.FileHandler(c, "content", "thumbnail")
	}

	var id string
	if err == nil {
		id, err = service.repository.CreateContent(domain.ContentEntity{
			Title:     form.Title,
			Thumbnail: thumbnail,
			Content:   form.Content,
		})
	}

	service.responseHandler(c, err, gin.H{"data": id})
}

func (service *service) UpdateContentHandler(c *gin.Context) {
	id := c.Param("id")

	var form domain.UpdateContentForm
	err := c.Bind(&form)

	var thumbnail string
	if err == nil {
		thumbnail, err = common.FileHandler(c, "content", "thumbnail")
	}

	if err == nil {
		err = service.repository.UpdateContent(id, domain.ContentEntity{
			Title:     form.Title,
			Thumbnail: thumbnail,
			Content:   form.Content,
		})
	}

	service.responseHandler(c, err, nil)
}

func (service *service) DeleteContentHandler(c *gin.Context) {
	id := c.Param("id")
	err := service.repository.DeleteContent(id)
	service.responseHandler(c, err, nil)
}
