# touchy

  little touch tool with gist support.

  really easy way to scaffold new views, models, controllers, components, etc.

# example

to create a component scaffolding:

```bash
mkdir -p components/{a,b,c}
ls components/ | xargs -I % touchy https://gist.github.com/matthewmueller/f0fc7637e1ff426c89dee977fbe612df components/%/{index.js,index.css}
```

Then find and replace all `__NAME__` with the name of the component

# installation

```bash
go get github.com/matthewmueller/touchy
```

# like-to-haves

- look for any __VARIABLE__ or {{VARIABLE}} and ask what you want to put there.
- better errors

# license

MIT
