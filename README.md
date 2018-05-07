# Twenty48
[![CircleCI](https://circleci.com/gh/alyyousuf7/twenty48.svg?style=shield)](https://circleci.com/gh/alyyousuf7/twenty48) [![codecov](https://codecov.io/gh/alyyousuf7/twenty48/branch/master/graph/badge.svg)](https://codecov.io/gh/alyyousuf7/twenty48)

Play 2048 from CLI

![2048 on CLI](https://user-images.githubusercontent.com/14050128/39704165-f1c7654a-5223-11e8-87fc-66ddc9ae2977.gif)

Use up, down, left and right arrow keys to play, and `ESC` to quit.

The default grid size is `8x8` which can be changed be setting `-width` and `-height` flags.

```bash
$ twenty48 [-width=8] [-height=8]
```

## Install
Requires Golang:
```bash
$ go get -u github.com/alyyousuf7/twenty48/cmd/twenty48
```

## Development
On host machine with Golang installed:
```bash
$ make test   # execute test
$ make binary # build the executable in bin/
```

Or using Docker:
```bash
$ make shell  # jump into a container
root@container:/go/src/github.com/alyyousuf7/twenty48# make test
root@container:/go/src/github.com/alyyousuf7/twenty48# make binary
root@container:/go/src/github.com/alyyousuf7/twenty48# exit
```

The Docker container mounts a volume to `bin/` to copy executable to host machine.

## TODOs
- [ ] Write more tests
- [ ] Maintain scoreboard
- [ ] Display game over text

PRs are most welcome!

## License
[MIT](./LICENSE)