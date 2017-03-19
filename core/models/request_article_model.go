package models



type RequestArticleModel struct{

	Title      string `form: "title" boil:"title" json:"title" toml:"title" yaml:"title"`
	Content    string `form: "title" boil:"content" json:"content" toml:"content" yaml:"content"`
}
