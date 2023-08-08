package valuex

func StrPtr(v string) *string {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func IntPtr(v int) *int {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

func NotZeroIntPtr(v *int) bool {
	return v != nil && *v != 0
}

func NotZeroInt32Ptr(v *int32) bool {
	return v != nil && *v != 0
}

func NotZeroInt64Ptr(v *int64) bool {
	return v != nil && *v != 0
}
