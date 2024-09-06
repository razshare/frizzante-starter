# Frizzante Example

This is an example of how to use [frizzante](https://github.com/razshare/frizzante) to render [Svelte](https://svelte.dev/) components using V8 bindings.

1. Enable CGO
   ```sh
   go env -w CGO_ENABLED='1'
   ```

2. Install your dependencies.
   ```sh
   npm i && \\
   go mod tidy
   ```
3. Run the server.
   ```sh
   go run main.go
   ```
4. Visit http://127.0.0.1:8080
