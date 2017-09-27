# traefik-fl - Traefik frontend list

This simple utility proxies request to [traefik](https://traefik.io/) container,
filters response and returns hash with running backends and its hosts.

The aim of this project is to cut traefik api private data (BasicAuth etc) from running host list
for use in [dcape CIS](https://github.com/dopos/dcape/tree/master/apps/cis).

Result parsing example: see javascript [index.html](https://github.com/dopos/dcape/blob/master/apps/cis/html/index.html) there.

## Usage
```
$ curl -s http://cis.dev.lan/frontends.json | jq '.'
{
  "cis": [
    "cis.dev.lan"
  ],
  "gitea": [
    "git.dev.lan"
  ],
  "pdns": [
    "ns.dev.lan"
  ],
  "portainer": [
    "port.dev.lan"
  ]
}

```

## License

The MIT License (MIT), see [LICENSE](LICENSE).

Copyright (c) 2017 Alexey Kovrizhkin <lekovr+dopos@gmail.com>
