PYTHON_WASM_REPO     ?= goccy/python-wasm
PYTHON_WASM_VERSION  ?= v0.1.5
# python-wasm emits its release attestations from release.yml (the v* tag
# workflow), NOT build.yml — releasing lives only in release.yml there.
PYTHON_WASM_WORKFLOW ?= goccy/python-wasm/.github/workflows/release.yml

TARBALL          := python_wasm2go.tar.gz
SHA256SUMS       := python_wasm2go.sha256
RELEASE_URL       = https://github.com/$(PYTHON_WASM_REPO)/releases/download/$(PYTHON_WASM_VERSION)
ATTESTATION_API   = https://api.github.com/repos/$(PYTHON_WASM_REPO)/attestations

.PHONY: bundle download verify verify-release verify-attestation build

## bundle: download release artifacts, verify their attestations, and
## leave the extracted module tree in place. Default to running this
## whenever PYTHON_WASM_VERSION bumps.
bundle: download verify

## download: fetch the wasm2go tarball + sha256 manifest and extract
## them into the repo root. The tarball is discarded once unpacked.
download:
	curl -fSL --proto '=https' --tlsv1.2 -o $(TARBALL)    $(RELEASE_URL)/$(TARBALL)
	curl -fSL --proto '=https' --tlsv1.2 -o $(SHA256SUMS) $(RELEASE_URL)/$(SHA256SUMS)
	tar xzf $(TARBALL)
	rm -f $(TARBALL)

## verify: byte-check each in-tree file against the sha256 manifest AND
## confirm each carries a valid GitHub artifact attestation signed by
## the upstream release.yml workflow. Either check failing aborts.
verify: verify-release verify-attestation

## verify-release: confirm every in-tree file matches the entries in
## $(SHA256SUMS). Fast sanity check; not a trust anchor on its own.
verify-release:
	@echo "==> verifying in-tree files against $(SHA256SUMS)"
	@shasum -a 256 -c $(SHA256SUMS)

## verify-attestation: confirm every in-tree artifact is a signed
## subject of the upstream SLSA build attestation. The release emits one
## attestation whose subject list covers every file in the tarball, so
## we fetch the bundle once from the public attestation API and then
## offline-verify each file via `gh attestation verify --bundle`.
## The fetch works anonymously, but anonymous api.github.com is rate
## limited to 60 req/hr per IP — enough to 403 on shared CI runners — so
## when GH_API_TOKEN is set (CI passes ${{ github.token }}) the request
## is authenticated, raising the limit to 5000/hr. The offline
## `gh attestation verify` below stays tokenless regardless.
verify-attestation:
	@set -eu; \
	tmpdir=$$(mktemp -d); \
	bundle=$$tmpdir/bundle.jsonl; \
	trap 'rm -rf $$tmpdir' EXIT; \
	probe=$$(awk 'NR==1 {print $$2}' $(SHA256SUMS) | sed 's|^\./||'); \
	digest=$$(shasum -a 256 $$probe | awk '{print $$1}'); \
	echo "==> fetching attestation bundle via $$probe (sha256:$$digest)"; \
	if [ -n "$${GH_API_TOKEN:-}" ]; then set -- -H "Authorization: Bearer $$GH_API_TOKEN"; else set --; fi; \
	curl -fsSL --proto '=https' --tlsv1.2 "$$@" \
	  "$(ATTESTATION_API)/sha256:$$digest" \
	  | jq -c '.attestations[].bundle' > $$bundle; \
	files=$$(awk '{print $$2}' $(SHA256SUMS) | sed 's|^\./||'); \
	for f in $$files; do \
	  echo "==> verifying $$f"; \
	  GH_TOKEN= GITHUB_TOKEN= gh attestation verify "$$f" \
	    -R $(PYTHON_WASM_REPO) \
	    --bundle $$bundle \
	    --signer-workflow $(PYTHON_WASM_WORKFLOW); \
	done

## build: sanity-build every package. The bundle is pure generated
## code, so a successful `go build` confirms the tarball extracted
## into a working module.
build:
	go build ./...
