### Path Params

```zsh
curl localhost:8080/api/users/:pathParam
```

Normalmente utilizamos quando precisamos identificar recursos expecificos

**Exemplos:**

Buscando usuario com id `1`
rota: /api/users/:id

```zsh
curl localhost:8080/api/users/1
```

Buscando produtos da categoria  `eletronicos`
rota: /api/products/:category

```zsh
curl localhost:8080/api/products/eletronicos
```

### Query Params

usado para filtros, paginacao e ordenacao

// filtro
?category=eletronicos&name=mo

// paginacao
?page=2&itemsPerPage=10

// ordenacao
?sortBy=name&sortOrder=desc