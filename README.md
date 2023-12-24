# Securify - Simple Password Generator

[![Go Report Card](https://goreportcard.com/badge/github.com/veyselaksin/securify)](https://goreportcard.com/report/github.com/veyselaksin/securify)
[![GoDoc](https://godoc.org/github.com/veyselaksin/securify?status.svg)](https://godoc.org/github.com/veyselaksin/securify)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/veyselaksin/securify)

## Overview

Securify is a command-line utility for creating strong and secure passwords. It's a handy tool for developers, system administrators, and anyone who needs a quick and reliable way to generate passwords.

## Features

- **Customizable Options:** Tailor your passwords with various options such as length, character sets, and more.
- **Secure and Random:** Utilizes a robust randomization algorithm to ensure the generated passwords are secure.
- **Easy to Use:** Simple command-line interface makes it easy to integrate into your workflow.
- **Cross-Platform:** Works on Windows, macOS, and Linux.

## Installation

You can download the latest version of Securify as a single binary for your operating system and architecture from below.

| OS      | Architecture                                         |                                                   |                                                  |
| ------- | ---------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------ |
| Linux   | [AMD](./dist/securify_linux_amd64_v1/securify)       | [ARM](./dist/securify_linux_arm64/securify)       | [i386](./dist/securify_linux_386/securify)       |
| macOS   | [AMD](./dist/securify_darwin_amd64_v1/securify)      | [ARM](./dist/securify_darwin_arm64/securify)      |                                                  |
| Windows | [AMD](./dist/securify_windows_amd64_v1/securify.exe) | [ARM](./dist/securify_windows_arm64/securify.exe) | [i386](./dist/securify_windows_386/securify.exe) |

### Go Installation

If you have Go installed, you can install Securify with the following command:

```bash
go install github.com/veyselaksin/securify@latest

```

## Usage

```bash

securify generate password [options]

```

## Options

The following options are available:

- **-h, --help:** Show help information.
- **-l, --letters:** Include letters in the password. (default: true)
- **-d, --digits:** Include digits in the password. (default: true)
- **-s, --special:** Include special characters in the password. (default: true)
- **-c, --capital:** Include capital letters in the password. (default: true)
- **-n, --length:** The length of the password. (default: 16)

## Examples

The following examples show how to use the command:

```bash
securify generate password // default options
securify generate password -n 32 // 32 characters
securify generate password -l -d -n 8 // only letters and digits with 8 characters
securify generate password -l -d -s -n 8 // only letters, digits, and special characters with 8 characters
securify generate password -l -d -s -c -n 8 // only letters, digits, special characters, and capital letters with 16 characters
```
