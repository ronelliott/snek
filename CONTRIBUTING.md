# Contributing

Thank you for your interest in contributing to `snek`. Bug fixes are the
highest priority and receive the most timely review. For other
contributions, review may take time — please be patient.

## Bug Fixes

Open a bug report issue first so the reproduction steps and expected
behavior are documented. A pull request without a linked issue may be
asked to wait until one is filed. Regression tests that prevent the bug
from recurring are strongly preferred.

## Features

Before opening a feature pull request, start a
[Discussion](https://github.com/ronelliott/snek/discussions) to describe
the use case and proposed approach. This avoids investing time in an
implementation that doesn't align with the project direction.

## Refactoring

Refactoring PRs are welcome when they make the project measurably better.
Include a short explanation of what the refactor improves — simpler code,
better performance, clearer structure — and why the improvement is worth
the churn.

## Pull Request Hygiene

Run the local validation command before marking a pull request ready for
review. Do not rely solely on CI as a first check:

```sh
test -z "$(gofmt -l .)" && go vet ./... && go build ./... && go test ./...
```

All four steps must pass locally before the PR is ready.

## Conventional Commits

This project uses [Conventional Commits](https://www.conventionalcommits.org/).
Release automation reads the commit history to determine the next version
and generate the changelog, so the format matters:

- `feat:` — a new feature (bumps the minor version after 1.0).
- `fix:` — a bug fix (bumps the patch version).
- `feat!:` or `fix!:` with a `BREAKING CHANGE:` footer — a breaking change
  (bumps the major version after 1.0; bumps the patch version before 1.0
  due to `bump-patch-for-minor-pre-major`).
- Other types (`docs:`, `chore:`, `refactor:`, `test:`, `ci:`) do not
  trigger a release on their own.

Keep subjects in the imperative mood, lowercase after the colon, and under
72 characters.

## License

By submitting a contribution you confirm that you own or have the rights to
the code you are contributing. All contributions are provided under the
[Apache-2.0](LICENSE.txt) license that covers this project.
