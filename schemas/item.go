package schemas

type Item struct {
	Checked     bool   `json:"checked" sql:"bit not null default 0"`
	Description string `json:"description" sql:"varchar(255) not null default ''"`
}
