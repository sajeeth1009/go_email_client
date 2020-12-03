package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailTemplate struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	MessageType     string              `bson:"messageType"`
	StudyKey        string              `bson:"studyKey,omitempty"`
	DefaultLanguage string              `bson:"defaultLanguage"`
	HeaderOverrides *HeaderOverrides    `bson:"headerOverrides"`
	Translations    []LocalizedTemplate `bson:"translations"`
}

type HeaderOverrides struct {
	From      string   `bson:"from"`
	Sender    string   `bson:"sender"`
	ReplyTo   []string `bson:"replyTo"`
	NoReplyTo bool     `bson:"noReplyTo"`
}

type LocalizedTemplate struct {
	Lang        string `bson:"languageCode"`
	Subject     string `bson:"subject"`
	TemplateDef string `bson:"templateDef"`
}
