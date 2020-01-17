package backend

import "github.com/openlibre/wikinota/backend/initbackend"

func Hello() string {
	return "Hello, from backend!" + initbackend.Hello()
}
