# mockery-proxy

A compatibility wrapper that enables legacy `mockery` v2-style CLI usage on projects that depend on the newer `mockery` v3.

## Purpose

`mockery` v3 removed support for direct CLI flags like `--name`, `--inpackage`, etc., and requires a YAML config file instead. This proxy allows you to continue using the older v2-style CLI syntax without needing to rewrite existing scripts or Makefiles.

`mockery-proxy` internally invokes the actual v2 command implementation from the [`github.com/vektra/mockery/v2/cmd"`](https://github.com/vektra/mockery/tree/v2/cmd) package.

## Installation

```bash
go install github.com/spirin/mockery-proxy@latest
```

## Usage

### v2-style CLI (internally handled via v2 logic)

```bash
mockery-proxy --name MyInterface --inpackage --case underscore
```

Behaves exactly as if you were using `mockery v2`.

### v3-style passthrough

If the `--config` flag is passed, or no arguments are provided, the proxy forwards the call directly to the `mockery` binary installed in your system:

```bash
mockery-proxy --config .mockery.yml
```

```bash
mockery-proxy
```

## Requirements

- `mockery` v3 must be installed and available in `PATH` for passthrough behavior.

## License

MIT
