# GitHub Workflows

## Create PR for LTS Branch

The repo that owns API definitions is responsible for deciding when to update API definitions (ie after a release) and pushing the new API to a branch in solo-apis.

This action only needs to know 2 pieces of information:
1. The name of the LTS branch it should merge the code into
1. The tag that should be applied to the merge commit

We decided to use the branch name to pass this information. The workflow only operates on branches in the following format:
```sync-apis/{LTS_BRANCH_NAME}/{LTS_TAG_NAME}```

This allows the workflow to operate successfully, knowing the least information possible about the source repo that triggered the sync.

The workflow has 2 responsibilities:
1. To open a pull request to the LTS branch name
1. To signify the tag that should be applied, when the PR merges

The first step is easy, since the LTS branch is extracted from the branch name.
The second step is more challenging since after the PR merges, the GitHub event that is triggered does not contain a reference to the name of the branch that merged. Therefore, we don't have a way to determine the name of the tag to apply. **We use a little trick to resolve this: providing a PR title with the tag included. When the PR is merged, that title becomes the commit message, and we parse that commit message when generating a tag name.**


*Note: This workflow opens a PR, but expects manual intervention by an engineer to approve the PR*

## Tag Commit on LTS Branch

After PRs are merged into LTS branches, we:
1. Extract the latest commit message (which the above workflow provided)
1. Determine the tag name from the commit message
1. Push a tag name for that commit
