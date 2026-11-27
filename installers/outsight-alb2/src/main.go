// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package main implements the Talos overlay installer for the Outsight ALB2.
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

const dtb = "freescale/imx8qm-iwg27m-alb2.dtb"

// The U-Boot environment lives in the eMMC user area at 0x400000 and 0x800000
// (redundant copies); Talos partitions must start past it.
const partitionOffsetSectors = 2048 * 16

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	adapter.Execute(ctx, &alb2Installer{})
}

type alb2Installer struct{}

type alb2ExtraOptions struct{}

func (i *alb2Installer) GetOptions(_ context.Context, _ alb2ExtraOptions) (overlay.Options, error) {
	return overlay.Options{
		Name: "outsight-alb2",
		KernelArgs: []string{
			"console=ttyLP4,115200",
			"sysctl.kernel.kexec_load_disabled=1",
			"talos.dashboard.disabled=1",
		},
		PartitionOptions: overlay.PartitionOptions{
			Offset: partitionOffsetSectors,
		},
	}, nil
}

// Install only places the DTB on the EFI partition. The bootloader lives in the
// eMMC boot0/boot1 hardware partitions, is signed for AHAB, and is updated A/B
// outside of Talos, so a Talos upgrade never touches the boot chain.
func (i *alb2Installer) Install(_ context.Context, options overlay.InstallOptions[alb2ExtraOptions]) error {
	src := filepath.Join(options.ArtifactsPath, "arm64/dtb", dtb)
	dst := filepath.Join(options.MountPrefix, "boot/EFI/dtb", dtb)

	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return err
	}

	return copy.File(src, dst)
}
