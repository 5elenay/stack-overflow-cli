# Stack Overflow Cli (Stovc)
**St**ack **Ov**erflow **C**li is a CLI tool that helps you to search a question in stack overflow in your terminal.

## Installation
- Download binary for your system from: [latest release](https://github.com/5elenay/stack-overflow-cli/releases/latest)
- Add to the path and done.

## Example
Run this example search command in your favorite terminal:
```bash
stovc search "postgresql" --sort=votes
```

You should get a result like:
```bash
Question #1
  Title: "PostgreSQL: Show tables in PostgreSQL"
  Tags: [ postgresql, ]
  Visit the Page? [y/n]:
```

if you type `y` or `yes` and enter, It will open the post url in your browser. If not, it will skip to the next result.

Also you can use `help` command for getting help about commands, options etc...

## License
This project is licensed under GPLv3 license. [See](https://raw.githubusercontent.com/5elenay/stack-overflow-cli/main/LICENSE)