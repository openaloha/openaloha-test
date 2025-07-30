package syncfunc

type InitFunc func() error
type RefreshFunc func() error