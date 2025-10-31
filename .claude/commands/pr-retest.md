---
description: Find and retest failed e2e CI jobs on a PR
argument-hint: <pr-number>
---

## Name
pr-retest

## Synopsis
```
/pr-retest <pr-number>
```

## Description
Analyzes a pull request in the openshift/ovn-kubernetes repository to find all failed e2e CI jobs and provides options
to retest them. The command identifies jobs matching the pattern `ci/prow/.*e2e.*` that are in a FAILURE or ERROR state,
then presents interactive options for retesting.

## Implementation

IMPORTANT: Execute this command immediately without asking for permission or confirmation first.

### Step 1: Fetch failed e2e jobs

Immediately use the GitHub API to fetch the PR's status checks from openshift/ovn-kubernetes and filter for:
- Jobs with state == "FAILURE" or "ERROR"
- Jobs matching the pattern `ci/prow/.*e2e.*`

Transform job names by stripping the `ci/prow/` prefix to get the test name for retest commands.

Example:
- Job context: `ci/prow/e2e-aws-ovn`
- Test name: `e2e-aws-ovn`
- Retest command: `/test e2e-aws-ovn`

### Step 2: Present options to user

If failed e2e jobs are found, display the list and use the AskUserQuestion tool to present these options:

1. **Retest selected** - Provide a space-separated list of job names to retest
2. **Use /retest** - Post a single `/retest` comment to rerun all failed tests
3. **Just show list** - Display the failed jobs without taking action

If no failed e2e jobs are found, inform the user and exit.

### Step 3: Execute user's choice

Based on the selected option:

**Option 1 - Retest selected:**
- Ask user to provide a space-separated list of job names (e.g., "e2e-aws-ovn e2e-gcp-ovn")
- Parse the input and split by whitespace to get individual job names
- Create a comment body with one `/test <job>` line per selected job
- Post the single comment to the PR using `gh pr comment`

Example: If user provides "e2e-aws-ovn e2e-gcp-ovn", post comment:
```
/test e2e-aws-ovn
/test e2e-gcp-ovn
```

**Option 2 - Use /retest:**
- Post a comment containing just `/retest`

**Option 3 - Just show list:**
- Display the list of failed e2e jobs to the user
- No further action

## Example Commands

Fetch failed jobs:
```bash
gh pr view <pr-number> --repo openshift/ovn-kubernetes --json statusCheckRollup | \
  jq -r '.statusCheckRollup[] |
    select(.state == "FAILURE" or .state == "ERROR") |
    select(.context | test("ci/prow/.*e2e")) |
    .context | sub("ci/prow/"; "")'
```

Post retest comment:
```bash
gh pr comment <pr-number> --repo openshift/ovn-kubernetes --body "/test <job-name>"
```

## Arguments
- $1: Pull request number (required)
  - Example: 2838
