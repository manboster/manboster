package agent

type Store struct {
	ID       string   `json:"id"`
	Data     Agent    `json:"data"`
	Father   string   `json:"father"`
	Children []string `json:"children"`
}

func Restore() {

}
