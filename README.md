> [!IMPORTANT]
> This repository has been transferred to https://github.com/tnagatomi/gh-fuda

# gh-mrlabel

`gh-mrlabel` is a gh extension for manipulating labels across multiple repositories

## Installation

Install as a [gh](https://cli.github.com/) extension (ref. [gh manual of gh extension install](https://cli.github.com/manual/gh_extension_install)).

```console
$ gh extension install tnagatomi/gh-mrlabel
```

## Usage

Login to GitHub with `gh auth login` (ref. [gh manual of gh auth login](https://cli.github.com/manual/gh_auth_login)) so that the extension can access the repositories.

### Global Options

- `-R`, `--repos`: Select repositories using the `OWNER/REPO` format separated by comma (e.g., `owner1/repo1,owner2/repo1`)
- `--dry-run`: Check what operations would be executed without actually operating on the repositories

### List of Commands

#### Create Labels

```console
$ gh mrlabel create
```

Create specified labels to the specified repositories.

##### Options

- `-l`, `--label`: Specify the labels to create in the format of `'label1:color1:description1[,label2:color2:description2,...]'` (description can be omitted)

##### Example

```console
$ gh mrlabel create -R "owner1/repo1,owner1/repo2,owner2/repo1" -l "label1:ff0000:description for label 1,label2:00ff00,label3:0000ff"
```

#### Delete Labels

```console
$ gh mrlabel delete
```

Delete specified labels from the specified repositories.

##### Options

- `-l`, `--label`: Specify the labels to delete in the format of `'label1[,label2,...]'`
- `--force`: Do not prompt for confirmation

##### Example

```console
$ gh mrlabel delete -R "owner1/repo1,owner1/repo2,owner2/repo1" -l "label1,label2,label3"
```

#### Sync Labels

```console
$ gh mrlabel sync
```

Sync the labels in the specified repositories with the specified labels.

##### Options

- `-l`, `--label`: Specify the labels to set in the format ofSpecify the labels to set in the format of `'label1:color1:description1[,label2:color2:description2,...]'` (description can be omitted)
- `--force`: Do not prompt for confirmation

##### Example

```console
$ gh mrlabel sync -R "owner1/repo1,owner1/repo2,owner2/repo1" -l "label1:ff0000:description for label 1,label2:00ff00,label3:0000ff"
```

#### Empty Labels

```console
$ gh mrlabel empty
```

Delete all labels from the specified repositories.

##### Options

- `--force`: Do not prompt for confirmation

##### Example

```console
$ gh mrlabel empty -R "owner1/repo1,owner1/repo2,owner2/repo1"
```
