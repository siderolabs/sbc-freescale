# sbc-freescale

Talos Linux overlays for Freescale/NXP single-board computers.

| Overlay Name  | Board                                    |
| ------------- | ---------------------------------------- |
| outsight-alb2 | Outsight ALB2 (iWave iWG27M i.MX8QM SoM) |

## Layout

Each board has its own:

* installer in `installers/<board>/`, listed in `installers/pkg.yaml` and `go.work`;
* profile in `profiles/<board>/`, listed in `profiles/pkg.yaml`;
* build artifacts in `artifacts/<board>/<component>/`, as a `<component>-<board>`
  stage listed in `installers/pkg.yaml`.

Artifacts shared by several boards of the same SoC go in
`artifacts/<component>/<soc>/`.

## Boards

### outsight-alb2

The overlay installs only the device tree. The bootloader is signed for AHAB,
lives in the eMMC boot partitions and is updated outside of Talos. Talos
partitions start at 16 MiB to keep the U-Boot environment.

The ALB2 device tree is not in Linux: `artifacts/outsight-alb2/dtb` builds it
against the SoC `.dtsi` files of the Linux version set in `Pkgfile`.

## Adding a board

* Add the installer, profile and artifacts as described in [Layout](#layout).
* Add the board to the table above and to `MAINTAINERS.md`.

## Building

* `make help` describes setting up a `buildx` builder.
* `make sbc-freescale` builds the overlay image.
* `make rekres` regenerates the kres-managed files after editing `.kres.yaml`.
