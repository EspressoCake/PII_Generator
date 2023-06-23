# Fake PII Generator

## Why
- `Python` loops are not incredibly performant.
- Due to `Golang`'s static compilation, the can *easily* be turned into shellcode
- We can *easily* turn this into a supported `Jenkins` build pipeline if we just want the artifacts themselves

## How
We take liberties with the strengths of `Golang`, namely its concurrency model. This splits the generation of "records" into "chunks" 40K at a time.

We use the synchronization primitive, a `WaitGroup`, to ensure these items finish before continuing. This is a convenient alternative to traditional `channel`s. However, we are smart about it, in using the aformentioned chunks to dictate the amount to generate per concurrent use, and not simply `async`ing every use. This leads to lower overhead and generally faster performance.

## Anticipated Performance
- This program reliably creates and writes 400K "records" in under four seconds

## Why not use the CSV module within the language?
- If we understand the data we are creating and using, this becomes redundant if they are pre-defined

## Usage
Modify the total number of records as you see fit in `main.go` (default: 400000)

```sh
make all
./bins/PII_Generator_[darwin|linux|windows.exe]
```
