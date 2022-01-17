package logger

// LogConf -
type LogConf struct {
	// ServiceName - 服務名稱
	ServiceName string `json:",optional"`

	// Mode - 日志模式(console - 輸出到console, file - 輸出到文件)
	Mode string `json:",default=console,options=console|file"`

	// Path - 儲存路徑
	Path string `json:",default=logs"`

	// Level - 級別
	Level string `json:",default=info,options=info|error|severe"`

	// Compress - 使否開啟 gzip 壓縮
	Compress bool `json:",optional"`

	// KeepDays - 保存幾天
	KeepDays int `json:",optional"`
}
