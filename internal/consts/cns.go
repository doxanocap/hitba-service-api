package consts

import "time"

const (
	AppDevelopmentEnv = "development"
	AppProductionEnv  = "production"
)

// App constants
const (
	// DateFormat default date format in the domain of Tasbih-API
	// to handle daily counter via this format
	DateFormat = "2006-01-02"
	// SingleSessionDuration when counting has started, all counted values
	// during this session time will be recorded as one session
	SingleSessionDuration = time.Hour

	// TokenMaxAge param for tokens time to live
	TokenMaxAge = 30 * 24 * 60 * 60 * 1000
	TokenPath   = "/"

	// AuthPrefix default
	AuthPrefix = "Bearer "

	MailingAddress = "mailing_address"
	MailSentTo     = "mail_sent_to"
)

var (
	ByteNewLine = []byte{'\n'}
	ByteSpace   = []byte{' '}
)

const (
	PictureFormatPNG  = "png"
	PictureFormatJPEG = "jpeg"
	PictureFormatIMG  = "img"
	PictureFormatJPG  = "jpg"
)

// Types
const (
	// NilString default nil value string across domain of the project
	NilString = ""
	// NilInt default nil value int across domain of the project
	NilInt = 0
)

// NilByte default nil value []byte across domain of the project
var (
	NilByte []byte
)
