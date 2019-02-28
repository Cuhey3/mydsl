process.env.TZ = 'Asia/Tokyo';

const express = require('express');
const path = require('path');
const PORT = process.env.PORT || 5000;
const bodyParser = require('body-parser');
var nunjucks = require('nunjucks');

function setRoutes(app) {
  app.set('view engine', 'html');
  nunjucks.configure('views', {
    autoscape: true,
    express: app,
    watch: true
  });
  app
    .use('/public', express.static(path.join(__dirname, 'public')))
    .use('/lib', express.static(path.join(__dirname, 'lib')))
    .use(bodyParser.urlencoded({
      extended: true
    }))
    .listen(PORT, () => console.log(`Listening on ${ PORT }`));
}

setRoutes(express());
