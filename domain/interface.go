package domain

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Config "github.com/joaosoft/clean-infrastructure/s3/config"
	"github.com/joaosoft/clean-infrastructure/utils/errors"

	"github.com/joaosoft/clean-infrastructure/datatable/database"
	"github.com/joaosoft/clean-infrastructure/datatable/elastic_search"

	"github.com/joaosoft/clean-infrastructure/utils/database/session"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	appConfig "github.com/joaosoft/clean-infrastructure/app/config"
	"github.com/joaosoft/clean-infrastructure/database/config"
	elasticSearchConfig "github.com/joaosoft/clean-infrastructure/elastic_search/config"
	grpcConfig "github.com/joaosoft/clean-infrastructure/grpc/config"
	httpConfig "github.com/joaosoft/clean-infrastructure/http/config"
	loggerConfig "github.com/joaosoft/clean-infrastructure/logger/config"
	rabbitmqConfig "github.com/joaosoft/clean-infrastructure/rabbitmq/config"
	redisConfig "github.com/joaosoft/clean-infrastructure/redis/config"
	sqsConfig "github.com/joaosoft/clean-infrastructure/sqs/config"
	stateMachineDomain "github.com/joaosoft/clean-infrastructure/state_machine/instance/domain"
	tracerConfig "github.com/joaosoft/clean-infrastructure/tracer/config"
	msg "github.com/joaosoft/clean-infrastructure/utils/pagination"
	"github.com/opentracing/opentracing-go"
	"github.com/redis/go-redis/v9"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

// IApp App interface
type IApp interface {
	IService

	// Config gets the configuration
	Config() *appConfig.Config
	// ConfigFile the configuration file
	ConfigFile() string

	// WithLogger sets the logger
	WithLogger(logger ILogger) IApp
	// Logger gets the logger
	Logger() ILogger
	// WithDatabase sets the database
	WithDatabase(database IDatabase) IApp
	// Database gets the database
	Database() IDatabase
	// WithRabbitmq sets the rabbitmq
	WithRabbitmq(rabbitmq IRabbitMQ) IApp
	// Rabbitmq gets the rabbitmq
	Rabbitmq() IRabbitMQ
	// WithSQS sets the sqs
	WithSQS(sqs ISQS) IApp
	// SQS gets the sqs
	SQS() ISQS
	// WithRedis sets the redis
	WithRedis(redis IRedis) IApp
	// Redis gets the redis
	Redis() IRedis
	// WithElasticSearch sets the elastic search
	WithElasticSearch(elasticSearch IElasticSearch) IApp
	// ElasticSearch gets the elastic search
	ElasticSearch() IElasticSearch
	// WithGrpc sets the grpc
	WithGrpc(grpc IGrpc) IApp
	// Grpc gets the grpc
	Grpc() IGrpc
	// WithHttp sets the http
	WithHttp(http IHttp) IApp
	// Http gets the http
	Http() IHttp
	// WithValidator sets custom validators
	WithValidator(validator IValidator) IApp
	// Validator gets the validator
	Validator() IValidator
	// WithAws sets the aws connection
	WithAws(aws IAws) IApp
	// Aws gets the http
	Aws() IAws
	// WithTracer sets the tracer
	WithTracer(tracer ITracer) IApp
	// Tracer gets the tracer
	Tracer() ITracer
	// WithDatatable sets the datatable
	WithDatatable(datatable IDatatable) IApp
	// Datatable gets the datatable
	Datatable() IDatatable
	// WithAws sets the aws connection
	WithS3(s3 IS3) IApp
	// S3 gets the s3 Connection
	S3() IS3
	// StateMachine gets the state machine service
	StateMachine() IStateMachineService
	// WithAdditionalConfigType sets an additional config type
	WithAdditionalConfigType(obj interface{}) IApp
}

// IService service interface
type IService interface {
	// Name name of the service
	Name() string
	// Start starts the service
	Start() error
	// Stop stops the service
	Stop() error
	// Started true if service started
	Started() bool
}

// IMiddleware the interface of the middlewares
type IMiddleware interface {
	RegisterMiddlewares()
	GetHandlers() []gin.HandlerFunc
}

// IController the interface of the controllers
type IController interface {
	App() IApp
	Register()
	Json(ctx IContext, data interface{}, err ...error)
	JsonWithPagination(ctx IContext, data interface{}, pagination *msg.Pagination, err ...error)
}

// IHttp the interface for the http service
type IHttp interface {
	IService

	// ConfigFile gets the configuration file name
	ConfigFile() string
	// Config gets the configurations
	Config() *httpConfig.Config

	// WithController adds a controller
	WithMiddleware(controller IMiddleware) IHttp
	// WithController adds a controller
	WithController(controller IController) IHttp
	// WithRouter sets the router
	WithRouter(router *gin.Engine) IHttp
	// Router gets the router
	Router() *gin.Engine
}

// ILogger logger service interface
type ILogger interface {
	IService

	// ConfigFile gets the configuration file
	ConfigFile() string
	// Config gets the configurations
	Config() *loggerConfig.Config

	//Log the error
	Log() ILogging

	// Database Log
	DBLog(error) error
	// SQS Log
	SQSLog(error) error
	//Elastic Log
	ElasticLog(error) error
	//Redis Log
	RedisLog(error) error
}

// IDatabase database service interface
type IDatabase interface {
	IService

	// ConfigFile gets the configuration file
	ConfigFile() string
	// Config gets the configurations
	Config() *config.Config

	// Read the read connection
	Read() session.ISession
	// Write the write connection
	Write() session.ISession
}

// IElasticSearch elastic search service interface
type IElasticSearch interface {
	IService

	ConfigFile() string
	Config() *elasticSearchConfig.Config

	Client() *elasticsearch.Client
}

// IRabbitMQ rabbitmq service interface
type IRabbitMQ interface {
	IService

	// ConfigFile gets the configuration file
	ConfigFile() string
	// Config gets the configurations
	Config() *rabbitmqConfig.Config

	// Produce produces to the rabbitmq
	Produce(message string, exchange string, routingKey string) error
	// Consume consumes from the rabbitmq
	Consume(app IApp, queues string, handlers map[string]func(msg amqp.Delivery) bool)
	// WithConsumer adds a consumer to the rabbitmq
	WithConsumer(consumer IRabbitMQConsumer) IRabbitMQ
}

// ISQS sqs service interface
type ISQS interface {
	IService

	// WithAdditionalConfigType sets an additional config type
	WithAdditionalConfigType(obj interface{}) ISQS
	// ConfigFile gets the configuration file
	ConfigFile() string
	// Config gets the configurations
	Config() *sqsConfig.Config
	// WithConsumer adds a consumer to the rabbitmq
	WithConsumer(consumer ISQSConsumer) ISQS
	// Connection
	Connection(name string) ISQSConnection
}

type ISQSConnection interface {
	// Connect connect
	Connect() error
	// Produce produces to the sqs
	Produce(queue string, messageAttributes map[string]*sqs.MessageAttributeValue, messages ...string) error
	// Consume consumes from the sqs
	Consume(maskedQueue string, consumer ISQSConsumer)
}

// IS3 s3 service interface
type IS3 interface {
	IService

	// ConfigFile gets the configuration file
	ConfigFile() string
	// Config gets the configurations
	Config() *s3Config.Config
	// Client S3 client
	Client() *s3.Client
}

// IRedis redis service interface
type IRedis interface {
	IService

	// WithAdditionalConfigType sets an additional config type
	WithAdditionalConfigType(obj interface{}) IRedis

	// ConfigFile gets the configuration file
	ConfigFile() string
	// Config gets the configurations
	Config() *redisConfig.Config

	// Client gets the redis client
	Client() *redis.Client
}

// IGrpc grpc service interface
type IGrpc interface {
	IService

	InitServer() error
	InitClients() error

	// Config gets the configurations
	Config() *grpcConfig.Configs
	// ConfigFile gets the configuration file
	ConfigFile() string
	// GetClient gets the client by name
	GetClient(name string) (conn *grpc.ClientConn)
	// GetServer gets the server
	GetServer() (conn *grpc.Server, err error)
	// WithController adds a controller
	WithController(controller IController) IGrpc
}

// IConsumer defines the rabbitmq consumers interface
type IRabbitMQConsumer interface {
	// GetHandlers gets the handlers
	GetHandlers() map[string]func(msg amqp.Delivery) bool
	// GetQueue gets the queue
	GetQueue() string
}

type ISQSConsumer interface {
	// GetHandlers gets the handlers
	GetHandlers() map[string]func(msg *sqs.Message) bool
	// GetConnection gets the connection name
	GetConnection() string
	// GetQueue gets the queue
	GetQueue() string
	// GetAttributeNames gets the queue attribute names
	GetAttributeNames() []*string
	// GetMessageAttributeNames gets the message attribute names
	GetMessageAttributeNames() []*string
}

// ILogging defines what our logging lib need
type ILogging interface {
	// Log the error
	Do(err error, info ...*LoggerInfo)
	// Log Multi
	Multi(err []error, info ...*LoggerInfo)
	// Frontend
	Frontend(error string, level errors.Level, fe *Frontend)
	// Initialize
	Init(cgf loggerConfig.Config) ILogging
}

// IValidator Interface
type IValidator interface {
	IService
	// AddFieldValidators adds a custom field validator
	AddFieldValidators(v ...IFieldValidator) IValidator
	// AddStructValidators adds a custom struct validator
	AddStructValidators(v ...IStructValidator) IValidator
	// Validate validates the struct
	Validate(v any) error
}

type IFieldValidator interface {
	Tag() string
	Func(a IApp) validator.Func
}

type IStructValidator interface {
	Struct() any
	Func(a IApp) validator.StructLevelFunc
}

// IAWS interface
type IAws interface {
	IService
	// AddValidator adds a custom validator
	Connection() *aws.Config
}

// ITracer interface
type ITracer interface {
	IService

	// Config gets the configurations
	Config() *tracerConfig.Config
	// ConfigFile gets the configuration file
	ConfigFile() string

	opentracing.Tracer
}

// IStateMachineService interface
type IStateMachineService interface {
	IService

	Get(name string) (stateMachineDomain.IStateMachine, error)
	AddOrchestrator(stateMachine IStateMachineOrchestrator)
}

// IStateMachine interface
type IStateMachine interface {
	GetName() string
	GetStateMachine() stateMachineDomain.IStateMachine
}

// IStateMachineOrchestrator interface
type IStateMachineOrchestrator interface {
	CheckHandlers() map[string]map[string]stateMachineDomain.CheckHandler
	ExecuteHandlers() map[string]map[string]stateMachineDomain.ExecuteHandler
	OnSuccessHandlers() map[string]map[string]stateMachineDomain.OnSuccessHandler
	OnErrorHandlers() map[string]map[string]stateMachineDomain.OnErrorHandler
	GetStateMachines() []IStateMachine
}

// IContext interface
type IContext interface {
	context.Context
	FullPath() string
	Next()
	Set(key string, value any)
	Get(key string) (value any, exists bool)
	MustGet(key string) any
	GetString(key string) (s string)
	GetBool(key string) (b bool)
	GetInt(key string) (i int)
	GetInt64(key string) (i64 int64)
	GetUint(key string) (ui uint)
	GetUint64(key string) (ui64 uint64)
	GetFloat64(key string) (f64 float64)
	GetTime(key string) (t time.Time)
	GetDuration(key string) (d time.Duration)
	GetStringSlice(key string) (ss []string)
	GetStringMap(key string) (sm map[string]any)
	GetStringMapString(key string) (sms map[string]string)
	GetStringMapStringSlice(key string) (smss map[string][]string)
	Param(key string) string
	AddParam(key, value string)
	Query(key string) (value string)
	DefaultQuery(key, defaultValue string) string
	GetQuery(key string) (string, bool)
	QueryArray(key string) (values []string)
	GetQueryArray(key string) (values []string, ok bool)
	QueryMap(key string) (dicts map[string]string)
	GetQueryMap(key string) (map[string]string, bool)
	PostForm(key string) (value string)
	DefaultPostForm(key, defaultValue string) string
	GetPostForm(key string) (string, bool)
	PostFormArray(key string) (values []string)
	GetPostFormArray(key string) (values []string, ok bool)
	PostFormMap(key string) (dicts map[string]string)
	GetPostFormMap(key string) (map[string]string, bool)
	FormFile(name string) (*multipart.FileHeader, error)
	MultipartForm() (*multipart.Form, error)
	SaveUploadedFile(file *multipart.FileHeader, dst string) error
	Bind(obj any) error
	BindJSON(obj any) error
	BindQuery(obj any) error
	BindHeader(obj any) error
	BindUri(obj any) error
	MustBindWith(obj any, b binding.Binding) error
	ShouldBind(obj any) error
	ShouldBindJSON(obj any) error
	ShouldBindQuery(obj any) error
	ShouldBindHeader(obj any) error
	ShouldBindUri(obj any) error
	ShouldBindWith(obj any, b binding.Binding) error
	ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error)
	ClientIP() string
	RemoteIP() string
	ContentType() string
	Status(code int)
	Header(key, value string)
	GetHeader(key string) string
	GetRawData() ([]byte, error)
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	SetSameSite(http.SameSite)
	Cookie(name string) (string, error)
	IndentedJSON(code int, obj any)
	JSONP(code int, obj any)
	JSON(code int, obj any)
	String(code int, format string, values ...any)
	Redirect(code int, location string)
	Data(code int, contentType string, data []byte)
	DataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string)
	File(filepath string)
	FileFromFS(filepath string, fs http.FileSystem)
	FileAttachment(filepath, filename string)
	Stream(step func(w io.Writer) bool) bool
	SetAccepted(formats ...string)
	Values(key any) any
	Params() gin.Params
	Keys() map[string]any
	Request() *http.Request
	Response() gin.ResponseWriter
	SetIdMarket(int)
	SetIdBu(int)
	SetIdShop(int)
	SetIdFascia(int)
	SetIdUserExternal(int)
	SetUsername(string)
	SetLanguageCode(string)
	SetBody([]byte)
	SetAuthorizations([]string)
	GetIdMarket() int
	GetIdBu() int
	GetIdShop() int
	GetIdFascia() int
	GetIdUserExternal() int
	GetUsername() string
	GetLanguageCode() string
	GetBody() []byte
	GetAuthorizations() []string
	Abort()
	AddMeta(meta any) IContext
	AddPagination(pagination *msg.Pagination) IContext
	GetMeta() any
	GetPagination() *msg.Pagination
	FromGrpc(ctx context.Context) IContext
	ToGrpc() context.Context
}

// IDatatable datatable service interface
type IDatatable interface {
	IService
	Database() database.IDatabase
	Elastic() elastic_search.IElastic
}

type FallbackReader interface {
	ReadLines() ([]string, error)
}

type FallbackWriter interface {
	io.Writer
	Remove() error
}
