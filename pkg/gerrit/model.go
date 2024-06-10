package gerrit

type Source struct {
	ID       int    `cdevents:"user_id"`
	FullName string `cdevents:"user_name"`
}

type Destination struct {
	UserID int    `cdevents:"user_id"`
	Name   string `cdevents:"user_name,omitempty"`
}
