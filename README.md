<a href="https://github.com/PryosCode/center/blob/master/assets/neofetch.txt"><img align="center" width="100%" alt="neofetch" src="https://github.com/PryosCode/center/raw/master/assets/neofetch.png"></a>

<p align="center">
    <a href="https://github.com/PryosCode/center/releases"><img alt="Version" src="https://img.shields.io/github/v/release/PryosCode/center?label=Version"></a>
    <a href="https://github.com/PryosCode/center/actions"><img alt="Linux" src="https://github.com/PryosCode/center/actions/workflows/linux.yml/badge.svg"></a>
    <a href="https://github.com/PryosCode/center/actions"><img alt="Windows" src="https://github.com/PryosCode/center/actions/workflows/windows.yml/badge.svg"></a>
    <a href="https://github.com/PryosCode/center/actions"><img alt="MacOS" src="https://github.com/PryosCode/center/actions/workflows/macos.yml/badge.svg"></a>
    <a href="https://github.com/PryosCode/center/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/PryosCode/center?label=License"></a>
</p>

# center

A cli tool to center text

## Download

[Releases](https://github.com/PryosCode/center/releases)

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