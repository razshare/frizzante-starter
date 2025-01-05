# Get started

> [!NOTE]
> Make sure you have [Go installed and visible in your path](https://go.dev/doc/install).

> [!NOTE]
> This project uses the [Bun](https://bun.sh) runtime for updating javascript dependencies,
> so make sure Bun is also installed.\
> If you'd rather use a different runtime see [makefile, section "update"](https://github.com/razshare/frizzante-starter/blob/master/makefile#L1-L6).

> [!NOTE]
> Bun itself is not required for your application to run,
> it's only required when building your application with Vite.\

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
