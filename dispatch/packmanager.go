package dispatch

import (
	"runtime"
)

// PackManager represents a PACkage MANager
type PackManager interface {
	Q([]string) error
	Qc([]string) error
	Qe([]string) error
	Qi([]string) error
	Qk([]string) error
	Ql([]string) error
	Qm([]string) error
	Qo([]string) error
	Qp([]string) error
	Qs([]string) error
	Qu([]string) error
	R([]string) error
	Rn([]string) error
	Rns([]string) error
	Rs([]string) error
	S([]string) error
	Sc([]string) error
	Scc([]string) error
	Sccc([]string) error
	Sg([]string) error
	Si([]string) error
	Sii([]string) error
	Sl([]string) error
	Ss([]string) error
	Su([]string) error
	Suy([]string) error
	Sw([]string) error
	Sy([]string) error
	U([]string) error
}

// DetectPackManager detects the package manager in use
// TODO: Make this function REALLY detect package managers
func DetectPackManager(dryRun bool, noConfirm bool) PackManager {
	switch runtime.GOOS {
	case "windows":
		return &Chocolatey{dryRun, noConfirm}
	case "darwin":
		return &Homebrew{dryRun}
	default:
		return &Unknown{dryRun, noConfirm}
	}
}
