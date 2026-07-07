# Making Temporal API and Go SDK Changes

This repository picks up its `temporalio/api` dependency via a Git submodule. This works great
when you are only interested in working with the `main` branch. But additional steps are required
if you are trying to test out changes that modify the gRPC protocol buffer definitions, and wire
those through to the SDK and eventually Temporal server.

After you've made changes to the `api` repo and pushed your changes to GitHub, you will need to
update the local submodule to point to your branch. (IMPORTANT: You will need to revert that change
before you submit your `api-go` changes, after you have first merged your `api` changes into `main`!)

```zsh
# Required to allow Git to pull data via local file:// transport.
git -c protocol.file.allow=always 
git config --global protocol.file.allow always

# Name of your feature branch in the api repo.
export LOCAL_API_BRANCH="feature/my-experimental-feature"

# Set the submodule in proto/api to your local api repo.
git submodule set-url proto/api "$(pwd)/../api"
git submodule set-branch --branch ${LOCAL_API_BRANCH} proto/api
git submodule sync   # propagates the URL change from .gitmodules -> .git/config
git submodule update --init --remote proto/api
```

Now, building via `make` or `make update-proto` will pickup any changes
from your feature branch. And not `main`.
