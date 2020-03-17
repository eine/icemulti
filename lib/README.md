# [icemulti] Library

## `.icem` file formats

### Uncompressed (`tar`)

- `<filename>.icem`
  - `.`
    - `bin`
    - `json`

Create:

```
ls -la
  bin
  json
tar -cvf <filename>.icem ./*
```

Extract:

```
mkdir img02_blink
cd img02_blink
tar -xvf img02_blink.icem
```

### Compressed (`.tgz`, `.tar.gz`)

- `<filename>.icem`
  - `<filename>.tar`
    - `.`
      - `bin`
      - `json`

Create:

```
ls -la
  bin
  json
tar -cvzf <filename>.icem ./*
```

Extract:

```
mkdir img02_blink
cd img02_blink
tar -xvzf img02_blink.icem
```
