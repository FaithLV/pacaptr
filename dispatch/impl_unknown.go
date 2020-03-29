package dispatch

// Unknown package manager config.
// This is a special "blank" PackManager reserved for copy-paste convenience
// as well as unknown package manager handling.
type Unknown struct {
	DryRun    bool
	NoConfirm bool
}

// For method implementation see: https://golang.org/src/os/exec/example_test.go
// For method explanation see: https://wiki.archlinux.org/index.php/Pacman/Rosetta
// and https://wiki.archlinux.org/index.php/Pacman

// RunIfNotDry prints out the command if DryRun, else it runs the command.
func (pm *Unknown) RunIfNotDry(cmd []string) (err error) {
	if pm.DryRun {
		PrintCommand(cmd)
		return
	}
	return RunCommand(cmd)
}

// Q generates a list of installed packages.
func (pm *Unknown) Q(kw []string) (err error) {
	return NotImplemented()
}

// Qc shows the changelog of a package.
func (pm *Unknown) Qc(kw []string) (err error) {
	return NotImplemented()
}

// Qe lists packages installed explicitly (not as dependencies).
func (pm *Unknown) Qe(kw []string) (err error) {
	return NotImplemented()
}

// Qi displays local package information: name, version, description, etc.
func (pm *Unknown) Qi(kw []string) (err error) {
	return NotImplemented()
}

// Qk verifies one or more packages.
func (pm *Unknown) Qk(kw []string) (err error) {
	return NotImplemented()
}

// Ql displays files provided by local package.
func (pm *Unknown) Ql(kw []string) (err error) {
	return NotImplemented()
}

// Qm lists packages that are installed but are not available in any installation source (anymore).
func (pm *Unknown) Qm(kw []string) (err error) {
	return NotImplemented()
}

// Qo queries the package which provides FILE.
func (pm *Unknown) Qo(kw []string) (err error) {
	return NotImplemented()
}

// Qp queries a package supplied on the command line rather than an entry in the package management database.
func (pm *Unknown) Qp(kw []string) (err error) {
	return NotImplemented()
}

// Qs searches locally installed package for names or descriptions.
func (pm *Unknown) Qs(kw []string) (err error) {
	return NotImplemented()
}

// Qu lists packages which have an update available.
func (pm *Unknown) Qu(kw []string) (err error) {
	return NotImplemented()
}

// R removes a single package, leaving all of its dependencies installed.
func (pm *Unknown) R(kw []string) (err error) {
	return NotImplemented()
}

// Rn removes a package and skips the generation of configuration backup files.
func (pm *Unknown) Rn(kw []string) (err error) {
	return NotImplemented()
}

// Rns removes a package and its dependencies which are not required by any other installed package, and skips the generation of configuration backup files.
func (pm *Unknown) Rns(kw []string) (err error) {
	return NotImplemented()
}

// Rs removes a package and its dependencies which are not required by any other installed package.
func (pm *Unknown) Rs(kw []string) (err error) {
	return NotImplemented()
}

// S installs one or more packages by name.
func (pm *Unknown) S(kw []string) (err error) {
	return NotImplemented()
}

// Sc removes all the cached packages that are not currently installed, and the unused sync database.
func (pm *Unknown) Sc(kw []string) (err error) {
	return NotImplemented()
}

// Scc removes all files from the cache.
func (pm *Unknown) Scc(kw []string) (err error) {
	return NotImplemented()
}

// Sccc ...
// ! What is this?
func (pm *Unknown) Sccc(kw []string) (err error) {
	return NotImplemented()
}

// Sg lists all packages belonging to the GROUP.
func (pm *Unknown) Sg(kw []string) (err error) {
	return NotImplemented()
}

// Si displays remote package information: name, version, description, etc.
func (pm *Unknown) Si(kw []string) (err error) {
	return NotImplemented()
}

// Sii displays packages which require X to be installed, aka reverse dependencies.
func (pm *Unknown) Sii(kw []string) (err error) {
	return NotImplemented()
}

// Sl displays a list of all packages in all installation sources that are handled by the packages management.
func (pm *Unknown) Sl(kw []string) (err error) {
	return NotImplemented()
}

// Ss searches for package(s) by searching the expression in name, description, short description.
func (pm *Unknown) Ss(kw []string) (err error) {
	return NotImplemented()
}

// Su updates outdated packages.
func (pm *Unknown) Su(kw []string) (err error) {
	return NotImplemented()
}

// Suy refreshes the local package database, then updates outdated packages.
func (pm *Unknown) Suy(kw []string) (err error) {
	return NotImplemented()
}

// Sw retrieves all packages from the server, but does not install/upgrade anything.
func (pm *Unknown) Sw(kw []string) (err error) {
	return NotImplemented()
}

// Sy refreshes the local package database.
func (pm *Unknown) Sy(kw []string) (err error) {
	return NotImplemented()
}

// U upgrades or adds package(s) to the system and installs the required dependencies from sync repositories.
func (pm *Unknown) U(kw []string) (err error) {
	return NotImplemented()
}
