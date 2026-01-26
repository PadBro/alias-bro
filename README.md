# alias-bro

`alias-bro` is a simple CLI tool for managing shell command aliases.
It lets you create, list, and manage aliases for your favorite commands, and automatically generates a shell source file so your aliases are ready to use.

---

## Features

- Add new aliases: `alias-bro add [alias] [command]`
- List all aliases in a nicely formatted table
- Remove aliases
- Automatically generates a shell source file (`aliases.sh`)
- Versioned releases with semantic versioning

---

## Installation

Download the latest release from GitHub:

```bash
curl -sSL https://github.com/PadBro/alias-bro/releases/latest/download/alias-bro -o alias-bro
chmod +x alias-bro
sudo mv alias-bro /usr/local/bin/
```

Or build from source:

```bash
git clone https://github.com/PadBro/alias-bro.git
cd alias-bro
go build -ldflags "-X github.com/PadBro/alias-bro/internal.Version=$(git describe --tags --abbrev=0)" -o alias-bro
sudo mv alias-bro /usr/local/bin/
```

---

## Setup

After installing, set up your shell to automatically source alias-bro’s aliases:

```bash
alias-bro setup
```

This detects your shell and adds the required source line to your shell’s rc file (`~/.bashrc`, `~/.zshrc`, etc.).

> If you add new aliases, remember to **reload your shell** or run:
> ```bash
> source ~/.config/alias-bro/aliases.sh
> ```

---

## Usage

### Add an alias

```bash
alias-bro add ll "ls -la"
```

- Adds a new alias `ll` for the command `ls -la`
- Automatically regenerates your `aliases.sh` file

### List aliases

```bash
alias-bro list
```

- Shows all aliases in a neat table

### Remove an alias

```bash
alias-bro remove ll
```

- Removes the alias `ll` from your configuration

---

## Configuration

Aliases are stored in:

```
~/.config/alias-bro/aliases.yaml
```

The generated shell source file is located at:

```
~/.config/alias-bro/aliases.sh
```

- `add` updates both the YAML and the shell file
- `remove` removes from both

---

## Versioning

`alias-bro` uses **semantic versioning**.

You can see the current version:

```bash
alias-bro version
```

Builds are versioned and included in the binary via Go `ldflags`.

---
