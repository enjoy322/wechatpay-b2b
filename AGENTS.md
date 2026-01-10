# Repository Guidelines

## Project Structure & Module Organization
- `go.mod`: Module definition (`github.com/enjoy322/wechatpay-b2b`, Go 1.21).
- `config/`: Shared configuration (app ids, app keys, env).
- `client/`: HTTP client wrapper and token-provider hook.
- `signer/`: HMAC utilities for paySig and user signature.
- `model/`: Core DTOs (amounts, orders, statuses, balance/withdraw types).
- `service/`: Service facades (payment param builder, order, refund, balance/withdraw, notify).
- Tests folder not present yet; place new tests alongside packages (e.g., `service/payment_test.go`).

## Language & Localization
- 项目文档、注释、提交信息优先使用中文，除非对外 API 名称需保持英文一致性。
- 面向微信支付 B2B 场景，保持术语与微信官方文档一致（如 “商户号”“支付签名”“退款” 等）。

## Build, Test, and Development Commands
- `go fmt ./...` — format code (use before committing).
- `go test ./...` — run unit tests across all packages.
- `go vet ./...` — static checks for common issues.
- `golangci-lint run` — if available locally, run aggregated linting (not vendored here).

## Coding Style & Naming Conventions
- Follow standard Go style; always run `gofmt`.
- Use Go module path imports: `github.com/enjoy322/wechatpay-b2b/...`.
- Keep exported types/functions commented when part of the public API; keep comments succinct.
- Prefer clear, value-oriented naming (e.g., `BuildCommonPaymentParams` over abbreviations).
- Errors: return wrapped errors with context; avoid panics in library code.

## Testing Guidelines
- Add table-driven tests per package (`*_test.go`), colocated with the code.
- Cover edge cases for signing (empty keys, invalid uri/body), request validation, and HTTP error handling once implemented.
- Use `t.Helper()` for shared assertions; avoid global state to keep tests parallel-safe.

## Commit & Pull Request Guidelines
- Commits: use imperative, concise subjects (e.g., “Add signer helpers”, “Implement refund client”); group related changes per commit.
- PRs: include a short summary, the main changes, and test evidence (`go test ./...`). Link related issues if applicable.
- Screenshots not required unless adding user-facing artifacts (none currently).

## Security & Configuration Tips
- Do not commit app secrets/app keys; load them via environment or external config.
- When adding token fetching, guard logging to avoid leaking tokens or signatures.
- Keep HTTP clients time-bounded (timeouts) and avoid disabling TLS verification.
