package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

func cg() {
	cgroups := "/sys/fs/cgroup/"
	mem := filepath.Join(cgroups, "memory")
	mgontainer := filepath.Join(mem, "gontainer")
	os.Mkdir(mgontainer, 0755)
	ioutil.WriteFile(filepath.Join(mgontainer, "memory.limit_in_bytes"), []byte("99942400"), 0700)
	ioutil.WriteFile(filepath.Join(mgontainer, "notify_on_release"), []byte("1"), 0700)
	ioutil.WriteFile(filepath.Join(mgontainer, "cgroup.procs"), []byte(pid), 0700)

	pids := filepath.Join(cgroups, "pids")
	pgontainer := filepath.Join(pids, "gontainer")
	os.Mkdir(pgontainer, 0755)
	pid := strconv.Itoa(os.Getpid())
	ioutil.WriteFile(filepath.Join(pids, "pids.max"), []byte("5"), 0700)
	ioutil.WriteFile(filepath.Join(pids, "notify_on_release"), []byte("1"), 0700)
	ioutil.WriteFile(filepath.Join(pids, "cgroup.procs"), []byte(pid), 0700)
}
