*Psst — looking for a shareable component template? Go here --> [sveltejs/component-template](https://github.com/sveltejs/component-template)*

---

# svelte app

This is a fork of the project template for [Svelte](https://svelte.dev) apps. It lives at https://github.com/sveltejs/template.
I updated it to work with chokidar and to run the compiler/server in docker.

You don't even need to install nodejs to get this baby going if you have docker.

Run this to make a directory for your project and get everything it needs.

`docker run --rm -it -v "%CD%":/app node /bin/bash -c "cd /app && npx degit ScienceVikings/svelte-template my-svelte-project && cd my-svelte-project && npm install"`

Note: If you're running in a Linux-y environment change the `%CD%` to `$PWD`

Change directories into the `my-svelte-project` directory.

Finally, gettergoin with a lil `docker-compose up`
