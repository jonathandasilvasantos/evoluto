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

  if actual.Trunk != "" {
    t.Errorf("Test Get Argument Tree expected trunk '' got '%s'", actual.Trunk)
  }

if actual.Branch != "" {
  t.Errorf("Test Get Argument Tree expected branch '' got '%s'", actual.Branch)
}

if actual.Parameter != "" {
  t.Errorf("Test Get Argument Tree expected parameter '' got '%s'", actual.Parameter)
}

}
