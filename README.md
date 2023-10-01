# Show Hardware (ShowHW)

**ShowHW** is a command line (CLI) tool written in [Go](https://go.dev), which shows the main hardware component specs of the local machine.

## Usage

```
showhw
```

```
 __ _                          __    __
/ _\ |__   _____      __/\  /\/ / /\ \ \
\ \| '_ \ / _ \ \ /\ / / /_/ /\ \/  \/ /
_\ \ | | | (_) \ V  V / __  /  \  /\  /
\__/_| |_|\___/ \_/\_/\/ /_/    \/  \/

version 0.1.0

Product   NUC11PAHi7 Intel(R) Client Systems
Memory    64GB
CPU1      11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz (4 cores, 8 threads)
Disc1     2TB SSD Samsung SSD 980 PRO 2TB
GPU1      TigerLake-LP GT2 [Iris Xe Graphics] Intel Corporation
```

## Notes

`showhw` is not recommnded on Mac OSX as some of the [ghw](https://github.com/jaypipes/ghw) library functions are not implemented and make the tool mostly moot.


## License

[MIT License](/LICENSE)
