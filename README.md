# batch tool

[![Build Status](https://travis-ci.org/hpeng526/batch-tool.svg?branch=master)](https://travis-ci.org/hpeng526/batch-tool)

### Why?

1. I want to exec some command in sub dirs
2. And ignore specified directory

### How?

```bash
go get github.com/hpeng526/batch-tool
```

#### setup your config

```json
{
  "ignore_paths": [".git"],
  "reg_exp": "(^\\.git$|^\\.DS_Store)"
}
```

```bash
batch-tool git pull
```

* `ignore_paths` is fully-matching
* `reg_exp` is using regular expression
* if both setting up, `reg_exp` will take place of `ignore_paths`
* enjoy