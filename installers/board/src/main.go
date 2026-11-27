// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/siderolabs/go-copy/copy"
	"github.com/siderolabs/talos/pkg/machinery/overlay"
	"github.com/siderolabs/talos/pkg/machinery/overlay/adapter"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	adapter.Execute(ctx, &BoardInstaller{})
}

type BoardInstaller struct{}

type boardExtraOptions struct {
	Console []string `yaml:"console,omitempty"`
}

func (i *BoardInstaller) GetOptions(_ context.Context, extra boardExtraOptions) (overlay.Options, error) {
	kernelArgs := []string{
		"console=tty0",
		"sysctl.kernel.kexec_load_disabled=1",
		"talos.dashboard.disabled=1",
	}

	kernelArgs = append(kernelArgs, extra.Console...)

	return overlay.Options{
		Name:       "board",
		KernelArgs: kernelArgs,
	}, nil
}

func (i *BoardInstaller) Install(_ context.Context, options overlay.InstallOptions[boardExtraOptions]) error {
	return copy.File(filepath.Join(options.ArtifactsPath, "arm64/u-boot/board/u-boot.bin"), filepath.Join(options.MountPrefix, "/boot/EFI/u-boot.bin"))
}
