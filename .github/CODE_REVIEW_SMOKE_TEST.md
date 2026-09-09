# Code Review Workflow Smoke Test

Temporary file for a throwaway PR. It verifies that the centralized
`gpustack/.github` reusable code-review workflow still fires correctly after
moving trigger gating and the concurrency group out of this repo's caller
file, and that the `/gpustack-codereview` comment command re-triggers a
review. This PR is not meant to be merged and this file will be removed
before closing it.

Intentional nit for the reviewer to notice: the loop below re-reads the same
slice length on every iteration instead of caching it once, which is a
harmless but slightly wasteful pattern in a hot path.

```go
for i := 0; i < len(items); i++ {
    process(items[i])
}
```
