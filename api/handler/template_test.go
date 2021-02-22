package handler

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/rknruben56/feedback-api/entity"
	"github.com/rknruben56/feedback-api/usecase/template/mock"
	"github.com/steinfletcher/apitest"
)

func Test_ListTemplates(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().
		ListTemplates().
		Return([]*entity.Template{temp}, nil).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Get("/v1/template").
		Expect(t).
		Status(http.StatusOK).
		End()
}

func Test_ListTemplates_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		ListTemplates().
		Return(nil, entity.ErrNotFound).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Get("/v1/template").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func Test_GetTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(temp, nil).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Getf("/v1/template/%s", temp.ID.String()).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func Test_GetTemplate_ServerError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(temp, errors.New("server error")).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Getf("/v1/template/%s", temp.ID.String()).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func Test_GetTemplate_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(nil, entity.ErrNotFound).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Getf("/v1/template/%s", temp.ID.String()).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func Test_ListTemplates_ServerError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		ListTemplates().
		Return(nil, errors.New("bad error")).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Get("/v1/template").
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func Test_CreateTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		CreateTemplate(gomock.Any(), gomock.Any()).
		Return(entity.NewID(), nil).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Post("/v1/template").
		Body(`{"class": "Class123", "content": "[Student] is doing well"}`).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func Test_CreateTemplate_ServerError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	service.EXPECT().
		CreateTemplate(gomock.Any(), gomock.Any()).
		Return(entity.NewID(), errors.New("error creating template")).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Post("/v1/template").
		Body(`{"class": "Class123", "content": "[Student] is doing well"}`).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func Test_UpdateTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(temp, nil).
		AnyTimes()
	service.EXPECT().
		UpdateTemplate(temp).
		Return(nil).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	body := fmt.Sprintf(`{"id": "%s", "class": "Class123", "content": "[Student] is doing well"}`, temp.ID.String())
	apitest.New().
		Handler(r).
		Put("/v1/template").
		Body(body).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func Test_UpdateTemplate_GetError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(nil, errors.New("service error")).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	body := fmt.Sprintf(`{"id": "%s", "class": "Class123", "content": "[Student] is doing well"}`, temp.ID.String())
	apitest.New().
		Handler(r).
		Put("/v1/template").
		Body(body).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func Test_UpdateTemplate_NotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(nil, entity.ErrNotFound).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	body := fmt.Sprintf(`{"id": "%s", "class": "Class123", "content": "[Student] is doing well"}`, temp.ID.String())
	apitest.New().
		Handler(r).
		Put("/v1/template").
		Body(body).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}

func Test_UpdateTemplate_UpdateError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().
		GetTemplate(temp.ID).
		Return(temp, nil).
		AnyTimes()
	service.EXPECT().
		UpdateTemplate(temp).
		Return(errors.New("update failed")).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	body := fmt.Sprintf(`{"id": "%s", "class": "Class123", "content": "[Student] is doing well"}`, temp.ID.String())
	apitest.New().
		Handler(r).
		Put("/v1/template").
		Body(body).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}

func Test_DeleteTemplate(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().DeleteTemplate(temp.ID).Return(nil).AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Deletef("/v1/template/%s", temp.ID.String()).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func Test_DeleteTemplate_ServerError(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	service := mock.NewMockUseCase(controller)
	temp := &entity.Template{
		ID: entity.NewID(),
	}
	service.EXPECT().
		DeleteTemplate(temp.ID).
		Return(errors.New("server error")).
		AnyTimes()
	r := mux.NewRouter()
	MakeTemplateHandlers(r, service)

	apitest.New().
		Handler(r).
		Deletef("/v1/template/%s", temp.ID.String()).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()
}
