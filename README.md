# Get started

> [!NOTE]
> This project uses the [Bun](https://bun.sh) runtime for updating dependencies.\
> If you'd rather use a different runtime see [makefile, section "update"](./makefile#L1-L6).

Update dependencies

```bash
make update
```

Then start the server

```bash
make start
```

or build it

```bash
make build
```

> [!NOTE]
> The `www/dist` directory is embedded, which makes the final executable completely portable.
>
> That being said, you can still create a "www/dist" directory near your executable.\
> Whenever a request is trying to access a file missing from the embedded file system, the server will fall
> back to the nearby "www/dist" directory instead.
