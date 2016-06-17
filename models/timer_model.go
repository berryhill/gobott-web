package models

type Timer struct {
	Name 		string                `json:"name"`
	Seconds 	int                   `json:"seconds"`
}

func (t *Timer)SetTimer(seconds int) error {
	t.Seconds = seconds

	return nil
}