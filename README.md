A small CLI tool for [Shamir's Secret Sharing](https://en.wikipedia.org/wiki/Shamir%27s_Secret_Sharing)
written in Go, using [Vault's](https://github.com/hashicorp/vault) Shamir
implementation.

![Go](https://github.com/kinvolk/go-shamir/workflows/Go/badge.svg)

## Usage

Split secret:

```
$ echo -n "very very secret" | ./bin/shamir split -p 4 -t 2
baa3e1b656d6b253052d293b99daf7fa4a
07cfbaa1bf6982413dd52abb2578ca6373
c9cc6036850debccca9dd598bebf27acd1
db7b57989fb3d27775c62f20fa858dd338
```

Combine secret:

```
$ cat <<EOF | ./bin/shamir combine
> 07cfbaa1bf6982413dd52abb2578ca6373
> c9cc6036850debccca9dd598bebf27acd1
> EOF
very very secret
```
