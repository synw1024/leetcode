var name = 'World!';
(function () {
  console.log(typeof name === 'undefined')
    if (typeof name === 'undefined') {
      var name = 'Jack';
      console.log('Goodbye ' + name);
    } else {
      console.log('Hello ' + name);
    }
})();