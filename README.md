![demo.gif](./demo/demo.gif)

# mdfmt

A simple markdown formatter.

Currently formatting of the following markdown elements is supported:

- headings
- paragraphs
- lists (with hyphens)
- tables

# Installation

```shell
  $ go install github.com/tuxikus/mdfmt@latest

  # or clone locally
  $ git clone https://github.com/tuxikus/mdfmt.git
  $ cd mdfmt
  $ go install
```

# Usage

## CLI

```shell
  $ cat README.md | mdfmt > README.md.tmp | mv README.md.tmp README.md
```

## Helix

Select text with `%`, pipe with `|` and call mdfmt.
