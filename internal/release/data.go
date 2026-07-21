package release

// Version defines manboster's application version.
const Version = "0.3.0"

// APILevel defines the current level(Tool) supported in Manboster.
const APILevel = 1

// V indicates config's version, now is 0
const V = 0

var (
	BuildCommit    string = "unknown"
	BuildTime      string = "unknown"
	CurrentChannel        = "unknown"
	Package               = "none"
)
