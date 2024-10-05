# Changelog
All notable changes to this project will be documented in this file.

The format is based on http://keepachangelog.com/en/1.0.0/
and this project adheres to http://semver.org/spec/v2.0.0.html.

## [0.4.0-dev]

- Add option -S, --consecutive-space, default 4
- Add field Trim.ConsecutiveSpace

## [0.3.0] 2022-12-21

- Add option --path-len to cmd/trim
- Add Trimmer.PathLen
- Add func TrimPaths
- Add type Trimmer
- Rename package to gregoryv/trim
- Add cmd/trim

## [0.2.0] 2022-11-30

- Add option --tab-width
- Reset terminal coloring if trimmed

## [0.1.0] 2022-11-30

- Add option --suffix, defaults to "..."
- Add option --columns, defaults to 72
