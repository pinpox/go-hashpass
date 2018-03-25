# go-hashpass
A hashpass (https://github.com/stepchowfun/hashpass) compatible password manager

Taken from the Hashpass Readme:



> Hashpass is a Chrome extension designed to make passwords less painful. It generates a unique password for every website you use, and you only have to memorize a single secret key.
>
> Hashpass is deterministic, meaning that it will always generate the same password for any given site and secret key. It uses a well-known formula to generate the passwords, so you could even compute them yourself.
>
> A key feature of Hashpass is that it's stateless. Hashpass never writes to the file system or makes network requests. There is no password database.

This is just a client that implements the used algorithm, yielding the same results. It is implemented as CLI tool and can be used without any browser.
