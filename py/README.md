## mydsl for Python

mydsl is yaml-based DSL library for JavaScript/Node.js/Go/Python.

mydsl can replace your code with a YAML file.

### Install mydsl via pip

```
pip install -e git+https://github.com/cuhey3/mydsl.git#egg=mydsl\&subdirectory=py
```

### Usage

```
from mydsl import dsl_core, dsl_mongo, dsl_server
import yaml

print(dsl_mongo.loadDslFunctions)
dsl_mongo.loadDslFunctions(dsl_core.dslFunctions)
dsl_server.loadDslFunctions(dsl_core.dslFunctions, dsl_core.dslAvailableFunctions)

with open('router.yml') as yaml_file:
  dsl = yaml.load(yaml_file)
  router, err = dsl_core.Argument(dsl).evaluate({})
```

If you use dsl_mongo, you need a MONGODB_URI set of environment variables.

```
export MONGODB_URI=mongodb://[username:password@]host1[:port1][,...hostN[:portN]]][/[database][?options]]
```
