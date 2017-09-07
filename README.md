# invoice

A dead simple way to edit invoices, mainly for freelancers.

*Comes bundled for French invoicing laws, contributions are welcome for localization!*

## Example

1. Source [header](example/header.txt) and [footer](example/footer.txt)
2. Source [YAML file](example/000001.yml)
3. And... [voilà! 🤑](example/000001.pdf)

## 1-minute setup

```bash
go get gitlab.com/Lesterpig/invoice

# Edit header and footer
vim header.txt
vim footer.txt

# Edit first invoice
invoice new
vim 000001.yml
invoice 000001.yml 0000001.pdf
evince 000001.yml

# Edit second invoice
invoice new
vim 000002.yml
...
```
