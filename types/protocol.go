package types

type Protocol string

var GIT Protocol = "git"
var SSH Protocol = "ssh"
var HTTP Protocol = "http"
var HTTPS Protocol = "https"

func CheckProto(str string) Protocol {
	switch str {
	case "git":
		return GIT
	case "ssh":
		return SSH
	case "http":
		return HTTP
	case "https":
		return HTTPS
	default:
		return HTTPS
	}
}
