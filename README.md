## CoalIDE

CoalIDE provides an online application that can be used for writing, compiling, and building console applications. Some of the features include:

- Uncluttered and visually appealing user interface
- Powerful text editor with syntax highlighting and code folding
- Builds take place in a container isolated from the host

### Requirements

CoalIDE consists of a single executable and some static files. The executable uses Docker for building and running the applications while ensuring complete isolation from the host. In order to do this, the executable needs to be able to interact with Docker through its remote API.

### Setup

[TODO]

### Development

The application is written in Go and the client-side JavaScript code uses [Bootstrap](http://getbootstrap.com/) and [Ember.js](http://emberjs.com/).
