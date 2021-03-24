import go
import TestUtilities.InlineExpectationsTest

class Sink extends DataFlow::Node {
  Sink() {
    exists(DataFlow::CallNode c | c.getTarget().getName() = "sink" | this = c.getAnArgument())
  }
}

class Source extends DataFlow::Node {
  Source() {
    exists(DataFlow::CallNode c | c.getTarget().getName() = "source" | this = c.getResult())
  }
}

class TestConfig extends DataFlow::Configuration {
  TestConfig() { this = "testconfig" }

  override predicate isSource(DataFlow::Node source) { source instanceof Source }

  override predicate isSink(DataFlow::Node sink) { sink instanceof Sink }
}

class DataFlowTest extends InlineExpectationsTest {
  DataFlowTest() { this = "DataFlowTest" }

  override string getARelevantTag() { result = "dataflow" }

  override predicate hasActualResult(string file, int line, string element, string tag, string value) {
    tag = "dataflow" and
    exists(DataFlow::Node source, DataFlow::Node sink | any(TestConfig c).hasFlow(source, sink) |
      element = sink.toString() and
      value = source.getStartLine().toString() and
      sink.hasLocationInfo(file, line, _, _, _)
    )
  }
}
