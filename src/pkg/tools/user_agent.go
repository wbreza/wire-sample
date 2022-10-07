package tools

import "fmt"

type UserAgent struct {
	Template string
	Version  *Version
}

func (ua *UserAgent) String() string {
	return fmt.Sprintf("azd version:%s, template:%s", ua.Version, ua.Template)
}

type Version struct {
	Major uint
	Minor uint
	Patch uint
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
