# Go by Example — exercises

Practice exercises for every section of [Go by Example](https://gobyexample.com),
one folder per section, in the same order as the site. Each folder has a few
small functions for you to write and tests that tell you whether you got them
right.

## How to use it

1. Read a section on the site, e.g. <https://gobyexample.com/slices>.
2. Open the matching folder (`09-slices`) and edit **`exercise.go`**. Every
   function has a comment describing what it should do and a `// TODO` where
   your code goes.
3. Run that section's tests from this folder:

   ```
   go test ./09-slices/
   ```

4. Read the failures, fix, repeat until you see `ok`.

Each folder contains:

| File               | What it is                                                        |
| ------------------ | ----------------------------------------------------------------- |
| `exercise.go`      | The exercises. This is the file you edit.                         |
| `exercise_test.go` | The tests. Read them — they show exactly what is expected.        |
| `solution.go`      | A reference solution. Try not to peek before you have a passing version of your own. |

Section 72 (Testing and Benchmarking) is the other way around: the code is
given and you write the tests, in `intmin_test.go`.

## Useful commands

```
go test ./09-slices/                       # one section
go test -v ./09-slices/                    # ...listing every test
go test -run TestParallelSquares ./29-channels/   # a single test
go test ./...                              # everything: shows your overall progress
go test -tags solution ./...               # run the tests against the reference solutions
```

In PowerShell, a compact progress overview:

```powershell
go test ./... 2>&1 | Select-String '^(ok|FAIL)\s'
```

Tips:

- Go caches test results. `(cached)` in the output is fine; add `-count=1` to
  force a re-run.
- The first line of each `exercise.go`, `//go:build !solution`, is a build
  constraint. It is what lets `exercise.go` and `solution.go` define the same
  functions in one folder: `-tags solution` switches from one file to the
  other. Leave that line in place.
- The stubs return placeholder values so that everything compiles; you will
  usually need to add `import`s as you go. Go refuses to compile unused imports
  and unused variables, so remove what you don't use.
- Stuck? Run a single test with `-v`, add `fmt.Println` calls, and re-read the
  page — every exercise uses what its page (or an earlier one) shows.

## Sections

| #  | Section | Folder |
| -- | ------- | ------ |
| 1  | [Hello World](https://gobyexample.com/hello-world) | `01-hello-world` |
| 2  | [Values](https://gobyexample.com/values) | `02-values` |
| 3  | [Variables](https://gobyexample.com/variables) | `03-variables` |
| 4  | [Constants](https://gobyexample.com/constants) | `04-constants` |
| 5  | [For](https://gobyexample.com/for) | `05-for` |
| 6  | [If/Else](https://gobyexample.com/if-else) | `06-if-else` |
| 7  | [Switch](https://gobyexample.com/switch) | `07-switch` |
| 8  | [Arrays](https://gobyexample.com/arrays) | `08-arrays` |
| 9  | [Slices](https://gobyexample.com/slices) | `09-slices` |
| 10 | [Maps](https://gobyexample.com/maps) | `10-maps` |
| 11 | [Functions](https://gobyexample.com/functions) | `11-functions` |
| 12 | [Multiple Return Values](https://gobyexample.com/multiple-return-values) | `12-multiple-return-values` |
| 13 | [Variadic Functions](https://gobyexample.com/variadic-functions) | `13-variadic-functions` |
| 14 | [Closures](https://gobyexample.com/closures) | `14-closures` |
| 15 | [Recursion](https://gobyexample.com/recursion) | `15-recursion` |
| 16 | [Range over Built-in Types](https://gobyexample.com/range-over-built-in-types) | `16-range-over-built-in-types` |
| 17 | [Pointers](https://gobyexample.com/pointers) | `17-pointers` |
| 18 | [Strings and Runes](https://gobyexample.com/strings-and-runes) | `18-strings-and-runes` |
| 19 | [Structs](https://gobyexample.com/structs) | `19-structs` |
| 20 | [Methods](https://gobyexample.com/methods) | `20-methods` |
| 21 | [Interfaces](https://gobyexample.com/interfaces) | `21-interfaces` |
| 22 | [Enums](https://gobyexample.com/enums) | `22-enums` |
| 23 | [Struct Embedding](https://gobyexample.com/struct-embedding) | `23-struct-embedding` |
| 24 | [Generics](https://gobyexample.com/generics) | `24-generics` |
| 25 | [Range over Iterators](https://gobyexample.com/range-over-iterators) | `25-range-over-iterators` |
| 26 | [Errors](https://gobyexample.com/errors) | `26-errors` |
| 27 | [Custom Errors](https://gobyexample.com/custom-errors) | `27-custom-errors` |
| 28 | [Goroutines](https://gobyexample.com/goroutines) | `28-goroutines` |
| 29 | [Channels](https://gobyexample.com/channels) | `29-channels` |
| 30 | [Channel Buffering](https://gobyexample.com/channel-buffering) | `30-channel-buffering` |
| 31 | [Channel Synchronization](https://gobyexample.com/channel-synchronization) | `31-channel-synchronization` |
| 32 | [Channel Directions](https://gobyexample.com/channel-directions) | `32-channel-directions` |
| 33 | [Select](https://gobyexample.com/select) | `33-select` |
| 34 | [Timeouts](https://gobyexample.com/timeouts) | `34-timeouts` |
| 35 | [Non-Blocking Channel Operations](https://gobyexample.com/non-blocking-channel-operations) | `35-non-blocking-channel-operations` |
| 36 | [Closing Channels](https://gobyexample.com/closing-channels) | `36-closing-channels` |
| 37 | [Range over Channels](https://gobyexample.com/range-over-channels) | `37-range-over-channels` |
| 38 | [Timers](https://gobyexample.com/timers) | `38-timers` |
| 39 | [Tickers](https://gobyexample.com/tickers) | `39-tickers` |
| 40 | [Worker Pools](https://gobyexample.com/worker-pools) | `40-worker-pools` |
| 41 | [WaitGroups](https://gobyexample.com/waitgroups) | `41-waitgroups` |
| 42 | [Rate Limiting](https://gobyexample.com/rate-limiting) | `42-rate-limiting` |
| 43 | [Atomic Counters](https://gobyexample.com/atomic-counters) | `43-atomic-counters` |
| 44 | [Mutexes](https://gobyexample.com/mutexes) | `44-mutexes` |
| 45 | [Stateful Goroutines](https://gobyexample.com/stateful-goroutines) | `45-stateful-goroutines` |
| 46 | [Sorting](https://gobyexample.com/sorting) | `46-sorting` |
| 47 | [Sorting by Functions](https://gobyexample.com/sorting-by-functions) | `47-sorting-by-functions` |
| 48 | [Panic](https://gobyexample.com/panic) | `48-panic` |
| 49 | [Defer](https://gobyexample.com/defer) | `49-defer` |
| 50 | [Recover](https://gobyexample.com/recover) | `50-recover` |
| 51 | [String Functions](https://gobyexample.com/string-functions) | `51-string-functions` |
| 52 | [String Formatting](https://gobyexample.com/string-formatting) | `52-string-formatting` |
| 53 | [Text Templates](https://gobyexample.com/text-templates) | `53-text-templates` |
| 54 | [Regular Expressions](https://gobyexample.com/regular-expressions) | `54-regular-expressions` |
| 55 | [JSON](https://gobyexample.com/json) | `55-json` |
| 56 | [XML](https://gobyexample.com/xml) | `56-xml` |
| 57 | [Time](https://gobyexample.com/time) | `57-time` |
| 58 | [Epoch](https://gobyexample.com/epoch) | `58-epoch` |
| 59 | [Time Formatting / Parsing](https://gobyexample.com/time-formatting-parsing) | `59-time-formatting-parsing` |
| 60 | [Random Numbers](https://gobyexample.com/random-numbers) | `60-random-numbers` |
| 61 | [Number Parsing](https://gobyexample.com/number-parsing) | `61-number-parsing` |
| 62 | [URL Parsing](https://gobyexample.com/url-parsing) | `62-url-parsing` |
| 63 | [SHA256 Hashes](https://gobyexample.com/sha256-hashes) | `63-sha256-hashes` |
| 64 | [Base64 Encoding](https://gobyexample.com/base64-encoding) | `64-base64-encoding` |
| 65 | [Reading Files](https://gobyexample.com/reading-files) | `65-reading-files` |
| 66 | [Writing Files](https://gobyexample.com/writing-files) | `66-writing-files` |
| 67 | [Line Filters](https://gobyexample.com/line-filters) | `67-line-filters` |
| 68 | [File Paths](https://gobyexample.com/file-paths) | `68-file-paths` |
| 69 | [Directories](https://gobyexample.com/directories) | `69-directories` |
| 70 | [Temporary Files and Directories](https://gobyexample.com/temporary-files-and-directories) | `70-temporary-files-and-directories` |
| 71 | [Embed Directive](https://gobyexample.com/embed-directive) | `71-embed-directive` |
| 72 | [Testing and Benchmarking](https://gobyexample.com/testing-and-benchmarking) | `72-testing-and-benchmarking` |
| 73 | [Command-Line Arguments](https://gobyexample.com/command-line-arguments) | `73-command-line-arguments` |
| 74 | [Command-Line Flags](https://gobyexample.com/command-line-flags) | `74-command-line-flags` |
| 75 | [Command-Line Subcommands](https://gobyexample.com/command-line-subcommands) | `75-command-line-subcommands` |
| 76 | [Environment Variables](https://gobyexample.com/environment-variables) | `76-environment-variables` |
| 77 | [Logging](https://gobyexample.com/logging) | `77-logging` |
| 78 | [HTTP Client](https://gobyexample.com/http-client) | `78-http-client` |
| 79 | [HTTP Server](https://gobyexample.com/http-server) | `79-http-server` |
| 80 | [TCP Server](https://gobyexample.com/tcp-server) | `80-tcp-server` |
| 81 | [Context](https://gobyexample.com/context) | `81-context` |
| 82 | [Spawning Processes](https://gobyexample.com/spawning-processes) | `82-spawning-processes` |
| 83 | [Exec'ing Processes](https://gobyexample.com/execing-processes) | `83-execing-processes` |
| 84 | [Signals](https://gobyexample.com/signals) | `84-signals` |
| 85 | [Exit](https://gobyexample.com/exit) | `85-exit` |
