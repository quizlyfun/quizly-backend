// Code generated by protoc-gen-twirp v8.1.3, DO NOT EDIT.
// source: editor/v1/round_question.proto

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

// ==============================
// RoundQuestionService Interface
// ==============================

type RoundQuestionService interface {
	// CreateRoundQuestion adds question for topic in pack round.
	CreateRoundQuestion(context.Context, *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error)

	// GetRoundQuestion returns round question.
	GetRoundQuestion(context.Context, *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error)
}

// ====================================
// RoundQuestionService Protobuf Client
// ====================================

type roundQuestionServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewRoundQuestionServiceProtobufClient creates a Protobuf client that implements the RoundQuestionService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewRoundQuestionServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) RoundQuestionService {
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
	serviceURL += baseServicePath(pathPrefix, "editor.v1", "RoundQuestionService")
	urls := [2]string{
		serviceURL + "CreateRoundQuestion",
		serviceURL + "GetRoundQuestion",
	}

	return &roundQuestionServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *roundQuestionServiceProtobufClient) CreateRoundQuestion(ctx context.Context, in *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "RoundQuestionService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateRoundQuestion")
	caller := c.callCreateRoundQuestion
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRoundQuestionRequest) when calling interceptor")
					}
					return c.callCreateRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *roundQuestionServiceProtobufClient) callCreateRoundQuestion(ctx context.Context, in *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
	out := new(CreateRoundQuestionResponse)
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

func (c *roundQuestionServiceProtobufClient) GetRoundQuestion(ctx context.Context, in *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "RoundQuestionService")
	ctx = ctxsetters.WithMethodName(ctx, "GetRoundQuestion")
	caller := c.callGetRoundQuestion
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetRoundQuestionRequest) when calling interceptor")
					}
					return c.callGetRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *roundQuestionServiceProtobufClient) callGetRoundQuestion(ctx context.Context, in *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
	out := new(GetRoundQuestionResponse)
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

// ================================
// RoundQuestionService JSON Client
// ================================

type roundQuestionServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewRoundQuestionServiceJSONClient creates a JSON client that implements the RoundQuestionService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewRoundQuestionServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) RoundQuestionService {
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
	serviceURL += baseServicePath(pathPrefix, "editor.v1", "RoundQuestionService")
	urls := [2]string{
		serviceURL + "CreateRoundQuestion",
		serviceURL + "GetRoundQuestion",
	}

	return &roundQuestionServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *roundQuestionServiceJSONClient) CreateRoundQuestion(ctx context.Context, in *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "RoundQuestionService")
	ctx = ctxsetters.WithMethodName(ctx, "CreateRoundQuestion")
	caller := c.callCreateRoundQuestion
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRoundQuestionRequest) when calling interceptor")
					}
					return c.callCreateRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *roundQuestionServiceJSONClient) callCreateRoundQuestion(ctx context.Context, in *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
	out := new(CreateRoundQuestionResponse)
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

func (c *roundQuestionServiceJSONClient) GetRoundQuestion(ctx context.Context, in *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "RoundQuestionService")
	ctx = ctxsetters.WithMethodName(ctx, "GetRoundQuestion")
	caller := c.callGetRoundQuestion
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetRoundQuestionRequest) when calling interceptor")
					}
					return c.callGetRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *roundQuestionServiceJSONClient) callGetRoundQuestion(ctx context.Context, in *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
	out := new(GetRoundQuestionResponse)
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

// ===================================
// RoundQuestionService Server Handler
// ===================================

type roundQuestionServiceServer struct {
	RoundQuestionService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewRoundQuestionServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewRoundQuestionServiceServer(svc RoundQuestionService, opts ...interface{}) TwirpServer {
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

	return &roundQuestionServiceServer{
		RoundQuestionService: svc,
		hooks:                serverOpts.Hooks,
		interceptor:          twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:           pathPrefix,
		jsonSkipDefaults:     jsonSkipDefaults,
		jsonCamelCase:        jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *roundQuestionServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *roundQuestionServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
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

// RoundQuestionServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const RoundQuestionServicePathPrefix = "/twirp/editor.v1.RoundQuestionService/"

func (s *roundQuestionServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "editor.v1")
	ctx = ctxsetters.WithServiceName(ctx, "RoundQuestionService")
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
	if pkgService != "editor.v1.RoundQuestionService" {
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
	case "CreateRoundQuestion":
		s.serveCreateRoundQuestion(ctx, resp, req)
		return
	case "GetRoundQuestion":
		s.serveGetRoundQuestion(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *roundQuestionServiceServer) serveCreateRoundQuestion(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateRoundQuestionJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateRoundQuestionProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *roundQuestionServiceServer) serveCreateRoundQuestionJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateRoundQuestion")
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
	reqContent := new(CreateRoundQuestionRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.RoundQuestionService.CreateRoundQuestion
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRoundQuestionRequest) when calling interceptor")
					}
					return s.RoundQuestionService.CreateRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateRoundQuestionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateRoundQuestionResponse and nil error while calling CreateRoundQuestion. nil responses are not supported"))
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

func (s *roundQuestionServiceServer) serveCreateRoundQuestionProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "CreateRoundQuestion")
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
	reqContent := new(CreateRoundQuestionRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.RoundQuestionService.CreateRoundQuestion
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *CreateRoundQuestionRequest) (*CreateRoundQuestionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*CreateRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*CreateRoundQuestionRequest) when calling interceptor")
					}
					return s.RoundQuestionService.CreateRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*CreateRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*CreateRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *CreateRoundQuestionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *CreateRoundQuestionResponse and nil error while calling CreateRoundQuestion. nil responses are not supported"))
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

func (s *roundQuestionServiceServer) serveGetRoundQuestion(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetRoundQuestionJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetRoundQuestionProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *roundQuestionServiceServer) serveGetRoundQuestionJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetRoundQuestion")
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
	reqContent := new(GetRoundQuestionRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.RoundQuestionService.GetRoundQuestion
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetRoundQuestionRequest) when calling interceptor")
					}
					return s.RoundQuestionService.GetRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *GetRoundQuestionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *GetRoundQuestionResponse and nil error while calling GetRoundQuestion. nil responses are not supported"))
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

func (s *roundQuestionServiceServer) serveGetRoundQuestionProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetRoundQuestion")
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
	reqContent := new(GetRoundQuestionRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.RoundQuestionService.GetRoundQuestion
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *GetRoundQuestionRequest) (*GetRoundQuestionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*GetRoundQuestionRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*GetRoundQuestionRequest) when calling interceptor")
					}
					return s.RoundQuestionService.GetRoundQuestion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*GetRoundQuestionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*GetRoundQuestionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *GetRoundQuestionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *GetRoundQuestionResponse and nil error while calling GetRoundQuestion. nil responses are not supported"))
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

func (s *roundQuestionServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor4, 0
}

func (s *roundQuestionServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.3"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *roundQuestionServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "editor.v1", "RoundQuestionService")
}

var twirpFileDescriptor4 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xc7, 0x6b, 0xe7, 0xcb, 0x39, 0x49, 0x16, 0x77, 0xa8, 0xb4, 0xb3, 0x29, 0x74, 0xd3, 0xac,
	0x40, 0xd1, 0x0a, 0x39, 0xda, 0x70, 0x81, 0x04, 0x45, 0xc8, 0x4e, 0x1c, 0x64, 0xb5, 0xca, 0x6e,
	0xc7, 0x0e, 0x12, 0xf4, 0xc2, 0xf2, 0xc6, 0xd3, 0x62, 0x91, 0xc4, 0xa9, 0x3d, 0x09, 0x54, 0xdc,
	0xf5, 0x31, 0x78, 0x04, 0x5e, 0x86, 0x6b, 0x5e, 0x65, 0xaf, 0x90, 0x67, 0x9c, 0xd4, 0xf6, 0x6e,
	0x36, 0x42, 0xbd, 0x3b, 0xe7, 0xcc, 0x99, 0xff, 0x39, 0x73, 0xe6, 0x37, 0x36, 0x3c, 0xa1, 0x7e,
	0xc0, 0xc2, 0xa8, 0xbf, 0xb9, 0xe8, 0x47, 0xe1, 0x7a, 0xe9, 0xbb, 0x6f, 0xd7, 0x34, 0x66, 0x41,
	0xb8, 0xd4, 0x56, 0x51, 0xc8, 0x42, 0x54, 0x17, 0xeb, 0xda, 0xe6, 0xa2, 0x7d, 0xbc, 0xf1, 0xe6,
	0x81, 0xef, 0x31, 0xda, 0xdf, 0x1a, 0x22, 0xa7, 0xfd, 0xe4, 0x4d, 0x18, 0xbe, 0x99, 0xd3, 0x3e,
	0xf7, 0xae, 0xd7, 0xaf, 0xfb, 0xfe, 0x3a, 0xf2, 0x3e, 0x68, 0x74, 0xff, 0xa9, 0x40, 0x8b, 0x24,
	0xe2, 0x2f, 0x53, 0x6d, 0x74, 0x04, 0x72, 0xe0, 0x63, 0xa9, 0x23, 0xf5, 0x2a, 0x44, 0x0e, 0x7c,
	0x74, 0x02, 0x8a, 0xa8, 0x1e, 0xf8, 0x58, 0xe6, 0xd1, 0x1a, 0xf7, 0x2d, 0xbe, 0xc4, 0xc2, 0x55,
	0x30, 0x4b, 0x96, 0x4a, 0x62, 0x89, 0xfb, 0x96, 0x8f, 0xbe, 0x07, 0x65, 0xdb, 0x2d, 0x2e, 0x77,
	0xa4, 0x5e, 0x63, 0xf0, 0x54, 0xdb, 0xb5, 0xab, 0xe5, 0x2a, 0x6a, 0x5b, 0x83, 0xec, 0xb6, 0x20,
	0x1d, 0x5a, 0x5b, 0xdb, 0x65, 0xef, 0x56, 0x14, 0x57, 0x3a, 0x52, 0xef, 0x68, 0xf0, 0xd9, 0x3e,
	0x0d, 0xe7, 0xdd, 0x8a, 0x92, 0xe6, 0xdb, 0x8c, 0x87, 0xce, 0x32, 0x12, 0xb3, 0x30, 0x66, 0xb8,
	0xca, 0x3b, 0xdc, 0x25, 0x0d, 0xc3, 0x98, 0xa1, 0x6f, 0xa0, 0xea, 0x2d, 0xe3, 0xdf, 0x69, 0x84,
	0x6b, 0xbc, 0xc9, 0xd3, 0xbd, 0x4d, 0xea, 0x3c, 0x8d, 0xa4, 0xe9, 0xe8, 0x5b, 0x68, 0x08, 0xcb,
	0x65, 0xc1, 0x82, 0x62, 0x85, 0xef, 0x3e, 0xd1, 0xc4, 0xb4, 0xb5, 0xed, 0xb4, 0xb5, 0x51, 0x3a,
	0x6d, 0x02, 0x22, 0xdb, 0x09, 0x16, 0x14, 0x3d, 0x85, 0xe6, 0xaf, 0x61, 0xcc, 0xdc, 0x59, 0xb8,
	0x58, 0xd0, 0x25, 0xc3, 0xf5, 0x8e, 0xd4, 0xab, 0x93, 0x46, 0x12, 0x1b, 0x8a, 0x50, 0x92, 0x12,
	0xd3, 0x59, 0x44, 0x99, 0xcb, 0x07, 0x8a, 0x41, 0xa4, 0x88, 0x98, 0x93, 0x84, 0xd0, 0x29, 0xa4,
	0xae, 0x38, 0x5d, 0x83, 0x9f, 0x0e, 0x44, 0x88, 0x9f, 0xed, 0x19, 0xb4, 0x58, 0xe4, 0x2d, 0xe3,
	0xd7, 0x49, 0x93, 0xc9, 0x0c, 0x9b, 0x7c, 0x86, 0xc7, 0x99, 0x23, 0x3a, 0xe9, 0xba, 0x18, 0x1f,
	0xcb, 0x78, 0x89, 0x7c, 0x10, 0xbb, 0xbf, 0x51, 0xba, 0xf2, 0xae, 0xe7, 0x14, 0xb7, 0x3a, 0x52,
	0x4f, 0x21, 0x10, 0xc4, 0xcf, 0xd3, 0x48, 0xfb, 0x39, 0x28, 0x7b, 0x99, 0x41, 0x50, 0x66, 0xf4,
	0x0f, 0xc6, 0x79, 0xa9, 0x13, 0x6e, 0xa3, 0xc7, 0x50, 0x5f, 0x50, 0x3f, 0xf0, 0xdc, 0x75, 0x34,
	0xe7, 0xb4, 0xd4, 0x89, 0xc2, 0x03, 0xd3, 0x68, 0xde, 0xb6, 0xa0, 0x2a, 0x06, 0xfc, 0xd1, 0x52,
	0xdd, 0xbf, 0xca, 0xd0, 0x1e, 0x46, 0xd4, 0x63, 0x34, 0x77, 0x81, 0x84, 0xf2, 0x8b, 0x4f, 0xce,
	0xb5, 0xc3, 0x62, 0x57, 0x08, 0xb6, 0xa1, 0x02, 0xd4, 0x72, 0x1e, 0xea, 0xec, 0x53, 0x28, 0xe5,
	0x9f, 0x82, 0x5d, 0x04, 0xb6, 0x7c, 0x18, 0x58, 0x43, 0xbd, 0x31, 0x5a, 0xef, 0x25, 0xc0, 0x12,
	0x96, 0x71, 0x09, 0x97, 0x71, 0xa5, 0x80, 0xf0, 0x57, 0x45, 0x84, 0x93, 0x57, 0x50, 0x31, 0x6a,
	0x37, 0x46, 0xb9, 0x2d, 0xf7, 0xa4, 0x02, 0xcb, 0x2f, 0xf2, 0x48, 0x56, 0x0f, 0x20, 0xc9, 0xab,
	0xff, 0x2d, 0x81, 0x22, 0x75, 0x65, 0xe5, 0xd9, 0x40, 0x56, 0x2a, 0xf7, 0x42, 0x5a, 0x3b, 0x0c,
	0xa9, 0x72, 0x10, 0xd2, 0xfa, 0x2d, 0x48, 0x0b, 0x98, 0x41, 0x11, 0x33, 0xf4, 0xa2, 0x48, 0x71,
	0xe3, 0x5e, 0x8a, 0x8d, 0xa3, 0x1b, 0xa3, 0xf1, 0x5e, 0x52, 0xf0, 0x03, 0x31, 0xd5, 0x3c, 0xd5,
	0x5d, 0x0b, 0x1e, 0xdf, 0xc9, 0x46, 0xbc, 0x0a, 0x97, 0x31, 0x45, 0xe7, 0xf0, 0x30, 0xff, 0xa5,
	0xfd, 0x80, 0xc8, 0x27, 0x51, 0x76, 0x87, 0xe5, 0x77, 0x4d, 0x38, 0xfe, 0x91, 0xb2, 0x3b, 0x19,
	0xfb, 0x3f, 0x32, 0xaf, 0x00, 0xdf, 0x96, 0x49, 0xdb, 0xf9, 0x01, 0x8e, 0xf2, 0x3a, 0x5c, 0xa4,
	0x31, 0xc0, 0xfb, 0xa8, 0x22, 0xad, 0x9c, 0xfc, 0xf9, 0x25, 0x34, 0xb3, 0xc3, 0x41, 0x9f, 0xc3,
	0x89, 0x43, 0xf4, 0x89, 0x3d, 0x36, 0x89, 0xeb, 0xfc, 0x7c, 0x65, 0xba, 0xd3, 0x89, 0x7d, 0x65,
	0x0e, 0xad, 0xb1, 0x65, 0x8e, 0xd4, 0x07, 0x08, 0xa0, 0x6a, 0x98, 0xe3, 0x4b, 0x62, 0xaa, 0x12,
	0xaa, 0x43, 0x45, 0x1f, 0x3b, 0x26, 0x51, 0xe5, 0xc4, 0x9c, 0x98, 0x3f, 0x99, 0x44, 0x2d, 0x9d,
	0xff, 0x09, 0x0f, 0x6f, 0x61, 0x8c, 0xce, 0xe0, 0x94, 0x5c, 0x4e, 0x27, 0x23, 0xf7, 0xe5, 0xd4,
	0xb4, 0x1d, 0xeb, 0x72, 0x72, 0x97, 0x76, 0x13, 0x14, 0xdb, 0xd1, 0x27, 0x23, 0x9d, 0x8c, 0x54,
	0x09, 0x29, 0x50, 0xb6, 0xf5, 0xb1, 0xa9, 0xca, 0x49, 0x4d, 0xdb, 0x1c, 0x12, 0xd3, 0x51, 0x4b,
	0x48, 0x85, 0xa6, 0x3d, 0xbd, 0x32, 0x89, 0x9b, 0x46, 0xca, 0xa8, 0x01, 0x35, 0x7d, 0x3a, 0x4c,
	0x34, 0xd5, 0xca, 0xe0, 0x5f, 0x09, 0x1e, 0xe5, 0xaa, 0xdb, 0x34, 0xda, 0x04, 0x33, 0x8a, 0x7c,
	0xf8, 0xf4, 0x8e, 0x5b, 0x45, 0x5f, 0x64, 0xc6, 0xb4, 0xff, 0x8b, 0xd0, 0xfe, 0xf2, 0x50, 0x5a,
	0x7a, 0x1b, 0xaf, 0x40, 0x2d, 0xde, 0x14, 0xea, 0x66, 0xf6, 0xee, 0xa1, 0xa1, 0x7d, 0x76, 0x6f,
	0x8e, 0x10, 0x37, 0x1e, 0xfd, 0x82, 0x76, 0x7f, 0xfb, 0xef, 0x84, 0xb5, 0xb9, 0xb8, 0xae, 0xf2,
	0x57, 0xfb, 0xf5, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x22, 0xbf, 0x0f, 0x0a, 0x08, 0x00,
	0x00,
}
