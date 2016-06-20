package model

const (
CommandTypeNotSpecified = iota
  CommandTypeNew = iota
  CommandTypeSet = iota
  CommandTypeDelete = iota
)
  type ArgumentTree struct {
  Trunk string
  Branch string
  Parameter string
  CommandPrincipal int
  CommandSecondary int
}
