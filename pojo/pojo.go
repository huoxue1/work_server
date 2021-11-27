package pojo

type Work struct {
	Id      int64  `json:"id" xorm:"'id' pk autoincr"`
	Name    string `json:"name" xorm:"'name'"`
	EndTime int64  `json:"end_time" xorm:"'end_time'"` //nolint:tagliatelle
	Created int64  `json:"created" xorm:"created"`     //nolint:gofmt
	Deleted int64  `json:"deleted" xorm:"deleted"`
}
type Files struct {
	Id         int64  `json:"id" xorm:"'id' pk autoincr"`
	Data       string `json:"data" xorm:"'data'"`
	WorkID     int    `json:"work_id" xorm:"'work_id'"`
	FileName   string `json:"file_name" xorm:"'file_name'"`
	Size       int64  `json:"size" xorm:"'size'"`
	UploadTime int64  `json:"upload_time" xorm:"updated"`
	Token      string `json:"token" xorm:"'token'"`
	Created    int64  `json:"created" xorm:"created"`
	Deleted    int64  `json:"deleted" xorm:"deleted"`
}
