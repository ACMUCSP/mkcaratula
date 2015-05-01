# UCSP cover pages generator

At [Universidad Católica San Pablo](http://ucsp.edu.pe) there is a standard cover page for papers. This webapp generates these cover
pages from a form.

## Colaborators

- [Aldo Culquicondor](https://github.com/alculquicondor) (maintainer)
- [Juan Salas](https://github.com/ratasxy)
- [Javier Quinte](https://github.com/jaqus)

## Development

### Requirements

- TeX Live installed on the server
- composer
- PHP 5.3+

### Install & Run

1. Install composer

        curl -sS https://getcomposer.org/installer | php

2. Run composer

        php composer.phar update

3. Create & edit configuration files

        cp config/config.yml.dist config/config.yml
        cp web/js/config.js.dist web/js/config.js

4. Create directory for pdfs

        mkdir web/s

4. Run server

        php -S localhost:8000 -t web
