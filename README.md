# Duck Dns Update

A simple command line tool to update your duckdns domain with your current public ip address
Built with Go using the [Cobra](https://github.com/spf13/cobra) library

## Build

```bash
go build .
```

## Usage

> [!IMPORTANT]
> domain should be the subdomain name without the .duckdns.org part and without https:// prefix

- This will update your domain with your current public ip address

```bash
duckdnsupdate <subdomain> <token>
```

- This will update your domain with the provide ip address

```bash
duckdnsupdate <subdomain> <token> --ip-addr <ip-address>
```

[LICENSE](LICENSE)
