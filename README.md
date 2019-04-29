# LibreCLI

LibreCLI is a small CLI tool that allows you to interact with LibreNMS.

You need to specify the following environment variables to be able to use it:

```
LIBRECLI_URL = Root URL of your LibreNMS setup
LIBRECLI_TOKEN = Your token
```

## Usage

```bash

# Global Flags
  -f, --format string   Output format: table|list|json (default "table")
  -V, --verbose         Enable verbose mode


$ librecli bgp peers list <Device name or remote ASN>

$ librecli bgp peers counters <Device Name>

$ librecli fdb lookup [MAC address]

$ librecli ipv4 lookup [IPv4 Address]

```

Much more commands to come, PR are welcome of course.

