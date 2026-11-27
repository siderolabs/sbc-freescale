# sbc-freescale

Talos Linux overlays for Freescale/NXP single-board computers.

No board is supported yet: `board` is a placeholder overlay that builds
U-Boot from `board_defconfig` and copies `u-boot.bin` to `/boot/EFI`.

## Adding a board

* Rename `board` to the overlay name in `installers/`, `profiles/`,
  `artifacts/u-boot/pkg.yaml`, `installers/pkg.yaml`, `profiles/pkg.yaml`
  and `go.work`.
* Set the U-Boot defconfig and boot firmware in `artifacts/u-boot/pkg.yaml`.
* Implement the install logic in `installers/<board>/src/main.go`.
* Add the board to `MAINTAINERS.md`.

## Building

* `make help` describes setting up a `buildx` builder.
* `make sbc-freescale` builds the overlay image.
* `make rekres` regenerates the kres-managed files after editing `.kres.yaml`.
