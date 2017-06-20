package structs


type GenericJSON struct {
	Stuff string `json:"some_stuff"`
}

type ErrorMsg struct {
	Msg string `json:"error"`
}

type RequestDataJSON struct {
	Data string `json:"thread_link"`
}

type ResponseJSON struct {
	Files []File `json:"files"`
}

type Board struct {
	Threads []Thread
	Error error
}

type Thread struct {
	Posts []InnerPost
}

type InnerPost struct {
	Banned int `json:"banned"`
	Comment string `json:"comment"`
	Files []File
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Thumbnail string `json:"thumbnail"`
}