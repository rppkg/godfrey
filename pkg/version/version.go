package version

var (
	// GitVersion 是语义化的版本号.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	GitCommit = "$Format:%H$"
)
