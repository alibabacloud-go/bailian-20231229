// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCategoryRequest struct {
	// This parameter is required.
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s AddCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddCategoryRequest) SetCategoryName(v string) *AddCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *AddCategoryRequest) SetCategoryType(v string) *AddCategoryRequest {
	s.CategoryType = &v
	return s
}

func (s *AddCategoryRequest) SetParentCategoryId(v string) *AddCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

type AddCategoryResponseBody struct {
	// example:
	//
	// success
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 778C0B3B-03C1-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBody) SetCode(v string) *AddCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *AddCategoryResponseBody) SetData(v *AddCategoryResponseBodyData) *AddCategoryResponseBody {
	s.Data = v
	return s
}

func (s *AddCategoryResponseBody) SetMessage(v string) *AddCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *AddCategoryResponseBody) SetRequestId(v string) *AddCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCategoryResponseBody) SetStatus(v string) *AddCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *AddCategoryResponseBody) SetSuccess(v bool) *AddCategoryResponseBody {
	s.Success = &v
	return s
}

type AddCategoryResponseBodyData struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
}

func (s AddCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBodyData) SetCategoryId(v string) *AddCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *AddCategoryResponseBodyData) SetCategoryName(v string) *AddCategoryResponseBodyData {
	s.CategoryName = &v
	return s
}

type AddCategoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCategoryResponse) GoString() string {
	return s.String()
}

func (s *AddCategoryResponse) SetHeaders(v map[string]*string) *AddCategoryResponse {
	s.Headers = v
	return s
}

func (s *AddCategoryResponse) SetStatusCode(v int32) *AddCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCategoryResponse) SetBody(v *AddCategoryResponseBody) *AddCategoryResponse {
	s.Body = v
	return s
}

type AddFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847
	LeaseId *string `json:"LeaseId,omitempty" xml:"LeaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFileRequest) GoString() string {
	return s.String()
}

func (s *AddFileRequest) SetCategoryId(v string) *AddFileRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFileRequest) SetLeaseId(v string) *AddFileRequest {
	s.LeaseId = &v
	return s
}

func (s *AddFileRequest) SetParser(v string) *AddFileRequest {
	s.Parser = &v
	return s
}

type AddFileResponseBody struct {
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AddFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBody) GoString() string {
	return s.String()
}

func (s *AddFileResponseBody) SetCode(v string) *AddFileResponseBody {
	s.Code = &v
	return s
}

func (s *AddFileResponseBody) SetData(v *AddFileResponseBodyData) *AddFileResponseBody {
	s.Data = v
	return s
}

func (s *AddFileResponseBody) SetMessage(v string) *AddFileResponseBody {
	s.Message = &v
	return s
}

func (s *AddFileResponseBody) SetRequestId(v string) *AddFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFileResponseBody) SetStatus(v string) *AddFileResponseBody {
	s.Status = &v
	return s
}

func (s *AddFileResponseBody) SetSuccess(v string) *AddFileResponseBody {
	s.Success = &v
	return s
}

type AddFileResponseBodyData struct {
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
}

func (s AddFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddFileResponseBodyData) SetFileId(v string) *AddFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *AddFileResponseBodyData) SetParser(v string) *AddFileResponseBodyData {
	s.Parser = &v
	return s
}

type AddFileResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFileResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFileResponse) GoString() string {
	return s.String()
}

func (s *AddFileResponse) SetHeaders(v map[string]*string) *AddFileResponse {
	s.Headers = v
	return s
}

func (s *AddFileResponse) SetStatusCode(v int32) *AddFileResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFileResponse) SetBody(v *AddFileResponseBody) *AddFileResponse {
	s.Body = v
	return s
}

type ApplyFileUploadLeaseRequest struct {
	// This parameter is required.
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19657c391f6c70bcea63c154d8606bb3
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	SizeInBytes *string `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
}

func (s ApplyFileUploadLeaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseRequest) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseRequest) SetFileName(v string) *ApplyFileUploadLeaseRequest {
	s.FileName = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetMd5(v string) *ApplyFileUploadLeaseRequest {
	s.Md5 = &v
	return s
}

func (s *ApplyFileUploadLeaseRequest) SetSizeInBytes(v string) *ApplyFileUploadLeaseRequest {
	s.SizeInBytes = &v
	return s
}

type ApplyFileUploadLeaseResponseBody struct {
	// example:
	//
	// DataCenter.FileTooLarge
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ApplyFileUploadLeaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 778C0B3B-xxxx-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBody) SetCode(v string) *ApplyFileUploadLeaseResponseBody {
	s.Code = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetData(v *ApplyFileUploadLeaseResponseBodyData) *ApplyFileUploadLeaseResponseBody {
	s.Data = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetMessage(v string) *ApplyFileUploadLeaseResponseBody {
	s.Message = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetRequestId(v string) *ApplyFileUploadLeaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetStatus(v string) *ApplyFileUploadLeaseResponseBody {
	s.Status = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBody) SetSuccess(v bool) *ApplyFileUploadLeaseResponseBody {
	s.Success = &v
	return s
}

type ApplyFileUploadLeaseResponseBodyData struct {
	// example:
	//
	// 1e6a159107384782be5e45ac4759b247.1719325231035
	FileUploadLeaseId *string                                    `json:"FileUploadLeaseId,omitempty" xml:"FileUploadLeaseId,omitempty"`
	Param             *ApplyFileUploadLeaseResponseBodyDataParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// example:
	//
	// HTTP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetFileUploadLeaseId(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.FileUploadLeaseId = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetParam(v *ApplyFileUploadLeaseResponseBodyDataParam) *ApplyFileUploadLeaseResponseBodyData {
	s.Param = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyData) SetType(v string) *ApplyFileUploadLeaseResponseBodyData {
	s.Type = &v
	return s
}

type ApplyFileUploadLeaseResponseBodyDataParam struct {
	// example:
	//
	// "X-bailian-extra": "MTAwNTQyNjQ5NTE2OTE3OA==",
	//
	//         "Content-Type": "application/pdf"
	Headers interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// PUT
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// https://bailian-datahub-data-origin-prod.oss-cn-hangzhou.aliyuncs.com/1005426495169178/10024405/68abd1dea7b6404d8f7d7b9f7fbd332d.1716698936847.pdf?Expires=1716699536&OSSAccessKeyId=TestID&Signature=HfwPUZo4pR6DatSDym0zFKVh9Wg%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponseBodyDataParam) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetHeaders(v interface{}) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetMethod(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Method = &v
	return s
}

func (s *ApplyFileUploadLeaseResponseBodyDataParam) SetUrl(v string) *ApplyFileUploadLeaseResponseBodyDataParam {
	s.Url = &v
	return s
}

type ApplyFileUploadLeaseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFileUploadLeaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFileUploadLeaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyFileUploadLeaseResponse) GoString() string {
	return s.String()
}

func (s *ApplyFileUploadLeaseResponse) SetHeaders(v map[string]*string) *ApplyFileUploadLeaseResponse {
	s.Headers = v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetStatusCode(v int32) *ApplyFileUploadLeaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFileUploadLeaseResponse) SetBody(v *ApplyFileUploadLeaseResponseBody) *ApplyFileUploadLeaseResponse {
	s.Body = v
	return s
}

type CreateAndPulishAgentRequest struct {
	ApplicationConfig *CreateAndPulishAgentRequestApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Instructions      *string                                       `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                       `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAndPulishAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequest) SetApplicationConfig(v *CreateAndPulishAgentRequestApplicationConfig) *CreateAndPulishAgentRequest {
	s.ApplicationConfig = v
	return s
}

func (s *CreateAndPulishAgentRequest) SetInstructions(v string) *CreateAndPulishAgentRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAndPulishAgentRequest) SetModelId(v string) *CreateAndPulishAgentRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAndPulishAgentRequest) SetName(v string) *CreateAndPulishAgentRequest {
	s.Name = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfig struct {
	HistoryConfig  *CreateAndPulishAgentRequestApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *CreateAndPulishAgentRequestApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *CreateAndPulishAgentRequestApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *CreateAndPulishAgentRequestApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	SecurityConfig *CreateAndPulishAgentRequestApplicationConfigSecurityConfig `json:"securityConfig,omitempty" xml:"securityConfig,omitempty" type:"Struct"`
	Tools          []*CreateAndPulishAgentRequestApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*CreateAndPulishAgentRequestApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s CreateAndPulishAgentRequestApplicationConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetHistoryConfig(v *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) *CreateAndPulishAgentRequestApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetLongTermMemory(v *CreateAndPulishAgentRequestApplicationConfigLongTermMemory) *CreateAndPulishAgentRequestApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetParameters(v *CreateAndPulishAgentRequestApplicationConfigParameters) *CreateAndPulishAgentRequestApplicationConfig {
	s.Parameters = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetRagConfig(v *CreateAndPulishAgentRequestApplicationConfigRagConfig) *CreateAndPulishAgentRequestApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetSecurityConfig(v *CreateAndPulishAgentRequestApplicationConfigSecurityConfig) *CreateAndPulishAgentRequestApplicationConfig {
	s.SecurityConfig = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetTools(v []*CreateAndPulishAgentRequestApplicationConfigTools) *CreateAndPulishAgentRequestApplicationConfig {
	s.Tools = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfig) SetWorkFlows(v []*CreateAndPulishAgentRequestApplicationConfigWorkFlows) *CreateAndPulishAgentRequestApplicationConfig {
	s.WorkFlows = v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigHistoryConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetEnableRecord(v bool) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetInstanceId(v string) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetRegion(v string) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigHistoryConfig) SetStoreCode(v string) *CreateAndPulishAgentRequestApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigLongTermMemory) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigLongTermMemory) SetEnable(v bool) *CreateAndPulishAgentRequestApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigParameters struct {
	DialogRound *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	MaxTokens   *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetDialogRound(v int32) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetMaxTokens(v int32) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigParameters) SetTemperature(v float64) *CreateAndPulishAgentRequestApplicationConfigParameters {
	s.Temperature = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigRagConfig struct {
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigRagConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetEnableCitation(v bool) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetEnableSearch(v bool) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *CreateAndPulishAgentRequestApplicationConfigRagConfig) SetTopK(v int32) *CreateAndPulishAgentRequestApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigSecurityConfig struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigSecurityConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigSecurityConfig) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigSecurityConfig) SetProcessingStrategy(v string) *CreateAndPulishAgentRequestApplicationConfigSecurityConfig {
	s.ProcessingStrategy = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigTools) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigTools) SetType(v string) *CreateAndPulishAgentRequestApplicationConfigTools {
	s.Type = &v
	return s
}

type CreateAndPulishAgentRequestApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAndPulishAgentRequestApplicationConfigWorkFlows) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentRequestApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentRequestApplicationConfigWorkFlows) SetType(v string) *CreateAndPulishAgentRequestApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

type CreateAndPulishAgentShrinkRequest struct {
	ApplicationConfigShrink *string `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty"`
	Instructions            *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId                 *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAndPulishAgentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentShrinkRequest) SetApplicationConfigShrink(v string) *CreateAndPulishAgentShrinkRequest {
	s.ApplicationConfigShrink = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetInstructions(v string) *CreateAndPulishAgentShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetModelId(v string) *CreateAndPulishAgentShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetName(v string) *CreateAndPulishAgentShrinkRequest {
	s.Name = &v
	return s
}

type CreateAndPulishAgentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAndPulishAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentResponseBody) SetCode(v string) *CreateAndPulishAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetData(v string) *CreateAndPulishAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetHttpStatusCode(v int32) *CreateAndPulishAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetMessage(v string) *CreateAndPulishAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetRequestId(v string) *CreateAndPulishAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetSuccess(v bool) *CreateAndPulishAgentResponseBody {
	s.Success = &v
	return s
}

type CreateAndPulishAgentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndPulishAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndPulishAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAndPulishAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentResponse) SetHeaders(v map[string]*string) *CreateAndPulishAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAndPulishAgentResponse) SetStatusCode(v int32) *CreateAndPulishAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndPulishAgentResponse) SetBody(v *CreateAndPulishAgentResponseBody) *CreateAndPulishAgentResponse {
	s.Body = v
	return s
}

type CreateIndexRequest struct {
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// example:
	//
	// 128
	ChunkSize   *int32                        `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	Columns     []*CreateIndexRequestColumns  `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	DataSource  *CreateIndexRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	Description *string                       `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentIds []*string                     `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// example:
	//
	// ,
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// example:
	//
	// gp-bp321093j84
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// structured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s CreateIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) SetCategoryIds(v []*string) *CreateIndexRequest {
	s.CategoryIds = v
	return s
}

func (s *CreateIndexRequest) SetChunkSize(v int32) *CreateIndexRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateIndexRequest) SetColumns(v []*CreateIndexRequestColumns) *CreateIndexRequest {
	s.Columns = v
	return s
}

func (s *CreateIndexRequest) SetDataSource(v *CreateIndexRequestDataSource) *CreateIndexRequest {
	s.DataSource = v
	return s
}

func (s *CreateIndexRequest) SetDescription(v string) *CreateIndexRequest {
	s.Description = &v
	return s
}

func (s *CreateIndexRequest) SetDocumentIds(v []*string) *CreateIndexRequest {
	s.DocumentIds = v
	return s
}

func (s *CreateIndexRequest) SetEmbeddingModelName(v string) *CreateIndexRequest {
	s.EmbeddingModelName = &v
	return s
}

func (s *CreateIndexRequest) SetName(v string) *CreateIndexRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexRequest) SetOverlapSize(v int32) *CreateIndexRequest {
	s.OverlapSize = &v
	return s
}

func (s *CreateIndexRequest) SetRerankMinScore(v float64) *CreateIndexRequest {
	s.RerankMinScore = &v
	return s
}

func (s *CreateIndexRequest) SetRerankModelName(v string) *CreateIndexRequest {
	s.RerankModelName = &v
	return s
}

func (s *CreateIndexRequest) SetSeparator(v string) *CreateIndexRequest {
	s.Separator = &v
	return s
}

func (s *CreateIndexRequest) SetSinkInstanceId(v string) *CreateIndexRequest {
	s.SinkInstanceId = &v
	return s
}

func (s *CreateIndexRequest) SetSinkRegion(v string) *CreateIndexRequest {
	s.SinkRegion = &v
	return s
}

func (s *CreateIndexRequest) SetSinkType(v string) *CreateIndexRequest {
	s.SinkType = &v
	return s
}

func (s *CreateIndexRequest) SetSourceType(v string) *CreateIndexRequest {
	s.SourceType = &v
	return s
}

func (s *CreateIndexRequest) SetStructureType(v string) *CreateIndexRequest {
	s.StructureType = &v
	return s
}

type CreateIndexRequestColumns struct {
	Column   *string `json:"Column,omitempty" xml:"Column,omitempty"`
	IsRecall *bool   `json:"IsRecall,omitempty" xml:"IsRecall,omitempty"`
	IsSearch *bool   `json:"IsSearch,omitempty" xml:"IsSearch,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIndexRequestColumns) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestColumns) SetColumn(v string) *CreateIndexRequestColumns {
	s.Column = &v
	return s
}

func (s *CreateIndexRequestColumns) SetIsRecall(v bool) *CreateIndexRequestColumns {
	s.IsRecall = &v
	return s
}

func (s *CreateIndexRequestColumns) SetIsSearch(v bool) *CreateIndexRequestColumns {
	s.IsSearch = &v
	return s
}

func (s *CreateIndexRequestColumns) SetName(v string) *CreateIndexRequestColumns {
	s.Name = &v
	return s
}

func (s *CreateIndexRequestColumns) SetType(v string) *CreateIndexRequestColumns {
	s.Type = &v
	return s
}

type CreateIndexRequestDataSource struct {
	CredentialId  *string `json:"CredentialId,omitempty" xml:"CredentialId,omitempty"`
	CredentialKey *string `json:"CredentialKey,omitempty" xml:"CredentialKey,omitempty"`
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Endpoint      *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	IsPrivateLink *bool   `json:"IsPrivateLink,omitempty" xml:"IsPrivateLink,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	SubPath       *string `json:"SubPath,omitempty" xml:"SubPath,omitempty"`
	SubType       *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	Table         *string `json:"Table,omitempty" xml:"Table,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIndexRequestDataSource) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexRequestDataSource) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSource) SetCredentialId(v string) *CreateIndexRequestDataSource {
	s.CredentialId = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetCredentialKey(v string) *CreateIndexRequestDataSource {
	s.CredentialKey = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetDatabase(v string) *CreateIndexRequestDataSource {
	s.Database = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetEndpoint(v string) *CreateIndexRequestDataSource {
	s.Endpoint = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetIsPrivateLink(v bool) *CreateIndexRequestDataSource {
	s.IsPrivateLink = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetRegion(v string) *CreateIndexRequestDataSource {
	s.Region = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetSubPath(v string) *CreateIndexRequestDataSource {
	s.SubPath = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetSubType(v string) *CreateIndexRequestDataSource {
	s.SubType = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetTable(v string) *CreateIndexRequestDataSource {
	s.Table = &v
	return s
}

func (s *CreateIndexRequestDataSource) SetType(v string) *CreateIndexRequestDataSource {
	s.Type = &v
	return s
}

type CreateIndexShrinkRequest struct {
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	// example:
	//
	// 128
	ChunkSize         *int32  `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	ColumnsShrink     *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	DataSourceShrink  *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// example:
	//
	// text-embedding-v2
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// example:
	//
	// 0.20
	RerankMinScore *float64 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// example:
	//
	// ,
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// example:
	//
	// gp-bp321093j84
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEFAULT
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// structured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s CreateIndexShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexShrinkRequest) SetCategoryIdsShrink(v string) *CreateIndexShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetChunkSize(v int32) *CreateIndexShrinkRequest {
	s.ChunkSize = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetColumnsShrink(v string) *CreateIndexShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDataSourceShrink(v string) *CreateIndexShrinkRequest {
	s.DataSourceShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDescription(v string) *CreateIndexShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetDocumentIdsShrink(v string) *CreateIndexShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetEmbeddingModelName(v string) *CreateIndexShrinkRequest {
	s.EmbeddingModelName = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetName(v string) *CreateIndexShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetOverlapSize(v int32) *CreateIndexShrinkRequest {
	s.OverlapSize = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankMinScore(v float64) *CreateIndexShrinkRequest {
	s.RerankMinScore = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetRerankModelName(v string) *CreateIndexShrinkRequest {
	s.RerankModelName = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSeparator(v string) *CreateIndexShrinkRequest {
	s.Separator = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkInstanceId(v string) *CreateIndexShrinkRequest {
	s.SinkInstanceId = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkRegion(v string) *CreateIndexShrinkRequest {
	s.SinkRegion = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSinkType(v string) *CreateIndexShrinkRequest {
	s.SinkType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetSourceType(v string) *CreateIndexShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateIndexShrinkRequest) SetStructureType(v string) *CreateIndexShrinkRequest {
	s.StructureType = &v
	return s
}

type CreateIndexResponseBody struct {
	// example:
	//
	// Forbidden
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid input, variable name is missing
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) SetCode(v string) *CreateIndexResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIndexResponseBody) SetData(v *CreateIndexResponseBodyData) *CreateIndexResponseBody {
	s.Data = v
	return s
}

func (s *CreateIndexResponseBody) SetMessage(v string) *CreateIndexResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetStatus(v string) *CreateIndexResponseBody {
	s.Status = &v
	return s
}

func (s *CreateIndexResponseBody) SetSuccess(v bool) *CreateIndexResponseBody {
	s.Success = &v
	return s
}

type CreateIndexResponseBodyData struct {
	// example:
	//
	// jkurxhju6b
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateIndexResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBodyData) SetId(v string) *CreateIndexResponseBodyData {
	s.Id = &v
	return s
}

type CreateIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIndexResponse) GoString() string {
	return s.String()
}

func (s *CreateIndexResponse) SetHeaders(v map[string]*string) *CreateIndexResponse {
	s.Headers = v
	return s
}

func (s *CreateIndexResponse) SetStatusCode(v int32) *CreateIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIndexResponse) SetBody(v *CreateIndexResponseBody) *CreateIndexResponse {
	s.Body = v
	return s
}

type CreateMemoryRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateMemoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMemoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRequest) SetDescription(v string) *CreateMemoryRequest {
	s.Description = &v
	return s
}

type CreateMemoryResponseBody struct {
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbd5fc3
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBody) SetMemoryId(v string) *CreateMemoryResponseBody {
	s.MemoryId = &v
	return s
}

func (s *CreateMemoryResponseBody) SetRequestId(v string) *CreateMemoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMemoryResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponse) SetHeaders(v map[string]*string) *CreateMemoryResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryResponse) SetStatusCode(v int32) *CreateMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryResponse) SetBody(v *CreateMemoryResponseBody) *CreateMemoryResponse {
	s.Body = v
	return s
}

type CreateMemoryNodeRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s CreateMemoryNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMemoryNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryNodeRequest) SetContent(v string) *CreateMemoryNodeRequest {
	s.Content = &v
	return s
}

type CreateMemoryNodeResponseBody struct {
	// example:
	//
	// 68de06c95368463a8be4a84efc872cc5
	MemoryNodeId *string `json:"memoryNodeId,omitempty" xml:"memoryNodeId,omitempty"`
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMemoryNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryNodeResponseBody) SetMemoryNodeId(v string) *CreateMemoryNodeResponseBody {
	s.MemoryNodeId = &v
	return s
}

func (s *CreateMemoryNodeResponseBody) SetRequestId(v string) *CreateMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

type CreateMemoryNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryNodeResponse) SetHeaders(v map[string]*string) *CreateMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryNodeResponse) SetStatusCode(v int32) *CreateMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryNodeResponse) SetBody(v *CreateMemoryNodeResponseBody) *CreateMemoryNodeResponse {
	s.Body = v
	return s
}

type DeleteAgentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetData(v string) *DeleteAgentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAgentResponseBody) SetHttpStatusCode(v int32) *DeleteAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

type DeleteAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponse) SetHeaders(v map[string]*string) *DeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentResponse) SetStatusCode(v int32) *DeleteAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentResponse) SetBody(v *DeleteAgentResponseBody) *DeleteAgentResponse {
	s.Body = v
	return s
}

type DeleteCategoryResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *DeleteCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBody) SetCode(v string) *DeleteCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetData(v *DeleteCategoryResponseBodyData) *DeleteCategoryResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCategoryResponseBody) SetMessage(v string) *DeleteCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetRequestId(v string) *DeleteCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetStatus(v string) *DeleteCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetSuccess(v bool) *DeleteCategoryResponseBody {
	s.Success = &v
	return s
}

type DeleteCategoryResponseBodyData struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBodyData) SetCategoryId(v string) *DeleteCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

type DeleteCategoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponse) SetHeaders(v map[string]*string) *DeleteCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteCategoryResponse) SetStatusCode(v int32) *DeleteCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCategoryResponse) SetBody(v *DeleteCategoryResponseBody) *DeleteCategoryResponse {
	s.Body = v
	return s
}

type DeleteFileResponseBody struct {
	// example:
	//
	// InvalidParameter
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Current file status does not support delete.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetCode(v string) *DeleteFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFileResponseBody) SetData(v *DeleteFileResponseBodyData) *DeleteFileResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFileResponseBody) SetMessage(v string) *DeleteFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFileResponseBody) SetRequestId(v string) *DeleteFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileResponseBody) SetStatus(v string) *DeleteFileResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteFileResponseBody) SetSuccess(v bool) *DeleteFileResponseBody {
	s.Success = &v
	return s
}

type DeleteFileResponseBodyData struct {
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s DeleteFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBodyData) SetFileId(v string) *DeleteFileResponseBodyData {
	s.FileId = &v
	return s
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DeleteIndexRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f89ie0fj5a
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s DeleteIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexRequest) SetIndexId(v string) *DeleteIndexRequest {
	s.IndexId = &v
	return s
}

type DeleteIndexResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponseBody) SetCode(v string) *DeleteIndexResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIndexResponseBody) SetMessage(v string) *DeleteIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIndexResponseBody) SetRequestId(v string) *DeleteIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexResponseBody) SetStatus(v string) *DeleteIndexResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndexResponseBody) SetSuccess(v bool) *DeleteIndexResponseBody {
	s.Success = &v
	return s
}

type DeleteIndexResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponse) SetHeaders(v map[string]*string) *DeleteIndexResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexResponse) SetStatusCode(v int32) *DeleteIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexResponse) SetBody(v *DeleteIndexResponseBody) *DeleteIndexResponse {
	s.Body = v
	return s
}

type DeleteIndexDocumentRequest struct {
	// This parameter is required.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s DeleteIndexDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexDocumentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentRequest) SetDocumentIds(v []*string) *DeleteIndexDocumentRequest {
	s.DocumentIds = v
	return s
}

func (s *DeleteIndexDocumentRequest) SetIndexId(v string) *DeleteIndexDocumentRequest {
	s.IndexId = &v
	return s
}

type DeleteIndexDocumentShrinkRequest struct {
	// This parameter is required.
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s DeleteIndexDocumentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentShrinkRequest) SetDocumentIdsShrink(v string) *DeleteIndexDocumentShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *DeleteIndexDocumentShrinkRequest) SetIndexId(v string) *DeleteIndexDocumentShrinkRequest {
	s.IndexId = &v
	return s
}

type DeleteIndexDocumentResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DeleteIndexDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIndexDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentResponseBody) SetCode(v string) *DeleteIndexDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetData(v *DeleteIndexDocumentResponseBodyData) *DeleteIndexDocumentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetMessage(v string) *DeleteIndexDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetRequestId(v string) *DeleteIndexDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetStatus(v string) *DeleteIndexDocumentResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetSuccess(v bool) *DeleteIndexDocumentResponseBody {
	s.Success = &v
	return s
}

type DeleteIndexDocumentResponseBodyData struct {
	DeletedDocument []*string `json:"DeletedDocument,omitempty" xml:"DeletedDocument,omitempty" type:"Repeated"`
}

func (s DeleteIndexDocumentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentResponseBodyData) SetDeletedDocument(v []*string) *DeleteIndexDocumentResponseBodyData {
	s.DeletedDocument = v
	return s
}

type DeleteIndexDocumentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIndexDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIndexDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIndexDocumentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentResponse) SetHeaders(v map[string]*string) *DeleteIndexDocumentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIndexDocumentResponse) SetStatusCode(v int32) *DeleteIndexDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIndexDocumentResponse) SetBody(v *DeleteIndexDocumentResponseBody) *DeleteIndexDocumentResponse {
	s.Body = v
	return s
}

type DeleteMemoryResponseBody struct {
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBody) SetRequestId(v string) *DeleteMemoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponse) SetHeaders(v map[string]*string) *DeleteMemoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryResponse) SetStatusCode(v int32) *DeleteMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryResponse) SetBody(v *DeleteMemoryResponseBody) *DeleteMemoryResponse {
	s.Body = v
	return s
}

type DeleteMemoryNodeResponseBody struct {
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteMemoryNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryNodeResponseBody) SetRequestId(v string) *DeleteMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMemoryNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryNodeResponse) SetHeaders(v map[string]*string) *DeleteMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryNodeResponse) SetStatusCode(v int32) *DeleteMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryNodeResponse) SetBody(v *DeleteMemoryNodeResponseBody) *DeleteMemoryNodeResponse {
	s.Body = v
	return s
}

type DescribeFileResponseBody struct {
	// example:
	//
	// Success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBody) SetCode(v string) *DescribeFileResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeFileResponseBody) SetData(v *DescribeFileResponseBodyData) *DescribeFileResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFileResponseBody) SetMessage(v string) *DescribeFileResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeFileResponseBody) SetRequestId(v string) *DescribeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileResponseBody) SetStatus(v string) *DescribeFileResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFileResponseBody) SetSuccess(v bool) *DescribeFileResponseBody {
	s.Success = &v
	return s
}

type DescribeFileResponseBodyData struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2024-05-26 12:45:43
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// test.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// example:
	//
	// 1234
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// example:
	//
	// PARSE_SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFileResponseBodyData) SetCategoryId(v string) *DescribeFileResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetCreateTime(v string) *DescribeFileResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileId(v string) *DescribeFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileName(v string) *DescribeFileResponseBodyData {
	s.FileName = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetFileType(v string) *DescribeFileResponseBodyData {
	s.FileType = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetParser(v string) *DescribeFileResponseBodyData {
	s.Parser = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetSizeInBytes(v int64) *DescribeFileResponseBodyData {
	s.SizeInBytes = &v
	return s
}

func (s *DescribeFileResponseBodyData) SetStatus(v string) *DescribeFileResponseBodyData {
	s.Status = &v
	return s
}

type DescribeFileResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileResponse) SetHeaders(v map[string]*string) *DescribeFileResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileResponse) SetStatusCode(v int32) *DescribeFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileResponse) SetBody(v *DescribeFileResponseBody) *DescribeFileResponse {
	s.Body = v
	return s
}

type GetIndexJobStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20230718xxxx-146c93bf
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetIndexJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusRequest) SetIndexId(v string) *GetIndexJobStatusRequest {
	s.IndexId = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetJobId(v string) *GetIndexJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetPageNumber(v int32) *GetIndexJobStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *GetIndexJobStatusRequest) SetPageSize(v int32) *GetIndexJobStatusRequest {
	s.PageSize = &v
	return s
}

type GetIndexJobStatusResponseBody struct {
	// example:
	//
	// Index.Forbidden
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetIndexJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetIndexJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBody) SetCode(v string) *GetIndexJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetData(v *GetIndexJobStatusResponseBodyData) *GetIndexJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetMessage(v string) *GetIndexJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetRequestId(v string) *GetIndexJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetStatus(v string) *GetIndexJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetSuccess(v bool) *GetIndexJobStatusResponseBody {
	s.Success = &v
	return s
}

type GetIndexJobStatusResponseBodyData struct {
	Documents []*GetIndexJobStatusResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// 66122af12a4e45ddae6bd6c845556647
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIndexJobStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBodyData) SetDocuments(v []*GetIndexJobStatusResponseBodyDataDocuments) *GetIndexJobStatusResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) SetJobId(v string) *GetIndexJobStatusResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) SetStatus(v string) *GetIndexJobStatusResponseBodyData {
	s.Status = &v
	return s
}

type GetIndexJobStatusResponseBodyDataDocuments struct {
	// example:
	//
	// Index.Document.ChunkError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	DocId   *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	DocName *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	// example:
	//
	// document parse error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIndexJobStatusResponseBodyDataDocuments) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetCode(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Code = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetDocId(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetDocName(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.DocName = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetMessage(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Message = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetStatus(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Status = &v
	return s
}

type GetIndexJobStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIndexJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIndexJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIndexJobStatusResponse) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponse) SetHeaders(v map[string]*string) *GetIndexJobStatusResponse {
	s.Headers = v
	return s
}

func (s *GetIndexJobStatusResponse) SetStatusCode(v int32) *GetIndexJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIndexJobStatusResponse) SetBody(v *GetIndexJobStatusResponseBody) *GetIndexJobStatusResponse {
	s.Body = v
	return s
}

type GetMemoryResponseBody struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbd5fc3
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// llm-us9hjmt32nysdm5v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetMemoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBody) SetDescription(v string) *GetMemoryResponseBody {
	s.Description = &v
	return s
}

func (s *GetMemoryResponseBody) SetMemoryId(v string) *GetMemoryResponseBody {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryResponseBody) SetRequestId(v string) *GetMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryResponseBody) SetWorkspaceId(v string) *GetMemoryResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetMemoryResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemoryResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryResponse) SetHeaders(v map[string]*string) *GetMemoryResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryResponse) SetStatusCode(v int32) *GetMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryResponse) SetBody(v *GetMemoryResponseBody) *GetMemoryResponse {
	s.Body = v
	return s
}

type GetMemoryNodeResponseBody struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 6bff4f317a14442fbc9f73d29dbd5fc3
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// example:
	//
	// 68de06c95368463a8be4a84efc872cc5
	MemoryNodeId *string `json:"memoryNodeId,omitempty" xml:"memoryNodeId,omitempty"`
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// llm-us9hjmt32nysdm5v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetMemoryNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryNodeResponseBody) SetContent(v string) *GetMemoryNodeResponseBody {
	s.Content = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetMemoryId(v string) *GetMemoryNodeResponseBody {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetMemoryNodeId(v string) *GetMemoryNodeResponseBody {
	s.MemoryNodeId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetRequestId(v string) *GetMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryNodeResponseBody) SetWorkspaceId(v string) *GetMemoryNodeResponseBody {
	s.WorkspaceId = &v
	return s
}

type GetMemoryNodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryNodeResponse) SetHeaders(v map[string]*string) *GetMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryNodeResponse) SetStatusCode(v int32) *GetMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryNodeResponse) SetBody(v *GetMemoryNodeResponseBody) *GetMemoryNodeResponse {
	s.Body = v
	return s
}

type GetPublishedAgentResponseBody struct {
	Code           *string                            `json:"code,omitempty" xml:"code,omitempty"`
	Data           *GetPublishedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                            `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                              `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPublishedAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBody) SetCode(v string) *GetPublishedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetData(v *GetPublishedAgentResponseBodyData) *GetPublishedAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetPublishedAgentResponseBody) SetHttpStatusCode(v int32) *GetPublishedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetMessage(v string) *GetPublishedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetRequestId(v string) *GetPublishedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublishedAgentResponseBody) SetSuccess(v bool) *GetPublishedAgentResponseBody {
	s.Success = &v
	return s
}

type GetPublishedAgentResponseBodyData struct {
	ApplicationConfig *GetPublishedAgentResponseBodyDataApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Code              *string                                             `json:"code,omitempty" xml:"code,omitempty"`
	Instructions      *string                                             `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                             `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                             `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetPublishedAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyData) SetApplicationConfig(v *GetPublishedAgentResponseBodyDataApplicationConfig) *GetPublishedAgentResponseBodyData {
	s.ApplicationConfig = v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetCode(v string) *GetPublishedAgentResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetInstructions(v string) *GetPublishedAgentResponseBodyData {
	s.Instructions = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetModelId(v string) *GetPublishedAgentResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *GetPublishedAgentResponseBodyData) SetName(v string) *GetPublishedAgentResponseBodyData {
	s.Name = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfig struct {
	HistoryConfig  *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *GetPublishedAgentResponseBodyDataApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *GetPublishedAgentResponseBodyDataApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*GetPublishedAgentResponseBodyDataApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfig) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetHistoryConfig(v *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetLongTermMemory(v *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetParameters(v *GetPublishedAgentResponseBodyDataApplicationConfigParameters) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.Parameters = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetRagConfig(v *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetSecurity(v *GetPublishedAgentResponseBodyDataApplicationConfigSecurity) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.Security = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetTools(v []*GetPublishedAgentResponseBodyDataApplicationConfigTools) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.Tools = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfig) SetWorkFlows(v []*GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) *GetPublishedAgentResponseBodyDataApplicationConfig {
	s.WorkFlows = v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetEnableRecord(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetInstanceId(v string) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetRegion(v string) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig) SetStoreCode(v string) *GetPublishedAgentResponseBodyDataApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory) SetEnable(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigParameters struct {
	DialogRound *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	MaxTokens   *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) SetDialogRound(v int32) *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) SetMaxTokens(v int32) *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigParameters) SetTemperature(v float64) *GetPublishedAgentResponseBodyDataApplicationConfigParameters {
	s.Temperature = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigRagConfig struct {
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetEnableCitation(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetEnableSearch(v bool) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig) SetTopK(v int32) *GetPublishedAgentResponseBodyDataApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigSecurity) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigSecurity) SetProcessingStrategy(v string) *GetPublishedAgentResponseBodyDataApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigTools) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigTools) SetType(v string) *GetPublishedAgentResponseBodyDataApplicationConfigTools {
	s.Type = &v
	return s
}

type GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows) SetType(v string) *GetPublishedAgentResponseBodyDataApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

type GetPublishedAgentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublishedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublishedAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublishedAgentResponse) GoString() string {
	return s.String()
}

func (s *GetPublishedAgentResponse) SetHeaders(v map[string]*string) *GetPublishedAgentResponse {
	s.Headers = v
	return s
}

func (s *GetPublishedAgentResponse) SetStatusCode(v int32) *GetPublishedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublishedAgentResponse) SetBody(v *GetPublishedAgentResponseBody) *GetPublishedAgentResponse {
	s.Body = v
	return s
}

type ListCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdH70eOCSCKtacdomNzak4U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListCategoryRequest) SetCategoryType(v string) *ListCategoryRequest {
	s.CategoryType = &v
	return s
}

func (s *ListCategoryRequest) SetMaxResults(v int32) *ListCategoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCategoryRequest) SetNextToken(v string) *ListCategoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListCategoryRequest) SetParentCategoryId(v string) *ListCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

type ListCategoryResponseBody struct {
	// example:
	//
	// success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBody) SetCode(v string) *ListCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoryResponseBody) SetData(v *ListCategoryResponseBodyData) *ListCategoryResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoryResponseBody) SetMessage(v string) *ListCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoryResponseBody) SetRequestId(v string) *ListCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoryResponseBody) SetStatus(v string) *ListCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *ListCategoryResponseBody) SetSuccess(v bool) *ListCategoryResponseBody {
	s.Success = &v
	return s
}

type ListCategoryResponseBodyData struct {
	CategoryList []*ListCategoryResponseBodyDataCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh7Z4G8jiXR6IjHYd+M9WQVJ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCategoryResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBodyData) SetCategoryList(v []*ListCategoryResponseBodyDataCategoryList) *ListCategoryResponseBodyData {
	s.CategoryList = v
	return s
}

func (s *ListCategoryResponseBodyData) SetHasNext(v bool) *ListCategoryResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *ListCategoryResponseBodyData) SetMaxResults(v int32) *ListCategoryResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListCategoryResponseBodyData) SetNextToken(v string) *ListCategoryResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCategoryResponseBodyData) SetTotalCount(v int32) *ListCategoryResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCategoryResponseBodyDataCategoryList struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// cate_addd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListCategoryResponseBodyDataCategoryList) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponseBodyDataCategoryList) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBodyDataCategoryList) SetCategoryId(v string) *ListCategoryResponseBodyDataCategoryList {
	s.CategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetCategoryName(v string) *ListCategoryResponseBodyDataCategoryList {
	s.CategoryName = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetCategoryType(v string) *ListCategoryResponseBodyDataCategoryList {
	s.CategoryType = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetIsDefault(v bool) *ListCategoryResponseBodyDataCategoryList {
	s.IsDefault = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetParentCategoryId(v string) *ListCategoryResponseBodyDataCategoryList {
	s.ParentCategoryId = &v
	return s
}

type ListCategoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCategoryResponse) GoString() string {
	return s.String()
}

func (s *ListCategoryResponse) SetHeaders(v map[string]*string) *ListCategoryResponse {
	s.Headers = v
	return s
}

func (s *ListCategoryResponse) SetStatusCode(v int32) *ListCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCategoryResponse) SetBody(v *ListCategoryResponseBody) *ListCategoryResponse {
	s.Body = v
	return s
}

type ListChunksRequest struct {
	Fields []*string `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// file_5f03dfea56da4050ab68d61871fc4cb3_10151493
	Filed *string `json:"Filed,omitempty" xml:"Filed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// otoru9en4v
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChunksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChunksRequest) GoString() string {
	return s.String()
}

func (s *ListChunksRequest) SetFields(v []*string) *ListChunksRequest {
	s.Fields = v
	return s
}

func (s *ListChunksRequest) SetFiled(v string) *ListChunksRequest {
	s.Filed = &v
	return s
}

func (s *ListChunksRequest) SetIndexId(v string) *ListChunksRequest {
	s.IndexId = &v
	return s
}

func (s *ListChunksRequest) SetPageNum(v int32) *ListChunksRequest {
	s.PageNum = &v
	return s
}

func (s *ListChunksRequest) SetPageSize(v int32) *ListChunksRequest {
	s.PageSize = &v
	return s
}

type ListChunksResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListChunksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8F97A63B-55F1-527F-9D6E-467B6A7E8CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChunksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBody) SetCode(v string) *ListChunksResponseBody {
	s.Code = &v
	return s
}

func (s *ListChunksResponseBody) SetData(v *ListChunksResponseBodyData) *ListChunksResponseBody {
	s.Data = v
	return s
}

func (s *ListChunksResponseBody) SetMessage(v string) *ListChunksResponseBody {
	s.Message = &v
	return s
}

func (s *ListChunksResponseBody) SetRequestId(v string) *ListChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChunksResponseBody) SetStatus(v string) *ListChunksResponseBody {
	s.Status = &v
	return s
}

func (s *ListChunksResponseBody) SetSuccess(v bool) *ListChunksResponseBody {
	s.Success = &v
	return s
}

type ListChunksResponseBodyData struct {
	Nodes []*ListChunksResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 16
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChunksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListChunksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyData) SetNodes(v []*ListChunksResponseBodyDataNodes) *ListChunksResponseBodyData {
	s.Nodes = v
	return s
}

func (s *ListChunksResponseBodyData) SetTotal(v int64) *ListChunksResponseBodyData {
	s.Total = &v
	return s
}

type ListChunksResponseBodyDataNodes struct {
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 0.3
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text  *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListChunksResponseBodyDataNodes) String() string {
	return tea.Prettify(s)
}

func (s ListChunksResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyDataNodes) SetMetadata(v interface{}) *ListChunksResponseBodyDataNodes {
	s.Metadata = v
	return s
}

func (s *ListChunksResponseBodyDataNodes) SetScore(v float64) *ListChunksResponseBodyDataNodes {
	s.Score = &v
	return s
}

func (s *ListChunksResponseBodyDataNodes) SetText(v string) *ListChunksResponseBodyDataNodes {
	s.Text = &v
	return s
}

type ListChunksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChunksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChunksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChunksResponse) GoString() string {
	return s.String()
}

func (s *ListChunksResponse) SetHeaders(v map[string]*string) *ListChunksResponse {
	s.Headers = v
	return s
}

func (s *ListChunksResponse) SetStatusCode(v int32) *ListChunksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChunksResponse) SetBody(v *ListChunksResponseBody) *ListChunksResponse {
	s.Body = v
	return s
}

type ListFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdH70eOCSCKtacdomNzak4U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileRequest) GoString() string {
	return s.String()
}

func (s *ListFileRequest) SetCategoryId(v string) *ListFileRequest {
	s.CategoryId = &v
	return s
}

func (s *ListFileRequest) SetMaxResults(v int32) *ListFileRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFileRequest) SetNextToken(v string) *ListFileRequest {
	s.NextToken = &v
	return s
}

type ListFileResponseBody struct {
	// example:
	//
	// success
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Requests throttling triggered.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8F97A63B-55F1-527F-9D6E-467B6A7E8CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileResponseBody) SetCode(v string) *ListFileResponseBody {
	s.Code = &v
	return s
}

func (s *ListFileResponseBody) SetData(v *ListFileResponseBodyData) *ListFileResponseBody {
	s.Data = v
	return s
}

func (s *ListFileResponseBody) SetMessage(v string) *ListFileResponseBody {
	s.Message = &v
	return s
}

func (s *ListFileResponseBody) SetRequestId(v string) *ListFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileResponseBody) SetStatus(v string) *ListFileResponseBody {
	s.Status = &v
	return s
}

func (s *ListFileResponseBody) SetSuccess(v bool) *ListFileResponseBody {
	s.Success = &v
	return s
}

type ListFileResponseBodyData struct {
	FileList []*ListFileResponseBodyDataFileList `json:"FileList,omitempty" xml:"FileList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4jzbJk9J6lNeuXD9hP0viA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 48
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyData) SetFileList(v []*ListFileResponseBodyDataFileList) *ListFileResponseBodyData {
	s.FileList = v
	return s
}

func (s *ListFileResponseBodyData) SetHasNext(v bool) *ListFileResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *ListFileResponseBodyData) SetMaxResults(v int32) *ListFileResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListFileResponseBodyData) SetNextToken(v string) *ListFileResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListFileResponseBodyData) SetTotalCount(v int32) *ListFileResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListFileResponseBodyDataFileList struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 2023-08-18 11:03:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// file_5ff599b3455a45db8c41b0054b361518_10098576
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// auto-test-1721096109278.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// DASHSCOPE_DOCMIND
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// example:
	//
	// 512
	SizeInBytes *int64 `json:"SizeInBytes,omitempty" xml:"SizeInBytes,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFileResponseBodyDataFileList) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBodyDataFileList) GoString() string {
	return s.String()
}

func (s *ListFileResponseBodyDataFileList) SetCategoryId(v string) *ListFileResponseBodyDataFileList {
	s.CategoryId = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetCreateTime(v string) *ListFileResponseBodyDataFileList {
	s.CreateTime = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetFileId(v string) *ListFileResponseBodyDataFileList {
	s.FileId = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetFileName(v string) *ListFileResponseBodyDataFileList {
	s.FileName = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetFileType(v string) *ListFileResponseBodyDataFileList {
	s.FileType = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetParser(v string) *ListFileResponseBodyDataFileList {
	s.Parser = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetSizeInBytes(v int64) *ListFileResponseBodyDataFileList {
	s.SizeInBytes = &v
	return s
}

func (s *ListFileResponseBodyDataFileList) SetStatus(v string) *ListFileResponseBodyDataFileList {
	s.Status = &v
	return s
}

type ListFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponse) GoString() string {
	return s.String()
}

func (s *ListFileResponse) SetHeaders(v map[string]*string) *ListFileResponse {
	s.Headers = v
	return s
}

func (s *ListFileResponse) SetStatusCode(v int32) *ListFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileResponse) SetBody(v *ListFileResponseBody) *ListFileResponse {
	s.Body = v
	return s
}

type ListIndexDocumentsRequest struct {
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// example:
	//
	// FINISH
	DocumentStatus *string `json:"DocumentStatus,omitempty" xml:"DocumentStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIndexDocumentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndexDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsRequest) SetDocumentName(v string) *ListIndexDocumentsRequest {
	s.DocumentName = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetDocumentStatus(v string) *ListIndexDocumentsRequest {
	s.DocumentStatus = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetIndexId(v string) *ListIndexDocumentsRequest {
	s.IndexId = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetPageNumber(v int32) *ListIndexDocumentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetPageSize(v int32) *ListIndexDocumentsRequest {
	s.PageSize = &v
	return s
}

type ListIndexDocumentsResponseBody struct {
	// example:
	//
	// InvalidParameter
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIndexDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35A267BF-xxxx-54DB-8394-AA3B0742D833
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndexDocumentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndexDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponseBody) SetCode(v string) *ListIndexDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetData(v *ListIndexDocumentsResponseBodyData) *ListIndexDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetMessage(v string) *ListIndexDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetRequestId(v string) *ListIndexDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetStatus(v string) *ListIndexDocumentsResponseBody {
	s.Status = &v
	return s
}

func (s *ListIndexDocumentsResponseBody) SetSuccess(v bool) *ListIndexDocumentsResponseBody {
	s.Success = &v
	return s
}

type ListIndexDocumentsResponseBodyData struct {
	Documents []*ListIndexDocumentsResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// example:
	//
	// pno97tn8iu
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2437
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIndexDocumentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIndexDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponseBodyData) SetDocuments(v []*ListIndexDocumentsResponseBodyDataDocuments) *ListIndexDocumentsResponseBodyData {
	s.Documents = v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetIndexId(v string) *ListIndexDocumentsResponseBodyData {
	s.IndexId = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetPageNumber(v int32) *ListIndexDocumentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetPageSize(v int32) *ListIndexDocumentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyData) SetTotalCount(v int64) *ListIndexDocumentsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListIndexDocumentsResponseBodyDataDocuments struct {
	// example:
	//
	// 110002
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// pdf
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// example:
	//
	// doc_c134aa2073204a5d936d870bf960f56a10024701
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// check fileUrlKey[file_path] / fileNameKey[null] / fileExtensionKey[file_extension] is invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 996764
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// cate_21a407a3372c4ba7aedc649709143f0c10021401
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListIndexDocumentsResponseBodyDataDocuments) String() string {
	return tea.Prettify(s)
}

func (s ListIndexDocumentsResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetCode(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Code = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetDocumentType(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.DocumentType = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetId(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Id = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetMessage(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Message = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetName(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Name = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetSize(v int32) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Size = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetSourceId(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.SourceId = &v
	return s
}

func (s *ListIndexDocumentsResponseBodyDataDocuments) SetStatus(v string) *ListIndexDocumentsResponseBodyDataDocuments {
	s.Status = &v
	return s
}

type ListIndexDocumentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexDocumentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndexDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsResponse) SetHeaders(v map[string]*string) *ListIndexDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListIndexDocumentsResponse) SetStatusCode(v int32) *ListIndexDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexDocumentsResponse) SetBody(v *ListIndexDocumentsResponseBody) *ListIndexDocumentsResponse {
	s.Body = v
	return s
}

type ListIndicesRequest struct {
	// example:
	//
	// idx_status_score
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIndicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesRequest) GoString() string {
	return s.String()
}

func (s *ListIndicesRequest) SetIndexName(v string) *ListIndicesRequest {
	s.IndexName = &v
	return s
}

func (s *ListIndicesRequest) SetPageNumber(v string) *ListIndicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIndicesRequest) SetPageSize(v string) *ListIndicesRequest {
	s.PageSize = &v
	return s
}

type ListIndicesResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListIndicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBody) SetCode(v string) *ListIndicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIndicesResponseBody) SetData(v *ListIndicesResponseBodyData) *ListIndicesResponseBody {
	s.Data = v
	return s
}

func (s *ListIndicesResponseBody) SetMessage(v string) *ListIndicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIndicesResponseBody) SetRequestId(v string) *ListIndicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndicesResponseBody) SetStatus(v string) *ListIndicesResponseBody {
	s.Status = &v
	return s
}

func (s *ListIndicesResponseBody) SetSuccess(v bool) *ListIndicesResponseBody {
	s.Success = &v
	return s
}

type ListIndicesResponseBodyData struct {
	Indices []*ListIndicesResponseBodyDataIndices `json:"Indices,omitempty" xml:"Indices,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 48
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIndicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBodyData) SetIndices(v []*ListIndicesResponseBodyDataIndices) *ListIndicesResponseBodyData {
	s.Indices = v
	return s
}

func (s *ListIndicesResponseBodyData) SetPageNumber(v int32) *ListIndicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListIndicesResponseBodyData) SetPageSize(v int32) *ListIndicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIndicesResponseBodyData) SetTotalCount(v int32) *ListIndicesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListIndicesResponseBodyDataIndices struct {
	// example:
	//
	// 5
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// example:
	//
	// If each RAM user belongs to a RAM group, the configuration is considered compliant.
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// example:
	//
	// conv-rewrite-qwen-1.8b
	EmbeddingModelName *string `json:"EmbeddingModelName,omitempty" xml:"EmbeddingModelName,omitempty"`
	// example:
	//
	// 259899
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// temp_mUB4j
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// example:
	//
	// 0.01
	RerankMinScore *string `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// gte-rerank-hybrid
	RerankModelName *string `json:"RerankModelName,omitempty" xml:"RerankModelName,omitempty"`
	// example:
	//
	// \\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// example:
	//
	// gp-bp1gq62t1788yw2ol
	SinkInstanceId *string `json:"SinkInstanceId,omitempty" xml:"SinkInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	SinkRegion *string `json:"SinkRegion,omitempty" xml:"SinkRegion,omitempty"`
	// example:
	//
	// es
	SinkType *string `json:"SinkType,omitempty" xml:"SinkType,omitempty"`
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// structured
	StructureType *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s ListIndicesResponseBodyDataIndices) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesResponseBodyDataIndices) GoString() string {
	return s.String()
}

func (s *ListIndicesResponseBodyDataIndices) SetChunkSize(v int32) *ListIndicesResponseBodyDataIndices {
	s.ChunkSize = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetDescription(v string) *ListIndicesResponseBodyDataIndices {
	s.Description = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetDocumentIds(v []*string) *ListIndicesResponseBodyDataIndices {
	s.DocumentIds = v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetEmbeddingModelName(v string) *ListIndicesResponseBodyDataIndices {
	s.EmbeddingModelName = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetId(v string) *ListIndicesResponseBodyDataIndices {
	s.Id = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetName(v string) *ListIndicesResponseBodyDataIndices {
	s.Name = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetOverlapSize(v int32) *ListIndicesResponseBodyDataIndices {
	s.OverlapSize = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetRerankMinScore(v string) *ListIndicesResponseBodyDataIndices {
	s.RerankMinScore = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetRerankModelName(v string) *ListIndicesResponseBodyDataIndices {
	s.RerankModelName = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSeparator(v string) *ListIndicesResponseBodyDataIndices {
	s.Separator = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSinkInstanceId(v string) *ListIndicesResponseBodyDataIndices {
	s.SinkInstanceId = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSinkRegion(v string) *ListIndicesResponseBodyDataIndices {
	s.SinkRegion = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSinkType(v string) *ListIndicesResponseBodyDataIndices {
	s.SinkType = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetSourceType(v string) *ListIndicesResponseBodyDataIndices {
	s.SourceType = &v
	return s
}

func (s *ListIndicesResponseBodyDataIndices) SetStructureType(v string) *ListIndicesResponseBodyDataIndices {
	s.StructureType = &v
	return s
}

type ListIndicesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIndicesResponse) GoString() string {
	return s.String()
}

func (s *ListIndicesResponse) SetHeaders(v map[string]*string) *ListIndicesResponse {
	s.Headers = v
	return s
}

func (s *ListIndicesResponse) SetStatusCode(v int32) *ListIndicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndicesResponse) SetBody(v *ListIndicesResponseBody) *ListIndicesResponse {
	s.Body = v
	return s
}

type ListMemoriesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMemoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMemoriesRequest) GoString() string {
	return s.String()
}

func (s *ListMemoriesRequest) SetMaxResults(v int32) *ListMemoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoriesRequest) SetNextToken(v string) *ListMemoriesRequest {
	s.NextToken = &v
	return s
}

type ListMemoriesResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32                              `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	Memories   []*ListMemoriesResponseBodyMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 105
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// llm-us9hjmt32nysdm5v
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListMemoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBody) SetMaxResults(v int32) *ListMemoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoriesResponseBody) SetMemories(v []*ListMemoriesResponseBodyMemories) *ListMemoriesResponseBody {
	s.Memories = v
	return s
}

func (s *ListMemoriesResponseBody) SetNextToken(v string) *ListMemoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoriesResponseBody) SetRequestId(v string) *ListMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoriesResponseBody) SetTotalCount(v int32) *ListMemoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMemoriesResponseBody) SetWorkspaceId(v string) *ListMemoriesResponseBody {
	s.WorkspaceId = &v
	return s
}

type ListMemoriesResponseBodyMemories struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 3fc531f4519444beaafffa4538f60667
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
}

func (s ListMemoriesResponseBodyMemories) String() string {
	return tea.Prettify(s)
}

func (s ListMemoriesResponseBodyMemories) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBodyMemories) SetDescription(v string) *ListMemoriesResponseBodyMemories {
	s.Description = &v
	return s
}

func (s *ListMemoriesResponseBodyMemories) SetMemoryId(v string) *ListMemoriesResponseBodyMemories {
	s.MemoryId = &v
	return s
}

type ListMemoriesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMemoriesResponse) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponse) SetHeaders(v map[string]*string) *ListMemoriesResponse {
	s.Headers = v
	return s
}

func (s *ListMemoriesResponse) SetStatusCode(v int32) *ListMemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoriesResponse) SetBody(v *ListMemoriesResponseBody) *ListMemoriesResponse {
	s.Body = v
	return s
}

type ListMemoryNodesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMemoryNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMemoryNodesRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesRequest) SetMaxResults(v int32) *ListMemoryNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryNodesRequest) SetNextToken(v string) *ListMemoryNodesRequest {
	s.NextToken = &v
	return s
}

type ListMemoryNodesResponseBody struct {
	// example:
	//
	// 20
	MaxResults  *int32                                    `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	MemoryNodes []*ListMemoryNodesResponseBodyMemoryNodes `json:"memoryNodes,omitempty" xml:"memoryNodes,omitempty" type:"Repeated"`
	// example:
	//
	// dc270401186b433f975d7e1faaa34e0e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMemoryNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMemoryNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesResponseBody) SetMaxResults(v int32) *ListMemoryNodesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryNodesResponseBody) SetMemoryNodes(v []*ListMemoryNodesResponseBodyMemoryNodes) *ListMemoryNodesResponseBody {
	s.MemoryNodes = v
	return s
}

func (s *ListMemoryNodesResponseBody) SetNextToken(v string) *ListMemoryNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoryNodesResponseBody) SetRequestId(v string) *ListMemoryNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryNodesResponseBody) SetTotalCount(v int32) *ListMemoryNodesResponseBody {
	s.TotalCount = &v
	return s
}

type ListMemoryNodesResponseBodyMemoryNodes struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 68de06c95368463a8be4a84efc872cc5
	MemoryNodeId *string `json:"memoryNodeId,omitempty" xml:"memoryNodeId,omitempty"`
}

func (s ListMemoryNodesResponseBodyMemoryNodes) String() string {
	return tea.Prettify(s)
}

func (s ListMemoryNodesResponseBodyMemoryNodes) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) SetContent(v string) *ListMemoryNodesResponseBodyMemoryNodes {
	s.Content = &v
	return s
}

func (s *ListMemoryNodesResponseBodyMemoryNodes) SetMemoryNodeId(v string) *ListMemoryNodesResponseBodyMemoryNodes {
	s.MemoryNodeId = &v
	return s
}

type ListMemoryNodesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMemoryNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMemoryNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMemoryNodesResponse) GoString() string {
	return s.String()
}

func (s *ListMemoryNodesResponse) SetHeaders(v map[string]*string) *ListMemoryNodesResponse {
	s.Headers = v
	return s
}

func (s *ListMemoryNodesResponse) SetStatusCode(v int32) *ListMemoryNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMemoryNodesResponse) SetBody(v *ListMemoryNodesResponseBody) *ListMemoryNodesResponse {
	s.Body = v
	return s
}

type ListPublishedAgentRequest struct {
	PageNo   *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPublishedAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentRequest) SetPageNo(v int32) *ListPublishedAgentRequest {
	s.PageNo = &v
	return s
}

func (s *ListPublishedAgentRequest) SetPageSize(v int32) *ListPublishedAgentRequest {
	s.PageSize = &v
	return s
}

type ListPublishedAgentResponseBody struct {
	Code           *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListPublishedAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                             `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *string                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPublishedAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBody) SetCode(v string) *ListPublishedAgentResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetData(v *ListPublishedAgentResponseBodyData) *ListPublishedAgentResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishedAgentResponseBody) SetHttpStatusCode(v int32) *ListPublishedAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetMessage(v string) *ListPublishedAgentResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetRequestId(v string) *ListPublishedAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedAgentResponseBody) SetSuccess(v string) *ListPublishedAgentResponseBody {
	s.Success = &v
	return s
}

type ListPublishedAgentResponseBodyData struct {
	List     []*ListPublishedAgentResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNo   *int32                                    `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	PageSize *int32                                    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                    `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListPublishedAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyData) SetList(v []*ListPublishedAgentResponseBodyDataList) *ListPublishedAgentResponseBodyData {
	s.List = v
	return s
}

func (s *ListPublishedAgentResponseBodyData) SetPageNo(v int32) *ListPublishedAgentResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListPublishedAgentResponseBodyData) SetPageSize(v int32) *ListPublishedAgentResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPublishedAgentResponseBodyData) SetTotal(v int32) *ListPublishedAgentResponseBodyData {
	s.Total = &v
	return s
}

type ListPublishedAgentResponseBodyDataList struct {
	ApplicationConfig *ListPublishedAgentResponseBodyDataListApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Code              *string                                                  `json:"code,omitempty" xml:"code,omitempty"`
	Instructions      *string                                                  `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                                  `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                                  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataList) SetApplicationConfig(v *ListPublishedAgentResponseBodyDataListApplicationConfig) *ListPublishedAgentResponseBodyDataList {
	s.ApplicationConfig = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetCode(v string) *ListPublishedAgentResponseBodyDataList {
	s.Code = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetInstructions(v string) *ListPublishedAgentResponseBodyDataList {
	s.Instructions = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetModelId(v string) *ListPublishedAgentResponseBodyDataList {
	s.ModelId = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataList) SetName(v string) *ListPublishedAgentResponseBodyDataList {
	s.Name = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfig struct {
	HistoryConfig  *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *ListPublishedAgentResponseBodyDataListApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*ListPublishedAgentResponseBodyDataListApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfig) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfig) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetHistoryConfig(v *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetLongTermMemory(v *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetParameters(v *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.Parameters = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetRagConfig(v *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetSecurity(v *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.Security = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetTools(v []*ListPublishedAgentResponseBodyDataListApplicationConfigTools) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.Tools = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfig) SetWorkFlows(v []*ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) *ListPublishedAgentResponseBodyDataListApplicationConfig {
	s.WorkFlows = v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetEnableRecord(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetInstanceId(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetRegion(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig) SetStoreCode(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory) SetEnable(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigParameters struct {
	DialogRound *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	MaxTokens   *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) SetDialogRound(v int32) *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) SetMaxTokens(v int32) *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigParameters) SetTemperature(v float64) *ListPublishedAgentResponseBodyDataListApplicationConfigParameters {
	s.Temperature = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig struct {
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetEnableCitation(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetEnableSearch(v bool) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig) SetTopK(v int32) *ListPublishedAgentResponseBodyDataListApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity) SetProcessingStrategy(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigTools) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigTools) SetType(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigTools {
	s.Type = &v
	return s
}

type ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows) SetType(v string) *ListPublishedAgentResponseBodyDataListApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

type ListPublishedAgentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublishedAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublishedAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAgentResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAgentResponse) SetHeaders(v map[string]*string) *ListPublishedAgentResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAgentResponse) SetStatusCode(v int32) *ListPublishedAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAgentResponse) SetBody(v *ListPublishedAgentResponseBody) *ListPublishedAgentResponse {
	s.Body = v
	return s
}

type RetrieveRequest struct {
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// example:
	//
	// false
	EnableRewrite *bool     `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	Images        []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0m2g6t
	IndexId *string                  `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Query   *string                  `json:"Query,omitempty" xml:"Query,omitempty"`
	Rerank  []*RetrieveRequestRerank `json:"Rerank,omitempty" xml:"Rerank,omitempty" type:"Repeated"`
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// 5
	RerankTopN *int32                    `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	Rewrite    []*RetrieveRequestRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SaveRetrieverHistory *bool                `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	SearchFilters        []map[string]*string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	SparseSimilarityTopK *int32 `json:"SparseSimilarityTopK,omitempty" xml:"SparseSimilarityTopK,omitempty"`
}

func (s RetrieveRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequest) GoString() string {
	return s.String()
}

func (s *RetrieveRequest) SetDenseSimilarityTopK(v int32) *RetrieveRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *RetrieveRequest) SetEnableReranking(v bool) *RetrieveRequest {
	s.EnableReranking = &v
	return s
}

func (s *RetrieveRequest) SetEnableRewrite(v bool) *RetrieveRequest {
	s.EnableRewrite = &v
	return s
}

func (s *RetrieveRequest) SetImages(v []*string) *RetrieveRequest {
	s.Images = v
	return s
}

func (s *RetrieveRequest) SetIndexId(v string) *RetrieveRequest {
	s.IndexId = &v
	return s
}

func (s *RetrieveRequest) SetQuery(v string) *RetrieveRequest {
	s.Query = &v
	return s
}

func (s *RetrieveRequest) SetRerank(v []*RetrieveRequestRerank) *RetrieveRequest {
	s.Rerank = v
	return s
}

func (s *RetrieveRequest) SetRerankMinScore(v float32) *RetrieveRequest {
	s.RerankMinScore = &v
	return s
}

func (s *RetrieveRequest) SetRerankTopN(v int32) *RetrieveRequest {
	s.RerankTopN = &v
	return s
}

func (s *RetrieveRequest) SetRewrite(v []*RetrieveRequestRewrite) *RetrieveRequest {
	s.Rewrite = v
	return s
}

func (s *RetrieveRequest) SetSaveRetrieverHistory(v bool) *RetrieveRequest {
	s.SaveRetrieverHistory = &v
	return s
}

func (s *RetrieveRequest) SetSearchFilters(v []map[string]*string) *RetrieveRequest {
	s.SearchFilters = v
	return s
}

func (s *RetrieveRequest) SetSparseSimilarityTopK(v int32) *RetrieveRequest {
	s.SparseSimilarityTopK = &v
	return s
}

type RetrieveRequestRerank struct {
	// example:
	//
	// gte-rerank-hybrid
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s RetrieveRequestRerank) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequestRerank) GoString() string {
	return s.String()
}

func (s *RetrieveRequestRerank) SetModelName(v string) *RetrieveRequestRerank {
	s.ModelName = &v
	return s
}

type RetrieveRequestRewrite struct {
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s RetrieveRequestRewrite) String() string {
	return tea.Prettify(s)
}

func (s RetrieveRequestRewrite) GoString() string {
	return s.String()
}

func (s *RetrieveRequestRewrite) SetModelName(v string) *RetrieveRequestRewrite {
	s.ModelName = &v
	return s
}

type RetrieveShrinkRequest struct {
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// example:
	//
	// false
	EnableRewrite *bool   `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	ImagesShrink  *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0m2g6t
	IndexId      *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	Query        *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RerankShrink *string `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// example:
	//
	// 5
	RerankTopN    *int32  `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	RewriteShrink *string `json:"Rewrite,omitempty" xml:"Rewrite,omitempty"`
	// example:
	//
	// false
	SaveRetrieverHistory *bool   `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	SearchFiltersShrink  *string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty"`
	// example:
	//
	// 100
	SparseSimilarityTopK *int32 `json:"SparseSimilarityTopK,omitempty" xml:"SparseSimilarityTopK,omitempty"`
}

func (s RetrieveShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RetrieveShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetrieveShrinkRequest) SetDenseSimilarityTopK(v int32) *RetrieveShrinkRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *RetrieveShrinkRequest) SetEnableReranking(v bool) *RetrieveShrinkRequest {
	s.EnableReranking = &v
	return s
}

func (s *RetrieveShrinkRequest) SetEnableRewrite(v bool) *RetrieveShrinkRequest {
	s.EnableRewrite = &v
	return s
}

func (s *RetrieveShrinkRequest) SetImagesShrink(v string) *RetrieveShrinkRequest {
	s.ImagesShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetIndexId(v string) *RetrieveShrinkRequest {
	s.IndexId = &v
	return s
}

func (s *RetrieveShrinkRequest) SetQuery(v string) *RetrieveShrinkRequest {
	s.Query = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankShrink(v string) *RetrieveShrinkRequest {
	s.RerankShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankMinScore(v float32) *RetrieveShrinkRequest {
	s.RerankMinScore = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankTopN(v int32) *RetrieveShrinkRequest {
	s.RerankTopN = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRewriteShrink(v string) *RetrieveShrinkRequest {
	s.RewriteShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSaveRetrieverHistory(v bool) *RetrieveShrinkRequest {
	s.SaveRetrieverHistory = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSearchFiltersShrink(v string) *RetrieveShrinkRequest {
	s.SearchFiltersShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSparseSimilarityTopK(v int32) *RetrieveShrinkRequest {
	s.SparseSimilarityTopK = &v
	return s
}

type RetrieveResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RetrieveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetrieveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBody) SetCode(v string) *RetrieveResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveResponseBody) SetData(v *RetrieveResponseBodyData) *RetrieveResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveResponseBody) SetMessage(v string) *RetrieveResponseBody {
	s.Message = &v
	return s
}

func (s *RetrieveResponseBody) SetRequestId(v string) *RetrieveResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveResponseBody) SetStatus(v string) *RetrieveResponseBody {
	s.Status = &v
	return s
}

func (s *RetrieveResponseBody) SetSuccess(v bool) *RetrieveResponseBody {
	s.Success = &v
	return s
}

type RetrieveResponseBodyData struct {
	Nodes []*RetrieveResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RetrieveResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyData) SetNodes(v []*RetrieveResponseBodyDataNodes) *RetrieveResponseBodyData {
	s.Nodes = v
	return s
}

type RetrieveResponseBodyDataNodes struct {
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 0.3
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	Text  *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RetrieveResponseBodyDataNodes) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *RetrieveResponseBodyDataNodes) SetMetadata(v interface{}) *RetrieveResponseBodyDataNodes {
	s.Metadata = v
	return s
}

func (s *RetrieveResponseBodyDataNodes) SetScore(v float64) *RetrieveResponseBodyDataNodes {
	s.Score = &v
	return s
}

func (s *RetrieveResponseBodyDataNodes) SetText(v string) *RetrieveResponseBodyDataNodes {
	s.Text = &v
	return s
}

type RetrieveResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveResponse) String() string {
	return tea.Prettify(s)
}

func (s RetrieveResponse) GoString() string {
	return s.String()
}

func (s *RetrieveResponse) SetHeaders(v map[string]*string) *RetrieveResponse {
	s.Headers = v
	return s
}

func (s *RetrieveResponse) SetStatusCode(v int32) *RetrieveResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveResponse) SetBody(v *RetrieveResponseBody) *RetrieveResponse {
	s.Body = v
	return s
}

type SubmitIndexAddDocumentsJobRequest struct {
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SubmitIndexAddDocumentsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobRequest) SetCategoryIds(v []*string) *SubmitIndexAddDocumentsJobRequest {
	s.CategoryIds = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetDocumentIds(v []*string) *SubmitIndexAddDocumentsJobRequest {
	s.DocumentIds = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetIndexId(v string) *SubmitIndexAddDocumentsJobRequest {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetSourceType(v string) *SubmitIndexAddDocumentsJobRequest {
	s.SourceType = &v
	return s
}

type SubmitIndexAddDocumentsJobShrinkRequest struct {
	CategoryIdsShrink *string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty"`
	DocumentIdsShrink *string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SubmitIndexAddDocumentsJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetCategoryIdsShrink(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.CategoryIdsShrink = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetDocumentIdsShrink(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.DocumentIdsShrink = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetIndexId(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobShrinkRequest) SetSourceType(v string) *SubmitIndexAddDocumentsJobShrinkRequest {
	s.SourceType = &v
	return s
}

type SubmitIndexAddDocumentsJobResponseBody struct {
	// example:
	//
	// Index.InvalidParameter
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitIndexAddDocumentsJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 778C0B3B-03C1-5FC1-A947-36EDD13606AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIndexAddDocumentsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetCode(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetData(v *SubmitIndexAddDocumentsJobResponseBodyData) *SubmitIndexAddDocumentsJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetMessage(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetRequestId(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetStatus(v string) *SubmitIndexAddDocumentsJobResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponseBody) SetSuccess(v bool) *SubmitIndexAddDocumentsJobResponseBody {
	s.Success = &v
	return s
}

type SubmitIndexAddDocumentsJobResponseBodyData struct {
	// example:
	//
	// 42687eb254a34802bed398357f5498ae
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitIndexAddDocumentsJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobResponseBodyData) SetId(v string) *SubmitIndexAddDocumentsJobResponseBodyData {
	s.Id = &v
	return s
}

type SubmitIndexAddDocumentsJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIndexAddDocumentsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIndexAddDocumentsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobResponse) SetHeaders(v map[string]*string) *SubmitIndexAddDocumentsJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponse) SetStatusCode(v int32) *SubmitIndexAddDocumentsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobResponse) SetBody(v *SubmitIndexAddDocumentsJobResponseBody) *SubmitIndexAddDocumentsJobResponse {
	s.Body = v
	return s
}

type SubmitIndexJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s SubmitIndexJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobRequest) SetIndexId(v string) *SubmitIndexJobRequest {
	s.IndexId = &v
	return s
}

type SubmitIndexJobResponseBody struct {
	// example:
	//
	// InvalidParameter
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitIndexJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitIndexJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponseBody) SetCode(v string) *SubmitIndexJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetData(v *SubmitIndexJobResponseBodyData) *SubmitIndexJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitIndexJobResponseBody) SetMessage(v string) *SubmitIndexJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetRequestId(v string) *SubmitIndexJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetStatus(v string) *SubmitIndexJobResponseBody {
	s.Status = &v
	return s
}

func (s *SubmitIndexJobResponseBody) SetSuccess(v bool) *SubmitIndexJobResponseBody {
	s.Success = &v
	return s
}

type SubmitIndexJobResponseBodyData struct {
	// example:
	//
	// eFDr2fGRzP9gdDZWAdo3YQ==
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// khdyak1uuj
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
}

func (s SubmitIndexJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponseBodyData) SetId(v string) *SubmitIndexJobResponseBodyData {
	s.Id = &v
	return s
}

func (s *SubmitIndexJobResponseBodyData) SetIndexId(v string) *SubmitIndexJobResponseBodyData {
	s.IndexId = &v
	return s
}

type SubmitIndexJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIndexJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIndexJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitIndexJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitIndexJobResponse) SetHeaders(v map[string]*string) *SubmitIndexJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitIndexJobResponse) SetStatusCode(v int32) *SubmitIndexJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIndexJobResponse) SetBody(v *SubmitIndexJobResponseBody) *SubmitIndexJobResponse {
	s.Body = v
	return s
}

type UpdateAndPublishAgentRequest struct {
	ApplicationConfig *UpdateAndPublishAgentRequestApplicationConfig `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty" type:"Struct"`
	Instructions      *string                                        `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId           *string                                        `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name              *string                                        `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateAndPublishAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequest) SetApplicationConfig(v *UpdateAndPublishAgentRequestApplicationConfig) *UpdateAndPublishAgentRequest {
	s.ApplicationConfig = v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetInstructions(v string) *UpdateAndPublishAgentRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetModelId(v string) *UpdateAndPublishAgentRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateAndPublishAgentRequest) SetName(v string) *UpdateAndPublishAgentRequest {
	s.Name = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfig struct {
	HistoryConfig  *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig  `json:"historyConfig,omitempty" xml:"historyConfig,omitempty" type:"Struct"`
	LongTermMemory *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory `json:"longTermMemory,omitempty" xml:"longTermMemory,omitempty" type:"Struct"`
	Parameters     *UpdateAndPublishAgentRequestApplicationConfigParameters     `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Struct"`
	RagConfig      *UpdateAndPublishAgentRequestApplicationConfigRagConfig      `json:"ragConfig,omitempty" xml:"ragConfig,omitempty" type:"Struct"`
	Security       *UpdateAndPublishAgentRequestApplicationConfigSecurity       `json:"security,omitempty" xml:"security,omitempty" type:"Struct"`
	Tools          []*UpdateAndPublishAgentRequestApplicationConfigTools        `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
	WorkFlows      []*UpdateAndPublishAgentRequestApplicationConfigWorkFlows    `json:"workFlows,omitempty" xml:"workFlows,omitempty" type:"Repeated"`
}

func (s UpdateAndPublishAgentRequestApplicationConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetHistoryConfig(v *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) *UpdateAndPublishAgentRequestApplicationConfig {
	s.HistoryConfig = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetLongTermMemory(v *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) *UpdateAndPublishAgentRequestApplicationConfig {
	s.LongTermMemory = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetParameters(v *UpdateAndPublishAgentRequestApplicationConfigParameters) *UpdateAndPublishAgentRequestApplicationConfig {
	s.Parameters = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetRagConfig(v *UpdateAndPublishAgentRequestApplicationConfigRagConfig) *UpdateAndPublishAgentRequestApplicationConfig {
	s.RagConfig = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetSecurity(v *UpdateAndPublishAgentRequestApplicationConfigSecurity) *UpdateAndPublishAgentRequestApplicationConfig {
	s.Security = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetTools(v []*UpdateAndPublishAgentRequestApplicationConfigTools) *UpdateAndPublishAgentRequestApplicationConfig {
	s.Tools = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfig) SetWorkFlows(v []*UpdateAndPublishAgentRequestApplicationConfigWorkFlows) *UpdateAndPublishAgentRequestApplicationConfig {
	s.WorkFlows = v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigHistoryConfig struct {
	EnableAdbRecord *bool   `json:"enableAdbRecord,omitempty" xml:"enableAdbRecord,omitempty"`
	EnableRecord    *bool   `json:"enableRecord,omitempty" xml:"enableRecord,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region          *string `json:"region,omitempty" xml:"region,omitempty"`
	StoreCode       *string `json:"storeCode,omitempty" xml:"storeCode,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetEnableAdbRecord(v bool) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.EnableAdbRecord = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetEnableRecord(v bool) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.EnableRecord = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetInstanceId(v string) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.InstanceId = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetRegion(v string) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.Region = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig) SetStoreCode(v string) *UpdateAndPublishAgentRequestApplicationConfigHistoryConfig {
	s.StoreCode = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigLongTermMemory struct {
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory) SetEnable(v bool) *UpdateAndPublishAgentRequestApplicationConfigLongTermMemory {
	s.Enable = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigParameters struct {
	DialogRound *int32   `json:"dialogRound,omitempty" xml:"dialogRound,omitempty"`
	MaxTokens   *int32   `json:"maxTokens,omitempty" xml:"maxTokens,omitempty"`
	Temperature *float64 `json:"temperature,omitempty" xml:"temperature,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigParameters) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetDialogRound(v int32) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.DialogRound = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetMaxTokens(v int32) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.MaxTokens = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigParameters) SetTemperature(v float64) *UpdateAndPublishAgentRequestApplicationConfigParameters {
	s.Temperature = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigRagConfig struct {
	EnableCitation        *bool     `json:"enableCitation,omitempty" xml:"enableCitation,omitempty"`
	EnableSearch          *bool     `json:"enableSearch,omitempty" xml:"enableSearch,omitempty"`
	KnowledgeBaseCodeList []*string `json:"knowledgeBaseCodeList,omitempty" xml:"knowledgeBaseCodeList,omitempty" type:"Repeated"`
	TopK                  *int32    `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigRagConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigRagConfig) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetEnableCitation(v bool) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.EnableCitation = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetEnableSearch(v bool) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *UpdateAndPublishAgentRequestApplicationConfigRagConfig) SetTopK(v int32) *UpdateAndPublishAgentRequestApplicationConfigRagConfig {
	s.TopK = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigSecurity struct {
	ProcessingStrategy *string `json:"processingStrategy,omitempty" xml:"processingStrategy,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigSecurity) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigSecurity) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigSecurity) SetProcessingStrategy(v string) *UpdateAndPublishAgentRequestApplicationConfigSecurity {
	s.ProcessingStrategy = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigTools struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigTools) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigTools) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigTools) SetType(v string) *UpdateAndPublishAgentRequestApplicationConfigTools {
	s.Type = &v
	return s
}

type UpdateAndPublishAgentRequestApplicationConfigWorkFlows struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateAndPublishAgentRequestApplicationConfigWorkFlows) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentRequestApplicationConfigWorkFlows) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentRequestApplicationConfigWorkFlows) SetType(v string) *UpdateAndPublishAgentRequestApplicationConfigWorkFlows {
	s.Type = &v
	return s
}

type UpdateAndPublishAgentShrinkRequest struct {
	ApplicationConfigShrink *string `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty"`
	Instructions            *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId                 *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s UpdateAndPublishAgentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentShrinkRequest) SetApplicationConfigShrink(v string) *UpdateAndPublishAgentShrinkRequest {
	s.ApplicationConfigShrink = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetInstructions(v string) *UpdateAndPublishAgentShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetModelId(v string) *UpdateAndPublishAgentShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetName(v string) *UpdateAndPublishAgentShrinkRequest {
	s.Name = &v
	return s
}

type UpdateAndPublishAgentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAndPublishAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentResponseBody) SetCode(v string) *UpdateAndPublishAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetData(v string) *UpdateAndPublishAgentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetHttpStatusCode(v int32) *UpdateAndPublishAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetMessage(v string) *UpdateAndPublishAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetRequestId(v string) *UpdateAndPublishAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetSuccess(v bool) *UpdateAndPublishAgentResponseBody {
	s.Success = &v
	return s
}

type UpdateAndPublishAgentResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAndPublishAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAndPublishAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAndPublishAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentResponse) SetHeaders(v map[string]*string) *UpdateAndPublishAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAndPublishAgentResponse) SetStatusCode(v int32) *UpdateAndPublishAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAndPublishAgentResponse) SetBody(v *UpdateAndPublishAgentResponseBody) *UpdateAndPublishAgentResponse {
	s.Body = v
	return s
}

type UpdateMemoryRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateMemoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) SetDescription(v string) *UpdateMemoryRequest {
	s.Description = &v
	return s
}

type UpdateMemoryResponseBody struct {
	// example:
	//
	// 6a71f2d9-f1c9-913b-818b-114029103cad
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponse) SetHeaders(v map[string]*string) *UpdateMemoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryResponse) SetStatusCode(v int32) *UpdateMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryResponse) SetBody(v *UpdateMemoryResponseBody) *UpdateMemoryResponse {
	s.Body = v
	return s
}

type UpdateMemoryNodeRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s UpdateMemoryNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemoryNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryNodeRequest) SetContent(v string) *UpdateMemoryNodeRequest {
	s.Content = &v
	return s
}

type UpdateMemoryNodeResponseBody struct {
	// example:
	//
	// 8C56C7AF-6573-19CE-B018-E05E1EDCF4C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMemoryNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemoryNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryNodeResponseBody) SetRequestId(v string) *UpdateMemoryNodeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMemoryNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryNodeResponse) SetHeaders(v map[string]*string) *UpdateMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryNodeResponse) SetStatusCode(v int32) *UpdateMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryNodeResponse) SetBody(v *UpdateMemoryNodeResponseBody) *UpdateMemoryNodeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bailian"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - AddCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCategoryResponse
func (client *Client) AddCategoryWithOptions(WorkspaceId *string, request *AddCategoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryName)) {
		body["CategoryName"] = request.CategoryName
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryType)) {
		body["CategoryType"] = request.CategoryType
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryId)) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCategory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/category/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddCategoryRequest
//
// @return AddCategoryResponse
func (client *Client) AddCategory(WorkspaceId *string, request *AddCategoryRequest) (_result *AddCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddCategoryResponse{}
	_body, _err := client.AddCategoryWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 将临时上传的文档导入百炼数据中心，导入成功之后会自动触发文档解析。
//
// @param request - AddFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFileResponse
func (client *Client) AddFileWithOptions(WorkspaceId *string, request *AddFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LeaseId)) {
		body["LeaseId"] = request.LeaseId
	}

	if !tea.BoolValue(util.IsUnset(request.Parser)) {
		body["Parser"] = request.Parser
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 将临时上传的文档导入百炼数据中心，导入成功之后会自动触发文档解析。
//
// @param request - AddFileRequest
//
// @return AddFileResponse
func (client *Client) AddFile(WorkspaceId *string, request *AddFileRequest) (_result *AddFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFileResponse{}
	_body, _err := client.AddFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 请求文档上传租约，进行文档上传。
//
// @param request - ApplyFileUploadLeaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLeaseWithOptions(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyFileUploadLeaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.Md5)) {
		body["Md5"] = request.Md5
	}

	if !tea.BoolValue(util.IsUnset(request.SizeInBytes)) {
		body["SizeInBytes"] = request.SizeInBytes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyFileUploadLease"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/category/" + tea.StringValue(openapiutil.GetEncodeParam(CategoryId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 请求文档上传租约，进行文档上传。
//
// @param request - ApplyFileUploadLeaseRequest
//
// @return ApplyFileUploadLeaseResponse
func (client *Client) ApplyFileUploadLease(CategoryId *string, WorkspaceId *string, request *ApplyFileUploadLeaseRequest) (_result *ApplyFileUploadLeaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyFileUploadLeaseResponse{}
	_body, _err := client.ApplyFileUploadLeaseWithOptions(CategoryId, WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建并发布智能体应用
//
// @param tmpReq - CreateAndPulishAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndPulishAgentResponse
func (client *Client) CreateAndPulishAgentWithOptions(workspaceId *string, tmpReq *CreateAndPulishAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAndPulishAgentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAndPulishAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApplicationConfig)) {
		request.ApplicationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationConfig, tea.String("applicationConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationConfigShrink)) {
		body["applicationConfig"] = request.ApplicationConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAndPulishAgent"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/application/agents"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAndPulishAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并发布智能体应用
//
// @param request - CreateAndPulishAgentRequest
//
// @return CreateAndPulishAgentResponse
func (client *Client) CreateAndPulishAgent(workspaceId *string, request *CreateAndPulishAgentRequest) (_result *CreateAndPulishAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAndPulishAgentResponse{}
	_body, _err := client.CreateAndPulishAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建并运行pipeline
//
// @param tmpReq - CreateIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndexResponse
func (client *Client) CreateIndexWithOptions(WorkspaceId *string, tmpReq *CreateIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIndexResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateIndexShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CategoryIds)) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, tea.String("CategoryIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Columns)) {
		request.ColumnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Columns, tea.String("Columns"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DataSource)) {
		request.DataSourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DataSource, tea.String("DataSource"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DocumentIds)) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, tea.String("DocumentIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryIdsShrink)) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ChunkSize)) {
		query["ChunkSize"] = request.ChunkSize
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnsShrink)) {
		query["Columns"] = request.ColumnsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceShrink)) {
		query["DataSource"] = request.DataSourceShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentIdsShrink)) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EmbeddingModelName)) {
		query["EmbeddingModelName"] = request.EmbeddingModelName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OverlapSize)) {
		query["OverlapSize"] = request.OverlapSize
	}

	if !tea.BoolValue(util.IsUnset(request.RerankMinScore)) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !tea.BoolValue(util.IsUnset(request.RerankModelName)) {
		query["RerankModelName"] = request.RerankModelName
	}

	if !tea.BoolValue(util.IsUnset(request.Separator)) {
		query["Separator"] = request.Separator
	}

	if !tea.BoolValue(util.IsUnset(request.SinkInstanceId)) {
		query["SinkInstanceId"] = request.SinkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SinkRegion)) {
		query["SinkRegion"] = request.SinkRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SinkType)) {
		query["SinkType"] = request.SinkType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StructureType)) {
		query["StructureType"] = request.StructureType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIndex"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建并运行pipeline
//
// @param request - CreateIndexRequest
//
// @return CreateIndexResponse
func (client *Client) CreateIndex(WorkspaceId *string, request *CreateIndexRequest) (_result *CreateIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndexResponse{}
	_body, _err := client.CreateIndexWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Memory
//
// @param request - CreateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryResponse
func (client *Client) CreateMemoryWithOptions(workspaceId *string, request *CreateMemoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMemoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMemory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Memory
//
// @param request - CreateMemoryRequest
//
// @return CreateMemoryResponse
func (client *Client) CreateMemory(workspaceId *string, request *CreateMemoryRequest) (_result *CreateMemoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryResponse{}
	_body, _err := client.CreateMemoryWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建记忆Node
//
// @param request - CreateMemoryNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryNodeResponse
func (client *Client) CreateMemoryNodeWithOptions(workspaceId *string, memoryId *string, request *CreateMemoryNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateMemoryNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMemoryNode"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId)) + "/memoryNodes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建记忆Node
//
// @param request - CreateMemoryNodeRequest
//
// @return CreateMemoryNodeResponse
func (client *Client) CreateMemoryNode(workspaceId *string, memoryId *string, request *CreateMemoryNodeRequest) (_result *CreateMemoryNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryNodeResponse{}
	_body, _err := client.CreateMemoryNodeWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除智能体
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithOptions(workspaceId *string, appCode *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAgent"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/application/agents/" + tea.StringValue(openapiutil.GetEncodeParam(appCode))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除智能体
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(workspaceId *string, appCode *string) (_result *DeleteAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(workspaceId, appCode, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithOptions(CategoryId *string, WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCategory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/category/" + tea.StringValue(openapiutil.GetEncodeParam(CategoryId)) + "/"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DeleteCategoryResponse
func (client *Client) DeleteCategory(CategoryId *string, WorkspaceId *string) (_result *DeleteCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(CategoryId, WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(FileId *string, WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file/" + tea.StringValue(openapiutil.GetEncodeParam(FileId)) + "/"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(FileId *string, WorkspaceId *string) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(FileId, WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Index
//
// @param request - DeleteIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndexWithOptions(WorkspaceId *string, request *DeleteIndexRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndex"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Index
//
// @param request - DeleteIndexRequest
//
// @return DeleteIndexResponse
func (client *Client) DeleteIndex(WorkspaceId *string, request *DeleteIndexRequest) (_result *DeleteIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexResponse{}
	_body, _err := client.DeleteIndexWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除index doc
//
// @param tmpReq - DeleteIndexDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndexDocumentResponse
func (client *Client) DeleteIndexDocumentWithOptions(WorkspaceId *string, tmpReq *DeleteIndexDocumentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteIndexDocumentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteIndexDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocumentIds)) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, tea.String("DocumentIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentIdsShrink)) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIndexDocument"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/delete_index_document"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIndexDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除index doc
//
// @param request - DeleteIndexDocumentRequest
//
// @return DeleteIndexDocumentResponse
func (client *Client) DeleteIndexDocument(WorkspaceId *string, request *DeleteIndexDocumentRequest) (_result *DeleteIndexDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndexDocumentResponse{}
	_body, _err := client.DeleteIndexDocumentWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除memory
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(workspaceId *string, memoryId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMemory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除memory
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(workspaceId *string, memoryId *string) (_result *DeleteMemoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(workspaceId, memoryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除记忆Node
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryNodeResponse
func (client *Client) DeleteMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMemoryNodeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMemoryNode"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId)) + "/memoryNodes/" + tea.StringValue(openapiutil.GetEncodeParam(memoryNodeId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除记忆Node
//
// @return DeleteMemoryNodeResponse
func (client *Client) DeleteMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string) (_result *DeleteMemoryNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryNodeResponse{}
	_body, _err := client.DeleteMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档基本信息，包括文档名称、类型、状态等。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileResponse
func (client *Client) DescribeFileWithOptions(WorkspaceId *string, FileId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFileResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/file/" + tea.StringValue(openapiutil.GetEncodeParam(FileId)) + "/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档基本信息，包括文档名称、类型、状态等。
//
// @return DescribeFileResponse
func (client *Client) DescribeFile(WorkspaceId *string, FileId *string) (_result *DescribeFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFileResponse{}
	_body, _err := client.DescribeFileWithOptions(WorkspaceId, FileId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Index运行状态
//
// @param request - GetIndexJobStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIndexJobStatusResponse
func (client *Client) GetIndexJobStatusWithOptions(WorkspaceId *string, request *GetIndexJobStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIndexJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIndexJobStatus"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/job/status"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIndexJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Index运行状态
//
// @param request - GetIndexJobStatusRequest
//
// @return GetIndexJobStatusResponse
func (client *Client) GetIndexJobStatus(WorkspaceId *string, request *GetIndexJobStatusRequest) (_result *GetIndexJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIndexJobStatusResponse{}
	_body, _err := client.GetIndexJobStatusWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(workspaceId *string, memoryId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMemory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @return GetMemoryResponse
func (client *Client) GetMemory(workspaceId *string, memoryId *string) (_result *GetMemoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(workspaceId, memoryId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取记忆Node
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryNodeResponse
func (client *Client) GetMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMemoryNodeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMemoryNode"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId)) + "/memoryNodes/" + tea.StringValue(openapiutil.GetEncodeParam(memoryNodeId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取记忆Node
//
// @return GetMemoryNodeResponse
func (client *Client) GetMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string) (_result *GetMemoryNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryNodeResponse{}
	_body, _err := client.GetMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发布态智能体应用
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublishedAgentResponse
func (client *Client) GetPublishedAgentWithOptions(workspaceId *string, appCode *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPublishedAgentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetPublishedAgent"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/application/agents/" + tea.StringValue(openapiutil.GetEncodeParam(appCode))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPublishedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发布态智能体应用
//
// @return GetPublishedAgentResponse
func (client *Client) GetPublishedAgent(workspaceId *string, appCode *string) (_result *GetPublishedAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPublishedAgentResponse{}
	_body, _err := client.GetPublishedAgentWithOptions(workspaceId, appCode, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// ListCategory
//
// @param request - ListCategoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoryResponse
func (client *Client) ListCategoryWithOptions(WorkspaceId *string, request *ListCategoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryType)) {
		body["CategoryType"] = request.CategoryType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCategoryId)) {
		body["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCategory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/categories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// ListCategory
//
// @param request - ListCategoryRequest
//
// @return ListCategoryResponse
func (client *Client) ListCategory(WorkspaceId *string, request *ListCategoryRequest) (_result *ListCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCategoryResponse{}
	_body, _err := client.ListCategoryWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Chunk
//
// @param request - ListChunksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChunksResponse
func (client *Client) ListChunksWithOptions(WorkspaceId *string, request *ListChunksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListChunksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Filed)) {
		body["Filed"] = request.Filed
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		body["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChunks"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/list_chunks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChunksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Chunk
//
// @param request - ListChunksRequest
//
// @return ListChunksResponse
func (client *Client) ListChunks(WorkspaceId *string, request *ListChunksRequest) (_result *ListChunksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListChunksResponse{}
	_body, _err := client.ListChunksWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档列表
//
// @param request - ListFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileResponse
func (client *Client) ListFileWithOptions(WorkspaceId *string, request *ListFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFile"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/datacenter/files"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档列表
//
// @param request - ListFileRequest
//
// @return ListFileResponse
func (client *Client) ListFile(WorkspaceId *string, request *ListFileRequest) (_result *ListFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFileResponse{}
	_body, _err := client.ListFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Index文件
//
// @param request - ListIndexDocumentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexDocumentsResponse
func (client *Client) ListIndexDocumentsWithOptions(WorkspaceId *string, request *ListIndexDocumentsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndexDocumentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentName)) {
		query["DocumentName"] = request.DocumentName
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentStatus)) {
		query["DocumentStatus"] = request.DocumentStatus
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndexDocuments"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/list_index_documents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndexDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Index文件
//
// @param request - ListIndexDocumentsRequest
//
// @return ListIndexDocumentsResponse
func (client *Client) ListIndexDocuments(WorkspaceId *string, request *ListIndexDocumentsRequest) (_result *ListIndexDocumentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndexDocumentsResponse{}
	_body, _err := client.ListIndexDocumentsWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询pipeline
//
// @param request - ListIndicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndicesResponse
func (client *Client) ListIndicesWithOptions(WorkspaceId *string, request *ListIndicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIndicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexName)) {
		query["IndexName"] = request.IndexName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIndices"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/list_indices"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIndicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询pipeline
//
// @param request - ListIndicesRequest
//
// @return ListIndicesResponse
func (client *Client) ListIndices(WorkspaceId *string, request *ListIndicesRequest) (_result *ListIndicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIndicesResponse{}
	_body, _err := client.ListIndicesWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @param request - ListMemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoriesResponse
func (client *Client) ListMemoriesWithOptions(workspaceId *string, request *ListMemoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMemoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMemories"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMemoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取memory
//
// @param request - ListMemoriesRequest
//
// @return ListMemoriesResponse
func (client *Client) ListMemories(workspaceId *string, request *ListMemoriesRequest) (_result *ListMemoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoriesResponse{}
	_body, _err := client.ListMemoriesWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取记忆Node列表
//
// @param request - ListMemoryNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryNodesResponse
func (client *Client) ListMemoryNodesWithOptions(workspaceId *string, memoryId *string, request *ListMemoryNodesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMemoryNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["maxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMemoryNodes"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId)) + "/memoryNodes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMemoryNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取记忆Node列表
//
// @param request - ListMemoryNodesRequest
//
// @return ListMemoryNodesResponse
func (client *Client) ListMemoryNodes(workspaceId *string, memoryId *string, request *ListMemoryNodesRequest) (_result *ListMemoryNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryNodesResponse{}
	_body, _err := client.ListMemoryNodesWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询已发布的智能体应用列表
//
// @param request - ListPublishedAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublishedAgentResponse
func (client *Client) ListPublishedAgentWithOptions(workspaceId *string, request *ListPublishedAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPublishedAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["pageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishedAgent"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/application/agents"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishedAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询已发布的智能体应用列表
//
// @param request - ListPublishedAgentRequest
//
// @return ListPublishedAgentResponse
func (client *Client) ListPublishedAgent(workspaceId *string, request *ListPublishedAgentRequest) (_result *ListPublishedAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPublishedAgentResponse{}
	_body, _err := client.ListPublishedAgentWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 召回测试
//
// @param tmpReq - RetrieveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveResponse
func (client *Client) RetrieveWithOptions(WorkspaceId *string, tmpReq *RetrieveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RetrieveResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RetrieveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Images)) {
		request.ImagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Images, tea.String("Images"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Rerank)) {
		request.RerankShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rerank, tea.String("Rerank"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Rewrite)) {
		request.RewriteShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Rewrite, tea.String("Rewrite"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SearchFilters)) {
		request.SearchFiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchFilters, tea.String("SearchFilters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DenseSimilarityTopK)) {
		query["DenseSimilarityTopK"] = request.DenseSimilarityTopK
	}

	if !tea.BoolValue(util.IsUnset(request.EnableReranking)) {
		query["EnableReranking"] = request.EnableReranking
	}

	if !tea.BoolValue(util.IsUnset(request.EnableRewrite)) {
		query["EnableRewrite"] = request.EnableRewrite
	}

	if !tea.BoolValue(util.IsUnset(request.ImagesShrink)) {
		query["Images"] = request.ImagesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RerankShrink)) {
		query["Rerank"] = request.RerankShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RerankMinScore)) {
		query["RerankMinScore"] = request.RerankMinScore
	}

	if !tea.BoolValue(util.IsUnset(request.RerankTopN)) {
		query["RerankTopN"] = request.RerankTopN
	}

	if !tea.BoolValue(util.IsUnset(request.RewriteShrink)) {
		query["Rewrite"] = request.RewriteShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SaveRetrieverHistory)) {
		query["SaveRetrieverHistory"] = request.SaveRetrieverHistory
	}

	if !tea.BoolValue(util.IsUnset(request.SearchFiltersShrink)) {
		query["SearchFilters"] = request.SearchFiltersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SparseSimilarityTopK)) {
		query["SparseSimilarityTopK"] = request.SparseSimilarityTopK
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Retrieve"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/retrieve"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RetrieveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 召回测试
//
// @param request - RetrieveRequest
//
// @return RetrieveResponse
func (client *Client) Retrieve(WorkspaceId *string, request *RetrieveRequest) (_result *RetrieveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetrieveResponse{}
	_body, _err := client.RetrieveWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 知识索引
//
// @param tmpReq - SubmitIndexAddDocumentsJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIndexAddDocumentsJobResponse
func (client *Client) SubmitIndexAddDocumentsJobWithOptions(WorkspaceId *string, tmpReq *SubmitIndexAddDocumentsJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitIndexAddDocumentsJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitIndexAddDocumentsJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CategoryIds)) {
		request.CategoryIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CategoryIds, tea.String("CategoryIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DocumentIds)) {
		request.DocumentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentIds, tea.String("DocumentIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryIdsShrink)) {
		query["CategoryIds"] = request.CategoryIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DocumentIdsShrink)) {
		query["DocumentIds"] = request.DocumentIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitIndexAddDocumentsJob"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/add_documents_to_index"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitIndexAddDocumentsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 知识索引
//
// @param request - SubmitIndexAddDocumentsJobRequest
//
// @return SubmitIndexAddDocumentsJobResponse
func (client *Client) SubmitIndexAddDocumentsJob(WorkspaceId *string, request *SubmitIndexAddDocumentsJobRequest) (_result *SubmitIndexAddDocumentsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitIndexAddDocumentsJobResponse{}
	_body, _err := client.SubmitIndexAddDocumentsJobWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交索引任务
//
// @param request - SubmitIndexJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIndexJobResponse
func (client *Client) SubmitIndexJobWithOptions(WorkspaceId *string, request *SubmitIndexJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitIndexJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IndexId)) {
		query["IndexId"] = request.IndexId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitIndexJob"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/index/submit_index_job"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitIndexJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交索引任务
//
// @param request - SubmitIndexJobRequest
//
// @return SubmitIndexJobResponse
func (client *Client) SubmitIndexJob(WorkspaceId *string, request *SubmitIndexJobRequest) (_result *SubmitIndexJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitIndexJobResponse{}
	_body, _err := client.SubmitIndexJobWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新并发布智能体应用
//
// @param tmpReq - UpdateAndPublishAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAndPublishAgentResponse
func (client *Client) UpdateAndPublishAgentWithOptions(workspaceId *string, appCode *string, tmpReq *UpdateAndPublishAgentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateAndPublishAgentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAndPublishAgentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApplicationConfig)) {
		request.ApplicationConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationConfig, tea.String("applicationConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationConfigShrink)) {
		body["applicationConfig"] = request.ApplicationConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Instructions)) {
		body["instructions"] = request.Instructions
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		body["modelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAndPublishAgent"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/application/agents/" + tea.StringValue(openapiutil.GetEncodeParam(appCode))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAndPublishAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新并发布智能体应用
//
// @param request - UpdateAndPublishAgentRequest
//
// @return UpdateAndPublishAgentResponse
func (client *Client) UpdateAndPublishAgent(workspaceId *string, appCode *string, request *UpdateAndPublishAgentRequest) (_result *UpdateAndPublishAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAndPublishAgentResponse{}
	_body, _err := client.UpdateAndPublishAgentWithOptions(workspaceId, appCode, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新memory
//
// @param request - UpdateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(workspaceId *string, memoryId *string, request *UpdateMemoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMemory"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新memory
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(workspaceId *string, memoryId *string, request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(workspaceId, memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新记忆Node
//
// @param request - UpdateMemoryNodeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryNodeResponse
func (client *Client) UpdateMemoryNodeWithOptions(workspaceId *string, memoryId *string, memoryNodeId *string, request *UpdateMemoryNodeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateMemoryNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["content"] = request.Content
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMemoryNode"),
		Version:     tea.String("2023-12-29"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/memories/" + tea.StringValue(openapiutil.GetEncodeParam(memoryId)) + "/memoryNodes/" + tea.StringValue(openapiutil.GetEncodeParam(memoryNodeId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMemoryNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新记忆Node
//
// @param request - UpdateMemoryNodeRequest
//
// @return UpdateMemoryNodeResponse
func (client *Client) UpdateMemoryNode(workspaceId *string, memoryId *string, memoryNodeId *string, request *UpdateMemoryNodeRequest) (_result *UpdateMemoryNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryNodeResponse{}
	_body, _err := client.UpdateMemoryNodeWithOptions(workspaceId, memoryId, memoryNodeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
