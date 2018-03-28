# [DTD] Vue.js web app

HTML + CSS + JavaScript sources of a web-based GUI to provide a responsive and customizable layout for visualization (frontend).

## Development

After the development environment is initialized, the frontend can be tested with:

```
cd app
yarn serve
```

In order to test the full app (i.e., both frontend and backend), first build the frontend:

```
yarn build
```

Then, start the backend:

```
cd ../api
go run main.go -v -p 8080 -d ../app/dist
```

---

- https://github.com/vuejs/vue-loader
  - https://vue-loader.vuejs.org/en/
- https://github.com/tholman/github-corners

---

- https://github.com/surmon-china/vue-codemirror
- https://github.com/mauricius/vue-draggable-resizable

---

- https://github.com/vuejs/awesome-vue
- https://vuematerial.io/
- https://vulma.org/
