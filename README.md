# LibreCLI

LibreCLI is a small CLI tool that allows you to interact with LibreNMS.

You need to specify the following environment variables to be able to use it:

```
LIBRECLI_URL = Root URL of your LibreNMS setup
LIBRECLI_TOKEN = Your token
```

It's still in very early development, for now the following commands are implemented

```bash

$ librecli bgp peers list <Device name or remote ASN>

+-----------------+-------------+-------------+---------+--------------+-------------+-------------+
| Device          | Local IP    | Peer IP     | Peer AS | AS Holder    | State       | Admin State |
+-----------------+-------------+-------------+---------+--------------+-------------+-------------+
| my-router-1     | 0.0.0.0     | 191.9.41.5  | 65000   | Private ASN  | established | running     | 
| my-router-1     | 0.0.0.0     | 191.9.41.6  | 65000   | Private ASN  | established | running     |
+-----------------+-------------+-------------+---------+--------------+-------------+-------------+ 

$ librecli bgp peers counter

+-----------------------+-------------------+--------------+------------+-----------+---------+                                                                             
| Device                | Peer IP           | Accepted Pfx | Denied Pfx | Pfx Thrsd | Adv Pfx |                                                                             
+-----------------------+-------------------+--------------+------------+-----------+---------+                                                                             
| my-router-1           | 191.9.41.5        | 0            | 0          | 0         | 0       |                                                                             
| my-router-1           | 191.9.41.6        | 109          | 109        | 0         | 1       |
+-----------------------+-------------------+--------------+------------+-----------+---------+ 

$ librecli fdb lookup [MAC address]

Should display all the FDB entries containing the MAC address, I need to update my LibreNMS to test it :)
```

Much more commands to come, PR are welcome of course.

