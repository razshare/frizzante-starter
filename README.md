# Get started

Configure the project
```bash
make configure
```

> [!NOTE]
> This will install [Bun](https://bun.sh).\
> If you'd rather use a different runtime see [makefile](./makefile), section `configure`.

Load dependencies

```bash
make load
```

And finally, either start or build your project

```bash
make start
```

```bash
make build
```

> [!NOTE]
> The `ui` directory is not embedded into the final executable.