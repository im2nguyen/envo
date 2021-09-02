## envo - Mask your environment variables

`envo` (environment variable obscuration) is a CLI tool that masks/obscurates your environment variables.

### Use case

You're in the middle of a demo and want to show the environment variables you 
set for the demo, without _showing_ your actual environment variables.

For example, you're running a Terraform demo and you want to show that you set 
your AWS credentials as environment variables.

Typically, you would run the `env | grep AWS`, but now your AWS credentials are 
exposed to everyone. You can get around by creating temporary AWS credentials, 
then revoking it, but this is also a hassle.

With `envo`, you can mask these values.

```
$ envo | grep AWS
AWS_SESSION_TOKEN=IQoJb3JpZ2luX2VjEMf//////////wEaCXVzLWVhc3QtMSJIMEYCIQCmfFtM4rtTmuk5yEBsY5rmy1hmRKp7yH3YRCyum7ACDQIhAIjrHzOpv+byWtSCfjpPoRaajzUS+yn05hDe8BY588RbKu4ECBAQARoMNTYxNjU2OTgwMTU5IgyezRiwDbMoMtHp5yYqywRau7B5fQ2COWvwrB0cQgS9Exy60Gg18sdxiSJwIFSv2lwcmVwV7XAXwUWm58MXkeQh8QDCT+qlk6lWbvOt0LI4bo4GZeqlAKkn95dMefGatI+X3JtcG1gj/mOLAlBtRMReih31sZBxxakbrvC7VcQC8vt+mQ79X+0J6Bftnp7dp4/YjkTl8OXegbQ9b/TJpypw5C9tPO3QCzbH0...
AWS_SESSION_EXPIRATION=-621...
AWS_SECRET_ACCESS_KEY=hOFE6NnaPqdAa...
AWS_ACCESS_KEY_ID=ASIAYF...
```

### Installation

To install on MacOS, run the following commands.

```
$ brew tap im2nguyen/envo
$ brew install envo
```

### Basic usage

`envo` has two main flags, `-maskMethod` (`-m`) and `-truncLength` (`-t`). 
If `-maskMethod` is not specified, `envo` will truncate the environment variable 
values.

- Set no flags to truncate environment variable value, defaults to a third of 
value length.

    ```
    $ envo | grep USER
    USER=d...
    ```

- Set `-maskMethod` to `random` to replace the environment variable value with 
random characters.

    ```
    $ envo -m=random | grep USER
    USER=me8
    ```

- Set `-maskMethod` to a value that is neither `trunc` nor `random` to replace 
the environment variable value with the provided value.

    ```
    $ envo -m=* | grep USER
    USER=***
    ```

    ```
    $ envo -m=0 | grep USER
    USER=000
    ```

- Set `-truncate` to a length to truncate the environment variable value. If the
specified length is greater than or equal to the length of the value, `envo` 
will truncate the value to a third of its original length.

    ```
    $ envo -t=1 | grep USER
    USER=d...
    ```

    ```
    $ envo -m=3 | grep USER
    USER=d...
    ```


