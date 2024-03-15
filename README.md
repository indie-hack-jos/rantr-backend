# RECON

---

## Contribution

---

1. Cloning the project
   - Always clone from the staging branch `git clone -b staging https://github.com/Recon-Africa/recon-api`
   - Run npm install `npm install`
   - Always create your branch from the `staging` branch i.e
     - `git checkout staging`
     - `git checkout -b branch-name`

---

2. PRs detail examples before committing

   #### Issue

   - QA engineer can't test phone verification because twilo not sending sms on staging

   #### What has changed:

   1. Changed line 97 on utils.js file
   2. restricted sending of sms to just staging and production

   #### What reviewers should know:

   - used the ternary expression

Example of a standard PR

```
### Issues
* I worked on a particular problem affecting our server

### What has changed?
* I updated this particular line or this file index.ts

### What reviewer should know?
* nothing for now, i ran migrations e.t.c
```

---

3. Creating Git branch names for adding features, bug fixes, & code cleanups respectively:

- feat/branch-name or feature/branch-name
- fix/branch-name or bug-fix/branch-name
- chore/branch-name

---

4. Tagging PR changes

- bug fixes: `fix:summary-of-fix` or `bug-fix:summary-of-fix`
- Features added: `feat:commit-message` or `feature:commit-message`
- Code cleanup `chore:commit-message`

---

5. Important thing to note while collaborating

- Always pull from develop. You can checkout to staging branch and pull or pull directly from develop into your current branch.
- Create new branch for every feature or fixes you are being assigned. Only create a fix branch if the previous tasks has been closed.
- Do not push directly to `develop` or `main`. Always create a PR and use our PR standard.

The above is simplified in this table:

| Task            | Branch naming          | PR changes tagging          |
| --------------- | ---------------------- | --------------------------- |
| Adding features | feat/branch-name       | `feat:commit-message`       |
| Bug fixes       | fix/branch-name        | `fix:summary-of-fix`        |
| Code cleanup    | chore/branch-name      | `chore:commit-message`      |
| Code deployment | deployment/branch-name | `deployment:commit-message` |

---

## Run Server

---

1. Load .env
   - get the development environmental variable from .env.sample and put into your .env

---

2. Run server
   - `go run main.go`

---
