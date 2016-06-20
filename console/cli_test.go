package console
import "testing"
import "reflect"
import "../model"

func TestGetCommandTypeInt(t *testing.T) {
  actual := GetCommandTypeInt(CommandTypeDescriptionNew)
  expected := model.CommandTypeNew

if expected != actual {
  t.Errorf("Test Get Command Type Int expected: %d got: %d", expected, actual)
}

actual = GetCommandTypeInt(CommandTypeDescriptionDelete)
expected = model.CommandTypeDelete

if expected != actual {
t.Errorf("Test Get Command Type Int expected: %d got: %d", expected, actual)
}

}

func TestGetFirstCommandTerm(t *testing.T) {
  actual := GetFirstCommandTerm("new project")
  expected := "new"

  if actual != expected {
    t.Errorf("Test Get First Command Term expected: '%s' got: '%s'", expected, actual);
  }
}

func TestCreateArgumentTree(t *testing.T) {

  var argumentTree model.ArgumentTree
  typeExpected := reflect.TypeOf(argumentTree).Kind()
  actual := CreateArgumentTree("")

  if typeExpected != reflect.TypeOf(actual).Kind() {
    t.Errorf("Test Create Argument Tree expected '%s' got: '%s'", typeExpected, actual);
  }

  var expected model.ArgumentTree
  expected.CommandPrincipal = model.CommandTypeNew
  actual = CreateArgumentTree("new project")

  if expected.CommandPrincipal != actual.CommandPrincipal {
    t.Errorf("Test Create Argument Tree Command Principal expected '%d' got: '%d'", expected.CommandPrincipal, actual.CommandPrincipal);
  }

  actual = CreateArgumentTree("delete project")
  expected.CommandPrincipal = model.CommandTypeDelete

  if expected.CommandPrincipal != actual.CommandPrincipal {
    t.Errorf("Test Create Argument Tree Command Principal expected '%d' got: '%d'", expected.CommandPrincipal, actual.CommandPrincipal);
  }








}

func TestGetArgumentTree(t *testing.T) {

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

  if actual.CommandPrincipal != 0 {
  t.Errorf("Test Get Argument Tree expected command principal '0' got '%d'", actual.CommandPrincipal)
  }

  if actual.CommandSecondary != 0 {
  t.Errorf("Test Get Argument Tree expected command secondary '0' got '%d'", actual.CommandSecondary)
  }
}
