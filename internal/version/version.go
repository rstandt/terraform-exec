package version

const version = "0.16.1"

// ModuleVersion returns the current version of the github.com/rstandt/terraform-exec Go module.
// This is a function to allow for future possible enhancement using debug.BuildInfo.
func ModuleVersion() string {
	return version
}
