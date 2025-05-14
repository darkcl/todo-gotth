# templUI todo list

Trying to make sense of the folder structure

## Usage

Run `make prepare` to create `gopath.css` or all the style will be fucked

## Folder structure

> Unsure will be marked with `(?)`

- assets, static files
- components
  - UI library from templUI
- moduels
  - core logic - each module is a separate use case
    - {{module name}}
      - {{module}}.go - interfaces
      - {{module}}_controller.go - controller
      - {{module}}_service.go - service
      - {{module}}_routes.go - routes
      - {{usecase}}_view.templ - view template (?)
      - {{module}}_page.templ - page template, expected to be the whole page of the module
- routes
  - api_routes.go - api routes, point to specific core logic, e.g. `/todos`
  - assets_routes.go - static assets, e.g. `/assets`, handle live mode / embdded mode here
  - pages_routes.go - page routes, point to specific page, e.g. `/` will show the landing page
- utils, common helper functions
