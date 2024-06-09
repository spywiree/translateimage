package translateimage

type Options struct {
	DebugMode bool
	VideoPath string
}

func getOptions(options []Options) Options {
	if len(options) > 0 {
		return options[0]
	} else {
		return Options{
			DebugMode: false,
			VideoPath: "videos",
		}
	}
}
