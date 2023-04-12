package main

type ServerConfig struct {
	Key           string `json:"key"`
	LocalName     string `json:"local_name"`
	DbType        string `json:"dbType"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	HasRecordPwd  bool   `json:"has_record_pwd"`
	HasUseSSH     bool   `json:"hasUseSSH"`
	SshHost       string `json:"ssh_host"`
	SshPort       string `json:"ssh_port"`
	SshUser       string `json:"ssh_user"`
	HasSshKeyfile bool   `json:"has_ssh_keyfile"`
	SshKeyfile    string `json:"ssh_keyfile"`
	HasSshPass    bool   `json:"has_ssh_pass"`
	SshPassword   string `json:"ssh_password"`
	ConState      bool   `json:"conState"`
}

func (a *App) ReturnSuccess(Data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"State": true,
		"Data":  Data,
	}
}

func (a *App) ReturnError(msg string) map[string]interface{} {
	return map[string]interface{}{
		"State":   false,
		"Message": msg,
	}
}

type TreeData struct {
	Key      string     `json:"key"`
	Label    string     `json:"label"`
	Children []TreeData `json:"children"`
	ConState bool       `json:"conState"`
	ObjType  string     `json:"obj_type"`
}
