var a = {}
var b = Object.prototype
console.log(a.prototype === b)
console.log(Object.getPrototypeOf(a) == b)