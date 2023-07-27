// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: editor/v1/tag.proto

package editorv1

import context "context"
import fmt "fmt"
import http "net/http"
import io "io"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// ====================
// TagService Interface
// ====================

type TagService interface {
	CreateTag(context.Context, *CreateTagRequest) (*CreateTagResponse, error)

	ListTags(context.Context, *ListTagsRequest) (*ListTagsResponse, error)
}

// ==========================
// TagService Protobuf Client
// ==========================

type tagServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTagServiceProtobufClient creates a Protobuf client that implements the TagService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewTagServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TagService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "editor.v1", "TagService")
	urls := [2]string{
		serviceURL + "CreateTag",
		serviceURL + "ListTags",
	}

	return &tagServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *tagServiceProtobufClient) CreateTag(ctx context.Context, in *CreateTagRequest) (*CreateTagResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TagService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateTag")
	caller := c.callCreateTag
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateTagRequest) (*CreateTagResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTagRequest) when calling interceptor")
					}
					return c.callCreateTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *tagServiceProtobufClient) callCreateTag(ctx context.Context, in *CreateTagRequest) (*CreateTagResponse, error) {
	out := new(CreateTagResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *tagServiceProtobufClient) ListTags(ctx context.Context, in *ListTagsRequest) (*ListTagsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TagService")
	ctx = ctxsetters.WithMethodName(ctx, "ListTags")
	caller := c.callListTags
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListTagsRequest) (*ListTagsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListTagsRequest) when calling interceptor")
					}
					return c.callListTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *tagServiceProtobufClient) callListTags(ctx context.Context, in *ListTagsRequest) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ======================
// TagService JSON Client
// ======================

type tagServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewTagServiceJSONClient creates a JSON client that implements the TagService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewTagServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) TagService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "editor.v1", "TagService")
	urls := [2]string{
		serviceURL + "CreateTag",
		serviceURL + "ListTags",
	}

	return &tagServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *tagServiceJSONClient) CreateTag(ctx context.Context, in *CreateTagRequest) (*CreateTagResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TagService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateTag")
	caller := c.callCreateTag
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateTagRequest) (*CreateTagResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTagRequest) when calling interceptor")
					}
					return c.callCreateTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *tagServiceJSONClient) callCreateTag(ctx context.Context, in *CreateTagRequest) (*CreateTagResponse, error) {
	out := new(CreateTagResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *tagServiceJSONClient) ListTags(ctx context.Context, in *ListTagsRequest) (*ListTagsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TagService")
	ctx = ctxsetters.WithMethodName(ctx, "ListTags")
	caller := c.callListTags
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ListTagsRequest) (*ListTagsResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListTagsRequest) when calling interceptor")
					}
					return c.callListTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *tagServiceJSONClient) callListTags(ctx context.Context, in *ListTagsRequest) (*ListTagsResponse, error) {
	out := new(ListTagsResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =========================
// TagService Server Handler
// =========================

type tagServiceServer struct {
	TagService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewTagServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewTagServiceServer(svc TagService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwards compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &tagServiceServer{
		TagService:       svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *tagServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *tagServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// TagServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const TagServicePathPrefix = "/twirp/editor.v1.TagService/"

func (s *tagServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "TagService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "editor.v1.TagService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "CreateTag":
		s.serveCreateTag(ctx, resp, req)
		return
	case "ListTags":
		s.serveListTags(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *tagServiceServer) serveCreateTag(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateTagJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateTagProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *tagServiceServer) serveCreateTagJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateTag")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(CreateTagRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TagService.CreateTag
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateTagRequest) (*CreateTagResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTagRequest) when calling interceptor")
					}
					return s.TagService.CreateTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateTagResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateTagResponse and nil error while calling CreateTag. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tagServiceServer) serveCreateTagProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateTag")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(CreateTagRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TagService.CreateTag
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateTagRequest) (*CreateTagResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateTagRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateTagRequest) when calling interceptor")
					}
					return s.TagService.CreateTag(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateTagResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateTagResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateTagResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateTagResponse and nil error while calling CreateTag. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tagServiceServer) serveListTags(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveListTagsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveListTagsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *tagServiceServer) serveListTagsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListTags")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(ListTagsRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.TagService.ListTags
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListTagsRequest) (*ListTagsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListTagsRequest) when calling interceptor")
					}
					return s.TagService.ListTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListTagsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListTagsResponse and nil error while calling ListTags. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tagServiceServer) serveListTagsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ListTags")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := io.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(ListTagsRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.TagService.ListTags
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ListTagsRequest) (*ListTagsResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ListTagsRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ListTagsRequest) when calling interceptor")
					}
					return s.TagService.ListTags(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*ListTagsResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*ListTagsResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *ListTagsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *ListTagsResponse and nil error while calling ListTags. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tagServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor5, 0
}

func (s *tagServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *tagServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "editor.v1", "TagService")
}

var twirpFileDescriptor5 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x8f, 0x93, 0x50,
	0x10, 0x0e, 0xd2, 0x74, 0xcb, 0x34, 0xda, 0xfa, 0x34, 0x8a, 0xac, 0x46, 0x42, 0x8c, 0xd9, 0x13,
	0xa4, 0x18, 0x13, 0x93, 0xbd, 0xb1, 0x89, 0x17, 0x8d, 0x31, 0xc8, 0xc9, 0x83, 0x64, 0xba, 0x1d,
	0x9f, 0xc4, 0x85, 0x87, 0x8f, 0x29, 0xb1, 0x3f, 0xc5, 0xbf, 0xea, 0xc9, 0xc0, 0x83, 0xb6, 0xa9,
	0xbd, 0xcd, 0x7c, 0xf3, 0xbd, 0xef, 0xfb, 0x66, 0x00, 0x1e, 0xd1, 0xa6, 0x60, 0xa5, 0xa3, 0x76,
	0x15, 0x31, 0xca, 0xb0, 0xd6, 0x8a, 0x95, 0x70, 0x0c, 0x18, 0xb6, 0x2b, 0xef, 0x69, 0x8b, 0x77,
	0xc5, 0x06, 0x99, 0xa2, 0xb1, 0x30, 0x1c, 0xef, 0xa5, 0x54, 0x4a, 0xde, 0x51, 0xd4, 0x77, 0xeb,
	0xed, 0xf7, 0x88, 0x8b, 0x92, 0x1a, 0xc6, 0xb2, 0x36, 0x84, 0xa0, 0x02, 0x3b, 0x43, 0x29, 0x04,
	0x4c, 0x2a, 0x2c, 0xc9, 0xb5, 0x7c, 0xeb, 0xca, 0x49, 0xfb, 0x5a, 0x3c, 0x81, 0x29, 0x6e, 0xf9,
	0x87, 0xd2, 0xee, 0xbd, 0x1e, 0x1d, 0x3a, 0x71, 0x0d, 0xf3, 0x5b, 0x4d, 0xc8, 0x94, 0x77, 0x62,
	0x6e, 0xec, 0x5b, 0x57, 0xf3, 0xd8, 0x0b, 0x8d, 0x53, 0x38, 0x3a, 0x85, 0xd9, 0xe8, 0x94, 0x82,
	0xa1, 0x77, 0x40, 0xf0, 0x0e, 0x96, 0x37, 0xa6, 0x43, 0x99, 0xd2, 0xaf, 0x2d, 0x35, 0x2c, 0x5e,
	0xc1, 0x8c, 0x51, 0xe6, 0x87, 0x00, 0x89, 0xf3, 0x37, 0x99, 0xea, 0xc9, 0xd2, 0x76, 0x17, 0xe9,
	0x05, 0xa3, 0xfc, 0x84, 0x25, 0x05, 0x6f, 0xe1, 0xe1, 0xd1, 0xcb, 0xa6, 0x56, 0x55, 0x43, 0xc2,
	0x07, 0x9b, 0x51, 0xf6, 0xaf, 0xe6, 0xf1, 0x83, 0x70, 0x7f, 0x91, 0xb0, 0x23, 0x75, 0xa3, 0xe0,
	0x03, 0x2c, 0x3e, 0x16, 0x0d, 0x67, 0x28, 0x9b, 0xd1, 0xef, 0x19, 0xcc, 0x94, 0xde, 0x90, 0xce,
	0xd7, 0xbb, 0x61, 0xe1, 0x8b, 0xbe, 0x4f, 0x76, 0xe2, 0x05, 0x40, 0x8d, 0x92, 0x72, 0x56, 0x3f,
	0xa9, 0x1a, 0xf6, 0x76, 0x3a, 0x24, 0xeb, 0x80, 0xe0, 0x1b, 0x2c, 0x0f, 0x62, 0x43, 0x84, 0x00,
	0x26, 0x8c, 0xb2, 0x71, 0x2d, 0xdf, 0x3e, 0x93, 0xa1, 0x9f, 0x89, 0xd7, 0xb0, 0xa8, 0xe8, 0x37,
	0xe7, 0xff, 0x69, 0xdf, 0xef, 0xe0, 0xcf, 0xa3, 0x7e, 0xfc, 0xc7, 0x02, 0xc8, 0x50, 0x7e, 0x21,
	0xdd, 0x16, 0xb7, 0x24, 0xde, 0x83, 0xb3, 0x5f, 0x59, 0x5c, 0x1e, 0x29, 0x9f, 0x9e, 0xd0, 0x7b,
	0x7e, 0x7e, 0x38, 0x44, 0xbc, 0x81, 0xd9, 0x18, 0x5b, 0x78, 0x47, 0xcc, 0x93, 0xc3, 0x78, 0x97,
	0x67, 0x67, 0x46, 0x24, 0x79, 0xfc, 0x55, 0xec, 0xff, 0xc2, 0x6b, 0x53, 0xb5, 0xab, 0xf5, 0xb4,
	0xff, 0xde, 0x6f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x28, 0x3f, 0x3e, 0x0c, 0xa2, 0x02, 0x00,
	0x00,
}
