<a href="assets/neofetch.txt"><img align="center" width="100%" alt="neofetch" src="assets/neofetch.png"></a>

<p align="center">
    <a href="https://github.com/PryosCode/center/actions"><img alt="Build" src="https://github.com/PryosCode/center/actions/workflows/go.yml/badge.svg"></a>
    <a href="https://github.com/PryosCode/center/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/PryosCode/center?label=License"></a>
    <a href="https://discord.gg/bF2GRHq"><img alt="Discord" src="https://discord.com/api/guilds/350302354639290379/widget.png"></a>
</p>

# center

A cli tool to center text

## Download

[![Version](https://img.shields.io/github/v/release/PryosCode/center?label=Version)](https://github.com/PryosCode/center/releases)

## Usage

```bash
Usage:
  center [text]... [flags]

Flags:
  -h, --help      help for center
  -v, --version   version for center
```

## Examples

```bash
center "This Text is centered"
```

```bash
center "This Text is centered" "With multiple lines"
```

```bash
echo "This Text is centered" | center
```

```bash
echo "This Text is centered \n With multiple lines" | center
```

## License

[Apache License 2.0](LICENSE)
