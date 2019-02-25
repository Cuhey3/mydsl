const {
  Argument,
  dslFunctions,
  dslDefinedFunctions
} = require('../lib/dsl-core.js');

const dsl_define_func = {};

dsl_define_func.define = function(arg1, arg2) {
  const functionName = arg1.evaluate(this);
  if (!(functionName in dslDefinedFunctions)) {
    return dslDefinedFunctions[functionName] = function() {
      return arg2.evaluate(this);
    };
  }
};

dsl_define_func.timer = function(arg1, arg2) {
  const _self = JSON.parse(JSON.stringify(this));
  const period = arg1.evaluate(this);
  setInterval(function() {
    arg2.evaluate(_self);
  }, period);
};

function addAll() {
  Object.assign(dslFunctions, dsl_define_func);
}

addAll();
module.exports = { Argument };
