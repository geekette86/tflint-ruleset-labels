# TFLint Ruleset Labels ( based on https://github.com/terraform-linters/tflint-ruleset-template)

This is a template repository for building a custom ruleset to have lables as mandatory attribute. 
## Requirements

- TFLint v0.40+
- Go v1.19
## In case of new release

To create auto release with `goreleaser` and push it to github , just type:
```sh
git tag v0.0.X -m dzrlblrule
goreleaser release --debug  --rm-dist 
```
You need to set :  
* GPG_FINGERPRINT
* GITHUB_TOKEN  
to get you public key from you GPG FINGERPRINT
```
gpg --output PUBLIC.gpg --export --armor YOUR_ID
```
## Installation

TODO: This template repository does not contain release binaries, so this installation will not work. Please rewrite for your repository. See the "Building the plugin" section to get this template ruleset working.

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
 enabled = true

    version = "0.0.3"
    source = "github.com/geekette86/tflint-ruleset-labels"
    signing_key = <<-KEY
    -----BEGIN PGP PUBLIC KEY BLOCK-----

    mQINBGOheScBEADD71zhHOLxpIq6haF4AVrDB9Bpd5e+O4h3MX4dzaeaUf+UHYAx
    jq5PfOnXfiWq6QaQV+KQrB3Tw3HkvTqiZaUS3RCf71hyaY+UQFc9ByxUNIm6gWse
    PEN9EDkC/5bGAW1aU56Vfv5PObnvHbRQ9n5/wJksO6heRRU9aoWnfYiVtnXbbaip
    AEPNemeK5NbCDQqi4QZO+n98lLZP/j++lTzVllIX0WG1oRfxjZ+PoD575kUYvBzo
    8aCWAZS/McUyXrktM7/2CoVg/A7buEGfaN4HRHZMve58WmJOtIgjtNnlWf9yg9WY
    ..Contact me for the rest
    -----END PGP PUBLIC KEY BLOCK-----
    KEY
}
```

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "template" {
  enabled = true
}
EOS
$ tflint
```
