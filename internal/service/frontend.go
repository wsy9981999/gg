// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IFrontend interface{}
)

var (
	localFrontend IFrontend
)

func Frontend() IFrontend {
	if localFrontend == nil {
		panic("implement not found for interface IFrontend, forgot register?")
	}
	return localFrontend
}

func RegisterFrontend(i IFrontend) {
	localFrontend = i
}
