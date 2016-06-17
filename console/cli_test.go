package console
import "testing"

func TestGetArgumentTreeTree(t *testing.T) {



  expected := "No command"
  actual := GetArgumentTree()

  if expected != actual {
    t.Errorf("Test Get Argument Tree expected '%s' got: '%s'", expected, actual);
  }
}
