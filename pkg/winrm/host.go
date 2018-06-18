package winrm

// Representation of a remote machine that WinRM can connect to.
type Host struct {
	Host string
	User string
	Pass string
}
