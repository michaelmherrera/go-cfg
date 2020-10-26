# Go CFG generator

> Creates phrases using context-free grammars

# Usage

```
go-cfg [N] [SEED]

Generate N pieces of text with the random-number generator seeded with SEED. 

By default, 
    N = 1
    SEED = time.Now().UnixNano().

```

# Grammar and template files

- Generates rules via the grammar file `bng-cfg.txt` which is in [Backus-Naur Form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form)
- Generates text based on the template defined in `template.txt`
- You can modify `bnf-cfg.txt` and `template.txt` with your own grammar and template. For demonstration purposes, the grammar and template are configured to so the generator produces fake addresses.

