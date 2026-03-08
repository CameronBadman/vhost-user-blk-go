# vhost-user-blk-go

A learning project implementing the [vhost-user protocol](https://qemu.readthedocs.io/en/latest/interop/vhost-user.html) for a block device backend in Go.

The server listens on a Unix socket (`/tmp/vhost-blk.sock`) and handles vhost-user negotiation from a QEMU guest.

## Requirements

- Go 1.25+
- [QEMU](https://www.qemu.org/) — required to connect a guest VM to the vhost-user backend

## Running

```sh
go run .
```

Then launch QEMU with a vhost-user-blk device pointing at `/tmp/vhost-blk.sock`.

## Notes

This is a work in progress. If you have feedback it's appreciated.
