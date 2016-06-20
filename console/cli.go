package console

import "../model"
import "strings"

const (
  TermSeparation = " "
  CommandTypeDescriptionNew = "new"
CommandTypeDescriptionSet = "set"
CommandTypeDescriptionDelete = "delete"
)

func GetArgumentTree() model.ArgumentTree {
  var created model.ArgumentTree
  return created
}

func CreateArgumentTree(line string) model.ArgumentTree {
  var created model.ArgumentTree
  created.CommandPrincipal = GetCommandTypeInt(GetFirstCommandTerm(line))
  return created
}

func GetFirstCommandTerm(line string) string {

  partial := strings.Split(line, TermSeparation)

    if len(partial) <= 0 {
      return ""
    }

  return partial[0]
}

func GetCommandTypeInt(line string) int {

if line == CommandTypeDescriptionNew {
  return model.CommandTypeNew
}

if line == CommandTypeDescriptionSet {
  return model.CommandTypeSet
}

if line == CommandTypeDescriptionDelete {
  return model.CommandTypeDelete
}
  return model.CommandTypeNotSpecified
}
