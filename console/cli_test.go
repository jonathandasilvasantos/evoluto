package console
import "testing"
import "reflect"
import "../model"

func TestGetArgumentTreeTree(t *testing.T) {



  var argumentTree model.ArgumentTree
  expected := reflect.TypeOf(argumentTree).Kind()

  actual := GetArgumentTree()

  if expected != reflect.TypeOf(actual).Kind() {
    t.Errorf("Test Get Argument Tree expected '%s' got: '%s'", expected, actual);
  }
}
