package url_shortener

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}
