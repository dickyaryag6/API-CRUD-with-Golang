package models

type Article struct {
  ArticleID      string `json:"ArticleID"  bson:"ArticleID"`
  Title          string `json:"Title"     bson:"Title"`
  Desc           string `json:"Desc"      bson:"Desc"`
  Content        string `json:"Content"   bson:"Content"`
}
