package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/answersuck/host/internal/domain/tag"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type tagHandler struct {
	log      *zap.Logger
	validate validate
	service  tagService
}

func newTagMux(d *Deps) *chi.Mux {
	h := tagHandler{
		log:      d.Logger,
		validate: d.Validate,
		service:  d.TagService,
	}

	m := chi.NewMux()
	verificator := mwVerificator(d.Logger, &d.Config.Session, d.SessionService)

	m.Post("/list", h.getAll)
	m.With(verificator).Post("/", h.createMultiple)

	return m
}

type tagGetAllReq struct {
	Filter struct {
		Name       string `json:"name"`
		LanguageId uint   `json:"language_id"`
	} `json:"filter"`
	LastId uint32 `json:"last_id"`
	Limit  uint64 `json:"limit"`
}

func (h *tagHandler) getAll(w http.ResponseWriter, r *http.Request) {
	var req tagGetAllReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.log.Info("http - v1 - tag - getAll - Decode", zap.Error(err))
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	f := tag.Filter{
		Name:       req.Filter.Name,
		LanguageId: req.Filter.LanguageId,
	}
	tagList, err := h.service.GetAll(r.Context(), tag.NewListParams(req.LastId, req.Limit, f))
	if err != nil {
		h.log.Error("http - v1 - tag - getAll - h.service.GetAll", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, tagList)
}

type tagCreateMultipleReq struct {
	Tags []struct {
		Name       string `json:"name" validate:"required,gte=1,lte=32"`
		LanguageId uint   `json:"language_id" validate:"required"`
	} `json:"tags" validate:"required,min=1,max=10"`
}

func (r tagCreateMultipleReq) validate(v validate) error {
	for _, reqTag := range r.Tags {
		if err := v.Struct(reqTag); err != nil {
			return err
		}
	}
	return nil
}

func (h *tagHandler) createMultiple(w http.ResponseWriter, r *http.Request) {
	var req tagCreateMultipleReq

	if err := h.validate.RequestBody(r.Body, &req); err != nil {
		h.log.Info("http - v1 - tag - createMultiple - RequestBody", zap.Error(err))
		writeValidationErr(w, http.StatusBadRequest, errInvalidRequestBody, h.validate.TranslateError(err))
		return
	}

	if err := req.validate(h.validate); err != nil {
		writeValidationErr(w, http.StatusBadRequest, errInvalidRequestBody, h.validate.TranslateError(err))
		return
	}

	t := make([]tag.Tag, 0, len(req.Tags))
	for _, rt := range req.Tags {
		t = append(t, tag.Tag{
			Name:       rt.Name,
			LanguageId: uint8(rt.LanguageId),
		})
	}

	tags, err := h.service.CreateMultiple(r.Context(), t)
	if err != nil {
		h.log.Error("http - v1 - tag - createMultiple - h.service.CreateMultiple", zap.Error(err))

		if errors.Is(err, tag.ErrLanguageIdNotFound) {
			writeErr(w, http.StatusBadRequest, tag.ErrLanguageIdNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, listResponse{tags})
}
