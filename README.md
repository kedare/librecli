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


// Supported
$ librecli bgp peers list <hostname or remote ASN>
$ librecli bgp peers counters <hostname>
$ librecli fdb lookup [MAC address]
$ librecli ip lookup [IPv4/IPv6 Address]

// TODO
$ librecli ipsec sa list <hostname>
$ librecli ipv6 lookup [IPv6 Address]
$ librecli logs events tail [hostname] [counts]
$ librecli logs syslogs tail [hostname] [counts]
$ librecli sensors list <hostname> <Filter> # Need API changes on LibreNMS, no way to filter by device on current version
$ librecli links list [hostname]
$ librecli alerts list <filters>
$ librecli alerts ack [Alert ID...]
$ librecli alerts mute [Alert ID...]
$ librecli alerts unmute [Alert ID...]
$ librecli alerts rules list
$ librecli inventory list [hostname]
```

Much more commands to come, PR are welcome of course.

