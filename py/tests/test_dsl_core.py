def test_all():
  from mydsl import dsl_core, dsl_mongo, dsl_server
  import yaml
  
  dsl_mongo.loadDslFunctions(dsl_core.dslFunctions)
  dsl_server.loadDslFunctions(dsl_core.dslFunctions, dsl_core.dslAvailableFunctions)
  
  with open('tests/testsuite.yml') as yaml_file:
    testsuites = yaml.safe_load(yaml_file)
    container = {}
    for testsuite in testsuites:
      evaluated, err = dsl_core.Argument(testsuite).evaluate(container)
      assert err == None and evaluated == None, "testsuite {} failed: {}".format(str(testsuite["testsuite"][0]), evaluated)
