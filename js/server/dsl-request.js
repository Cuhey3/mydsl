const { Argument, dslFunctions } = require('../lib/dsl-core.js');

const request = require('request-promise');
const { JSDOM } = require('jsdom');
const dsl_request_func = {};
const hash = require('object-hash');

dsl_request_func.request = async function(arg1, arg2, arg3) {
  if (arg1.toString() === 'get') {
    const url = arg2.evaluate(this);
    const result = await request.get(url);
    if (arg3 && arg3.toString() === 'json') {
      return JSON.parse(result);
    }
    else {
      return result;
    }
  }
  else if (arg1.toString() === 'post') {
    // TBD
    // return await request.post(arg2.evaluate(this), arg3.evaluate(this));
  }
};

dsl_request_func.toDocument = function(arg1) {
  return new JSDOM(arg1.evaluate(this)).window.document;
};

dsl_request_func.selectOne = function(arg1, arg2) {
  return arg1.evaluate(this).querySelector(arg2.evaluate(this));
};

dsl_request_func.selectAll = function(arg1, arg2) {
  return Array.from(arg1.evaluate(this).querySelectorAll(arg2.evaluate(this)));
};

{
  const entryKind = {};
  dsl_request_func.isNewEntry = function(arg1, arg2, arg3) {
    const kind = arg1.evaluate(this);
    const entry = arg2.evaluate(this);
    const kindSize = arg3.evaluate(this);
    if (!(kind in entryKind)) {
      entryKind[kind] = { map: {}, list: [] };
    }
    const { map, list } = entryKind[kind];
    const _hash = hash(entry);
    const isNew = !(_hash in map);
    if (isNew) {
      map[_hash] = '';
      list.push(_hash);
      while (list.length > kindSize) {
        delete map[list.shift()];
      }
    }
    return isNew;
  };
}

{
  const throttles = {};

  dsl_request_func.semaphoreQueue = function(arg1, arg2, arg3) {
    const throttleName = arg1.evaluate(this);
    const semaphoreSize = arg2.evaluate(this);
    if (!(throttleName in throttles)) {
      throttles[throttleName] = {
        queue: [],
        promises: {},
        consumerRunning: false
      };
      throttles[throttleName].consume = function() {
        if (throttles[throttleName].consumerRunning === false) {
          const { promises, queue } = throttles[throttleName];
          while (Object.keys(promises).length < semaphoreSize &&
            queue.length > 0) {
            const key = new Date().getTime().toString();
            promises[key] = (async function() {
              const [arg, copied] = queue.shift();
              await arg.evaluate(copied);
              return key;
            })();
          }
          if (Object.keys(promises).length > 0) {
            throttles[throttleName].consumerRunning = true;
            Promise.race(Object.keys(promises).map(function(key) {
              return promises[key];
            })).then(function(key) {
              delete promises[key];
            }).finally(function() {
              throttles[throttleName].consumerRunning = false;
              throttles[throttleName].consume();
            });
          }
        }
      };
    }
    const copied = JSON.parse(JSON.stringify(this));
    throttles[throttleName].queue.push([arg3, copied]);
    throttles[throttleName].consume();
  };
}

function addAll() {
  Object.assign(dslFunctions, dsl_request_func);
}

addAll();
module.exports = { Argument };
