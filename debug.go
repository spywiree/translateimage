package translateimage

type DebugOptions struct {
	Enabled   bool
	VideoPath string
}

var Debug = DebugOptions{
	Enabled:   false,
	VideoPath: "videos",
}
