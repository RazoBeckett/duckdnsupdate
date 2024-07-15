# Duck Dns Update

this will update your duck dns domain with your current ip address or a custom ip address

## Build

```bash
go build .
```

## Usage

> [!IMPORTANT]
> domain should be the domain name without the .duckdns.org part and without https:// prefix

- This will update your domain with your current public ip address

```bash
ddnsupdate --domain <your-domain> --token <your-token>
```

- This will update your domain with the provide ip address

```bash
ddnsupdate --domain <your-domain> --token <your-token> --ip <your-ip>
```
