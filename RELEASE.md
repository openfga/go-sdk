# Release guide

This project uses [release-please](https://github.com/googleapis/release-please) via a
`workflow_dispatch`-triggered GitHub Actions workflow. This document explains how to cut
a release and what to watch out for.

---

## Versioning rules for this project

We are pre-1.0.0. Semver conventions are relaxed:

| Change type      | Bump                  | Example             |
|---               |---                    |---                  |
| Breaking change  | **Minor** (`0.x.0`)   | `0.7.0` → `0.8.0`  |
| Everything else  | **Patch** (`0.0.x`)   | `0.7.5` → `0.7.6`  |

Major bumps (`1.0.0`) are reserved for a deliberate stable-API graduation decision — not for
routine breaking changes.

---

## Cutting a release

1. Go to **Actions → release-please** and click **Run workflow**.
2. Choose a bump type:
   - `patch` — bugfixes, docs, small changes
   - `minor` — breaking changes (see above)
   - `explicit` — you specify the exact version string (e.g. `0.8.0` or `0.8.0-beta.1`)
3. The workflow creates a release PR. Review it, then merge.
4. The GitHub Release and tag are created automatically on merge.

> **Note — release-please only understands `auto` or an explicit version string.**
> The `patch`, `minor`, and `major` options in the workflow dropdown are conveniences
> implemented in the workflow. The workflow reads the current manifest version, computes
> the next version (e.g. `0.7.5` + patch = `0.7.6`), and passes that computed string
> to release-please as an explicit `Release-As:` commit — exactly the same as choosing
> `explicit` and typing it yourself. There is no native patch/minor/major mode in
> release-please. This is why `explicit` is always the safest option when in doubt —
> you are just skipping the arithmetic step.

---

## When to use `explicit`

Use `explicit` and type the version yourself in any of these situations:

**After a beta or non-conventional tag.**
If the previous release was something like `0.7.5-beta.1`, release-please tracks the
base semver (`0.7.5`) but cannot reliably decide whether the next release should be
`0.7.5`, `0.7.6`, or `0.8.0`. It will often guess wrong.

The rule of thumb: **if the last tag had a pre-release suffix, always use `explicit` for
the next release.**

**After a manually created tag.**
Any tag created outside of the release-please workflow (e.g. hotfixes, manual git tags)
is invisible to release-please's version logic. Use `explicit` to anchor the next version
correctly.

**When you want a beta.**
Release-please does not increment pre-release suffixes automatically. Use `explicit` for
every beta, incrementing the suffix manually:
```
0.8.0-beta.1  →  explicit: 0.8.0-beta.2  →  explicit: 0.8.0
```

---

## What goes in the changelog

Commit messages must follow [Conventional Commits](https://www.conventionalcommits.org/)
for release-please to group them correctly:

```
feat: add support for batch check          → Added
fix: correct retry logic for transient errors  → Fixed
docs: update API reference                 → Documentation
perf: cache DNS lookups                    → Changed
refactor: extract auth helper              → (hidden)
chore: bump dependencies                   → (hidden)
```

---

## Troubleshooting

**"Invalid previous_tag parameter" error.**
The manifest version does not have a corresponding GitHub Release object. Reset the
manifest to the last valid tag:
```bash
echo '{ ".": "0.x.y" }' > .release-please-manifest.json
git commit -am "chore: reset manifest to v0.x.y"
git push origin main
```

**Duplicate release PRs.**
Close all stale ones. The workflow auto-closes stale open PRs on each dispatch, but
merged duplicates need manual labelling with `autorelease: tagged`.

**Changelog shows everything ungrouped.**
Make sure `changelog-type` in `release-please-config.json` is set to `"default"`, not
`"github"`.
